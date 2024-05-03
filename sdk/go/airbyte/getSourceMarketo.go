// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceMarketo DataSource
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
//			_, err := airbyte.LookupSourceMarketo(ctx, &airbyte.LookupSourceMarketoArgs{
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
func LookupSourceMarketo(ctx *pulumi.Context, args *LookupSourceMarketoArgs, opts ...pulumi.InvokeOption) (*LookupSourceMarketoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceMarketoResult
	err := ctx.Invoke("airbyte:index/getSourceMarketo:getSourceMarketo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceMarketo.
type LookupSourceMarketoArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceMarketo.
type LookupSourceMarketoResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceMarketoOutput(ctx *pulumi.Context, args LookupSourceMarketoOutputArgs, opts ...pulumi.InvokeOption) LookupSourceMarketoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceMarketoResult, error) {
			args := v.(LookupSourceMarketoArgs)
			r, err := LookupSourceMarketo(ctx, &args, opts...)
			var s LookupSourceMarketoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceMarketoResultOutput)
}

// A collection of arguments for invoking getSourceMarketo.
type LookupSourceMarketoOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceMarketoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMarketoArgs)(nil)).Elem()
}

// A collection of values returned by getSourceMarketo.
type LookupSourceMarketoResultOutput struct{ *pulumi.OutputState }

func (LookupSourceMarketoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceMarketoResult)(nil)).Elem()
}

func (o LookupSourceMarketoResultOutput) ToLookupSourceMarketoResultOutput() LookupSourceMarketoResultOutput {
	return o
}

func (o LookupSourceMarketoResultOutput) ToLookupSourceMarketoResultOutputWithContext(ctx context.Context) LookupSourceMarketoResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceMarketoResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceMarketoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceMarketoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceMarketoResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceMarketoResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceMarketoResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceMarketoResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceMarketoResultOutput{})
}