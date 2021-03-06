// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package daxstub

import (
	"github.com/aws/aws-sdk-go/service/dax"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *CreateClusterFuture

	CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error)
	CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *CreateParameterGroupFuture

	CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error)
	CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *CreateSubnetGroupFuture

	DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error)
	DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DecreaseReplicationFactorFuture

	DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DeleteClusterFuture

	DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error)
	DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DeleteParameterGroupFuture

	DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error)
	DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DeleteSubnetGroupFuture

	DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DescribeClustersFuture

	DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error)
	DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DescribeDefaultParametersFuture

	DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DescribeEventsFuture

	DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DescribeParameterGroupsFuture

	DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error)
	DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DescribeParametersFuture

	DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error)
	DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DescribeSubnetGroupsFuture

	IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error)
	IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *IncreaseReplicationFactorFuture

	ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *ListTagsFuture

	RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error)
	RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *RebootNodeFuture

	TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *UntagResourceFuture

	UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error)
	UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *UpdateClusterFuture

	UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error)
	UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *UpdateParameterGroupFuture

	UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error)
	UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *UpdateSubnetGroupFuture
}

func NewClient() Client {
	return &stub{}
}
