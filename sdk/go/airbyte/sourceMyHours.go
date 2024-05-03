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

// SourceMyHours Resource
type SourceMyHours struct {
	pulumi.CustomResourceState

	Configuration SourceMyHoursConfigurationOutput `pulumi:"configuration"`
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

// NewSourceMyHours registers a new resource with the given unique name, arguments, and options.
func NewSourceMyHours(ctx *pulumi.Context,
	name string, args *SourceMyHoursArgs, opts ...pulumi.ResourceOption) (*SourceMyHours, error) {
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
	var resource SourceMyHours
	err := ctx.RegisterResource("airbyte:index/sourceMyHours:SourceMyHours", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMyHours gets an existing SourceMyHours resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMyHours(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMyHoursState, opts ...pulumi.ResourceOption) (*SourceMyHours, error) {
	var resource SourceMyHours
	err := ctx.ReadResource("airbyte:index/sourceMyHours:SourceMyHours", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMyHours resources.
type sourceMyHoursState struct {
	Configuration *SourceMyHoursConfiguration `pulumi:"configuration"`
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

type SourceMyHoursState struct {
	Configuration SourceMyHoursConfigurationPtrInput
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

func (SourceMyHoursState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMyHoursState)(nil)).Elem()
}

type sourceMyHoursArgs struct {
	Configuration SourceMyHoursConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMyHours resource.
type SourceMyHoursArgs struct {
	Configuration SourceMyHoursConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMyHoursArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMyHoursArgs)(nil)).Elem()
}

type SourceMyHoursInput interface {
	pulumi.Input

	ToSourceMyHoursOutput() SourceMyHoursOutput
	ToSourceMyHoursOutputWithContext(ctx context.Context) SourceMyHoursOutput
}

func (*SourceMyHours) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMyHours)(nil)).Elem()
}

func (i *SourceMyHours) ToSourceMyHoursOutput() SourceMyHoursOutput {
	return i.ToSourceMyHoursOutputWithContext(context.Background())
}

func (i *SourceMyHours) ToSourceMyHoursOutputWithContext(ctx context.Context) SourceMyHoursOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMyHoursOutput)
}

// SourceMyHoursArrayInput is an input type that accepts SourceMyHoursArray and SourceMyHoursArrayOutput values.
// You can construct a concrete instance of `SourceMyHoursArrayInput` via:
//
//	SourceMyHoursArray{ SourceMyHoursArgs{...} }
type SourceMyHoursArrayInput interface {
	pulumi.Input

	ToSourceMyHoursArrayOutput() SourceMyHoursArrayOutput
	ToSourceMyHoursArrayOutputWithContext(context.Context) SourceMyHoursArrayOutput
}

type SourceMyHoursArray []SourceMyHoursInput

func (SourceMyHoursArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMyHours)(nil)).Elem()
}

func (i SourceMyHoursArray) ToSourceMyHoursArrayOutput() SourceMyHoursArrayOutput {
	return i.ToSourceMyHoursArrayOutputWithContext(context.Background())
}

func (i SourceMyHoursArray) ToSourceMyHoursArrayOutputWithContext(ctx context.Context) SourceMyHoursArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMyHoursArrayOutput)
}

// SourceMyHoursMapInput is an input type that accepts SourceMyHoursMap and SourceMyHoursMapOutput values.
// You can construct a concrete instance of `SourceMyHoursMapInput` via:
//
//	SourceMyHoursMap{ "key": SourceMyHoursArgs{...} }
type SourceMyHoursMapInput interface {
	pulumi.Input

	ToSourceMyHoursMapOutput() SourceMyHoursMapOutput
	ToSourceMyHoursMapOutputWithContext(context.Context) SourceMyHoursMapOutput
}

type SourceMyHoursMap map[string]SourceMyHoursInput

func (SourceMyHoursMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMyHours)(nil)).Elem()
}

func (i SourceMyHoursMap) ToSourceMyHoursMapOutput() SourceMyHoursMapOutput {
	return i.ToSourceMyHoursMapOutputWithContext(context.Background())
}

func (i SourceMyHoursMap) ToSourceMyHoursMapOutputWithContext(ctx context.Context) SourceMyHoursMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMyHoursMapOutput)
}

type SourceMyHoursOutput struct{ *pulumi.OutputState }

func (SourceMyHoursOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMyHours)(nil)).Elem()
}

func (o SourceMyHoursOutput) ToSourceMyHoursOutput() SourceMyHoursOutput {
	return o
}

func (o SourceMyHoursOutput) ToSourceMyHoursOutputWithContext(ctx context.Context) SourceMyHoursOutput {
	return o
}

func (o SourceMyHoursOutput) Configuration() SourceMyHoursConfigurationOutput {
	return o.ApplyT(func(v *SourceMyHours) SourceMyHoursConfigurationOutput { return v.Configuration }).(SourceMyHoursConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceMyHoursOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceMyHoursOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceMyHoursOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMyHoursOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMyHoursOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMyHoursOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMyHours) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceMyHoursArrayOutput struct{ *pulumi.OutputState }

func (SourceMyHoursArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMyHours)(nil)).Elem()
}

func (o SourceMyHoursArrayOutput) ToSourceMyHoursArrayOutput() SourceMyHoursArrayOutput {
	return o
}

func (o SourceMyHoursArrayOutput) ToSourceMyHoursArrayOutputWithContext(ctx context.Context) SourceMyHoursArrayOutput {
	return o
}

func (o SourceMyHoursArrayOutput) Index(i pulumi.IntInput) SourceMyHoursOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceMyHours {
		return vs[0].([]*SourceMyHours)[vs[1].(int)]
	}).(SourceMyHoursOutput)
}

type SourceMyHoursMapOutput struct{ *pulumi.OutputState }

func (SourceMyHoursMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMyHours)(nil)).Elem()
}

func (o SourceMyHoursMapOutput) ToSourceMyHoursMapOutput() SourceMyHoursMapOutput {
	return o
}

func (o SourceMyHoursMapOutput) ToSourceMyHoursMapOutputWithContext(ctx context.Context) SourceMyHoursMapOutput {
	return o
}

func (o SourceMyHoursMapOutput) MapIndex(k pulumi.StringInput) SourceMyHoursOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceMyHours {
		return vs[0].(map[string]*SourceMyHours)[vs[1].(string)]
	}).(SourceMyHoursOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMyHoursInput)(nil)).Elem(), &SourceMyHours{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMyHoursArrayInput)(nil)).Elem(), SourceMyHoursArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMyHoursMapInput)(nil)).Elem(), SourceMyHoursMap{})
	pulumi.RegisterOutputType(SourceMyHoursOutput{})
	pulumi.RegisterOutputType(SourceMyHoursArrayOutput{})
	pulumi.RegisterOutputType(SourceMyHoursMapOutput{})
}