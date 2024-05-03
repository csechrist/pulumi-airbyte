// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceLemlist DataSource
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
//			_, err := airbyte.LookupSourceLemlist(ctx, &airbyte.LookupSourceLemlistArgs{
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
func LookupSourceLemlist(ctx *pulumi.Context, args *LookupSourceLemlistArgs, opts ...pulumi.InvokeOption) (*LookupSourceLemlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceLemlistResult
	err := ctx.Invoke("airbyte:index/getSourceLemlist:getSourceLemlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceLemlist.
type LookupSourceLemlistArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceLemlist.
type LookupSourceLemlistResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceLemlistOutput(ctx *pulumi.Context, args LookupSourceLemlistOutputArgs, opts ...pulumi.InvokeOption) LookupSourceLemlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceLemlistResult, error) {
			args := v.(LookupSourceLemlistArgs)
			r, err := LookupSourceLemlist(ctx, &args, opts...)
			var s LookupSourceLemlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceLemlistResultOutput)
}

// A collection of arguments for invoking getSourceLemlist.
type LookupSourceLemlistOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceLemlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLemlistArgs)(nil)).Elem()
}

// A collection of values returned by getSourceLemlist.
type LookupSourceLemlistResultOutput struct{ *pulumi.OutputState }

func (LookupSourceLemlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceLemlistResult)(nil)).Elem()
}

func (o LookupSourceLemlistResultOutput) ToLookupSourceLemlistResultOutput() LookupSourceLemlistResultOutput {
	return o
}

func (o LookupSourceLemlistResultOutput) ToLookupSourceLemlistResultOutputWithContext(ctx context.Context) LookupSourceLemlistResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceLemlistResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceLemlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceLemlistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceLemlistResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceLemlistResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceLemlistResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceLemlistResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceLemlistResultOutput{})
}