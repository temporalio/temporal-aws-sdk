// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoT1ClickProjectsActivities struct {
	client iot1clickprojectsiface.IoT1ClickProjectsAPI

	sessionFactory SessionFactory
}

func NewIoT1ClickProjectsActivities(sess *session.Session, config ...*aws.Config) *IoT1ClickProjectsActivities {
	client := iot1clickprojects.New(sess, config...)
	return &IoT1ClickProjectsActivities{client: client}
}

func NewIoT1ClickProjectsActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoT1ClickProjectsActivities {
	return &IoT1ClickProjectsActivities{sessionFactory: sessionFactory}
}

func (a *IoT1ClickProjectsActivities) getClient(ctx context.Context) (iot1clickprojectsiface.IoT1ClickProjectsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iot1clickprojects.New(sess), nil
}

func (a *IoT1ClickProjectsActivities) AssociateDeviceWithPlacement(ctx context.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateDeviceWithPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) CreatePlacement(ctx context.Context, input *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) CreateProject(ctx context.Context, input *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DeletePlacement(ctx context.Context, input *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DeleteProject(ctx context.Context, input *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DescribePlacement(ctx context.Context, input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DescribeProject(ctx context.Context, input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProjectWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) DisassociateDeviceFromPlacement(ctx context.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateDeviceFromPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) GetDevicesInPlacement(ctx context.Context, input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDevicesInPlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListPlacements(ctx context.Context, input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPlacementsWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListProjects(ctx context.Context, input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectsWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) ListTagsForResource(ctx context.Context, input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) TagResource(ctx context.Context, input *iot1clickprojects.TagResourceInput) (*iot1clickprojects.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UntagResource(ctx context.Context, input *iot1clickprojects.UntagResourceInput) (*iot1clickprojects.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UpdatePlacement(ctx context.Context, input *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePlacementWithContext(ctx, input)
}

func (a *IoT1ClickProjectsActivities) UpdateProject(ctx context.Context, input *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProjectWithContext(ctx, input)
}
