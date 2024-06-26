// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceOutbrainAmplify DataSource
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
//			_, err := airbyte.LookupSourceOutbrainAmplify(ctx, &airbyte.LookupSourceOutbrainAmplifyArgs{
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
func LookupSourceOutbrainAmplify(ctx *pulumi.Context, args *LookupSourceOutbrainAmplifyArgs, opts ...pulumi.InvokeOption) (*LookupSourceOutbrainAmplifyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceOutbrainAmplifyResult
	err := ctx.Invoke("airbyte:index/getSourceOutbrainAmplify:getSourceOutbrainAmplify", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceOutbrainAmplifyOutput(ctx *pulumi.Context, args LookupSourceOutbrainAmplifyOutputArgs, opts ...pulumi.InvokeOption) LookupSourceOutbrainAmplifyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceOutbrainAmplifyResult, error) {
			args := v.(LookupSourceOutbrainAmplifyArgs)
			r, err := LookupSourceOutbrainAmplify(ctx, &args, opts...)
			var s LookupSourceOutbrainAmplifyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceOutbrainAmplifyResultOutput)
}

// A collection of arguments for invoking getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceOutbrainAmplifyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutbrainAmplifyArgs)(nil)).Elem()
}

// A collection of values returned by getSourceOutbrainAmplify.
type LookupSourceOutbrainAmplifyResultOutput struct{ *pulumi.OutputState }

func (LookupSourceOutbrainAmplifyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceOutbrainAmplifyResult)(nil)).Elem()
}

func (o LookupSourceOutbrainAmplifyResultOutput) ToLookupSourceOutbrainAmplifyResultOutput() LookupSourceOutbrainAmplifyResultOutput {
	return o
}

func (o LookupSourceOutbrainAmplifyResultOutput) ToLookupSourceOutbrainAmplifyResultOutputWithContext(ctx context.Context) LookupSourceOutbrainAmplifyResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceOutbrainAmplifyResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceOutbrainAmplifyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceOutbrainAmplifyResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceOutbrainAmplifyResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceOutbrainAmplifyResultOutput{})
}
