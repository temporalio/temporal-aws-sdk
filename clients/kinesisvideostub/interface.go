// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisvideostub

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateSignalingChannel(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error)
	CreateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) *KinesisVideoCreateSignalingChannelFuture

	CreateStream(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error)
	CreateStreamAsync(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) *KinesisVideoCreateStreamFuture

	DeleteSignalingChannel(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error)
	DeleteSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) *KinesisVideoDeleteSignalingChannelFuture

	DeleteStream(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error)
	DeleteStreamAsync(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) *KinesisVideoDeleteStreamFuture

	DescribeSignalingChannel(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error)
	DescribeSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) *KinesisVideoDescribeSignalingChannelFuture

	DescribeStream(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error)
	DescribeStreamAsync(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) *KinesisVideoDescribeStreamFuture

	GetDataEndpoint(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error)
	GetDataEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) *KinesisVideoGetDataEndpointFuture

	GetSignalingChannelEndpoint(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error)
	GetSignalingChannelEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) *KinesisVideoGetSignalingChannelEndpointFuture

	ListSignalingChannels(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error)
	ListSignalingChannelsAsync(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) *KinesisVideoListSignalingChannelsFuture

	ListStreams(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) *KinesisVideoListStreamsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) *KinesisVideoListTagsForResourceFuture

	ListTagsForStream(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error)
	ListTagsForStreamAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) *KinesisVideoListTagsForStreamFuture

	TagResource(ctx workflow.Context, input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisvideo.TagResourceInput) *KinesisVideoTagResourceFuture

	TagStream(ctx workflow.Context, input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error)
	TagStreamAsync(ctx workflow.Context, input *kinesisvideo.TagStreamInput) *KinesisVideoTagStreamFuture

	UntagResource(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) *KinesisVideoUntagResourceFuture

	UntagStream(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error)
	UntagStreamAsync(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) *KinesisVideoUntagStreamFuture

	UpdateDataRetention(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error)
	UpdateDataRetentionAsync(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) *KinesisVideoUpdateDataRetentionFuture

	UpdateSignalingChannel(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error)
	UpdateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) *KinesisVideoUpdateSignalingChannelFuture

	UpdateStream(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error)
	UpdateStreamAsync(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) *KinesisVideoUpdateStreamFuture
}

func NewClient() Client {
	return &stub{}
}
