// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package forecastservice

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
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
	client forecastserviceiface.ForecastServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := forecastservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (forecastserviceiface.ForecastServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return forecastservice.New(sess), nil
}

func (a *Activities) CreateDataset(ctx context.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatasetWithContext(ctx, input)
}

func (a *Activities) CreateDatasetGroup(ctx context.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatasetGroupWithContext(ctx, input)
}

func (a *Activities) CreateDatasetImportJob(ctx context.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatasetImportJobWithContext(ctx, input)
}

func (a *Activities) CreateForecast(ctx context.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateForecastWithContext(ctx, input)
}

func (a *Activities) CreateForecastExportJob(ctx context.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateForecastExportJobWithContext(ctx, input)
}

func (a *Activities) CreatePredictor(ctx context.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePredictorWithContext(ctx, input)
}

func (a *Activities) DeleteDataset(ctx context.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetWithContext(ctx, input)
}

func (a *Activities) DeleteDatasetGroup(ctx context.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetGroupWithContext(ctx, input)
}

func (a *Activities) DeleteDatasetImportJob(ctx context.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetImportJobWithContext(ctx, input)
}

func (a *Activities) DeleteForecast(ctx context.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteForecastWithContext(ctx, input)
}

func (a *Activities) DeleteForecastExportJob(ctx context.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteForecastExportJobWithContext(ctx, input)
}

func (a *Activities) DeletePredictor(ctx context.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePredictorWithContext(ctx, input)
}

func (a *Activities) DescribeDataset(ctx context.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatasetWithContext(ctx, input)
}

func (a *Activities) DescribeDatasetGroup(ctx context.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatasetGroupWithContext(ctx, input)
}

func (a *Activities) DescribeDatasetImportJob(ctx context.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatasetImportJobWithContext(ctx, input)
}

func (a *Activities) DescribeForecast(ctx context.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeForecastWithContext(ctx, input)
}

func (a *Activities) DescribeForecastExportJob(ctx context.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeForecastExportJobWithContext(ctx, input)
}

func (a *Activities) DescribePredictor(ctx context.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePredictorWithContext(ctx, input)
}

func (a *Activities) GetAccuracyMetrics(ctx context.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccuracyMetricsWithContext(ctx, input)
}

func (a *Activities) ListDatasetGroups(ctx context.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetGroupsWithContext(ctx, input)
}

func (a *Activities) ListDatasetImportJobs(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetImportJobsWithContext(ctx, input)
}

func (a *Activities) ListDatasets(ctx context.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetsWithContext(ctx, input)
}

func (a *Activities) ListForecastExportJobs(ctx context.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListForecastExportJobsWithContext(ctx, input)
}

func (a *Activities) ListForecasts(ctx context.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListForecastsWithContext(ctx, input)
}

func (a *Activities) ListPredictors(ctx context.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPredictorsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateDatasetGroup(ctx context.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDatasetGroupWithContext(ctx, input)
}