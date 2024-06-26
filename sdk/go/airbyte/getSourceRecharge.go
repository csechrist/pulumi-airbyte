// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceRecharge DataSource
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
//			_, err := airbyte.LookupSourceRecharge(ctx, &airbyte.LookupSourceRechargeArgs{
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
func LookupSourceRecharge(ctx *pulumi.Context, args *LookupSourceRechargeArgs, opts ...pulumi.InvokeOption) (*LookupSourceRechargeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceRechargeResult
	err := ctx.Invoke("airbyte:index/getSourceRecharge:getSourceRecharge", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRecharge.
type LookupSourceRechargeArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRecharge.
type LookupSourceRechargeResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceRechargeOutput(ctx *pulumi.Context, args LookupSourceRechargeOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRechargeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRechargeResult, error) {
			args := v.(LookupSourceRechargeArgs)
			r, err := LookupSourceRecharge(ctx, &args, opts...)
			var s LookupSourceRechargeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRechargeResultOutput)
}

// A collection of arguments for invoking getSourceRecharge.
type LookupSourceRechargeOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceRechargeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRechargeArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRecharge.
type LookupSourceRechargeResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRechargeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRechargeResult)(nil)).Elem()
}

func (o LookupSourceRechargeResultOutput) ToLookupSourceRechargeResultOutput() LookupSourceRechargeResultOutput {
	return o
}

func (o LookupSourceRechargeResultOutput) ToLookupSourceRechargeResultOutputWithContext(ctx context.Context) LookupSourceRechargeResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceRechargeResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRechargeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRechargeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceRechargeResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRechargeResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceRechargeResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRechargeResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRechargeResultOutput{})
}
