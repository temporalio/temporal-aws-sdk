// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package servicediscoverystub

import (
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateHttpNamespace(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error)
	CreateHttpNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreateHttpNamespaceInput) *CreateHttpNamespaceFuture

	CreatePrivateDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error)
	CreatePrivateDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) *CreatePrivateDnsNamespaceFuture

	CreatePublicDnsNamespace(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error)
	CreatePublicDnsNamespaceAsync(ctx workflow.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) *CreatePublicDnsNamespaceFuture

	CreateService(ctx workflow.Context, input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error)
	CreateServiceAsync(ctx workflow.Context, input *servicediscovery.CreateServiceInput) *CreateServiceFuture

	DeleteNamespace(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error)
	DeleteNamespaceAsync(ctx workflow.Context, input *servicediscovery.DeleteNamespaceInput) *DeleteNamespaceFuture

	DeleteService(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error)
	DeleteServiceAsync(ctx workflow.Context, input *servicediscovery.DeleteServiceInput) *DeleteServiceFuture

	DeregisterInstance(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error)
	DeregisterInstanceAsync(ctx workflow.Context, input *servicediscovery.DeregisterInstanceInput) *DeregisterInstanceFuture

	DiscoverInstances(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error)
	DiscoverInstancesAsync(ctx workflow.Context, input *servicediscovery.DiscoverInstancesInput) *DiscoverInstancesFuture

	GetInstance(ctx workflow.Context, input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error)
	GetInstanceAsync(ctx workflow.Context, input *servicediscovery.GetInstanceInput) *GetInstanceFuture

	GetInstancesHealthStatus(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error)
	GetInstancesHealthStatusAsync(ctx workflow.Context, input *servicediscovery.GetInstancesHealthStatusInput) *GetInstancesHealthStatusFuture

	GetNamespace(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error)
	GetNamespaceAsync(ctx workflow.Context, input *servicediscovery.GetNamespaceInput) *GetNamespaceFuture

	GetOperation(ctx workflow.Context, input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error)
	GetOperationAsync(ctx workflow.Context, input *servicediscovery.GetOperationInput) *GetOperationFuture

	GetService(ctx workflow.Context, input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error)
	GetServiceAsync(ctx workflow.Context, input *servicediscovery.GetServiceInput) *GetServiceFuture

	ListInstances(ctx workflow.Context, input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error)
	ListInstancesAsync(ctx workflow.Context, input *servicediscovery.ListInstancesInput) *ListInstancesFuture

	ListNamespaces(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error)
	ListNamespacesAsync(ctx workflow.Context, input *servicediscovery.ListNamespacesInput) *ListNamespacesFuture

	ListOperations(ctx workflow.Context, input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error)
	ListOperationsAsync(ctx workflow.Context, input *servicediscovery.ListOperationsInput) *ListOperationsFuture

	ListServices(ctx workflow.Context, input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error)
	ListServicesAsync(ctx workflow.Context, input *servicediscovery.ListServicesInput) *ListServicesFuture

	ListTagsForResource(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *servicediscovery.ListTagsForResourceInput) *ListTagsForResourceFuture

	RegisterInstance(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error)
	RegisterInstanceAsync(ctx workflow.Context, input *servicediscovery.RegisterInstanceInput) *RegisterInstanceFuture

	TagResource(ctx workflow.Context, input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *servicediscovery.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *servicediscovery.UntagResourceInput) *UntagResourceFuture

	UpdateInstanceCustomHealthStatus(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error)
	UpdateInstanceCustomHealthStatusAsync(ctx workflow.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) *UpdateInstanceCustomHealthStatusFuture

	UpdateService(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error)
	UpdateServiceAsync(ctx workflow.Context, input *servicediscovery.UpdateServiceInput) *UpdateServiceFuture
}

func NewClient() Client {
	return &stub{}
}
