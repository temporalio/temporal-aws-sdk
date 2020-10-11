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

type Client interface {
	CreateGatewayRoute(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error)
	CreateGatewayRouteAsync(ctx workflow.Context, input *appmesh.CreateGatewayRouteInput) *AppMeshCreateGatewayRouteFuture

	CreateMesh(ctx workflow.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error)
	CreateMeshAsync(ctx workflow.Context, input *appmesh.CreateMeshInput) *AppMeshCreateMeshFuture

	CreateRoute(ctx workflow.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error)
	CreateRouteAsync(ctx workflow.Context, input *appmesh.CreateRouteInput) *AppMeshCreateRouteFuture

	CreateVirtualGateway(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error)
	CreateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.CreateVirtualGatewayInput) *AppMeshCreateVirtualGatewayFuture

	CreateVirtualNode(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error)
	CreateVirtualNodeAsync(ctx workflow.Context, input *appmesh.CreateVirtualNodeInput) *AppMeshCreateVirtualNodeFuture

	CreateVirtualRouter(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error)
	CreateVirtualRouterAsync(ctx workflow.Context, input *appmesh.CreateVirtualRouterInput) *AppMeshCreateVirtualRouterFuture

	CreateVirtualService(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error)
	CreateVirtualServiceAsync(ctx workflow.Context, input *appmesh.CreateVirtualServiceInput) *AppMeshCreateVirtualServiceFuture

	DeleteGatewayRoute(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error)
	DeleteGatewayRouteAsync(ctx workflow.Context, input *appmesh.DeleteGatewayRouteInput) *AppMeshDeleteGatewayRouteFuture

	DeleteMesh(ctx workflow.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error)
	DeleteMeshAsync(ctx workflow.Context, input *appmesh.DeleteMeshInput) *AppMeshDeleteMeshFuture

	DeleteRoute(ctx workflow.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error)
	DeleteRouteAsync(ctx workflow.Context, input *appmesh.DeleteRouteInput) *AppMeshDeleteRouteFuture

	DeleteVirtualGateway(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error)
	DeleteVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DeleteVirtualGatewayInput) *AppMeshDeleteVirtualGatewayFuture

	DeleteVirtualNode(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error)
	DeleteVirtualNodeAsync(ctx workflow.Context, input *appmesh.DeleteVirtualNodeInput) *AppMeshDeleteVirtualNodeFuture

	DeleteVirtualRouter(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error)
	DeleteVirtualRouterAsync(ctx workflow.Context, input *appmesh.DeleteVirtualRouterInput) *AppMeshDeleteVirtualRouterFuture

	DeleteVirtualService(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error)
	DeleteVirtualServiceAsync(ctx workflow.Context, input *appmesh.DeleteVirtualServiceInput) *AppMeshDeleteVirtualServiceFuture

	DescribeGatewayRoute(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error)
	DescribeGatewayRouteAsync(ctx workflow.Context, input *appmesh.DescribeGatewayRouteInput) *AppMeshDescribeGatewayRouteFuture

	DescribeMesh(ctx workflow.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error)
	DescribeMeshAsync(ctx workflow.Context, input *appmesh.DescribeMeshInput) *AppMeshDescribeMeshFuture

	DescribeRoute(ctx workflow.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error)
	DescribeRouteAsync(ctx workflow.Context, input *appmesh.DescribeRouteInput) *AppMeshDescribeRouteFuture

	DescribeVirtualGateway(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error)
	DescribeVirtualGatewayAsync(ctx workflow.Context, input *appmesh.DescribeVirtualGatewayInput) *AppMeshDescribeVirtualGatewayFuture

	DescribeVirtualNode(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error)
	DescribeVirtualNodeAsync(ctx workflow.Context, input *appmesh.DescribeVirtualNodeInput) *AppMeshDescribeVirtualNodeFuture

	DescribeVirtualRouter(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error)
	DescribeVirtualRouterAsync(ctx workflow.Context, input *appmesh.DescribeVirtualRouterInput) *AppMeshDescribeVirtualRouterFuture

	DescribeVirtualService(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error)
	DescribeVirtualServiceAsync(ctx workflow.Context, input *appmesh.DescribeVirtualServiceInput) *AppMeshDescribeVirtualServiceFuture

	ListGatewayRoutes(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error)
	ListGatewayRoutesAsync(ctx workflow.Context, input *appmesh.ListGatewayRoutesInput) *AppMeshListGatewayRoutesFuture

	ListMeshes(ctx workflow.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error)
	ListMeshesAsync(ctx workflow.Context, input *appmesh.ListMeshesInput) *AppMeshListMeshesFuture

	ListRoutes(ctx workflow.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error)
	ListRoutesAsync(ctx workflow.Context, input *appmesh.ListRoutesInput) *AppMeshListRoutesFuture

	ListTagsForResource(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *appmesh.ListTagsForResourceInput) *AppMeshListTagsForResourceFuture

	ListVirtualGateways(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error)
	ListVirtualGatewaysAsync(ctx workflow.Context, input *appmesh.ListVirtualGatewaysInput) *AppMeshListVirtualGatewaysFuture

	ListVirtualNodes(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error)
	ListVirtualNodesAsync(ctx workflow.Context, input *appmesh.ListVirtualNodesInput) *AppMeshListVirtualNodesFuture

	ListVirtualRouters(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error)
	ListVirtualRoutersAsync(ctx workflow.Context, input *appmesh.ListVirtualRoutersInput) *AppMeshListVirtualRoutersFuture

	ListVirtualServices(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error)
	ListVirtualServicesAsync(ctx workflow.Context, input *appmesh.ListVirtualServicesInput) *AppMeshListVirtualServicesFuture

	TagResource(ctx workflow.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *appmesh.TagResourceInput) *AppMeshTagResourceFuture

	UntagResource(ctx workflow.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *appmesh.UntagResourceInput) *AppMeshUntagResourceFuture

	UpdateGatewayRoute(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error)
	UpdateGatewayRouteAsync(ctx workflow.Context, input *appmesh.UpdateGatewayRouteInput) *AppMeshUpdateGatewayRouteFuture

	UpdateMesh(ctx workflow.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error)
	UpdateMeshAsync(ctx workflow.Context, input *appmesh.UpdateMeshInput) *AppMeshUpdateMeshFuture

	UpdateRoute(ctx workflow.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error)
	UpdateRouteAsync(ctx workflow.Context, input *appmesh.UpdateRouteInput) *AppMeshUpdateRouteFuture

	UpdateVirtualGateway(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error)
	UpdateVirtualGatewayAsync(ctx workflow.Context, input *appmesh.UpdateVirtualGatewayInput) *AppMeshUpdateVirtualGatewayFuture

	UpdateVirtualNode(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error)
	UpdateVirtualNodeAsync(ctx workflow.Context, input *appmesh.UpdateVirtualNodeInput) *AppMeshUpdateVirtualNodeFuture

	UpdateVirtualRouter(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error)
	UpdateVirtualRouterAsync(ctx workflow.Context, input *appmesh.UpdateVirtualRouterInput) *AppMeshUpdateVirtualRouterFuture

	UpdateVirtualService(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error)
	UpdateVirtualServiceAsync(ctx workflow.Context, input *appmesh.UpdateVirtualServiceInput) *AppMeshUpdateVirtualServiceFuture
}

func NewClient() Client {
	return &stub{}
}

