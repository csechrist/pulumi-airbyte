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

// SourceTheGuardianAPI Resource
type SourceTheGuardianApi struct {
	pulumi.CustomResourceState

	Configuration SourceTheGuardianApiConfigurationOutput `pulumi:"configuration"`
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

// NewSourceTheGuardianApi registers a new resource with the given unique name, arguments, and options.
func NewSourceTheGuardianApi(ctx *pulumi.Context,
	name string, args *SourceTheGuardianApiArgs, opts ...pulumi.ResourceOption) (*SourceTheGuardianApi, error) {
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
	var resource SourceTheGuardianApi
	err := ctx.RegisterResource("airbyte:index/sourceTheGuardianApi:SourceTheGuardianApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceTheGuardianApi gets an existing SourceTheGuardianApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceTheGuardianApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceTheGuardianApiState, opts ...pulumi.ResourceOption) (*SourceTheGuardianApi, error) {
	var resource SourceTheGuardianApi
	err := ctx.ReadResource("airbyte:index/sourceTheGuardianApi:SourceTheGuardianApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceTheGuardianApi resources.
type sourceTheGuardianApiState struct {
	Configuration *SourceTheGuardianApiConfiguration `pulumi:"configuration"`
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

type SourceTheGuardianApiState struct {
	Configuration SourceTheGuardianApiConfigurationPtrInput
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

func (SourceTheGuardianApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTheGuardianApiState)(nil)).Elem()
}

type sourceTheGuardianApiArgs struct {
	Configuration SourceTheGuardianApiConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceTheGuardianApi resource.
type SourceTheGuardianApiArgs struct {
	Configuration SourceTheGuardianApiConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceTheGuardianApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceTheGuardianApiArgs)(nil)).Elem()
}

type SourceTheGuardianApiInput interface {
	pulumi.Input

	ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput
	ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput
}

func (*SourceTheGuardianApi) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTheGuardianApi)(nil)).Elem()
}

func (i *SourceTheGuardianApi) ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput {
	return i.ToSourceTheGuardianApiOutputWithContext(context.Background())
}

func (i *SourceTheGuardianApi) ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTheGuardianApiOutput)
}

// SourceTheGuardianApiArrayInput is an input type that accepts SourceTheGuardianApiArray and SourceTheGuardianApiArrayOutput values.
// You can construct a concrete instance of `SourceTheGuardianApiArrayInput` via:
//
//	SourceTheGuardianApiArray{ SourceTheGuardianApiArgs{...} }
type SourceTheGuardianApiArrayInput interface {
	pulumi.Input

	ToSourceTheGuardianApiArrayOutput() SourceTheGuardianApiArrayOutput
	ToSourceTheGuardianApiArrayOutputWithContext(context.Context) SourceTheGuardianApiArrayOutput
}

type SourceTheGuardianApiArray []SourceTheGuardianApiInput

func (SourceTheGuardianApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTheGuardianApi)(nil)).Elem()
}

func (i SourceTheGuardianApiArray) ToSourceTheGuardianApiArrayOutput() SourceTheGuardianApiArrayOutput {
	return i.ToSourceTheGuardianApiArrayOutputWithContext(context.Background())
}

func (i SourceTheGuardianApiArray) ToSourceTheGuardianApiArrayOutputWithContext(ctx context.Context) SourceTheGuardianApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTheGuardianApiArrayOutput)
}

// SourceTheGuardianApiMapInput is an input type that accepts SourceTheGuardianApiMap and SourceTheGuardianApiMapOutput values.
// You can construct a concrete instance of `SourceTheGuardianApiMapInput` via:
//
//	SourceTheGuardianApiMap{ "key": SourceTheGuardianApiArgs{...} }
type SourceTheGuardianApiMapInput interface {
	pulumi.Input

	ToSourceTheGuardianApiMapOutput() SourceTheGuardianApiMapOutput
	ToSourceTheGuardianApiMapOutputWithContext(context.Context) SourceTheGuardianApiMapOutput
}

type SourceTheGuardianApiMap map[string]SourceTheGuardianApiInput

func (SourceTheGuardianApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTheGuardianApi)(nil)).Elem()
}

func (i SourceTheGuardianApiMap) ToSourceTheGuardianApiMapOutput() SourceTheGuardianApiMapOutput {
	return i.ToSourceTheGuardianApiMapOutputWithContext(context.Background())
}

func (i SourceTheGuardianApiMap) ToSourceTheGuardianApiMapOutputWithContext(ctx context.Context) SourceTheGuardianApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceTheGuardianApiMapOutput)
}

type SourceTheGuardianApiOutput struct{ *pulumi.OutputState }

func (SourceTheGuardianApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTheGuardianApi)(nil)).Elem()
}

func (o SourceTheGuardianApiOutput) ToSourceTheGuardianApiOutput() SourceTheGuardianApiOutput {
	return o
}

func (o SourceTheGuardianApiOutput) ToSourceTheGuardianApiOutputWithContext(ctx context.Context) SourceTheGuardianApiOutput {
	return o
}

func (o SourceTheGuardianApiOutput) Configuration() SourceTheGuardianApiConfigurationOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) SourceTheGuardianApiConfigurationOutput { return v.Configuration }).(SourceTheGuardianApiConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceTheGuardianApiOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceTheGuardianApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceTheGuardianApiOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceTheGuardianApiOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceTheGuardianApiOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceTheGuardianApiOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceTheGuardianApi) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceTheGuardianApiArrayOutput struct{ *pulumi.OutputState }

func (SourceTheGuardianApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceTheGuardianApi)(nil)).Elem()
}

func (o SourceTheGuardianApiArrayOutput) ToSourceTheGuardianApiArrayOutput() SourceTheGuardianApiArrayOutput {
	return o
}

func (o SourceTheGuardianApiArrayOutput) ToSourceTheGuardianApiArrayOutputWithContext(ctx context.Context) SourceTheGuardianApiArrayOutput {
	return o
}

func (o SourceTheGuardianApiArrayOutput) Index(i pulumi.IntInput) SourceTheGuardianApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceTheGuardianApi {
		return vs[0].([]*SourceTheGuardianApi)[vs[1].(int)]
	}).(SourceTheGuardianApiOutput)
}

type SourceTheGuardianApiMapOutput struct{ *pulumi.OutputState }

func (SourceTheGuardianApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceTheGuardianApi)(nil)).Elem()
}

func (o SourceTheGuardianApiMapOutput) ToSourceTheGuardianApiMapOutput() SourceTheGuardianApiMapOutput {
	return o
}

func (o SourceTheGuardianApiMapOutput) ToSourceTheGuardianApiMapOutputWithContext(ctx context.Context) SourceTheGuardianApiMapOutput {
	return o
}

func (o SourceTheGuardianApiMapOutput) MapIndex(k pulumi.StringInput) SourceTheGuardianApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceTheGuardianApi {
		return vs[0].(map[string]*SourceTheGuardianApi)[vs[1].(string)]
	}).(SourceTheGuardianApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTheGuardianApiInput)(nil)).Elem(), &SourceTheGuardianApi{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTheGuardianApiArrayInput)(nil)).Elem(), SourceTheGuardianApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTheGuardianApiMapInput)(nil)).Elem(), SourceTheGuardianApiMap{})
	pulumi.RegisterOutputType(SourceTheGuardianApiOutput{})
	pulumi.RegisterOutputType(SourceTheGuardianApiArrayOutput{})
	pulumi.RegisterOutputType(SourceTheGuardianApiMapOutput{})
}
