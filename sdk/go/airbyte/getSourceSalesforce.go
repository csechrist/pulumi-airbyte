// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSalesforce DataSource
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
//			_, err := airbyte.LookupSourceSalesforce(ctx, &airbyte.LookupSourceSalesforceArgs{
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
func LookupSourceSalesforce(ctx *pulumi.Context, args *LookupSourceSalesforceArgs, opts ...pulumi.InvokeOption) (*LookupSourceSalesforceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSalesforceResult
	err := ctx.Invoke("airbyte:index/getSourceSalesforce:getSourceSalesforce", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSalesforce.
type LookupSourceSalesforceArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSalesforce.
type LookupSourceSalesforceResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSalesforceOutput(ctx *pulumi.Context, args LookupSourceSalesforceOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSalesforceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSalesforceResult, error) {
			args := v.(LookupSourceSalesforceArgs)
			r, err := LookupSourceSalesforce(ctx, &args, opts...)
			var s LookupSourceSalesforceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSalesforceResultOutput)
}

// A collection of arguments for invoking getSourceSalesforce.
type LookupSourceSalesforceOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSalesforceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSalesforceArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSalesforce.
type LookupSourceSalesforceResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSalesforceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSalesforceResult)(nil)).Elem()
}

func (o LookupSourceSalesforceResultOutput) ToLookupSourceSalesforceResultOutput() LookupSourceSalesforceResultOutput {
	return o
}

func (o LookupSourceSalesforceResultOutput) ToLookupSourceSalesforceResultOutputWithContext(ctx context.Context) LookupSourceSalesforceResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceSalesforceResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSalesforceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSalesforceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSalesforceResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSalesforceResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSalesforceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSalesforceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSalesforceResultOutput{})
}