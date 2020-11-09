// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediastorestub

import (
	"github.com/aws/aws-sdk-go/service/mediastore"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateContainer(ctx workflow.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error)
	CreateContainerAsync(ctx workflow.Context, input *mediastore.CreateContainerInput) *CreateContainerFuture

	DeleteContainer(ctx workflow.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error)
	DeleteContainerAsync(ctx workflow.Context, input *mediastore.DeleteContainerInput) *DeleteContainerFuture

	DeleteContainerPolicy(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error)
	DeleteContainerPolicyAsync(ctx workflow.Context, input *mediastore.DeleteContainerPolicyInput) *DeleteContainerPolicyFuture

	DeleteCorsPolicy(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error)
	DeleteCorsPolicyAsync(ctx workflow.Context, input *mediastore.DeleteCorsPolicyInput) *DeleteCorsPolicyFuture

	DeleteLifecyclePolicy(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.DeleteLifecyclePolicyInput) *DeleteLifecyclePolicyFuture

	DeleteMetricPolicy(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error)
	DeleteMetricPolicyAsync(ctx workflow.Context, input *mediastore.DeleteMetricPolicyInput) *DeleteMetricPolicyFuture

	DescribeContainer(ctx workflow.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error)
	DescribeContainerAsync(ctx workflow.Context, input *mediastore.DescribeContainerInput) *DescribeContainerFuture

	GetContainerPolicy(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error)
	GetContainerPolicyAsync(ctx workflow.Context, input *mediastore.GetContainerPolicyInput) *GetContainerPolicyFuture

	GetCorsPolicy(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error)
	GetCorsPolicyAsync(ctx workflow.Context, input *mediastore.GetCorsPolicyInput) *GetCorsPolicyFuture

	GetLifecyclePolicy(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.GetLifecyclePolicyInput) *GetLifecyclePolicyFuture

	GetMetricPolicy(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error)
	GetMetricPolicyAsync(ctx workflow.Context, input *mediastore.GetMetricPolicyInput) *GetMetricPolicyFuture

	ListContainers(ctx workflow.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error)
	ListContainersAsync(ctx workflow.Context, input *mediastore.ListContainersInput) *ListContainersFuture

	ListTagsForResource(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediastore.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutContainerPolicy(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error)
	PutContainerPolicyAsync(ctx workflow.Context, input *mediastore.PutContainerPolicyInput) *PutContainerPolicyFuture

	PutCorsPolicy(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error)
	PutCorsPolicyAsync(ctx workflow.Context, input *mediastore.PutCorsPolicyInput) *PutCorsPolicyFuture

	PutLifecyclePolicy(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error)
	PutLifecyclePolicyAsync(ctx workflow.Context, input *mediastore.PutLifecyclePolicyInput) *PutLifecyclePolicyFuture

	PutMetricPolicy(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error)
	PutMetricPolicyAsync(ctx workflow.Context, input *mediastore.PutMetricPolicyInput) *PutMetricPolicyFuture

	StartAccessLogging(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error)
	StartAccessLoggingAsync(ctx workflow.Context, input *mediastore.StartAccessLoggingInput) *StartAccessLoggingFuture

	StopAccessLogging(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error)
	StopAccessLoggingAsync(ctx workflow.Context, input *mediastore.StopAccessLoggingInput) *StopAccessLoggingFuture

	TagResource(ctx workflow.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediastore.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediastore.UntagResourceInput) *UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
