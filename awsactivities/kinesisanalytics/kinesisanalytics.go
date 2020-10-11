// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisanalytics

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client kinesisanalyticsiface.KinesisAnalyticsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := kinesisanalytics.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (kinesisanalyticsiface.KinesisAnalyticsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return kinesisanalytics.New(sess), nil
}

func (a *Activities) AddApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *Activities) AddApplicationInput(ctx context.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationInputWithContext(ctx, input)
}

func (a *Activities) AddApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *Activities) AddApplicationOutput(ctx context.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationOutputWithContext(ctx, input)
}

func (a *Activities) AddApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *Activities) CreateApplication(ctx context.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *Activities) DeleteApplication(ctx context.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationWithContext(ctx, input)
}

func (a *Activities) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *Activities) DeleteApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *Activities) DeleteApplicationOutput(ctx context.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationOutputWithContext(ctx, input)
}

func (a *Activities) DeleteApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *Activities) DescribeApplication(ctx context.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeApplicationWithContext(ctx, input)
}

func (a *Activities) DiscoverInputSchema(ctx context.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DiscoverInputSchemaWithContext(ctx, input)
}

func (a *Activities) ListApplications(ctx context.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) StartApplication(ctx context.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartApplicationWithContext(ctx, input)
}

func (a *Activities) StopApplication(ctx context.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopApplicationWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}