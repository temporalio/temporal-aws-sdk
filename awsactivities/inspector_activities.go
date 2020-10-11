// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type InspectorActivities struct {
	client inspectoriface.InspectorAPI

	sessionFactory SessionFactory
}

func NewInspectorActivities(sess *session.Session, config ...*aws.Config) *InspectorActivities {
	client := inspector.New(sess, config...)
	return &InspectorActivities{client: client}
}

func NewInspectorActivitiesWithSessionFactory(sessionFactory SessionFactory) *InspectorActivities {
	return &InspectorActivities{sessionFactory: sessionFactory}
}

func (a *InspectorActivities) getClient(ctx context.Context) (inspectoriface.InspectorAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return inspector.New(sess), nil
}

func (a *InspectorActivities) AddAttributesToFindings(ctx context.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddAttributesToFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) CreateAssessmentTarget(ctx context.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAssessmentTargetWithContext(ctx, input)
}

func (a *InspectorActivities) CreateAssessmentTemplate(ctx context.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAssessmentTemplateWithContext(ctx, input)
}

func (a *InspectorActivities) CreateExclusionsPreview(ctx context.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateExclusionsPreviewWithContext(ctx, input)
}

func (a *InspectorActivities) CreateResourceGroup(ctx context.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResourceGroupWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentRun(ctx context.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentTarget(ctx context.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssessmentTargetWithContext(ctx, input)
}

func (a *InspectorActivities) DeleteAssessmentTemplate(ctx context.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssessmentTemplateWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentRuns(ctx context.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssessmentRunsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentTargets(ctx context.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssessmentTargetsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeAssessmentTemplates(ctx context.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssessmentTemplatesWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeCrossAccountAccessRole(ctx context.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCrossAccountAccessRoleWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeExclusions(ctx context.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeExclusionsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeFindings(ctx context.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeResourceGroups(ctx context.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeResourceGroupsWithContext(ctx, input)
}

func (a *InspectorActivities) DescribeRulesPackages(ctx context.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRulesPackagesWithContext(ctx, input)
}

func (a *InspectorActivities) GetAssessmentReport(ctx context.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssessmentReportWithContext(ctx, input)
}

func (a *InspectorActivities) GetExclusionsPreview(ctx context.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetExclusionsPreviewWithContext(ctx, input)
}

func (a *InspectorActivities) GetTelemetryMetadata(ctx context.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTelemetryMetadataWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentRunAgents(ctx context.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssessmentRunAgentsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentRuns(ctx context.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssessmentRunsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentTargets(ctx context.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssessmentTargetsWithContext(ctx, input)
}

func (a *InspectorActivities) ListAssessmentTemplates(ctx context.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssessmentTemplatesWithContext(ctx, input)
}

func (a *InspectorActivities) ListEventSubscriptions(ctx context.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventSubscriptionsWithContext(ctx, input)
}

func (a *InspectorActivities) ListExclusions(ctx context.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListExclusionsWithContext(ctx, input)
}

func (a *InspectorActivities) ListFindings(ctx context.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) ListRulesPackages(ctx context.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRulesPackagesWithContext(ctx, input)
}

func (a *InspectorActivities) ListTagsForResource(ctx context.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *InspectorActivities) PreviewAgents(ctx context.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PreviewAgentsWithContext(ctx, input)
}

func (a *InspectorActivities) RegisterCrossAccountAccessRole(ctx context.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterCrossAccountAccessRoleWithContext(ctx, input)
}

func (a *InspectorActivities) RemoveAttributesFromFindings(ctx context.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveAttributesFromFindingsWithContext(ctx, input)
}

func (a *InspectorActivities) SetTagsForResource(ctx context.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetTagsForResourceWithContext(ctx, input)
}

func (a *InspectorActivities) StartAssessmentRun(ctx context.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) StopAssessmentRun(ctx context.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopAssessmentRunWithContext(ctx, input)
}

func (a *InspectorActivities) SubscribeToEvent(ctx context.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SubscribeToEventWithContext(ctx, input)
}

func (a *InspectorActivities) UnsubscribeFromEvent(ctx context.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnsubscribeFromEventWithContext(ctx, input)
}

func (a *InspectorActivities) UpdateAssessmentTarget(ctx context.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAssessmentTargetWithContext(ctx, input)
}
