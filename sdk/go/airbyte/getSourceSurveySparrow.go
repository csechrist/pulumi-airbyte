// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSurveySparrow DataSource
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
//			_, err := airbyte.LookupSourceSurveySparrow(ctx, &airbyte.LookupSourceSurveySparrowArgs{
//				SourceId: "...my_source_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSourceSurveySparrow(ctx *pulumi.Context, args *LookupSourceSurveySparrowArgs, opts ...pulumi.InvokeOption) (*LookupSourceSurveySparrowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSurveySparrowResult
	err := ctx.Invoke("airbyte:index/getSourceSurveySparrow:getSourceSurveySparrow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSurveySparrow.
type LookupSourceSurveySparrowArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSurveySparrow.
type LookupSourceSurveySparrowResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSurveySparrowOutput(ctx *pulumi.Context, args LookupSourceSurveySparrowOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSurveySparrowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSurveySparrowResult, error) {
			args := v.(LookupSourceSurveySparrowArgs)
			r, err := LookupSourceSurveySparrow(ctx, &args, opts...)
			var s LookupSourceSurveySparrowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSurveySparrowResultOutput)
}

// A collection of arguments for invoking getSourceSurveySparrow.
type LookupSourceSurveySparrowOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSurveySparrowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSurveySparrowArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSurveySparrow.
type LookupSourceSurveySparrowResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSurveySparrowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSurveySparrowResult)(nil)).Elem()
}

func (o LookupSourceSurveySparrowResultOutput) ToLookupSourceSurveySparrowResultOutput() LookupSourceSurveySparrowResultOutput {
	return o
}

func (o LookupSourceSurveySparrowResultOutput) ToLookupSourceSurveySparrowResultOutputWithContext(ctx context.Context) LookupSourceSurveySparrowResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceSurveySparrowResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSurveySparrowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSurveySparrowResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSurveySparrowResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSurveySparrowResultOutput{})
}
