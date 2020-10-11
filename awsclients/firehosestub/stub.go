// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package firehosestub

import (
	"github.com/aws/aws-sdk-go/service/firehose"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type FirehoseCreateDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseCreateDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.CreateDeliveryStreamOutput, error) {
	var output firehose.CreateDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseDeleteDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseDeleteDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.DeleteDeliveryStreamOutput, error) {
	var output firehose.DeleteDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseDescribeDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseDescribeDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.DescribeDeliveryStreamOutput, error) {
	var output firehose.DescribeDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseListDeliveryStreamsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseListDeliveryStreamsFuture) Get(ctx workflow.Context) (*firehose.ListDeliveryStreamsOutput, error) {
	var output firehose.ListDeliveryStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseListTagsForDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseListTagsForDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	var output firehose.ListTagsForDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehosePutRecordFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehosePutRecordFuture) Get(ctx workflow.Context) (*firehose.PutRecordOutput, error) {
	var output firehose.PutRecordOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehosePutRecordBatchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehosePutRecordBatchFuture) Get(ctx workflow.Context) (*firehose.PutRecordBatchOutput, error) {
	var output firehose.PutRecordBatchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseStartDeliveryStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseStartDeliveryStreamEncryptionFuture) Get(ctx workflow.Context) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	var output firehose.StartDeliveryStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseStopDeliveryStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseStopDeliveryStreamEncryptionFuture) Get(ctx workflow.Context) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	var output firehose.StopDeliveryStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseTagDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseTagDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.TagDeliveryStreamOutput, error) {
	var output firehose.TagDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseUntagDeliveryStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseUntagDeliveryStreamFuture) Get(ctx workflow.Context) (*firehose.UntagDeliveryStreamOutput, error) {
	var output firehose.UntagDeliveryStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FirehoseUpdateDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FirehoseUpdateDestinationFuture) Get(ctx workflow.Context) (*firehose.UpdateDestinationOutput, error) {
	var output firehose.UpdateDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeliveryStream(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
	var output firehose.CreateDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.CreateDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeliveryStreamAsync(ctx workflow.Context, input *firehose.CreateDeliveryStreamInput) *FirehoseCreateDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.CreateDeliveryStream", input)
	return &FirehoseCreateDeliveryStreamFuture{Future: future}
}

func (a *stub) DeleteDeliveryStream(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
	var output firehose.DeleteDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.DeleteDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDeliveryStreamAsync(ctx workflow.Context, input *firehose.DeleteDeliveryStreamInput) *FirehoseDeleteDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.DeleteDeliveryStream", input)
	return &FirehoseDeleteDeliveryStreamFuture{Future: future}
}

func (a *stub) DescribeDeliveryStream(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
	var output firehose.DescribeDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.DescribeDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDeliveryStreamAsync(ctx workflow.Context, input *firehose.DescribeDeliveryStreamInput) *FirehoseDescribeDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.DescribeDeliveryStream", input)
	return &FirehoseDescribeDeliveryStreamFuture{Future: future}
}

func (a *stub) ListDeliveryStreams(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
	var output firehose.ListDeliveryStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.ListDeliveryStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDeliveryStreamsAsync(ctx workflow.Context, input *firehose.ListDeliveryStreamsInput) *FirehoseListDeliveryStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.ListDeliveryStreams", input)
	return &FirehoseListDeliveryStreamsFuture{Future: future}
}

func (a *stub) ListTagsForDeliveryStream(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	var output firehose.ListTagsForDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.ListTagsForDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForDeliveryStreamAsync(ctx workflow.Context, input *firehose.ListTagsForDeliveryStreamInput) *FirehoseListTagsForDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.ListTagsForDeliveryStream", input)
	return &FirehoseListTagsForDeliveryStreamFuture{Future: future}
}

func (a *stub) PutRecord(ctx workflow.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
	var output firehose.PutRecordOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.PutRecord", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordAsync(ctx workflow.Context, input *firehose.PutRecordInput) *FirehosePutRecordFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.PutRecord", input)
	return &FirehosePutRecordFuture{Future: future}
}

func (a *stub) PutRecordBatch(ctx workflow.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
	var output firehose.PutRecordBatchOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.PutRecordBatch", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordBatchAsync(ctx workflow.Context, input *firehose.PutRecordBatchInput) *FirehosePutRecordBatchFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.PutRecordBatch", input)
	return &FirehosePutRecordBatchFuture{Future: future}
}

func (a *stub) StartDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	var output firehose.StartDeliveryStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.StartDeliveryStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StartDeliveryStreamEncryptionInput) *FirehoseStartDeliveryStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.StartDeliveryStreamEncryption", input)
	return &FirehoseStartDeliveryStreamEncryptionFuture{Future: future}
}

func (a *stub) StopDeliveryStreamEncryption(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	var output firehose.StopDeliveryStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.StopDeliveryStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopDeliveryStreamEncryptionAsync(ctx workflow.Context, input *firehose.StopDeliveryStreamEncryptionInput) *FirehoseStopDeliveryStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.StopDeliveryStreamEncryption", input)
	return &FirehoseStopDeliveryStreamEncryptionFuture{Future: future}
}

func (a *stub) TagDeliveryStream(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
	var output firehose.TagDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.TagDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagDeliveryStreamAsync(ctx workflow.Context, input *firehose.TagDeliveryStreamInput) *FirehoseTagDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.TagDeliveryStream", input)
	return &FirehoseTagDeliveryStreamFuture{Future: future}
}

func (a *stub) UntagDeliveryStream(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
	var output firehose.UntagDeliveryStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.UntagDeliveryStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagDeliveryStreamAsync(ctx workflow.Context, input *firehose.UntagDeliveryStreamInput) *FirehoseUntagDeliveryStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.UntagDeliveryStream", input)
	return &FirehoseUntagDeliveryStreamFuture{Future: future}
}

func (a *stub) UpdateDestination(ctx workflow.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
	var output firehose.UpdateDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.firehose.UpdateDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDestinationAsync(ctx workflow.Context, input *firehose.UpdateDestinationInput) *FirehoseUpdateDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.firehose.UpdateDestination", input)
	return &FirehoseUpdateDestinationFuture{Future: future}
}
