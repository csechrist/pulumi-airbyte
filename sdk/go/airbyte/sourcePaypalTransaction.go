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

// SourcePaypalTransaction Resource
type SourcePaypalTransaction struct {
	pulumi.CustomResourceState

	Configuration SourcePaypalTransactionConfigurationOutput `pulumi:"configuration"`
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

// NewSourcePaypalTransaction registers a new resource with the given unique name, arguments, and options.
func NewSourcePaypalTransaction(ctx *pulumi.Context,
	name string, args *SourcePaypalTransactionArgs, opts ...pulumi.ResourceOption) (*SourcePaypalTransaction, error) {
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
	var resource SourcePaypalTransaction
	err := ctx.RegisterResource("airbyte:index/sourcePaypalTransaction:SourcePaypalTransaction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourcePaypalTransaction gets an existing SourcePaypalTransaction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourcePaypalTransaction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourcePaypalTransactionState, opts ...pulumi.ResourceOption) (*SourcePaypalTransaction, error) {
	var resource SourcePaypalTransaction
	err := ctx.ReadResource("airbyte:index/sourcePaypalTransaction:SourcePaypalTransaction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourcePaypalTransaction resources.
type sourcePaypalTransactionState struct {
	Configuration *SourcePaypalTransactionConfiguration `pulumi:"configuration"`
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

type SourcePaypalTransactionState struct {
	Configuration SourcePaypalTransactionConfigurationPtrInput
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

func (SourcePaypalTransactionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePaypalTransactionState)(nil)).Elem()
}

type sourcePaypalTransactionArgs struct {
	Configuration SourcePaypalTransactionConfiguration `pulumi:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId *string `pulumi:"definitionId"`
	// Name of the source e.g. dev-mysql-instance.
	Name *string `pulumi:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    *string `pulumi:"secretId"`
	WorkspaceId string  `pulumi:"workspaceId"`
}

// The set of arguments for constructing a SourcePaypalTransaction resource.
type SourcePaypalTransactionArgs struct {
	Configuration SourcePaypalTransactionConfigurationInput
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
	DefinitionId pulumi.StringPtrInput
	// Name of the source e.g. dev-mysql-instance.
	Name pulumi.StringPtrInput
	// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
	SecretId    pulumi.StringPtrInput
	WorkspaceId pulumi.StringInput
}

func (SourcePaypalTransactionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourcePaypalTransactionArgs)(nil)).Elem()
}

type SourcePaypalTransactionInput interface {
	pulumi.Input

	ToSourcePaypalTransactionOutput() SourcePaypalTransactionOutput
	ToSourcePaypalTransactionOutputWithContext(ctx context.Context) SourcePaypalTransactionOutput
}

func (*SourcePaypalTransaction) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePaypalTransaction)(nil)).Elem()
}

func (i *SourcePaypalTransaction) ToSourcePaypalTransactionOutput() SourcePaypalTransactionOutput {
	return i.ToSourcePaypalTransactionOutputWithContext(context.Background())
}

func (i *SourcePaypalTransaction) ToSourcePaypalTransactionOutputWithContext(ctx context.Context) SourcePaypalTransactionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePaypalTransactionOutput)
}

// SourcePaypalTransactionArrayInput is an input type that accepts SourcePaypalTransactionArray and SourcePaypalTransactionArrayOutput values.
// You can construct a concrete instance of `SourcePaypalTransactionArrayInput` via:
//
//	SourcePaypalTransactionArray{ SourcePaypalTransactionArgs{...} }
type SourcePaypalTransactionArrayInput interface {
	pulumi.Input

	ToSourcePaypalTransactionArrayOutput() SourcePaypalTransactionArrayOutput
	ToSourcePaypalTransactionArrayOutputWithContext(context.Context) SourcePaypalTransactionArrayOutput
}

type SourcePaypalTransactionArray []SourcePaypalTransactionInput

func (SourcePaypalTransactionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePaypalTransaction)(nil)).Elem()
}

func (i SourcePaypalTransactionArray) ToSourcePaypalTransactionArrayOutput() SourcePaypalTransactionArrayOutput {
	return i.ToSourcePaypalTransactionArrayOutputWithContext(context.Background())
}

func (i SourcePaypalTransactionArray) ToSourcePaypalTransactionArrayOutputWithContext(ctx context.Context) SourcePaypalTransactionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePaypalTransactionArrayOutput)
}

// SourcePaypalTransactionMapInput is an input type that accepts SourcePaypalTransactionMap and SourcePaypalTransactionMapOutput values.
// You can construct a concrete instance of `SourcePaypalTransactionMapInput` via:
//
//	SourcePaypalTransactionMap{ "key": SourcePaypalTransactionArgs{...} }
type SourcePaypalTransactionMapInput interface {
	pulumi.Input

	ToSourcePaypalTransactionMapOutput() SourcePaypalTransactionMapOutput
	ToSourcePaypalTransactionMapOutputWithContext(context.Context) SourcePaypalTransactionMapOutput
}

type SourcePaypalTransactionMap map[string]SourcePaypalTransactionInput

func (SourcePaypalTransactionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePaypalTransaction)(nil)).Elem()
}

func (i SourcePaypalTransactionMap) ToSourcePaypalTransactionMapOutput() SourcePaypalTransactionMapOutput {
	return i.ToSourcePaypalTransactionMapOutputWithContext(context.Background())
}

func (i SourcePaypalTransactionMap) ToSourcePaypalTransactionMapOutputWithContext(ctx context.Context) SourcePaypalTransactionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePaypalTransactionMapOutput)
}

type SourcePaypalTransactionOutput struct{ *pulumi.OutputState }

func (SourcePaypalTransactionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourcePaypalTransaction)(nil)).Elem()
}

func (o SourcePaypalTransactionOutput) ToSourcePaypalTransactionOutput() SourcePaypalTransactionOutput {
	return o
}

func (o SourcePaypalTransactionOutput) ToSourcePaypalTransactionOutputWithContext(ctx context.Context) SourcePaypalTransactionOutput {
	return o
}

func (o SourcePaypalTransactionOutput) Configuration() SourcePaypalTransactionConfigurationOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) SourcePaypalTransactionConfigurationOutput { return v.Configuration }).(SourcePaypalTransactionConfigurationOutput)
}

// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
func (o SourcePaypalTransactionOutput) DefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringPtrOutput { return v.DefinitionId }).(pulumi.StringPtrOutput)
}

// Name of the source e.g. dev-mysql-instance.
func (o SourcePaypalTransactionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
func (o SourcePaypalTransactionOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringPtrOutput { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o SourcePaypalTransactionOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringOutput { return v.SourceId }).(pulumi.StringOutput)
}

func (o SourcePaypalTransactionOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o SourcePaypalTransactionOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SourcePaypalTransaction) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type SourcePaypalTransactionArrayOutput struct{ *pulumi.OutputState }

func (SourcePaypalTransactionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SourcePaypalTransaction)(nil)).Elem()
}

func (o SourcePaypalTransactionArrayOutput) ToSourcePaypalTransactionArrayOutput() SourcePaypalTransactionArrayOutput {
	return o
}

func (o SourcePaypalTransactionArrayOutput) ToSourcePaypalTransactionArrayOutputWithContext(ctx context.Context) SourcePaypalTransactionArrayOutput {
	return o
}

func (o SourcePaypalTransactionArrayOutput) Index(i pulumi.IntInput) SourcePaypalTransactionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SourcePaypalTransaction {
		return vs[0].([]*SourcePaypalTransaction)[vs[1].(int)]
	}).(SourcePaypalTransactionOutput)
}

type SourcePaypalTransactionMapOutput struct{ *pulumi.OutputState }

func (SourcePaypalTransactionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SourcePaypalTransaction)(nil)).Elem()
}

func (o SourcePaypalTransactionMapOutput) ToSourcePaypalTransactionMapOutput() SourcePaypalTransactionMapOutput {
	return o
}

func (o SourcePaypalTransactionMapOutput) ToSourcePaypalTransactionMapOutputWithContext(ctx context.Context) SourcePaypalTransactionMapOutput {
	return o
}

func (o SourcePaypalTransactionMapOutput) MapIndex(k pulumi.StringInput) SourcePaypalTransactionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SourcePaypalTransaction {
		return vs[0].(map[string]*SourcePaypalTransaction)[vs[1].(string)]
	}).(SourcePaypalTransactionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePaypalTransactionInput)(nil)).Elem(), &SourcePaypalTransaction{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePaypalTransactionArrayInput)(nil)).Elem(), SourcePaypalTransactionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourcePaypalTransactionMapInput)(nil)).Elem(), SourcePaypalTransactionMap{})
	pulumi.RegisterOutputType(SourcePaypalTransactionOutput{})
	pulumi.RegisterOutputType(SourcePaypalTransactionArrayOutput{})
	pulumi.RegisterOutputType(SourcePaypalTransactionMapOutput{})
}
