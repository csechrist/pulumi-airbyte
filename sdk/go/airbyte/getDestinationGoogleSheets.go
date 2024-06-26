// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationGoogleSheets DataSource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airbyte.LookupDestinationGoogleSheets(ctx, &airbyte.LookupDestinationGoogleSheetsArgs{
//				DestinationId: "...my_destination_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDestinationGoogleSheets(ctx *pulumi.Context, args *LookupDestinationGoogleSheetsArgs, opts ...pulumi.InvokeOption) (*LookupDestinationGoogleSheetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationGoogleSheetsResult
	err := ctx.Invoke("airbyte:index/getDestinationGoogleSheets:getDestinationGoogleSheets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationGoogleSheets.
type LookupDestinationGoogleSheetsArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationGoogleSheets.
type LookupDestinationGoogleSheetsResult struct {
	// The values required to configure the destination. Parsed as JSON.
	Configuration   string `pulumi:"configuration"`
	DestinationId   string `pulumi:"destinationId"`
	DestinationType string `pulumi:"destinationType"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationGoogleSheetsOutput(ctx *pulumi.Context, args LookupDestinationGoogleSheetsOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationGoogleSheetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationGoogleSheetsResult, error) {
			args := v.(LookupDestinationGoogleSheetsArgs)
			r, err := LookupDestinationGoogleSheets(ctx, &args, opts...)
			var s LookupDestinationGoogleSheetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationGoogleSheetsResultOutput)
}

// A collection of arguments for invoking getDestinationGoogleSheets.
type LookupDestinationGoogleSheetsOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationGoogleSheetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationGoogleSheetsArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationGoogleSheets.
type LookupDestinationGoogleSheetsResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationGoogleSheetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationGoogleSheetsResult)(nil)).Elem()
}

func (o LookupDestinationGoogleSheetsResultOutput) ToLookupDestinationGoogleSheetsResultOutput() LookupDestinationGoogleSheetsResultOutput {
	return o
}

func (o LookupDestinationGoogleSheetsResultOutput) ToLookupDestinationGoogleSheetsResultOutputWithContext(ctx context.Context) LookupDestinationGoogleSheetsResultOutput {
	return o
}

// The values required to configure the destination. Parsed as JSON.
func (o LookupDestinationGoogleSheetsResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o LookupDestinationGoogleSheetsResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

func (o LookupDestinationGoogleSheetsResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationGoogleSheetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationGoogleSheetsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationGoogleSheetsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGoogleSheetsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationGoogleSheetsResultOutput{})
}
