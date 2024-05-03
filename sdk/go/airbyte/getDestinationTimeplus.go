// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DestinationTimeplus DataSource
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
//			_, err := airbyte.LookupDestinationTimeplus(ctx, &airbyte.LookupDestinationTimeplusArgs{
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
func LookupDestinationTimeplus(ctx *pulumi.Context, args *LookupDestinationTimeplusArgs, opts ...pulumi.InvokeOption) (*LookupDestinationTimeplusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationTimeplusResult
	err := ctx.Invoke("airbyte:index/getDestinationTimeplus:getDestinationTimeplus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDestinationTimeplus.
type LookupDestinationTimeplusArgs struct {
	DestinationId string `pulumi:"destinationId"`
}

// A collection of values returned by getDestinationTimeplus.
type LookupDestinationTimeplusResult struct {
	// The values required to configure the destination. Parsed as JSON.
	Configuration   string `pulumi:"configuration"`
	DestinationId   string `pulumi:"destinationId"`
	DestinationType string `pulumi:"destinationType"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupDestinationTimeplusOutput(ctx *pulumi.Context, args LookupDestinationTimeplusOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationTimeplusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationTimeplusResult, error) {
			args := v.(LookupDestinationTimeplusArgs)
			r, err := LookupDestinationTimeplus(ctx, &args, opts...)
			var s LookupDestinationTimeplusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationTimeplusResultOutput)
}

// A collection of arguments for invoking getDestinationTimeplus.
type LookupDestinationTimeplusOutputArgs struct {
	DestinationId pulumi.StringInput `pulumi:"destinationId"`
}

func (LookupDestinationTimeplusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationTimeplusArgs)(nil)).Elem()
}

// A collection of values returned by getDestinationTimeplus.
type LookupDestinationTimeplusResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationTimeplusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationTimeplusResult)(nil)).Elem()
}

func (o LookupDestinationTimeplusResultOutput) ToLookupDestinationTimeplusResultOutput() LookupDestinationTimeplusResultOutput {
	return o
}

func (o LookupDestinationTimeplusResultOutput) ToLookupDestinationTimeplusResultOutputWithContext(ctx context.Context) LookupDestinationTimeplusResultOutput {
	return o
}

// The values required to configure the destination. Parsed as JSON.
func (o LookupDestinationTimeplusResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o LookupDestinationTimeplusResultOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.DestinationId }).(pulumi.StringOutput)
}

func (o LookupDestinationTimeplusResultOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.DestinationType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDestinationTimeplusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDestinationTimeplusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDestinationTimeplusResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDestinationTimeplusResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationTimeplusResultOutput{})
}