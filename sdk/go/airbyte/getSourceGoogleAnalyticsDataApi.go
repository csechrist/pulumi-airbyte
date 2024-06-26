// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceGoogleAnalyticsDataAPI DataSource
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
//			_, err := airbyte.LookupSourceGoogleAnalyticsDataApi(ctx, &airbyte.LookupSourceGoogleAnalyticsDataApiArgs{
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
func LookupSourceGoogleAnalyticsDataApi(ctx *pulumi.Context, args *LookupSourceGoogleAnalyticsDataApiArgs, opts ...pulumi.InvokeOption) (*LookupSourceGoogleAnalyticsDataApiResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceGoogleAnalyticsDataApiResult
	err := ctx.Invoke("airbyte:index/getSourceGoogleAnalyticsDataApi:getSourceGoogleAnalyticsDataApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceGoogleAnalyticsDataApi.
type LookupSourceGoogleAnalyticsDataApiArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceGoogleAnalyticsDataApi.
type LookupSourceGoogleAnalyticsDataApiResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceGoogleAnalyticsDataApiOutput(ctx *pulumi.Context, args LookupSourceGoogleAnalyticsDataApiOutputArgs, opts ...pulumi.InvokeOption) LookupSourceGoogleAnalyticsDataApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceGoogleAnalyticsDataApiResult, error) {
			args := v.(LookupSourceGoogleAnalyticsDataApiArgs)
			r, err := LookupSourceGoogleAnalyticsDataApi(ctx, &args, opts...)
			var s LookupSourceGoogleAnalyticsDataApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceGoogleAnalyticsDataApiResultOutput)
}

// A collection of arguments for invoking getSourceGoogleAnalyticsDataApi.
type LookupSourceGoogleAnalyticsDataApiOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceGoogleAnalyticsDataApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleAnalyticsDataApiArgs)(nil)).Elem()
}

// A collection of values returned by getSourceGoogleAnalyticsDataApi.
type LookupSourceGoogleAnalyticsDataApiResultOutput struct{ *pulumi.OutputState }

func (LookupSourceGoogleAnalyticsDataApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceGoogleAnalyticsDataApiResult)(nil)).Elem()
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) ToLookupSourceGoogleAnalyticsDataApiResultOutput() LookupSourceGoogleAnalyticsDataApiResultOutput {
	return o
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) ToLookupSourceGoogleAnalyticsDataApiResultOutputWithContext(ctx context.Context) LookupSourceGoogleAnalyticsDataApiResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceGoogleAnalyticsDataApiResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceGoogleAnalyticsDataApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceGoogleAnalyticsDataApiResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceGoogleAnalyticsDataApiResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceGoogleAnalyticsDataApiResultOutput{})
}
