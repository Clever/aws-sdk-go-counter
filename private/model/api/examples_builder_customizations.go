// +build codegen

package api

import (
	"bytes"
	"fmt"
)

type wafregionalExamplesBuilder struct {
	defaultExamplesBuilder
}

func (builder wafregionalExamplesBuilder) Imports(a *API) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`"fmt"
	"strings"
	"time"

	"github.com/Clever/aws-sdk-go-counter/aws"
	"github.com/Clever/aws-sdk-go-counter/aws/awserr"
	"github.com/Clever/aws-sdk-go-counter/aws/session"
	"github.com/Clever/aws-sdk-go-counter/service/waf"
	`)

	buf.WriteString(fmt.Sprintf("\"%s/%s\"", "github.com/aws/aws-sdk-go/service", a.PackageName()))
	return buf.String()
}
