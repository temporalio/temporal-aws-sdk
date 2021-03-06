// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package globalacceleratorstub

import (
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error)
	AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *AdvertiseByoipCidrFuture

	CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *CreateAcceleratorFuture

	CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *CreateEndpointGroupFuture

	CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *CreateListenerFuture

	DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *DeleteAcceleratorFuture

	DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *DeleteEndpointGroupFuture

	DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *DeleteListenerFuture

	DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error)
	DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *DeprovisionByoipCidrFuture

	DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *DescribeAcceleratorFuture

	DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *DescribeAcceleratorAttributesFuture

	DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *DescribeEndpointGroupFuture

	DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *DescribeListenerFuture

	ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *ListAcceleratorsFuture

	ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error)
	ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *ListByoipCidrsFuture

	ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *ListEndpointGroupsFuture

	ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error)
	ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *ListListenersFuture

	ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *ListTagsForResourceFuture

	ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error)
	ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *ProvisionByoipCidrFuture

	TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *UntagResourceFuture

	UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *UpdateAcceleratorFuture

	UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *UpdateAcceleratorAttributesFuture

	UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *UpdateEndpointGroupFuture

	UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *UpdateListenerFuture

	WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error)
	WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *WithdrawByoipCidrFuture
}

func NewClient() Client {
	return &stub{}
}
