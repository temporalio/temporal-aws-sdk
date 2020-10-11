// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ebsstub

import (
	"github.com/aws/aws-sdk-go/service/ebs"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type EBSCompleteSnapshotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSCompleteSnapshotFuture) Get(ctx workflow.Context) (*ebs.CompleteSnapshotOutput, error) {
	var output ebs.CompleteSnapshotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EBSGetSnapshotBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSGetSnapshotBlockFuture) Get(ctx workflow.Context) (*ebs.GetSnapshotBlockOutput, error) {
	var output ebs.GetSnapshotBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EBSListChangedBlocksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSListChangedBlocksFuture) Get(ctx workflow.Context) (*ebs.ListChangedBlocksOutput, error) {
	var output ebs.ListChangedBlocksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EBSListSnapshotBlocksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSListSnapshotBlocksFuture) Get(ctx workflow.Context) (*ebs.ListSnapshotBlocksOutput, error) {
	var output ebs.ListSnapshotBlocksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EBSPutSnapshotBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSPutSnapshotBlockFuture) Get(ctx workflow.Context) (*ebs.PutSnapshotBlockOutput, error) {
	var output ebs.PutSnapshotBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EBSStartSnapshotFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EBSStartSnapshotFuture) Get(ctx workflow.Context) (*ebs.StartSnapshotOutput, error) {
	var output ebs.StartSnapshotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
	var output ebs.CompleteSnapshotOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.CompleteSnapshot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CompleteSnapshotAsync(ctx workflow.Context, input *ebs.CompleteSnapshotInput) *EBSCompleteSnapshotFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.CompleteSnapshot", input)
	return &EBSCompleteSnapshotFuture{Future: future}
}

func (a *stub) GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
	var output ebs.GetSnapshotBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.GetSnapshotBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSnapshotBlockAsync(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) *EBSGetSnapshotBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.GetSnapshotBlock", input)
	return &EBSGetSnapshotBlockFuture{Future: future}
}

func (a *stub) ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
	var output ebs.ListChangedBlocksOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.ListChangedBlocks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListChangedBlocksAsync(ctx workflow.Context, input *ebs.ListChangedBlocksInput) *EBSListChangedBlocksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.ListChangedBlocks", input)
	return &EBSListChangedBlocksFuture{Future: future}
}

func (a *stub) ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
	var output ebs.ListSnapshotBlocksOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.ListSnapshotBlocks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSnapshotBlocksAsync(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) *EBSListSnapshotBlocksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.ListSnapshotBlocks", input)
	return &EBSListSnapshotBlocksFuture{Future: future}
}

func (a *stub) PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error) {
	var output ebs.PutSnapshotBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.PutSnapshotBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutSnapshotBlockAsync(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) *EBSPutSnapshotBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.PutSnapshotBlock", input)
	return &EBSPutSnapshotBlockFuture{Future: future}
}

func (a *stub) StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error) {
	var output ebs.StartSnapshotOutput
	err := workflow.ExecuteActivity(ctx, "aws.ebs.StartSnapshot", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartSnapshotAsync(ctx workflow.Context, input *ebs.StartSnapshotInput) *EBSStartSnapshotFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ebs.StartSnapshot", input)
	return &EBSStartSnapshotFuture{Future: future}
}