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

// SourceEmailoctopus Resource
type SourceEmailoctopus struct {
	pulumi.CustomResourceState

	Configuration SourceEmailoctopusConfigurationOutput `pulumi:"configuration"`
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

// NewSourceEmailoctopus registers a new resource with the given unique name, arguments, and options.
func NewSourceEmailoctopus(ctx *pulumi.Context,
	name string, args *SourceEmailoctopusArgs, opts ...pulumi.ResourceOption) (*SourceEmailoctopus, error) {
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
	var resource SourceEmailoctopus
	err := ctx.RegisterResource("airbyte:index/sourceEmailoctopus:SourceEmailoctopus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceEmailoctopus gets an existing SourceEmailoctopus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceEmailoctopus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceEmailoctopusState, opts ...pulumi.ResourceOption) (*SourceEmailoctopus, error) {
	var resource SourceEmailoctopus
	err := ctx.ReadResource("airbyte:index/sourceEmailoctopus:SourceEmailoctopus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceEmailoctopus resources.
type sourceEmailoctopusState struct {
	Configuration *SourceEmailoctopusConfiguration `pulumi:"configuration"`
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

type SourceEmailoctopusState struct {
	Configuration SourceEmailoctopusConfigurationPtrInput
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

func (SourceEmailoctopusState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceEmailoctopusState)(nil)).Elem()
}

type sourceEmailoctopusArgs struct {
	Configuration SourceEmailoctopusConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceEmailoctopus resource.
type SourceEmailoctopusArgs struct {
	Configuration SourceEmailoctopusConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceEmailoctopusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceEmailoctopusArgs)(nil)).Elem()
}

type SourceEmailoctopusInput interface {
	pulumi.Input

	ToSourceEmailoctopusOutput() SourceEmailoctopusOutput
	ToSourceEmailoctopusOutputWithContext(ctx context.Context) SourceEmailoctopusOutput
}

func (*SourceEmailoctopus) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceEmailoctopus)(nil)).Elem()
}

func (i *SourceEmailoctopus) ToSourceEmailoctopusOutput() SourceEmailoctopusOutput {
	return i.ToSourceEmailoctopusOutputWithContext(context.Background())
}

func (i *SourceEmailoctopus) ToSourceEmailoctopusOutputWithContext(ctx context.Context) SourceEmailoctopusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceEmailoctopusOutput)
}

// SourceEmailoctopusArrayInput is an input type that accepts SourceEmailoctopusArray and SourceEmailoctopusArrayOutput values.
// You can construct a concrete instance of `SourceEmailoctopusArrayInput` via:
//
//	SourceEmailoctopusArray{ SourceEmailoctopusArgs{...} }
type SourceEmailoctopusArrayInput interface {
	pulumi.Input

	ToSourceEmailoctopusArrayOutput() SourceEmailoctopusArrayOutput
	ToSourceEmailoctopusArrayOutputWithContext(context.Context) SourceEmailoctopusArrayOutput
}

type SourceEmailoctopusArray []SourceEmailoctopusInput

func (SourceEmailoctopusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceEmailoctopus)(nil)).Elem()
}

func (i SourceEmailoctopusArray) ToSourceEmailoctopusArrayOutput() SourceEmailoctopusArrayOutput {
	return i.ToSourceEmailoctopusArrayOutputWithContext(context.Background())
}

func (i SourceEmailoctopusArray) ToSourceEmailoctopusArrayOutputWithContext(ctx context.Context) SourceEmailoctopusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceEmailoctopusArrayOutput)
}

// SourceEmailoctopusMapInput is an input type that accepts SourceEmailoctopusMap and SourceEmailoctopusMapOutput values.
// You can construct a concrete instance of `SourceEmailoctopusMapInput` via:
//
//	SourceEmailoctopusMap{ "key": SourceEmailoctopusArgs{...} }
type SourceEmailoctopusMapInput interface {
	pulumi.Input

	ToSourceEmailoctopusMapOutput() SourceEmailoctopusMapOutput
	ToSourceEmailoctopusMapOutputWithContext(context.Context) SourceEmailoctopusMapOutput
}

type SourceEmailoctopusMap map[string]SourceEmailoctopusInput

func (SourceEmailoctopusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceEmailoctopus)(nil)).Elem()
}

func (i SourceEmailoctopusMap) ToSourceEmailoctopusMapOutput() SourceEmailoctopusMapOutput {
	return i.ToSourceEmailoctopusMapOutputWithContext(context.Background())
}

func (i SourceEmailoctopusMap) ToSourceEmailoctopusMapOutputWithContext(ctx context.Context) SourceEmailoctopusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceEmailoctopusMapOutput)
}

type SourceEmailoctopusOutput struct{ *pulumi.OutputState }

func (SourceEmailoctopusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceEmailoctopus)(nil)).Elem()
}

func (o SourceEmailoctopusOutput) ToSourceEmailoctopusOutput() SourceEmailoctopusOutput {
	return o
}

func (o SourceEmailoctopusOutput) ToSourceEmailoctopusOutputWithContext(ctx context.Context) SourceEmailoctopusOutput {
	return o
}

func (o SourceEmailoctopusOutput) Configuration() SourceEmailoctopusConfigurationOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) SourceEmailoctopusConfigurationOutput { return v.Configuration }).(SourceEmailoctopusConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceEmailoctopusOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceEmailoctopusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceEmailoctopusOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceEmailoctopusOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceEmailoctopusOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceEmailoctopusOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceEmailoctopus) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceEmailoctopusArrayOutput struct{ *pulumi.OutputState }

func (SourceEmailoctopusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceEmailoctopus)(nil)).Elem()
}

func (o SourceEmailoctopusArrayOutput) ToSourceEmailoctopusArrayOutput() SourceEmailoctopusArrayOutput {
	return o
}

func (o SourceEmailoctopusArrayOutput) ToSourceEmailoctopusArrayOutputWithContext(ctx context.Context) SourceEmailoctopusArrayOutput {
	return o
}

func (o SourceEmailoctopusArrayOutput) Index(i pulumi.IntInput) SourceEmailoctopusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceEmailoctopus {
		return vs[0].([]*SourceEmailoctopus)[vs[1].(int)]
	}).(SourceEmailoctopusOutput)
}

type SourceEmailoctopusMapOutput struct{ *pulumi.OutputState }

func (SourceEmailoctopusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceEmailoctopus)(nil)).Elem()
}

func (o SourceEmailoctopusMapOutput) ToSourceEmailoctopusMapOutput() SourceEmailoctopusMapOutput {
	return o
}

func (o SourceEmailoctopusMapOutput) ToSourceEmailoctopusMapOutputWithContext(ctx context.Context) SourceEmailoctopusMapOutput {
	return o
}

func (o SourceEmailoctopusMapOutput) MapIndex(k pulumi.StringInput) SourceEmailoctopusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceEmailoctopus {
		return vs[0].(map[string]*SourceEmailoctopus)[vs[1].(string)]
	}).(SourceEmailoctopusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceEmailoctopusInput)(nil)).Elem(), &SourceEmailoctopus{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceEmailoctopusArrayInput)(nil)).Elem(), SourceEmailoctopusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceEmailoctopusMapInput)(nil)).Elem(), SourceEmailoctopusMap{})
	pulumi.RegisterOutputType(SourceEmailoctopusOutput{})
	pulumi.RegisterOutputType(SourceEmailoctopusArrayOutput{})
	pulumi.RegisterOutputType(SourceEmailoctopusMapOutput{})
}