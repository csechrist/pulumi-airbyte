// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceVantage DataSource
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
//			_, err := airbyte.LookupSourceVantage(ctx, &airbyte.LookupSourceVantageArgs{
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
func LookupSourceVantage(ctx *pulumi.Context, args *LookupSourceVantageArgs, opts ...pulumi.InvokeOption) (*LookupSourceVantageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceVantageResult
	err := ctx.Invoke("airbyte:index/getSourceVantage:getSourceVantage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceVantage.
type LookupSourceVantageArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceVantage.
type LookupSourceVantageResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceVantageOutput(ctx *pulumi.Context, args LookupSourceVantageOutputArgs, opts ...pulumi.InvokeOption) LookupSourceVantageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceVantageResult, error) {
			args := v.(LookupSourceVantageArgs)
			r, err := LookupSourceVantage(ctx, &args, opts...)
			var s LookupSourceVantageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceVantageResultOutput)
}

// A collection of arguments for invoking getSourceVantage.
type LookupSourceVantageOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceVantageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceVantageArgs)(nil)).Elem()
}

// A collection of values returned by getSourceVantage.
type LookupSourceVantageResultOutput struct{ *pulumi.OutputState }

func (LookupSourceVantageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceVantageResult)(nil)).Elem()
}

func (o LookupSourceVantageResultOutput) ToLookupSourceVantageResultOutput() LookupSourceVantageResultOutput {
	return o
}

func (o LookupSourceVantageResultOutput) ToLookupSourceVantageResultOutputWithContext(ctx context.Context) LookupSourceVantageResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceVantageResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceVantageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceVantageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceVantageResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceVantageResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceVantageResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceVantageResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceVantageResultOutput{})
}