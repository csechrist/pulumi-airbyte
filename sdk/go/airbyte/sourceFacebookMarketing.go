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

// SourceFacebookMarketing Resource
type SourceFacebookMarketing struct {
	pulumi.CustomResourceState

	Configuration SourceFacebookMarketingConfigurationOutput `pulumi:"configuration"`
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

// NewSourceFacebookMarketing registers a new resource with the given unique name, arguments, and options.
func NewSourceFacebookMarketing(ctx *pulumi.Context,
	name string, args *SourceFacebookMarketingArgs, opts ...pulumi.ResourceOption) (*SourceFacebookMarketing, error) {
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
	var resource SourceFacebookMarketing
	err := ctx.RegisterResource("airbyte:index/sourceFacebookMarketing:SourceFacebookMarketing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceFacebookMarketing gets an existing SourceFacebookMarketing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceFacebookMarketing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceFacebookMarketingState, opts ...pulumi.ResourceOption) (*SourceFacebookMarketing, error) {
	var resource SourceFacebookMarketing
	err := ctx.ReadResource("airbyte:index/sourceFacebookMarketing:SourceFacebookMarketing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceFacebookMarketing resources.
type sourceFacebookMarketingState struct {
	Configuration *SourceFacebookMarketingConfiguration `pulumi:"configuration"`
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

type SourceFacebookMarketingState struct {
	Configuration SourceFacebookMarketingConfigurationPtrInput
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

func (SourceFacebookMarketingState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFacebookMarketingState)(nil)).Elem()
}

type sourceFacebookMarketingArgs struct {
	Configuration SourceFacebookMarketingConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceFacebookMarketing resource.
type SourceFacebookMarketingArgs struct {
	Configuration SourceFacebookMarketingConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceFacebookMarketingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceFacebookMarketingArgs)(nil)).Elem()
}

type SourceFacebookMarketingInput interface {
	pulumi.Input

	ToSourceFacebookMarketingOutput() SourceFacebookMarketingOutput
	ToSourceFacebookMarketingOutputWithContext(ctx context.Context) SourceFacebookMarketingOutput
}

func (*SourceFacebookMarketing) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFacebookMarketing)(nil)).Elem()
}

func (i *SourceFacebookMarketing) ToSourceFacebookMarketingOutput() SourceFacebookMarketingOutput {
	return i.ToSourceFacebookMarketingOutputWithContext(context.Background())
}

func (i *SourceFacebookMarketing) ToSourceFacebookMarketingOutputWithContext(ctx context.Context) SourceFacebookMarketingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFacebookMarketingOutput)
}

// SourceFacebookMarketingArrayInput is an input type that accepts SourceFacebookMarketingArray and SourceFacebookMarketingArrayOutput values.
// You can construct a concrete instance of `SourceFacebookMarketingArrayInput` via:
//
//	SourceFacebookMarketingArray{ SourceFacebookMarketingArgs{...} }
type SourceFacebookMarketingArrayInput interface {
	pulumi.Input

	ToSourceFacebookMarketingArrayOutput() SourceFacebookMarketingArrayOutput
	ToSourceFacebookMarketingArrayOutputWithContext(context.Context) SourceFacebookMarketingArrayOutput
}

type SourceFacebookMarketingArray []SourceFacebookMarketingInput

func (SourceFacebookMarketingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFacebookMarketing)(nil)).Elem()
}

func (i SourceFacebookMarketingArray) ToSourceFacebookMarketingArrayOutput() SourceFacebookMarketingArrayOutput {
	return i.ToSourceFacebookMarketingArrayOutputWithContext(context.Background())
}

func (i SourceFacebookMarketingArray) ToSourceFacebookMarketingArrayOutputWithContext(ctx context.Context) SourceFacebookMarketingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFacebookMarketingArrayOutput)
}

// SourceFacebookMarketingMapInput is an input type that accepts SourceFacebookMarketingMap and SourceFacebookMarketingMapOutput values.
// You can construct a concrete instance of `SourceFacebookMarketingMapInput` via:
//
//	SourceFacebookMarketingMap{ "key": SourceFacebookMarketingArgs{...} }
type SourceFacebookMarketingMapInput interface {
	pulumi.Input

	ToSourceFacebookMarketingMapOutput() SourceFacebookMarketingMapOutput
	ToSourceFacebookMarketingMapOutputWithContext(context.Context) SourceFacebookMarketingMapOutput
}

type SourceFacebookMarketingMap map[string]SourceFacebookMarketingInput

func (SourceFacebookMarketingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFacebookMarketing)(nil)).Elem()
}

func (i SourceFacebookMarketingMap) ToSourceFacebookMarketingMapOutput() SourceFacebookMarketingMapOutput {
	return i.ToSourceFacebookMarketingMapOutputWithContext(context.Background())
}

func (i SourceFacebookMarketingMap) ToSourceFacebookMarketingMapOutputWithContext(ctx context.Context) SourceFacebookMarketingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceFacebookMarketingMapOutput)
}

type SourceFacebookMarketingOutput struct{ *pulumi.OutputState }

func (SourceFacebookMarketingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceFacebookMarketing)(nil)).Elem()
}

func (o SourceFacebookMarketingOutput) ToSourceFacebookMarketingOutput() SourceFacebookMarketingOutput {
	return o
}

func (o SourceFacebookMarketingOutput) ToSourceFacebookMarketingOutputWithContext(ctx context.Context) SourceFacebookMarketingOutput {
	return o
}

func (o SourceFacebookMarketingOutput) Configuration() SourceFacebookMarketingConfigurationOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) SourceFacebookMarketingConfigurationOutput { return v.Configuration }).(SourceFacebookMarketingConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceFacebookMarketingOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceFacebookMarketingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceFacebookMarketingOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceFacebookMarketingOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceFacebookMarketingOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceFacebookMarketingOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceFacebookMarketing) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceFacebookMarketingArrayOutput struct{ *pulumi.OutputState }

func (SourceFacebookMarketingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceFacebookMarketing)(nil)).Elem()
}

func (o SourceFacebookMarketingArrayOutput) ToSourceFacebookMarketingArrayOutput() SourceFacebookMarketingArrayOutput {
	return o
}

func (o SourceFacebookMarketingArrayOutput) ToSourceFacebookMarketingArrayOutputWithContext(ctx context.Context) SourceFacebookMarketingArrayOutput {
	return o
}

func (o SourceFacebookMarketingArrayOutput) Index(i pulumi.IntInput) SourceFacebookMarketingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceFacebookMarketing {
		return vs[0].([]*SourceFacebookMarketing)[vs[1].(int)]
	}).(SourceFacebookMarketingOutput)
}

type SourceFacebookMarketingMapOutput struct{ *pulumi.OutputState }

func (SourceFacebookMarketingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceFacebookMarketing)(nil)).Elem()
}

func (o SourceFacebookMarketingMapOutput) ToSourceFacebookMarketingMapOutput() SourceFacebookMarketingMapOutput {
	return o
}

func (o SourceFacebookMarketingMapOutput) ToSourceFacebookMarketingMapOutputWithContext(ctx context.Context) SourceFacebookMarketingMapOutput {
	return o
}

func (o SourceFacebookMarketingMapOutput) MapIndex(k pulumi.StringInput) SourceFacebookMarketingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceFacebookMarketing {
		return vs[0].(map[string]*SourceFacebookMarketing)[vs[1].(string)]
	}).(SourceFacebookMarketingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFacebookMarketingInput)(nil)).Elem(), &SourceFacebookMarketing{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFacebookMarketingArrayInput)(nil)).Elem(), SourceFacebookMarketingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceFacebookMarketingMapInput)(nil)).Elem(), SourceFacebookMarketingMap{})
	pulumi.RegisterOutputType(SourceFacebookMarketingOutput{})
	pulumi.RegisterOutputType(SourceFacebookMarketingArrayOutput{})
	pulumi.RegisterOutputType(SourceFacebookMarketingMapOutput{})
}