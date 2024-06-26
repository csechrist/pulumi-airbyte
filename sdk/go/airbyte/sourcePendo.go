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

// SourcePendo Resource
type SourcePendo struct {
	pulumi.CustomResourceState

	Configuration SourcePendoConfigurationOutput `pulumi:"configuration"`
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

// NewSourcePendo registers a new resource with the given unique name, arguments, and options.
func NewSourcePendo(ctx *pulumi.Context,
	name string, args *SourcePendoArgs, opts ...pulumi.ResourceOption) (*SourcePendo, error) {
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
	var resource SourcePendo
	err := ctx.RegisterResource("airbyte:index/sourcePendo:SourcePendo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePendo gets an existing SourcePendo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePendo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePendoState, opts ...pulumi.ResourceOption) (*SourcePendo, error) {
	var resource SourcePendo
	err := ctx.ReadResource("airbyte:index/sourcePendo:SourcePendo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePendo resources.
type sourcePendoState struct {
	Configuration *SourcePendoConfiguration `pulumi:"configuration"`
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

type SourcePendoState struct {
	Configuration SourcePendoConfigurationPtrInput
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

func (SourcePendoState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePendoState)(nil)).Elem()
}

type sourcePendoArgs struct {
	Configuration SourcePendoConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePendo resource.
type SourcePendoArgs struct {
	Configuration SourcePendoConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePendoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePendoArgs)(nil)).Elem()
}

type SourcePendoInput interface {
	pulumi.Input

	ToSourcePendoOutput() SourcePendoOutput
	ToSourcePendoOutputWithContext(ctx context.Context) SourcePendoOutput
}

func (*SourcePendo) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePendo)(nil)).Elem()
}

func (i *SourcePendo) ToSourcePendoOutput() SourcePendoOutput {
	return i.ToSourcePendoOutputWithContext(context.Background())
}

func (i *SourcePendo) ToSourcePendoOutputWithContext(ctx context.Context) SourcePendoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePendoOutput)
}

// SourcePendoArrayInput is an input type that accepts SourcePendoArray and SourcePendoArrayOutput values.
// You can construct a concrete instance of `SourcePendoArrayInput` via:
//
//	SourcePendoArray{ SourcePendoArgs{...} }
type SourcePendoArrayInput interface {
	pulumi.Input

	ToSourcePendoArrayOutput() SourcePendoArrayOutput
	ToSourcePendoArrayOutputWithContext(context.Context) SourcePendoArrayOutput
}

type SourcePendoArray []SourcePendoInput

func (SourcePendoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePendo)(nil)).Elem()
}

func (i SourcePendoArray) ToSourcePendoArrayOutput() SourcePendoArrayOutput {
	return i.ToSourcePendoArrayOutputWithContext(context.Background())
}

func (i SourcePendoArray) ToSourcePendoArrayOutputWithContext(ctx context.Context) SourcePendoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePendoArrayOutput)
}

// SourcePendoMapInput is an input type that accepts SourcePendoMap and SourcePendoMapOutput values.
// You can construct a concrete instance of `SourcePendoMapInput` via:
//
//	SourcePendoMap{ "key": SourcePendoArgs{...} }
type SourcePendoMapInput interface {
	pulumi.Input

	ToSourcePendoMapOutput() SourcePendoMapOutput
	ToSourcePendoMapOutputWithContext(context.Context) SourcePendoMapOutput
}

type SourcePendoMap map[string]SourcePendoInput

func (SourcePendoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePendo)(nil)).Elem()
}

func (i SourcePendoMap) ToSourcePendoMapOutput() SourcePendoMapOutput {
	return i.ToSourcePendoMapOutputWithContext(context.Background())
}

func (i SourcePendoMap) ToSourcePendoMapOutputWithContext(ctx context.Context) SourcePendoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePendoMapOutput)
}

type SourcePendoOutput struct{ *pulumi.OutputState }

func (SourcePendoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePendo)(nil)).Elem()
}

func (o SourcePendoOutput) ToSourcePendoOutput() SourcePendoOutput {
	return o
}

func (o SourcePendoOutput) ToSourcePendoOutputWithContext(ctx context.Context) SourcePendoOutput {
	return o
}

func (o SourcePendoOutput) Configuration() SourcePendoConfigurationOutput {
	return o.ApplyT(func(v *SourcePendo) SourcePendoConfigurationOutput { return v.Configuration }).(SourcePendoConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourcePendoOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourcePendoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourcePendoOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePendoOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePendoOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePendoOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePendo) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourcePendoArrayOutput struct{ *pulumi.OutputState }

func (SourcePendoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePendo)(nil)).Elem()
}

func (o SourcePendoArrayOutput) ToSourcePendoArrayOutput() SourcePendoArrayOutput {
	return o
}

func (o SourcePendoArrayOutput) ToSourcePendoArrayOutputWithContext(ctx context.Context) SourcePendoArrayOutput {
	return o
}

func (o SourcePendoArrayOutput) Index(i pulumi.IntInput) SourcePendoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourcePendo {
		return vs[0].([]*SourcePendo)[vs[1].(int)]
	}).(SourcePendoOutput)
}

type SourcePendoMapOutput struct{ *pulumi.OutputState }

func (SourcePendoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePendo)(nil)).Elem()
}

func (o SourcePendoMapOutput) ToSourcePendoMapOutput() SourcePendoMapOutput {
	return o
}

func (o SourcePendoMapOutput) ToSourcePendoMapOutputWithContext(ctx context.Context) SourcePendoMapOutput {
	return o
}

func (o SourcePendoMapOutput) MapIndex(k pulumi.StringInput) SourcePendoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourcePendo {
		return vs[0].(map[string]*SourcePendo)[vs[1].(string)]
	}).(SourcePendoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePendoInput)(nil)).Elem(), &SourcePendo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePendoArrayInput)(nil)).Elem(), SourcePendoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePendoMapInput)(nil)).Elem(), SourcePendoMap{})
	pulumi.RegisterOutputType(SourcePendoOutput{})
	pulumi.RegisterOutputType(SourcePendoArrayOutput{})
	pulumi.RegisterOutputType(SourcePendoMapOutput{})
}
