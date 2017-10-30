// +build codegen

package api

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// An Operation defines a specific API Operation.
type Operation struct {
	API           *API `json:"-"`
	ExportedName  string
	Name          string
	Documentation string
	HTTP          HTTPInfo
	InputRef      ShapeRef   `json:"input"`
	OutputRef     ShapeRef   `json:"output"`
	ErrorRefs     []ShapeRef `json:"errors"`
	Paginator     *Paginator
	Deprecated    bool   `json:"deprecated"`
	AuthType      string `json:"authtype"`
	imports       map[string]bool
}

// A HTTPInfo defines the method of HTTP request for the Operation.
type HTTPInfo struct {
	Method       string
	RequestURI   string
	ResponseCode uint
}

// HasInput returns if the Operation accepts an input paramater
func (o *Operation) HasInput() bool {
	return o.InputRef.ShapeName != ""
}

// HasOutput returns if the Operation accepts an output parameter
func (o *Operation) HasOutput() bool {
	return o.OutputRef.ShapeName != ""
}

func (o *Operation) GetSigner() string {
	if o.AuthType == "v4-unsigned-body" {
		o.API.imports["github.com/aws/aws-sdk-go/aws/signer/v4"] = true
	}

	buf := bytes.NewBuffer(nil)

	switch o.AuthType {
	case "none":
		buf.WriteString("req.Config.Credentials = credentials.AnonymousCredentials")
	case "v4-unsigned-body":
		buf.WriteString("req.Handlers.Sign.Remove(v4.SignRequestHandler)\n")
		buf.WriteString("handler := v4.BuildNamedHandler(\"v4.CustomSignerHandler\", v4.WithUnsignedPayload)\n")
		buf.WriteString("req.Handlers.Sign.PushFrontNamed(handler)")
	}

	buf.WriteString("\n")
	return buf.String()
}

// tplOperation defines a template for rendering an API Operation
var tplOperation = template.Must(template.New("operation").Parse(`
// {{ .ExportedName }}Request is a passthrough to the underlying {{ .ExportedName }}Request.
// It will increment the count of requests made to {{ .ExportedName }}.
func (c *{{ .API.StructName }}) {{ .ExportedName }}Request(` +
	`input {{ .InputRef.GoTypeWithPkgName }}) (req *request.Request, output {{ .OutputRef.GoTypeWithPkgName }}) {
	c.inc("{{ .ExportedName }}")
	return c.svc.{{ .ExportedName }}Request(input)
}

// {{ .ExportedName }} is a passthrough to the underlying {{ .ExportedName }} method.
// It will increment the count of requests made to {{ .ExportedName }}.
func (c *{{ .API.StructName }}) {{ .ExportedName }}(` +
	`input {{ .InputRef.GoTypeWithPkgName }}) ({{ .OutputRef.GoTypeWithPkgName }}, error) {
	c.inc("{{ .ExportedName }}")
	return c.svc.{{ .ExportedName }}(input)
}

// {{ .ExportedName }}WithContext is a passthrough to the underlying {{ .ExportedName }}WithContext method.
// It will increment the count of requests made to {{ .ExportedName }}.
func (c *{{ .API.StructName }}) {{ .ExportedName }}WithContext(` +
	`ctx aws.Context, input {{ .InputRef.GoTypeWithPkgName }}, opts ...request.Option) ` +
	`({{ .OutputRef.GoTypeWithPkgName }}, error) {
	c.inc("{{ .ExportedName }}")
    return c.svc.{{ .ExportedName }}WithContext(ctx, input, opts...)
}

{{ if .Paginator }}
// {{ .ExportedName }}Pages is a passthrough to the underlying {{ .ExportedName }}Pages method.
// It will increment the count of requests made to {{ .ExportedName }} on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use {{ .ExportedName }}PagesWithContext to avoid this.
func (c *{{ .API.StructName }}) {{ .ExportedName }}Pages(` +
	`input {{ .InputRef.GoTypeWithPkgName }}, fn func({{ .OutputRef.GoTypeWithPkgName }}, bool) bool) error {
	wrappedFn := func(page {{ .OutputRef.GoTypeWithPkgName }}, lastPage bool) bool {
		c.inc("{{ .ExportedName }}")
		return fn(page, lastPage)
	}
	return c.{{ .ExportedName }}Pages(input, wrappedFn)
}

// {{ .ExportedName }}PagesWithContext is a passthrough to the underlying {{ .ExportedName }}PagesWithContext method.
// It will add a request.Option that will increment the count of requests made to {{ .ExportedName }} when applied to the request.
func (c *{{ .API.StructName }}) {{ .ExportedName }}PagesWithContext(` +
	`ctx aws.Context, ` +
	`input {{ .InputRef.GoTypeWithPkgName }}, ` +
	`fn func({{ .OutputRef.GoTypeWithPkgName }}, bool) bool, ` +
	`opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("{{ .ExportedName }}"))
	return c.{{ .ExportedName }}PagesWithContext(ctx, input, fn, opts...)
}
{{ end }}
`))

// GoCode returns a string of rendered GoCode for this Operation
func (o *Operation) GoCode() string {
	var buf bytes.Buffer
	err := tplOperation.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

// tplInfSig defines the template for rendering an Operation's signature within an Interface definition.
var tplInfSig = template.Must(template.New("opsig").Parse(`
{{ .ExportedName }}({{ .InputRef.GoTypeWithPkgName }}) ({{ .OutputRef.GoTypeWithPkgName }}, error)
{{ .ExportedName }}WithContext(aws.Context, {{ .InputRef.GoTypeWithPkgName }}, ...request.Option) ({{ .OutputRef.GoTypeWithPkgName }}, error)
{{ .ExportedName }}Request({{ .InputRef.GoTypeWithPkgName }}) (*request.Request, {{ .OutputRef.GoTypeWithPkgName }})

{{ if .Paginator -}}
{{ .ExportedName }}Pages({{ .InputRef.GoTypeWithPkgName }}, func({{ .OutputRef.GoTypeWithPkgName }}, bool) bool) error
{{ .ExportedName }}PagesWithContext(aws.Context, {{ .InputRef.GoTypeWithPkgName }}, func({{ .OutputRef.GoTypeWithPkgName }}, bool) bool, ...request.Option) error
{{- end }}
`))

// InterfaceSignature returns a string representing the Operation's interface{}
// functional signature.
func (o *Operation) InterfaceSignature() string {
	var buf bytes.Buffer
	err := tplInfSig.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

// tplExample defines the template for rendering an Operation example
var tplExample = template.Must(template.New("operationExample").Parse(`
func Example{{ .API.StructName }}_{{ .ExportedName }}() {
	sess := session.Must(session.NewSession())

	svc := {{ .API.PackageName }}.New(sess)

	{{ .ExampleInput }}
	resp, err := svc.{{ .ExportedName }}(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
`))

// Example returns a string of the rendered Go code for the Operation
func (o *Operation) Example() string {
	var buf bytes.Buffer
	err := tplExample.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

// ExampleInput return a string of the rendered Go code for an example's input parameters
func (o *Operation) ExampleInput() string {
	if len(o.InputRef.Shape.MemberRefs) == 0 {
		if strings.Contains(o.InputRef.GoTypeElem(), ".") {
			o.imports["github.com/aws/aws-sdk-go/service/"+strings.Split(o.InputRef.GoTypeElem(), ".")[0]] = true
			return fmt.Sprintf("var params *%s", o.InputRef.GoTypeElem())
		}
		return fmt.Sprintf("var params *%s.%s",
			o.API.PackageName(), o.InputRef.GoTypeElem())
	}
	e := example{o, map[string]int{}}
	return "params := " + e.traverseAny(o.InputRef.Shape, false, false)
}

// A example provides
type example struct {
	*Operation
	visited map[string]int
}

// traverseAny returns rendered Go code for the shape.
func (e *example) traverseAny(s *Shape, required, payload bool) string {
	str := ""
	e.visited[s.ShapeName]++

	switch s.Type {
	case "structure":
		str = e.traverseStruct(s, required, payload)
	case "list":
		str = e.traverseList(s, required, payload)
	case "map":
		str = e.traverseMap(s, required, payload)
	case "jsonvalue":
		str = "aws.JSONValue{\"key\": \"value\"}"
		if required {
			str += " // Required"
		}
	default:
		str = e.traverseScalar(s, required, payload)
	}

	e.visited[s.ShapeName]--

	return str
}

var reType = regexp.MustCompile(`\b([A-Z])`)

// traverseStruct returns rendered Go code for a structure type shape.
func (e *example) traverseStruct(s *Shape, required, payload bool) string {
	var buf bytes.Buffer

	if s.resolvePkg != "" {
		e.imports[s.resolvePkg] = true
		buf.WriteString("&" + s.GoTypeElem() + "{")
	} else {
		buf.WriteString("&" + s.API.PackageName() + "." + s.GoTypeElem() + "{")
	}

	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	req := make([]string, len(s.Required))
	copy(req, s.Required)
	sort.Strings(req)

	if e.visited[s.ShapeName] < 2 {
		for _, n := range req {
			m := s.MemberRefs[n].Shape
			p := n == s.Payload && (s.MemberRefs[n].Streaming || m.Streaming)
			buf.WriteString(n + ": " + e.traverseAny(m, true, p) + ",")
			if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
				buf.WriteString(" // Required")
			}
			buf.WriteString("\n")
		}

		for _, n := range s.MemberNames() {
			if s.IsRequired(n) {
				continue
			}
			m := s.MemberRefs[n].Shape
			p := n == s.Payload && (s.MemberRefs[n].Streaming || m.Streaming)
			buf.WriteString(n + ": " + e.traverseAny(m, false, p) + ",\n")
		}
	} else {
		buf.WriteString("// Recursive values...\n")
	}

	buf.WriteString("}")
	return buf.String()
}

// traverseMap returns rendered Go code for a map type shape.
func (e *example) traverseMap(s *Shape, required, payload bool) string {
	var buf bytes.Buffer

	t := ""
	if s.resolvePkg != "" {
		e.imports[s.resolvePkg] = true
		t = s.GoTypeElem()
	} else {
		t = reType.ReplaceAllString(s.GoTypeElem(), s.API.PackageName()+".$1")
	}
	buf.WriteString(t + "{")
	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	if e.visited[s.ShapeName] < 2 {
		m := s.ValueRef.Shape
		buf.WriteString("\"Key\": " + e.traverseAny(m, true, false) + ",")
		if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
			buf.WriteString(" // Required")
		}
		buf.WriteString("\n// More values...\n")
	} else {
		buf.WriteString("// Recursive values...\n")
	}
	buf.WriteString("}")

	return buf.String()
}

// traverseList returns rendered Go code for a list type shape.
func (e *example) traverseList(s *Shape, required, payload bool) string {
	var buf bytes.Buffer
	t := ""
	if s.resolvePkg != "" {
		e.imports[s.resolvePkg] = true
		t = s.GoTypeElem()
	} else {
		t = reType.ReplaceAllString(s.GoTypeElem(), s.API.PackageName()+".$1")
	}

	buf.WriteString(t + "{")
	if required {
		buf.WriteString(" // Required")
	}
	buf.WriteString("\n")

	if e.visited[s.ShapeName] < 2 {
		m := s.MemberRef.Shape
		buf.WriteString(e.traverseAny(m, true, false) + ",")
		if m.Type != "list" && m.Type != "structure" && m.Type != "map" {
			buf.WriteString(" // Required")
		}
		buf.WriteString("\n// More values...\n")
	} else {
		buf.WriteString("// Recursive values...\n")
	}
	buf.WriteString("}")

	return buf.String()
}

// traverseScalar returns an AWS Type string representation initialized to a value.
// Will panic if s is an unsupported shape type.
func (e *example) traverseScalar(s *Shape, required, payload bool) string {
	str := ""
	switch s.Type {
	case "integer", "long":
		str = `aws.Int64(1)`
	case "float", "double":
		str = `aws.Float64(1.0)`
	case "string", "character":
		str = `aws.String("` + s.ShapeName + `")`
	case "blob":
		if payload {
			str = `bytes.NewReader([]byte("PAYLOAD"))`
		} else {
			str = `[]byte("PAYLOAD")`
		}
	case "boolean":
		str = `aws.Bool(true)`
	case "timestamp":
		str = `aws.Time(time.Now())`
	default:
		panic("unsupported shape " + s.Type)
	}

	return str
}
