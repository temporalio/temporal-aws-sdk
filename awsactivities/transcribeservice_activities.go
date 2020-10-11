// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TranscribeServiceActivities struct {
	client transcribeserviceiface.TranscribeServiceAPI

	sessionFactory SessionFactory
}

func NewTranscribeServiceActivities(sess *session.Session, config ...*aws.Config) *TranscribeServiceActivities {
	client := transcribeservice.New(sess, config...)
	return &TranscribeServiceActivities{client: client}
}

func NewTranscribeServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *TranscribeServiceActivities {
	return &TranscribeServiceActivities{sessionFactory: sessionFactory}
}

func (a *TranscribeServiceActivities) getClient(ctx context.Context) (transcribeserviceiface.TranscribeServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return transcribeservice.New(sess), nil
}

func (a *TranscribeServiceActivities) CreateLanguageModel(ctx context.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateMedicalVocabulary(ctx context.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateVocabulary(ctx context.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateVocabularyFilter(ctx context.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteLanguageModel(ctx context.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteMedicalVocabulary(ctx context.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteVocabulary(ctx context.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteVocabularyFilter(ctx context.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DescribeLanguageModel(ctx context.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetMedicalVocabulary(ctx context.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetTranscriptionJob(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetVocabulary(ctx context.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetVocabularyFilter(ctx context.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListLanguageModels(ctx context.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLanguageModelsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListMedicalTranscriptionJobs(ctx context.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMedicalTranscriptionJobsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListMedicalVocabularies(ctx context.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMedicalVocabulariesWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListTranscriptionJobs(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTranscriptionJobsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListVocabularies(ctx context.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVocabulariesWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListVocabularyFilters(ctx context.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVocabularyFiltersWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) StartMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) StartTranscriptionJob(ctx context.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateMedicalVocabulary(ctx context.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateVocabulary(ctx context.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateVocabularyFilter(ctx context.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVocabularyFilterWithContext(ctx, input)
}
