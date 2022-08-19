// Code generated by pulumi-gen-awsx DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LifecycleTagStatus string

const (
	// Evaluate rule against all images
	LifecycleTagStatusAny = LifecycleTagStatus("any")
	// Only evaluate rule against untagged images
	LifecycleTagStatusUntagged = LifecycleTagStatus("untagged")
	// Only evaluated rule against images with specified prefixes
	LifecycleTagStatusTagged = LifecycleTagStatus("tagged")
)

func (LifecycleTagStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecycleTagStatus)(nil)).Elem()
}

func (e LifecycleTagStatus) ToLifecycleTagStatusOutput() LifecycleTagStatusOutput {
	return pulumi.ToOutput(e).(LifecycleTagStatusOutput)
}

func (e LifecycleTagStatus) ToLifecycleTagStatusOutputWithContext(ctx context.Context) LifecycleTagStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LifecycleTagStatusOutput)
}

func (e LifecycleTagStatus) ToLifecycleTagStatusPtrOutput() LifecycleTagStatusPtrOutput {
	return e.ToLifecycleTagStatusPtrOutputWithContext(context.Background())
}

func (e LifecycleTagStatus) ToLifecycleTagStatusPtrOutputWithContext(ctx context.Context) LifecycleTagStatusPtrOutput {
	return LifecycleTagStatus(e).ToLifecycleTagStatusOutputWithContext(ctx).ToLifecycleTagStatusPtrOutputWithContext(ctx)
}

func (e LifecycleTagStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LifecycleTagStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LifecycleTagStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LifecycleTagStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LifecycleTagStatusOutput struct{ *pulumi.OutputState }

func (LifecycleTagStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecycleTagStatus)(nil)).Elem()
}

func (o LifecycleTagStatusOutput) ToLifecycleTagStatusOutput() LifecycleTagStatusOutput {
	return o
}

func (o LifecycleTagStatusOutput) ToLifecycleTagStatusOutputWithContext(ctx context.Context) LifecycleTagStatusOutput {
	return o
}

func (o LifecycleTagStatusOutput) ToLifecycleTagStatusPtrOutput() LifecycleTagStatusPtrOutput {
	return o.ToLifecycleTagStatusPtrOutputWithContext(context.Background())
}

func (o LifecycleTagStatusOutput) ToLifecycleTagStatusPtrOutputWithContext(ctx context.Context) LifecycleTagStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LifecycleTagStatus) *LifecycleTagStatus {
		return &v
	}).(LifecycleTagStatusPtrOutput)
}

func (o LifecycleTagStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LifecycleTagStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LifecycleTagStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LifecycleTagStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LifecycleTagStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LifecycleTagStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LifecycleTagStatusPtrOutput struct{ *pulumi.OutputState }

func (LifecycleTagStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecycleTagStatus)(nil)).Elem()
}

func (o LifecycleTagStatusPtrOutput) ToLifecycleTagStatusPtrOutput() LifecycleTagStatusPtrOutput {
	return o
}

func (o LifecycleTagStatusPtrOutput) ToLifecycleTagStatusPtrOutputWithContext(ctx context.Context) LifecycleTagStatusPtrOutput {
	return o
}

func (o LifecycleTagStatusPtrOutput) Elem() LifecycleTagStatusOutput {
	return o.ApplyT(func(v *LifecycleTagStatus) LifecycleTagStatus {
		if v != nil {
			return *v
		}
		var ret LifecycleTagStatus
		return ret
	}).(LifecycleTagStatusOutput)
}

func (o LifecycleTagStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LifecycleTagStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LifecycleTagStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// LifecycleTagStatusInput is an input type that accepts LifecycleTagStatusArgs and LifecycleTagStatusOutput values.
// You can construct a concrete instance of `LifecycleTagStatusInput` via:
//
//	LifecycleTagStatusArgs{...}
type LifecycleTagStatusInput interface {
	pulumi.Input

	ToLifecycleTagStatusOutput() LifecycleTagStatusOutput
	ToLifecycleTagStatusOutputWithContext(context.Context) LifecycleTagStatusOutput
}

var lifecycleTagStatusPtrType = reflect.TypeOf((**LifecycleTagStatus)(nil)).Elem()

type LifecycleTagStatusPtrInput interface {
	pulumi.Input

	ToLifecycleTagStatusPtrOutput() LifecycleTagStatusPtrOutput
	ToLifecycleTagStatusPtrOutputWithContext(context.Context) LifecycleTagStatusPtrOutput
}

type lifecycleTagStatusPtr string

func LifecycleTagStatusPtr(v string) LifecycleTagStatusPtrInput {
	return (*lifecycleTagStatusPtr)(&v)
}

func (*lifecycleTagStatusPtr) ElementType() reflect.Type {
	return lifecycleTagStatusPtrType
}

func (in *lifecycleTagStatusPtr) ToLifecycleTagStatusPtrOutput() LifecycleTagStatusPtrOutput {
	return pulumi.ToOutput(in).(LifecycleTagStatusPtrOutput)
}

func (in *lifecycleTagStatusPtr) ToLifecycleTagStatusPtrOutputWithContext(ctx context.Context) LifecycleTagStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LifecycleTagStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LifecycleTagStatusInput)(nil)).Elem(), LifecycleTagStatus("any"))
	pulumi.RegisterInputType(reflect.TypeOf((*LifecycleTagStatusPtrInput)(nil)).Elem(), LifecycleTagStatus("any"))
	pulumi.RegisterOutputType(LifecycleTagStatusOutput{})
	pulumi.RegisterOutputType(LifecycleTagStatusPtrOutput{})
}
