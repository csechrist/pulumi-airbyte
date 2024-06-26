// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceUsCensus DataSource
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
//			_, err := airbyte.LookupSourceUsCensus(ctx, &airbyte.LookupSourceUsCensusArgs{
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
func LookupSourceUsCensus(ctx *pulumi.Context, args *LookupSourceUsCensusArgs, opts ...pulumi.InvokeOption) (*LookupSourceUsCensusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceUsCensusResult
	err := ctx.Invoke("airbyte:index/getSourceUsCensus:getSourceUsCensus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceUsCensus.
type LookupSourceUsCensusArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceUsCensus.
type LookupSourceUsCensusResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceUsCensusOutput(ctx *pulumi.Context, args LookupSourceUsCensusOutputArgs, opts ...pulumi.InvokeOption) LookupSourceUsCensusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceUsCensusResult, error) {
			args := v.(LookupSourceUsCensusArgs)
			r, err := LookupSourceUsCensus(ctx, &args, opts...)
			var s LookupSourceUsCensusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceUsCensusResultOutput)
}

// A collection of arguments for invoking getSourceUsCensus.
type LookupSourceUsCensusOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceUsCensusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceUsCensusArgs)(nil)).Elem()
}

// A collection of values returned by getSourceUsCensus.
type LookupSourceUsCensusResultOutput struct{ *pulumi.OutputState }

func (LookupSourceUsCensusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceUsCensusResult)(nil)).Elem()
}

func (o LookupSourceUsCensusResultOutput) ToLookupSourceUsCensusResultOutput() LookupSourceUsCensusResultOutput {
	return o
}

func (o LookupSourceUsCensusResultOutput) ToLookupSourceUsCensusResultOutputWithContext(ctx context.Context) LookupSourceUsCensusResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceUsCensusResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceUsCensusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceUsCensusResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceUsCensusResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceUsCensusResultOutput{})
}
