// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TranscribeStreamingServiceActivities struct {
	client transcribestreamingserviceiface.TranscribeStreamingServiceAPI

	sessionFactory SessionFactory
}

func NewTranscribeStreamingServiceActivities(sess *session.Session, config ...*aws.Config) *TranscribeStreamingServiceActivities {
	client := transcribestreamingservice.New(sess, config...)
	return &TranscribeStreamingServiceActivities{client: client}
}

func NewTranscribeStreamingServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *TranscribeStreamingServiceActivities {
	return &TranscribeStreamingServiceActivities{sessionFactory: sessionFactory}
}

func (a *TranscribeStreamingServiceActivities) getClient(ctx context.Context) (transcribestreamingserviceiface.TranscribeStreamingServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return transcribestreamingservice.New(sess), nil
}

func (a *TranscribeStreamingServiceActivities) StartStreamTranscription(ctx context.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartStreamTranscriptionWithContext(ctx, input)
}
