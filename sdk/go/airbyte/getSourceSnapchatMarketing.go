// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSnapchatMarketing DataSource
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
//			_, err := airbyte.LookupSourceSnapchatMarketing(ctx, &airbyte.LookupSourceSnapchatMarketingArgs{
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
func LookupSourceSnapchatMarketing(ctx *pulumi.Context, args *LookupSourceSnapchatMarketingArgs, opts ...pulumi.InvokeOption) (*LookupSourceSnapchatMarketingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSnapchatMarketingResult
	err := ctx.Invoke("airbyte:index/getSourceSnapchatMarketing:getSourceSnapchatMarketing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSnapchatMarketing.
type LookupSourceSnapchatMarketingArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSnapchatMarketing.
type LookupSourceSnapchatMarketingResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSnapchatMarketingOutput(ctx *pulumi.Context, args LookupSourceSnapchatMarketingOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSnapchatMarketingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSnapchatMarketingResult, error) {
			args := v.(LookupSourceSnapchatMarketingArgs)
			r, err := LookupSourceSnapchatMarketing(ctx, &args, opts...)
			var s LookupSourceSnapchatMarketingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSnapchatMarketingResultOutput)
}

// A collection of arguments for invoking getSourceSnapchatMarketing.
type LookupSourceSnapchatMarketingOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSnapchatMarketingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSnapchatMarketingArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSnapchatMarketing.
type LookupSourceSnapchatMarketingResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSnapchatMarketingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSnapchatMarketingResult)(nil)).Elem()
}

func (o LookupSourceSnapchatMarketingResultOutput) ToLookupSourceSnapchatMarketingResultOutput() LookupSourceSnapchatMarketingResultOutput {
	return o
}

func (o LookupSourceSnapchatMarketingResultOutput) ToLookupSourceSnapchatMarketingResultOutputWithContext(ctx context.Context) LookupSourceSnapchatMarketingResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceSnapchatMarketingResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSnapchatMarketingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSnapchatMarketingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSnapchatMarketingResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSnapchatMarketingResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSnapchatMarketingResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSnapchatMarketingResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSnapchatMarketingResultOutput{})
}