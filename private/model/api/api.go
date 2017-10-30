// +build codegen

// Package api represents API abstractions for rendering service generated files.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// An API defines a service API's definition. and logic to serialize the definition.
type API struct {
	Metadata      Metadata
	Operations    map[string]*Operation
	Shapes        map[string]*Shape
	Waiters       []Waiter
	Documentation string
	Examples      Examples

	// Set to true to avoid removing unused shapes
	NoRemoveUnusedShapes bool

	// Set to true to avoid renaming to 'Input/Output' postfixed shapes
	NoRenameToplevelShapes bool

	// Set to true to ignore service/request init methods (for testing)
	NoInitMethods bool

	// Set to true to ignore String() and GoString methods (for generated tests)
	NoStringerMethods bool

	// Set to true to not generate API service name constants
	NoConstServiceNames bool

	// Set to true to not generate validation shapes
	NoValidataShapeMethods bool

	// Set to true to not generate struct field accessors
	NoGenStructFieldAccessors bool

	SvcClientImportPath string

	initialized bool
	imports     map[string]bool
	name        string
	path        string

	BaseCrosslinkURL string
}

// A Metadata is the metadata about an API's definition.
type Metadata struct {
	APIVersion          string
	EndpointPrefix      string
	SigningName         string
	ServiceAbbreviation string
	ServiceFullName     string
	SignatureVersion    string
	JSONVersion         string
	TargetPrefix        string
	Protocol            string
	UID                 string
	EndpointsID         string

	NoResolveEndpoint bool
}

var serviceAliases map[string]string

func Bootstrap() error {
	b, err := ioutil.ReadFile(filepath.Join("..", "models", "customizations", "service-aliases.json"))
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &serviceAliases)
}

// PackageName name of the API package
func (a *API) PackageName() string {
	return strings.ToLower(a.StructName())
}

// InterfacePackageName returns the package name for the interface.
func (a *API) InterfacePackageName() string {
	return a.PackageName() + "iface"
}

var nameRegex = regexp.MustCompile(`^Amazon|AWS\s*|\(.*|\s+|\W+`)

// StructName returns the struct name for a given API.
func (a *API) StructName() string {
	if a.name == "" {
		name := a.Metadata.ServiceAbbreviation
		if name == "" {
			name = a.Metadata.ServiceFullName
		}

		name = nameRegex.ReplaceAllString(name, "")

		a.name = name
		if name, ok := serviceAliases[strings.ToLower(name)]; ok {
			a.name = name
		}
	}
	return a.name
}

// UseInitMethods returns if the service's init method should be rendered.
func (a *API) UseInitMethods() bool {
	return !a.NoInitMethods
}

// NiceName returns the human friendly API name.
func (a *API) NiceName() string {
	if a.Metadata.ServiceAbbreviation != "" {
		return a.Metadata.ServiceAbbreviation
	}
	return a.Metadata.ServiceFullName
}

// ProtocolPackage returns the package name of the protocol this API uses.
func (a *API) ProtocolPackage() string {
	switch a.Metadata.Protocol {
	case "json":
		return "jsonrpc"
	case "ec2":
		return "ec2query"
	default:
		return strings.Replace(a.Metadata.Protocol, "-", "", -1)
	}
}

// OperationNames returns a slice of API operations supported.
func (a *API) OperationNames() []string {
	i, names := 0, make([]string, len(a.Operations))
	for n := range a.Operations {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

// OperationList returns a slice of API operation pointers
func (a *API) OperationList() []*Operation {
	list := make([]*Operation, len(a.Operations))
	for i, n := range a.OperationNames() {
		list[i] = a.Operations[n]
	}
	return list
}

// OperationHasOutputPlaceholder returns if any of the API operation input
// or output shapes are place holders.
func (a *API) OperationHasOutputPlaceholder() bool {
	for _, op := range a.Operations {
		if op.OutputRef.Shape.Placeholder {
			return true
		}
	}
	return false
}

// ShapeNames returns a slice of names for each shape used by the API.
func (a *API) ShapeNames() []string {
	i, names := 0, make([]string, len(a.Shapes))
	for n := range a.Shapes {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

// ShapeList returns a slice of shape pointers used by the API.
//
// Will exclude error shapes from the list of shapes returned.
func (a *API) ShapeList() []*Shape {
	list := make([]*Shape, 0, len(a.Shapes))
	for _, n := range a.ShapeNames() {
		// Ignore error shapes in list
		if s := a.Shapes[n]; !s.IsError {
			list = append(list, s)
		}
	}
	return list
}

// ShapeListErrors returns a list of the errors defined by the API model
func (a *API) ShapeListErrors() []*Shape {
	list := []*Shape{}
	for _, n := range a.ShapeNames() {
		// Ignore error shapes in list
		if s := a.Shapes[n]; s.IsError {
			list = append(list, s)
		}
	}
	return list
}

// resetImports resets the import map to default values.
func (a *API) resetImports() {
	a.imports = map[string]bool{}
}

// importsGoCode returns the generated Go import code.
func (a *API) importsGoCode() string {
	if len(a.imports) == 0 {
		return ""
	}

	corePkgs, extPkgs := []string{}, []string{}
	for i := range a.imports {
		if strings.Contains(i, ".") {
			extPkgs = append(extPkgs, i)
		} else {
			corePkgs = append(corePkgs, i)
		}
	}
	sort.Strings(corePkgs)
	sort.Strings(extPkgs)

	code := "import (\n"
	for _, i := range corePkgs {
		code += fmt.Sprintf("\t%q\n", i)
	}
	if len(corePkgs) > 0 {
		code += "\n"
	}
	for _, i := range extPkgs {
		code += fmt.Sprintf("\t%q\n", i)
	}
	code += ")\n\n"
	return code
}

// A tplAPI is the top level template for the API
var tplAPI = template.Must(template.New("api").Parse(`
{{ range $_, $o := .OperationList }}
{{ $o.GoCode }}
{{ end }}
`))

// APIGoCode renders the API in Go code. Returning it as a string
func (a *API) APIGoCode() string {
	a.resetImports()
	a.imports["github.com/aws/aws-sdk-go/service/"+a.PackageName()] = true
	a.imports["github.com/aws/aws-sdk-go/aws"] = true
	a.imports["github.com/aws/aws-sdk-go/aws/request"] = true

	var buf bytes.Buffer
	err := tplAPI.Execute(&buf, a)
	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + strings.TrimSpace(buf.String())
	return code
}

var noCrossLinkServices = map[string]struct{}{
	"apigateway":        {},
	"budgets":           {},
	"cloudsearch":       {},
	"cloudsearchdomain": {},
	"elastictranscoder": {},
	"es":                {},
	"glacier":           {},
	"importexport":      {},
	"iot":               {},
	"iot-data":          {},
	"machinelearning":   {},
	"rekognition":       {},
	"sdb":               {},
	"swf":               {},
}

// HasCrosslinks will return whether or not a service has crosslinking .
func HasCrosslinks(service string) bool {
	_, ok := noCrossLinkServices[service]
	return !ok
}

// GetCrosslinkURL returns the crosslinking URL for the shape based on the name and
// uid provided. Empty string is returned if no crosslink link could be determined.
func GetCrosslinkURL(baseURL, uid string, params ...string) string {
	if uid == "" || baseURL == "" {
		return ""
	}

	if !HasCrosslinks(strings.ToLower(ServiceIDFromUID(uid))) {
		return ""
	}

	return strings.Join(append([]string{baseURL, "goto", "WebAPI", uid}, params...), "/")
}

// ServiceIDFromUID will parse the service id from the uid and return
// the service id that was found.
func ServiceIDFromUID(uid string) string {
	found := 0
	i := len(uid) - 1
	for ; i >= 0; i-- {
		if uid[i] == '-' {
			found++
		}
		// Terminate after the date component is found, e.g. es-2017-11-11
		if found == 3 {
			break
		}
	}

	return uid[0:i]
}

// APIName returns the API's service name.
func (a *API) APIName() string {
	return a.name
}

var tplServiceDoc = template.Must(template.New("service docs").Funcs(template.FuncMap{
	"GetCrosslinkURL": GetCrosslinkURL,
}).
	Parse(`
// Package {{ .PackageName }} provides the client and types for making API
// requests to {{ .Metadata.ServiceFullName }}.
{{ if .Documentation -}}
//
{{ .Documentation }}
{{ end -}}
{{ $crosslinkURL := GetCrosslinkURL $.BaseCrosslinkURL $.Metadata.UID -}}
{{ if $crosslinkURL -}}
//
// See {{ $crosslinkURL }} for more information on this service.
{{ end -}}
//
// See {{ .PackageName }} package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/{{ .PackageName }}/
//
// Using the Client
//
// To {{ .Metadata.ServiceFullName }} with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
// 
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the {{ .Metadata.ServiceFullName }} client {{ .StructName }} for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/{{ .PackageName }}/#New
`))

// A tplService defines the template for the service generated code.
var tplService = template.Must(template.New("service").Parse(`
// {{ .StructName }} counts the API operations made to {{ .Metadata.ServiceFullName }}.
type {{ .StructName }} struct {
	svc {{ .InterfacePackageName }}.{{ .StructName }}API
	counts sync.Map
}

var _ {{ .InterfacePackageName }}.{{ .StructName }}API = &{{ .StructName }}{} 

// New creates a new instance of the {{ .StructName }} counter.
func New(svc {{ .InterfacePackageName }}.{{ .StructName }}API) *{{ .StructName }} {
	return &{{ .StructName }}{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *{{ .StructName }}) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *{{ .StructName }}) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *{{ .StructName }}) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *{{ .StructName }}) count(op string) int64 {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	return cnt.(*counter).count()
}

// counter is a threadsafe cumulative counter.
type counter struct {
	c int64
}

// count returns the current count.
func (c *counter) count() int64 {
	return atomic.LoadInt64(&c.c)
}

// inc increments the counter by one.
func (c *counter) inc() {
	atomic.AddInt64(&c.c, 1)
}
`))

// ServicePackageDoc generates the contents of the doc file for the service.
//
// Will also read in the custom doc templates for the service if found.
func (a *API) ServicePackageDoc() string {
	a.imports = map[string]bool{}

	var buf bytes.Buffer
	if err := tplServiceDoc.Execute(&buf, a); err != nil {
		panic(err)
	}

	return buf.String()
}

// ServiceGoCode renders service go code. Returning it as a string.
func (a *API) ServiceGoCode() string {
	a.resetImports()
	a.imports["sync"] = true
	a.imports["sync/atomic"] = true
	a.imports["github.com/aws/aws-sdk-go/aws/request"] = true
	a.imports["github.com/aws/aws-sdk-go/service/"+a.PackageName()+"/"+a.InterfacePackageName()] = true

	var buf bytes.Buffer
	err := tplService.Execute(&buf, a)
	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + buf.String()
	return code
}

// ExampleGoCode renders service example code. Returning it as a string.
func (a *API) ExampleGoCode() string {
	exs := []string{}
	imports := map[string]bool{}
	for _, o := range a.OperationList() {
		o.imports = map[string]bool{}
		exs = append(exs, o.Example())
		for k, v := range o.imports {
			imports[k] = v
		}
	}

	code := fmt.Sprintf("import (\n%q\n%q\n%q\n\n%q\n%q\n%q\n",
		"bytes",
		"fmt",
		"time",
		"github.com/aws/aws-sdk-go/aws",
		"github.com/aws/aws-sdk-go/aws/session",
		path.Join(a.SvcClientImportPath, a.PackageName()),
	)
	for k := range imports {
		code += fmt.Sprintf("%q\n", k)
	}
	code += ")\n\n"
	code += "var _ time.Duration\nvar _ bytes.Buffer\n\n"
	code += strings.Join(exs, "\n\n")
	return code
}

// NewAPIGoCodeWithPkgName returns a string of instantiating the API prefixed
// with its package name. Takes a string depicting the Config.
func (a *API) NewAPIGoCodeWithPkgName(cfg string) string {
	return fmt.Sprintf("%s.New(%s)", a.PackageName(), cfg)
}

// computes the validation chain for all input shapes
func (a *API) addShapeValidations() {
	for _, o := range a.Operations {
		resolveShapeValidations(o.InputRef.Shape)
	}
}

// Updates the source shape and all nested shapes with the validations that
// could possibly be needed.
func resolveShapeValidations(s *Shape, ancestry ...*Shape) {
	for _, a := range ancestry {
		if a == s {
			return
		}
	}

	children := []string{}
	for _, name := range s.MemberNames() {
		ref := s.MemberRefs[name]

		if s.IsRequired(name) && !s.Validations.Has(ref, ShapeValidationRequired) {
			s.Validations = append(s.Validations, ShapeValidation{
				Name: name, Ref: ref, Type: ShapeValidationRequired,
			})
		}

		if ref.Shape.Min != 0 && !s.Validations.Has(ref, ShapeValidationMinVal) {
			s.Validations = append(s.Validations, ShapeValidation{
				Name: name, Ref: ref, Type: ShapeValidationMinVal,
			})
		}

		switch ref.Shape.Type {
		case "map", "list", "structure":
			children = append(children, name)
		}
	}

	ancestry = append(ancestry, s)
	for _, name := range children {
		ref := s.MemberRefs[name]
		// Since this is a grab bag we will just continue since
		// we can't validate because we don't know the valued shape.
		if ref.JSONValue {
			continue
		}

		nestedShape := ref.Shape.NestedShape()

		var v *ShapeValidation
		if len(nestedShape.Validations) > 0 {
			v = &ShapeValidation{
				Name: name, Ref: ref, Type: ShapeValidationNested,
			}
		} else {
			resolveShapeValidations(nestedShape, ancestry...)
			if len(nestedShape.Validations) > 0 {
				v = &ShapeValidation{
					Name: name, Ref: ref, Type: ShapeValidationNested,
				}
			}
		}

		if v != nil && !s.Validations.Has(v.Ref, v.Type) {
			s.Validations = append(s.Validations, *v)
		}
	}
	ancestry = ancestry[:len(ancestry)-1]
}
