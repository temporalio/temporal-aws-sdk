// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package applicationdiscoveryservicestub

import (
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateConfigurationItemsToApplication(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error)
	AssociateConfigurationItemsToApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) *ApplicationDiscoveryServiceAssociateConfigurationItemsToApplicationFuture

	BatchDeleteImportData(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error)
	BatchDeleteImportDataAsync(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) *ApplicationDiscoveryServiceBatchDeleteImportDataFuture

	CreateApplication(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) *ApplicationDiscoveryServiceCreateApplicationFuture

	CreateTags(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) *ApplicationDiscoveryServiceCreateTagsFuture

	DeleteApplications(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error)
	DeleteApplicationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) *ApplicationDiscoveryServiceDeleteApplicationsFuture

	DeleteTags(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) *ApplicationDiscoveryServiceDeleteTagsFuture

	DescribeAgents(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error)
	DescribeAgentsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) *ApplicationDiscoveryServiceDescribeAgentsFuture

	DescribeConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error)
	DescribeConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) *ApplicationDiscoveryServiceDescribeConfigurationsFuture

	DescribeContinuousExports(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error)
	DescribeContinuousExportsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) *ApplicationDiscoveryServiceDescribeContinuousExportsFuture

	DescribeExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error)
	DescribeExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) *ApplicationDiscoveryServiceDescribeExportConfigurationsFuture

	DescribeExportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error)
	DescribeExportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) *ApplicationDiscoveryServiceDescribeExportTasksFuture

	DescribeImportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error)
	DescribeImportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) *ApplicationDiscoveryServiceDescribeImportTasksFuture

	DescribeTags(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) *ApplicationDiscoveryServiceDescribeTagsFuture

	DisassociateConfigurationItemsFromApplication(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error)
	DisassociateConfigurationItemsFromApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) *ApplicationDiscoveryServiceDisassociateConfigurationItemsFromApplicationFuture

	ExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error)
	ExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) *ApplicationDiscoveryServiceExportConfigurationsFuture

	GetDiscoverySummary(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error)
	GetDiscoverySummaryAsync(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) *ApplicationDiscoveryServiceGetDiscoverySummaryFuture

	ListConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error)
	ListConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) *ApplicationDiscoveryServiceListConfigurationsFuture

	ListServerNeighbors(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error)
	ListServerNeighborsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) *ApplicationDiscoveryServiceListServerNeighborsFuture

	StartContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error)
	StartContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) *ApplicationDiscoveryServiceStartContinuousExportFuture

	StartDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error)
	StartDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) *ApplicationDiscoveryServiceStartDataCollectionByAgentIdsFuture

	StartExportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error)
	StartExportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) *ApplicationDiscoveryServiceStartExportTaskFuture

	StartImportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error)
	StartImportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) *ApplicationDiscoveryServiceStartImportTaskFuture

	StopContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error)
	StopContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) *ApplicationDiscoveryServiceStopContinuousExportFuture

	StopDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error)
	StopDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) *ApplicationDiscoveryServiceStopDataCollectionByAgentIdsFuture

	UpdateApplication(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) *ApplicationDiscoveryServiceUpdateApplicationFuture
}

func NewClient() Client {
	return &stub{}
}
