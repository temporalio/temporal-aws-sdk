// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SQSActivities struct {
	client sqsiface.SQSAPI

	sessionFactory SessionFactory
}

func NewSQSActivities(sess *session.Session, config ...*aws.Config) *SQSActivities {
	client := sqs.New(sess, config...)
	return &SQSActivities{client: client}
}

func NewSQSActivitiesWithSessionFactory(sessionFactory SessionFactory) *SQSActivities {
	return &SQSActivities{sessionFactory: sessionFactory}
}

func (a *SQSActivities) getClient(ctx context.Context) (sqsiface.SQSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return sqs.New(sess), nil
}

func (a *SQSActivities) AddPermission(ctx context.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddPermissionWithContext(ctx, input)
}

func (a *SQSActivities) ChangeMessageVisibility(ctx context.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ChangeMessageVisibilityWithContext(ctx, input)
}

func (a *SQSActivities) ChangeMessageVisibilityBatch(ctx context.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ChangeMessageVisibilityBatchWithContext(ctx, input)
}

func (a *SQSActivities) CreateQueue(ctx context.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateQueueWithContext(ctx, input)
}

func (a *SQSActivities) DeleteMessage(ctx context.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMessageWithContext(ctx, input)
}

func (a *SQSActivities) DeleteMessageBatch(ctx context.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMessageBatchWithContext(ctx, input)
}

func (a *SQSActivities) DeleteQueue(ctx context.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteQueueWithContext(ctx, input)
}

func (a *SQSActivities) GetQueueAttributes(ctx context.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetQueueAttributesWithContext(ctx, input)
}

func (a *SQSActivities) GetQueueUrl(ctx context.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetQueueUrlWithContext(ctx, input)
}

func (a *SQSActivities) ListDeadLetterSourceQueues(ctx context.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeadLetterSourceQueuesWithContext(ctx, input)
}

func (a *SQSActivities) ListQueueTags(ctx context.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListQueueTagsWithContext(ctx, input)
}

func (a *SQSActivities) ListQueues(ctx context.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListQueuesWithContext(ctx, input)
}

func (a *SQSActivities) PurgeQueue(ctx context.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PurgeQueueWithContext(ctx, input)
}

func (a *SQSActivities) ReceiveMessage(ctx context.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReceiveMessageWithContext(ctx, input)
}

func (a *SQSActivities) RemovePermission(ctx context.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemovePermissionWithContext(ctx, input)
}

func (a *SQSActivities) SendMessage(ctx context.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendMessageWithContext(ctx, input)
}

func (a *SQSActivities) SendMessageBatch(ctx context.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendMessageBatchWithContext(ctx, input)
}

func (a *SQSActivities) SetQueueAttributes(ctx context.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetQueueAttributesWithContext(ctx, input)
}

func (a *SQSActivities) TagQueue(ctx context.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagQueueWithContext(ctx, input)
}

func (a *SQSActivities) UntagQueue(ctx context.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagQueueWithContext(ctx, input)
}
