// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package emrstub

import (
	"github.com/aws/aws-sdk-go/service/emr"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddInstanceFleetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddInstanceFleetFuture) Get(ctx workflow.Context) (*emr.AddInstanceFleetOutput, error) {
	var output emr.AddInstanceFleetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddInstanceGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddInstanceGroupsFuture) Get(ctx workflow.Context) (*emr.AddInstanceGroupsOutput, error) {
	var output emr.AddInstanceGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddJobFlowStepsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddJobFlowStepsFuture) Get(ctx workflow.Context) (*emr.AddJobFlowStepsOutput, error) {
	var output emr.AddJobFlowStepsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsFuture) Get(ctx workflow.Context) (*emr.AddTagsOutput, error) {
	var output emr.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CancelStepsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CancelStepsFuture) Get(ctx workflow.Context) (*emr.CancelStepsOutput, error) {
	var output emr.CancelStepsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSecurityConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSecurityConfigurationFuture) Get(ctx workflow.Context) (*emr.CreateSecurityConfigurationOutput, error) {
	var output emr.CreateSecurityConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSecurityConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSecurityConfigurationFuture) Get(ctx workflow.Context) (*emr.DeleteSecurityConfigurationOutput, error) {
	var output emr.DeleteSecurityConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeClusterFuture) Get(ctx workflow.Context) (*emr.DescribeClusterOutput, error) {
	var output emr.DescribeClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobFlowsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobFlowsFuture) Get(ctx workflow.Context) (*emr.DescribeJobFlowsOutput, error) {
	var output emr.DescribeJobFlowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeNotebookExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeNotebookExecutionFuture) Get(ctx workflow.Context) (*emr.DescribeNotebookExecutionOutput, error) {
	var output emr.DescribeNotebookExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSecurityConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSecurityConfigurationFuture) Get(ctx workflow.Context) (*emr.DescribeSecurityConfigurationOutput, error) {
	var output emr.DescribeSecurityConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStepFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStepFuture) Get(ctx workflow.Context) (*emr.DescribeStepOutput, error) {
	var output emr.DescribeStepOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBlockPublicAccessConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBlockPublicAccessConfigurationFuture) Get(ctx workflow.Context) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	var output emr.GetBlockPublicAccessConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetManagedScalingPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetManagedScalingPolicyFuture) Get(ctx workflow.Context) (*emr.GetManagedScalingPolicyOutput, error) {
	var output emr.GetManagedScalingPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBootstrapActionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBootstrapActionsFuture) Get(ctx workflow.Context) (*emr.ListBootstrapActionsOutput, error) {
	var output emr.ListBootstrapActionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListClustersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListClustersFuture) Get(ctx workflow.Context) (*emr.ListClustersOutput, error) {
	var output emr.ListClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInstanceFleetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInstanceFleetsFuture) Get(ctx workflow.Context) (*emr.ListInstanceFleetsOutput, error) {
	var output emr.ListInstanceFleetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInstanceGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInstanceGroupsFuture) Get(ctx workflow.Context) (*emr.ListInstanceGroupsOutput, error) {
	var output emr.ListInstanceGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInstancesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInstancesFuture) Get(ctx workflow.Context) (*emr.ListInstancesOutput, error) {
	var output emr.ListInstancesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListNotebookExecutionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListNotebookExecutionsFuture) Get(ctx workflow.Context) (*emr.ListNotebookExecutionsOutput, error) {
	var output emr.ListNotebookExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSecurityConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSecurityConfigurationsFuture) Get(ctx workflow.Context) (*emr.ListSecurityConfigurationsOutput, error) {
	var output emr.ListSecurityConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListStepsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListStepsFuture) Get(ctx workflow.Context) (*emr.ListStepsOutput, error) {
	var output emr.ListStepsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyClusterFuture) Get(ctx workflow.Context) (*emr.ModifyClusterOutput, error) {
	var output emr.ModifyClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyInstanceFleetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyInstanceFleetFuture) Get(ctx workflow.Context) (*emr.ModifyInstanceFleetOutput, error) {
	var output emr.ModifyInstanceFleetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyInstanceGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyInstanceGroupsFuture) Get(ctx workflow.Context) (*emr.ModifyInstanceGroupsOutput, error) {
	var output emr.ModifyInstanceGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutAutoScalingPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutAutoScalingPolicyFuture) Get(ctx workflow.Context) (*emr.PutAutoScalingPolicyOutput, error) {
	var output emr.PutAutoScalingPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutBlockPublicAccessConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutBlockPublicAccessConfigurationFuture) Get(ctx workflow.Context) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
	var output emr.PutBlockPublicAccessConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutManagedScalingPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutManagedScalingPolicyFuture) Get(ctx workflow.Context) (*emr.PutManagedScalingPolicyOutput, error) {
	var output emr.PutManagedScalingPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveAutoScalingPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveAutoScalingPolicyFuture) Get(ctx workflow.Context) (*emr.RemoveAutoScalingPolicyOutput, error) {
	var output emr.RemoveAutoScalingPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveManagedScalingPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveManagedScalingPolicyFuture) Get(ctx workflow.Context) (*emr.RemoveManagedScalingPolicyOutput, error) {
	var output emr.RemoveManagedScalingPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFuture) Get(ctx workflow.Context) (*emr.RemoveTagsOutput, error) {
	var output emr.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RunJobFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RunJobFlowFuture) Get(ctx workflow.Context) (*emr.RunJobFlowOutput, error) {
	var output emr.RunJobFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetTerminationProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetTerminationProtectionFuture) Get(ctx workflow.Context) (*emr.SetTerminationProtectionOutput, error) {
	var output emr.SetTerminationProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetVisibleToAllUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetVisibleToAllUsersFuture) Get(ctx workflow.Context) (*emr.SetVisibleToAllUsersOutput, error) {
	var output emr.SetVisibleToAllUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartNotebookExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartNotebookExecutionFuture) Get(ctx workflow.Context) (*emr.StartNotebookExecutionOutput, error) {
	var output emr.StartNotebookExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopNotebookExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopNotebookExecutionFuture) Get(ctx workflow.Context) (*emr.StopNotebookExecutionOutput, error) {
	var output emr.StopNotebookExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TerminateJobFlowsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TerminateJobFlowsFuture) Get(ctx workflow.Context) (*emr.TerminateJobFlowsOutput, error) {
	var output emr.TerminateJobFlowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddInstanceFleet(ctx workflow.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error) {
	var output emr.AddInstanceFleetOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.AddInstanceFleet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddInstanceFleetAsync(ctx workflow.Context, input *emr.AddInstanceFleetInput) *AddInstanceFleetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.AddInstanceFleet", input)
	return &AddInstanceFleetFuture{Future: future}
}

func (a *stub) AddInstanceGroups(ctx workflow.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error) {
	var output emr.AddInstanceGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.AddInstanceGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddInstanceGroupsAsync(ctx workflow.Context, input *emr.AddInstanceGroupsInput) *AddInstanceGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.AddInstanceGroups", input)
	return &AddInstanceGroupsFuture{Future: future}
}

func (a *stub) AddJobFlowSteps(ctx workflow.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error) {
	var output emr.AddJobFlowStepsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.AddJobFlowSteps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddJobFlowStepsAsync(ctx workflow.Context, input *emr.AddJobFlowStepsInput) *AddJobFlowStepsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.AddJobFlowSteps", input)
	return &AddJobFlowStepsFuture{Future: future}
}

func (a *stub) AddTags(ctx workflow.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error) {
	var output emr.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *emr.AddTagsInput) *AddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.AddTags", input)
	return &AddTagsFuture{Future: future}
}

func (a *stub) CancelSteps(ctx workflow.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error) {
	var output emr.CancelStepsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.CancelSteps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelStepsAsync(ctx workflow.Context, input *emr.CancelStepsInput) *CancelStepsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.CancelSteps", input)
	return &CancelStepsFuture{Future: future}
}

func (a *stub) CreateSecurityConfiguration(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error) {
	var output emr.CreateSecurityConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.CreateSecurityConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSecurityConfigurationAsync(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) *CreateSecurityConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.CreateSecurityConfiguration", input)
	return &CreateSecurityConfigurationFuture{Future: future}
}

func (a *stub) DeleteSecurityConfiguration(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error) {
	var output emr.DeleteSecurityConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DeleteSecurityConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSecurityConfigurationAsync(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) *DeleteSecurityConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DeleteSecurityConfiguration", input)
	return &DeleteSecurityConfigurationFuture{Future: future}
}

func (a *stub) DescribeCluster(ctx workflow.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
	var output emr.DescribeClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DescribeCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeClusterAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *DescribeClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DescribeCluster", input)
	return &DescribeClusterFuture{Future: future}
}

func (a *stub) DescribeJobFlows(ctx workflow.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
	var output emr.DescribeJobFlowsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DescribeJobFlows", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobFlowsAsync(ctx workflow.Context, input *emr.DescribeJobFlowsInput) *DescribeJobFlowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DescribeJobFlows", input)
	return &DescribeJobFlowsFuture{Future: future}
}

func (a *stub) DescribeNotebookExecution(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
	var output emr.DescribeNotebookExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DescribeNotebookExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeNotebookExecutionAsync(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) *DescribeNotebookExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DescribeNotebookExecution", input)
	return &DescribeNotebookExecutionFuture{Future: future}
}

func (a *stub) DescribeSecurityConfiguration(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
	var output emr.DescribeSecurityConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DescribeSecurityConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSecurityConfigurationAsync(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) *DescribeSecurityConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DescribeSecurityConfiguration", input)
	return &DescribeSecurityConfigurationFuture{Future: future}
}

func (a *stub) DescribeStep(ctx workflow.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
	var output emr.DescribeStepOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.DescribeStep", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStepAsync(ctx workflow.Context, input *emr.DescribeStepInput) *DescribeStepFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.DescribeStep", input)
	return &DescribeStepFuture{Future: future}
}

func (a *stub) GetBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	var output emr.GetBlockPublicAccessConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.GetBlockPublicAccessConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) *GetBlockPublicAccessConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.GetBlockPublicAccessConfiguration", input)
	return &GetBlockPublicAccessConfigurationFuture{Future: future}
}

func (a *stub) GetManagedScalingPolicy(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
	var output emr.GetManagedScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.GetManagedScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetManagedScalingPolicyAsync(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) *GetManagedScalingPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.GetManagedScalingPolicy", input)
	return &GetManagedScalingPolicyFuture{Future: future}
}

func (a *stub) ListBootstrapActions(ctx workflow.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
	var output emr.ListBootstrapActionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListBootstrapActions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBootstrapActionsAsync(ctx workflow.Context, input *emr.ListBootstrapActionsInput) *ListBootstrapActionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListBootstrapActions", input)
	return &ListBootstrapActionsFuture{Future: future}
}

func (a *stub) ListClusters(ctx workflow.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
	var output emr.ListClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListClustersAsync(ctx workflow.Context, input *emr.ListClustersInput) *ListClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListClusters", input)
	return &ListClustersFuture{Future: future}
}

func (a *stub) ListInstanceFleets(ctx workflow.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
	var output emr.ListInstanceFleetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListInstanceFleets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInstanceFleetsAsync(ctx workflow.Context, input *emr.ListInstanceFleetsInput) *ListInstanceFleetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListInstanceFleets", input)
	return &ListInstanceFleetsFuture{Future: future}
}

func (a *stub) ListInstanceGroups(ctx workflow.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
	var output emr.ListInstanceGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListInstanceGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInstanceGroupsAsync(ctx workflow.Context, input *emr.ListInstanceGroupsInput) *ListInstanceGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListInstanceGroups", input)
	return &ListInstanceGroupsFuture{Future: future}
}

func (a *stub) ListInstances(ctx workflow.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
	var output emr.ListInstancesOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListInstances", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInstancesAsync(ctx workflow.Context, input *emr.ListInstancesInput) *ListInstancesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListInstances", input)
	return &ListInstancesFuture{Future: future}
}

func (a *stub) ListNotebookExecutions(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
	var output emr.ListNotebookExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListNotebookExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNotebookExecutionsAsync(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) *ListNotebookExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListNotebookExecutions", input)
	return &ListNotebookExecutionsFuture{Future: future}
}

func (a *stub) ListSecurityConfigurations(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
	var output emr.ListSecurityConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListSecurityConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSecurityConfigurationsAsync(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) *ListSecurityConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListSecurityConfigurations", input)
	return &ListSecurityConfigurationsFuture{Future: future}
}

func (a *stub) ListSteps(ctx workflow.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
	var output emr.ListStepsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ListSteps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStepsAsync(ctx workflow.Context, input *emr.ListStepsInput) *ListStepsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ListSteps", input)
	return &ListStepsFuture{Future: future}
}

func (a *stub) ModifyCluster(ctx workflow.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error) {
	var output emr.ModifyClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ModifyCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyClusterAsync(ctx workflow.Context, input *emr.ModifyClusterInput) *ModifyClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ModifyCluster", input)
	return &ModifyClusterFuture{Future: future}
}

func (a *stub) ModifyInstanceFleet(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error) {
	var output emr.ModifyInstanceFleetOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ModifyInstanceFleet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyInstanceFleetAsync(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) *ModifyInstanceFleetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ModifyInstanceFleet", input)
	return &ModifyInstanceFleetFuture{Future: future}
}

func (a *stub) ModifyInstanceGroups(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error) {
	var output emr.ModifyInstanceGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.ModifyInstanceGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyInstanceGroupsAsync(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) *ModifyInstanceGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.ModifyInstanceGroups", input)
	return &ModifyInstanceGroupsFuture{Future: future}
}

func (a *stub) PutAutoScalingPolicy(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error) {
	var output emr.PutAutoScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.PutAutoScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAutoScalingPolicyAsync(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) *PutAutoScalingPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.PutAutoScalingPolicy", input)
	return &PutAutoScalingPolicyFuture{Future: future}
}

func (a *stub) PutBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error) {
	var output emr.PutBlockPublicAccessConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.PutBlockPublicAccessConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) *PutBlockPublicAccessConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.PutBlockPublicAccessConfiguration", input)
	return &PutBlockPublicAccessConfigurationFuture{Future: future}
}

func (a *stub) PutManagedScalingPolicy(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error) {
	var output emr.PutManagedScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.PutManagedScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutManagedScalingPolicyAsync(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) *PutManagedScalingPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.PutManagedScalingPolicy", input)
	return &PutManagedScalingPolicyFuture{Future: future}
}

func (a *stub) RemoveAutoScalingPolicy(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error) {
	var output emr.RemoveAutoScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.RemoveAutoScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveAutoScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) *RemoveAutoScalingPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.RemoveAutoScalingPolicy", input)
	return &RemoveAutoScalingPolicyFuture{Future: future}
}

func (a *stub) RemoveManagedScalingPolicy(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error) {
	var output emr.RemoveManagedScalingPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.RemoveManagedScalingPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveManagedScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) *RemoveManagedScalingPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.RemoveManagedScalingPolicy", input)
	return &RemoveManagedScalingPolicyFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error) {
	var output emr.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *emr.RemoveTagsInput) *RemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.RemoveTags", input)
	return &RemoveTagsFuture{Future: future}
}

func (a *stub) RunJobFlow(ctx workflow.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error) {
	var output emr.RunJobFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.RunJobFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RunJobFlowAsync(ctx workflow.Context, input *emr.RunJobFlowInput) *RunJobFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.RunJobFlow", input)
	return &RunJobFlowFuture{Future: future}
}

func (a *stub) SetTerminationProtection(ctx workflow.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error) {
	var output emr.SetTerminationProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.SetTerminationProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetTerminationProtectionAsync(ctx workflow.Context, input *emr.SetTerminationProtectionInput) *SetTerminationProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.SetTerminationProtection", input)
	return &SetTerminationProtectionFuture{Future: future}
}

func (a *stub) SetVisibleToAllUsers(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error) {
	var output emr.SetVisibleToAllUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.SetVisibleToAllUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetVisibleToAllUsersAsync(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) *SetVisibleToAllUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.SetVisibleToAllUsers", input)
	return &SetVisibleToAllUsersFuture{Future: future}
}

func (a *stub) StartNotebookExecution(ctx workflow.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error) {
	var output emr.StartNotebookExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.StartNotebookExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartNotebookExecutionAsync(ctx workflow.Context, input *emr.StartNotebookExecutionInput) *StartNotebookExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.StartNotebookExecution", input)
	return &StartNotebookExecutionFuture{Future: future}
}

func (a *stub) StopNotebookExecution(ctx workflow.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error) {
	var output emr.StopNotebookExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.StopNotebookExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopNotebookExecutionAsync(ctx workflow.Context, input *emr.StopNotebookExecutionInput) *StopNotebookExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.StopNotebookExecution", input)
	return &StopNotebookExecutionFuture{Future: future}
}

func (a *stub) TerminateJobFlows(ctx workflow.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error) {
	var output emr.TerminateJobFlowsOutput
	err := workflow.ExecuteActivity(ctx, "aws.emr.TerminateJobFlows", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TerminateJobFlowsAsync(ctx workflow.Context, input *emr.TerminateJobFlowsInput) *TerminateJobFlowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.TerminateJobFlows", input)
	return &TerminateJobFlowsFuture{Future: future}
}

func (a *stub) WaitUntilClusterRunning(ctx workflow.Context, input *emr.DescribeClusterInput) error {
	return workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilClusterRunning", input).Get(ctx, nil)
}

func (a *stub) WaitUntilClusterRunningAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilClusterRunning", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilClusterTerminated(ctx workflow.Context, input *emr.DescribeClusterInput) error {
	return workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilClusterTerminated", input).Get(ctx, nil)
}

func (a *stub) WaitUntilClusterTerminatedAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilClusterTerminated", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilStepComplete(ctx workflow.Context, input *emr.DescribeStepInput) error {
	return workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilStepComplete", input).Get(ctx, nil)
}

func (a *stub) WaitUntilStepCompleteAsync(ctx workflow.Context, input *emr.DescribeStepInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.emr.WaitUntilStepComplete", input)
	return clients.NewVoidFuture(future)
}
