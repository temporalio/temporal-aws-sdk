// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package resourcegroupstaggingapistub

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ResourceGroupsTaggingAPIDescribeReportCreationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIDescribeReportCreationFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	var output resourcegroupstaggingapi.DescribeReportCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIGetComplianceSummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIGetComplianceSummaryFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	var output resourcegroupstaggingapi.GetComplianceSummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIGetResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIGetResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	var output resourcegroupstaggingapi.GetResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIGetTagKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIGetTagKeysFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	var output resourcegroupstaggingapi.GetTagKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIGetTagValuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIGetTagValuesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	var output resourcegroupstaggingapi.GetTagValuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIStartReportCreationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIStartReportCreationFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	var output resourcegroupstaggingapi.StartReportCreationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPITagResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPITagResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	var output resourcegroupstaggingapi.TagResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTaggingAPIUntagResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ResourceGroupsTaggingAPIUntagResourcesFuture) Get(ctx workflow.Context) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	var output resourcegroupstaggingapi.UntagResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	var output resourcegroupstaggingapi.DescribeReportCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.DescribeReportCreation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *ResourceGroupsTaggingAPIDescribeReportCreationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.DescribeReportCreation", input)
	return &ResourceGroupsTaggingAPIDescribeReportCreationFuture{Future: future}
}

func (a *stub) GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	var output resourcegroupstaggingapi.GetComplianceSummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetComplianceSummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *ResourceGroupsTaggingAPIGetComplianceSummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetComplianceSummary", input)
	return &ResourceGroupsTaggingAPIGetComplianceSummaryFuture{Future: future}
}

func (a *stub) GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	var output resourcegroupstaggingapi.GetResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *ResourceGroupsTaggingAPIGetResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetResources", input)
	return &ResourceGroupsTaggingAPIGetResourcesFuture{Future: future}
}

func (a *stub) GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	var output resourcegroupstaggingapi.GetTagKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetTagKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *ResourceGroupsTaggingAPIGetTagKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetTagKeys", input)
	return &ResourceGroupsTaggingAPIGetTagKeysFuture{Future: future}
}

func (a *stub) GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	var output resourcegroupstaggingapi.GetTagValuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetTagValues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *ResourceGroupsTaggingAPIGetTagValuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.GetTagValues", input)
	return &ResourceGroupsTaggingAPIGetTagValuesFuture{Future: future}
}

func (a *stub) StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	var output resourcegroupstaggingapi.StartReportCreationOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.StartReportCreation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *ResourceGroupsTaggingAPIStartReportCreationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.StartReportCreation", input)
	return &ResourceGroupsTaggingAPIStartReportCreationFuture{Future: future}
}

func (a *stub) TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	var output resourcegroupstaggingapi.TagResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.TagResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *ResourceGroupsTaggingAPITagResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.TagResources", input)
	return &ResourceGroupsTaggingAPITagResourcesFuture{Future: future}
}

func (a *stub) UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	var output resourcegroupstaggingapi.UntagResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.UntagResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *ResourceGroupsTaggingAPIUntagResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroupstaggingapi.UntagResources", input)
	return &ResourceGroupsTaggingAPIUntagResourcesFuture{Future: future}
}
