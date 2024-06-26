// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationGcs DataSource
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
//			_, err := airbyte.LookupDestinationGcs(ctx, &airbyte.LookupDestinationGcsArgs{
//				DestinationId: "...my_destination_id...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDestinationGcs(ctx *pulumi.Context, args *LookupDestinationGcsArgs, opts ...pulumi.InvokeOption) (*LookupDestinationGcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationGcsResult
	err := ctx.Invoke("airbyte:index/getDestinationGcs:getDestinationGcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationGcs.
type LookupDestinationGcsArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationGcs.
type LookupDestinationGcsResult struct {
	// The values required to configure the destination. Parsed as JSON.
	Configuration   string `pulumi:"configuration"`
	DestinationId   string `pulumi:"destinationId"`
	DestinationType string `pulumi:"destinationType"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationGcsOutput(ctx *pulumi.Context, args LookupDestinationGcsOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationGcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationGcsResult, error) {
			args := v.(LookupDestinationGcsArgs)
			r, err := LookupDestinationGcs(ctx, &args, opts...)
			var s LookupDestinationGcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationGcsResultOutput)
}

// A collection of arguments for invoking getDestinationGcs.
type LookupDestinationGcsOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationGcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationGcsArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationGcs.
type LookupDestinationGcsResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationGcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationGcsResult)(nil)).Elem()
}

func (o LookupDestinationGcsResultOutput) ToLookupDestinationGcsResultOutput() LookupDestinationGcsResultOutput {
	return o
}

func (o LookupDestinationGcsResultOutput) ToLookupDestinationGcsResultOutputWithContext(ctx context.Context) LookupDestinationGcsResultOutput {
	return o
}

// The values required to configure the destination. Parsed as JSON.
func (o LookupDestinationGcsResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o LookupDestinationGcsResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

func (o LookupDestinationGcsResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationGcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationGcsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationGcsResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationGcsResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationGcsResultOutput{})
}
