// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package codepipelinestub

import (
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcknowledgeJob(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error)
	AcknowledgeJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) *CodePipelineAcknowledgeJobFuture

	AcknowledgeThirdPartyJob(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)
	AcknowledgeThirdPartyJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) *CodePipelineAcknowledgeThirdPartyJobFuture

	CreateCustomActionType(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error)
	CreateCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) *CodePipelineCreateCustomActionTypeFuture

	CreatePipeline(ctx workflow.Context, input *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error)
	CreatePipelineAsync(ctx workflow.Context, input *codepipeline.CreatePipelineInput) *CodePipelineCreatePipelineFuture

	DeleteCustomActionType(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error)
	DeleteCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) *CodePipelineDeleteCustomActionTypeFuture

	DeletePipeline(ctx workflow.Context, input *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error)
	DeletePipelineAsync(ctx workflow.Context, input *codepipeline.DeletePipelineInput) *CodePipelineDeletePipelineFuture

	DeleteWebhook(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error)
	DeleteWebhookAsync(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) *CodePipelineDeleteWebhookFuture

	DeregisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error)
	DeregisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) *CodePipelineDeregisterWebhookWithThirdPartyFuture

	DisableStageTransition(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error)
	DisableStageTransitionAsync(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) *CodePipelineDisableStageTransitionFuture

	EnableStageTransition(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error)
	EnableStageTransitionAsync(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) *CodePipelineEnableStageTransitionFuture

	GetJobDetails(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error)
	GetJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) *CodePipelineGetJobDetailsFuture

	GetPipeline(ctx workflow.Context, input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error)
	GetPipelineAsync(ctx workflow.Context, input *codepipeline.GetPipelineInput) *CodePipelineGetPipelineFuture

	GetPipelineExecution(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) *CodePipelineGetPipelineExecutionFuture

	GetPipelineState(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error)
	GetPipelineStateAsync(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) *CodePipelineGetPipelineStateFuture

	GetThirdPartyJobDetails(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	GetThirdPartyJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) *CodePipelineGetThirdPartyJobDetailsFuture

	ListActionExecutions(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionExecutionsAsync(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) *CodePipelineListActionExecutionsFuture

	ListActionTypes(ctx workflow.Context, input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error)
	ListActionTypesAsync(ctx workflow.Context, input *codepipeline.ListActionTypesInput) *CodePipelineListActionTypesFuture

	ListPipelineExecutions(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelineExecutionsAsync(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) *CodePipelineListPipelineExecutionsFuture

	ListPipelines(ctx workflow.Context, input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error)
	ListPipelinesAsync(ctx workflow.Context, input *codepipeline.ListPipelinesInput) *CodePipelineListPipelinesFuture

	ListTagsForResource(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) *CodePipelineListTagsForResourceFuture

	ListWebhooks(ctx workflow.Context, input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error)
	ListWebhooksAsync(ctx workflow.Context, input *codepipeline.ListWebhooksInput) *CodePipelineListWebhooksFuture

	PollForJobs(ctx workflow.Context, input *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error)
	PollForJobsAsync(ctx workflow.Context, input *codepipeline.PollForJobsInput) *CodePipelinePollForJobsFuture

	PollForThirdPartyJobs(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error)
	PollForThirdPartyJobsAsync(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) *CodePipelinePollForThirdPartyJobsFuture

	PutActionRevision(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error)
	PutActionRevisionAsync(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) *CodePipelinePutActionRevisionFuture

	PutApprovalResult(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error)
	PutApprovalResultAsync(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) *CodePipelinePutApprovalResultFuture

	PutJobFailureResult(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error)
	PutJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) *CodePipelinePutJobFailureResultFuture

	PutJobSuccessResult(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error)
	PutJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) *CodePipelinePutJobSuccessResultFuture

	PutThirdPartyJobFailureResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)
	PutThirdPartyJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) *CodePipelinePutThirdPartyJobFailureResultFuture

	PutThirdPartyJobSuccessResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)
	PutThirdPartyJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) *CodePipelinePutThirdPartyJobSuccessResultFuture

	PutWebhook(ctx workflow.Context, input *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error)
	PutWebhookAsync(ctx workflow.Context, input *codepipeline.PutWebhookInput) *CodePipelinePutWebhookFuture

	RegisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error)
	RegisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) *CodePipelineRegisterWebhookWithThirdPartyFuture

	RetryStageExecution(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error)
	RetryStageExecutionAsync(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) *CodePipelineRetryStageExecutionFuture

	StartPipelineExecution(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error)
	StartPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) *CodePipelineStartPipelineExecutionFuture

	StopPipelineExecution(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error)
	StopPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) *CodePipelineStopPipelineExecutionFuture

	TagResource(ctx workflow.Context, input *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codepipeline.TagResourceInput) *CodePipelineTagResourceFuture

	UntagResource(ctx workflow.Context, input *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codepipeline.UntagResourceInput) *CodePipelineUntagResourceFuture

	UpdatePipeline(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error)
	UpdatePipelineAsync(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) *CodePipelineUpdatePipelineFuture
}

func NewClient() Client {
	return &stub{}
}
