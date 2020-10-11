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

type Client interface {
	CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error)
	CompleteSnapshotAsync(ctx workflow.Context, input *ebs.CompleteSnapshotInput) *EBSCompleteSnapshotFuture

	GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error)
	GetSnapshotBlockAsync(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) *EBSGetSnapshotBlockFuture

	ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error)
	ListChangedBlocksAsync(ctx workflow.Context, input *ebs.ListChangedBlocksInput) *EBSListChangedBlocksFuture

	ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error)
	ListSnapshotBlocksAsync(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) *EBSListSnapshotBlocksFuture

	PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error)
	PutSnapshotBlockAsync(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) *EBSPutSnapshotBlockFuture

	StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error)
	StartSnapshotAsync(ctx workflow.Context, input *ebs.StartSnapshotInput) *EBSStartSnapshotFuture
}

func NewClient() Client {
	return &stub{}
}
