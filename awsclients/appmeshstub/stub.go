// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appmeshstub

import (
	"github.com/aws/aws-sdk-go/service/appmesh"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type AppMeshCreateGatewayRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateGatewayRouteFuture) Get(ctx workflow.Context) (*appmesh.CreateGatewayRouteOutput, error) {
	var output appmesh.CreateGatewayRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateMeshFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateMeshFuture) Get(ctx workflow.Context) (*appmesh.CreateMeshOutput, error) {
	var output appmesh.CreateMeshOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateRouteFuture) Get(ctx workflow.Context) (*appmesh.CreateRouteOutput, error) {
	var output appmesh.CreateRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateVirtualGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateVirtualGatewayFuture) Get(ctx workflow.Context) (*appmesh.CreateVirtualGatewayOutput, error) {
	var output appmesh.CreateVirtualGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateVirtualNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateVirtualNodeFuture) Get(ctx workflow.Context) (*appmesh.CreateVirtualNodeOutput, error) {
	var output appmesh.CreateVirtualNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateVirtualRouterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateVirtualRouterFuture) Get(ctx workflow.Context) (*appmesh.CreateVirtualRouterOutput, error) {
	var output appmesh.CreateVirtualRouterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshCreateVirtualServiceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshCreateVirtualServiceFuture) Get(ctx workflow.Context) (*appmesh.CreateVirtualServiceOutput, error) {
	var output appmesh.CreateVirtualServiceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteGatewayRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteGatewayRouteFuture) Get(ctx workflow.Context) (*appmesh.DeleteGatewayRouteOutput, error) {
	var output appmesh.DeleteGatewayRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteMeshFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteMeshFuture) Get(ctx workflow.Context) (*appmesh.DeleteMeshOutput, error) {
	var output appmesh.DeleteMeshOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteRouteFuture) Get(ctx workflow.Context) (*appmesh.DeleteRouteOutput, error) {
	var output appmesh.DeleteRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteVirtualGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteVirtualGatewayFuture) Get(ctx workflow.Context) (*appmesh.DeleteVirtualGatewayOutput, error) {
	var output appmesh.DeleteVirtualGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteVirtualNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteVirtualNodeFuture) Get(ctx workflow.Context) (*appmesh.DeleteVirtualNodeOutput, error) {
	var output appmesh.DeleteVirtualNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteVirtualRouterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteVirtualRouterFuture) Get(ctx workflow.Context) (*appmesh.DeleteVirtualRouterOutput, error) {
	var output appmesh.DeleteVirtualRouterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDeleteVirtualServiceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDeleteVirtualServiceFuture) Get(ctx workflow.Context) (*appmesh.DeleteVirtualServiceOutput, error) {
	var output appmesh.DeleteVirtualServiceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeGatewayRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeGatewayRouteFuture) Get(ctx workflow.Context) (*appmesh.DescribeGatewayRouteOutput, error) {
	var output appmesh.DescribeGatewayRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeMeshFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeMeshFuture) Get(ctx workflow.Context) (*appmesh.DescribeMeshOutput, error) {
	var output appmesh.DescribeMeshOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeRouteFuture) Get(ctx workflow.Context) (*appmesh.DescribeRouteOutput, error) {
	var output appmesh.DescribeRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeVirtualGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeVirtualGatewayFuture) Get(ctx workflow.Context) (*appmesh.DescribeVirtualGatewayOutput, error) {
	var output appmesh.DescribeVirtualGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeVirtualNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeVirtualNodeFuture) Get(ctx workflow.Context) (*appmesh.DescribeVirtualNodeOutput, error) {
	var output appmesh.DescribeVirtualNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeVirtualRouterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeVirtualRouterFuture) Get(ctx workflow.Context) (*appmesh.DescribeVirtualRouterOutput, error) {
	var output appmesh.DescribeVirtualRouterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshDescribeVirtualServiceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshDescribeVirtualServiceFuture) Get(ctx workflow.Context) (*appmesh.DescribeVirtualServiceOutput, error) {
	var output appmesh.DescribeVirtualServiceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListGatewayRoutesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListGatewayRoutesFuture) Get(ctx workflow.Context) (*appmesh.ListGatewayRoutesOutput, error) {
	var output appmesh.ListGatewayRoutesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListMeshesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListMeshesFuture) Get(ctx workflow.Context) (*appmesh.ListMeshesOutput, error) {
	var output appmesh.ListMeshesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListRoutesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListRoutesFuture) Get(ctx workflow.Context) (*appmesh.ListRoutesOutput, error) {
	var output appmesh.ListRoutesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListTagsForResourceFuture) Get(ctx workflow.Context) (*appmesh.ListTagsForResourceOutput, error) {
	var output appmesh.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListVirtualGatewaysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListVirtualGatewaysFuture) Get(ctx workflow.Context) (*appmesh.ListVirtualGatewaysOutput, error) {
	var output appmesh.ListVirtualGatewaysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListVirtualNodesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListVirtualNodesFuture) Get(ctx workflow.Context) (*appmesh.ListVirtualNodesOutput, error) {
	var output appmesh.ListVirtualNodesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListVirtualRoutersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListVirtualRoutersFuture) Get(ctx workflow.Context) (*appmesh.ListVirtualRoutersOutput, error) {
	var output appmesh.ListVirtualRoutersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshListVirtualServicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshListVirtualServicesFuture) Get(ctx workflow.Context) (*appmesh.ListVirtualServicesOutput, error) {
	var output appmesh.ListVirtualServicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshTagResourceFuture) Get(ctx workflow.Context) (*appmesh.TagResourceOutput, error) {
	var output appmesh.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUntagResourceFuture) Get(ctx workflow.Context) (*appmesh.UntagResourceOutput, error) {
	var output appmesh.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateGatewayRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateGatewayRouteFuture) Get(ctx workflow.Context) (*appmesh.UpdateGatewayRouteOutput, error) {
	var output appmesh.UpdateGatewayRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateMeshFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateMeshFuture) Get(ctx workflow.Context) (*appmesh.UpdateMeshOutput, error) {
	var output appmesh.UpdateMeshOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateRouteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateRouteFuture) Get(ctx workflow.Context) (*appmesh.UpdateRouteOutput, error) {
	var output appmesh.UpdateRouteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateVirtualGatewayFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateVirtualGatewayFuture) Get(ctx workflow.Context) (*appmesh.UpdateVirtualGatewayOutput, error) {
	var output appmesh.UpdateVirtualGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateVirtualNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateVirtualNodeFuture) Get(ctx workflow.Context) (*appmesh.UpdateVirtualNodeOutput, error) {
	var output appmesh.UpdateVirtualNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateVirtualRouterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateVirtualRouterFuture) Get(ctx workflow.Context) (*appmesh.UpdateVirtualRouterOutput, error) {
	var output appmesh.UpdateVirtualRouterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppMeshUpdateVirtualServiceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppMeshUpdateVirtualServiceFuture) Get(ctx workflow.Context) (*appmesh.UpdateVirtualServiceOutput, error) {
	var output appmesh.UpdateVirtualServiceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGatewayRoute(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
	var output appmesh.CreateGatewayRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateGatewayRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGatewayRouteAsync(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) *AppMeshCreateGatewayRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateGatewayRoute", input)
	return &AppMeshCreateGatewayRouteFuture{Future: future}
}

func (a *stub) CreateMesh(ctx workflow.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
	var output appmesh.CreateMeshOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateMesh", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMeshAsync(ctx workflow.Context, input *appmesh.CreateMeshInput) *AppMeshCreateMeshFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateMesh", input)
	return &AppMeshCreateMeshFuture{Future: future}
}

func (a *stub) CreateRoute(ctx workflow.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
	var output appmesh.CreateRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRouteAsync(ctx workflow.Context, input *appmesh.CreateRouteInput) *AppMeshCreateRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateRoute", input)
	return &AppMeshCreateRouteFuture{Future: future}
}

func (a *stub) CreateVirtualGateway(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
	var output appmesh.CreateVirtualGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) *AppMeshCreateVirtualGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualGateway", input)
	return &AppMeshCreateVirtualGatewayFuture{Future: future}
}

func (a *stub) CreateVirtualNode(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
	var output appmesh.CreateVirtualNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVirtualNodeAsync(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) *AppMeshCreateVirtualNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualNode", input)
	return &AppMeshCreateVirtualNodeFuture{Future: future}
}

func (a *stub) CreateVirtualRouter(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
	var output appmesh.CreateVirtualRouterOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualRouter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVirtualRouterAsync(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) *AppMeshCreateVirtualRouterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualRouter", input)
	return &AppMeshCreateVirtualRouterFuture{Future: future}
}

func (a *stub) CreateVirtualService(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
	var output appmesh.CreateVirtualServiceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualService", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVirtualServiceAsync(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) *AppMeshCreateVirtualServiceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.CreateVirtualService", input)
	return &AppMeshCreateVirtualServiceFuture{Future: future}
}

func (a *stub) DeleteGatewayRoute(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
	var output appmesh.DeleteGatewayRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteGatewayRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteGatewayRouteAsync(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) *AppMeshDeleteGatewayRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteGatewayRoute", input)
	return &AppMeshDeleteGatewayRouteFuture{Future: future}
}

func (a *stub) DeleteMesh(ctx workflow.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
	var output appmesh.DeleteMeshOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteMesh", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMeshAsync(ctx workflow.Context, input *appmesh.DeleteMeshInput) *AppMeshDeleteMeshFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteMesh", input)
	return &AppMeshDeleteMeshFuture{Future: future}
}

func (a *stub) DeleteRoute(ctx workflow.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
	var output appmesh.DeleteRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRouteAsync(ctx workflow.Context, input *appmesh.DeleteRouteInput) *AppMeshDeleteRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteRoute", input)
	return &AppMeshDeleteRouteFuture{Future: future}
}

func (a *stub) DeleteVirtualGateway(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
	var output appmesh.DeleteVirtualGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) *AppMeshDeleteVirtualGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualGateway", input)
	return &AppMeshDeleteVirtualGatewayFuture{Future: future}
}

func (a *stub) DeleteVirtualNode(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
	var output appmesh.DeleteVirtualNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVirtualNodeAsync(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) *AppMeshDeleteVirtualNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualNode", input)
	return &AppMeshDeleteVirtualNodeFuture{Future: future}
}

func (a *stub) DeleteVirtualRouter(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
	var output appmesh.DeleteVirtualRouterOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualRouter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVirtualRouterAsync(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) *AppMeshDeleteVirtualRouterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualRouter", input)
	return &AppMeshDeleteVirtualRouterFuture{Future: future}
}

func (a *stub) DeleteVirtualService(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
	var output appmesh.DeleteVirtualServiceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualService", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVirtualServiceAsync(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) *AppMeshDeleteVirtualServiceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DeleteVirtualService", input)
	return &AppMeshDeleteVirtualServiceFuture{Future: future}
}

func (a *stub) DescribeGatewayRoute(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
	var output appmesh.DescribeGatewayRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeGatewayRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGatewayRouteAsync(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) *AppMeshDescribeGatewayRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeGatewayRoute", input)
	return &AppMeshDescribeGatewayRouteFuture{Future: future}
}

func (a *stub) DescribeMesh(ctx workflow.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
	var output appmesh.DescribeMeshOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeMesh", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeMeshAsync(ctx workflow.Context, input *appmesh.DescribeMeshInput) *AppMeshDescribeMeshFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeMesh", input)
	return &AppMeshDescribeMeshFuture{Future: future}
}

func (a *stub) DescribeRoute(ctx workflow.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
	var output appmesh.DescribeRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRouteAsync(ctx workflow.Context, input *appmesh.DescribeRouteInput) *AppMeshDescribeRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeRoute", input)
	return &AppMeshDescribeRouteFuture{Future: future}
}

func (a *stub) DescribeVirtualGateway(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
	var output appmesh.DescribeVirtualGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) *AppMeshDescribeVirtualGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualGateway", input)
	return &AppMeshDescribeVirtualGatewayFuture{Future: future}
}

func (a *stub) DescribeVirtualNode(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
	var output appmesh.DescribeVirtualNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeVirtualNodeAsync(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) *AppMeshDescribeVirtualNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualNode", input)
	return &AppMeshDescribeVirtualNodeFuture{Future: future}
}

func (a *stub) DescribeVirtualRouter(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
	var output appmesh.DescribeVirtualRouterOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualRouter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeVirtualRouterAsync(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) *AppMeshDescribeVirtualRouterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualRouter", input)
	return &AppMeshDescribeVirtualRouterFuture{Future: future}
}

func (a *stub) DescribeVirtualService(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
	var output appmesh.DescribeVirtualServiceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualService", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeVirtualServiceAsync(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) *AppMeshDescribeVirtualServiceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.DescribeVirtualService", input)
	return &AppMeshDescribeVirtualServiceFuture{Future: future}
}

func (a *stub) ListGatewayRoutes(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
	var output appmesh.ListGatewayRoutesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListGatewayRoutes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGatewayRoutesAsync(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) *AppMeshListGatewayRoutesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListGatewayRoutes", input)
	return &AppMeshListGatewayRoutesFuture{Future: future}
}

func (a *stub) ListMeshes(ctx workflow.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	var output appmesh.ListMeshesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListMeshes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMeshesAsync(ctx workflow.Context, input *appmesh.ListMeshesInput) *AppMeshListMeshesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListMeshes", input)
	return &AppMeshListMeshesFuture{Future: future}
}

func (a *stub) ListRoutes(ctx workflow.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
	var output appmesh.ListRoutesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListRoutes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRoutesAsync(ctx workflow.Context, input *appmesh.ListRoutesInput) *AppMeshListRoutesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListRoutes", input)
	return &AppMeshListRoutesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	var output appmesh.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) *AppMeshListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListTagsForResource", input)
	return &AppMeshListTagsForResourceFuture{Future: future}
}

func (a *stub) ListVirtualGateways(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
	var output appmesh.ListVirtualGatewaysOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualGateways", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVirtualGatewaysAsync(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) *AppMeshListVirtualGatewaysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualGateways", input)
	return &AppMeshListVirtualGatewaysFuture{Future: future}
}

func (a *stub) ListVirtualNodes(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
	var output appmesh.ListVirtualNodesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualNodes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVirtualNodesAsync(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) *AppMeshListVirtualNodesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualNodes", input)
	return &AppMeshListVirtualNodesFuture{Future: future}
}

func (a *stub) ListVirtualRouters(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
	var output appmesh.ListVirtualRoutersOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualRouters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVirtualRoutersAsync(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) *AppMeshListVirtualRoutersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualRouters", input)
	return &AppMeshListVirtualRoutersFuture{Future: future}
}

func (a *stub) ListVirtualServices(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
	var output appmesh.ListVirtualServicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualServices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVirtualServicesAsync(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) *AppMeshListVirtualServicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.ListVirtualServices", input)
	return &AppMeshListVirtualServicesFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
	var output appmesh.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *appmesh.TagResourceInput) *AppMeshTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.TagResource", input)
	return &AppMeshTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
	var output appmesh.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *appmesh.UntagResourceInput) *AppMeshUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UntagResource", input)
	return &AppMeshUntagResourceFuture{Future: future}
}

func (a *stub) UpdateGatewayRoute(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
	var output appmesh.UpdateGatewayRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateGatewayRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateGatewayRouteAsync(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) *AppMeshUpdateGatewayRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateGatewayRoute", input)
	return &AppMeshUpdateGatewayRouteFuture{Future: future}
}

func (a *stub) UpdateMesh(ctx workflow.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
	var output appmesh.UpdateMeshOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateMesh", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateMeshAsync(ctx workflow.Context, input *appmesh.UpdateMeshInput) *AppMeshUpdateMeshFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateMesh", input)
	return &AppMeshUpdateMeshFuture{Future: future}
}

func (a *stub) UpdateRoute(ctx workflow.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
	var output appmesh.UpdateRouteOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateRoute", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRouteAsync(ctx workflow.Context, input *appmesh.UpdateRouteInput) *AppMeshUpdateRouteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateRoute", input)
	return &AppMeshUpdateRouteFuture{Future: future}
}

func (a *stub) UpdateVirtualGateway(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
	var output appmesh.UpdateVirtualGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) *AppMeshUpdateVirtualGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualGateway", input)
	return &AppMeshUpdateVirtualGatewayFuture{Future: future}
}

func (a *stub) UpdateVirtualNode(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
	var output appmesh.UpdateVirtualNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVirtualNodeAsync(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) *AppMeshUpdateVirtualNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualNode", input)
	return &AppMeshUpdateVirtualNodeFuture{Future: future}
}

func (a *stub) UpdateVirtualRouter(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
	var output appmesh.UpdateVirtualRouterOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualRouter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVirtualRouterAsync(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) *AppMeshUpdateVirtualRouterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualRouter", input)
	return &AppMeshUpdateVirtualRouterFuture{Future: future}
}

func (a *stub) UpdateVirtualService(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
	var output appmesh.UpdateVirtualServiceOutput
	err := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualService", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVirtualServiceAsync(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) *AppMeshUpdateVirtualServiceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.appmesh.UpdateVirtualService", input)
	return &AppMeshUpdateVirtualServiceFuture{Future: future}
}
