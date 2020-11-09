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

type stub struct{}

type CreateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateClusterFuture) Get(ctx workflow.Context) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateParameterGroupFuture) Get(ctx workflow.Context) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DecreaseReplicationFactorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DecreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteClusterFuture) Get(ctx workflow.Context) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteParameterGroupFuture) Get(ctx workflow.Context) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSubnetGroupFuture) Get(ctx workflow.Context) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeClustersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeClustersFuture) Get(ctx workflow.Context) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDefaultParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDefaultParametersFuture) Get(ctx workflow.Context) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEventsFuture) Get(ctx workflow.Context) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeParameterGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeParameterGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeParametersFuture) Get(ctx workflow.Context) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSubnetGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSubnetGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IncreaseReplicationFactorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IncreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsFuture) Get(ctx workflow.Context) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RebootNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RebootNodeFuture) Get(ctx workflow.Context) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateClusterFuture) Get(ctx workflow.Context) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateParameterGroupFuture) Get(ctx workflow.Context) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *CreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input)
	return &CreateClusterFuture{Future: future}
}

func (a *stub) CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *CreateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input)
	return &CreateParameterGroupFuture{Future: future}
}

func (a *stub) CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *CreateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input)
	return &CreateSubnetGroupFuture{Future: future}
}

func (a *stub) DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DecreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input)
	return &DecreaseReplicationFactorFuture{Future: future}
}

func (a *stub) DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DeleteClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input)
	return &DeleteClusterFuture{Future: future}
}

func (a *stub) DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DeleteParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input)
	return &DeleteParameterGroupFuture{Future: future}
}

func (a *stub) DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DeleteSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input)
	return &DeleteSubnetGroupFuture{Future: future}
}

func (a *stub) DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DescribeClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input)
	return &DescribeClustersFuture{Future: future}
}

func (a *stub) DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DescribeDefaultParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input)
	return &DescribeDefaultParametersFuture{Future: future}
}

func (a *stub) DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DescribeEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input)
	return &DescribeEventsFuture{Future: future}
}

func (a *stub) DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DescribeParameterGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input)
	return &DescribeParameterGroupsFuture{Future: future}
}

func (a *stub) DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DescribeParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input)
	return &DescribeParametersFuture{Future: future}
}

func (a *stub) DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DescribeSubnetGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input)
	return &DescribeSubnetGroupsFuture{Future: future}
}

func (a *stub) IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *IncreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input)
	return &IncreaseReplicationFactorFuture{Future: future}
}

func (a *stub) ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *ListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input)
	return &ListTagsFuture{Future: future}
}

func (a *stub) RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *RebootNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input)
	return &RebootNodeFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *UpdateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input)
	return &UpdateClusterFuture{Future: future}
}

func (a *stub) UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *UpdateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input)
	return &UpdateParameterGroupFuture{Future: future}
}

func (a *stub) UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *UpdateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input)
	return &UpdateSubnetGroupFuture{Future: future}
}
