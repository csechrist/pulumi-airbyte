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

// SourceOutbrainAmplify Resource
type SourceOutbrainAmplify struct {
	pulumi.CustomResourceState

	Configuration SourceOutbrainAmplifyConfigurationOutput `pulumi:"configuration"`
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

// NewSourceOutbrainAmplify registers a new resource with the given unique name, arguments, and options.
func NewSourceOutbrainAmplify(ctx *pulumi.Context,
	name string, args *SourceOutbrainAmplifyArgs, opts ...pulumi.ResourceOption) (*SourceOutbrainAmplify, error) {
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
	var resource SourceOutbrainAmplify
	err := ctx.RegisterResource("airbyte:index/sourceOutbrainAmplify:SourceOutbrainAmplify", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceOutbrainAmplify gets an existing SourceOutbrainAmplify resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceOutbrainAmplify(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceOutbrainAmplifyState, opts ...pulumi.ResourceOption) (*SourceOutbrainAmplify, error) {
	var resource SourceOutbrainAmplify
	err := ctx.ReadResource("airbyte:index/sourceOutbrainAmplify:SourceOutbrainAmplify", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceOutbrainAmplify resources.
type sourceOutbrainAmplifyState struct {
	Configuration *SourceOutbrainAmplifyConfiguration `pulumi:"configuration"`
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

type SourceOutbrainAmplifyState struct {
	Configuration SourceOutbrainAmplifyConfigurationPtrInput
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

func (SourceOutbrainAmplifyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutbrainAmplifyState)(nil)).Elem()
}

type sourceOutbrainAmplifyArgs struct {
	Configuration SourceOutbrainAmplifyConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceOutbrainAmplify resource.
type SourceOutbrainAmplifyArgs struct {
	Configuration SourceOutbrainAmplifyConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceOutbrainAmplifyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceOutbrainAmplifyArgs)(nil)).Elem()
}

type SourceOutbrainAmplifyInput interface {
	pulumi.Input

	ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput
	ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput
}

func (*SourceOutbrainAmplify) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutbrainAmplify)(nil)).Elem()
}

func (i *SourceOutbrainAmplify) ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput {
	return i.ToSourceOutbrainAmplifyOutputWithContext(context.Background())
}

func (i *SourceOutbrainAmplify) ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutbrainAmplifyOutput)
}

// SourceOutbrainAmplifyArrayInput is an input type that accepts SourceOutbrainAmplifyArray and SourceOutbrainAmplifyArrayOutput values.
// You can construct a concrete instance of `SourceOutbrainAmplifyArrayInput` via:
//
//	SourceOutbrainAmplifyArray{ SourceOutbrainAmplifyArgs{...} }
type SourceOutbrainAmplifyArrayInput interface {
	pulumi.Input

	ToSourceOutbrainAmplifyArrayOutput() SourceOutbrainAmplifyArrayOutput
	ToSourceOutbrainAmplifyArrayOutputWithContext(context.Context) SourceOutbrainAmplifyArrayOutput
}

type SourceOutbrainAmplifyArray []SourceOutbrainAmplifyInput

func (SourceOutbrainAmplifyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOutbrainAmplify)(nil)).Elem()
}

func (i SourceOutbrainAmplifyArray) ToSourceOutbrainAmplifyArrayOutput() SourceOutbrainAmplifyArrayOutput {
	return i.ToSourceOutbrainAmplifyArrayOutputWithContext(context.Background())
}

func (i SourceOutbrainAmplifyArray) ToSourceOutbrainAmplifyArrayOutputWithContext(ctx context.Context) SourceOutbrainAmplifyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutbrainAmplifyArrayOutput)
}

// SourceOutbrainAmplifyMapInput is an input type that accepts SourceOutbrainAmplifyMap and SourceOutbrainAmplifyMapOutput values.
// You can construct a concrete instance of `SourceOutbrainAmplifyMapInput` via:
//
//	SourceOutbrainAmplifyMap{ "key": SourceOutbrainAmplifyArgs{...} }
type SourceOutbrainAmplifyMapInput interface {
	pulumi.Input

	ToSourceOutbrainAmplifyMapOutput() SourceOutbrainAmplifyMapOutput
	ToSourceOutbrainAmplifyMapOutputWithContext(context.Context) SourceOutbrainAmplifyMapOutput
}

type SourceOutbrainAmplifyMap map[string]SourceOutbrainAmplifyInput

func (SourceOutbrainAmplifyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOutbrainAmplify)(nil)).Elem()
}

func (i SourceOutbrainAmplifyMap) ToSourceOutbrainAmplifyMapOutput() SourceOutbrainAmplifyMapOutput {
	return i.ToSourceOutbrainAmplifyMapOutputWithContext(context.Background())
}

func (i SourceOutbrainAmplifyMap) ToSourceOutbrainAmplifyMapOutputWithContext(ctx context.Context) SourceOutbrainAmplifyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutbrainAmplifyMapOutput)
}

type SourceOutbrainAmplifyOutput struct{ *pulumi.OutputState }

func (SourceOutbrainAmplifyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceOutbrainAmplify)(nil)).Elem()
}

func (o SourceOutbrainAmplifyOutput) ToSourceOutbrainAmplifyOutput() SourceOutbrainAmplifyOutput {
	return o
}

func (o SourceOutbrainAmplifyOutput) ToSourceOutbrainAmplifyOutputWithContext(ctx context.Context) SourceOutbrainAmplifyOutput {
	return o
}

func (o SourceOutbrainAmplifyOutput) Configuration() SourceOutbrainAmplifyConfigurationOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) SourceOutbrainAmplifyConfigurationOutput { return v.Configuration }).(SourceOutbrainAmplifyConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceOutbrainAmplifyOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceOutbrainAmplifyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceOutbrainAmplifyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceOutbrainAmplifyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceOutbrainAmplifyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceOutbrainAmplifyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceOutbrainAmplify) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceOutbrainAmplifyArrayOutput struct{ *pulumi.OutputState }

func (SourceOutbrainAmplifyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceOutbrainAmplify)(nil)).Elem()
}

func (o SourceOutbrainAmplifyArrayOutput) ToSourceOutbrainAmplifyArrayOutput() SourceOutbrainAmplifyArrayOutput {
	return o
}

func (o SourceOutbrainAmplifyArrayOutput) ToSourceOutbrainAmplifyArrayOutputWithContext(ctx context.Context) SourceOutbrainAmplifyArrayOutput {
	return o
}

func (o SourceOutbrainAmplifyArrayOutput) Index(i pulumi.IntInput) SourceOutbrainAmplifyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceOutbrainAmplify {
		return vs[0].([]*SourceOutbrainAmplify)[vs[1].(int)]
	}).(SourceOutbrainAmplifyOutput)
}

type SourceOutbrainAmplifyMapOutput struct{ *pulumi.OutputState }

func (SourceOutbrainAmplifyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceOutbrainAmplify)(nil)).Elem()
}

func (o SourceOutbrainAmplifyMapOutput) ToSourceOutbrainAmplifyMapOutput() SourceOutbrainAmplifyMapOutput {
	return o
}

func (o SourceOutbrainAmplifyMapOutput) ToSourceOutbrainAmplifyMapOutputWithContext(ctx context.Context) SourceOutbrainAmplifyMapOutput {
	return o
}

func (o SourceOutbrainAmplifyMapOutput) MapIndex(k pulumi.StringInput) SourceOutbrainAmplifyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceOutbrainAmplify {
		return vs[0].(map[string]*SourceOutbrainAmplify)[vs[1].(string)]
	}).(SourceOutbrainAmplifyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutbrainAmplifyInput)(nil)).Elem(), &SourceOutbrainAmplify{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutbrainAmplifyArrayInput)(nil)).Elem(), SourceOutbrainAmplifyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceOutbrainAmplifyMapInput)(nil)).Elem(), SourceOutbrainAmplifyMap{})
	pulumi.RegisterOutputType(SourceOutbrainAmplifyOutput{})
	pulumi.RegisterOutputType(SourceOutbrainAmplifyArrayOutput{})
	pulumi.RegisterOutputType(SourceOutbrainAmplifyMapOutput{})
}
