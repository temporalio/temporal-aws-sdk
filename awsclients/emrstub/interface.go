// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package emrstub

import (
	"github.com/aws/aws-sdk-go/service/emr"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AddInstanceFleet(ctx workflow.Context, input *emr.AddInstanceFleetInput) (*emr.AddInstanceFleetOutput, error)
	AddInstanceFleetAsync(ctx workflow.Context, input *emr.AddInstanceFleetInput) *EMRAddInstanceFleetFuture

	AddInstanceGroups(ctx workflow.Context, input *emr.AddInstanceGroupsInput) (*emr.AddInstanceGroupsOutput, error)
	AddInstanceGroupsAsync(ctx workflow.Context, input *emr.AddInstanceGroupsInput) *EMRAddInstanceGroupsFuture

	AddJobFlowSteps(ctx workflow.Context, input *emr.AddJobFlowStepsInput) (*emr.AddJobFlowStepsOutput, error)
	AddJobFlowStepsAsync(ctx workflow.Context, input *emr.AddJobFlowStepsInput) *EMRAddJobFlowStepsFuture

	AddTags(ctx workflow.Context, input *emr.AddTagsInput) (*emr.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *emr.AddTagsInput) *EMRAddTagsFuture

	CancelSteps(ctx workflow.Context, input *emr.CancelStepsInput) (*emr.CancelStepsOutput, error)
	CancelStepsAsync(ctx workflow.Context, input *emr.CancelStepsInput) *EMRCancelStepsFuture

	CreateSecurityConfiguration(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) (*emr.CreateSecurityConfigurationOutput, error)
	CreateSecurityConfigurationAsync(ctx workflow.Context, input *emr.CreateSecurityConfigurationInput) *EMRCreateSecurityConfigurationFuture

	DeleteSecurityConfiguration(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) (*emr.DeleteSecurityConfigurationOutput, error)
	DeleteSecurityConfigurationAsync(ctx workflow.Context, input *emr.DeleteSecurityConfigurationInput) *EMRDeleteSecurityConfigurationFuture

	DescribeCluster(ctx workflow.Context, input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *EMRDescribeClusterFuture

	DescribeJobFlows(ctx workflow.Context, input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error)
	DescribeJobFlowsAsync(ctx workflow.Context, input *emr.DescribeJobFlowsInput) *EMRDescribeJobFlowsFuture

	DescribeNotebookExecution(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error)
	DescribeNotebookExecutionAsync(ctx workflow.Context, input *emr.DescribeNotebookExecutionInput) *EMRDescribeNotebookExecutionFuture

	DescribeSecurityConfiguration(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error)
	DescribeSecurityConfigurationAsync(ctx workflow.Context, input *emr.DescribeSecurityConfigurationInput) *EMRDescribeSecurityConfigurationFuture

	DescribeStep(ctx workflow.Context, input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error)
	DescribeStepAsync(ctx workflow.Context, input *emr.DescribeStepInput) *EMRDescribeStepFuture

	GetBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	GetBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.GetBlockPublicAccessConfigurationInput) *EMRGetBlockPublicAccessConfigurationFuture

	GetManagedScalingPolicy(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error)
	GetManagedScalingPolicyAsync(ctx workflow.Context, input *emr.GetManagedScalingPolicyInput) *EMRGetManagedScalingPolicyFuture

	ListBootstrapActions(ctx workflow.Context, input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error)
	ListBootstrapActionsAsync(ctx workflow.Context, input *emr.ListBootstrapActionsInput) *EMRListBootstrapActionsFuture

	ListClusters(ctx workflow.Context, input *emr.ListClustersInput) (*emr.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *emr.ListClustersInput) *EMRListClustersFuture

	ListInstanceFleets(ctx workflow.Context, input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error)
	ListInstanceFleetsAsync(ctx workflow.Context, input *emr.ListInstanceFleetsInput) *EMRListInstanceFleetsFuture

	ListInstanceGroups(ctx workflow.Context, input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error)
	ListInstanceGroupsAsync(ctx workflow.Context, input *emr.ListInstanceGroupsInput) *EMRListInstanceGroupsFuture

	ListInstances(ctx workflow.Context, input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error)
	ListInstancesAsync(ctx workflow.Context, input *emr.ListInstancesInput) *EMRListInstancesFuture

	ListNotebookExecutions(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error)
	ListNotebookExecutionsAsync(ctx workflow.Context, input *emr.ListNotebookExecutionsInput) *EMRListNotebookExecutionsFuture

	ListSecurityConfigurations(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error)
	ListSecurityConfigurationsAsync(ctx workflow.Context, input *emr.ListSecurityConfigurationsInput) *EMRListSecurityConfigurationsFuture

	ListSteps(ctx workflow.Context, input *emr.ListStepsInput) (*emr.ListStepsOutput, error)
	ListStepsAsync(ctx workflow.Context, input *emr.ListStepsInput) *EMRListStepsFuture

	ModifyCluster(ctx workflow.Context, input *emr.ModifyClusterInput) (*emr.ModifyClusterOutput, error)
	ModifyClusterAsync(ctx workflow.Context, input *emr.ModifyClusterInput) *EMRModifyClusterFuture

	ModifyInstanceFleet(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) (*emr.ModifyInstanceFleetOutput, error)
	ModifyInstanceFleetAsync(ctx workflow.Context, input *emr.ModifyInstanceFleetInput) *EMRModifyInstanceFleetFuture

	ModifyInstanceGroups(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) (*emr.ModifyInstanceGroupsOutput, error)
	ModifyInstanceGroupsAsync(ctx workflow.Context, input *emr.ModifyInstanceGroupsInput) *EMRModifyInstanceGroupsFuture

	PutAutoScalingPolicy(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) (*emr.PutAutoScalingPolicyOutput, error)
	PutAutoScalingPolicyAsync(ctx workflow.Context, input *emr.PutAutoScalingPolicyInput) *EMRPutAutoScalingPolicyFuture

	PutBlockPublicAccessConfiguration(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) (*emr.PutBlockPublicAccessConfigurationOutput, error)
	PutBlockPublicAccessConfigurationAsync(ctx workflow.Context, input *emr.PutBlockPublicAccessConfigurationInput) *EMRPutBlockPublicAccessConfigurationFuture

	PutManagedScalingPolicy(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) (*emr.PutManagedScalingPolicyOutput, error)
	PutManagedScalingPolicyAsync(ctx workflow.Context, input *emr.PutManagedScalingPolicyInput) *EMRPutManagedScalingPolicyFuture

	RemoveAutoScalingPolicy(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) (*emr.RemoveAutoScalingPolicyOutput, error)
	RemoveAutoScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveAutoScalingPolicyInput) *EMRRemoveAutoScalingPolicyFuture

	RemoveManagedScalingPolicy(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) (*emr.RemoveManagedScalingPolicyOutput, error)
	RemoveManagedScalingPolicyAsync(ctx workflow.Context, input *emr.RemoveManagedScalingPolicyInput) *EMRRemoveManagedScalingPolicyFuture

	RemoveTags(ctx workflow.Context, input *emr.RemoveTagsInput) (*emr.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *emr.RemoveTagsInput) *EMRRemoveTagsFuture

	RunJobFlow(ctx workflow.Context, input *emr.RunJobFlowInput) (*emr.RunJobFlowOutput, error)
	RunJobFlowAsync(ctx workflow.Context, input *emr.RunJobFlowInput) *EMRRunJobFlowFuture

	SetTerminationProtection(ctx workflow.Context, input *emr.SetTerminationProtectionInput) (*emr.SetTerminationProtectionOutput, error)
	SetTerminationProtectionAsync(ctx workflow.Context, input *emr.SetTerminationProtectionInput) *EMRSetTerminationProtectionFuture

	SetVisibleToAllUsers(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) (*emr.SetVisibleToAllUsersOutput, error)
	SetVisibleToAllUsersAsync(ctx workflow.Context, input *emr.SetVisibleToAllUsersInput) *EMRSetVisibleToAllUsersFuture

	StartNotebookExecution(ctx workflow.Context, input *emr.StartNotebookExecutionInput) (*emr.StartNotebookExecutionOutput, error)
	StartNotebookExecutionAsync(ctx workflow.Context, input *emr.StartNotebookExecutionInput) *EMRStartNotebookExecutionFuture

	StopNotebookExecution(ctx workflow.Context, input *emr.StopNotebookExecutionInput) (*emr.StopNotebookExecutionOutput, error)
	StopNotebookExecutionAsync(ctx workflow.Context, input *emr.StopNotebookExecutionInput) *EMRStopNotebookExecutionFuture

	TerminateJobFlows(ctx workflow.Context, input *emr.TerminateJobFlowsInput) (*emr.TerminateJobFlowsOutput, error)
	TerminateJobFlowsAsync(ctx workflow.Context, input *emr.TerminateJobFlowsInput) *EMRTerminateJobFlowsFuture

	WaitUntilClusterRunning(ctx workflow.Context, input *emr.DescribeClusterInput) error
	WaitUntilClusterRunningAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *awsclients.VoidFuture

	WaitUntilClusterTerminated(ctx workflow.Context, input *emr.DescribeClusterInput) error
	WaitUntilClusterTerminatedAsync(ctx workflow.Context, input *emr.DescribeClusterInput) *awsclients.VoidFuture

	WaitUntilStepComplete(ctx workflow.Context, input *emr.DescribeStepInput) error
	WaitUntilStepCompleteAsync(ctx workflow.Context, input *emr.DescribeStepInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

