// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client elastictranscoderiface.ElasticTranscoderAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elastictranscoder.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elastictranscoderiface.ElasticTranscoderAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return elastictranscoder.New(sess), nil
}

func (a *Activities) CancelJob(ctx context.Context, input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateJob(ctx context.Context, input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePipeline(ctx context.Context, input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePreset(ctx context.Context, input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePipeline(ctx context.Context, input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePreset(ctx context.Context, input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobsByPipeline(ctx context.Context, input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsByPipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobsByStatus(ctx context.Context, input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsByStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPipelines(ctx context.Context, input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPipelinesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPresets(ctx context.Context, input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPresetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ReadJob(ctx context.Context, input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ReadJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ReadPipeline(ctx context.Context, input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ReadPipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ReadPreset(ctx context.Context, input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ReadPresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TestRole(ctx context.Context, input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TestRoleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePipeline(ctx context.Context, input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePipelineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePipelineNotifications(ctx context.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePipelineNotificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePipelineStatus(ctx context.Context, input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePipelineStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilJobComplete(ctx context.Context, input *elastictranscoder.ReadJobInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilJobCompleteWithContext(ctx, input, options...))
	})
}
