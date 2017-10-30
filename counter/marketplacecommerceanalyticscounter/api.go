// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalyticscounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
)

// GenerateDataSetRequest is a passthrough to the underlying GenerateDataSetRequest.
// It will increment the count of requests made to GenerateDataSet.
func (c *MarketplaceCommerceAnalytics) GenerateDataSetRequest(input *marketplacecommerceanalytics.GenerateDataSetInput) (req *request.Request, output *marketplacecommerceanalytics.GenerateDataSetOutput) {
	c.inc("GenerateDataSet")
	return c.svc.GenerateDataSetRequest(input)
}

// GenerateDataSet is a passthrough to the underlying GenerateDataSet method.
// It will increment the count of requests made to GenerateDataSet.
func (c *MarketplaceCommerceAnalytics) GenerateDataSet(input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	c.inc("GenerateDataSet")
	return c.svc.GenerateDataSet(input)
}

// GenerateDataSetWithContext is a passthrough to the underlying GenerateDataSetWithContext method.
// It will increment the count of requests made to GenerateDataSet.
func (c *MarketplaceCommerceAnalytics) GenerateDataSetWithContext(ctx aws.Context, input *marketplacecommerceanalytics.GenerateDataSetInput, opts ...request.Option) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	c.inc("GenerateDataSet")
	return c.svc.GenerateDataSetWithContext(ctx, input, opts...)
}

// StartSupportDataExportRequest is a passthrough to the underlying StartSupportDataExportRequest.
// It will increment the count of requests made to StartSupportDataExport.
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportRequest(input *marketplacecommerceanalytics.StartSupportDataExportInput) (req *request.Request, output *marketplacecommerceanalytics.StartSupportDataExportOutput) {
	c.inc("StartSupportDataExport")
	return c.svc.StartSupportDataExportRequest(input)
}

// StartSupportDataExport is a passthrough to the underlying StartSupportDataExport method.
// It will increment the count of requests made to StartSupportDataExport.
func (c *MarketplaceCommerceAnalytics) StartSupportDataExport(input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	c.inc("StartSupportDataExport")
	return c.svc.StartSupportDataExport(input)
}

// StartSupportDataExportWithContext is a passthrough to the underlying StartSupportDataExportWithContext method.
// It will increment the count of requests made to StartSupportDataExport.
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportWithContext(ctx aws.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput, opts ...request.Option) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	c.inc("StartSupportDataExport")
	return c.svc.StartSupportDataExportWithContext(ctx, input, opts...)
}
