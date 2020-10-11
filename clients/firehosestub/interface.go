// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package firehosestub

import (
	"github.com/aws/aws-sdk-go/service/firehose"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateDeliveryStream(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error)
	CreateDeliveryStreamAsync(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) *FirehoseCreateDeliveryStreamFuture

	DeleteDeliveryStream(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error)
	DeleteDeliveryStreamAsync(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) *FirehoseDeleteDeliveryStreamFuture

	DescribeDeliveryStream(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error)
	DescribeDeliveryStreamAsync(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) *FirehoseDescribeDeliveryStreamFuture

	ListDeliveryStreams(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error)
	ListDeliveryStreamsAsync(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) *FirehoseListDeliveryStreamsFuture

	ListTagsForDeliveryStream(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error)
	ListTagsForDeliveryStreamAsync(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) *FirehoseListTagsForDeliveryStreamFuture

	PutRecord(ctx workflow.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
	PutRecordAsync(ctx workflow.Context, input *firehose.PutRecordInput) *FirehosePutRecordFuture

	PutRecordBatch(ctx workflow.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error)
	PutRecordBatchAsync(ctx workflow.Context, input *firehose.PutRecordBatchInput) *FirehosePutRecordBatchFuture

	StartDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error)
	StartDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) *FirehoseStartDeliveryStreamEncryptionFuture

	StopDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error)
	StopDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) *FirehoseStopDeliveryStreamEncryptionFuture

	TagDeliveryStream(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error)
	TagDeliveryStreamAsync(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) *FirehoseTagDeliveryStreamFuture

	UntagDeliveryStream(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error)
	UntagDeliveryStreamAsync(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) *FirehoseUntagDeliveryStreamFuture

	UpdateDestination(ctx workflow.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error)
	UpdateDestinationAsync(ctx workflow.Context, input *firehose.UpdateDestinationInput) *FirehoseUpdateDestinationFuture
}

func NewClient() Client {
	return &stub{}
}
