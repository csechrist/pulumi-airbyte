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

// SourceShopify Resource
type SourceShopify struct {
	pulumi.CustomResourceState

	Configuration SourceShopifyConfigurationOutput `pulumi:"configuration"`
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

// NewSourceShopify registers a new resource with the given unique name, arguments, and options.
func NewSourceShopify(ctx *pulumi.Context,
	name string, args *SourceShopifyArgs, opts ...pulumi.ResourceOption) (*SourceShopify, error) {
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
	var resource SourceShopify
	err := ctx.RegisterResource("airbyte:index/sourceShopify:SourceShopify", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceShopify gets an existing SourceShopify resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceShopify(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceShopifyState, opts ...pulumi.ResourceOption) (*SourceShopify, error) {
	var resource SourceShopify
	err := ctx.ReadResource("airbyte:index/sourceShopify:SourceShopify", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceShopify resources.
type sourceShopifyState struct {
	Configuration *SourceShopifyConfiguration `pulumi:"configuration"`
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

type SourceShopifyState struct {
	Configuration SourceShopifyConfigurationPtrInput
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

func (SourceShopifyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceShopifyState)(nil)).Elem()
}

type sourceShopifyArgs struct {
	Configuration SourceShopifyConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceShopify resource.
type SourceShopifyArgs struct {
	Configuration SourceShopifyConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceShopifyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceShopifyArgs)(nil)).Elem()
}

type SourceShopifyInput interface {
	pulumi.Input

	ToSourceShopifyOutput() SourceShopifyOutput
	ToSourceShopifyOutputWithContext(ctx context.Context) SourceShopifyOutput
}

func (*SourceShopify) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceShopify)(nil)).Elem()
}

func (i *SourceShopify) ToSourceShopifyOutput() SourceShopifyOutput {
	return i.ToSourceShopifyOutputWithContext(context.Background())
}

func (i *SourceShopify) ToSourceShopifyOutputWithContext(ctx context.Context) SourceShopifyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceShopifyOutput)
}

// SourceShopifyArrayInput is an input type that accepts SourceShopifyArray and SourceShopifyArrayOutput values.
// You can construct a concrete instance of `SourceShopifyArrayInput` via:
//
//	SourceShopifyArray{ SourceShopifyArgs{...} }
type SourceShopifyArrayInput interface {
	pulumi.Input

	ToSourceShopifyArrayOutput() SourceShopifyArrayOutput
	ToSourceShopifyArrayOutputWithContext(context.Context) SourceShopifyArrayOutput
}

type SourceShopifyArray []SourceShopifyInput

func (SourceShopifyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceShopify)(nil)).Elem()
}

func (i SourceShopifyArray) ToSourceShopifyArrayOutput() SourceShopifyArrayOutput {
	return i.ToSourceShopifyArrayOutputWithContext(context.Background())
}

func (i SourceShopifyArray) ToSourceShopifyArrayOutputWithContext(ctx context.Context) SourceShopifyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceShopifyArrayOutput)
}

// SourceShopifyMapInput is an input type that accepts SourceShopifyMap and SourceShopifyMapOutput values.
// You can construct a concrete instance of `SourceShopifyMapInput` via:
//
//	SourceShopifyMap{ "key": SourceShopifyArgs{...} }
type SourceShopifyMapInput interface {
	pulumi.Input

	ToSourceShopifyMapOutput() SourceShopifyMapOutput
	ToSourceShopifyMapOutputWithContext(context.Context) SourceShopifyMapOutput
}

type SourceShopifyMap map[string]SourceShopifyInput

func (SourceShopifyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceShopify)(nil)).Elem()
}

func (i SourceShopifyMap) ToSourceShopifyMapOutput() SourceShopifyMapOutput {
	return i.ToSourceShopifyMapOutputWithContext(context.Background())
}

func (i SourceShopifyMap) ToSourceShopifyMapOutputWithContext(ctx context.Context) SourceShopifyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceShopifyMapOutput)
}

type SourceShopifyOutput struct{ *pulumi.OutputState }

func (SourceShopifyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceShopify)(nil)).Elem()
}

func (o SourceShopifyOutput) ToSourceShopifyOutput() SourceShopifyOutput {
	return o
}

func (o SourceShopifyOutput) ToSourceShopifyOutputWithContext(ctx context.Context) SourceShopifyOutput {
	return o
}

func (o SourceShopifyOutput) Configuration() SourceShopifyConfigurationOutput {
	return o.ApplyT(func(v *SourceShopify) SourceShopifyConfigurationOutput { return v.Configuration }).(SourceShopifyConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceShopifyOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceShopifyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceShopifyOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceShopifyOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceShopifyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceShopifyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceShopify) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceShopifyArrayOutput struct{ *pulumi.OutputState }

func (SourceShopifyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceShopify)(nil)).Elem()
}

func (o SourceShopifyArrayOutput) ToSourceShopifyArrayOutput() SourceShopifyArrayOutput {
	return o
}

func (o SourceShopifyArrayOutput) ToSourceShopifyArrayOutputWithContext(ctx context.Context) SourceShopifyArrayOutput {
	return o
}

func (o SourceShopifyArrayOutput) Index(i pulumi.IntInput) SourceShopifyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceShopify {
		return vs[0].([]*SourceShopify)[vs[1].(int)]
	}).(SourceShopifyOutput)
}

type SourceShopifyMapOutput struct{ *pulumi.OutputState }

func (SourceShopifyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceShopify)(nil)).Elem()
}

func (o SourceShopifyMapOutput) ToSourceShopifyMapOutput() SourceShopifyMapOutput {
	return o
}

func (o SourceShopifyMapOutput) ToSourceShopifyMapOutputWithContext(ctx context.Context) SourceShopifyMapOutput {
	return o
}

func (o SourceShopifyMapOutput) MapIndex(k pulumi.StringInput) SourceShopifyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceShopify {
		return vs[0].(map[string]*SourceShopify)[vs[1].(string)]
	}).(SourceShopifyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceShopifyInput)(nil)).Elem(), &SourceShopify{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceShopifyArrayInput)(nil)).Elem(), SourceShopifyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceShopifyMapInput)(nil)).Elem(), SourceShopifyMap{})
	pulumi.RegisterOutputType(SourceShopifyOutput{})
	pulumi.RegisterOutputType(SourceShopifyArrayOutput{})
	pulumi.RegisterOutputType(SourceShopifyMapOutput{})
}
