// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudwatchevents

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
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
	client cloudwatcheventsiface.CloudWatchEventsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := cloudwatchevents.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (cloudwatcheventsiface.CloudWatchEventsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cloudwatchevents.New(sess), nil
}

func (a *Activities) ActivateEventSource(ctx context.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ActivateEventSourceWithContext(ctx, input)
}

func (a *Activities) CreateEventBus(ctx context.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEventBusWithContext(ctx, input)
}

func (a *Activities) CreatePartnerEventSource(ctx context.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePartnerEventSourceWithContext(ctx, input)
}

func (a *Activities) DeactivateEventSource(ctx context.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeactivateEventSourceWithContext(ctx, input)
}

func (a *Activities) DeleteEventBus(ctx context.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventBusWithContext(ctx, input)
}

func (a *Activities) DeletePartnerEventSource(ctx context.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePartnerEventSourceWithContext(ctx, input)
}

func (a *Activities) DeleteRule(ctx context.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRuleWithContext(ctx, input)
}

func (a *Activities) DescribeEventBus(ctx context.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventBusWithContext(ctx, input)
}

func (a *Activities) DescribeEventSource(ctx context.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventSourceWithContext(ctx, input)
}

func (a *Activities) DescribePartnerEventSource(ctx context.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePartnerEventSourceWithContext(ctx, input)
}

func (a *Activities) DescribeRule(ctx context.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRuleWithContext(ctx, input)
}

func (a *Activities) DisableRule(ctx context.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableRuleWithContext(ctx, input)
}

func (a *Activities) EnableRule(ctx context.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableRuleWithContext(ctx, input)
}

func (a *Activities) ListEventBuses(ctx context.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventBusesWithContext(ctx, input)
}

func (a *Activities) ListEventSources(ctx context.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventSourcesWithContext(ctx, input)
}

func (a *Activities) ListPartnerEventSourceAccounts(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartnerEventSourceAccountsWithContext(ctx, input)
}

func (a *Activities) ListPartnerEventSources(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartnerEventSourcesWithContext(ctx, input)
}

func (a *Activities) ListRuleNamesByTarget(ctx context.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRuleNamesByTargetWithContext(ctx, input)
}

func (a *Activities) ListRules(ctx context.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRulesWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ListTargetsByRule(ctx context.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTargetsByRuleWithContext(ctx, input)
}

func (a *Activities) PutEvents(ctx context.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEventsWithContext(ctx, input)
}

func (a *Activities) PutPartnerEvents(ctx context.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPartnerEventsWithContext(ctx, input)
}

func (a *Activities) PutPermission(ctx context.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPermissionWithContext(ctx, input)
}

func (a *Activities) PutRule(ctx context.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutRuleWithContext(ctx, input)
}

func (a *Activities) PutTargets(ctx context.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutTargetsWithContext(ctx, input)
}

func (a *Activities) RemovePermission(ctx context.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemovePermissionWithContext(ctx, input)
}

func (a *Activities) RemoveTargets(ctx context.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTargetsWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) TestEventPattern(ctx context.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestEventPatternWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}