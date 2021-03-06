// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudhsmv2stub

import (
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CopyBackupToRegionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CopyBackupToRegionFuture) Get(ctx workflow.Context) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	var output cloudhsmv2.CopyBackupToRegionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.CreateClusterOutput, error) {
	var output cloudhsmv2.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHsmFuture) Get(ctx workflow.Context) (*cloudhsmv2.CreateHsmOutput, error) {
	var output cloudhsmv2.CreateHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBackupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBackupFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteBackupOutput, error) {
	var output cloudhsmv2.DeleteBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteClusterOutput, error) {
	var output cloudhsmv2.DeleteClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteHsmFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteHsmOutput, error) {
	var output cloudhsmv2.DeleteHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBackupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBackupsFuture) Get(ctx workflow.Context) (*cloudhsmv2.DescribeBackupsOutput, error) {
	var output cloudhsmv2.DescribeBackupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeClustersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeClustersFuture) Get(ctx workflow.Context) (*cloudhsmv2.DescribeClustersOutput, error) {
	var output cloudhsmv2.DescribeClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InitializeClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InitializeClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.InitializeClusterOutput, error) {
	var output cloudhsmv2.InitializeClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsFuture) Get(ctx workflow.Context) (*cloudhsmv2.ListTagsOutput, error) {
	var output cloudhsmv2.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RestoreBackupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RestoreBackupFuture) Get(ctx workflow.Context) (*cloudhsmv2.RestoreBackupOutput, error) {
	var output cloudhsmv2.RestoreBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*cloudhsmv2.TagResourceOutput, error) {
	var output cloudhsmv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*cloudhsmv2.UntagResourceOutput, error) {
	var output cloudhsmv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	var output cloudhsmv2.CopyBackupToRegionOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CopyBackupToRegion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *CopyBackupToRegionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CopyBackupToRegion", input)
	return &CopyBackupToRegionFuture{Future: future}
}

func (a *stub) CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
	var output cloudhsmv2.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *CreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateCluster", input)
	return &CreateClusterFuture{Future: future}
}

func (a *stub) CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
	var output cloudhsmv2.CreateHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *CreateHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateHsm", input)
	return &CreateHsmFuture{Future: future}
}

func (a *stub) DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
	var output cloudhsmv2.DeleteBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *DeleteBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteBackup", input)
	return &DeleteBackupFuture{Future: future}
}

func (a *stub) DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
	var output cloudhsmv2.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *DeleteClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteCluster", input)
	return &DeleteClusterFuture{Future: future}
}

func (a *stub) DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
	var output cloudhsmv2.DeleteHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *DeleteHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteHsm", input)
	return &DeleteHsmFuture{Future: future}
}

func (a *stub) DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
	var output cloudhsmv2.DescribeBackupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeBackups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *DescribeBackupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeBackups", input)
	return &DescribeBackupsFuture{Future: future}
}

func (a *stub) DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
	var output cloudhsmv2.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *DescribeClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeClusters", input)
	return &DescribeClustersFuture{Future: future}
}

func (a *stub) InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
	var output cloudhsmv2.InitializeClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.InitializeCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *InitializeClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.InitializeCluster", input)
	return &InitializeClusterFuture{Future: future}
}

func (a *stub) ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
	var output cloudhsmv2.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *ListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.ListTags", input)
	return &ListTagsFuture{Future: future}
}

func (a *stub) RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
	var output cloudhsmv2.RestoreBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.RestoreBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *RestoreBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.RestoreBackup", input)
	return &RestoreBackupFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
	var output cloudhsmv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
	var output cloudhsmv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
