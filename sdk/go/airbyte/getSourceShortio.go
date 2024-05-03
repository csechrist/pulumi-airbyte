// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceShortio DataSource
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
//			_, err := airbyte.LookupSourceShortio(ctx, &airbyte.LookupSourceShortioArgs{
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
func LookupSourceShortio(ctx *pulumi.Context, args *LookupSourceShortioArgs, opts ...pulumi.InvokeOption) (*LookupSourceShortioResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceShortioResult
	err := ctx.Invoke("airbyte:index/getSourceShortio:getSourceShortio", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceShortio.
type LookupSourceShortioArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceShortio.
type LookupSourceShortioResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceShortioOutput(ctx *pulumi.Context, args LookupSourceShortioOutputArgs, opts ...pulumi.InvokeOption) LookupSourceShortioResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceShortioResult, error) {
			args := v.(LookupSourceShortioArgs)
			r, err := LookupSourceShortio(ctx, &args, opts...)
			var s LookupSourceShortioResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceShortioResultOutput)
}

// A collection of arguments for invoking getSourceShortio.
type LookupSourceShortioOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceShortioOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceShortioArgs)(nil)).Elem()
}

// A collection of values returned by getSourceShortio.
type LookupSourceShortioResultOutput struct{ *pulumi.OutputState }

func (LookupSourceShortioResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceShortioResult)(nil)).Elem()
}

func (o LookupSourceShortioResultOutput) ToLookupSourceShortioResultOutput() LookupSourceShortioResultOutput {
	return o
}

func (o LookupSourceShortioResultOutput) ToLookupSourceShortioResultOutputWithContext(ctx context.Context) LookupSourceShortioResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceShortioResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceShortioResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceShortioResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceShortioResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceShortioResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceShortioResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceShortioResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceShortioResultOutput{})
}