// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudsearchdomainstub

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SearchFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SearchFuture) Get(ctx workflow.Context) (*cloudsearchdomain.SearchOutput, error) {
	var output cloudsearchdomain.SearchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SuggestFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SuggestFuture) Get(ctx workflow.Context) (*cloudsearchdomain.SuggestOutput, error) {
	var output cloudsearchdomain.SuggestOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UploadDocumentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UploadDocumentsFuture) Get(ctx workflow.Context) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	var output cloudsearchdomain.UploadDocumentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	var output cloudsearchdomain.SearchOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.Search", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *SearchFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.Search", input)
	return &SearchFuture{Future: future}
}

func (a *stub) Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
	var output cloudsearchdomain.SuggestOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.Suggest", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *SuggestFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.Suggest", input)
	return &SuggestFuture{Future: future}
}

func (a *stub) UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	var output cloudsearchdomain.UploadDocumentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.UploadDocuments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *UploadDocumentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudsearchdomain.UploadDocuments", input)
	return &UploadDocumentsFuture{Future: future}
}
