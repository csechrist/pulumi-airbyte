// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSftp DataSource
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
//			_, err := airbyte.LookupSourceSftp(ctx, &airbyte.LookupSourceSftpArgs{
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
func LookupSourceSftp(ctx *pulumi.Context, args *LookupSourceSftpArgs, opts ...pulumi.InvokeOption) (*LookupSourceSftpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceSftpResult
	err := ctx.Invoke("airbyte:index/getSourceSftp:getSourceSftp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceSftp.
type LookupSourceSftpArgs struct {
	SourceId string `pulumi:"sourceId"`
}

// A collection of values returned by getSourceSftp.
type LookupSourceSftpResult struct {
	// The values required to configure the source. Parsed as JSON.
	Configuration string `pulumi:"configuration"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	SourceId    string `pulumi:"sourceId"`
	SourceType  string `pulumi:"sourceType"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupSourceSftpOutput(ctx *pulumi.Context, args LookupSourceSftpOutputArgs, opts ...pulumi.InvokeOption) LookupSourceSftpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceSftpResult, error) {
			args := v.(LookupSourceSftpArgs)
			r, err := LookupSourceSftp(ctx, &args, opts...)
			var s LookupSourceSftpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceSftpResultOutput)
}

// A collection of arguments for invoking getSourceSftp.
type LookupSourceSftpOutputArgs struct {
	SourceId pulumi.StringInput `pulumi:"sourceId"`
}

func (LookupSourceSftpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpArgs)(nil)).Elem()
}

// A collection of values returned by getSourceSftp.
type LookupSourceSftpResultOutput struct{ *pulumi.OutputState }

func (LookupSourceSftpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceSftpResult)(nil)).Elem()
}

func (o LookupSourceSftpResultOutput) ToLookupSourceSftpResultOutput() LookupSourceSftpResultOutput {
	return o
}

func (o LookupSourceSftpResultOutput) ToLookupSourceSftpResultOutputWithContext(ctx context.Context) LookupSourceSftpResultOutput {
	return o
}

// The values required to configure the source. Parsed as JSON.
func (o LookupSourceSftpResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.Configuration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSourceSftpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o LookupSourceSftpResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceSftpResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceSftpResultOutput{})
}
