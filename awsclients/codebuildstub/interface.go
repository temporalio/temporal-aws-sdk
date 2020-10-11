// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package codebuildstub

import (
	"github.com/aws/aws-sdk-go/service/codebuild"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	BatchDeleteBuilds(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error)
	BatchDeleteBuildsAsync(ctx workflow.Context, input *codebuild.BatchDeleteBuildsInput) *CodeBuildBatchDeleteBuildsFuture

	BatchGetBuildBatches(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error)
	BatchGetBuildBatchesAsync(ctx workflow.Context, input *codebuild.BatchGetBuildBatchesInput) *CodeBuildBatchGetBuildBatchesFuture

	BatchGetBuilds(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsAsync(ctx workflow.Context, input *codebuild.BatchGetBuildsInput) *CodeBuildBatchGetBuildsFuture

	BatchGetProjects(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsAsync(ctx workflow.Context, input *codebuild.BatchGetProjectsInput) *CodeBuildBatchGetProjectsFuture

	BatchGetReportGroups(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error)
	BatchGetReportGroupsAsync(ctx workflow.Context, input *codebuild.BatchGetReportGroupsInput) *CodeBuildBatchGetReportGroupsFuture

	BatchGetReports(ctx workflow.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error)
	BatchGetReportsAsync(ctx workflow.Context, input *codebuild.BatchGetReportsInput) *CodeBuildBatchGetReportsFuture

	CreateProject(ctx workflow.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *codebuild.CreateProjectInput) *CodeBuildCreateProjectFuture

	CreateReportGroup(ctx workflow.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error)
	CreateReportGroupAsync(ctx workflow.Context, input *codebuild.CreateReportGroupInput) *CodeBuildCreateReportGroupFuture

	CreateWebhook(ctx workflow.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error)
	CreateWebhookAsync(ctx workflow.Context, input *codebuild.CreateWebhookInput) *CodeBuildCreateWebhookFuture

	DeleteBuildBatch(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error)
	DeleteBuildBatchAsync(ctx workflow.Context, input *codebuild.DeleteBuildBatchInput) *CodeBuildDeleteBuildBatchFuture

	DeleteProject(ctx workflow.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *codebuild.DeleteProjectInput) *CodeBuildDeleteProjectFuture

	DeleteReport(ctx workflow.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error)
	DeleteReportAsync(ctx workflow.Context, input *codebuild.DeleteReportInput) *CodeBuildDeleteReportFuture

	DeleteReportGroup(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error)
	DeleteReportGroupAsync(ctx workflow.Context, input *codebuild.DeleteReportGroupInput) *CodeBuildDeleteReportGroupFuture

	DeleteResourcePolicy(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *codebuild.DeleteResourcePolicyInput) *CodeBuildDeleteResourcePolicyFuture

	DeleteSourceCredentials(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error)
	DeleteSourceCredentialsAsync(ctx workflow.Context, input *codebuild.DeleteSourceCredentialsInput) *CodeBuildDeleteSourceCredentialsFuture

	DeleteWebhook(ctx workflow.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error)
	DeleteWebhookAsync(ctx workflow.Context, input *codebuild.DeleteWebhookInput) *CodeBuildDeleteWebhookFuture

	DescribeCodeCoverages(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error)
	DescribeCodeCoveragesAsync(ctx workflow.Context, input *codebuild.DescribeCodeCoveragesInput) *CodeBuildDescribeCodeCoveragesFuture

	DescribeTestCases(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error)
	DescribeTestCasesAsync(ctx workflow.Context, input *codebuild.DescribeTestCasesInput) *CodeBuildDescribeTestCasesFuture

	GetResourcePolicy(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *codebuild.GetResourcePolicyInput) *CodeBuildGetResourcePolicyFuture

	ImportSourceCredentials(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error)
	ImportSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ImportSourceCredentialsInput) *CodeBuildImportSourceCredentialsFuture

	InvalidateProjectCache(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error)
	InvalidateProjectCacheAsync(ctx workflow.Context, input *codebuild.InvalidateProjectCacheInput) *CodeBuildInvalidateProjectCacheFuture

	ListBuildBatches(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error)
	ListBuildBatchesAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesInput) *CodeBuildListBuildBatchesFuture

	ListBuildBatchesForProject(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error)
	ListBuildBatchesForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildBatchesForProjectInput) *CodeBuildListBuildBatchesForProjectFuture

	ListBuilds(ctx workflow.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error)
	ListBuildsAsync(ctx workflow.Context, input *codebuild.ListBuildsInput) *CodeBuildListBuildsFuture

	ListBuildsForProject(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectAsync(ctx workflow.Context, input *codebuild.ListBuildsForProjectInput) *CodeBuildListBuildsForProjectFuture

	ListCuratedEnvironmentImages(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesAsync(ctx workflow.Context, input *codebuild.ListCuratedEnvironmentImagesInput) *CodeBuildListCuratedEnvironmentImagesFuture

	ListProjects(ctx workflow.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *codebuild.ListProjectsInput) *CodeBuildListProjectsFuture

	ListReportGroups(ctx workflow.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error)
	ListReportGroupsAsync(ctx workflow.Context, input *codebuild.ListReportGroupsInput) *CodeBuildListReportGroupsFuture

	ListReports(ctx workflow.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error)
	ListReportsAsync(ctx workflow.Context, input *codebuild.ListReportsInput) *CodeBuildListReportsFuture

	ListReportsForReportGroup(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error)
	ListReportsForReportGroupAsync(ctx workflow.Context, input *codebuild.ListReportsForReportGroupInput) *CodeBuildListReportsForReportGroupFuture

	ListSharedProjects(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error)
	ListSharedProjectsAsync(ctx workflow.Context, input *codebuild.ListSharedProjectsInput) *CodeBuildListSharedProjectsFuture

	ListSharedReportGroups(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error)
	ListSharedReportGroupsAsync(ctx workflow.Context, input *codebuild.ListSharedReportGroupsInput) *CodeBuildListSharedReportGroupsFuture

	ListSourceCredentials(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error)
	ListSourceCredentialsAsync(ctx workflow.Context, input *codebuild.ListSourceCredentialsInput) *CodeBuildListSourceCredentialsFuture

	PutResourcePolicy(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *codebuild.PutResourcePolicyInput) *CodeBuildPutResourcePolicyFuture

	RetryBuild(ctx workflow.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error)
	RetryBuildAsync(ctx workflow.Context, input *codebuild.RetryBuildInput) *CodeBuildRetryBuildFuture

	RetryBuildBatch(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error)
	RetryBuildBatchAsync(ctx workflow.Context, input *codebuild.RetryBuildBatchInput) *CodeBuildRetryBuildBatchFuture

	StartBuild(ctx workflow.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error)
	StartBuildAsync(ctx workflow.Context, input *codebuild.StartBuildInput) *CodeBuildStartBuildFuture

	StartBuildBatch(ctx workflow.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error)
	StartBuildBatchAsync(ctx workflow.Context, input *codebuild.StartBuildBatchInput) *CodeBuildStartBuildBatchFuture

	StopBuild(ctx workflow.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error)
	StopBuildAsync(ctx workflow.Context, input *codebuild.StopBuildInput) *CodeBuildStopBuildFuture

	StopBuildBatch(ctx workflow.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error)
	StopBuildBatchAsync(ctx workflow.Context, input *codebuild.StopBuildBatchInput) *CodeBuildStopBuildBatchFuture

	UpdateProject(ctx workflow.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *codebuild.UpdateProjectInput) *CodeBuildUpdateProjectFuture

	UpdateReportGroup(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error)
	UpdateReportGroupAsync(ctx workflow.Context, input *codebuild.UpdateReportGroupInput) *CodeBuildUpdateReportGroupFuture

	UpdateWebhook(ctx workflow.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error)
	UpdateWebhookAsync(ctx workflow.Context, input *codebuild.UpdateWebhookInput) *CodeBuildUpdateWebhookFuture
}

func NewClient() Client {
	return &stub{}
}

