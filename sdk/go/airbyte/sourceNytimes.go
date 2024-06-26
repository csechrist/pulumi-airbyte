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

// SourceNytimes Resource
type SourceNytimes struct {
	pulumi.CustomResourceState

	Configuration SourceNytimesConfigurationOutput `pulumi:"configuration"`
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

// NewSourceNytimes registers a new resource with the given unique name, arguments, and options.
func NewSourceNytimes(ctx *pulumi.Context,
	name string, args *SourceNytimesArgs, opts ...pulumi.ResourceOption) (*SourceNytimes, error) {
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
	var resource SourceNytimes
	err := ctx.RegisterResource("airbyte:index/sourceNytimes:SourceNytimes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceNytimes gets an existing SourceNytimes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceNytimes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceNytimesState, opts ...pulumi.ResourceOption) (*SourceNytimes, error) {
	var resource SourceNytimes
	err := ctx.ReadResource("airbyte:index/sourceNytimes:SourceNytimes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceNytimes resources.
type sourceNytimesState struct {
	Configuration *SourceNytimesConfiguration `pulumi:"configuration"`
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

type SourceNytimesState struct {
	Configuration SourceNytimesConfigurationPtrInput
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

func (SourceNytimesState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNytimesState)(nil)).Elem()
}

type sourceNytimesArgs struct {
	Configuration SourceNytimesConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceNytimes resource.
type SourceNytimesArgs struct {
	Configuration SourceNytimesConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceNytimesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceNytimesArgs)(nil)).Elem()
}

type SourceNytimesInput interface {
	pulumi.Input

	ToSourceNytimesOutput() SourceNytimesOutput
	ToSourceNytimesOutputWithContext(ctx context.Context) SourceNytimesOutput
}

func (*SourceNytimes) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNytimes)(nil)).Elem()
}

func (i *SourceNytimes) ToSourceNytimesOutput() SourceNytimesOutput {
	return i.ToSourceNytimesOutputWithContext(context.Background())
}

func (i *SourceNytimes) ToSourceNytimesOutputWithContext(ctx context.Context) SourceNytimesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNytimesOutput)
}

// SourceNytimesArrayInput is an input type that accepts SourceNytimesArray and SourceNytimesArrayOutput values.
// You can construct a concrete instance of `SourceNytimesArrayInput` via:
//
//	SourceNytimesArray{ SourceNytimesArgs{...} }
type SourceNytimesArrayInput interface {
	pulumi.Input

	ToSourceNytimesArrayOutput() SourceNytimesArrayOutput
	ToSourceNytimesArrayOutputWithContext(context.Context) SourceNytimesArrayOutput
}

type SourceNytimesArray []SourceNytimesInput

func (SourceNytimesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNytimes)(nil)).Elem()
}

func (i SourceNytimesArray) ToSourceNytimesArrayOutput() SourceNytimesArrayOutput {
	return i.ToSourceNytimesArrayOutputWithContext(context.Background())
}

func (i SourceNytimesArray) ToSourceNytimesArrayOutputWithContext(ctx context.Context) SourceNytimesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNytimesArrayOutput)
}

// SourceNytimesMapInput is an input type that accepts SourceNytimesMap and SourceNytimesMapOutput values.
// You can construct a concrete instance of `SourceNytimesMapInput` via:
//
//	SourceNytimesMap{ "key": SourceNytimesArgs{...} }
type SourceNytimesMapInput interface {
	pulumi.Input

	ToSourceNytimesMapOutput() SourceNytimesMapOutput
	ToSourceNytimesMapOutputWithContext(context.Context) SourceNytimesMapOutput
}

type SourceNytimesMap map[string]SourceNytimesInput

func (SourceNytimesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNytimes)(nil)).Elem()
}

func (i SourceNytimesMap) ToSourceNytimesMapOutput() SourceNytimesMapOutput {
	return i.ToSourceNytimesMapOutputWithContext(context.Background())
}

func (i SourceNytimesMap) ToSourceNytimesMapOutputWithContext(ctx context.Context) SourceNytimesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceNytimesMapOutput)
}

type SourceNytimesOutput struct{ *pulumi.OutputState }

func (SourceNytimesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceNytimes)(nil)).Elem()
}

func (o SourceNytimesOutput) ToSourceNytimesOutput() SourceNytimesOutput {
	return o
}

func (o SourceNytimesOutput) ToSourceNytimesOutputWithContext(ctx context.Context) SourceNytimesOutput {
	return o
}

func (o SourceNytimesOutput) Configuration() SourceNytimesConfigurationOutput {
	return o.ApplyT(func(v *SourceNytimes) SourceNytimesConfigurationOutput { return v.Configuration }).(SourceNytimesConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceNytimesOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceNytimesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceNytimesOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceNytimesOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceNytimesOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceNytimesOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceNytimes) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceNytimesArrayOutput struct{ *pulumi.OutputState }

func (SourceNytimesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceNytimes)(nil)).Elem()
}

func (o SourceNytimesArrayOutput) ToSourceNytimesArrayOutput() SourceNytimesArrayOutput {
	return o
}

func (o SourceNytimesArrayOutput) ToSourceNytimesArrayOutputWithContext(ctx context.Context) SourceNytimesArrayOutput {
	return o
}

func (o SourceNytimesArrayOutput) Index(i pulumi.IntInput) SourceNytimesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceNytimes {
		return vs[0].([]*SourceNytimes)[vs[1].(int)]
	}).(SourceNytimesOutput)
}

type SourceNytimesMapOutput struct{ *pulumi.OutputState }

func (SourceNytimesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceNytimes)(nil)).Elem()
}

func (o SourceNytimesMapOutput) ToSourceNytimesMapOutput() SourceNytimesMapOutput {
	return o
}

func (o SourceNytimesMapOutput) ToSourceNytimesMapOutputWithContext(ctx context.Context) SourceNytimesMapOutput {
	return o
}

func (o SourceNytimesMapOutput) MapIndex(k pulumi.StringInput) SourceNytimesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceNytimes {
		return vs[0].(map[string]*SourceNytimes)[vs[1].(string)]
	}).(SourceNytimesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNytimesInput)(nil)).Elem(), &SourceNytimes{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNytimesArrayInput)(nil)).Elem(), SourceNytimesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceNytimesMapInput)(nil)).Elem(), SourceNytimesMap{})
	pulumi.RegisterOutputType(SourceNytimesOutput{})
	pulumi.RegisterOutputType(SourceNytimesArrayOutput{})
	pulumi.RegisterOutputType(SourceNytimesMapOutput{})
}
