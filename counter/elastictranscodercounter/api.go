// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscodercounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
)

// CancelJobRequest is a passthrough to the underlying CancelJobRequest.
// It will increment the count of requests made to CancelJob.
func (c *ElasticTranscoder) CancelJobRequest(input *elastictranscoder.CancelJobInput) (req *request.Request, output *elastictranscoder.CancelJobOutput) {
	c.inc("CancelJob")
	return c.svc.CancelJobRequest(input)
}

// CancelJob is a passthrough to the underlying CancelJob method.
// It will increment the count of requests made to CancelJob.
func (c *ElasticTranscoder) CancelJob(input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
	c.inc("CancelJob")
	return c.svc.CancelJob(input)
}

// CancelJobWithContext is a passthrough to the underlying CancelJobWithContext method.
// It will increment the count of requests made to CancelJob.
func (c *ElasticTranscoder) CancelJobWithContext(ctx aws.Context, input *elastictranscoder.CancelJobInput, opts ...request.Option) (*elastictranscoder.CancelJobOutput, error) {
	c.inc("CancelJob")
	return c.svc.CancelJobWithContext(ctx, input, opts...)
}

// CreateJobRequest is a passthrough to the underlying CreateJobRequest.
// It will increment the count of requests made to CreateJob.
func (c *ElasticTranscoder) CreateJobRequest(input *elastictranscoder.CreateJobInput) (req *request.Request, output *elastictranscoder.CreateJobResponse) {
	c.inc("CreateJob")
	return c.svc.CreateJobRequest(input)
}

// CreateJob is a passthrough to the underlying CreateJob method.
// It will increment the count of requests made to CreateJob.
func (c *ElasticTranscoder) CreateJob(input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
	c.inc("CreateJob")
	return c.svc.CreateJob(input)
}

// CreateJobWithContext is a passthrough to the underlying CreateJobWithContext method.
// It will increment the count of requests made to CreateJob.
func (c *ElasticTranscoder) CreateJobWithContext(ctx aws.Context, input *elastictranscoder.CreateJobInput, opts ...request.Option) (*elastictranscoder.CreateJobResponse, error) {
	c.inc("CreateJob")
	return c.svc.CreateJobWithContext(ctx, input, opts...)
}

// CreatePipelineRequest is a passthrough to the underlying CreatePipelineRequest.
// It will increment the count of requests made to CreatePipeline.
func (c *ElasticTranscoder) CreatePipelineRequest(input *elastictranscoder.CreatePipelineInput) (req *request.Request, output *elastictranscoder.CreatePipelineOutput) {
	c.inc("CreatePipeline")
	return c.svc.CreatePipelineRequest(input)
}

// CreatePipeline is a passthrough to the underlying CreatePipeline method.
// It will increment the count of requests made to CreatePipeline.
func (c *ElasticTranscoder) CreatePipeline(input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
	c.inc("CreatePipeline")
	return c.svc.CreatePipeline(input)
}

// CreatePipelineWithContext is a passthrough to the underlying CreatePipelineWithContext method.
// It will increment the count of requests made to CreatePipeline.
func (c *ElasticTranscoder) CreatePipelineWithContext(ctx aws.Context, input *elastictranscoder.CreatePipelineInput, opts ...request.Option) (*elastictranscoder.CreatePipelineOutput, error) {
	c.inc("CreatePipeline")
	return c.svc.CreatePipelineWithContext(ctx, input, opts...)
}

// CreatePresetRequest is a passthrough to the underlying CreatePresetRequest.
// It will increment the count of requests made to CreatePreset.
func (c *ElasticTranscoder) CreatePresetRequest(input *elastictranscoder.CreatePresetInput) (req *request.Request, output *elastictranscoder.CreatePresetOutput) {
	c.inc("CreatePreset")
	return c.svc.CreatePresetRequest(input)
}

// CreatePreset is a passthrough to the underlying CreatePreset method.
// It will increment the count of requests made to CreatePreset.
func (c *ElasticTranscoder) CreatePreset(input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error) {
	c.inc("CreatePreset")
	return c.svc.CreatePreset(input)
}

// CreatePresetWithContext is a passthrough to the underlying CreatePresetWithContext method.
// It will increment the count of requests made to CreatePreset.
func (c *ElasticTranscoder) CreatePresetWithContext(ctx aws.Context, input *elastictranscoder.CreatePresetInput, opts ...request.Option) (*elastictranscoder.CreatePresetOutput, error) {
	c.inc("CreatePreset")
	return c.svc.CreatePresetWithContext(ctx, input, opts...)
}

// DeletePipelineRequest is a passthrough to the underlying DeletePipelineRequest.
// It will increment the count of requests made to DeletePipeline.
func (c *ElasticTranscoder) DeletePipelineRequest(input *elastictranscoder.DeletePipelineInput) (req *request.Request, output *elastictranscoder.DeletePipelineOutput) {
	c.inc("DeletePipeline")
	return c.svc.DeletePipelineRequest(input)
}

// DeletePipeline is a passthrough to the underlying DeletePipeline method.
// It will increment the count of requests made to DeletePipeline.
func (c *ElasticTranscoder) DeletePipeline(input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error) {
	c.inc("DeletePipeline")
	return c.svc.DeletePipeline(input)
}

// DeletePipelineWithContext is a passthrough to the underlying DeletePipelineWithContext method.
// It will increment the count of requests made to DeletePipeline.
func (c *ElasticTranscoder) DeletePipelineWithContext(ctx aws.Context, input *elastictranscoder.DeletePipelineInput, opts ...request.Option) (*elastictranscoder.DeletePipelineOutput, error) {
	c.inc("DeletePipeline")
	return c.svc.DeletePipelineWithContext(ctx, input, opts...)
}

// DeletePresetRequest is a passthrough to the underlying DeletePresetRequest.
// It will increment the count of requests made to DeletePreset.
func (c *ElasticTranscoder) DeletePresetRequest(input *elastictranscoder.DeletePresetInput) (req *request.Request, output *elastictranscoder.DeletePresetOutput) {
	c.inc("DeletePreset")
	return c.svc.DeletePresetRequest(input)
}

// DeletePreset is a passthrough to the underlying DeletePreset method.
// It will increment the count of requests made to DeletePreset.
func (c *ElasticTranscoder) DeletePreset(input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error) {
	c.inc("DeletePreset")
	return c.svc.DeletePreset(input)
}

// DeletePresetWithContext is a passthrough to the underlying DeletePresetWithContext method.
// It will increment the count of requests made to DeletePreset.
func (c *ElasticTranscoder) DeletePresetWithContext(ctx aws.Context, input *elastictranscoder.DeletePresetInput, opts ...request.Option) (*elastictranscoder.DeletePresetOutput, error) {
	c.inc("DeletePreset")
	return c.svc.DeletePresetWithContext(ctx, input, opts...)
}

// ListJobsByPipelineRequest is a passthrough to the underlying ListJobsByPipelineRequest.
// It will increment the count of requests made to ListJobsByPipeline.
func (c *ElasticTranscoder) ListJobsByPipelineRequest(input *elastictranscoder.ListJobsByPipelineInput) (req *request.Request, output *elastictranscoder.ListJobsByPipelineOutput) {
	c.inc("ListJobsByPipeline")
	return c.svc.ListJobsByPipelineRequest(input)
}

// ListJobsByPipeline is a passthrough to the underlying ListJobsByPipeline method.
// It will increment the count of requests made to ListJobsByPipeline.
func (c *ElasticTranscoder) ListJobsByPipeline(input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	c.inc("ListJobsByPipeline")
	return c.svc.ListJobsByPipeline(input)
}

// ListJobsByPipelineWithContext is a passthrough to the underlying ListJobsByPipelineWithContext method.
// It will increment the count of requests made to ListJobsByPipeline.
func (c *ElasticTranscoder) ListJobsByPipelineWithContext(ctx aws.Context, input *elastictranscoder.ListJobsByPipelineInput, opts ...request.Option) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	c.inc("ListJobsByPipeline")
	return c.svc.ListJobsByPipelineWithContext(ctx, input, opts...)
}

// ListJobsByPipelinePages is a passthrough to the underlying ListJobsByPipelinePages method.
// It will increment the count of requests made to ListJobsByPipeline on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListJobsByPipelinePagesWithContext to avoid this.
func (c *ElasticTranscoder) ListJobsByPipelinePages(input *elastictranscoder.ListJobsByPipelineInput, fn func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool) error {
	wrappedFn := func(page *elastictranscoder.ListJobsByPipelineOutput, lastPage bool) bool {
		c.inc("ListJobsByPipeline")
		return fn(page, lastPage)
	}
	return c.ListJobsByPipelinePages(input, wrappedFn)
}

// ListJobsByPipelinePagesWithContext is a passthrough to the underlying ListJobsByPipelinePagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListJobsByPipeline when applied to the request.
func (c *ElasticTranscoder) ListJobsByPipelinePagesWithContext(ctx aws.Context, input *elastictranscoder.ListJobsByPipelineInput, fn func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListJobsByPipeline"))
	return c.ListJobsByPipelinePagesWithContext(ctx, input, fn, opts...)
}

// ListJobsByStatusRequest is a passthrough to the underlying ListJobsByStatusRequest.
// It will increment the count of requests made to ListJobsByStatus.
func (c *ElasticTranscoder) ListJobsByStatusRequest(input *elastictranscoder.ListJobsByStatusInput) (req *request.Request, output *elastictranscoder.ListJobsByStatusOutput) {
	c.inc("ListJobsByStatus")
	return c.svc.ListJobsByStatusRequest(input)
}

// ListJobsByStatus is a passthrough to the underlying ListJobsByStatus method.
// It will increment the count of requests made to ListJobsByStatus.
func (c *ElasticTranscoder) ListJobsByStatus(input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
	c.inc("ListJobsByStatus")
	return c.svc.ListJobsByStatus(input)
}

// ListJobsByStatusWithContext is a passthrough to the underlying ListJobsByStatusWithContext method.
// It will increment the count of requests made to ListJobsByStatus.
func (c *ElasticTranscoder) ListJobsByStatusWithContext(ctx aws.Context, input *elastictranscoder.ListJobsByStatusInput, opts ...request.Option) (*elastictranscoder.ListJobsByStatusOutput, error) {
	c.inc("ListJobsByStatus")
	return c.svc.ListJobsByStatusWithContext(ctx, input, opts...)
}

// ListJobsByStatusPages is a passthrough to the underlying ListJobsByStatusPages method.
// It will increment the count of requests made to ListJobsByStatus on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListJobsByStatusPagesWithContext to avoid this.
func (c *ElasticTranscoder) ListJobsByStatusPages(input *elastictranscoder.ListJobsByStatusInput, fn func(*elastictranscoder.ListJobsByStatusOutput, bool) bool) error {
	wrappedFn := func(page *elastictranscoder.ListJobsByStatusOutput, lastPage bool) bool {
		c.inc("ListJobsByStatus")
		return fn(page, lastPage)
	}
	return c.ListJobsByStatusPages(input, wrappedFn)
}

// ListJobsByStatusPagesWithContext is a passthrough to the underlying ListJobsByStatusPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListJobsByStatus when applied to the request.
func (c *ElasticTranscoder) ListJobsByStatusPagesWithContext(ctx aws.Context, input *elastictranscoder.ListJobsByStatusInput, fn func(*elastictranscoder.ListJobsByStatusOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListJobsByStatus"))
	return c.ListJobsByStatusPagesWithContext(ctx, input, fn, opts...)
}

// ListPipelinesRequest is a passthrough to the underlying ListPipelinesRequest.
// It will increment the count of requests made to ListPipelines.
func (c *ElasticTranscoder) ListPipelinesRequest(input *elastictranscoder.ListPipelinesInput) (req *request.Request, output *elastictranscoder.ListPipelinesOutput) {
	c.inc("ListPipelines")
	return c.svc.ListPipelinesRequest(input)
}

// ListPipelines is a passthrough to the underlying ListPipelines method.
// It will increment the count of requests made to ListPipelines.
func (c *ElasticTranscoder) ListPipelines(input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	c.inc("ListPipelines")
	return c.svc.ListPipelines(input)
}

// ListPipelinesWithContext is a passthrough to the underlying ListPipelinesWithContext method.
// It will increment the count of requests made to ListPipelines.
func (c *ElasticTranscoder) ListPipelinesWithContext(ctx aws.Context, input *elastictranscoder.ListPipelinesInput, opts ...request.Option) (*elastictranscoder.ListPipelinesOutput, error) {
	c.inc("ListPipelines")
	return c.svc.ListPipelinesWithContext(ctx, input, opts...)
}

// ListPipelinesPages is a passthrough to the underlying ListPipelinesPages method.
// It will increment the count of requests made to ListPipelines on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListPipelinesPagesWithContext to avoid this.
func (c *ElasticTranscoder) ListPipelinesPages(input *elastictranscoder.ListPipelinesInput, fn func(*elastictranscoder.ListPipelinesOutput, bool) bool) error {
	wrappedFn := func(page *elastictranscoder.ListPipelinesOutput, lastPage bool) bool {
		c.inc("ListPipelines")
		return fn(page, lastPage)
	}
	return c.ListPipelinesPages(input, wrappedFn)
}

// ListPipelinesPagesWithContext is a passthrough to the underlying ListPipelinesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListPipelines when applied to the request.
func (c *ElasticTranscoder) ListPipelinesPagesWithContext(ctx aws.Context, input *elastictranscoder.ListPipelinesInput, fn func(*elastictranscoder.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListPipelines"))
	return c.ListPipelinesPagesWithContext(ctx, input, fn, opts...)
}

// ListPresetsRequest is a passthrough to the underlying ListPresetsRequest.
// It will increment the count of requests made to ListPresets.
func (c *ElasticTranscoder) ListPresetsRequest(input *elastictranscoder.ListPresetsInput) (req *request.Request, output *elastictranscoder.ListPresetsOutput) {
	c.inc("ListPresets")
	return c.svc.ListPresetsRequest(input)
}

// ListPresets is a passthrough to the underlying ListPresets method.
// It will increment the count of requests made to ListPresets.
func (c *ElasticTranscoder) ListPresets(input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
	c.inc("ListPresets")
	return c.svc.ListPresets(input)
}

// ListPresetsWithContext is a passthrough to the underlying ListPresetsWithContext method.
// It will increment the count of requests made to ListPresets.
func (c *ElasticTranscoder) ListPresetsWithContext(ctx aws.Context, input *elastictranscoder.ListPresetsInput, opts ...request.Option) (*elastictranscoder.ListPresetsOutput, error) {
	c.inc("ListPresets")
	return c.svc.ListPresetsWithContext(ctx, input, opts...)
}

// ListPresetsPages is a passthrough to the underlying ListPresetsPages method.
// It will increment the count of requests made to ListPresets on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListPresetsPagesWithContext to avoid this.
func (c *ElasticTranscoder) ListPresetsPages(input *elastictranscoder.ListPresetsInput, fn func(*elastictranscoder.ListPresetsOutput, bool) bool) error {
	wrappedFn := func(page *elastictranscoder.ListPresetsOutput, lastPage bool) bool {
		c.inc("ListPresets")
		return fn(page, lastPage)
	}
	return c.ListPresetsPages(input, wrappedFn)
}

// ListPresetsPagesWithContext is a passthrough to the underlying ListPresetsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListPresets when applied to the request.
func (c *ElasticTranscoder) ListPresetsPagesWithContext(ctx aws.Context, input *elastictranscoder.ListPresetsInput, fn func(*elastictranscoder.ListPresetsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListPresets"))
	return c.ListPresetsPagesWithContext(ctx, input, fn, opts...)
}

// ReadJobRequest is a passthrough to the underlying ReadJobRequest.
// It will increment the count of requests made to ReadJob.
func (c *ElasticTranscoder) ReadJobRequest(input *elastictranscoder.ReadJobInput) (req *request.Request, output *elastictranscoder.ReadJobOutput) {
	c.inc("ReadJob")
	return c.svc.ReadJobRequest(input)
}

// ReadJob is a passthrough to the underlying ReadJob method.
// It will increment the count of requests made to ReadJob.
func (c *ElasticTranscoder) ReadJob(input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
	c.inc("ReadJob")
	return c.svc.ReadJob(input)
}

// ReadJobWithContext is a passthrough to the underlying ReadJobWithContext method.
// It will increment the count of requests made to ReadJob.
func (c *ElasticTranscoder) ReadJobWithContext(ctx aws.Context, input *elastictranscoder.ReadJobInput, opts ...request.Option) (*elastictranscoder.ReadJobOutput, error) {
	c.inc("ReadJob")
	return c.svc.ReadJobWithContext(ctx, input, opts...)
}

// ReadPipelineRequest is a passthrough to the underlying ReadPipelineRequest.
// It will increment the count of requests made to ReadPipeline.
func (c *ElasticTranscoder) ReadPipelineRequest(input *elastictranscoder.ReadPipelineInput) (req *request.Request, output *elastictranscoder.ReadPipelineOutput) {
	c.inc("ReadPipeline")
	return c.svc.ReadPipelineRequest(input)
}

// ReadPipeline is a passthrough to the underlying ReadPipeline method.
// It will increment the count of requests made to ReadPipeline.
func (c *ElasticTranscoder) ReadPipeline(input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error) {
	c.inc("ReadPipeline")
	return c.svc.ReadPipeline(input)
}

// ReadPipelineWithContext is a passthrough to the underlying ReadPipelineWithContext method.
// It will increment the count of requests made to ReadPipeline.
func (c *ElasticTranscoder) ReadPipelineWithContext(ctx aws.Context, input *elastictranscoder.ReadPipelineInput, opts ...request.Option) (*elastictranscoder.ReadPipelineOutput, error) {
	c.inc("ReadPipeline")
	return c.svc.ReadPipelineWithContext(ctx, input, opts...)
}

// ReadPresetRequest is a passthrough to the underlying ReadPresetRequest.
// It will increment the count of requests made to ReadPreset.
func (c *ElasticTranscoder) ReadPresetRequest(input *elastictranscoder.ReadPresetInput) (req *request.Request, output *elastictranscoder.ReadPresetOutput) {
	c.inc("ReadPreset")
	return c.svc.ReadPresetRequest(input)
}

// ReadPreset is a passthrough to the underlying ReadPreset method.
// It will increment the count of requests made to ReadPreset.
func (c *ElasticTranscoder) ReadPreset(input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error) {
	c.inc("ReadPreset")
	return c.svc.ReadPreset(input)
}

// ReadPresetWithContext is a passthrough to the underlying ReadPresetWithContext method.
// It will increment the count of requests made to ReadPreset.
func (c *ElasticTranscoder) ReadPresetWithContext(ctx aws.Context, input *elastictranscoder.ReadPresetInput, opts ...request.Option) (*elastictranscoder.ReadPresetOutput, error) {
	c.inc("ReadPreset")
	return c.svc.ReadPresetWithContext(ctx, input, opts...)
}

// TestRoleRequest is a passthrough to the underlying TestRoleRequest.
// It will increment the count of requests made to TestRole.
func (c *ElasticTranscoder) TestRoleRequest(input *elastictranscoder.TestRoleInput) (req *request.Request, output *elastictranscoder.TestRoleOutput) {
	c.inc("TestRole")
	return c.svc.TestRoleRequest(input)
}

// TestRole is a passthrough to the underlying TestRole method.
// It will increment the count of requests made to TestRole.
func (c *ElasticTranscoder) TestRole(input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error) {
	c.inc("TestRole")
	return c.svc.TestRole(input)
}

// TestRoleWithContext is a passthrough to the underlying TestRoleWithContext method.
// It will increment the count of requests made to TestRole.
func (c *ElasticTranscoder) TestRoleWithContext(ctx aws.Context, input *elastictranscoder.TestRoleInput, opts ...request.Option) (*elastictranscoder.TestRoleOutput, error) {
	c.inc("TestRole")
	return c.svc.TestRoleWithContext(ctx, input, opts...)
}

// UpdatePipelineRequest is a passthrough to the underlying UpdatePipelineRequest.
// It will increment the count of requests made to UpdatePipeline.
func (c *ElasticTranscoder) UpdatePipelineRequest(input *elastictranscoder.UpdatePipelineInput) (req *request.Request, output *elastictranscoder.UpdatePipelineOutput) {
	c.inc("UpdatePipeline")
	return c.svc.UpdatePipelineRequest(input)
}

// UpdatePipeline is a passthrough to the underlying UpdatePipeline method.
// It will increment the count of requests made to UpdatePipeline.
func (c *ElasticTranscoder) UpdatePipeline(input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error) {
	c.inc("UpdatePipeline")
	return c.svc.UpdatePipeline(input)
}

// UpdatePipelineWithContext is a passthrough to the underlying UpdatePipelineWithContext method.
// It will increment the count of requests made to UpdatePipeline.
func (c *ElasticTranscoder) UpdatePipelineWithContext(ctx aws.Context, input *elastictranscoder.UpdatePipelineInput, opts ...request.Option) (*elastictranscoder.UpdatePipelineOutput, error) {
	c.inc("UpdatePipeline")
	return c.svc.UpdatePipelineWithContext(ctx, input, opts...)
}

// UpdatePipelineNotificationsRequest is a passthrough to the underlying UpdatePipelineNotificationsRequest.
// It will increment the count of requests made to UpdatePipelineNotifications.
func (c *ElasticTranscoder) UpdatePipelineNotificationsRequest(input *elastictranscoder.UpdatePipelineNotificationsInput) (req *request.Request, output *elastictranscoder.UpdatePipelineNotificationsOutput) {
	c.inc("UpdatePipelineNotifications")
	return c.svc.UpdatePipelineNotificationsRequest(input)
}

// UpdatePipelineNotifications is a passthrough to the underlying UpdatePipelineNotifications method.
// It will increment the count of requests made to UpdatePipelineNotifications.
func (c *ElasticTranscoder) UpdatePipelineNotifications(input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	c.inc("UpdatePipelineNotifications")
	return c.svc.UpdatePipelineNotifications(input)
}

// UpdatePipelineNotificationsWithContext is a passthrough to the underlying UpdatePipelineNotificationsWithContext method.
// It will increment the count of requests made to UpdatePipelineNotifications.
func (c *ElasticTranscoder) UpdatePipelineNotificationsWithContext(ctx aws.Context, input *elastictranscoder.UpdatePipelineNotificationsInput, opts ...request.Option) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	c.inc("UpdatePipelineNotifications")
	return c.svc.UpdatePipelineNotificationsWithContext(ctx, input, opts...)
}

// UpdatePipelineStatusRequest is a passthrough to the underlying UpdatePipelineStatusRequest.
// It will increment the count of requests made to UpdatePipelineStatus.
func (c *ElasticTranscoder) UpdatePipelineStatusRequest(input *elastictranscoder.UpdatePipelineStatusInput) (req *request.Request, output *elastictranscoder.UpdatePipelineStatusOutput) {
	c.inc("UpdatePipelineStatus")
	return c.svc.UpdatePipelineStatusRequest(input)
}

// UpdatePipelineStatus is a passthrough to the underlying UpdatePipelineStatus method.
// It will increment the count of requests made to UpdatePipelineStatus.
func (c *ElasticTranscoder) UpdatePipelineStatus(input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	c.inc("UpdatePipelineStatus")
	return c.svc.UpdatePipelineStatus(input)
}

// UpdatePipelineStatusWithContext is a passthrough to the underlying UpdatePipelineStatusWithContext method.
// It will increment the count of requests made to UpdatePipelineStatus.
func (c *ElasticTranscoder) UpdatePipelineStatusWithContext(ctx aws.Context, input *elastictranscoder.UpdatePipelineStatusInput, opts ...request.Option) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	c.inc("UpdatePipelineStatus")
	return c.svc.UpdatePipelineStatusWithContext(ctx, input, opts...)
}
