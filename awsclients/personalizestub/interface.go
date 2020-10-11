// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package personalizestub

import (
	"github.com/aws/aws-sdk-go/service/personalize"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateBatchInferenceJob(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error)
	CreateBatchInferenceJobAsync(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) *PersonalizeCreateBatchInferenceJobFuture

	CreateCampaign(ctx workflow.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error)
	CreateCampaignAsync(ctx workflow.Context, input *personalize.CreateCampaignInput) *PersonalizeCreateCampaignFuture

	CreateDataset(ctx workflow.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error)
	CreateDatasetAsync(ctx workflow.Context, input *personalize.CreateDatasetInput) *PersonalizeCreateDatasetFuture

	CreateDatasetGroup(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error)
	CreateDatasetGroupAsync(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) *PersonalizeCreateDatasetGroupFuture

	CreateDatasetImportJob(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error)
	CreateDatasetImportJobAsync(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) *PersonalizeCreateDatasetImportJobFuture

	CreateEventTracker(ctx workflow.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error)
	CreateEventTrackerAsync(ctx workflow.Context, input *personalize.CreateEventTrackerInput) *PersonalizeCreateEventTrackerFuture

	CreateFilter(ctx workflow.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error)
	CreateFilterAsync(ctx workflow.Context, input *personalize.CreateFilterInput) *PersonalizeCreateFilterFuture

	CreateSchema(ctx workflow.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error)
	CreateSchemaAsync(ctx workflow.Context, input *personalize.CreateSchemaInput) *PersonalizeCreateSchemaFuture

	CreateSolution(ctx workflow.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error)
	CreateSolutionAsync(ctx workflow.Context, input *personalize.CreateSolutionInput) *PersonalizeCreateSolutionFuture

	CreateSolutionVersion(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error)
	CreateSolutionVersionAsync(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) *PersonalizeCreateSolutionVersionFuture

	DeleteCampaign(ctx workflow.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error)
	DeleteCampaignAsync(ctx workflow.Context, input *personalize.DeleteCampaignInput) *PersonalizeDeleteCampaignFuture

	DeleteDataset(ctx workflow.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error)
	DeleteDatasetAsync(ctx workflow.Context, input *personalize.DeleteDatasetInput) *PersonalizeDeleteDatasetFuture

	DeleteDatasetGroup(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error)
	DeleteDatasetGroupAsync(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) *PersonalizeDeleteDatasetGroupFuture

	DeleteEventTracker(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error)
	DeleteEventTrackerAsync(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) *PersonalizeDeleteEventTrackerFuture

	DeleteFilter(ctx workflow.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error)
	DeleteFilterAsync(ctx workflow.Context, input *personalize.DeleteFilterInput) *PersonalizeDeleteFilterFuture

	DeleteSchema(ctx workflow.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error)
	DeleteSchemaAsync(ctx workflow.Context, input *personalize.DeleteSchemaInput) *PersonalizeDeleteSchemaFuture

	DeleteSolution(ctx workflow.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error)
	DeleteSolutionAsync(ctx workflow.Context, input *personalize.DeleteSolutionInput) *PersonalizeDeleteSolutionFuture

	DescribeAlgorithm(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error)
	DescribeAlgorithmAsync(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) *PersonalizeDescribeAlgorithmFuture

	DescribeBatchInferenceJob(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error)
	DescribeBatchInferenceJobAsync(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) *PersonalizeDescribeBatchInferenceJobFuture

	DescribeCampaign(ctx workflow.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error)
	DescribeCampaignAsync(ctx workflow.Context, input *personalize.DescribeCampaignInput) *PersonalizeDescribeCampaignFuture

	DescribeDataset(ctx workflow.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error)
	DescribeDatasetAsync(ctx workflow.Context, input *personalize.DescribeDatasetInput) *PersonalizeDescribeDatasetFuture

	DescribeDatasetGroup(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error)
	DescribeDatasetGroupAsync(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) *PersonalizeDescribeDatasetGroupFuture

	DescribeDatasetImportJob(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error)
	DescribeDatasetImportJobAsync(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) *PersonalizeDescribeDatasetImportJobFuture

	DescribeEventTracker(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error)
	DescribeEventTrackerAsync(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) *PersonalizeDescribeEventTrackerFuture

	DescribeFeatureTransformation(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error)
	DescribeFeatureTransformationAsync(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) *PersonalizeDescribeFeatureTransformationFuture

	DescribeFilter(ctx workflow.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error)
	DescribeFilterAsync(ctx workflow.Context, input *personalize.DescribeFilterInput) *PersonalizeDescribeFilterFuture

	DescribeRecipe(ctx workflow.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error)
	DescribeRecipeAsync(ctx workflow.Context, input *personalize.DescribeRecipeInput) *PersonalizeDescribeRecipeFuture

	DescribeSchema(ctx workflow.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error)
	DescribeSchemaAsync(ctx workflow.Context, input *personalize.DescribeSchemaInput) *PersonalizeDescribeSchemaFuture

	DescribeSolution(ctx workflow.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error)
	DescribeSolutionAsync(ctx workflow.Context, input *personalize.DescribeSolutionInput) *PersonalizeDescribeSolutionFuture

	DescribeSolutionVersion(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error)
	DescribeSolutionVersionAsync(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) *PersonalizeDescribeSolutionVersionFuture

	GetSolutionMetrics(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error)
	GetSolutionMetricsAsync(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) *PersonalizeGetSolutionMetricsFuture

	ListBatchInferenceJobs(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error)
	ListBatchInferenceJobsAsync(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) *PersonalizeListBatchInferenceJobsFuture

	ListCampaigns(ctx workflow.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error)
	ListCampaignsAsync(ctx workflow.Context, input *personalize.ListCampaignsInput) *PersonalizeListCampaignsFuture

	ListDatasetGroups(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error)
	ListDatasetGroupsAsync(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) *PersonalizeListDatasetGroupsFuture

	ListDatasetImportJobs(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error)
	ListDatasetImportJobsAsync(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) *PersonalizeListDatasetImportJobsFuture

	ListDatasets(ctx workflow.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error)
	ListDatasetsAsync(ctx workflow.Context, input *personalize.ListDatasetsInput) *PersonalizeListDatasetsFuture

	ListEventTrackers(ctx workflow.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error)
	ListEventTrackersAsync(ctx workflow.Context, input *personalize.ListEventTrackersInput) *PersonalizeListEventTrackersFuture

	ListFilters(ctx workflow.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error)
	ListFiltersAsync(ctx workflow.Context, input *personalize.ListFiltersInput) *PersonalizeListFiltersFuture

	ListRecipes(ctx workflow.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error)
	ListRecipesAsync(ctx workflow.Context, input *personalize.ListRecipesInput) *PersonalizeListRecipesFuture

	ListSchemas(ctx workflow.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *personalize.ListSchemasInput) *PersonalizeListSchemasFuture

	ListSolutionVersions(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error)
	ListSolutionVersionsAsync(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) *PersonalizeListSolutionVersionsFuture

	ListSolutions(ctx workflow.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error)
	ListSolutionsAsync(ctx workflow.Context, input *personalize.ListSolutionsInput) *PersonalizeListSolutionsFuture

	UpdateCampaign(ctx workflow.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error)
	UpdateCampaignAsync(ctx workflow.Context, input *personalize.UpdateCampaignInput) *PersonalizeUpdateCampaignFuture
}

func NewClient() Client {
	return &stub{}
}

