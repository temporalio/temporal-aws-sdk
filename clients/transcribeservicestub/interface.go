// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package transcribeservicestub

import (
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateLanguageModel(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error)
	CreateLanguageModelAsync(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) *CreateLanguageModelFuture

	CreateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error)
	CreateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) *CreateMedicalVocabularyFuture

	CreateVocabulary(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error)
	CreateVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) *CreateVocabularyFuture

	CreateVocabularyFilter(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error)
	CreateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) *CreateVocabularyFilterFuture

	DeleteLanguageModel(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error)
	DeleteLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) *DeleteLanguageModelFuture

	DeleteMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error)
	DeleteMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) *DeleteMedicalTranscriptionJobFuture

	DeleteMedicalVocabulary(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error)
	DeleteMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) *DeleteMedicalVocabularyFuture

	DeleteTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error)
	DeleteTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) *DeleteTranscriptionJobFuture

	DeleteVocabulary(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error)
	DeleteVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) *DeleteVocabularyFuture

	DeleteVocabularyFilter(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error)
	DeleteVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) *DeleteVocabularyFilterFuture

	DescribeLanguageModel(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error)
	DescribeLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) *DescribeLanguageModelFuture

	GetMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error)
	GetMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) *GetMedicalTranscriptionJobFuture

	GetMedicalVocabulary(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error)
	GetMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) *GetMedicalVocabularyFuture

	GetTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error)
	GetTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) *GetTranscriptionJobFuture

	GetVocabulary(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error)
	GetVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) *GetVocabularyFuture

	GetVocabularyFilter(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error)
	GetVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) *GetVocabularyFilterFuture

	ListLanguageModels(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error)
	ListLanguageModelsAsync(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) *ListLanguageModelsFuture

	ListMedicalTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error)
	ListMedicalTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) *ListMedicalTranscriptionJobsFuture

	ListMedicalVocabularies(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error)
	ListMedicalVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) *ListMedicalVocabulariesFuture

	ListTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error)
	ListTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) *ListTranscriptionJobsFuture

	ListVocabularies(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error)
	ListVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) *ListVocabulariesFuture

	ListVocabularyFilters(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error)
	ListVocabularyFiltersAsync(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) *ListVocabularyFiltersFuture

	StartMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error)
	StartMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) *StartMedicalTranscriptionJobFuture

	StartTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error)
	StartTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) *StartTranscriptionJobFuture

	UpdateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error)
	UpdateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) *UpdateMedicalVocabularyFuture

	UpdateVocabulary(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error)
	UpdateVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) *UpdateVocabularyFuture

	UpdateVocabularyFilter(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error)
	UpdateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) *UpdateVocabularyFilterFuture
}

func NewClient() Client {
	return &stub{}
}
