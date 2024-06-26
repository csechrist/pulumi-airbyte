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

// SourceCloseCom Resource
type SourceCloseCom struct {
	pulumi.CustomResourceState

	Configuration SourceCloseComConfigurationOutput `pulumi:"configuration"`
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

// NewSourceCloseCom registers a new resource with the given unique name, arguments, and options.
func NewSourceCloseCom(ctx *pulumi.Context,
	name string, args *SourceCloseComArgs, opts ...pulumi.ResourceOption) (*SourceCloseCom, error) {
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
	var resource SourceCloseCom
	err := ctx.RegisterResource("airbyte:index/sourceCloseCom:SourceCloseCom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCloseCom gets an existing SourceCloseCom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCloseCom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCloseComState, opts ...pulumi.ResourceOption) (*SourceCloseCom, error) {
	var resource SourceCloseCom
	err := ctx.ReadResource("airbyte:index/sourceCloseCom:SourceCloseCom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCloseCom resources.
type sourceCloseComState struct {
	Configuration *SourceCloseComConfiguration `pulumi:"configuration"`
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

type SourceCloseComState struct {
	Configuration SourceCloseComConfigurationPtrInput
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

func (SourceCloseComState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCloseComState)(nil)).Elem()
}

type sourceCloseComArgs struct {
	Configuration SourceCloseComConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceCloseCom resource.
type SourceCloseComArgs struct {
	Configuration SourceCloseComConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceCloseComArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCloseComArgs)(nil)).Elem()
}

type SourceCloseComInput interface {
	pulumi.Input

	ToSourceCloseComOutput() SourceCloseComOutput
	ToSourceCloseComOutputWithContext(ctx context.Context) SourceCloseComOutput
}

func (*SourceCloseCom) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCloseCom)(nil)).Elem()
}

func (i *SourceCloseCom) ToSourceCloseComOutput() SourceCloseComOutput {
	return i.ToSourceCloseComOutputWithContext(context.Background())
}

func (i *SourceCloseCom) ToSourceCloseComOutputWithContext(ctx context.Context) SourceCloseComOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCloseComOutput)
}

// SourceCloseComArrayInput is an input type that accepts SourceCloseComArray and SourceCloseComArrayOutput values.
// You can construct a concrete instance of `SourceCloseComArrayInput` via:
//
//	SourceCloseComArray{ SourceCloseComArgs{...} }
type SourceCloseComArrayInput interface {
	pulumi.Input

	ToSourceCloseComArrayOutput() SourceCloseComArrayOutput
	ToSourceCloseComArrayOutputWithContext(context.Context) SourceCloseComArrayOutput
}

type SourceCloseComArray []SourceCloseComInput

func (SourceCloseComArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCloseCom)(nil)).Elem()
}

func (i SourceCloseComArray) ToSourceCloseComArrayOutput() SourceCloseComArrayOutput {
	return i.ToSourceCloseComArrayOutputWithContext(context.Background())
}

func (i SourceCloseComArray) ToSourceCloseComArrayOutputWithContext(ctx context.Context) SourceCloseComArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCloseComArrayOutput)
}

// SourceCloseComMapInput is an input type that accepts SourceCloseComMap and SourceCloseComMapOutput values.
// You can construct a concrete instance of `SourceCloseComMapInput` via:
//
//	SourceCloseComMap{ "key": SourceCloseComArgs{...} }
type SourceCloseComMapInput interface {
	pulumi.Input

	ToSourceCloseComMapOutput() SourceCloseComMapOutput
	ToSourceCloseComMapOutputWithContext(context.Context) SourceCloseComMapOutput
}

type SourceCloseComMap map[string]SourceCloseComInput

func (SourceCloseComMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCloseCom)(nil)).Elem()
}

func (i SourceCloseComMap) ToSourceCloseComMapOutput() SourceCloseComMapOutput {
	return i.ToSourceCloseComMapOutputWithContext(context.Background())
}

func (i SourceCloseComMap) ToSourceCloseComMapOutputWithContext(ctx context.Context) SourceCloseComMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCloseComMapOutput)
}

type SourceCloseComOutput struct{ *pulumi.OutputState }

func (SourceCloseComOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCloseCom)(nil)).Elem()
}

func (o SourceCloseComOutput) ToSourceCloseComOutput() SourceCloseComOutput {
	return o
}

func (o SourceCloseComOutput) ToSourceCloseComOutputWithContext(ctx context.Context) SourceCloseComOutput {
	return o
}

func (o SourceCloseComOutput) Configuration() SourceCloseComConfigurationOutput {
	return o.ApplyT(func(v *SourceCloseCom) SourceCloseComConfigurationOutput { return v.Configuration }).(SourceCloseComConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceCloseComOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceCloseComOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceCloseComOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceCloseComOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceCloseComOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceCloseComOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCloseCom) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceCloseComArrayOutput struct{ *pulumi.OutputState }

func (SourceCloseComArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceCloseCom)(nil)).Elem()
}

func (o SourceCloseComArrayOutput) ToSourceCloseComArrayOutput() SourceCloseComArrayOutput {
	return o
}

func (o SourceCloseComArrayOutput) ToSourceCloseComArrayOutputWithContext(ctx context.Context) SourceCloseComArrayOutput {
	return o
}

func (o SourceCloseComArrayOutput) Index(i pulumi.IntInput) SourceCloseComOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceCloseCom {
		return vs[0].([]*SourceCloseCom)[vs[1].(int)]
	}).(SourceCloseComOutput)
}

type SourceCloseComMapOutput struct{ *pulumi.OutputState }

func (SourceCloseComMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceCloseCom)(nil)).Elem()
}

func (o SourceCloseComMapOutput) ToSourceCloseComMapOutput() SourceCloseComMapOutput {
	return o
}

func (o SourceCloseComMapOutput) ToSourceCloseComMapOutputWithContext(ctx context.Context) SourceCloseComMapOutput {
	return o
}

func (o SourceCloseComMapOutput) MapIndex(k pulumi.StringInput) SourceCloseComOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceCloseCom {
		return vs[0].(map[string]*SourceCloseCom)[vs[1].(string)]
	}).(SourceCloseComOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCloseComInput)(nil)).Elem(), &SourceCloseCom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCloseComArrayInput)(nil)).Elem(), SourceCloseComArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCloseComMapInput)(nil)).Elem(), SourceCloseComMap{})
	pulumi.RegisterOutputType(SourceCloseComOutput{})
	pulumi.RegisterOutputType(SourceCloseComArrayOutput{})
	pulumi.RegisterOutputType(SourceCloseComMapOutput{})
}
