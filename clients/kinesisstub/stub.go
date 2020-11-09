// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisstub

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddTagsToStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsToStreamFuture) Get(ctx workflow.Context) (*kinesis.AddTagsToStreamOutput, error) {
	var output kinesis.AddTagsToStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateStreamFuture) Get(ctx workflow.Context) (*kinesis.CreateStreamOutput, error) {
	var output kinesis.CreateStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DecreaseStreamRetentionPeriodFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DecreaseStreamRetentionPeriodFuture) Get(ctx workflow.Context) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.DecreaseStreamRetentionPeriodOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteStreamFuture) Get(ctx workflow.Context) (*kinesis.DeleteStreamOutput, error) {
	var output kinesis.DeleteStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeregisterStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeregisterStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.DeregisterStreamConsumerOutput, error) {
	var output kinesis.DeregisterStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLimitsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLimitsFuture) Get(ctx workflow.Context) (*kinesis.DescribeLimitsOutput, error) {
	var output kinesis.DescribeLimitsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStreamFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamOutput, error) {
	var output kinesis.DescribeStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamConsumerOutput, error) {
	var output kinesis.DescribeStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeStreamSummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeStreamSummaryFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamSummaryOutput, error) {
	var output kinesis.DescribeStreamSummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableEnhancedMonitoringFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableEnhancedMonitoringFuture) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableEnhancedMonitoringFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableEnhancedMonitoringFuture) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRecordsFuture) Get(ctx workflow.Context) (*kinesis.GetRecordsOutput, error) {
	var output kinesis.GetRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetShardIteratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetShardIteratorFuture) Get(ctx workflow.Context) (*kinesis.GetShardIteratorOutput, error) {
	var output kinesis.GetShardIteratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IncreaseStreamRetentionPeriodFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IncreaseStreamRetentionPeriodFuture) Get(ctx workflow.Context) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.IncreaseStreamRetentionPeriodOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListShardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListShardsFuture) Get(ctx workflow.Context) (*kinesis.ListShardsOutput, error) {
	var output kinesis.ListShardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListStreamConsumersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListStreamConsumersFuture) Get(ctx workflow.Context) (*kinesis.ListStreamConsumersOutput, error) {
	var output kinesis.ListStreamConsumersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListStreamsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListStreamsFuture) Get(ctx workflow.Context) (*kinesis.ListStreamsOutput, error) {
	var output kinesis.ListStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForStreamFuture) Get(ctx workflow.Context) (*kinesis.ListTagsForStreamOutput, error) {
	var output kinesis.ListTagsForStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MergeShardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MergeShardsFuture) Get(ctx workflow.Context) (*kinesis.MergeShardsOutput, error) {
	var output kinesis.MergeShardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutRecordFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutRecordFuture) Get(ctx workflow.Context) (*kinesis.PutRecordOutput, error) {
	var output kinesis.PutRecordOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutRecordsFuture) Get(ctx workflow.Context) (*kinesis.PutRecordsOutput, error) {
	var output kinesis.PutRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RegisterStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RegisterStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.RegisterStreamConsumerOutput, error) {
	var output kinesis.RegisterStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFromStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFromStreamFuture) Get(ctx workflow.Context) (*kinesis.RemoveTagsFromStreamOutput, error) {
	var output kinesis.RemoveTagsFromStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SplitShardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SplitShardFuture) Get(ctx workflow.Context) (*kinesis.SplitShardOutput, error) {
	var output kinesis.SplitShardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartStreamEncryptionFuture) Get(ctx workflow.Context) (*kinesis.StartStreamEncryptionOutput, error) {
	var output kinesis.StartStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopStreamEncryptionFuture) Get(ctx workflow.Context) (*kinesis.StopStreamEncryptionOutput, error) {
	var output kinesis.StopStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SubscribeToShardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SubscribeToShardFuture) Get(ctx workflow.Context) (*kinesis.SubscribeToShardOutput, error) {
	var output kinesis.SubscribeToShardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateShardCountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateShardCountFuture) Get(ctx workflow.Context) (*kinesis.UpdateShardCountOutput, error) {
	var output kinesis.UpdateShardCountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToStream(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
	var output kinesis.AddTagsToStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.AddTagsToStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToStreamAsync(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) *AddTagsToStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.AddTagsToStream", input)
	return &AddTagsToStreamFuture{Future: future}
}

func (a *stub) CreateStream(ctx workflow.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
	var output kinesis.CreateStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.CreateStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateStreamAsync(ctx workflow.Context, input *kinesis.CreateStreamInput) *CreateStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.CreateStream", input)
	return &CreateStreamFuture{Future: future}
}

func (a *stub) DecreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.DecreaseStreamRetentionPeriodOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DecreaseStreamRetentionPeriod", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DecreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) *DecreaseStreamRetentionPeriodFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DecreaseStreamRetentionPeriod", input)
	return &DecreaseStreamRetentionPeriodFuture{Future: future}
}

func (a *stub) DeleteStream(ctx workflow.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
	var output kinesis.DeleteStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DeleteStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteStreamAsync(ctx workflow.Context, input *kinesis.DeleteStreamInput) *DeleteStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DeleteStream", input)
	return &DeleteStreamFuture{Future: future}
}

func (a *stub) DeregisterStreamConsumer(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
	var output kinesis.DeregisterStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DeregisterStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) *DeregisterStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DeregisterStreamConsumer", input)
	return &DeregisterStreamConsumerFuture{Future: future}
}

func (a *stub) DescribeLimits(ctx workflow.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
	var output kinesis.DescribeLimitsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLimitsAsync(ctx workflow.Context, input *kinesis.DescribeLimitsInput) *DescribeLimitsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeLimits", input)
	return &DescribeLimitsFuture{Future: future}
}

func (a *stub) DescribeStream(ctx workflow.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	var output kinesis.DescribeStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *DescribeStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStream", input)
	return &DescribeStreamFuture{Future: future}
}

func (a *stub) DescribeStreamConsumer(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
	var output kinesis.DescribeStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamConsumerAsync(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) *DescribeStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamConsumer", input)
	return &DescribeStreamConsumerFuture{Future: future}
}

func (a *stub) DescribeStreamSummary(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
	var output kinesis.DescribeStreamSummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamSummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamSummaryAsync(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) *DescribeStreamSummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamSummary", input)
	return &DescribeStreamSummaryFuture{Future: future}
}

func (a *stub) DisableEnhancedMonitoring(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DisableEnhancedMonitoring", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) *DisableEnhancedMonitoringFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DisableEnhancedMonitoring", input)
	return &DisableEnhancedMonitoringFuture{Future: future}
}

func (a *stub) EnableEnhancedMonitoring(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.EnableEnhancedMonitoring", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) *EnableEnhancedMonitoringFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.EnableEnhancedMonitoring", input)
	return &EnableEnhancedMonitoringFuture{Future: future}
}

func (a *stub) GetRecords(ctx workflow.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	var output kinesis.GetRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.GetRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecordsAsync(ctx workflow.Context, input *kinesis.GetRecordsInput) *GetRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.GetRecords", input)
	return &GetRecordsFuture{Future: future}
}

func (a *stub) GetShardIterator(ctx workflow.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
	var output kinesis.GetShardIteratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.GetShardIterator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetShardIteratorAsync(ctx workflow.Context, input *kinesis.GetShardIteratorInput) *GetShardIteratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.GetShardIterator", input)
	return &GetShardIteratorFuture{Future: future}
}

func (a *stub) IncreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.IncreaseStreamRetentionPeriodOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.IncreaseStreamRetentionPeriod", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IncreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) *IncreaseStreamRetentionPeriodFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.IncreaseStreamRetentionPeriod", input)
	return &IncreaseStreamRetentionPeriodFuture{Future: future}
}

func (a *stub) ListShards(ctx workflow.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	var output kinesis.ListShardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListShards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListShardsAsync(ctx workflow.Context, input *kinesis.ListShardsInput) *ListShardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListShards", input)
	return &ListShardsFuture{Future: future}
}

func (a *stub) ListStreamConsumers(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
	var output kinesis.ListStreamConsumersOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreamConsumers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamConsumersAsync(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) *ListStreamConsumersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreamConsumers", input)
	return &ListStreamConsumersFuture{Future: future}
}

func (a *stub) ListStreams(ctx workflow.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	var output kinesis.ListStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamsAsync(ctx workflow.Context, input *kinesis.ListStreamsInput) *ListStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreams", input)
	return &ListStreamsFuture{Future: future}
}

func (a *stub) ListTagsForStream(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
	var output kinesis.ListTagsForStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListTagsForStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForStreamAsync(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) *ListTagsForStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListTagsForStream", input)
	return &ListTagsForStreamFuture{Future: future}
}

func (a *stub) MergeShards(ctx workflow.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
	var output kinesis.MergeShardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.MergeShards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) MergeShardsAsync(ctx workflow.Context, input *kinesis.MergeShardsInput) *MergeShardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.MergeShards", input)
	return &MergeShardsFuture{Future: future}
}

func (a *stub) PutRecord(ctx workflow.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	var output kinesis.PutRecordOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecord", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordAsync(ctx workflow.Context, input *kinesis.PutRecordInput) *PutRecordFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecord", input)
	return &PutRecordFuture{Future: future}
}

func (a *stub) PutRecords(ctx workflow.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
	var output kinesis.PutRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordsAsync(ctx workflow.Context, input *kinesis.PutRecordsInput) *PutRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecords", input)
	return &PutRecordsFuture{Future: future}
}

func (a *stub) RegisterStreamConsumer(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
	var output kinesis.RegisterStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.RegisterStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) *RegisterStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.RegisterStreamConsumer", input)
	return &RegisterStreamConsumerFuture{Future: future}
}

func (a *stub) RemoveTagsFromStream(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
	var output kinesis.RemoveTagsFromStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.RemoveTagsFromStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsFromStreamAsync(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) *RemoveTagsFromStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.RemoveTagsFromStream", input)
	return &RemoveTagsFromStreamFuture{Future: future}
}

func (a *stub) SplitShard(ctx workflow.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
	var output kinesis.SplitShardOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.SplitShard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SplitShardAsync(ctx workflow.Context, input *kinesis.SplitShardInput) *SplitShardFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.SplitShard", input)
	return &SplitShardFuture{Future: future}
}

func (a *stub) StartStreamEncryption(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
	var output kinesis.StartStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.StartStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) *StartStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.StartStreamEncryption", input)
	return &StartStreamEncryptionFuture{Future: future}
}

func (a *stub) StopStreamEncryption(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
	var output kinesis.StopStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.StopStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) *StopStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.StopStreamEncryption", input)
	return &StopStreamEncryptionFuture{Future: future}
}

func (a *stub) SubscribeToShard(ctx workflow.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
	var output kinesis.SubscribeToShardOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.SubscribeToShard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubscribeToShardAsync(ctx workflow.Context, input *kinesis.SubscribeToShardInput) *SubscribeToShardFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.SubscribeToShard", input)
	return &SubscribeToShardFuture{Future: future}
}

func (a *stub) UpdateShardCount(ctx workflow.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
	var output kinesis.UpdateShardCountOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.UpdateShardCount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateShardCountAsync(ctx workflow.Context, input *kinesis.UpdateShardCountInput) *UpdateShardCountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.UpdateShardCount", input)
	return &UpdateShardCountFuture{Future: future}
}

func (a *stub) WaitUntilStreamExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
	return workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilStreamExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamExists", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilStreamNotExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
	return workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamNotExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilStreamNotExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamNotExists", input)
	return clients.NewVoidFuture(future)
}
