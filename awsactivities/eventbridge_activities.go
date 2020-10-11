// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type EventBridgeActivities struct {
	client eventbridgeiface.EventBridgeAPI

	sessionFactory SessionFactory
}

func NewEventBridgeActivities(sess *session.Session, config ...*aws.Config) *EventBridgeActivities {
	client := eventbridge.New(sess, config...)
	return &EventBridgeActivities{client: client}
}

func NewEventBridgeActivitiesWithSessionFactory(sessionFactory SessionFactory) *EventBridgeActivities {
	return &EventBridgeActivities{sessionFactory: sessionFactory}
}

func (a *EventBridgeActivities) getClient(ctx context.Context) (eventbridgeiface.EventBridgeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return eventbridge.New(sess), nil
}

func (a *EventBridgeActivities) ActivateEventSource(ctx context.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ActivateEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) CreateEventBus(ctx context.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) CreatePartnerEventSource(ctx context.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeactivateEventSource(ctx context.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeactivateEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeleteEventBus(ctx context.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeletePartnerEventSource(ctx context.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DeleteRule(ctx context.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeEventBus(ctx context.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventBusWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeEventSource(ctx context.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribePartnerEventSource(ctx context.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePartnerEventSourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) DescribeRule(ctx context.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) DisableRule(ctx context.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) EnableRule(ctx context.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListEventBuses(ctx context.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventBusesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListEventSources(ctx context.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventSourcesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListPartnerEventSourceAccounts(ctx context.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartnerEventSourceAccountsWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListPartnerEventSources(ctx context.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartnerEventSourcesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListRuleNamesByTarget(ctx context.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRuleNamesByTargetWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListRules(ctx context.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRulesWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListTagsForResource(ctx context.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) ListTargetsByRule(ctx context.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTargetsByRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutEvents(ctx context.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEventsWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutPartnerEvents(ctx context.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPartnerEventsWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutPermission(ctx context.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPermissionWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutRule(ctx context.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutRuleWithContext(ctx, input)
}

func (a *EventBridgeActivities) PutTargets(ctx context.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutTargetsWithContext(ctx, input)
}

func (a *EventBridgeActivities) RemovePermission(ctx context.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemovePermissionWithContext(ctx, input)
}

func (a *EventBridgeActivities) RemoveTargets(ctx context.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTargetsWithContext(ctx, input)
}

func (a *EventBridgeActivities) TagResource(ctx context.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *EventBridgeActivities) TestEventPattern(ctx context.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestEventPatternWithContext(ctx, input)
}

func (a *EventBridgeActivities) UntagResource(ctx context.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
