// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSendinblue DataSource
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
//			_, err := airbyte.LookupSourceSendinblue(ctx, &airbyte.LookupSourceSendinblueArgs{
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
func LookupSourceSendinblue(ctx *pulumi.Context, args *LookupSourceSendinblueArgs, opts ...pulumi.InvokeOption) (*LookupSourceSendinblueResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSendinblueResult
	err := ctx.Invoke("airbyte:index/getSourceSendinblue:getSourceSendinblue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSendinblue.
type LookupSourceSendinblueArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSendinblue.
type LookupSourceSendinblueResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSendinblueOutput(ctx *pulumi.Context, args LookupSourceSendinblueOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSendinblueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSendinblueResult, error) {
			args := v.(LookupSourceSendinblueArgs)
			r, err := LookupSourceSendinblue(ctx, &args, opts...)
			var s LookupSourceSendinblueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSendinblueResultOutput)
}

// A collection of arguments for invoking getSourceSendinblue.
type LookupSourceSendinblueOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSendinblueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSendinblueArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSendinblue.
type LookupSourceSendinblueResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSendinblueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSendinblueResult)(nil)).Elem()
}

func (o LookupSourceSendinblueResultOutput) ToLookupSourceSendinblueResultOutput() LookupSourceSendinblueResultOutput {
	return o
}

func (o LookupSourceSendinblueResultOutput) ToLookupSourceSendinblueResultOutputWithContext(ctx context.Context) LookupSourceSendinblueResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceSendinblueResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSendinblueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSendinblueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSendinblueResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSendinblueResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSendinblueResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSendinblueResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSendinblueResultOutput{})
}