// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package healthstub

import (
	"github.com/aws/aws-sdk-go/service/health"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	DescribeAffectedAccountsForOrganization(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error)
	DescribeAffectedAccountsForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) *HealthDescribeAffectedAccountsForOrganizationFuture

	DescribeAffectedEntities(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error)
	DescribeAffectedEntitiesAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) *HealthDescribeAffectedEntitiesFuture

	DescribeAffectedEntitiesForOrganization(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error)
	DescribeAffectedEntitiesForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) *HealthDescribeAffectedEntitiesForOrganizationFuture

	DescribeEntityAggregates(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEntityAggregatesAsync(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) *HealthDescribeEntityAggregatesFuture

	DescribeEventAggregates(ctx workflow.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventAggregatesAsync(ctx workflow.Context, input *health.DescribeEventAggregatesInput) *HealthDescribeEventAggregatesFuture

	DescribeEventDetails(ctx workflow.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error)
	DescribeEventDetailsAsync(ctx workflow.Context, input *health.DescribeEventDetailsInput) *HealthDescribeEventDetailsFuture

	DescribeEventDetailsForOrganization(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error)
	DescribeEventDetailsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) *HealthDescribeEventDetailsForOrganizationFuture

	DescribeEventTypes(ctx workflow.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error)
	DescribeEventTypesAsync(ctx workflow.Context, input *health.DescribeEventTypesInput) *HealthDescribeEventTypesFuture

	DescribeEvents(ctx workflow.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *health.DescribeEventsInput) *HealthDescribeEventsFuture

	DescribeEventsForOrganization(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error)
	DescribeEventsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) *HealthDescribeEventsForOrganizationFuture

	DescribeHealthServiceStatusForOrganization(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error)
	DescribeHealthServiceStatusForOrganizationAsync(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) *HealthDescribeHealthServiceStatusForOrganizationFuture

	DisableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error)
	DisableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) *HealthDisableHealthServiceAccessForOrganizationFuture

	EnableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error)
	EnableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) *HealthEnableHealthServiceAccessForOrganizationFuture
}

func NewClient() Client {
	return &stub{}
}

