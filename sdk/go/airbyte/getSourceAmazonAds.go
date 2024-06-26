// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceAmazonAds DataSource
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
//			_, err := airbyte.LookupSourceAmazonAds(ctx, &airbyte.LookupSourceAmazonAdsArgs{
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
func LookupSourceAmazonAds(ctx *pulumi.Context, args *LookupSourceAmazonAdsArgs, opts ...pulumi.InvokeOption) (*LookupSourceAmazonAdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceAmazonAdsResult
	err := ctx.Invoke("airbyte:index/getSourceAmazonAds:getSourceAmazonAds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceAmazonAds.
type LookupSourceAmazonAdsArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceAmazonAds.
type LookupSourceAmazonAdsResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceAmazonAdsOutput(ctx *pulumi.Context, args LookupSourceAmazonAdsOutputArgs, opts ...pulumi.InvokeOption) LookupSourceAmazonAdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceAmazonAdsResult, error) {
			args := v.(LookupSourceAmazonAdsArgs)
			r, err := LookupSourceAmazonAds(ctx, &args, opts...)
			var s LookupSourceAmazonAdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceAmazonAdsResultOutput)
}

// A collection of arguments for invoking getSourceAmazonAds.
type LookupSourceAmazonAdsOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceAmazonAdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonAdsArgs)(nil)).Elem()
}

// A collection of values returned by getSourceAmazonAds.
type LookupSourceAmazonAdsResultOutput struct{ *pulumi.OutputState }

func (LookupSourceAmazonAdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceAmazonAdsResult)(nil)).Elem()
}

func (o LookupSourceAmazonAdsResultOutput) ToLookupSourceAmazonAdsResultOutput() LookupSourceAmazonAdsResultOutput {
	return o
}

func (o LookupSourceAmazonAdsResultOutput) ToLookupSourceAmazonAdsResultOutputWithContext(ctx context.Context) LookupSourceAmazonAdsResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceAmazonAdsResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceAmazonAdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonAdsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonAdsResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonAdsResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceAmazonAdsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceAmazonAdsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceAmazonAdsResultOutput{})
}
