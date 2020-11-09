// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediatailorstub

import (
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DeletePlaybackConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePlaybackConfigurationFuture) Get(ctx workflow.Context) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	var output mediatailor.DeletePlaybackConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetPlaybackConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetPlaybackConfigurationFuture) Get(ctx workflow.Context) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	var output mediatailor.GetPlaybackConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPlaybackConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPlaybackConfigurationsFuture) Get(ctx workflow.Context) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	var output mediatailor.ListPlaybackConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*mediatailor.ListTagsForResourceOutput, error) {
	var output mediatailor.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutPlaybackConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutPlaybackConfigurationFuture) Get(ctx workflow.Context) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	var output mediatailor.PutPlaybackConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*mediatailor.TagResourceOutput, error) {
	var output mediatailor.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*mediatailor.UntagResourceOutput, error) {
	var output mediatailor.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePlaybackConfiguration(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	var output mediatailor.DeletePlaybackConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.DeletePlaybackConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) *DeletePlaybackConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.DeletePlaybackConfiguration", input)
	return &DeletePlaybackConfigurationFuture{Future: future}
}

func (a *stub) GetPlaybackConfiguration(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	var output mediatailor.GetPlaybackConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.GetPlaybackConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) *GetPlaybackConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.GetPlaybackConfiguration", input)
	return &GetPlaybackConfigurationFuture{Future: future}
}

func (a *stub) ListPlaybackConfigurations(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	var output mediatailor.ListPlaybackConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.ListPlaybackConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPlaybackConfigurationsAsync(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) *ListPlaybackConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.ListPlaybackConfigurations", input)
	return &ListPlaybackConfigurationsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	var output mediatailor.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutPlaybackConfiguration(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	var output mediatailor.PutPlaybackConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.PutPlaybackConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) *PutPlaybackConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.PutPlaybackConfiguration", input)
	return &PutPlaybackConfigurationFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
	var output mediatailor.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *mediatailor.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
	var output mediatailor.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediatailor.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *mediatailor.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediatailor.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
