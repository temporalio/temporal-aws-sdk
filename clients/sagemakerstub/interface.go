// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package sagemakerstub

import (
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTags(ctx workflow.Context, input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *sagemaker.AddTagsInput) *AddTagsFuture

	AssociateTrialComponent(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error)
	AssociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) *AssociateTrialComponentFuture

	CreateAlgorithm(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error)
	CreateAlgorithmAsync(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) *CreateAlgorithmFuture

	CreateApp(ctx workflow.Context, input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *sagemaker.CreateAppInput) *CreateAppFuture

	CreateAutoMLJob(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error)
	CreateAutoMLJobAsync(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) *CreateAutoMLJobFuture

	CreateCodeRepository(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error)
	CreateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) *CreateCodeRepositoryFuture

	CreateCompilationJob(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error)
	CreateCompilationJobAsync(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) *CreateCompilationJobFuture

	CreateDomain(ctx workflow.Context, input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *sagemaker.CreateDomainInput) *CreateDomainFuture

	CreateEndpoint(ctx workflow.Context, input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *sagemaker.CreateEndpointInput) *CreateEndpointFuture

	CreateEndpointConfig(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error)
	CreateEndpointConfigAsync(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) *CreateEndpointConfigFuture

	CreateExperiment(ctx workflow.Context, input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error)
	CreateExperimentAsync(ctx workflow.Context, input *sagemaker.CreateExperimentInput) *CreateExperimentFuture

	CreateFlowDefinition(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error)
	CreateFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) *CreateFlowDefinitionFuture

	CreateHumanTaskUi(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error)
	CreateHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) *CreateHumanTaskUiFuture

	CreateHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error)
	CreateHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) *CreateHyperParameterTuningJobFuture

	CreateLabelingJob(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error)
	CreateLabelingJobAsync(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) *CreateLabelingJobFuture

	CreateModel(ctx workflow.Context, input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *sagemaker.CreateModelInput) *CreateModelFuture

	CreateModelPackage(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error)
	CreateModelPackageAsync(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) *CreateModelPackageFuture

	CreateMonitoringSchedule(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error)
	CreateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) *CreateMonitoringScheduleFuture

	CreateNotebookInstance(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error)
	CreateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) *CreateNotebookInstanceFuture

	CreateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error)
	CreateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) *CreateNotebookInstanceLifecycleConfigFuture

	CreatePresignedDomainUrl(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error)
	CreatePresignedDomainUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) *CreatePresignedDomainUrlFuture

	CreatePresignedNotebookInstanceUrl(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error)
	CreatePresignedNotebookInstanceUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) *CreatePresignedNotebookInstanceUrlFuture

	CreateProcessingJob(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error)
	CreateProcessingJobAsync(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) *CreateProcessingJobFuture

	CreateTrainingJob(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error)
	CreateTrainingJobAsync(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) *CreateTrainingJobFuture

	CreateTransformJob(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error)
	CreateTransformJobAsync(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) *CreateTransformJobFuture

	CreateTrial(ctx workflow.Context, input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error)
	CreateTrialAsync(ctx workflow.Context, input *sagemaker.CreateTrialInput) *CreateTrialFuture

	CreateTrialComponent(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error)
	CreateTrialComponentAsync(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) *CreateTrialComponentFuture

	CreateUserProfile(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) *CreateUserProfileFuture

	CreateWorkforce(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error)
	CreateWorkforceAsync(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) *CreateWorkforceFuture

	CreateWorkteam(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error)
	CreateWorkteamAsync(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) *CreateWorkteamFuture

	DeleteAlgorithm(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error)
	DeleteAlgorithmAsync(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) *DeleteAlgorithmFuture

	DeleteApp(ctx workflow.Context, input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *sagemaker.DeleteAppInput) *DeleteAppFuture

	DeleteCodeRepository(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error)
	DeleteCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) *DeleteCodeRepositoryFuture

	DeleteDomain(ctx workflow.Context, input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *sagemaker.DeleteDomainInput) *DeleteDomainFuture

	DeleteEndpoint(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) *DeleteEndpointFuture

	DeleteEndpointConfig(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error)
	DeleteEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) *DeleteEndpointConfigFuture

	DeleteExperiment(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error)
	DeleteExperimentAsync(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) *DeleteExperimentFuture

	DeleteFlowDefinition(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error)
	DeleteFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) *DeleteFlowDefinitionFuture

	DeleteHumanTaskUi(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error)
	DeleteHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) *DeleteHumanTaskUiFuture

	DeleteModel(ctx workflow.Context, input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *sagemaker.DeleteModelInput) *DeleteModelFuture

	DeleteModelPackage(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error)
	DeleteModelPackageAsync(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) *DeleteModelPackageFuture

	DeleteMonitoringSchedule(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error)
	DeleteMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) *DeleteMonitoringScheduleFuture

	DeleteNotebookInstance(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error)
	DeleteNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) *DeleteNotebookInstanceFuture

	DeleteNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error)
	DeleteNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) *DeleteNotebookInstanceLifecycleConfigFuture

	DeleteTags(ctx workflow.Context, input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *sagemaker.DeleteTagsInput) *DeleteTagsFuture

	DeleteTrial(ctx workflow.Context, input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error)
	DeleteTrialAsync(ctx workflow.Context, input *sagemaker.DeleteTrialInput) *DeleteTrialFuture

	DeleteTrialComponent(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error)
	DeleteTrialComponentAsync(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) *DeleteTrialComponentFuture

	DeleteUserProfile(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) *DeleteUserProfileFuture

	DeleteWorkforce(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error)
	DeleteWorkforceAsync(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) *DeleteWorkforceFuture

	DeleteWorkteam(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error)
	DeleteWorkteamAsync(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) *DeleteWorkteamFuture

	DescribeAlgorithm(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error)
	DescribeAlgorithmAsync(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) *DescribeAlgorithmFuture

	DescribeApp(ctx workflow.Context, input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error)
	DescribeAppAsync(ctx workflow.Context, input *sagemaker.DescribeAppInput) *DescribeAppFuture

	DescribeAutoMLJob(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error)
	DescribeAutoMLJobAsync(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) *DescribeAutoMLJobFuture

	DescribeCodeRepository(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error)
	DescribeCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) *DescribeCodeRepositoryFuture

	DescribeCompilationJob(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error)
	DescribeCompilationJobAsync(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) *DescribeCompilationJobFuture

	DescribeDomain(ctx workflow.Context, input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error)
	DescribeDomainAsync(ctx workflow.Context, input *sagemaker.DescribeDomainInput) *DescribeDomainFuture

	DescribeEndpoint(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error)
	DescribeEndpointAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *DescribeEndpointFuture

	DescribeEndpointConfig(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error)
	DescribeEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) *DescribeEndpointConfigFuture

	DescribeExperiment(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error)
	DescribeExperimentAsync(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) *DescribeExperimentFuture

	DescribeFlowDefinition(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error)
	DescribeFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) *DescribeFlowDefinitionFuture

	DescribeHumanTaskUi(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error)
	DescribeHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) *DescribeHumanTaskUiFuture

	DescribeHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error)
	DescribeHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) *DescribeHyperParameterTuningJobFuture

	DescribeLabelingJob(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error)
	DescribeLabelingJobAsync(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) *DescribeLabelingJobFuture

	DescribeModel(ctx workflow.Context, input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error)
	DescribeModelAsync(ctx workflow.Context, input *sagemaker.DescribeModelInput) *DescribeModelFuture

	DescribeModelPackage(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error)
	DescribeModelPackageAsync(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) *DescribeModelPackageFuture

	DescribeMonitoringSchedule(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error)
	DescribeMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) *DescribeMonitoringScheduleFuture

	DescribeNotebookInstance(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error)
	DescribeNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *DescribeNotebookInstanceFuture

	DescribeNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error)
	DescribeNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) *DescribeNotebookInstanceLifecycleConfigFuture

	DescribeProcessingJob(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error)
	DescribeProcessingJobAsync(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) *DescribeProcessingJobFuture

	DescribeSubscribedWorkteam(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error)
	DescribeSubscribedWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) *DescribeSubscribedWorkteamFuture

	DescribeTrainingJob(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error)
	DescribeTrainingJobAsync(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) *DescribeTrainingJobFuture

	DescribeTransformJob(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error)
	DescribeTransformJobAsync(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) *DescribeTransformJobFuture

	DescribeTrial(ctx workflow.Context, input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error)
	DescribeTrialAsync(ctx workflow.Context, input *sagemaker.DescribeTrialInput) *DescribeTrialFuture

	DescribeTrialComponent(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error)
	DescribeTrialComponentAsync(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) *DescribeTrialComponentFuture

	DescribeUserProfile(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error)
	DescribeUserProfileAsync(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) *DescribeUserProfileFuture

	DescribeWorkforce(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error)
	DescribeWorkforceAsync(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) *DescribeWorkforceFuture

	DescribeWorkteam(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error)
	DescribeWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) *DescribeWorkteamFuture

	DisassociateTrialComponent(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error)
	DisassociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) *DisassociateTrialComponentFuture

	GetSearchSuggestions(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error)
	GetSearchSuggestionsAsync(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) *GetSearchSuggestionsFuture

	ListAlgorithms(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error)
	ListAlgorithmsAsync(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) *ListAlgorithmsFuture

	ListApps(ctx workflow.Context, input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error)
	ListAppsAsync(ctx workflow.Context, input *sagemaker.ListAppsInput) *ListAppsFuture

	ListAutoMLJobs(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error)
	ListAutoMLJobsAsync(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) *ListAutoMLJobsFuture

	ListCandidatesForAutoMLJob(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error)
	ListCandidatesForAutoMLJobAsync(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) *ListCandidatesForAutoMLJobFuture

	ListCodeRepositories(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error)
	ListCodeRepositoriesAsync(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) *ListCodeRepositoriesFuture

	ListCompilationJobs(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error)
	ListCompilationJobsAsync(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) *ListCompilationJobsFuture

	ListDomains(ctx workflow.Context, input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *sagemaker.ListDomainsInput) *ListDomainsFuture

	ListEndpointConfigs(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error)
	ListEndpointConfigsAsync(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) *ListEndpointConfigsFuture

	ListEndpoints(ctx workflow.Context, input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *sagemaker.ListEndpointsInput) *ListEndpointsFuture

	ListExperiments(ctx workflow.Context, input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error)
	ListExperimentsAsync(ctx workflow.Context, input *sagemaker.ListExperimentsInput) *ListExperimentsFuture

	ListFlowDefinitions(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error)
	ListFlowDefinitionsAsync(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) *ListFlowDefinitionsFuture

	ListHumanTaskUis(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error)
	ListHumanTaskUisAsync(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) *ListHumanTaskUisFuture

	ListHyperParameterTuningJobs(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error)
	ListHyperParameterTuningJobsAsync(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) *ListHyperParameterTuningJobsFuture

	ListLabelingJobs(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error)
	ListLabelingJobsAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) *ListLabelingJobsFuture

	ListLabelingJobsForWorkteam(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error)
	ListLabelingJobsForWorkteamAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) *ListLabelingJobsForWorkteamFuture

	ListModelPackages(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error)
	ListModelPackagesAsync(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) *ListModelPackagesFuture

	ListModels(ctx workflow.Context, input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error)
	ListModelsAsync(ctx workflow.Context, input *sagemaker.ListModelsInput) *ListModelsFuture

	ListMonitoringExecutions(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error)
	ListMonitoringExecutionsAsync(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) *ListMonitoringExecutionsFuture

	ListMonitoringSchedules(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error)
	ListMonitoringSchedulesAsync(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) *ListMonitoringSchedulesFuture

	ListNotebookInstanceLifecycleConfigs(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error)
	ListNotebookInstanceLifecycleConfigsAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) *ListNotebookInstanceLifecycleConfigsFuture

	ListNotebookInstances(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error)
	ListNotebookInstancesAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) *ListNotebookInstancesFuture

	ListProcessingJobs(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error)
	ListProcessingJobsAsync(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) *ListProcessingJobsFuture

	ListSubscribedWorkteams(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error)
	ListSubscribedWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) *ListSubscribedWorkteamsFuture

	ListTags(ctx workflow.Context, input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *sagemaker.ListTagsInput) *ListTagsFuture

	ListTrainingJobs(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error)
	ListTrainingJobsAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) *ListTrainingJobsFuture

	ListTrainingJobsForHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error)
	ListTrainingJobsForHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) *ListTrainingJobsForHyperParameterTuningJobFuture

	ListTransformJobs(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error)
	ListTransformJobsAsync(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) *ListTransformJobsFuture

	ListTrialComponents(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error)
	ListTrialComponentsAsync(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) *ListTrialComponentsFuture

	ListTrials(ctx workflow.Context, input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error)
	ListTrialsAsync(ctx workflow.Context, input *sagemaker.ListTrialsInput) *ListTrialsFuture

	ListUserProfiles(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error)
	ListUserProfilesAsync(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) *ListUserProfilesFuture

	ListWorkforces(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error)
	ListWorkforcesAsync(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) *ListWorkforcesFuture

	ListWorkteams(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error)
	ListWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) *ListWorkteamsFuture

	RenderUiTemplate(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error)
	RenderUiTemplateAsync(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) *RenderUiTemplateFuture

	Search(ctx workflow.Context, input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error)
	SearchAsync(ctx workflow.Context, input *sagemaker.SearchInput) *SearchFuture

	StartMonitoringSchedule(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error)
	StartMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) *StartMonitoringScheduleFuture

	StartNotebookInstance(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error)
	StartNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) *StartNotebookInstanceFuture

	StopAutoMLJob(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error)
	StopAutoMLJobAsync(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) *StopAutoMLJobFuture

	StopCompilationJob(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error)
	StopCompilationJobAsync(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) *StopCompilationJobFuture

	StopHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error)
	StopHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) *StopHyperParameterTuningJobFuture

	StopLabelingJob(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error)
	StopLabelingJobAsync(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) *StopLabelingJobFuture

	StopMonitoringSchedule(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error)
	StopMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) *StopMonitoringScheduleFuture

	StopNotebookInstance(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error)
	StopNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) *StopNotebookInstanceFuture

	StopProcessingJob(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error)
	StopProcessingJobAsync(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) *StopProcessingJobFuture

	StopTrainingJob(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error)
	StopTrainingJobAsync(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) *StopTrainingJobFuture

	StopTransformJob(ctx workflow.Context, input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error)
	StopTransformJobAsync(ctx workflow.Context, input *sagemaker.StopTransformJobInput) *StopTransformJobFuture

	UpdateCodeRepository(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error)
	UpdateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) *UpdateCodeRepositoryFuture

	UpdateDomain(ctx workflow.Context, input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error)
	UpdateDomainAsync(ctx workflow.Context, input *sagemaker.UpdateDomainInput) *UpdateDomainFuture

	UpdateEndpoint(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error)
	UpdateEndpointAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) *UpdateEndpointFuture

	UpdateEndpointWeightsAndCapacities(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error)
	UpdateEndpointWeightsAndCapacitiesAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) *UpdateEndpointWeightsAndCapacitiesFuture

	UpdateExperiment(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error)
	UpdateExperimentAsync(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) *UpdateExperimentFuture

	UpdateMonitoringSchedule(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error)
	UpdateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) *UpdateMonitoringScheduleFuture

	UpdateNotebookInstance(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error)
	UpdateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) *UpdateNotebookInstanceFuture

	UpdateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error)
	UpdateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) *UpdateNotebookInstanceLifecycleConfigFuture

	UpdateTrial(ctx workflow.Context, input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error)
	UpdateTrialAsync(ctx workflow.Context, input *sagemaker.UpdateTrialInput) *UpdateTrialFuture

	UpdateTrialComponent(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error)
	UpdateTrialComponentAsync(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) *UpdateTrialComponentFuture

	UpdateUserProfile(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) *UpdateUserProfileFuture

	UpdateWorkforce(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error)
	UpdateWorkforceAsync(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) *UpdateWorkforceFuture

	UpdateWorkteam(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error)
	UpdateWorkteamAsync(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) *UpdateWorkteamFuture

	WaitUntilEndpointDeleted(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
	WaitUntilEndpointDeletedAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *clients.VoidFuture

	WaitUntilEndpointInService(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
	WaitUntilEndpointInServiceAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *clients.VoidFuture

	WaitUntilNotebookInstanceDeleted(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceDeletedAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *clients.VoidFuture

	WaitUntilNotebookInstanceInService(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceInServiceAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *clients.VoidFuture

	WaitUntilNotebookInstanceStopped(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *clients.VoidFuture

	WaitUntilProcessingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) error
	WaitUntilProcessingJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) *clients.VoidFuture

	WaitUntilTrainingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) error
	WaitUntilTrainingJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) *clients.VoidFuture

	WaitUntilTransformJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) error
	WaitUntilTransformJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
