// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package dlmstub

import (
	"github.com/aws/aws-sdk-go/service/dlm"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error)
	CreateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) *DLMCreateLifecyclePolicyFuture

	DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyAsync(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) *DLMDeleteLifecyclePolicyFuture

	GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error)
	GetLifecyclePoliciesAsync(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) *DLMGetLifecyclePoliciesFuture

	GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyAsync(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) *DLMGetLifecyclePolicyFuture

	ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *dlm.ListTagsForResourceInput) *DLMListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dlm.TagResourceInput) *DLMTagResourceFuture

	UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dlm.UntagResourceInput) *DLMUntagResourceFuture

	UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error)
	UpdateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) *DLMUpdateLifecyclePolicyFuture
}

func NewClient() Client {
	return &stub{}
}