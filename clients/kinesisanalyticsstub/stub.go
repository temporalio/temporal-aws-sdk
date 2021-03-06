// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisanalyticsstub

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddApplicationCloudWatchLoggingOptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddApplicationInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddApplicationInputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddApplicationInputProcessingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddApplicationOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddApplicationReferenceDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteApplicationCloudWatchLoggingOptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteApplicationInputProcessingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteApplicationOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteApplicationReferenceDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DiscoverInputSchemaFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DiscoverInputSchemaFuture) Get(ctx workflow.Context) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListApplicationsFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *AddApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input)
	return &AddApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *stub) AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *AddApplicationInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input)
	return &AddApplicationInputFuture{Future: future}
}

func (a *stub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *AddApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input)
	return &AddApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *stub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *AddApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input)
	return &AddApplicationOutputFuture{Future: future}
}

func (a *stub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *AddApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input)
	return &AddApplicationReferenceDataSourceFuture{Future: future}
}

func (a *stub) CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *CreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input)
	return &CreateApplicationFuture{Future: future}
}

func (a *stub) DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *DeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input)
	return &DeleteApplicationFuture{Future: future}
}

func (a *stub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *DeleteApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input)
	return &DeleteApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *stub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *DeleteApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input)
	return &DeleteApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *stub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *DeleteApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input)
	return &DeleteApplicationOutputFuture{Future: future}
}

func (a *stub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *DeleteApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input)
	return &DeleteApplicationReferenceDataSourceFuture{Future: future}
}

func (a *stub) DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *DescribeApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input)
	return &DescribeApplicationFuture{Future: future}
}

func (a *stub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *DiscoverInputSchemaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input)
	return &DiscoverInputSchemaFuture{Future: future}
}

func (a *stub) ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *ListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input)
	return &ListApplicationsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *StartApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input)
	return &StartApplicationFuture{Future: future}
}

func (a *stub) StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *StopApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input)
	return &StopApplicationFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *UpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input)
	return &UpdateApplicationFuture{Future: future}
}
