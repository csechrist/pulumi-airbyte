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

// SourceAmazonSqs Resource
type SourceAmazonSqs struct {
	pulumi.CustomResourceState

	Configuration SourceAmazonSqsConfigurationOutput `pulumi:"configuration"`
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

// NewSourceAmazonSqs registers a new resource with the given unique name, arguments, and options.
func NewSourceAmazonSqs(ctx *pulumi.Context,
	name string, args *SourceAmazonSqsArgs, opts ...pulumi.ResourceOption) (*SourceAmazonSqs, error) {
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
	var resource SourceAmazonSqs
	err := ctx.RegisterResource("airbyte:index/sourceAmazonSqs:SourceAmazonSqs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceAmazonSqs gets an existing SourceAmazonSqs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceAmazonSqs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceAmazonSqsState, opts ...pulumi.ResourceOption) (*SourceAmazonSqs, error) {
	var resource SourceAmazonSqs
	err := ctx.ReadResource("airbyte:index/sourceAmazonSqs:SourceAmazonSqs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceAmazonSqs resources.
type sourceAmazonSqsState struct {
	Configuration *SourceAmazonSqsConfiguration `pulumi:"configuration"`
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

type SourceAmazonSqsState struct {
	Configuration SourceAmazonSqsConfigurationPtrInput
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

func (SourceAmazonSqsState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmazonSqsState)(nil)).Elem()
}

type sourceAmazonSqsArgs struct {
	Configuration SourceAmazonSqsConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceAmazonSqs resource.
type SourceAmazonSqsArgs struct {
	Configuration SourceAmazonSqsConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceAmazonSqsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceAmazonSqsArgs)(nil)).Elem()
}

type SourceAmazonSqsInput interface {
	pulumi.Input

	ToSourceAmazonSqsOutput() SourceAmazonSqsOutput
	ToSourceAmazonSqsOutputWithContext(ctx context.Context) SourceAmazonSqsOutput
}

func (*SourceAmazonSqs) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmazonSqs)(nil)).Elem()
}

func (i *SourceAmazonSqs) ToSourceAmazonSqsOutput() SourceAmazonSqsOutput {
	return i.ToSourceAmazonSqsOutputWithContext(context.Background())
}

func (i *SourceAmazonSqs) ToSourceAmazonSqsOutputWithContext(ctx context.Context) SourceAmazonSqsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonSqsOutput)
}

// SourceAmazonSqsArrayInput is an input type that accepts SourceAmazonSqsArray and SourceAmazonSqsArrayOutput values.
// You can construct a concrete instance of `SourceAmazonSqsArrayInput` via:
//
//	SourceAmazonSqsArray{ SourceAmazonSqsArgs{...} }
type SourceAmazonSqsArrayInput interface {
	pulumi.Input

	ToSourceAmazonSqsArrayOutput() SourceAmazonSqsArrayOutput
	ToSourceAmazonSqsArrayOutputWithContext(context.Context) SourceAmazonSqsArrayOutput
}

type SourceAmazonSqsArray []SourceAmazonSqsInput

func (SourceAmazonSqsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmazonSqs)(nil)).Elem()
}

func (i SourceAmazonSqsArray) ToSourceAmazonSqsArrayOutput() SourceAmazonSqsArrayOutput {
	return i.ToSourceAmazonSqsArrayOutputWithContext(context.Background())
}

func (i SourceAmazonSqsArray) ToSourceAmazonSqsArrayOutputWithContext(ctx context.Context) SourceAmazonSqsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonSqsArrayOutput)
}

// SourceAmazonSqsMapInput is an input type that accepts SourceAmazonSqsMap and SourceAmazonSqsMapOutput values.
// You can construct a concrete instance of `SourceAmazonSqsMapInput` via:
//
//	SourceAmazonSqsMap{ "key": SourceAmazonSqsArgs{...} }
type SourceAmazonSqsMapInput interface {
	pulumi.Input

	ToSourceAmazonSqsMapOutput() SourceAmazonSqsMapOutput
	ToSourceAmazonSqsMapOutputWithContext(context.Context) SourceAmazonSqsMapOutput
}

type SourceAmazonSqsMap map[string]SourceAmazonSqsInput

func (SourceAmazonSqsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmazonSqs)(nil)).Elem()
}

func (i SourceAmazonSqsMap) ToSourceAmazonSqsMapOutput() SourceAmazonSqsMapOutput {
	return i.ToSourceAmazonSqsMapOutputWithContext(context.Background())
}

func (i SourceAmazonSqsMap) ToSourceAmazonSqsMapOutputWithContext(ctx context.Context) SourceAmazonSqsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceAmazonSqsMapOutput)
}

type SourceAmazonSqsOutput struct{ *pulumi.OutputState }

func (SourceAmazonSqsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceAmazonSqs)(nil)).Elem()
}

func (o SourceAmazonSqsOutput) ToSourceAmazonSqsOutput() SourceAmazonSqsOutput {
	return o
}

func (o SourceAmazonSqsOutput) ToSourceAmazonSqsOutputWithContext(ctx context.Context) SourceAmazonSqsOutput {
	return o
}

func (o SourceAmazonSqsOutput) Configuration() SourceAmazonSqsConfigurationOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) SourceAmazonSqsConfigurationOutput { return v.Configuration }).(SourceAmazonSqsConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceAmazonSqsOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceAmazonSqsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceAmazonSqsOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceAmazonSqsOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceAmazonSqsOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceAmazonSqsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceAmazonSqs) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceAmazonSqsArrayOutput struct{ *pulumi.OutputState }

func (SourceAmazonSqsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceAmazonSqs)(nil)).Elem()
}

func (o SourceAmazonSqsArrayOutput) ToSourceAmazonSqsArrayOutput() SourceAmazonSqsArrayOutput {
	return o
}

func (o SourceAmazonSqsArrayOutput) ToSourceAmazonSqsArrayOutputWithContext(ctx context.Context) SourceAmazonSqsArrayOutput {
	return o
}

func (o SourceAmazonSqsArrayOutput) Index(i pulumi.IntInput) SourceAmazonSqsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceAmazonSqs {
		return vs[0].([]*SourceAmazonSqs)[vs[1].(int)]
	}).(SourceAmazonSqsOutput)
}

type SourceAmazonSqsMapOutput struct{ *pulumi.OutputState }

func (SourceAmazonSqsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceAmazonSqs)(nil)).Elem()
}

func (o SourceAmazonSqsMapOutput) ToSourceAmazonSqsMapOutput() SourceAmazonSqsMapOutput {
	return o
}

func (o SourceAmazonSqsMapOutput) ToSourceAmazonSqsMapOutputWithContext(ctx context.Context) SourceAmazonSqsMapOutput {
	return o
}

func (o SourceAmazonSqsMapOutput) MapIndex(k pulumi.StringInput) SourceAmazonSqsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceAmazonSqs {
		return vs[0].(map[string]*SourceAmazonSqs)[vs[1].(string)]
	}).(SourceAmazonSqsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonSqsInput)(nil)).Elem(), &SourceAmazonSqs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonSqsArrayInput)(nil)).Elem(), SourceAmazonSqsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceAmazonSqsMapInput)(nil)).Elem(), SourceAmazonSqsMap{})
	pulumi.RegisterOutputType(SourceAmazonSqsOutput{})
	pulumi.RegisterOutputType(SourceAmazonSqsArrayOutput{})
	pulumi.RegisterOutputType(SourceAmazonSqsMapOutput{})
}
