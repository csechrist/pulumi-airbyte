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

// SourceMicrosoftSharepoint Resource
type SourceMicrosoftSharepoint struct {
	pulumi.CustomResourceState

	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration SourceMicrosoftSharepointConfigurationOutput `pulumi:"configuration"`
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

// NewSourceMicrosoftSharepoint registers a new resource with the given unique name, arguments, and options.
func NewSourceMicrosoftSharepoint(ctx *pulumi.Context,
	name string, args *SourceMicrosoftSharepointArgs, opts ...pulumi.ResourceOption) (*SourceMicrosoftSharepoint, error) {
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
	var resource SourceMicrosoftSharepoint
	err := ctx.RegisterResource("airbyte:index/sourceMicrosoftSharepoint:SourceMicrosoftSharepoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceMicrosoftSharepoint gets an existing SourceMicrosoftSharepoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceMicrosoftSharepoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceMicrosoftSharepointState, opts ...pulumi.ResourceOption) (*SourceMicrosoftSharepoint, error) {
	var resource SourceMicrosoftSharepoint
	err := ctx.ReadResource("airbyte:index/sourceMicrosoftSharepoint:SourceMicrosoftSharepoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceMicrosoftSharepoint resources.
type sourceMicrosoftSharepointState struct {
	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration *SourceMicrosoftSharepointConfiguration `pulumi:"configuration"`
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

type SourceMicrosoftSharepointState struct {
	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration SourceMicrosoftSharepointConfigurationPtrInput
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

func (SourceMicrosoftSharepointState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftSharepointState)(nil)).Elem()
}

type sourceMicrosoftSharepointArgs struct {
	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration SourceMicrosoftSharepointConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourceMicrosoftSharepoint resource.
type SourceMicrosoftSharepointArgs struct {
	// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
	// This class combines the authentication details with additional configuration for the SharePoint API.
	Configuration SourceMicrosoftSharepointConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourceMicrosoftSharepointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceMicrosoftSharepointArgs)(nil)).Elem()
}

type SourceMicrosoftSharepointInput interface {
	pulumi.Input

	ToSourceMicrosoftSharepointOutput() SourceMicrosoftSharepointOutput
	ToSourceMicrosoftSharepointOutputWithContext(ctx context.Context) SourceMicrosoftSharepointOutput
}

func (*SourceMicrosoftSharepoint) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftSharepoint)(nil)).Elem()
}

func (i *SourceMicrosoftSharepoint) ToSourceMicrosoftSharepointOutput() SourceMicrosoftSharepointOutput {
	return i.ToSourceMicrosoftSharepointOutputWithContext(context.Background())
}

func (i *SourceMicrosoftSharepoint) ToSourceMicrosoftSharepointOutputWithContext(ctx context.Context) SourceMicrosoftSharepointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftSharepointOutput)
}

// SourceMicrosoftSharepointArrayInput is an input type that accepts SourceMicrosoftSharepointArray and SourceMicrosoftSharepointArrayOutput values.
// You can construct a concrete instance of `SourceMicrosoftSharepointArrayInput` via:
//
//	SourceMicrosoftSharepointArray{ SourceMicrosoftSharepointArgs{...} }
type SourceMicrosoftSharepointArrayInput interface {
	pulumi.Input

	ToSourceMicrosoftSharepointArrayOutput() SourceMicrosoftSharepointArrayOutput
	ToSourceMicrosoftSharepointArrayOutputWithContext(context.Context) SourceMicrosoftSharepointArrayOutput
}

type SourceMicrosoftSharepointArray []SourceMicrosoftSharepointInput

func (SourceMicrosoftSharepointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMicrosoftSharepoint)(nil)).Elem()
}

func (i SourceMicrosoftSharepointArray) ToSourceMicrosoftSharepointArrayOutput() SourceMicrosoftSharepointArrayOutput {
	return i.ToSourceMicrosoftSharepointArrayOutputWithContext(context.Background())
}

func (i SourceMicrosoftSharepointArray) ToSourceMicrosoftSharepointArrayOutputWithContext(ctx context.Context) SourceMicrosoftSharepointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftSharepointArrayOutput)
}

// SourceMicrosoftSharepointMapInput is an input type that accepts SourceMicrosoftSharepointMap and SourceMicrosoftSharepointMapOutput values.
// You can construct a concrete instance of `SourceMicrosoftSharepointMapInput` via:
//
//	SourceMicrosoftSharepointMap{ "key": SourceMicrosoftSharepointArgs{...} }
type SourceMicrosoftSharepointMapInput interface {
	pulumi.Input

	ToSourceMicrosoftSharepointMapOutput() SourceMicrosoftSharepointMapOutput
	ToSourceMicrosoftSharepointMapOutputWithContext(context.Context) SourceMicrosoftSharepointMapOutput
}

type SourceMicrosoftSharepointMap map[string]SourceMicrosoftSharepointInput

func (SourceMicrosoftSharepointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMicrosoftSharepoint)(nil)).Elem()
}

func (i SourceMicrosoftSharepointMap) ToSourceMicrosoftSharepointMapOutput() SourceMicrosoftSharepointMapOutput {
	return i.ToSourceMicrosoftSharepointMapOutputWithContext(context.Background())
}

func (i SourceMicrosoftSharepointMap) ToSourceMicrosoftSharepointMapOutputWithContext(ctx context.Context) SourceMicrosoftSharepointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceMicrosoftSharepointMapOutput)
}

type SourceMicrosoftSharepointOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftSharepointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceMicrosoftSharepoint)(nil)).Elem()
}

func (o SourceMicrosoftSharepointOutput) ToSourceMicrosoftSharepointOutput() SourceMicrosoftSharepointOutput {
	return o
}

func (o SourceMicrosoftSharepointOutput) ToSourceMicrosoftSharepointOutputWithContext(ctx context.Context) SourceMicrosoftSharepointOutput {
	return o
}

// SourceMicrosoftSharePointSpec class for Microsoft SharePoint Source Specification.
// This class combines the authentication details with additional configuration for the SharePoint API.
func (o SourceMicrosoftSharepointOutput) Configuration() SourceMicrosoftSharepointConfigurationOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) SourceMicrosoftSharepointConfigurationOutput {
		return v.Configuration
	}).(SourceMicrosoftSharepointConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourceMicrosoftSharepointOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourceMicrosoftSharepointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourceMicrosoftSharepointOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourceMicrosoftSharepointOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourceMicrosoftSharepointOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourceMicrosoftSharepointOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceMicrosoftSharepoint) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourceMicrosoftSharepointArrayOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftSharepointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourceMicrosoftSharepoint)(nil)).Elem()
}

func (o SourceMicrosoftSharepointArrayOutput) ToSourceMicrosoftSharepointArrayOutput() SourceMicrosoftSharepointArrayOutput {
	return o
}

func (o SourceMicrosoftSharepointArrayOutput) ToSourceMicrosoftSharepointArrayOutputWithContext(ctx context.Context) SourceMicrosoftSharepointArrayOutput {
	return o
}

func (o SourceMicrosoftSharepointArrayOutput) Index(i pulumi.IntInput) SourceMicrosoftSharepointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourceMicrosoftSharepoint {
		return vs[0].([]*SourceMicrosoftSharepoint)[vs[1].(int)]
	}).(SourceMicrosoftSharepointOutput)
}

type SourceMicrosoftSharepointMapOutput struct{ *pulumi.OutputState }

func (SourceMicrosoftSharepointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourceMicrosoftSharepoint)(nil)).Elem()
}

func (o SourceMicrosoftSharepointMapOutput) ToSourceMicrosoftSharepointMapOutput() SourceMicrosoftSharepointMapOutput {
	return o
}

func (o SourceMicrosoftSharepointMapOutput) ToSourceMicrosoftSharepointMapOutputWithContext(ctx context.Context) SourceMicrosoftSharepointMapOutput {
	return o
}

func (o SourceMicrosoftSharepointMapOutput) MapIndex(k pulumi.StringInput) SourceMicrosoftSharepointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourceMicrosoftSharepoint {
		return vs[0].(map[string]*SourceMicrosoftSharepoint)[vs[1].(string)]
	}).(SourceMicrosoftSharepointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftSharepointInput)(nil)).Elem(), &SourceMicrosoftSharepoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftSharepointArrayInput)(nil)).Elem(), SourceMicrosoftSharepointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceMicrosoftSharepointMapInput)(nil)).Elem(), SourceMicrosoftSharepointMap{})
	pulumi.RegisterOutputType(SourceMicrosoftSharepointOutput{})
	pulumi.RegisterOutputType(SourceMicrosoftSharepointArrayOutput{})
	pulumi.RegisterOutputType(SourceMicrosoftSharepointMapOutput{})
}