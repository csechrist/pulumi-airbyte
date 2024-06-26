// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourcePostmarkapp DataSource
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
//			_, err := airbyte.LookupSourcePostmarkapp(ctx, &airbyte.LookupSourcePostmarkappArgs{
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
func LookupSourcePostmarkapp(ctx *pulumi.Context, args *LookupSourcePostmarkappArgs, opts ...pulumi.InvokeOption) (*LookupSourcePostmarkappResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourcePostmarkappResult
	err := ctx.Invoke("airbyte:index/getSourcePostmarkapp:getSourcePostmarkapp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourcePostmarkapp.
type LookupSourcePostmarkappArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourcePostmarkapp.
type LookupSourcePostmarkappResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourcePostmarkappOutput(ctx *pulumi.Context, args LookupSourcePostmarkappOutputArgs, opts ...pulumi.InvokeOption) LookupSourcePostmarkappResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourcePostmarkappResult, error) {
			args := v.(LookupSourcePostmarkappArgs)
			r, err := LookupSourcePostmarkapp(ctx, &args, opts...)
			var s LookupSourcePostmarkappResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourcePostmarkappResultOutput)
}

// A collection of arguments for invoking getSourcePostmarkapp.
type LookupSourcePostmarkappOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourcePostmarkappOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePostmarkappArgs)(nil)).Elem()
}

// A collection of values returned by getSourcePostmarkapp.
type LookupSourcePostmarkappResultOutput struct{ *pulumi.OutputState }

func (LookupSourcePostmarkappResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourcePostmarkappResult)(nil)).Elem()
}

func (o LookupSourcePostmarkappResultOutput) ToLookupSourcePostmarkappResultOutput() LookupSourcePostmarkappResultOutput {
	return o
}

func (o LookupSourcePostmarkappResultOutput) ToLookupSourcePostmarkappResultOutputWithContext(ctx context.Context) LookupSourcePostmarkappResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourcePostmarkappResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourcePostmarkappResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourcePostmarkappResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourcePostmarkappResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourcePostmarkappResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourcePostmarkappResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourcePostmarkappResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourcePostmarkappResultOutput{})
}
