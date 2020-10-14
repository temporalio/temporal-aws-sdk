// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package sfnstub

import (
	"github.com/aws/aws-sdk-go/service/sfn"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SFNCreateActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNCreateActivityFuture) Get(ctx workflow.Context) (*sfn.CreateActivityOutput, error) {
	var output sfn.CreateActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNCreateStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNCreateStateMachineFuture) Get(ctx workflow.Context) (*sfn.CreateStateMachineOutput, error) {
	var output sfn.CreateStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDeleteActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDeleteActivityFuture) Get(ctx workflow.Context) (*sfn.DeleteActivityOutput, error) {
	var output sfn.DeleteActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDeleteStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDeleteStateMachineFuture) Get(ctx workflow.Context) (*sfn.DeleteStateMachineOutput, error) {
	var output sfn.DeleteStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDescribeActivityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDescribeActivityFuture) Get(ctx workflow.Context) (*sfn.DescribeActivityOutput, error) {
	var output sfn.DescribeActivityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDescribeExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDescribeExecutionFuture) Get(ctx workflow.Context) (*sfn.DescribeExecutionOutput, error) {
	var output sfn.DescribeExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDescribeStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDescribeStateMachineFuture) Get(ctx workflow.Context) (*sfn.DescribeStateMachineOutput, error) {
	var output sfn.DescribeStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNDescribeStateMachineForExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNDescribeStateMachineForExecutionFuture) Get(ctx workflow.Context) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	var output sfn.DescribeStateMachineForExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNGetActivityTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNGetActivityTaskFuture) Get(ctx workflow.Context) (*sfn.GetActivityTaskOutput, error) {
	var output sfn.GetActivityTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNGetExecutionHistoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNGetExecutionHistoryFuture) Get(ctx workflow.Context) (*sfn.GetExecutionHistoryOutput, error) {
	var output sfn.GetExecutionHistoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNListActivitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNListActivitiesFuture) Get(ctx workflow.Context) (*sfn.ListActivitiesOutput, error) {
	var output sfn.ListActivitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNListExecutionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNListExecutionsFuture) Get(ctx workflow.Context) (*sfn.ListExecutionsOutput, error) {
	var output sfn.ListExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNListStateMachinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNListStateMachinesFuture) Get(ctx workflow.Context) (*sfn.ListStateMachinesOutput, error) {
	var output sfn.ListStateMachinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNListTagsForResourceFuture) Get(ctx workflow.Context) (*sfn.ListTagsForResourceOutput, error) {
	var output sfn.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNSendTaskFailureFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNSendTaskFailureFuture) Get(ctx workflow.Context) (*sfn.SendTaskFailureOutput, error) {
	var output sfn.SendTaskFailureOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNSendTaskHeartbeatFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNSendTaskHeartbeatFuture) Get(ctx workflow.Context) (*sfn.SendTaskHeartbeatOutput, error) {
	var output sfn.SendTaskHeartbeatOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNSendTaskSuccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNSendTaskSuccessFuture) Get(ctx workflow.Context) (*sfn.SendTaskSuccessOutput, error) {
	var output sfn.SendTaskSuccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNStartExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNStartExecutionFuture) Get(ctx workflow.Context) (*sfn.StartExecutionOutput, error) {
	var output sfn.StartExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNStopExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNStopExecutionFuture) Get(ctx workflow.Context) (*sfn.StopExecutionOutput, error) {
	var output sfn.StopExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNTagResourceFuture) Get(ctx workflow.Context) (*sfn.TagResourceOutput, error) {
	var output sfn.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNUntagResourceFuture) Get(ctx workflow.Context) (*sfn.UntagResourceOutput, error) {
	var output sfn.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SFNUpdateStateMachineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SFNUpdateStateMachineFuture) Get(ctx workflow.Context) (*sfn.UpdateStateMachineOutput, error) {
	var output sfn.UpdateStateMachineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateActivity(ctx workflow.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	var output sfn.CreateActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.CreateActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateActivityAsync(ctx workflow.Context, input *sfn.CreateActivityInput) *SFNCreateActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.CreateActivity", input)
	return &SFNCreateActivityFuture{Future: future}
}

func (a *stub) CreateStateMachine(ctx workflow.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	var output sfn.CreateStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.CreateStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateStateMachineAsync(ctx workflow.Context, input *sfn.CreateStateMachineInput) *SFNCreateStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.CreateStateMachine", input)
	return &SFNCreateStateMachineFuture{Future: future}
}

func (a *stub) DeleteActivity(ctx workflow.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	var output sfn.DeleteActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DeleteActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteActivityAsync(ctx workflow.Context, input *sfn.DeleteActivityInput) *SFNDeleteActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DeleteActivity", input)
	return &SFNDeleteActivityFuture{Future: future}
}

func (a *stub) DeleteStateMachine(ctx workflow.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	var output sfn.DeleteStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DeleteStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteStateMachineAsync(ctx workflow.Context, input *sfn.DeleteStateMachineInput) *SFNDeleteStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DeleteStateMachine", input)
	return &SFNDeleteStateMachineFuture{Future: future}
}

func (a *stub) DescribeActivity(ctx workflow.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	var output sfn.DescribeActivityOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeActivity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeActivityAsync(ctx workflow.Context, input *sfn.DescribeActivityInput) *SFNDescribeActivityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeActivity", input)
	return &SFNDescribeActivityFuture{Future: future}
}

func (a *stub) DescribeExecution(ctx workflow.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	var output sfn.DescribeExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExecutionAsync(ctx workflow.Context, input *sfn.DescribeExecutionInput) *SFNDescribeExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeExecution", input)
	return &SFNDescribeExecutionFuture{Future: future}
}

func (a *stub) DescribeStateMachine(ctx workflow.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	var output sfn.DescribeStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStateMachineAsync(ctx workflow.Context, input *sfn.DescribeStateMachineInput) *SFNDescribeStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeStateMachine", input)
	return &SFNDescribeStateMachineFuture{Future: future}
}

func (a *stub) DescribeStateMachineForExecution(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	var output sfn.DescribeStateMachineForExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeStateMachineForExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStateMachineForExecutionAsync(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) *SFNDescribeStateMachineForExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.DescribeStateMachineForExecution", input)
	return &SFNDescribeStateMachineForExecutionFuture{Future: future}
}

func (a *stub) GetActivityTask(ctx workflow.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	var output sfn.GetActivityTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.GetActivityTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetActivityTaskAsync(ctx workflow.Context, input *sfn.GetActivityTaskInput) *SFNGetActivityTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.GetActivityTask", input)
	return &SFNGetActivityTaskFuture{Future: future}
}

func (a *stub) GetExecutionHistory(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	var output sfn.GetExecutionHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.GetExecutionHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetExecutionHistoryAsync(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) *SFNGetExecutionHistoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.GetExecutionHistory", input)
	return &SFNGetExecutionHistoryFuture{Future: future}
}

func (a *stub) ListActivities(ctx workflow.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	var output sfn.ListActivitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.ListActivities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListActivitiesAsync(ctx workflow.Context, input *sfn.ListActivitiesInput) *SFNListActivitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.ListActivities", input)
	return &SFNListActivitiesFuture{Future: future}
}

func (a *stub) ListExecutions(ctx workflow.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	var output sfn.ListExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.ListExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListExecutionsAsync(ctx workflow.Context, input *sfn.ListExecutionsInput) *SFNListExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.ListExecutions", input)
	return &SFNListExecutionsFuture{Future: future}
}

func (a *stub) ListStateMachines(ctx workflow.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	var output sfn.ListStateMachinesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.ListStateMachines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStateMachinesAsync(ctx workflow.Context, input *sfn.ListStateMachinesInput) *SFNListStateMachinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.ListStateMachines", input)
	return &SFNListStateMachinesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
	var output sfn.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *sfn.ListTagsForResourceInput) *SFNListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.ListTagsForResource", input)
	return &SFNListTagsForResourceFuture{Future: future}
}

func (a *stub) SendTaskFailure(ctx workflow.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	var output sfn.SendTaskFailureOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskFailure", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskFailureAsync(ctx workflow.Context, input *sfn.SendTaskFailureInput) *SFNSendTaskFailureFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskFailure", input)
	return &SFNSendTaskFailureFuture{Future: future}
}

func (a *stub) SendTaskHeartbeat(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	var output sfn.SendTaskHeartbeatOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskHeartbeat", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskHeartbeatAsync(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) *SFNSendTaskHeartbeatFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskHeartbeat", input)
	return &SFNSendTaskHeartbeatFuture{Future: future}
}

func (a *stub) SendTaskSuccess(ctx workflow.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	var output sfn.SendTaskSuccessOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskSuccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTaskSuccessAsync(ctx workflow.Context, input *sfn.SendTaskSuccessInput) *SFNSendTaskSuccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.SendTaskSuccess", input)
	return &SFNSendTaskSuccessFuture{Future: future}
}

func (a *stub) StartExecution(ctx workflow.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	var output sfn.StartExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.StartExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartExecutionAsync(ctx workflow.Context, input *sfn.StartExecutionInput) *SFNStartExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.StartExecution", input)
	return &SFNStartExecutionFuture{Future: future}
}

func (a *stub) StopExecution(ctx workflow.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	var output sfn.StopExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.StopExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopExecutionAsync(ctx workflow.Context, input *sfn.StopExecutionInput) *SFNStopExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.StopExecution", input)
	return &SFNStopExecutionFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
	var output sfn.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *sfn.TagResourceInput) *SFNTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.TagResource", input)
	return &SFNTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
	var output sfn.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *sfn.UntagResourceInput) *SFNUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.UntagResource", input)
	return &SFNUntagResourceFuture{Future: future}
}

func (a *stub) UpdateStateMachine(ctx workflow.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
	var output sfn.UpdateStateMachineOutput
	err := workflow.ExecuteActivity(ctx, "aws.sfn.UpdateStateMachine", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateStateMachineAsync(ctx workflow.Context, input *sfn.UpdateStateMachineInput) *SFNUpdateStateMachineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sfn.UpdateStateMachine", input)
	return &SFNUpdateStateMachineFuture{Future: future}
}
