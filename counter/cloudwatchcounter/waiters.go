// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// WaitUntilAlarmExists calls WaitUntilAlarmExistsWithContext with aws.BackgroundContext().
func (c *CloudWatch) WaitUntilAlarmExists(input *cloudwatch.DescribeAlarmsInput) error {
	return c.WaitUntilAlarmExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAlarmExistsWithContext calls the underlying client method with a request option that
// will count DescribeAlarms requests.
func (c *CloudWatch) WaitUntilAlarmExistsWithContext(ctx aws.Context, input *cloudwatch.DescribeAlarmsInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeAlarms")))
	return c.svc.WaitUntilAlarmExistsWithContext(ctx, input, opts...)
}
