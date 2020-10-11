// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CognitoSyncActivities struct {
	client cognitosynciface.CognitoSyncAPI

	sessionFactory SessionFactory
}

func NewCognitoSyncActivities(sess *session.Session, config ...*aws.Config) *CognitoSyncActivities {
	client := cognitosync.New(sess, config...)
	return &CognitoSyncActivities{client: client}
}

func NewCognitoSyncActivitiesWithSessionFactory(sessionFactory SessionFactory) *CognitoSyncActivities {
	return &CognitoSyncActivities{sessionFactory: sessionFactory}
}

func (a *CognitoSyncActivities) getClient(ctx context.Context) (cognitosynciface.CognitoSyncAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cognitosync.New(sess), nil
}

func (a *CognitoSyncActivities) BulkPublish(ctx context.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BulkPublishWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DeleteDataset(ctx context.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeDataset(ctx context.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeIdentityPoolUsage(ctx context.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityPoolUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeIdentityUsage(ctx context.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetBulkPublishDetails(ctx context.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBulkPublishDetailsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetCognitoEvents(ctx context.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCognitoEventsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetIdentityPoolConfiguration(ctx context.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIdentityPoolConfigurationWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListDatasets(ctx context.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListIdentityPoolUsage(ctx context.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIdentityPoolUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListRecords(ctx context.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRecordsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) RegisterDevice(ctx context.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterDeviceWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SetCognitoEvents(ctx context.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetCognitoEventsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SetIdentityPoolConfiguration(ctx context.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetIdentityPoolConfigurationWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SubscribeToDataset(ctx context.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SubscribeToDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) UnsubscribeFromDataset(ctx context.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnsubscribeFromDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) UpdateRecords(ctx context.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRecordsWithContext(ctx, input)
}
