// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package batchstub

import (
	"github.com/aws/aws-sdk-go/service/batch"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CancelJob(ctx workflow.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *batch.CancelJobInput) *BatchCancelJobFuture

	CreateComputeEnvironment(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error)
	CreateComputeEnvironmentAsync(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) *BatchCreateComputeEnvironmentFuture

	CreateJobQueue(ctx workflow.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error)
	CreateJobQueueAsync(ctx workflow.Context, input *batch.CreateJobQueueInput) *BatchCreateJobQueueFuture

	DeleteComputeEnvironment(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error)
	DeleteComputeEnvironmentAsync(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) *BatchDeleteComputeEnvironmentFuture

	DeleteJobQueue(ctx workflow.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error)
	DeleteJobQueueAsync(ctx workflow.Context, input *batch.DeleteJobQueueInput) *BatchDeleteJobQueueFuture

	DeregisterJobDefinition(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error)
	DeregisterJobDefinitionAsync(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) *BatchDeregisterJobDefinitionFuture

	DescribeComputeEnvironments(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error)
	DescribeComputeEnvironmentsAsync(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) *BatchDescribeComputeEnvironmentsFuture

	DescribeJobDefinitions(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error)
	DescribeJobDefinitionsAsync(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) *BatchDescribeJobDefinitionsFuture

	DescribeJobQueues(ctx workflow.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error)
	DescribeJobQueuesAsync(ctx workflow.Context, input *batch.DescribeJobQueuesInput) *BatchDescribeJobQueuesFuture

	DescribeJobs(ctx workflow.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error)
	DescribeJobsAsync(ctx workflow.Context, input *batch.DescribeJobsInput) *BatchDescribeJobsFuture

	ListJobs(ctx workflow.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *batch.ListJobsInput) *BatchListJobsFuture

	ListTagsForResource(ctx workflow.Context, input *batch.ListTagsForResourceInput) (*batch.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *batch.ListTagsForResourceInput) *BatchListTagsForResourceFuture

	RegisterJobDefinition(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error)
	RegisterJobDefinitionAsync(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) *BatchRegisterJobDefinitionFuture

	SubmitJob(ctx workflow.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error)
	SubmitJobAsync(ctx workflow.Context, input *batch.SubmitJobInput) *BatchSubmitJobFuture

	TagResource(ctx workflow.Context, input *batch.TagResourceInput) (*batch.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *batch.TagResourceInput) *BatchTagResourceFuture

	TerminateJob(ctx workflow.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error)
	TerminateJobAsync(ctx workflow.Context, input *batch.TerminateJobInput) *BatchTerminateJobFuture

	UntagResource(ctx workflow.Context, input *batch.UntagResourceInput) (*batch.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *batch.UntagResourceInput) *BatchUntagResourceFuture

	UpdateComputeEnvironment(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error)
	UpdateComputeEnvironmentAsync(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) *BatchUpdateComputeEnvironmentFuture

	UpdateJobQueue(ctx workflow.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error)
	UpdateJobQueueAsync(ctx workflow.Context, input *batch.UpdateJobQueueInput) *BatchUpdateJobQueueFuture
}

func NewClient() Client {
	return &stub{}
}

