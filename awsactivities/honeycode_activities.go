// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type HoneycodeActivities struct {
	client honeycodeiface.HoneycodeAPI

	sessionFactory SessionFactory
}

func NewHoneycodeActivities(sess *session.Session, config ...*aws.Config) *HoneycodeActivities {
	client := honeycode.New(sess, config...)
	return &HoneycodeActivities{client: client}
}

func NewHoneycodeActivitiesWithSessionFactory(sessionFactory SessionFactory) *HoneycodeActivities {
	return &HoneycodeActivities{sessionFactory: sessionFactory}
}

func (a *HoneycodeActivities) getClient(ctx context.Context) (honeycodeiface.HoneycodeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return honeycode.New(sess), nil
}

func (a *HoneycodeActivities) GetScreenData(ctx context.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetScreenDataWithContext(ctx, input)
}

func (a *HoneycodeActivities) InvokeScreenAutomation(ctx context.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InvokeScreenAutomationWithContext(ctx, input)
}
