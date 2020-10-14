// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package dataexchangestub

import (
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DataExchangeCancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeCancelJobFuture) Get(ctx workflow.Context) (*dataexchange.CancelJobOutput, error) {
	var output dataexchange.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeCreateDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeCreateDataSetFuture) Get(ctx workflow.Context) (*dataexchange.CreateDataSetOutput, error) {
	var output dataexchange.CreateDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeCreateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeCreateJobFuture) Get(ctx workflow.Context) (*dataexchange.CreateJobOutput, error) {
	var output dataexchange.CreateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeCreateRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeCreateRevisionFuture) Get(ctx workflow.Context) (*dataexchange.CreateRevisionOutput, error) {
	var output dataexchange.CreateRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeDeleteAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeDeleteAssetFuture) Get(ctx workflow.Context) (*dataexchange.DeleteAssetOutput, error) {
	var output dataexchange.DeleteAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeDeleteDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeDeleteDataSetFuture) Get(ctx workflow.Context) (*dataexchange.DeleteDataSetOutput, error) {
	var output dataexchange.DeleteDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeDeleteRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeDeleteRevisionFuture) Get(ctx workflow.Context) (*dataexchange.DeleteRevisionOutput, error) {
	var output dataexchange.DeleteRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeGetAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeGetAssetFuture) Get(ctx workflow.Context) (*dataexchange.GetAssetOutput, error) {
	var output dataexchange.GetAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeGetDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeGetDataSetFuture) Get(ctx workflow.Context) (*dataexchange.GetDataSetOutput, error) {
	var output dataexchange.GetDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeGetJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeGetJobFuture) Get(ctx workflow.Context) (*dataexchange.GetJobOutput, error) {
	var output dataexchange.GetJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeGetRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeGetRevisionFuture) Get(ctx workflow.Context) (*dataexchange.GetRevisionOutput, error) {
	var output dataexchange.GetRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeListDataSetRevisionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeListDataSetRevisionsFuture) Get(ctx workflow.Context) (*dataexchange.ListDataSetRevisionsOutput, error) {
	var output dataexchange.ListDataSetRevisionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeListDataSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeListDataSetsFuture) Get(ctx workflow.Context) (*dataexchange.ListDataSetsOutput, error) {
	var output dataexchange.ListDataSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeListJobsFuture) Get(ctx workflow.Context) (*dataexchange.ListJobsOutput, error) {
	var output dataexchange.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeListRevisionAssetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeListRevisionAssetsFuture) Get(ctx workflow.Context) (*dataexchange.ListRevisionAssetsOutput, error) {
	var output dataexchange.ListRevisionAssetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeListTagsForResourceFuture) Get(ctx workflow.Context) (*dataexchange.ListTagsForResourceOutput, error) {
	var output dataexchange.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeStartJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeStartJobFuture) Get(ctx workflow.Context) (*dataexchange.StartJobOutput, error) {
	var output dataexchange.StartJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeTagResourceFuture) Get(ctx workflow.Context) (*dataexchange.TagResourceOutput, error) {
	var output dataexchange.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeUntagResourceFuture) Get(ctx workflow.Context) (*dataexchange.UntagResourceOutput, error) {
	var output dataexchange.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeUpdateAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeUpdateAssetFuture) Get(ctx workflow.Context) (*dataexchange.UpdateAssetOutput, error) {
	var output dataexchange.UpdateAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeUpdateDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeUpdateDataSetFuture) Get(ctx workflow.Context) (*dataexchange.UpdateDataSetOutput, error) {
	var output dataexchange.UpdateDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataExchangeUpdateRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataExchangeUpdateRevisionFuture) Get(ctx workflow.Context) (*dataexchange.UpdateRevisionOutput, error) {
	var output dataexchange.UpdateRevisionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJob(ctx workflow.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
	var output dataexchange.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *dataexchange.CancelJobInput) *DataExchangeCancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.CancelJob", input)
	return &DataExchangeCancelJobFuture{Future: future}
}

func (a *stub) CreateDataSet(ctx workflow.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
	var output dataexchange.CreateDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSetAsync(ctx workflow.Context, input *dataexchange.CreateDataSetInput) *DataExchangeCreateDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateDataSet", input)
	return &DataExchangeCreateDataSetFuture{Future: future}
}

func (a *stub) CreateJob(ctx workflow.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
	var output dataexchange.CreateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobAsync(ctx workflow.Context, input *dataexchange.CreateJobInput) *DataExchangeCreateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateJob", input)
	return &DataExchangeCreateJobFuture{Future: future}
}

func (a *stub) CreateRevision(ctx workflow.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
	var output dataexchange.CreateRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRevisionAsync(ctx workflow.Context, input *dataexchange.CreateRevisionInput) *DataExchangeCreateRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.CreateRevision", input)
	return &DataExchangeCreateRevisionFuture{Future: future}
}

func (a *stub) DeleteAsset(ctx workflow.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
	var output dataexchange.DeleteAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssetAsync(ctx workflow.Context, input *dataexchange.DeleteAssetInput) *DataExchangeDeleteAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteAsset", input)
	return &DataExchangeDeleteAssetFuture{Future: future}
}

func (a *stub) DeleteDataSet(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
	var output dataexchange.DeleteDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataSetAsync(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) *DataExchangeDeleteDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteDataSet", input)
	return &DataExchangeDeleteDataSetFuture{Future: future}
}

func (a *stub) DeleteRevision(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
	var output dataexchange.DeleteRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRevisionAsync(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) *DataExchangeDeleteRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.DeleteRevision", input)
	return &DataExchangeDeleteRevisionFuture{Future: future}
}

func (a *stub) GetAsset(ctx workflow.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	var output dataexchange.GetAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAssetAsync(ctx workflow.Context, input *dataexchange.GetAssetInput) *DataExchangeGetAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetAsset", input)
	return &DataExchangeGetAssetFuture{Future: future}
}

func (a *stub) GetDataSet(ctx workflow.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	var output dataexchange.GetDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataSetAsync(ctx workflow.Context, input *dataexchange.GetDataSetInput) *DataExchangeGetDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetDataSet", input)
	return &DataExchangeGetDataSetFuture{Future: future}
}

func (a *stub) GetJob(ctx workflow.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	var output dataexchange.GetJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobAsync(ctx workflow.Context, input *dataexchange.GetJobInput) *DataExchangeGetJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetJob", input)
	return &DataExchangeGetJobFuture{Future: future}
}

func (a *stub) GetRevision(ctx workflow.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	var output dataexchange.GetRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRevisionAsync(ctx workflow.Context, input *dataexchange.GetRevisionInput) *DataExchangeGetRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.GetRevision", input)
	return &DataExchangeGetRevisionFuture{Future: future}
}

func (a *stub) ListDataSetRevisions(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	var output dataexchange.ListDataSetRevisionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListDataSetRevisions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSetRevisionsAsync(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) *DataExchangeListDataSetRevisionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListDataSetRevisions", input)
	return &DataExchangeListDataSetRevisionsFuture{Future: future}
}

func (a *stub) ListDataSets(ctx workflow.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	var output dataexchange.ListDataSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListDataSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSetsAsync(ctx workflow.Context, input *dataexchange.ListDataSetsInput) *DataExchangeListDataSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListDataSets", input)
	return &DataExchangeListDataSetsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	var output dataexchange.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *dataexchange.ListJobsInput) *DataExchangeListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListJobs", input)
	return &DataExchangeListJobsFuture{Future: future}
}

func (a *stub) ListRevisionAssets(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	var output dataexchange.ListRevisionAssetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListRevisionAssets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRevisionAssetsAsync(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) *DataExchangeListRevisionAssetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListRevisionAssets", input)
	return &DataExchangeListRevisionAssetsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	var output dataexchange.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) *DataExchangeListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.ListTagsForResource", input)
	return &DataExchangeListTagsForResourceFuture{Future: future}
}

func (a *stub) StartJob(ctx workflow.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
	var output dataexchange.StartJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.StartJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartJobAsync(ctx workflow.Context, input *dataexchange.StartJobInput) *DataExchangeStartJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.StartJob", input)
	return &DataExchangeStartJobFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
	var output dataexchange.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dataexchange.TagResourceInput) *DataExchangeTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.TagResource", input)
	return &DataExchangeTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
	var output dataexchange.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dataexchange.UntagResourceInput) *DataExchangeUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.UntagResource", input)
	return &DataExchangeUntagResourceFuture{Future: future}
}

func (a *stub) UpdateAsset(ctx workflow.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
	var output dataexchange.UpdateAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAssetAsync(ctx workflow.Context, input *dataexchange.UpdateAssetInput) *DataExchangeUpdateAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateAsset", input)
	return &DataExchangeUpdateAssetFuture{Future: future}
}

func (a *stub) UpdateDataSet(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
	var output dataexchange.UpdateDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataSetAsync(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) *DataExchangeUpdateDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateDataSet", input)
	return &DataExchangeUpdateDataSetFuture{Future: future}
}

func (a *stub) UpdateRevision(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
	var output dataexchange.UpdateRevisionOutput
	err := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRevisionAsync(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) *DataExchangeUpdateRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dataexchange.UpdateRevision", input)
	return &DataExchangeUpdateRevisionFuture{Future: future}
}
