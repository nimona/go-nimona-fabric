package main

import (
	"bytes"
	"strings"
	"text/template"
)

// nolint
var tpl = `// Code generated by nimona.io/tools/codegen. DO NOT EDIT.

package {{ .PackageAlias }}

import (
	"fmt"
	{{ range $alias, $pkg := .Imports }}
	{{ $alias }} "{{ $pkg }}"
	{{- end }}
)

{{ if .Domains }}
type (
	{{- range $domain := .Domains }}
	{{- range $object := $domain.Objects }}
	{{ $object.Name }} struct {
		{{- range $member := $object.Members }}
		{{ $member.Name }} {{ memberType $member.Type }} ` + "`" + `json:"{{ $member.Tag }},omitempty"` + "`" + `
		{{- end }}
	}
	{{- end }}
	{{- range $event := .Events }}
	{{ $event.Name }} struct {
		{{- range $member := $event.Members }}
		{{ $member.Name }} {{ memberType $member.Type }} ` + "`" + `json:"{{ $member.Tag }},omitempty"` + "`" + `
		{{- end }}
	}
	{{- end }}
	{{- end }}
)
{{ end }}

{{ range $domain := .Domains }}
{{ range $object := .Objects }}
func (e *{{ structName $object.Name }}) GetType() string {
	return "{{ $domain.Name }}.{{ $object.Name }}"
}

func (e *{{ structName $object.Name }}) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@ctx:s"] = "{{ $domain.Name }}.{{ $object.Name }}"
	{{- range $member := $object.Members }}
		{{- if $member.IsObject }}
			{{- if $member.IsRepeated }}
			if len(e.{{ $member.Name }}) > 0 {
				m["{{ $member.Tag }}"] = func() []interface{} {
					a := make([]interface{}, len(e.{{ $member.Name }}))
					for i, v := range e.{{ $member.Name }} {
						a[i] = v.ToObject().ToMap()
					}
					return a
				}()
			}
			{{- else }}
			if e.{{ $member.Name }} != nil {
				m["{{ $member.Tag }}"] = e.{{ $member.Name }}.ToObject().ToMap()
			}
			{{- end }}
		{{- else }}
			{{- if $member.IsRepeated }}
				if len(e.{{ $member.Name }}) > 0 {
					m["{{ $member.Tag }}"] = e.{{ $member.Name }}
				}
			{{- else }}
				m["{{ $member.Tag }}"] = e.{{ $member.Name }}
			{{- end }}
		{{- end }}
	{{- end }}
	return object.Object(m)
}

func (e *{{ structName $object.Name }}) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
{{ end }}
{{ range $event := .Events }}
func (e *{{ structName $event.Name }}) GetType() string {
	return "{{ $domain.Name }}.{{ $event.Name }}"
}

func (e *{{ structName $event.Name }}) ToObject() object.Object {
	m := map[string]interface{}{}
	m["@ctx:s"] = "{{ $domain.Name }}.{{ $event.Name }}"
	{{- range $member := $event.Members }}
		{{- if $member.IsObject }}
			{{- if $member.IsRepeated }}
			if len(e.{{ $member.Name }}) > 0 {
				m["{{ $member.Tag }}"] = func() []interface{} {
					a := make([]interface{}, len(e.{{ $member.Name }}))
					for i, v := range e.{{ $member.Name }} {
						a[i] = v.ToObject().ToMap()
					}
					return a
				}()
			}
			{{- else }}
			if e.{{ $member.Name }} != nil {
				m["{{ $member.Tag }}"] = e.{{ $member.Name }}.ToObject().ToMap()
			}
			{{- end }}
		{{- else }}
			{{- if $member.IsRepeated }}
				if len(e.{{ $member.Name }}) > 0 {
					m["{{ $member.Tag }}"] = e.{{ $member.Name }}
				}
			{{- else }}
				m["{{ $member.Tag }}"] = e.{{ $member.Name }}
			{{- end }}
		{{- end }}
	{{- end }}
	return object.Object(m)
}

func (e *{{ structName $event.Name }}) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
{{ end }}
{{ end }}
`

func Generate(doc *Document, output string) ([]byte, error) {
	originalImports := map[string]string{}
	t, err := template.New("tpl").Funcs(template.FuncMap{
		"structName": func(name string) string {
			ps := strings.Split(name, "/")
			ps = strings.Split(ps[len(ps)-1], ".")
			return ucFirst(ps[len(ps)-1])
		},
		"memberType": func(name string) string {
			for alias, pkg := range originalImports {
				name = strings.Replace(name, pkg, alias, 1)
			}
			return name
		},
	}).Parse(tpl)
	if err != nil {
		return nil, err
	}

	lPackage := strings.ToLower(lastSegment(doc.Package))
	for i, s := range doc.Domains {
		for k, e := range s.Events {
			lDomain := strings.ToLower(lastSegment(s.Name))
			if lDomain != lPackage {
				e.Name = ucFirst(lDomain) + ucFirst(e.Name)
			}
			if e.IsSigned {
				streamPkg := "stream."
				if doc.Package == "nimona.io/stream" {
					streamPkg = ""
				}
				doc.Domains[i].Events[k].Members = append(
					doc.Domains[i].Events[k].Members,
					&Member{
						Name:     "Signature",
						Type:     "*crypto.Signature",
						Tag:      "@signature:o",
						IsObject: true,
					},
					&Member{
						Name:       "Authors",
						Type:       "[]*" + streamPkg + "Author",
						Tag:        "@authors:ao",
						IsRepeated: true,
						IsObject:   true,
					},
				)
			}
		}
	}

	doc.Imports["json"] = "encoding/json"
	if doc.Package != "nimona.io/object" {
		doc.Imports["object"] = "nimona.io/object"
	}
	if doc.Package != "nimona.io/stream" {
		doc.Imports["stream"] = "nimona.io/stream"
	}

	for alias, pkg := range doc.Imports {
		originalImports[alias] = pkg
	}

	for i, pkg := range doc.Imports {
		doc.Imports[i] = strings.Replace(pkg, "nimona.io/", "nimona.io/pkg/", 1)
	}

	out := bytes.NewBuffer([]byte{})
	if err := t.Execute(out, doc); err != nil {
		return nil, err
	}

	res := out.String()
	if doc.Package == "nimona.io/object" {
		res = strings.ReplaceAll(res, "object.", "")
	}

	return []byte(res), nil
}

// lastSegment returns the last part of a namespace,
// ie lastSegment(nimona.io/stream) returns stream
func lastSegment(s string) string {
	ps := strings.Split(s, "/")
	return ps[len(ps)-1]
}
