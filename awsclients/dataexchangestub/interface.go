// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package dataexchangestub

import (
	"github.com/aws/aws-sdk-go/service/dataexchange"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CancelJob(ctx workflow.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *dataexchange.CancelJobInput) *DataExchangeCancelJobFuture

	CreateDataSet(ctx workflow.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error)
	CreateDataSetAsync(ctx workflow.Context, input *dataexchange.CreateDataSetInput) *DataExchangeCreateDataSetFuture

	CreateJob(ctx workflow.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *dataexchange.CreateJobInput) *DataExchangeCreateJobFuture

	CreateRevision(ctx workflow.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error)
	CreateRevisionAsync(ctx workflow.Context, input *dataexchange.CreateRevisionInput) *DataExchangeCreateRevisionFuture

	DeleteAsset(ctx workflow.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error)
	DeleteAssetAsync(ctx workflow.Context, input *dataexchange.DeleteAssetInput) *DataExchangeDeleteAssetFuture

	DeleteDataSet(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error)
	DeleteDataSetAsync(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) *DataExchangeDeleteDataSetFuture

	DeleteRevision(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error)
	DeleteRevisionAsync(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) *DataExchangeDeleteRevisionFuture

	GetAsset(ctx workflow.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error)
	GetAssetAsync(ctx workflow.Context, input *dataexchange.GetAssetInput) *DataExchangeGetAssetFuture

	GetDataSet(ctx workflow.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error)
	GetDataSetAsync(ctx workflow.Context, input *dataexchange.GetDataSetInput) *DataExchangeGetDataSetFuture

	GetJob(ctx workflow.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *dataexchange.GetJobInput) *DataExchangeGetJobFuture

	GetRevision(ctx workflow.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error)
	GetRevisionAsync(ctx workflow.Context, input *dataexchange.GetRevisionInput) *DataExchangeGetRevisionFuture

	ListDataSetRevisions(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error)
	ListDataSetRevisionsAsync(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) *DataExchangeListDataSetRevisionsFuture

	ListDataSets(ctx workflow.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error)
	ListDataSetsAsync(ctx workflow.Context, input *dataexchange.ListDataSetsInput) *DataExchangeListDataSetsFuture

	ListJobs(ctx workflow.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *dataexchange.ListJobsInput) *DataExchangeListJobsFuture

	ListRevisionAssets(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error)
	ListRevisionAssetsAsync(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) *DataExchangeListRevisionAssetsFuture

	ListTagsForResource(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) *DataExchangeListTagsForResourceFuture

	StartJob(ctx workflow.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error)
	StartJobAsync(ctx workflow.Context, input *dataexchange.StartJobInput) *DataExchangeStartJobFuture

	TagResource(ctx workflow.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dataexchange.TagResourceInput) *DataExchangeTagResourceFuture

	UntagResource(ctx workflow.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dataexchange.UntagResourceInput) *DataExchangeUntagResourceFuture

	UpdateAsset(ctx workflow.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error)
	UpdateAssetAsync(ctx workflow.Context, input *dataexchange.UpdateAssetInput) *DataExchangeUpdateAssetFuture

	UpdateDataSet(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error)
	UpdateDataSetAsync(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) *DataExchangeUpdateDataSetFuture

	UpdateRevision(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error)
	UpdateRevisionAsync(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) *DataExchangeUpdateRevisionFuture
}

func NewClient() Client {
	return &stub{}
}

