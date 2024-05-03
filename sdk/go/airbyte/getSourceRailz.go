// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceRailz DataSource
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
//			_, err := airbyte.LookupSourceRailz(ctx, &airbyte.LookupSourceRailzArgs{
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
func LookupSourceRailz(ctx *pulumi.Context, args *LookupSourceRailzArgs, opts ...pulumi.InvokeOption) (*LookupSourceRailzResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceRailzResult
	err := ctx.Invoke("airbyte:index/getSourceRailz:getSourceRailz", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceRailz.
type LookupSourceRailzArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceRailz.
type LookupSourceRailzResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceRailzOutput(ctx *pulumi.Context, args LookupSourceRailzOutputArgs, opts ...pulumi.InvokeOption) LookupSourceRailzResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceRailzResult, error) {
			args := v.(LookupSourceRailzArgs)
			r, err := LookupSourceRailz(ctx, &args, opts...)
			var s LookupSourceRailzResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceRailzResultOutput)
}

// A collection of arguments for invoking getSourceRailz.
type LookupSourceRailzOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceRailzOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRailzArgs)(nil)).Elem()
}

// A collection of values returned by getSourceRailz.
type LookupSourceRailzResultOutput struct{ *pulumi.OutputState }

func (LookupSourceRailzResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceRailzResult)(nil)).Elem()
}

func (o LookupSourceRailzResultOutput) ToLookupSourceRailzResultOutput() LookupSourceRailzResultOutput {
	return o
}

func (o LookupSourceRailzResultOutput) ToLookupSourceRailzResultOutputWithContext(ctx context.Context) LookupSourceRailzResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceRailzResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceRailzResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceRailzResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceRailzResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceRailzResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceRailzResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceRailzResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceRailzResultOutput{})
}