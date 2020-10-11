// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package datapipelinestub

import (
	"github.com/aws/aws-sdk-go/service/datapipeline"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type DataPipelineActivatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineActivatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.ActivatePipelineOutput, error) {
	var output datapipeline.ActivatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineAddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineAddTagsFuture) Get(ctx workflow.Context) (*datapipeline.AddTagsOutput, error) {
	var output datapipeline.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineCreatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineCreatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.CreatePipelineOutput, error) {
	var output datapipeline.CreatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineDeactivatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineDeactivatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.DeactivatePipelineOutput, error) {
	var output datapipeline.DeactivatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineDeletePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineDeletePipelineFuture) Get(ctx workflow.Context) (*datapipeline.DeletePipelineOutput, error) {
	var output datapipeline.DeletePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineDescribeObjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineDescribeObjectsFuture) Get(ctx workflow.Context) (*datapipeline.DescribeObjectsOutput, error) {
	var output datapipeline.DescribeObjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineDescribePipelinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineDescribePipelinesFuture) Get(ctx workflow.Context) (*datapipeline.DescribePipelinesOutput, error) {
	var output datapipeline.DescribePipelinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineEvaluateExpressionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineEvaluateExpressionFuture) Get(ctx workflow.Context) (*datapipeline.EvaluateExpressionOutput, error) {
	var output datapipeline.EvaluateExpressionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineGetPipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineGetPipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.GetPipelineDefinitionOutput, error) {
	var output datapipeline.GetPipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineListPipelinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineListPipelinesFuture) Get(ctx workflow.Context) (*datapipeline.ListPipelinesOutput, error) {
	var output datapipeline.ListPipelinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelinePollForTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelinePollForTaskFuture) Get(ctx workflow.Context) (*datapipeline.PollForTaskOutput, error) {
	var output datapipeline.PollForTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelinePutPipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelinePutPipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.PutPipelineDefinitionOutput, error) {
	var output datapipeline.PutPipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineQueryObjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineQueryObjectsFuture) Get(ctx workflow.Context) (*datapipeline.QueryObjectsOutput, error) {
	var output datapipeline.QueryObjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineRemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineRemoveTagsFuture) Get(ctx workflow.Context) (*datapipeline.RemoveTagsOutput, error) {
	var output datapipeline.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineReportTaskProgressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineReportTaskProgressFuture) Get(ctx workflow.Context) (*datapipeline.ReportTaskProgressOutput, error) {
	var output datapipeline.ReportTaskProgressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineReportTaskRunnerHeartbeatFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineReportTaskRunnerHeartbeatFuture) Get(ctx workflow.Context) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	var output datapipeline.ReportTaskRunnerHeartbeatOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineSetStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineSetStatusFuture) Get(ctx workflow.Context) (*datapipeline.SetStatusOutput, error) {
	var output datapipeline.SetStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineSetTaskStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineSetTaskStatusFuture) Get(ctx workflow.Context) (*datapipeline.SetTaskStatusOutput, error) {
	var output datapipeline.SetTaskStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataPipelineValidatePipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataPipelineValidatePipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	var output datapipeline.ValidatePipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivatePipeline(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
	var output datapipeline.ActivatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ActivatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivatePipelineAsync(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) *DataPipelineActivatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ActivatePipeline", input)
	return &DataPipelineActivatePipelineFuture{Future: future}
}

func (a *stub) AddTags(ctx workflow.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
	var output datapipeline.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *datapipeline.AddTagsInput) *DataPipelineAddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.AddTags", input)
	return &DataPipelineAddTagsFuture{Future: future}
}

func (a *stub) CreatePipeline(ctx workflow.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
	var output datapipeline.CreatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.CreatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePipelineAsync(ctx workflow.Context, input *datapipeline.CreatePipelineInput) *DataPipelineCreatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.CreatePipeline", input)
	return &DataPipelineCreatePipelineFuture{Future: future}
}

func (a *stub) DeactivatePipeline(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
	var output datapipeline.DeactivatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeactivatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivatePipelineAsync(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) *DataPipelineDeactivatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeactivatePipeline", input)
	return &DataPipelineDeactivatePipelineFuture{Future: future}
}

func (a *stub) DeletePipeline(ctx workflow.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
	var output datapipeline.DeletePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeletePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePipelineAsync(ctx workflow.Context, input *datapipeline.DeletePipelineInput) *DataPipelineDeletePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeletePipeline", input)
	return &DataPipelineDeletePipelineFuture{Future: future}
}

func (a *stub) DescribeObjects(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
	var output datapipeline.DescribeObjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribeObjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeObjectsAsync(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) *DataPipelineDescribeObjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribeObjects", input)
	return &DataPipelineDescribeObjectsFuture{Future: future}
}

func (a *stub) DescribePipelines(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
	var output datapipeline.DescribePipelinesOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribePipelines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePipelinesAsync(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) *DataPipelineDescribePipelinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribePipelines", input)
	return &DataPipelineDescribePipelinesFuture{Future: future}
}

func (a *stub) EvaluateExpression(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
	var output datapipeline.EvaluateExpressionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.EvaluateExpression", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EvaluateExpressionAsync(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) *DataPipelineEvaluateExpressionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.EvaluateExpression", input)
	return &DataPipelineEvaluateExpressionFuture{Future: future}
}

func (a *stub) GetPipelineDefinition(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
	var output datapipeline.GetPipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.GetPipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) *DataPipelineGetPipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.GetPipelineDefinition", input)
	return &DataPipelineGetPipelineDefinitionFuture{Future: future}
}

func (a *stub) ListPipelines(ctx workflow.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
	var output datapipeline.ListPipelinesOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ListPipelines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPipelinesAsync(ctx workflow.Context, input *datapipeline.ListPipelinesInput) *DataPipelineListPipelinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ListPipelines", input)
	return &DataPipelineListPipelinesFuture{Future: future}
}

func (a *stub) PollForTask(ctx workflow.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
	var output datapipeline.PollForTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.PollForTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PollForTaskAsync(ctx workflow.Context, input *datapipeline.PollForTaskInput) *DataPipelinePollForTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.PollForTask", input)
	return &DataPipelinePollForTaskFuture{Future: future}
}

func (a *stub) PutPipelineDefinition(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
	var output datapipeline.PutPipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.PutPipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) *DataPipelinePutPipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.PutPipelineDefinition", input)
	return &DataPipelinePutPipelineDefinitionFuture{Future: future}
}

func (a *stub) QueryObjects(ctx workflow.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
	var output datapipeline.QueryObjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.QueryObjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryObjectsAsync(ctx workflow.Context, input *datapipeline.QueryObjectsInput) *DataPipelineQueryObjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.QueryObjects", input)
	return &DataPipelineQueryObjectsFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
	var output datapipeline.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *datapipeline.RemoveTagsInput) *DataPipelineRemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.RemoveTags", input)
	return &DataPipelineRemoveTagsFuture{Future: future}
}

func (a *stub) ReportTaskProgress(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
	var output datapipeline.ReportTaskProgressOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskProgress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReportTaskProgressAsync(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) *DataPipelineReportTaskProgressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskProgress", input)
	return &DataPipelineReportTaskProgressFuture{Future: future}
}

func (a *stub) ReportTaskRunnerHeartbeat(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	var output datapipeline.ReportTaskRunnerHeartbeatOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskRunnerHeartbeat", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReportTaskRunnerHeartbeatAsync(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) *DataPipelineReportTaskRunnerHeartbeatFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskRunnerHeartbeat", input)
	return &DataPipelineReportTaskRunnerHeartbeatFuture{Future: future}
}

func (a *stub) SetStatus(ctx workflow.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
	var output datapipeline.SetStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetStatusAsync(ctx workflow.Context, input *datapipeline.SetStatusInput) *DataPipelineSetStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetStatus", input)
	return &DataPipelineSetStatusFuture{Future: future}
}

func (a *stub) SetTaskStatus(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
	var output datapipeline.SetTaskStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetTaskStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetTaskStatusAsync(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) *DataPipelineSetTaskStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetTaskStatus", input)
	return &DataPipelineSetTaskStatusFuture{Future: future}
}

func (a *stub) ValidatePipelineDefinition(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	var output datapipeline.ValidatePipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ValidatePipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ValidatePipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) *DataPipelineValidatePipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ValidatePipelineDefinition", input)
	return &DataPipelineValidatePipelineDefinitionFuture{Future: future}
}
