// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package gluestub

import (
	"github.com/aws/aws-sdk-go/service/glue"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchCreatePartition(ctx workflow.Context, input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error)
	BatchCreatePartitionAsync(ctx workflow.Context, input *glue.BatchCreatePartitionInput) *GlueBatchCreatePartitionFuture

	BatchDeleteConnection(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error)
	BatchDeleteConnectionAsync(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) *GlueBatchDeleteConnectionFuture

	BatchDeletePartition(ctx workflow.Context, input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error)
	BatchDeletePartitionAsync(ctx workflow.Context, input *glue.BatchDeletePartitionInput) *GlueBatchDeletePartitionFuture

	BatchDeleteTable(ctx workflow.Context, input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error)
	BatchDeleteTableAsync(ctx workflow.Context, input *glue.BatchDeleteTableInput) *GlueBatchDeleteTableFuture

	BatchDeleteTableVersion(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error)
	BatchDeleteTableVersionAsync(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) *GlueBatchDeleteTableVersionFuture

	BatchGetCrawlers(ctx workflow.Context, input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error)
	BatchGetCrawlersAsync(ctx workflow.Context, input *glue.BatchGetCrawlersInput) *GlueBatchGetCrawlersFuture

	BatchGetDevEndpoints(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error)
	BatchGetDevEndpointsAsync(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) *GlueBatchGetDevEndpointsFuture

	BatchGetJobs(ctx workflow.Context, input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error)
	BatchGetJobsAsync(ctx workflow.Context, input *glue.BatchGetJobsInput) *GlueBatchGetJobsFuture

	BatchGetPartition(ctx workflow.Context, input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error)
	BatchGetPartitionAsync(ctx workflow.Context, input *glue.BatchGetPartitionInput) *GlueBatchGetPartitionFuture

	BatchGetTriggers(ctx workflow.Context, input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error)
	BatchGetTriggersAsync(ctx workflow.Context, input *glue.BatchGetTriggersInput) *GlueBatchGetTriggersFuture

	BatchGetWorkflows(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error)
	BatchGetWorkflowsAsync(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) *GlueBatchGetWorkflowsFuture

	BatchStopJobRun(ctx workflow.Context, input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error)
	BatchStopJobRunAsync(ctx workflow.Context, input *glue.BatchStopJobRunInput) *GlueBatchStopJobRunFuture

	BatchUpdatePartition(ctx workflow.Context, input *glue.BatchUpdatePartitionInput) (*glue.BatchUpdatePartitionOutput, error)
	BatchUpdatePartitionAsync(ctx workflow.Context, input *glue.BatchUpdatePartitionInput) *GlueBatchUpdatePartitionFuture

	CancelMLTaskRun(ctx workflow.Context, input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error)
	CancelMLTaskRunAsync(ctx workflow.Context, input *glue.CancelMLTaskRunInput) *GlueCancelMLTaskRunFuture

	CreateClassifier(ctx workflow.Context, input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error)
	CreateClassifierAsync(ctx workflow.Context, input *glue.CreateClassifierInput) *GlueCreateClassifierFuture

	CreateConnection(ctx workflow.Context, input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error)
	CreateConnectionAsync(ctx workflow.Context, input *glue.CreateConnectionInput) *GlueCreateConnectionFuture

	CreateCrawler(ctx workflow.Context, input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error)
	CreateCrawlerAsync(ctx workflow.Context, input *glue.CreateCrawlerInput) *GlueCreateCrawlerFuture

	CreateDatabase(ctx workflow.Context, input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error)
	CreateDatabaseAsync(ctx workflow.Context, input *glue.CreateDatabaseInput) *GlueCreateDatabaseFuture

	CreateDevEndpoint(ctx workflow.Context, input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error)
	CreateDevEndpointAsync(ctx workflow.Context, input *glue.CreateDevEndpointInput) *GlueCreateDevEndpointFuture

	CreateJob(ctx workflow.Context, input *glue.CreateJobInput) (*glue.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *glue.CreateJobInput) *GlueCreateJobFuture

	CreateMLTransform(ctx workflow.Context, input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error)
	CreateMLTransformAsync(ctx workflow.Context, input *glue.CreateMLTransformInput) *GlueCreateMLTransformFuture

	CreatePartition(ctx workflow.Context, input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error)
	CreatePartitionAsync(ctx workflow.Context, input *glue.CreatePartitionInput) *GlueCreatePartitionFuture

	CreateScript(ctx workflow.Context, input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error)
	CreateScriptAsync(ctx workflow.Context, input *glue.CreateScriptInput) *GlueCreateScriptFuture

	CreateSecurityConfiguration(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error)
	CreateSecurityConfigurationAsync(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) *GlueCreateSecurityConfigurationFuture

	CreateTable(ctx workflow.Context, input *glue.CreateTableInput) (*glue.CreateTableOutput, error)
	CreateTableAsync(ctx workflow.Context, input *glue.CreateTableInput) *GlueCreateTableFuture

	CreateTrigger(ctx workflow.Context, input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error)
	CreateTriggerAsync(ctx workflow.Context, input *glue.CreateTriggerInput) *GlueCreateTriggerFuture

	CreateUserDefinedFunction(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error)
	CreateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) *GlueCreateUserDefinedFunctionFuture

	CreateWorkflow(ctx workflow.Context, input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error)
	CreateWorkflowAsync(ctx workflow.Context, input *glue.CreateWorkflowInput) *GlueCreateWorkflowFuture

	DeleteClassifier(ctx workflow.Context, input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error)
	DeleteClassifierAsync(ctx workflow.Context, input *glue.DeleteClassifierInput) *GlueDeleteClassifierFuture

	DeleteColumnStatisticsForPartition(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error)
	DeleteColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) *GlueDeleteColumnStatisticsForPartitionFuture

	DeleteColumnStatisticsForTable(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error)
	DeleteColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) *GlueDeleteColumnStatisticsForTableFuture

	DeleteConnection(ctx workflow.Context, input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *glue.DeleteConnectionInput) *GlueDeleteConnectionFuture

	DeleteCrawler(ctx workflow.Context, input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error)
	DeleteCrawlerAsync(ctx workflow.Context, input *glue.DeleteCrawlerInput) *GlueDeleteCrawlerFuture

	DeleteDatabase(ctx workflow.Context, input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error)
	DeleteDatabaseAsync(ctx workflow.Context, input *glue.DeleteDatabaseInput) *GlueDeleteDatabaseFuture

	DeleteDevEndpoint(ctx workflow.Context, input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error)
	DeleteDevEndpointAsync(ctx workflow.Context, input *glue.DeleteDevEndpointInput) *GlueDeleteDevEndpointFuture

	DeleteJob(ctx workflow.Context, input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error)
	DeleteJobAsync(ctx workflow.Context, input *glue.DeleteJobInput) *GlueDeleteJobFuture

	DeleteMLTransform(ctx workflow.Context, input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error)
	DeleteMLTransformAsync(ctx workflow.Context, input *glue.DeleteMLTransformInput) *GlueDeleteMLTransformFuture

	DeletePartition(ctx workflow.Context, input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error)
	DeletePartitionAsync(ctx workflow.Context, input *glue.DeletePartitionInput) *GlueDeletePartitionFuture

	DeleteResourcePolicy(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) *GlueDeleteResourcePolicyFuture

	DeleteSecurityConfiguration(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error)
	DeleteSecurityConfigurationAsync(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) *GlueDeleteSecurityConfigurationFuture

	DeleteTable(ctx workflow.Context, input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error)
	DeleteTableAsync(ctx workflow.Context, input *glue.DeleteTableInput) *GlueDeleteTableFuture

	DeleteTableVersion(ctx workflow.Context, input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error)
	DeleteTableVersionAsync(ctx workflow.Context, input *glue.DeleteTableVersionInput) *GlueDeleteTableVersionFuture

	DeleteTrigger(ctx workflow.Context, input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error)
	DeleteTriggerAsync(ctx workflow.Context, input *glue.DeleteTriggerInput) *GlueDeleteTriggerFuture

	DeleteUserDefinedFunction(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error)
	DeleteUserDefinedFunctionAsync(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) *GlueDeleteUserDefinedFunctionFuture

	DeleteWorkflow(ctx workflow.Context, input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error)
	DeleteWorkflowAsync(ctx workflow.Context, input *glue.DeleteWorkflowInput) *GlueDeleteWorkflowFuture

	GetCatalogImportStatus(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error)
	GetCatalogImportStatusAsync(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) *GlueGetCatalogImportStatusFuture

	GetClassifier(ctx workflow.Context, input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error)
	GetClassifierAsync(ctx workflow.Context, input *glue.GetClassifierInput) *GlueGetClassifierFuture

	GetClassifiers(ctx workflow.Context, input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error)
	GetClassifiersAsync(ctx workflow.Context, input *glue.GetClassifiersInput) *GlueGetClassifiersFuture

	GetColumnStatisticsForPartition(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error)
	GetColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) *GlueGetColumnStatisticsForPartitionFuture

	GetColumnStatisticsForTable(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error)
	GetColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) *GlueGetColumnStatisticsForTableFuture

	GetConnection(ctx workflow.Context, input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error)
	GetConnectionAsync(ctx workflow.Context, input *glue.GetConnectionInput) *GlueGetConnectionFuture

	GetConnections(ctx workflow.Context, input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error)
	GetConnectionsAsync(ctx workflow.Context, input *glue.GetConnectionsInput) *GlueGetConnectionsFuture

	GetCrawler(ctx workflow.Context, input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error)
	GetCrawlerAsync(ctx workflow.Context, input *glue.GetCrawlerInput) *GlueGetCrawlerFuture

	GetCrawlerMetrics(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlerMetricsAsync(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) *GlueGetCrawlerMetricsFuture

	GetCrawlers(ctx workflow.Context, input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error)
	GetCrawlersAsync(ctx workflow.Context, input *glue.GetCrawlersInput) *GlueGetCrawlersFuture

	GetDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
	GetDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) *GlueGetDataCatalogEncryptionSettingsFuture

	GetDatabase(ctx workflow.Context, input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error)
	GetDatabaseAsync(ctx workflow.Context, input *glue.GetDatabaseInput) *GlueGetDatabaseFuture

	GetDatabases(ctx workflow.Context, input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error)
	GetDatabasesAsync(ctx workflow.Context, input *glue.GetDatabasesInput) *GlueGetDatabasesFuture

	GetDataflowGraph(ctx workflow.Context, input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error)
	GetDataflowGraphAsync(ctx workflow.Context, input *glue.GetDataflowGraphInput) *GlueGetDataflowGraphFuture

	GetDevEndpoint(ctx workflow.Context, input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error)
	GetDevEndpointAsync(ctx workflow.Context, input *glue.GetDevEndpointInput) *GlueGetDevEndpointFuture

	GetDevEndpoints(ctx workflow.Context, input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error)
	GetDevEndpointsAsync(ctx workflow.Context, input *glue.GetDevEndpointsInput) *GlueGetDevEndpointsFuture

	GetJob(ctx workflow.Context, input *glue.GetJobInput) (*glue.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *glue.GetJobInput) *GlueGetJobFuture

	GetJobBookmark(ctx workflow.Context, input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error)
	GetJobBookmarkAsync(ctx workflow.Context, input *glue.GetJobBookmarkInput) *GlueGetJobBookmarkFuture

	GetJobRun(ctx workflow.Context, input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error)
	GetJobRunAsync(ctx workflow.Context, input *glue.GetJobRunInput) *GlueGetJobRunFuture

	GetJobRuns(ctx workflow.Context, input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error)
	GetJobRunsAsync(ctx workflow.Context, input *glue.GetJobRunsInput) *GlueGetJobRunsFuture

	GetJobs(ctx workflow.Context, input *glue.GetJobsInput) (*glue.GetJobsOutput, error)
	GetJobsAsync(ctx workflow.Context, input *glue.GetJobsInput) *GlueGetJobsFuture

	GetMLTaskRun(ctx workflow.Context, input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error)
	GetMLTaskRunAsync(ctx workflow.Context, input *glue.GetMLTaskRunInput) *GlueGetMLTaskRunFuture

	GetMLTaskRuns(ctx workflow.Context, input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error)
	GetMLTaskRunsAsync(ctx workflow.Context, input *glue.GetMLTaskRunsInput) *GlueGetMLTaskRunsFuture

	GetMLTransform(ctx workflow.Context, input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error)
	GetMLTransformAsync(ctx workflow.Context, input *glue.GetMLTransformInput) *GlueGetMLTransformFuture

	GetMLTransforms(ctx workflow.Context, input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error)
	GetMLTransformsAsync(ctx workflow.Context, input *glue.GetMLTransformsInput) *GlueGetMLTransformsFuture

	GetMapping(ctx workflow.Context, input *glue.GetMappingInput) (*glue.GetMappingOutput, error)
	GetMappingAsync(ctx workflow.Context, input *glue.GetMappingInput) *GlueGetMappingFuture

	GetPartition(ctx workflow.Context, input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error)
	GetPartitionAsync(ctx workflow.Context, input *glue.GetPartitionInput) *GlueGetPartitionFuture

	GetPartitionIndexes(ctx workflow.Context, input *glue.GetPartitionIndexesInput) (*glue.GetPartitionIndexesOutput, error)
	GetPartitionIndexesAsync(ctx workflow.Context, input *glue.GetPartitionIndexesInput) *GlueGetPartitionIndexesFuture

	GetPartitions(ctx workflow.Context, input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error)
	GetPartitionsAsync(ctx workflow.Context, input *glue.GetPartitionsInput) *GlueGetPartitionsFuture

	GetPlan(ctx workflow.Context, input *glue.GetPlanInput) (*glue.GetPlanOutput, error)
	GetPlanAsync(ctx workflow.Context, input *glue.GetPlanInput) *GlueGetPlanFuture

	GetResourcePolicies(ctx workflow.Context, input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error)
	GetResourcePoliciesAsync(ctx workflow.Context, input *glue.GetResourcePoliciesInput) *GlueGetResourcePoliciesFuture

	GetResourcePolicy(ctx workflow.Context, input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *glue.GetResourcePolicyInput) *GlueGetResourcePolicyFuture

	GetSecurityConfiguration(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error)
	GetSecurityConfigurationAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) *GlueGetSecurityConfigurationFuture

	GetSecurityConfigurations(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error)
	GetSecurityConfigurationsAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) *GlueGetSecurityConfigurationsFuture

	GetTable(ctx workflow.Context, input *glue.GetTableInput) (*glue.GetTableOutput, error)
	GetTableAsync(ctx workflow.Context, input *glue.GetTableInput) *GlueGetTableFuture

	GetTableVersion(ctx workflow.Context, input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error)
	GetTableVersionAsync(ctx workflow.Context, input *glue.GetTableVersionInput) *GlueGetTableVersionFuture

	GetTableVersions(ctx workflow.Context, input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error)
	GetTableVersionsAsync(ctx workflow.Context, input *glue.GetTableVersionsInput) *GlueGetTableVersionsFuture

	GetTables(ctx workflow.Context, input *glue.GetTablesInput) (*glue.GetTablesOutput, error)
	GetTablesAsync(ctx workflow.Context, input *glue.GetTablesInput) *GlueGetTablesFuture

	GetTags(ctx workflow.Context, input *glue.GetTagsInput) (*glue.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *glue.GetTagsInput) *GlueGetTagsFuture

	GetTrigger(ctx workflow.Context, input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error)
	GetTriggerAsync(ctx workflow.Context, input *glue.GetTriggerInput) *GlueGetTriggerFuture

	GetTriggers(ctx workflow.Context, input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error)
	GetTriggersAsync(ctx workflow.Context, input *glue.GetTriggersInput) *GlueGetTriggersFuture

	GetUserDefinedFunction(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctionAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) *GlueGetUserDefinedFunctionFuture

	GetUserDefinedFunctions(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error)
	GetUserDefinedFunctionsAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) *GlueGetUserDefinedFunctionsFuture

	GetWorkflow(ctx workflow.Context, input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error)
	GetWorkflowAsync(ctx workflow.Context, input *glue.GetWorkflowInput) *GlueGetWorkflowFuture

	GetWorkflowRun(ctx workflow.Context, input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error)
	GetWorkflowRunAsync(ctx workflow.Context, input *glue.GetWorkflowRunInput) *GlueGetWorkflowRunFuture

	GetWorkflowRunProperties(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error)
	GetWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) *GlueGetWorkflowRunPropertiesFuture

	GetWorkflowRuns(ctx workflow.Context, input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error)
	GetWorkflowRunsAsync(ctx workflow.Context, input *glue.GetWorkflowRunsInput) *GlueGetWorkflowRunsFuture

	ImportCatalogToGlue(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error)
	ImportCatalogToGlueAsync(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) *GlueImportCatalogToGlueFuture

	ListCrawlers(ctx workflow.Context, input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error)
	ListCrawlersAsync(ctx workflow.Context, input *glue.ListCrawlersInput) *GlueListCrawlersFuture

	ListDevEndpoints(ctx workflow.Context, input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error)
	ListDevEndpointsAsync(ctx workflow.Context, input *glue.ListDevEndpointsInput) *GlueListDevEndpointsFuture

	ListJobs(ctx workflow.Context, input *glue.ListJobsInput) (*glue.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *glue.ListJobsInput) *GlueListJobsFuture

	ListMLTransforms(ctx workflow.Context, input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error)
	ListMLTransformsAsync(ctx workflow.Context, input *glue.ListMLTransformsInput) *GlueListMLTransformsFuture

	ListTriggers(ctx workflow.Context, input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error)
	ListTriggersAsync(ctx workflow.Context, input *glue.ListTriggersInput) *GlueListTriggersFuture

	ListWorkflows(ctx workflow.Context, input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error)
	ListWorkflowsAsync(ctx workflow.Context, input *glue.ListWorkflowsInput) *GlueListWorkflowsFuture

	PutDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error)
	PutDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) *GluePutDataCatalogEncryptionSettingsFuture

	PutResourcePolicy(ctx workflow.Context, input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *glue.PutResourcePolicyInput) *GluePutResourcePolicyFuture

	PutWorkflowRunProperties(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error)
	PutWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) *GluePutWorkflowRunPropertiesFuture

	ResetJobBookmark(ctx workflow.Context, input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error)
	ResetJobBookmarkAsync(ctx workflow.Context, input *glue.ResetJobBookmarkInput) *GlueResetJobBookmarkFuture

	ResumeWorkflowRun(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error)
	ResumeWorkflowRunAsync(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) *GlueResumeWorkflowRunFuture

	SearchTables(ctx workflow.Context, input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error)
	SearchTablesAsync(ctx workflow.Context, input *glue.SearchTablesInput) *GlueSearchTablesFuture

	StartCrawler(ctx workflow.Context, input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error)
	StartCrawlerAsync(ctx workflow.Context, input *glue.StartCrawlerInput) *GlueStartCrawlerFuture

	StartCrawlerSchedule(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error)
	StartCrawlerScheduleAsync(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) *GlueStartCrawlerScheduleFuture

	StartExportLabelsTaskRun(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error)
	StartExportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) *GlueStartExportLabelsTaskRunFuture

	StartImportLabelsTaskRun(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error)
	StartImportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) *GlueStartImportLabelsTaskRunFuture

	StartJobRun(ctx workflow.Context, input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error)
	StartJobRunAsync(ctx workflow.Context, input *glue.StartJobRunInput) *GlueStartJobRunFuture

	StartMLEvaluationTaskRun(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error)
	StartMLEvaluationTaskRunAsync(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) *GlueStartMLEvaluationTaskRunFuture

	StartMLLabelingSetGenerationTaskRun(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error)
	StartMLLabelingSetGenerationTaskRunAsync(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) *GlueStartMLLabelingSetGenerationTaskRunFuture

	StartTrigger(ctx workflow.Context, input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error)
	StartTriggerAsync(ctx workflow.Context, input *glue.StartTriggerInput) *GlueStartTriggerFuture

	StartWorkflowRun(ctx workflow.Context, input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error)
	StartWorkflowRunAsync(ctx workflow.Context, input *glue.StartWorkflowRunInput) *GlueStartWorkflowRunFuture

	StopCrawler(ctx workflow.Context, input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error)
	StopCrawlerAsync(ctx workflow.Context, input *glue.StopCrawlerInput) *GlueStopCrawlerFuture

	StopCrawlerSchedule(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error)
	StopCrawlerScheduleAsync(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) *GlueStopCrawlerScheduleFuture

	StopTrigger(ctx workflow.Context, input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error)
	StopTriggerAsync(ctx workflow.Context, input *glue.StopTriggerInput) *GlueStopTriggerFuture

	StopWorkflowRun(ctx workflow.Context, input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error)
	StopWorkflowRunAsync(ctx workflow.Context, input *glue.StopWorkflowRunInput) *GlueStopWorkflowRunFuture

	TagResource(ctx workflow.Context, input *glue.TagResourceInput) (*glue.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *glue.TagResourceInput) *GlueTagResourceFuture

	UntagResource(ctx workflow.Context, input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *glue.UntagResourceInput) *GlueUntagResourceFuture

	UpdateClassifier(ctx workflow.Context, input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error)
	UpdateClassifierAsync(ctx workflow.Context, input *glue.UpdateClassifierInput) *GlueUpdateClassifierFuture

	UpdateColumnStatisticsForPartition(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error)
	UpdateColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) *GlueUpdateColumnStatisticsForPartitionFuture

	UpdateColumnStatisticsForTable(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error)
	UpdateColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) *GlueUpdateColumnStatisticsForTableFuture

	UpdateConnection(ctx workflow.Context, input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error)
	UpdateConnectionAsync(ctx workflow.Context, input *glue.UpdateConnectionInput) *GlueUpdateConnectionFuture

	UpdateCrawler(ctx workflow.Context, input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error)
	UpdateCrawlerAsync(ctx workflow.Context, input *glue.UpdateCrawlerInput) *GlueUpdateCrawlerFuture

	UpdateCrawlerSchedule(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error)
	UpdateCrawlerScheduleAsync(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) *GlueUpdateCrawlerScheduleFuture

	UpdateDatabase(ctx workflow.Context, input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error)
	UpdateDatabaseAsync(ctx workflow.Context, input *glue.UpdateDatabaseInput) *GlueUpdateDatabaseFuture

	UpdateDevEndpoint(ctx workflow.Context, input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error)
	UpdateDevEndpointAsync(ctx workflow.Context, input *glue.UpdateDevEndpointInput) *GlueUpdateDevEndpointFuture

	UpdateJob(ctx workflow.Context, input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error)
	UpdateJobAsync(ctx workflow.Context, input *glue.UpdateJobInput) *GlueUpdateJobFuture

	UpdateMLTransform(ctx workflow.Context, input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error)
	UpdateMLTransformAsync(ctx workflow.Context, input *glue.UpdateMLTransformInput) *GlueUpdateMLTransformFuture

	UpdatePartition(ctx workflow.Context, input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error)
	UpdatePartitionAsync(ctx workflow.Context, input *glue.UpdatePartitionInput) *GlueUpdatePartitionFuture

	UpdateTable(ctx workflow.Context, input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error)
	UpdateTableAsync(ctx workflow.Context, input *glue.UpdateTableInput) *GlueUpdateTableFuture

	UpdateTrigger(ctx workflow.Context, input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error)
	UpdateTriggerAsync(ctx workflow.Context, input *glue.UpdateTriggerInput) *GlueUpdateTriggerFuture

	UpdateUserDefinedFunction(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error)
	UpdateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) *GlueUpdateUserDefinedFunctionFuture

	UpdateWorkflow(ctx workflow.Context, input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error)
	UpdateWorkflowAsync(ctx workflow.Context, input *glue.UpdateWorkflowInput) *GlueUpdateWorkflowFuture
}

func NewClient() Client {
	return &stub{}
}
