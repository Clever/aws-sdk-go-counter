// Package counter contains automatically generated counters around AWS API clients
package counter

//go:generate go run -tags codegen ../private/model/cli/gen-api/main.go -path=../counter ../models/apis/*/*/api-2.json
//go:generate gofmt -s -w ../counter
