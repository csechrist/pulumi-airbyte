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

// SourceAha Resource
type SourceAha struct {
	pulumi.CustomResourceState

	Configuration SourceAhaConfigurationOutput `pulumi:"configuration"`
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

// NewSourceAha registers a new resource with the given unique name, arguments, and options.
func NewSourceAha(ctx *pulumi.Context,
	name string, args *SourceAhaArgs, opts ...pulumi.ResourceOption) (*SourceAha, error) {
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
	var resource SourceAha
	err := ctx.RegisterResource("airbyte:index/sourceAha:SourceAha", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAha gets an existing SourceAha resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAha(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAhaState, opts ...pulumi.ResourceOption) (*SourceAha, error) {
	var resource SourceAha
	err := ctx.ReadResource("airbyte:index/sourceAha:SourceAha", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAha resources.
type sourceAhaState struct {
	Configuration *SourceAhaConfiguration `pulumi:"configuration"`
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

type SourceAhaState struct {
	Configuration SourceAhaConfigurationPtrInput
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

func (SourceAhaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAhaState)(nil)).Elem()
}

type sourceAhaArgs struct {
	Configuration SourceAhaConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAha resource.
type SourceAhaArgs struct {
	Configuration SourceAhaConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAhaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAhaArgs)(nil)).Elem()
}

type SourceAhaInput interface {
	pulumi.Input

	ToSourceAhaOutput() SourceAhaOutput
	ToSourceAhaOutputWithContext(ctx context.Context) SourceAhaOutput
}

func (*SourceAha) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAha)(nil)).Elem()
}

func (i *SourceAha) ToSourceAhaOutput() SourceAhaOutput {
	return i.ToSourceAhaOutputWithContext(context.Background())
}

func (i *SourceAha) ToSourceAhaOutputWithContext(ctx context.Context) SourceAhaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAhaOutput)
}

// SourceAhaArrayInput is an input type that accepts SourceAhaArray and SourceAhaArrayOutput values.
// You can construct a concrete instance of `SourceAhaArrayInput` via:
//
//	SourceAhaArray{ SourceAhaArgs{...} }
type SourceAhaArrayInput interface {
	pulumi.Input

	ToSourceAhaArrayOutput() SourceAhaArrayOutput
	ToSourceAhaArrayOutputWithContext(context.Context) SourceAhaArrayOutput
}

type SourceAhaArray []SourceAhaInput

func (SourceAhaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAha)(nil)).Elem()
}

func (i SourceAhaArray) ToSourceAhaArrayOutput() SourceAhaArrayOutput {
	return i.ToSourceAhaArrayOutputWithContext(context.Background())
}

func (i SourceAhaArray) ToSourceAhaArrayOutputWithContext(ctx context.Context) SourceAhaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAhaArrayOutput)
}

// SourceAhaMapInput is an input type that accepts SourceAhaMap and SourceAhaMapOutput values.
// You can construct a concrete instance of `SourceAhaMapInput` via:
//
//	SourceAhaMap{ "key": SourceAhaArgs{...} }
type SourceAhaMapInput interface {
	pulumi.Input

	ToSourceAhaMapOutput() SourceAhaMapOutput
	ToSourceAhaMapOutputWithContext(context.Context) SourceAhaMapOutput
}

type SourceAhaMap map[string]SourceAhaInput

func (SourceAhaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAha)(nil)).Elem()
}

func (i SourceAhaMap) ToSourceAhaMapOutput() SourceAhaMapOutput {
	return i.ToSourceAhaMapOutputWithContext(context.Background())
}

func (i SourceAhaMap) ToSourceAhaMapOutputWithContext(ctx context.Context) SourceAhaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAhaMapOutput)
}

type SourceAhaOutput struct{ *pulumi.OutputState }

func (SourceAhaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAha)(nil)).Elem()
}

func (o SourceAhaOutput) ToSourceAhaOutput() SourceAhaOutput {
	return o
}

func (o SourceAhaOutput) ToSourceAhaOutputWithContext(ctx context.Context) SourceAhaOutput {
	return o
}

func (o SourceAhaOutput) Configuration() SourceAhaConfigurationOutput {
	return o.ApplyT(func(v *SourceAha) SourceAhaConfigurationOutput { return v.Configuration }).(SourceAhaConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceAhaOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceAhaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceAhaOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAhaOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAhaOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAhaOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAha) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAhaArrayOutput struct{ *pulumi.OutputState }

func (SourceAhaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAha)(nil)).Elem()
}

func (o SourceAhaArrayOutput) ToSourceAhaArrayOutput() SourceAhaArrayOutput {
	return o
}

func (o SourceAhaArrayOutput) ToSourceAhaArrayOutputWithContext(ctx context.Context) SourceAhaArrayOutput {
	return o
}

func (o SourceAhaArrayOutput) Index(i pulumi.IntInput) SourceAhaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAha {
		return vs[0].([]*SourceAha)[vs[1].(int)]
	}).(SourceAhaOutput)
}

type SourceAhaMapOutput struct{ *pulumi.OutputState }

func (SourceAhaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAha)(nil)).Elem()
}

func (o SourceAhaMapOutput) ToSourceAhaMapOutput() SourceAhaMapOutput {
	return o
}

func (o SourceAhaMapOutput) ToSourceAhaMapOutputWithContext(ctx context.Context) SourceAhaMapOutput {
	return o
}

func (o SourceAhaMapOutput) MapIndex(k pulumi.StringInput) SourceAhaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAha {
		return vs[0].(map[string]*SourceAha)[vs[1].(string)]
	}).(SourceAhaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAhaInput)(nil)).Elem(), &SourceAha{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAhaArrayInput)(nil)).Elem(), SourceAhaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAhaMapInput)(nil)).Elem(), SourceAhaMap{})
	pulumi.RegisterOutputType(SourceAhaOutput{})
	pulumi.RegisterOutputType(SourceAhaArrayOutput{})
	pulumi.RegisterOutputType(SourceAhaMapOutput{})
}