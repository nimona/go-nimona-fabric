package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/importer"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
	"nimona.io/pkg/object"
)

var (
	schema   = flag.String("schema", "", "schema for struct")
	input    = flag.String("in", "", "input file")
	output   = flag.String("out", "-", "output file (default is stdout)")
	typename = flag.String("type", "", "type to generate methods for")
)

func init() {
	flag.Parse()
}

func main() {
	gen := Generator{
		InputFile: *input,
		Type:      *typename,
	}

	code, err := gen.process()
	if err != nil {
		log.Fatal(err)
	}

	if *output == "-" {
		os.Stdout.Write(code)
	} else if err := ioutil.WriteFile(*output, code, 0644); err != nil {
		log.Fatal(err)
	}
}

type Generator struct {
	InputFile string
	Type      string
	Importer  types.Importer
	FileSet   *token.FileSet
}

type Values struct {
	Package      string
	StructName   string
	StructFields []*Field
	Schema       string
	Imports      map[string]bool
}

type Field struct {
	Skip        bool
	Name        string
	Tag         string
	TypePtr     string
	ElemTypePtr string
	Type        string
	ElemType    string
	Hint        string
	IsBasic     bool
	IsObject    bool
	IsSlice     bool
	CanBeNil    bool
}

func (gen *Generator) process() (code []byte, err error) {
	if gen.FileSet == nil {
		gen.FileSet = token.NewFileSet()
	}
	if gen.Importer == nil {
		gen.Importer = importer.Default()
	}
	pkg, err := gen.loadPackage()
	if err != nil {
		return nil, err
	}

	typ, err := lookupStructType(pkg.Types.Scope(), gen.Type)
	if err != nil {
		return nil, fmt.Errorf("can't find %s in %q: %v", gen.Type, pkg.PkgPath, err)
	}

	values := &Values{
		Package:      pkg.Name,
		StructName:   gen.Type,
		StructFields: []*Field{},
		Schema:       *schema,
		Imports:      map[string]bool{},
	}

	if pkg.Name != "encoding" {
		values.Imports["nimona.io/pkg/object"] = true
	}

	fmt.Printf("Objectifying %s.%s\n", values.Package, gen.Type)

	styp := typ.Underlying().(*types.Struct)

	// obj := pkg.Scope().Lookup(gen.Type)
	// ptr:=types.NewPointer(obj.Type())
	// imp:=types.Implements(obj.Name(), ifff)

	for i := 0; i < styp.NumFields(); i++ {
		f := styp.Field(i)
		if !f.Exported() {
			continue
		}
		if f.Anonymous() {
			fmt.Fprintf(os.Stderr, "Warning: ignoring embedded field %s\n", f.Name())
			continue
		}

		tag := reflect.StructTag(styp.Tag(i))
		ftag := tag.Get("fluffy")
		if ftag == "" {
			// TODO should we fallback to json?
			ftag = tag.Get("json")
		}
		vf := getMetaFromTag(ftag)
		if vf == nil {
			// no tags found
			vf = &Field{
				Tag: toLowerFirst(f.Name()),
			}
		}
		if vf.Skip {
			continue
		}
		vf.Name = f.Name()
		// vf.TypePtr = removePackageFromTypePtr(f.Type().String(), pkg.PkgPath, f.Pkg().Name())
		// vf.Type = removePackageFromType(f.Type().String())

		tp, tpkg := getPackageAndType(f.Type().String(), pkg.PkgPath, false)
		vf.TypePtr = tp

		if tpkg != "" {
			values.Imports[tpkg] = true
		}

		tp, _ = getPackageAndType(f.Type().String(), pkg.PkgPath, true)
		vf.Type = tp

		// vf.Type = removePackageFromType(f.Type().String())

		hint := getHint(f.Type())
		if vf.Hint == "" {
			vf.Hint = hint
		} else if vf.Hint != hint {
			panic(fmt.Errorf("existing hint of %s for field %s does not match infered %s", vf.Hint, vf.Name, hint))
		}

		if strings.Contains(vf.Hint, "o") {
			vf.IsObject = true
		}

		if vf.TypePtr[0] == '*' {
			vf.CanBeNil = true
		}

		if _, ok := f.Type().(*types.Map); ok {
			vf.CanBeNil = true
		}

		if fi, ok := f.Type().(*types.Slice); ok {
			vf.IsSlice = true
			vf.CanBeNil = true

			etp, _ := getPackageAndType(fi.Elem().String(), pkg.PkgPath, false)
			vf.ElemType = etp

			etp, _ = getPackageAndType(fi.Elem().String(), pkg.PkgPath, true)
			vf.ElemTypePtr = tp

			if _, ok := fi.Elem().(*types.Basic); ok {
				vf.IsBasic = true
			}

			if vf.IsSlice {
				vf.Type = "[]" + vf.Type
				vf.TypePtr = "[]" + vf.TypePtr
			}

		}

		values.StructFields = append(values.StructFields, vf)

		fmt.Printf("  - field=%s; tag=%s, type=%s, hint=%s, skipping=%t\n", vf.Name, vf.Tag, vf.Type, vf.Hint, vf.Skip)
	}

	tpl := `// Code generated by nimona.io/tools/objectify. DO NOT EDIT.

// +build !generate

package {{ .Package }}

import (
	"fmt"

	{{- range $pkg, $ok := .Imports }}
	"{{ $pkg }}"
	{{- end }}
)

// ToMap returns a map compatible with f12n
func (s {{ .StructName }}) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"@ctx:s": "{{ .Schema }}",
		{{- range .StructFields }}
		{{- if eq .Tag "@" }}
		{{- else if .CanBeNil }}
		{{- else if .IsObject }}
		"{{ .Tag }}:{{ .Hint }}": s.{{ .Name }}.ToMap(),
		{{- else }}
		"{{ .Tag }}:{{ .Hint }}": s.{{ .Name }},
		{{- end }}
		{{- end }}
	}
	{{- range .StructFields }}
	{{- if eq .Tag "@" }}
	{{- else if .CanBeNil }}
	if s.{{ .Name }} != nil {
		{{- if and .IsObject .IsSlice }}
		s{{ .Name }} := []map[string]interface{}{}
		for _, v := range s.{{ .Name }} {
			s{{ .Name }} = append(s{{ .Name }}, v.ToMap())
		}
		m["{{ .Tag }}:{{ .Hint }}"] = s{{ .Name }}
		{{- else if .IsObject }}
		m["{{ .Tag }}:{{ .Hint }}"] = s.{{ .Name }}.ToMap()
		{{- else }}
		m["{{ .Tag }}:{{ .Hint }}"] = s.{{ .Name }}
		{{- end }}
	}
	{{- end }}
	{{- end }}
	return m
}

// ToObject returns a f12n object
func (s {{ .StructName }}) ToObject() *object.Object {
	return object.FromMap(s.ToMap())
}

// FromMap populates the struct from a f12n compatible map
func (s *{{ .StructName }}) FromMap(m map[string]interface{}) error {
	{{- range .StructFields }}
	{{- if eq .Tag "@" }}
	s.{{ .Name }} = object.FromMap(m)
	{{- else if .IsSlice }}
	s.{{ .Name }} = []{{ .ElemType }}{}
	if ss, ok := m["{{ .Tag }}:{{ .Hint }}"].([]interface{}); ok {
		for _, si := range ss {
			if v, ok := si.({{ .ElemType }}); ok {
				s.{{ .Name }} = append(s.{{ .Name }}, v)
			}
			{{- if not .IsBasic }} else if v, ok := si.(map[string]interface{}); ok {
				s{{ .Name }} := {{ .ElemTypePtr }}{}
				if err := s{{ .Name }}.FromMap(v); err != nil {
					return err
				}
				s.{{ .Name }} = append(s.{{ .Name }}, s{{ .Name }})
			}
			{{- end }}
		}
	}
	{{- else if .IsObject }}
	if v, ok := m["{{ .Tag }}:{{ .Hint }}"].(map[string]interface{}); ok {
		s.{{ .Name }} = {{ .Type }}{}
		if err := s.{{ .Name }}.FromMap(v); err != nil {
			return err
		}
	} else if v, ok := m["{{ .Tag }}:{{ .Hint }}"].({{ .TypePtr }}); ok {
		s.{{ .Name }} = v
	}
	{{- end }}
	if v, ok := m["{{ .Tag }}:{{ .Hint }}"].({{ .TypePtr }}); ok {
		s.{{ .Name }} = v
	}
	{{- end }}
	return nil
}

// FromObject populates the struct from a f12n object
func (s *{{ .StructName }}) FromObject(o *object.Object) error {
	return s.FromMap(o.ToMap())
}

// GetType returns the object's type
func (s {{ .StructName }}) GetType() string {
	return "{{ .Schema }}"
}`

	f, err := os.Create(*output)
	if err != nil {
		panic(err)
	}

	t, err := template.New("t").Parse(tpl)
	if err != nil {
		panic(err)
	}

	out := bytes.NewBuffer([]byte{})
	if err := t.Execute(out, values); err != nil {
		panic(err)
	}

	if values.Package == "encoding" {
		sout := strings.Replace(string(out.Bytes()), "object.", "", -1)
		out = bytes.NewBuffer([]byte(sout))
	}

	if _, err := f.Write(out.Bytes()); err != nil {
		panic(err)
	}

	opt := &imports.Options{
		Comments:  true,
		TabIndent: true,
		TabWidth:  8,
	}
	code, err = imports.Process(*output, code, opt)
	if err != nil {
		panic(fmt.Errorf("BUG: can't gofmt generated code: %v", err))
	}
	return code, nil
}

func (gen *Generator) loadPackage() (*packages.Package, error) {
	pattern := "file=" + gen.InputFile
	lcfg := &packages.Config{
		Mode: packages.LoadAllSyntax,
		Fset: gen.FileSet,
		Env:  os.Environ(),
	}
	pkgs, err := packages.Load(lcfg, pattern)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, errors.New("no files found")
	}
	return pkgs[0], nil
}

func lookupStructType(scope *types.Scope, name string) (*types.Named, error) {
	typ, err := lookupType(scope, name)
	if err != nil {
		return nil, err
	}
	_, ok := typ.Underlying().(*types.Struct)
	if !ok {
		return nil, errors.New("not a struct type")
	}
	return typ, nil
}

func lookupType(scope *types.Scope, name string) (*types.Named, error) {
	obj := scope.Lookup(name)
	if obj == nil {
		return nil, errors.New("no such identifier")
	}
	typ, ok := obj.(*types.TypeName)
	if !ok {
		return nil, errors.New("not a type")
	}
	return typ.Type().(*types.Named), nil
}

func getMetaFromTag(tag string) *Field {
	if tag == "" {
		return nil
	}

	args := strings.Split(tag, ",")

	vf := &Field{
		Tag: args[0],
	}

	tp := strings.Split(vf.Tag, ":")
	if len(tp) > 1 {
		vf.Tag = tp[0]
		vf.Hint = tp[1]
	}

	for _, t := range args {
		switch t {
		case "object":
			vf.IsObject = true
		}
	}

	if vf.Tag == "-" {
		vf.Skip = true
	}

	return vf
}

func getPackageAndType(t, pkg string, deref bool) (string, string) {
	t = strings.Replace(t, "[]", "", 1)
	ptr := false
	if t[0] == '*' {
		ptr = true
		t = t[1:]
	}

	ct := strings.Replace(t, pkg, "", 1)
	ts := strings.Split(ct, ".")
	tpkg := strings.Join(ts[:len(ts)-1], ".")
	ts = strings.Split(ct, "/")
	tt := ts[len(ts)-1]

	tt = strings.TrimLeft(tt, ".")

	if ptr {
		if deref {
			tt = "&" + tt
		} else {
			tt = "*" + tt
		}
	}
	return tt, tpkg
}

func toLowerFirst(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	s = string(a)
	return s
}

func getHint(t types.Type) string {
	if t.String() == "[]byte" {
		return object.HintBytes
	}
	switch v := t.(type) {
	case *types.Basic:
		switch v.Kind() {
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
			return object.HintInt
		case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			return object.HintUint
		case types.Float32, types.Float64:
			return object.HintFloat
		case types.String:
			return object.HintString
		}
	case *types.Array:
		st := v.Elem()
		ss := getHint(st)
		if ss != "" {
			return object.HintArray + "<" + ss + ">"
		}
	case *types.Slice:
		st := v.Elem()
		ss := getHint(st)
		if ss != "" {
			return object.HintArray + "<" + ss + ">"
		}
	case *types.Struct:
		return object.HintMap
	case *types.Pointer:
		st := v.Elem()
		return getHint(st)
	case *types.Tuple:
	case *types.Signature:
	case *types.Interface:
	case *types.Map:
		return object.HintMap
	case *types.Chan:
	case *types.Named:
	}
	// TODO(geoah) insane hack/assumption
	return object.HintMap
}
