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

// SourceTwilio Resource
type SourceTwilio struct {
	pulumi.CustomResourceState

	Configuration SourceTwilioConfigurationOutput `pulumi:"configuration"`
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

// NewSourceTwilio registers a new resource with the given unique name, arguments, and options.
func NewSourceTwilio(ctx *pulumi.Context,
	name string, args *SourceTwilioArgs, opts ...pulumi.ResourceOption) (*SourceTwilio, error) {
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
	var resource SourceTwilio
	err := ctx.RegisterResource("airbyte:index/sourceTwilio:SourceTwilio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceTwilio gets an existing SourceTwilio resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceTwilio(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceTwilioState, opts ...pulumi.ResourceOption) (*SourceTwilio, error) {
	var resource SourceTwilio
	err := ctx.ReadResource("airbyte:index/sourceTwilio:SourceTwilio", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceTwilio resources.
type sourceTwilioState struct {
	Configuration *SourceTwilioConfiguration `pulumi:"configuration"`
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

type SourceTwilioState struct {
	Configuration SourceTwilioConfigurationPtrInput
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

func (SourceTwilioState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTwilioState)(nil)).Elem()
}

type sourceTwilioArgs struct {
	Configuration SourceTwilioConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceTwilio resource.
type SourceTwilioArgs struct {
	Configuration SourceTwilioConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceTwilioArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTwilioArgs)(nil)).Elem()
}

type SourceTwilioInput interface {
	pulumi.Input

	ToSourceTwilioOutput() SourceTwilioOutput
	ToSourceTwilioOutputWithContext(ctx context.Context) SourceTwilioOutput
}

func (*SourceTwilio) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTwilio)(nil)).Elem()
}

func (i *SourceTwilio) ToSourceTwilioOutput() SourceTwilioOutput {
	return i.ToSourceTwilioOutputWithContext(context.Background())
}

func (i *SourceTwilio) ToSourceTwilioOutputWithContext(ctx context.Context) SourceTwilioOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwilioOutput)
}

// SourceTwilioArrayInput is an input type that accepts SourceTwilioArray and SourceTwilioArrayOutput values.
// You can construct a concrete instance of `SourceTwilioArrayInput` via:
//
//	SourceTwilioArray{ SourceTwilioArgs{...} }
type SourceTwilioArrayInput interface {
	pulumi.Input

	ToSourceTwilioArrayOutput() SourceTwilioArrayOutput
	ToSourceTwilioArrayOutputWithContext(context.Context) SourceTwilioArrayOutput
}

type SourceTwilioArray []SourceTwilioInput

func (SourceTwilioArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTwilio)(nil)).Elem()
}

func (i SourceTwilioArray) ToSourceTwilioArrayOutput() SourceTwilioArrayOutput {
	return i.ToSourceTwilioArrayOutputWithContext(context.Background())
}

func (i SourceTwilioArray) ToSourceTwilioArrayOutputWithContext(ctx context.Context) SourceTwilioArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwilioArrayOutput)
}

// SourceTwilioMapInput is an input type that accepts SourceTwilioMap and SourceTwilioMapOutput values.
// You can construct a concrete instance of `SourceTwilioMapInput` via:
//
//	SourceTwilioMap{ "key": SourceTwilioArgs{...} }
type SourceTwilioMapInput interface {
	pulumi.Input

	ToSourceTwilioMapOutput() SourceTwilioMapOutput
	ToSourceTwilioMapOutputWithContext(context.Context) SourceTwilioMapOutput
}

type SourceTwilioMap map[string]SourceTwilioInput

func (SourceTwilioMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTwilio)(nil)).Elem()
}

func (i SourceTwilioMap) ToSourceTwilioMapOutput() SourceTwilioMapOutput {
	return i.ToSourceTwilioMapOutputWithContext(context.Background())
}

func (i SourceTwilioMap) ToSourceTwilioMapOutputWithContext(ctx context.Context) SourceTwilioMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTwilioMapOutput)
}

type SourceTwilioOutput struct{ *pulumi.OutputState }

func (SourceTwilioOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTwilio)(nil)).Elem()
}

func (o SourceTwilioOutput) ToSourceTwilioOutput() SourceTwilioOutput {
	return o
}

func (o SourceTwilioOutput) ToSourceTwilioOutputWithContext(ctx context.Context) SourceTwilioOutput {
	return o
}

func (o SourceTwilioOutput) Configuration() SourceTwilioConfigurationOutput {
	return o.ApplyT(func(v *SourceTwilio) SourceTwilioConfigurationOutput { return v.Configuration }).(SourceTwilioConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceTwilioOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceTwilioOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceTwilioOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceTwilioOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceTwilioOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceTwilioOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTwilio) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceTwilioArrayOutput struct{ *pulumi.OutputState }

func (SourceTwilioArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTwilio)(nil)).Elem()
}

func (o SourceTwilioArrayOutput) ToSourceTwilioArrayOutput() SourceTwilioArrayOutput {
	return o
}

func (o SourceTwilioArrayOutput) ToSourceTwilioArrayOutputWithContext(ctx context.Context) SourceTwilioArrayOutput {
	return o
}

func (o SourceTwilioArrayOutput) Index(i pulumi.IntInput) SourceTwilioOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceTwilio {
		return vs[0].([]*SourceTwilio)[vs[1].(int)]
	}).(SourceTwilioOutput)
}

type SourceTwilioMapOutput struct{ *pulumi.OutputState }

func (SourceTwilioMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTwilio)(nil)).Elem()
}

func (o SourceTwilioMapOutput) ToSourceTwilioMapOutput() SourceTwilioMapOutput {
	return o
}

func (o SourceTwilioMapOutput) ToSourceTwilioMapOutputWithContext(ctx context.Context) SourceTwilioMapOutput {
	return o
}

func (o SourceTwilioMapOutput) MapIndex(k pulumi.StringInput) SourceTwilioOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceTwilio {
		return vs[0].(map[string]*SourceTwilio)[vs[1].(string)]
	}).(SourceTwilioOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwilioInput)(nil)).Elem(), &SourceTwilio{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwilioArrayInput)(nil)).Elem(), SourceTwilioArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTwilioMapInput)(nil)).Elem(), SourceTwilioMap{})
	pulumi.RegisterOutputType(SourceTwilioOutput{})
	pulumi.RegisterOutputType(SourceTwilioArrayOutput{})
	pulumi.RegisterOutputType(SourceTwilioMapOutput{})
}