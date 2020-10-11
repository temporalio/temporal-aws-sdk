// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package costexplorer

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
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
	client costexploreriface.CostExplorerAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := costexplorer.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (costexploreriface.CostExplorerAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return costexplorer.New(sess), nil
}

func (a *Activities) CreateAnomalyMonitor(ctx context.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAnomalyMonitorWithContext(ctx, input)
}

func (a *Activities) CreateAnomalySubscription(ctx context.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAnomalySubscriptionWithContext(ctx, input)
}

func (a *Activities) CreateCostCategoryDefinition(ctx context.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCostCategoryDefinitionWithContext(ctx, input)
}

func (a *Activities) DeleteAnomalyMonitor(ctx context.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAnomalyMonitorWithContext(ctx, input)
}

func (a *Activities) DeleteAnomalySubscription(ctx context.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAnomalySubscriptionWithContext(ctx, input)
}

func (a *Activities) DeleteCostCategoryDefinition(ctx context.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCostCategoryDefinitionWithContext(ctx, input)
}

func (a *Activities) DescribeCostCategoryDefinition(ctx context.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCostCategoryDefinitionWithContext(ctx, input)
}

func (a *Activities) GetAnomalies(ctx context.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAnomaliesWithContext(ctx, input)
}

func (a *Activities) GetAnomalyMonitors(ctx context.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAnomalyMonitorsWithContext(ctx, input)
}

func (a *Activities) GetAnomalySubscriptions(ctx context.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAnomalySubscriptionsWithContext(ctx, input)
}

func (a *Activities) GetCostAndUsage(ctx context.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCostAndUsageWithContext(ctx, input)
}

func (a *Activities) GetCostAndUsageWithResources(ctx context.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCostAndUsageWithResourcesWithContext(ctx, input)
}

func (a *Activities) GetCostForecast(ctx context.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCostForecastWithContext(ctx, input)
}

func (a *Activities) GetDimensionValues(ctx context.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDimensionValuesWithContext(ctx, input)
}

func (a *Activities) GetReservationCoverage(ctx context.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetReservationCoverageWithContext(ctx, input)
}

func (a *Activities) GetReservationPurchaseRecommendation(ctx context.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetReservationPurchaseRecommendationWithContext(ctx, input)
}

func (a *Activities) GetReservationUtilization(ctx context.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetReservationUtilizationWithContext(ctx, input)
}

func (a *Activities) GetRightsizingRecommendation(ctx context.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRightsizingRecommendationWithContext(ctx, input)
}

func (a *Activities) GetSavingsPlansCoverage(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSavingsPlansCoverageWithContext(ctx, input)
}

func (a *Activities) GetSavingsPlansPurchaseRecommendation(ctx context.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSavingsPlansPurchaseRecommendationWithContext(ctx, input)
}

func (a *Activities) GetSavingsPlansUtilization(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSavingsPlansUtilizationWithContext(ctx, input)
}

func (a *Activities) GetSavingsPlansUtilizationDetails(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSavingsPlansUtilizationDetailsWithContext(ctx, input)
}

func (a *Activities) GetTags(ctx context.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTagsWithContext(ctx, input)
}

func (a *Activities) GetUsageForecast(ctx context.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUsageForecastWithContext(ctx, input)
}

func (a *Activities) ListCostCategoryDefinitions(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCostCategoryDefinitionsWithContext(ctx, input)
}

func (a *Activities) ProvideAnomalyFeedback(ctx context.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ProvideAnomalyFeedbackWithContext(ctx, input)
}

func (a *Activities) UpdateAnomalyMonitor(ctx context.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAnomalyMonitorWithContext(ctx, input)
}

func (a *Activities) UpdateAnomalySubscription(ctx context.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAnomalySubscriptionWithContext(ctx, input)
}

func (a *Activities) UpdateCostCategoryDefinition(ctx context.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateCostCategoryDefinitionWithContext(ctx, input)
}