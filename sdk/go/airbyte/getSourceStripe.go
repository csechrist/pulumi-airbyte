// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceStripe DataSource
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
//			_, err := airbyte.LookupSourceStripe(ctx, &airbyte.LookupSourceStripeArgs{
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
func LookupSourceStripe(ctx *pulumi.Context, args *LookupSourceStripeArgs, opts ...pulumi.InvokeOption) (*LookupSourceStripeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceStripeResult
	err := ctx.Invoke("airbyte:index/getSourceStripe:getSourceStripe", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceStripe.
type LookupSourceStripeArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceStripe.
type LookupSourceStripeResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceStripeOutput(ctx *pulumi.Context, args LookupSourceStripeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceStripeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceStripeResult, error) {
			args := v.(LookupSourceStripeArgs)
			r, err := LookupSourceStripe(ctx, &args, opts...)
			var s LookupSourceStripeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceStripeResultOutput)
}

// A collection of arguments for invoking getSourceStripe.
type LookupSourceStripeOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceStripeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceStripeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceStripe.
type LookupSourceStripeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceStripeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceStripeResult)(nil)).Elem()
}

func (o LookupSourceStripeResultOutput) ToLookupSourceStripeResultOutput() LookupSourceStripeResultOutput {
	return o
}

func (o LookupSourceStripeResultOutput) ToLookupSourceStripeResultOutputWithContext(ctx context.Context) LookupSourceStripeResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceStripeResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceStripeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceStripeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceStripeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceStripeResultOutput{})
}