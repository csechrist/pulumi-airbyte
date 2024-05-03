// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airbyte

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-airbyte/sdk/go/airbyte/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SourceSpacexAPI Resource
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
//			_, err := airbyte.NewSourceSpacexApi(ctx, "mySourceSpacexapi", &airbyte.SourceSpacexApiArgs{
//				Configuration: &airbyte.SourceSpacexApiConfigurationArgs{
//					Id:      pulumi.String("cbdb35a9-6cd0-4e48-b1e4-b30469b6ca0b"),
//					Options: pulumi.String("...my_options..."),
//				},
//				DefinitionId: pulumi.String("303cf017-cd97-4836-bf1b-e7e9b4aabfc5"),
//				SecretId:     pulumi.String("...my_secret_id..."),
//				WorkspaceId:  pulumi.String("c36bb733-7bf0-4bec-a93a-8ae78e1e537d"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SourceSpacexApi struct {
	pulumi.CustomResourceState

	Configuration SourceSpacexApiConfigurationOutput `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrOutput `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrOutput `pulumi:"secretId"`
	SourceId    pulumi.StringOutput    `pulumi:"sourceId"`
	SourceType  pulumi.StringOutput    `pulumi:"sourceType"`
	WorkspaceId pulumi.StringOutput    `pulumi:"workspaceId"`
}

// NewSourceSpacexApi registers a new resource with the given unique name, arguments, and options.
func NewSourceSpacexApi(ctx *pulumi.Context,
	name string, args *SourceSpacexApiArgs, opts ...pulumi.ResourceOption) (*SourceSpacexApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourceSpacexApi
	err := ctx.RegisterResource("airbyte:index/sourceSpacexApi:SourceSpacexApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceSpacexApi gets an existing SourceSpacexApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceSpacexApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceSpacexApiState, opts ...pulumi.ResourceOption) (*SourceSpacexApi, error) {
	var resource SourceSpacexApi
	err := ctx.ReadResource("airbyte:index/sourceSpacexApi:SourceSpacexApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceSpacexApi resources.
type sourceSpacexApiState struct {
	Configuration *SourceSpacexApiConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	SourceId    *string `pulumi:"sourceId"`
	SourceType  *string `pulumi:"sourceType"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type SourceSpacexApiState struct {
	Configuration SourceSpacexApiConfigurationPtrInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	SourceId    pulumi.StringPtrInput
	SourceType  pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (SourceSpacexApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSpacexApiState)(nil)).Elem()
}

type sourceSpacexApiArgs struct {
	Configuration SourceSpacexApiConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceSpacexApi resource.
type SourceSpacexApiArgs struct {
	Configuration SourceSpacexApiConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceSpacexApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceSpacexApiArgs)(nil)).Elem()
}

type SourceSpacexApiInput interface {
	pulumi.Input

	ToSourceSpacexApiOutput() SourceSpacexApiOutput
	ToSourceSpacexApiOutputWithContext(ctx context.Context) SourceSpacexApiOutput
}

func (*SourceSpacexApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSpacexApi)(nil)).Elem()
}

func (i *SourceSpacexApi) ToSourceSpacexApiOutput() SourceSpacexApiOutput {
	return i.ToSourceSpacexApiOutputWithContext(context.Background())
}

func (i *SourceSpacexApi) ToSourceSpacexApiOutputWithContext(ctx context.Context) SourceSpacexApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSpacexApiOutput)
}

// SourceSpacexApiArrayInput is an input type that accepts SourceSpacexApiArray and SourceSpacexApiArrayOutput values.
// You can construct a concrete instance of `SourceSpacexApiArrayInput` via:
//
//	SourceSpacexApiArray{ SourceSpacexApiArgs{...} }
type SourceSpacexApiArrayInput interface {
	pulumi.Input

	ToSourceSpacexApiArrayOutput() SourceSpacexApiArrayOutput
	ToSourceSpacexApiArrayOutputWithContext(context.Context) SourceSpacexApiArrayOutput
}

type SourceSpacexApiArray []SourceSpacexApiInput

func (SourceSpacexApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSpacexApi)(nil)).Elem()
}

func (i SourceSpacexApiArray) ToSourceSpacexApiArrayOutput() SourceSpacexApiArrayOutput {
	return i.ToSourceSpacexApiArrayOutputWithContext(context.Background())
}

func (i SourceSpacexApiArray) ToSourceSpacexApiArrayOutputWithContext(ctx context.Context) SourceSpacexApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSpacexApiArrayOutput)
}

// SourceSpacexApiMapInput is an input type that accepts SourceSpacexApiMap and SourceSpacexApiMapOutput values.
// You can construct a concrete instance of `SourceSpacexApiMapInput` via:
//
//	SourceSpacexApiMap{ "key": SourceSpacexApiArgs{...} }
type SourceSpacexApiMapInput interface {
	pulumi.Input

	ToSourceSpacexApiMapOutput() SourceSpacexApiMapOutput
	ToSourceSpacexApiMapOutputWithContext(context.Context) SourceSpacexApiMapOutput
}

type SourceSpacexApiMap map[string]SourceSpacexApiInput

func (SourceSpacexApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSpacexApi)(nil)).Elem()
}

func (i SourceSpacexApiMap) ToSourceSpacexApiMapOutput() SourceSpacexApiMapOutput {
	return i.ToSourceSpacexApiMapOutputWithContext(context.Background())
}

func (i SourceSpacexApiMap) ToSourceSpacexApiMapOutputWithContext(ctx context.Context) SourceSpacexApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceSpacexApiMapOutput)
}

type SourceSpacexApiOutput struct{ *pulumi.OutputState }

func (SourceSpacexApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceSpacexApi)(nil)).Elem()
}

func (o SourceSpacexApiOutput) ToSourceSpacexApiOutput() SourceSpacexApiOutput {
	return o
}

func (o SourceSpacexApiOutput) ToSourceSpacexApiOutputWithContext(ctx context.Context) SourceSpacexApiOutput {
	return o
}

func (o SourceSpacexApiOutput) Configuration() SourceSpacexApiConfigurationOutput {
	return o.ApplyT(func(v *SourceSpacexApi) SourceSpacexApiConfigurationOutput { return v.Configuration }).(SourceSpacexApiConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceSpacexApiOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceSpacexApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceSpacexApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceSpacexApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceSpacexApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceSpacexApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceSpacexApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceSpacexApiArrayOutput struct{ *pulumi.OutputState }

func (SourceSpacexApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceSpacexApi)(nil)).Elem()
}

func (o SourceSpacexApiArrayOutput) ToSourceSpacexApiArrayOutput() SourceSpacexApiArrayOutput {
	return o
}

func (o SourceSpacexApiArrayOutput) ToSourceSpacexApiArrayOutputWithContext(ctx context.Context) SourceSpacexApiArrayOutput {
	return o
}

func (o SourceSpacexApiArrayOutput) Index(i pulumi.IntInput) SourceSpacexApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceSpacexApi {
		return vs[0].([]*SourceSpacexApi)[vs[1].(int)]
	}).(SourceSpacexApiOutput)
}

type SourceSpacexApiMapOutput struct{ *pulumi.OutputState }

func (SourceSpacexApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceSpacexApi)(nil)).Elem()
}

func (o SourceSpacexApiMapOutput) ToSourceSpacexApiMapOutput() SourceSpacexApiMapOutput {
	return o
}

func (o SourceSpacexApiMapOutput) ToSourceSpacexApiMapOutputWithContext(ctx context.Context) SourceSpacexApiMapOutput {
	return o
}

func (o SourceSpacexApiMapOutput) MapIndex(k pulumi.StringInput) SourceSpacexApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceSpacexApi {
		return vs[0].(map[string]*SourceSpacexApi)[vs[1].(string)]
	}).(SourceSpacexApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSpacexApiInput)(nil)).Elem(), &SourceSpacexApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSpacexApiArrayInput)(nil)).Elem(), SourceSpacexApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceSpacexApiMapInput)(nil)).Elem(), SourceSpacexApiMap{})
	pulumi.RegisterOutputType(SourceSpacexApiOutput{})
	pulumi.RegisterOutputType(SourceSpacexApiArrayOutput{})
	pulumi.RegisterOutputType(SourceSpacexApiMapOutput{})
}