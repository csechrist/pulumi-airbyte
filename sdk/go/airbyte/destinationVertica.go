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

// DestinationVertica Resource
type DestinationVertica struct {
	pulumi.CustomResourceState

	Configuration DestinationVerticaConfigurationOutput `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
	DefinitionId    pulumi.StringPtrOutput `pulumi:"definitionId"`
	DestinationId   pulumi.StringOutput    `pulumi:"destinationId"`
	DestinationType pulumi.StringOutput    `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringOutput `pulumi:"name"`
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewDestinationVertica registers a new resource with the given unique name, arguments, and options.
func NewDestinationVertica(ctx *pulumi.Context,
	name string, args *DestinationVerticaArgs, opts ...pulumi.ResourceOption) (*DestinationVertica, error) {
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
	var resource DestinationVertica
	err := ctx.RegisterResource("airbyte:index/destinationVertica:DestinationVertica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestinationVertica gets an existing DestinationVertica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestinationVertica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestinationVerticaState, opts ...pulumi.ResourceOption) (*DestinationVertica, error) {
	var resource DestinationVertica
	err := ctx.ReadResource("airbyte:index/destinationVertica:DestinationVertica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestinationVertica resources.
type destinationVerticaState struct {
	Configuration *DestinationVerticaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
	DefinitionId    *string `pulumi:"definitionId"`
	DestinationId   *string `pulumi:"destinationId"`
	DestinationType *string `pulumi:"destinationType"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId *string `pulumi:"workspaceId"`
}

type DestinationVerticaState struct {
	Configuration DestinationVerticaConfigurationPtrInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
	DefinitionId    pulumi.StringPtrInput
	DestinationId   pulumi.StringPtrInput
	DestinationType pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringPtrInput
}

func (DestinationVerticaState) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationVerticaState)(nil)).Elem()
}

type destinationVerticaArgs struct {
	Configuration DestinationVerticaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        *string `pulumi:"name"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a DestinationVertica resource.
type DestinationVerticaArgs struct {
	Configuration DestinationVerticaConfigurationInput
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the destination e.g. dev-mysql-instance.
	Name        pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (DestinationVerticaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destinationVerticaArgs)(nil)).Elem()
}

type DestinationVerticaInput interface {
	pulumi.Input

	ToDestinationVerticaOutput() DestinationVerticaOutput
	ToDestinationVerticaOutputWithContext(ctx context.Context) DestinationVerticaOutput
}

func (*DestinationVertica) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationVertica)(nil)).Elem()
}

func (i *DestinationVertica) ToDestinationVerticaOutput() DestinationVerticaOutput {
	return i.ToDestinationVerticaOutputWithContext(context.Background())
}

func (i *DestinationVertica) ToDestinationVerticaOutputWithContext(ctx context.Context) DestinationVerticaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationVerticaOutput)
}

// DestinationVerticaArrayInput is an input type that accepts DestinationVerticaArray and DestinationVerticaArrayOutput values.
// You can construct a concrete instance of `DestinationVerticaArrayInput` via:
//
//	DestinationVerticaArray{ DestinationVerticaArgs{...} }
type DestinationVerticaArrayInput interface {
	pulumi.Input

	ToDestinationVerticaArrayOutput() DestinationVerticaArrayOutput
	ToDestinationVerticaArrayOutputWithContext(context.Context) DestinationVerticaArrayOutput
}

type DestinationVerticaArray []DestinationVerticaInput

func (DestinationVerticaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationVertica)(nil)).Elem()
}

func (i DestinationVerticaArray) ToDestinationVerticaArrayOutput() DestinationVerticaArrayOutput {
	return i.ToDestinationVerticaArrayOutputWithContext(context.Background())
}

func (i DestinationVerticaArray) ToDestinationVerticaArrayOutputWithContext(ctx context.Context) DestinationVerticaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationVerticaArrayOutput)
}

// DestinationVerticaMapInput is an input type that accepts DestinationVerticaMap and DestinationVerticaMapOutput values.
// You can construct a concrete instance of `DestinationVerticaMapInput` via:
//
//	DestinationVerticaMap{ "key": DestinationVerticaArgs{...} }
type DestinationVerticaMapInput interface {
	pulumi.Input

	ToDestinationVerticaMapOutput() DestinationVerticaMapOutput
	ToDestinationVerticaMapOutputWithContext(context.Context) DestinationVerticaMapOutput
}

type DestinationVerticaMap map[string]DestinationVerticaInput

func (DestinationVerticaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationVertica)(nil)).Elem()
}

func (i DestinationVerticaMap) ToDestinationVerticaMapOutput() DestinationVerticaMapOutput {
	return i.ToDestinationVerticaMapOutputWithContext(context.Background())
}

func (i DestinationVerticaMap) ToDestinationVerticaMapOutputWithContext(ctx context.Context) DestinationVerticaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationVerticaMapOutput)
}

type DestinationVerticaOutput struct{ *pulumi.OutputState }

func (DestinationVerticaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationVertica)(nil)).Elem()
}

func (o DestinationVerticaOutput) ToDestinationVerticaOutput() DestinationVerticaOutput {
	return o
}

func (o DestinationVerticaOutput) ToDestinationVerticaOutputWithContext(ctx context.Context) DestinationVerticaOutput {
	return o
}

func (o DestinationVerticaOutput) Configuration() DestinationVerticaConfigurationOutput {
	return o.ApplyT(func(v *DestinationVertica) DestinationVerticaConfigurationOutput { return v.Configuration }).(DestinationVerticaConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
func (o DestinationVerticaOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationVertica) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

func (o DestinationVerticaOutput) DestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationVertica) pulumi.StringOutput { return v.DestinationId }).(pulumi.StringOutput)
}

func (o DestinationVerticaOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationVertica) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// Name of the destination e.g. dev-mysql-instance.
func (o DestinationVerticaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationVertica) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DestinationVerticaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DestinationVertica) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type DestinationVerticaArrayOutput struct{ *pulumi.OutputState }

func (DestinationVerticaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DestinationVertica)(nil)).Elem()
}

func (o DestinationVerticaArrayOutput) ToDestinationVerticaArrayOutput() DestinationVerticaArrayOutput {
	return o
}

func (o DestinationVerticaArrayOutput) ToDestinationVerticaArrayOutputWithContext(ctx context.Context) DestinationVerticaArrayOutput {
	return o
}

func (o DestinationVerticaArrayOutput) Index(i pulumi.IntInput) DestinationVerticaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DestinationVertica {
		return vs[0].([]*DestinationVertica)[vs[1].(int)]
	}).(DestinationVerticaOutput)
}

type DestinationVerticaMapOutput struct{ *pulumi.OutputState }

func (DestinationVerticaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DestinationVertica)(nil)).Elem()
}

func (o DestinationVerticaMapOutput) ToDestinationVerticaMapOutput() DestinationVerticaMapOutput {
	return o
}

func (o DestinationVerticaMapOutput) ToDestinationVerticaMapOutputWithContext(ctx context.Context) DestinationVerticaMapOutput {
	return o
}

func (o DestinationVerticaMapOutput) MapIndex(k pulumi.StringInput) DestinationVerticaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DestinationVertica {
		return vs[0].(map[string]*DestinationVertica)[vs[1].(string)]
	}).(DestinationVerticaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationVerticaInput)(nil)).Elem(), &DestinationVertica{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationVerticaArrayInput)(nil)).Elem(), DestinationVerticaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DestinationVerticaMapInput)(nil)).Elem(), DestinationVerticaMap{})
	pulumi.RegisterOutputType(DestinationVerticaOutput{})
	pulumi.RegisterOutputType(DestinationVerticaArrayOutput{})
	pulumi.RegisterOutputType(DestinationVerticaMapOutput{})
}