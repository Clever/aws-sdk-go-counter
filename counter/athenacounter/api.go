// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athenacounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
)

// BatchGetNamedQueryRequest is a passthrough to the underlying BatchGetNamedQueryRequest.
// It will increment the count of requests made to BatchGetNamedQuery.
func (c *Athena) BatchGetNamedQueryRequest(input *athena.BatchGetNamedQueryInput) (req *request.Request, output *athena.BatchGetNamedQueryOutput) {
	c.inc("BatchGetNamedQuery")
	return c.svc.BatchGetNamedQueryRequest(input)
}

// BatchGetNamedQuery is a passthrough to the underlying BatchGetNamedQuery method.
// It will increment the count of requests made to BatchGetNamedQuery.
func (c *Athena) BatchGetNamedQuery(input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
	c.inc("BatchGetNamedQuery")
	return c.svc.BatchGetNamedQuery(input)
}

// BatchGetNamedQueryWithContext is a passthrough to the underlying BatchGetNamedQueryWithContext method.
// It will increment the count of requests made to BatchGetNamedQuery.
func (c *Athena) BatchGetNamedQueryWithContext(ctx aws.Context, input *athena.BatchGetNamedQueryInput, opts ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {
	c.inc("BatchGetNamedQuery")
	return c.svc.BatchGetNamedQueryWithContext(ctx, input, opts...)
}

// BatchGetQueryExecutionRequest is a passthrough to the underlying BatchGetQueryExecutionRequest.
// It will increment the count of requests made to BatchGetQueryExecution.
func (c *Athena) BatchGetQueryExecutionRequest(input *athena.BatchGetQueryExecutionInput) (req *request.Request, output *athena.BatchGetQueryExecutionOutput) {
	c.inc("BatchGetQueryExecution")
	return c.svc.BatchGetQueryExecutionRequest(input)
}

// BatchGetQueryExecution is a passthrough to the underlying BatchGetQueryExecution method.
// It will increment the count of requests made to BatchGetQueryExecution.
func (c *Athena) BatchGetQueryExecution(input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
	c.inc("BatchGetQueryExecution")
	return c.svc.BatchGetQueryExecution(input)
}

// BatchGetQueryExecutionWithContext is a passthrough to the underlying BatchGetQueryExecutionWithContext method.
// It will increment the count of requests made to BatchGetQueryExecution.
func (c *Athena) BatchGetQueryExecutionWithContext(ctx aws.Context, input *athena.BatchGetQueryExecutionInput, opts ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {
	c.inc("BatchGetQueryExecution")
	return c.svc.BatchGetQueryExecutionWithContext(ctx, input, opts...)
}

// CreateNamedQueryRequest is a passthrough to the underlying CreateNamedQueryRequest.
// It will increment the count of requests made to CreateNamedQuery.
func (c *Athena) CreateNamedQueryRequest(input *athena.CreateNamedQueryInput) (req *request.Request, output *athena.CreateNamedQueryOutput) {
	c.inc("CreateNamedQuery")
	return c.svc.CreateNamedQueryRequest(input)
}

// CreateNamedQuery is a passthrough to the underlying CreateNamedQuery method.
// It will increment the count of requests made to CreateNamedQuery.
func (c *Athena) CreateNamedQuery(input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
	c.inc("CreateNamedQuery")
	return c.svc.CreateNamedQuery(input)
}

// CreateNamedQueryWithContext is a passthrough to the underlying CreateNamedQueryWithContext method.
// It will increment the count of requests made to CreateNamedQuery.
func (c *Athena) CreateNamedQueryWithContext(ctx aws.Context, input *athena.CreateNamedQueryInput, opts ...request.Option) (*athena.CreateNamedQueryOutput, error) {
	c.inc("CreateNamedQuery")
	return c.svc.CreateNamedQueryWithContext(ctx, input, opts...)
}

// DeleteNamedQueryRequest is a passthrough to the underlying DeleteNamedQueryRequest.
// It will increment the count of requests made to DeleteNamedQuery.
func (c *Athena) DeleteNamedQueryRequest(input *athena.DeleteNamedQueryInput) (req *request.Request, output *athena.DeleteNamedQueryOutput) {
	c.inc("DeleteNamedQuery")
	return c.svc.DeleteNamedQueryRequest(input)
}

// DeleteNamedQuery is a passthrough to the underlying DeleteNamedQuery method.
// It will increment the count of requests made to DeleteNamedQuery.
func (c *Athena) DeleteNamedQuery(input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
	c.inc("DeleteNamedQuery")
	return c.svc.DeleteNamedQuery(input)
}

// DeleteNamedQueryWithContext is a passthrough to the underlying DeleteNamedQueryWithContext method.
// It will increment the count of requests made to DeleteNamedQuery.
func (c *Athena) DeleteNamedQueryWithContext(ctx aws.Context, input *athena.DeleteNamedQueryInput, opts ...request.Option) (*athena.DeleteNamedQueryOutput, error) {
	c.inc("DeleteNamedQuery")
	return c.svc.DeleteNamedQueryWithContext(ctx, input, opts...)
}

// GetNamedQueryRequest is a passthrough to the underlying GetNamedQueryRequest.
// It will increment the count of requests made to GetNamedQuery.
func (c *Athena) GetNamedQueryRequest(input *athena.GetNamedQueryInput) (req *request.Request, output *athena.GetNamedQueryOutput) {
	c.inc("GetNamedQuery")
	return c.svc.GetNamedQueryRequest(input)
}

// GetNamedQuery is a passthrough to the underlying GetNamedQuery method.
// It will increment the count of requests made to GetNamedQuery.
func (c *Athena) GetNamedQuery(input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
	c.inc("GetNamedQuery")
	return c.svc.GetNamedQuery(input)
}

// GetNamedQueryWithContext is a passthrough to the underlying GetNamedQueryWithContext method.
// It will increment the count of requests made to GetNamedQuery.
func (c *Athena) GetNamedQueryWithContext(ctx aws.Context, input *athena.GetNamedQueryInput, opts ...request.Option) (*athena.GetNamedQueryOutput, error) {
	c.inc("GetNamedQuery")
	return c.svc.GetNamedQueryWithContext(ctx, input, opts...)
}

// GetQueryExecutionRequest is a passthrough to the underlying GetQueryExecutionRequest.
// It will increment the count of requests made to GetQueryExecution.
func (c *Athena) GetQueryExecutionRequest(input *athena.GetQueryExecutionInput) (req *request.Request, output *athena.GetQueryExecutionOutput) {
	c.inc("GetQueryExecution")
	return c.svc.GetQueryExecutionRequest(input)
}

// GetQueryExecution is a passthrough to the underlying GetQueryExecution method.
// It will increment the count of requests made to GetQueryExecution.
func (c *Athena) GetQueryExecution(input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
	c.inc("GetQueryExecution")
	return c.svc.GetQueryExecution(input)
}

// GetQueryExecutionWithContext is a passthrough to the underlying GetQueryExecutionWithContext method.
// It will increment the count of requests made to GetQueryExecution.
func (c *Athena) GetQueryExecutionWithContext(ctx aws.Context, input *athena.GetQueryExecutionInput, opts ...request.Option) (*athena.GetQueryExecutionOutput, error) {
	c.inc("GetQueryExecution")
	return c.svc.GetQueryExecutionWithContext(ctx, input, opts...)
}

// GetQueryResultsRequest is a passthrough to the underlying GetQueryResultsRequest.
// It will increment the count of requests made to GetQueryResults.
func (c *Athena) GetQueryResultsRequest(input *athena.GetQueryResultsInput) (req *request.Request, output *athena.GetQueryResultsOutput) {
	c.inc("GetQueryResults")
	return c.svc.GetQueryResultsRequest(input)
}

// GetQueryResults is a passthrough to the underlying GetQueryResults method.
// It will increment the count of requests made to GetQueryResults.
func (c *Athena) GetQueryResults(input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	c.inc("GetQueryResults")
	return c.svc.GetQueryResults(input)
}

// GetQueryResultsWithContext is a passthrough to the underlying GetQueryResultsWithContext method.
// It will increment the count of requests made to GetQueryResults.
func (c *Athena) GetQueryResultsWithContext(ctx aws.Context, input *athena.GetQueryResultsInput, opts ...request.Option) (*athena.GetQueryResultsOutput, error) {
	c.inc("GetQueryResults")
	return c.svc.GetQueryResultsWithContext(ctx, input, opts...)
}

// GetQueryResultsPages is a passthrough to the underlying GetQueryResultsPages method.
// It will increment the count of requests made to GetQueryResults on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use GetQueryResultsPagesWithContext to avoid this.
func (c *Athena) GetQueryResultsPages(input *athena.GetQueryResultsInput, fn func(*athena.GetQueryResultsOutput, bool) bool) error {
	wrappedFn := func(page *athena.GetQueryResultsOutput, lastPage bool) bool {
		c.inc("GetQueryResults")
		return fn(page, lastPage)
	}
	return c.GetQueryResultsPages(input, wrappedFn)
}

// GetQueryResultsPagesWithContext is a passthrough to the underlying GetQueryResultsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to GetQueryResults when applied to the request.
func (c *Athena) GetQueryResultsPagesWithContext(ctx aws.Context, input *athena.GetQueryResultsInput, fn func(*athena.GetQueryResultsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("GetQueryResults"))
	return c.GetQueryResultsPagesWithContext(ctx, input, fn, opts...)
}

// ListNamedQueriesRequest is a passthrough to the underlying ListNamedQueriesRequest.
// It will increment the count of requests made to ListNamedQueries.
func (c *Athena) ListNamedQueriesRequest(input *athena.ListNamedQueriesInput) (req *request.Request, output *athena.ListNamedQueriesOutput) {
	c.inc("ListNamedQueries")
	return c.svc.ListNamedQueriesRequest(input)
}

// ListNamedQueries is a passthrough to the underlying ListNamedQueries method.
// It will increment the count of requests made to ListNamedQueries.
func (c *Athena) ListNamedQueries(input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
	c.inc("ListNamedQueries")
	return c.svc.ListNamedQueries(input)
}

// ListNamedQueriesWithContext is a passthrough to the underlying ListNamedQueriesWithContext method.
// It will increment the count of requests made to ListNamedQueries.
func (c *Athena) ListNamedQueriesWithContext(ctx aws.Context, input *athena.ListNamedQueriesInput, opts ...request.Option) (*athena.ListNamedQueriesOutput, error) {
	c.inc("ListNamedQueries")
	return c.svc.ListNamedQueriesWithContext(ctx, input, opts...)
}

// ListNamedQueriesPages is a passthrough to the underlying ListNamedQueriesPages method.
// It will increment the count of requests made to ListNamedQueries on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListNamedQueriesPagesWithContext to avoid this.
func (c *Athena) ListNamedQueriesPages(input *athena.ListNamedQueriesInput, fn func(*athena.ListNamedQueriesOutput, bool) bool) error {
	wrappedFn := func(page *athena.ListNamedQueriesOutput, lastPage bool) bool {
		c.inc("ListNamedQueries")
		return fn(page, lastPage)
	}
	return c.ListNamedQueriesPages(input, wrappedFn)
}

// ListNamedQueriesPagesWithContext is a passthrough to the underlying ListNamedQueriesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListNamedQueries when applied to the request.
func (c *Athena) ListNamedQueriesPagesWithContext(ctx aws.Context, input *athena.ListNamedQueriesInput, fn func(*athena.ListNamedQueriesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListNamedQueries"))
	return c.ListNamedQueriesPagesWithContext(ctx, input, fn, opts...)
}

// ListQueryExecutionsRequest is a passthrough to the underlying ListQueryExecutionsRequest.
// It will increment the count of requests made to ListQueryExecutions.
func (c *Athena) ListQueryExecutionsRequest(input *athena.ListQueryExecutionsInput) (req *request.Request, output *athena.ListQueryExecutionsOutput) {
	c.inc("ListQueryExecutions")
	return c.svc.ListQueryExecutionsRequest(input)
}

// ListQueryExecutions is a passthrough to the underlying ListQueryExecutions method.
// It will increment the count of requests made to ListQueryExecutions.
func (c *Athena) ListQueryExecutions(input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
	c.inc("ListQueryExecutions")
	return c.svc.ListQueryExecutions(input)
}

// ListQueryExecutionsWithContext is a passthrough to the underlying ListQueryExecutionsWithContext method.
// It will increment the count of requests made to ListQueryExecutions.
func (c *Athena) ListQueryExecutionsWithContext(ctx aws.Context, input *athena.ListQueryExecutionsInput, opts ...request.Option) (*athena.ListQueryExecutionsOutput, error) {
	c.inc("ListQueryExecutions")
	return c.svc.ListQueryExecutionsWithContext(ctx, input, opts...)
}

// ListQueryExecutionsPages is a passthrough to the underlying ListQueryExecutionsPages method.
// It will increment the count of requests made to ListQueryExecutions on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListQueryExecutionsPagesWithContext to avoid this.
func (c *Athena) ListQueryExecutionsPages(input *athena.ListQueryExecutionsInput, fn func(*athena.ListQueryExecutionsOutput, bool) bool) error {
	wrappedFn := func(page *athena.ListQueryExecutionsOutput, lastPage bool) bool {
		c.inc("ListQueryExecutions")
		return fn(page, lastPage)
	}
	return c.ListQueryExecutionsPages(input, wrappedFn)
}

// ListQueryExecutionsPagesWithContext is a passthrough to the underlying ListQueryExecutionsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListQueryExecutions when applied to the request.
func (c *Athena) ListQueryExecutionsPagesWithContext(ctx aws.Context, input *athena.ListQueryExecutionsInput, fn func(*athena.ListQueryExecutionsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListQueryExecutions"))
	return c.ListQueryExecutionsPagesWithContext(ctx, input, fn, opts...)
}

// StartQueryExecutionRequest is a passthrough to the underlying StartQueryExecutionRequest.
// It will increment the count of requests made to StartQueryExecution.
func (c *Athena) StartQueryExecutionRequest(input *athena.StartQueryExecutionInput) (req *request.Request, output *athena.StartQueryExecutionOutput) {
	c.inc("StartQueryExecution")
	return c.svc.StartQueryExecutionRequest(input)
}

// StartQueryExecution is a passthrough to the underlying StartQueryExecution method.
// It will increment the count of requests made to StartQueryExecution.
func (c *Athena) StartQueryExecution(input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	c.inc("StartQueryExecution")
	return c.svc.StartQueryExecution(input)
}

// StartQueryExecutionWithContext is a passthrough to the underlying StartQueryExecutionWithContext method.
// It will increment the count of requests made to StartQueryExecution.
func (c *Athena) StartQueryExecutionWithContext(ctx aws.Context, input *athena.StartQueryExecutionInput, opts ...request.Option) (*athena.StartQueryExecutionOutput, error) {
	c.inc("StartQueryExecution")
	return c.svc.StartQueryExecutionWithContext(ctx, input, opts...)
}

// StopQueryExecutionRequest is a passthrough to the underlying StopQueryExecutionRequest.
// It will increment the count of requests made to StopQueryExecution.
func (c *Athena) StopQueryExecutionRequest(input *athena.StopQueryExecutionInput) (req *request.Request, output *athena.StopQueryExecutionOutput) {
	c.inc("StopQueryExecution")
	return c.svc.StopQueryExecutionRequest(input)
}

// StopQueryExecution is a passthrough to the underlying StopQueryExecution method.
// It will increment the count of requests made to StopQueryExecution.
func (c *Athena) StopQueryExecution(input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
	c.inc("StopQueryExecution")
	return c.svc.StopQueryExecution(input)
}

// StopQueryExecutionWithContext is a passthrough to the underlying StopQueryExecutionWithContext method.
// It will increment the count of requests made to StopQueryExecution.
func (c *Athena) StopQueryExecutionWithContext(ctx aws.Context, input *athena.StopQueryExecutionInput, opts ...request.Option) (*athena.StopQueryExecutionOutput, error) {
	c.inc("StopQueryExecution")
	return c.svc.StopQueryExecutionWithContext(ctx, input, opts...)
}
