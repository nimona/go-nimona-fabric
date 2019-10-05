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

{{ if .Structs }}
type (
	{{- range $struct := .Structs }}
	{{ $struct.Name }} struct {
		{{- range $member := $struct.Members }}
		{{ $member.Name }} {{ memberType $member.Type }} ` + "`" + `json:"{{ $member.Tag }},omitempty"` + "`" + `
		{{- end }}
	}
	{{- end }}
)
{{ end }}

{{ range $struct := .Structs }}
func (e *{{ $struct.Name }}) ContextName() string {
	return "{{ $.Package }}"
}

func (e *{{ $struct.Name }}) GetType() string {
	return "{{ $struct.Name }}"
}

func (e *{{ $struct.Name }}) ToObject() object.Object {
	m := map[string]interface{}{
		"@ctx:s": "{{ $struct.Name }}",
		"@struct:s": "{{ $struct.Name }}",
	}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &m)
	return object.Object(m)
}

func (e *{{ $struct.Name }}) FromObject(o object.Object) error {
	b, _ := json.Marshal(map[string]interface{}(o))
	return json.Unmarshal(b, e)
}
{{ end }}

{{ if .Domains }}
type (
	{{- range $domain := .Domains }}
	{{- range $struct := $domain.Structs }}
	{{ $struct.Name }} struct {
		{{- range $member := $struct.Members }}
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
{{ range $event := .Events }}
func (e *{{ structName $event.Name }}) EventName() string {
	return "{{ $event.Name }}"
}

func (e *{{ structName $event.Name }}) GetType() string {
	return "{{ $domain.Name }}.{{ $event.Name }}"
}

func (e *{{ structName $event.Name }}) ToObject() object.Object {
	m := map[string]interface{}{
		"@ctx:s": "{{ $domain.Name }}.{{ $event.Name }}",
		"@domain:s": "{{ $domain.Name }}",
		"@event:s": "{{ $event.Name }}",
	}
	b, _ := json.Marshal(e)
	json.Unmarshal(b, &m)
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
	t, err := template.New("tpl").Funcs(template.FuncMap{
		"structName": func(name string) string {
			ps := strings.Split(name, "/")
			ps = strings.Split(ps[len(ps)-1], ".")
			return ucFirst(ps[len(ps)-1])
		},
		"memberType": func(name string) string {
			ps := strings.Split(name, "/")
			p := ps[len(ps)-1]
			return p
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
				doc.Domains[i].Events[k].Members = append(
					doc.Domains[i].Events[k].Members,
					&Member{
						Name: "Signature",
						Type: "*crypto.Signature",
						Tag:  "@signature:o",
					},
					&Member{
						Name: "Authors",
						Type: "[]*crypto.PublicKey",
						Tag:  "@authors:ao",
					},
				)
			}
		}
	}

	doc.Imports["json"] = "encoding/json"
	doc.Imports["object"] = "nimona.io/object"

	for i, pkg := range doc.Imports {
		doc.Imports[i] = strings.Replace(pkg, "nimona.io/", "nimona.io/pkg/", 1)
	}

	out := bytes.NewBuffer([]byte{})
	if err := t.Execute(out, doc); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// lastSegment returns the last part of a namespace,
// ie lastSegment(nimona.io/stream) returns stream
func lastSegment(s string) string {
	ps := strings.Split(s, "/")
	return ps[len(ps)-1]
}
