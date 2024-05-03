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

// SourceTiktokMarketing Resource
type SourceTiktokMarketing struct {
	pulumi.CustomResourceState

	Configuration SourceTiktokMarketingConfigurationOutput `pulumi:"configuration"`
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

// NewSourceTiktokMarketing registers a new resource with the given unique name, arguments, and options.
func NewSourceTiktokMarketing(ctx *pulumi.Context,
	name string, args *SourceTiktokMarketingArgs, opts ...pulumi.ResourceOption) (*SourceTiktokMarketing, error) {
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
	var resource SourceTiktokMarketing
	err := ctx.RegisterResource("airbyte:index/sourceTiktokMarketing:SourceTiktokMarketing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceTiktokMarketing gets an existing SourceTiktokMarketing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceTiktokMarketing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceTiktokMarketingState, opts ...pulumi.ResourceOption) (*SourceTiktokMarketing, error) {
	var resource SourceTiktokMarketing
	err := ctx.ReadResource("airbyte:index/sourceTiktokMarketing:SourceTiktokMarketing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceTiktokMarketing resources.
type sourceTiktokMarketingState struct {
	Configuration *SourceTiktokMarketingConfiguration `pulumi:"configuration"`
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

type SourceTiktokMarketingState struct {
	Configuration SourceTiktokMarketingConfigurationPtrInput
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

func (SourceTiktokMarketingState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTiktokMarketingState)(nil)).Elem()
}

type sourceTiktokMarketingArgs struct {
	Configuration SourceTiktokMarketingConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceTiktokMarketing resource.
type SourceTiktokMarketingArgs struct {
	Configuration SourceTiktokMarketingConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceTiktokMarketingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTiktokMarketingArgs)(nil)).Elem()
}

type SourceTiktokMarketingInput interface {
	pulumi.Input

	ToSourceTiktokMarketingOutput() SourceTiktokMarketingOutput
	ToSourceTiktokMarketingOutputWithContext(ctx context.Context) SourceTiktokMarketingOutput
}

func (*SourceTiktokMarketing) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTiktokMarketing)(nil)).Elem()
}

func (i *SourceTiktokMarketing) ToSourceTiktokMarketingOutput() SourceTiktokMarketingOutput {
	return i.ToSourceTiktokMarketingOutputWithContext(context.Background())
}

func (i *SourceTiktokMarketing) ToSourceTiktokMarketingOutputWithContext(ctx context.Context) SourceTiktokMarketingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTiktokMarketingOutput)
}

// SourceTiktokMarketingArrayInput is an input type that accepts SourceTiktokMarketingArray and SourceTiktokMarketingArrayOutput values.
// You can construct a concrete instance of `SourceTiktokMarketingArrayInput` via:
//
//	SourceTiktokMarketingArray{ SourceTiktokMarketingArgs{...} }
type SourceTiktokMarketingArrayInput interface {
	pulumi.Input

	ToSourceTiktokMarketingArrayOutput() SourceTiktokMarketingArrayOutput
	ToSourceTiktokMarketingArrayOutputWithContext(context.Context) SourceTiktokMarketingArrayOutput
}

type SourceTiktokMarketingArray []SourceTiktokMarketingInput

func (SourceTiktokMarketingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTiktokMarketing)(nil)).Elem()
}

func (i SourceTiktokMarketingArray) ToSourceTiktokMarketingArrayOutput() SourceTiktokMarketingArrayOutput {
	return i.ToSourceTiktokMarketingArrayOutputWithContext(context.Background())
}

func (i SourceTiktokMarketingArray) ToSourceTiktokMarketingArrayOutputWithContext(ctx context.Context) SourceTiktokMarketingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTiktokMarketingArrayOutput)
}

// SourceTiktokMarketingMapInput is an input type that accepts SourceTiktokMarketingMap and SourceTiktokMarketingMapOutput values.
// You can construct a concrete instance of `SourceTiktokMarketingMapInput` via:
//
//	SourceTiktokMarketingMap{ "key": SourceTiktokMarketingArgs{...} }
type SourceTiktokMarketingMapInput interface {
	pulumi.Input

	ToSourceTiktokMarketingMapOutput() SourceTiktokMarketingMapOutput
	ToSourceTiktokMarketingMapOutputWithContext(context.Context) SourceTiktokMarketingMapOutput
}

type SourceTiktokMarketingMap map[string]SourceTiktokMarketingInput

func (SourceTiktokMarketingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTiktokMarketing)(nil)).Elem()
}

func (i SourceTiktokMarketingMap) ToSourceTiktokMarketingMapOutput() SourceTiktokMarketingMapOutput {
	return i.ToSourceTiktokMarketingMapOutputWithContext(context.Background())
}

func (i SourceTiktokMarketingMap) ToSourceTiktokMarketingMapOutputWithContext(ctx context.Context) SourceTiktokMarketingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTiktokMarketingMapOutput)
}

type SourceTiktokMarketingOutput struct{ *pulumi.OutputState }

func (SourceTiktokMarketingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTiktokMarketing)(nil)).Elem()
}

func (o SourceTiktokMarketingOutput) ToSourceTiktokMarketingOutput() SourceTiktokMarketingOutput {
	return o
}

func (o SourceTiktokMarketingOutput) ToSourceTiktokMarketingOutputWithContext(ctx context.Context) SourceTiktokMarketingOutput {
	return o
}

func (o SourceTiktokMarketingOutput) Configuration() SourceTiktokMarketingConfigurationOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) SourceTiktokMarketingConfigurationOutput { return v.Configuration }).(SourceTiktokMarketingConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceTiktokMarketingOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceTiktokMarketingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceTiktokMarketingOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceTiktokMarketingOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceTiktokMarketingOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceTiktokMarketingOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTiktokMarketing) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceTiktokMarketingArrayOutput struct{ *pulumi.OutputState }

func (SourceTiktokMarketingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTiktokMarketing)(nil)).Elem()
}

func (o SourceTiktokMarketingArrayOutput) ToSourceTiktokMarketingArrayOutput() SourceTiktokMarketingArrayOutput {
	return o
}

func (o SourceTiktokMarketingArrayOutput) ToSourceTiktokMarketingArrayOutputWithContext(ctx context.Context) SourceTiktokMarketingArrayOutput {
	return o
}

func (o SourceTiktokMarketingArrayOutput) Index(i pulumi.IntInput) SourceTiktokMarketingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceTiktokMarketing {
		return vs[0].([]*SourceTiktokMarketing)[vs[1].(int)]
	}).(SourceTiktokMarketingOutput)
}

type SourceTiktokMarketingMapOutput struct{ *pulumi.OutputState }

func (SourceTiktokMarketingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTiktokMarketing)(nil)).Elem()
}

func (o SourceTiktokMarketingMapOutput) ToSourceTiktokMarketingMapOutput() SourceTiktokMarketingMapOutput {
	return o
}

func (o SourceTiktokMarketingMapOutput) ToSourceTiktokMarketingMapOutputWithContext(ctx context.Context) SourceTiktokMarketingMapOutput {
	return o
}

func (o SourceTiktokMarketingMapOutput) MapIndex(k pulumi.StringInput) SourceTiktokMarketingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceTiktokMarketing {
		return vs[0].(map[string]*SourceTiktokMarketing)[vs[1].(string)]
	}).(SourceTiktokMarketingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTiktokMarketingInput)(nil)).Elem(), &SourceTiktokMarketing{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTiktokMarketingArrayInput)(nil)).Elem(), SourceTiktokMarketingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTiktokMarketingMapInput)(nil)).Elem(), SourceTiktokMarketingMap{})
	pulumi.RegisterOutputType(SourceTiktokMarketingOutput{})
	pulumi.RegisterOutputType(SourceTiktokMarketingArrayOutput{})
	pulumi.RegisterOutputType(SourceTiktokMarketingMapOutput{})
}