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

// SourceGoogleWebfonts Resource
type SourceGoogleWebfonts struct {
	pulumi.CustomResourceState

	Configuration SourceGoogleWebfontsConfigurationOutput `pulumi:"configuration"`
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

// NewSourceGoogleWebfonts registers a new resource with the given unique name, arguments, and options.
func NewSourceGoogleWebfonts(ctx *pulumi.Context,
	name string, args *SourceGoogleWebfontsArgs, opts ...pulumi.ResourceOption) (*SourceGoogleWebfonts, error) {
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
	var resource SourceGoogleWebfonts
	err := ctx.RegisterResource("airbyte:index/sourceGoogleWebfonts:SourceGoogleWebfonts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceGoogleWebfonts gets an existing SourceGoogleWebfonts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceGoogleWebfonts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceGoogleWebfontsState, opts ...pulumi.ResourceOption) (*SourceGoogleWebfonts, error) {
	var resource SourceGoogleWebfonts
	err := ctx.ReadResource("airbyte:index/sourceGoogleWebfonts:SourceGoogleWebfonts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceGoogleWebfonts resources.
type sourceGoogleWebfontsState struct {
	Configuration *SourceGoogleWebfontsConfiguration `pulumi:"configuration"`
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

type SourceGoogleWebfontsState struct {
	Configuration SourceGoogleWebfontsConfigurationPtrInput
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

func (SourceGoogleWebfontsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleWebfontsState)(nil)).Elem()
}

type sourceGoogleWebfontsArgs struct {
	Configuration SourceGoogleWebfontsConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceGoogleWebfonts resource.
type SourceGoogleWebfontsArgs struct {
	Configuration SourceGoogleWebfontsConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceGoogleWebfontsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceGoogleWebfontsArgs)(nil)).Elem()
}

type SourceGoogleWebfontsInput interface {
	pulumi.Input

	ToSourceGoogleWebfontsOutput() SourceGoogleWebfontsOutput
	ToSourceGoogleWebfontsOutputWithContext(ctx context.Context) SourceGoogleWebfontsOutput
}

func (*SourceGoogleWebfonts) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleWebfonts)(nil)).Elem()
}

func (i *SourceGoogleWebfonts) ToSourceGoogleWebfontsOutput() SourceGoogleWebfontsOutput {
	return i.ToSourceGoogleWebfontsOutputWithContext(context.Background())
}

func (i *SourceGoogleWebfonts) ToSourceGoogleWebfontsOutputWithContext(ctx context.Context) SourceGoogleWebfontsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleWebfontsOutput)
}

// SourceGoogleWebfontsArrayInput is an input type that accepts SourceGoogleWebfontsArray and SourceGoogleWebfontsArrayOutput values.
// You can construct a concrete instance of `SourceGoogleWebfontsArrayInput` via:
//
//	SourceGoogleWebfontsArray{ SourceGoogleWebfontsArgs{...} }
type SourceGoogleWebfontsArrayInput interface {
	pulumi.Input

	ToSourceGoogleWebfontsArrayOutput() SourceGoogleWebfontsArrayOutput
	ToSourceGoogleWebfontsArrayOutputWithContext(context.Context) SourceGoogleWebfontsArrayOutput
}

type SourceGoogleWebfontsArray []SourceGoogleWebfontsInput

func (SourceGoogleWebfontsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleWebfonts)(nil)).Elem()
}

func (i SourceGoogleWebfontsArray) ToSourceGoogleWebfontsArrayOutput() SourceGoogleWebfontsArrayOutput {
	return i.ToSourceGoogleWebfontsArrayOutputWithContext(context.Background())
}

func (i SourceGoogleWebfontsArray) ToSourceGoogleWebfontsArrayOutputWithContext(ctx context.Context) SourceGoogleWebfontsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleWebfontsArrayOutput)
}

// SourceGoogleWebfontsMapInput is an input type that accepts SourceGoogleWebfontsMap and SourceGoogleWebfontsMapOutput values.
// You can construct a concrete instance of `SourceGoogleWebfontsMapInput` via:
//
//	SourceGoogleWebfontsMap{ "key": SourceGoogleWebfontsArgs{...} }
type SourceGoogleWebfontsMapInput interface {
	pulumi.Input

	ToSourceGoogleWebfontsMapOutput() SourceGoogleWebfontsMapOutput
	ToSourceGoogleWebfontsMapOutputWithContext(context.Context) SourceGoogleWebfontsMapOutput
}

type SourceGoogleWebfontsMap map[string]SourceGoogleWebfontsInput

func (SourceGoogleWebfontsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleWebfonts)(nil)).Elem()
}

func (i SourceGoogleWebfontsMap) ToSourceGoogleWebfontsMapOutput() SourceGoogleWebfontsMapOutput {
	return i.ToSourceGoogleWebfontsMapOutputWithContext(context.Background())
}

func (i SourceGoogleWebfontsMap) ToSourceGoogleWebfontsMapOutputWithContext(ctx context.Context) SourceGoogleWebfontsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceGoogleWebfontsMapOutput)
}

type SourceGoogleWebfontsOutput struct{ *pulumi.OutputState }

func (SourceGoogleWebfontsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceGoogleWebfonts)(nil)).Elem()
}

func (o SourceGoogleWebfontsOutput) ToSourceGoogleWebfontsOutput() SourceGoogleWebfontsOutput {
	return o
}

func (o SourceGoogleWebfontsOutput) ToSourceGoogleWebfontsOutputWithContext(ctx context.Context) SourceGoogleWebfontsOutput {
	return o
}

func (o SourceGoogleWebfontsOutput) Configuration() SourceGoogleWebfontsConfigurationOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) SourceGoogleWebfontsConfigurationOutput { return v.Configuration }).(SourceGoogleWebfontsConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceGoogleWebfontsOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceGoogleWebfontsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceGoogleWebfontsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceGoogleWebfontsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceGoogleWebfontsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceGoogleWebfontsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceGoogleWebfonts) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceGoogleWebfontsArrayOutput struct{ *pulumi.OutputState }

func (SourceGoogleWebfontsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceGoogleWebfonts)(nil)).Elem()
}

func (o SourceGoogleWebfontsArrayOutput) ToSourceGoogleWebfontsArrayOutput() SourceGoogleWebfontsArrayOutput {
	return o
}

func (o SourceGoogleWebfontsArrayOutput) ToSourceGoogleWebfontsArrayOutputWithContext(ctx context.Context) SourceGoogleWebfontsArrayOutput {
	return o
}

func (o SourceGoogleWebfontsArrayOutput) Index(i pulumi.IntInput) SourceGoogleWebfontsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceGoogleWebfonts {
		return vs[0].([]*SourceGoogleWebfonts)[vs[1].(int)]
	}).(SourceGoogleWebfontsOutput)
}

type SourceGoogleWebfontsMapOutput struct{ *pulumi.OutputState }

func (SourceGoogleWebfontsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceGoogleWebfonts)(nil)).Elem()
}

func (o SourceGoogleWebfontsMapOutput) ToSourceGoogleWebfontsMapOutput() SourceGoogleWebfontsMapOutput {
	return o
}

func (o SourceGoogleWebfontsMapOutput) ToSourceGoogleWebfontsMapOutputWithContext(ctx context.Context) SourceGoogleWebfontsMapOutput {
	return o
}

func (o SourceGoogleWebfontsMapOutput) MapIndex(k pulumi.StringInput) SourceGoogleWebfontsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceGoogleWebfonts {
		return vs[0].(map[string]*SourceGoogleWebfonts)[vs[1].(string)]
	}).(SourceGoogleWebfontsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleWebfontsInput)(nil)).Elem(), &SourceGoogleWebfonts{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleWebfontsArrayInput)(nil)).Elem(), SourceGoogleWebfontsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceGoogleWebfontsMapInput)(nil)).Elem(), SourceGoogleWebfontsMap{})
	pulumi.RegisterOutputType(SourceGoogleWebfontsOutput{})
	pulumi.RegisterOutputType(SourceGoogleWebfontsArrayOutput{})
	pulumi.RegisterOutputType(SourceGoogleWebfontsMapOutput{})
}
