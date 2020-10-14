// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package augmentedairuntimestub

import (
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AugmentedAIRuntimeDeleteHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AugmentedAIRuntimeDeleteHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	var output augmentedairuntime.DeleteHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AugmentedAIRuntimeDescribeHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AugmentedAIRuntimeDescribeHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	var output augmentedairuntime.DescribeHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AugmentedAIRuntimeListHumanLoopsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AugmentedAIRuntimeListHumanLoopsFuture) Get(ctx workflow.Context) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	var output augmentedairuntime.ListHumanLoopsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AugmentedAIRuntimeStartHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AugmentedAIRuntimeStartHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.StartHumanLoopOutput, error) {
	var output augmentedairuntime.StartHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AugmentedAIRuntimeStopHumanLoopFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AugmentedAIRuntimeStopHumanLoopFuture) Get(ctx workflow.Context) (*augmentedairuntime.StopHumanLoopOutput, error) {
	var output augmentedairuntime.StopHumanLoopOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	var output augmentedairuntime.DeleteHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.DeleteHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *AugmentedAIRuntimeDeleteHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.DeleteHumanLoop", input)
	return &AugmentedAIRuntimeDeleteHumanLoopFuture{Future: future}
}

func (a *stub) DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	var output augmentedairuntime.DescribeHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.DescribeHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *AugmentedAIRuntimeDescribeHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.DescribeHumanLoop", input)
	return &AugmentedAIRuntimeDescribeHumanLoopFuture{Future: future}
}

func (a *stub) ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	var output augmentedairuntime.ListHumanLoopsOutput
	err := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.ListHumanLoops", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *AugmentedAIRuntimeListHumanLoopsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.ListHumanLoops", input)
	return &AugmentedAIRuntimeListHumanLoopsFuture{Future: future}
}

func (a *stub) StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
	var output augmentedairuntime.StartHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.StartHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *AugmentedAIRuntimeStartHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.StartHumanLoop", input)
	return &AugmentedAIRuntimeStartHumanLoopFuture{Future: future}
}

func (a *stub) StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
	var output augmentedairuntime.StopHumanLoopOutput
	err := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.StopHumanLoop", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *AugmentedAIRuntimeStopHumanLoopFuture {
	future := workflow.ExecuteActivity(ctx, "aws.augmentedairuntime.StopHumanLoop", input)
	return &AugmentedAIRuntimeStopHumanLoopFuture{Future: future}
}
