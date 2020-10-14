// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client appmeshiface.AppMeshAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := appmesh.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (appmeshiface.AppMeshAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return appmesh.New(sess), nil
}

func (a *Activities) CreateGatewayRoute(ctx context.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateGatewayRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMesh(ctx context.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateMeshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRoute(ctx context.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVirtualGateway(ctx context.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateVirtualGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVirtualNode(ctx context.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateVirtualNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVirtualRouter(ctx context.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateVirtualRouterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVirtualService(ctx context.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateVirtualServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGatewayRoute(ctx context.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGatewayRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMesh(ctx context.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMeshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRoute(ctx context.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVirtualGateway(ctx context.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVirtualGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVirtualNode(ctx context.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVirtualNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVirtualRouter(ctx context.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVirtualRouterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVirtualService(ctx context.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVirtualServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGatewayRoute(ctx context.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGatewayRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMesh(ctx context.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMeshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRoute(ctx context.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVirtualGateway(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVirtualGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVirtualNode(ctx context.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVirtualNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVirtualRouter(ctx context.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVirtualRouterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVirtualService(ctx context.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVirtualServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGatewayRoutes(ctx context.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGatewayRoutesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMeshes(ctx context.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMeshesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRoutes(ctx context.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRoutesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVirtualGateways(ctx context.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVirtualGatewaysWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVirtualNodes(ctx context.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVirtualNodesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVirtualRouters(ctx context.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVirtualRoutersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVirtualServices(ctx context.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVirtualServicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGatewayRoute(ctx context.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateGatewayRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMesh(ctx context.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateMeshWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoute(ctx context.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateRouteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVirtualGateway(ctx context.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateVirtualGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVirtualNode(ctx context.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateVirtualNodeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVirtualRouter(ctx context.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateVirtualRouterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVirtualService(ctx context.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.UpdateVirtualServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
