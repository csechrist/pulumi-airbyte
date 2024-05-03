// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceNytimes DataSource
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
//			_, err := airbyte.LookupSourceNytimes(ctx, &airbyte.LookupSourceNytimesArgs{
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
func LookupSourceNytimes(ctx *pulumi.Context, args *LookupSourceNytimesArgs, opts ...pulumi.InvokeOption) (*LookupSourceNytimesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceNytimesResult
	err := ctx.Invoke("airbyte:index/getSourceNytimes:getSourceNytimes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceNytimes.
type LookupSourceNytimesArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceNytimes.
type LookupSourceNytimesResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceNytimesOutput(ctx *pulumi.Context, args LookupSourceNytimesOutputArgs, opts ...pulumi.InvokeOption) LookupSourceNytimesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceNytimesResult, error) {
			args := v.(LookupSourceNytimesArgs)
			r, err := LookupSourceNytimes(ctx, &args, opts...)
			var s LookupSourceNytimesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceNytimesResultOutput)
}

// A collection of arguments for invoking getSourceNytimes.
type LookupSourceNytimesOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceNytimesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNytimesArgs)(nil)).Elem()
}

// A collection of values returned by getSourceNytimes.
type LookupSourceNytimesResultOutput struct{ *pulumi.OutputState }

func (LookupSourceNytimesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceNytimesResult)(nil)).Elem()
}

func (o LookupSourceNytimesResultOutput) ToLookupSourceNytimesResultOutput() LookupSourceNytimesResultOutput {
	return o
}

func (o LookupSourceNytimesResultOutput) ToLookupSourceNytimesResultOutputWithContext(ctx context.Context) LookupSourceNytimesResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceNytimesResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceNytimesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceNytimesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceNytimesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceNytimesResultOutput{})
}