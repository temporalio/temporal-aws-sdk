// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SupportActivities struct {
	client supportiface.SupportAPI

	sessionFactory SessionFactory
}

func NewSupportActivities(sess *session.Session, config ...*aws.Config) *SupportActivities {
	client := support.New(sess, config...)
	return &SupportActivities{client: client}
}

func NewSupportActivitiesWithSessionFactory(sessionFactory SessionFactory) *SupportActivities {
	return &SupportActivities{sessionFactory: sessionFactory}
}

func (a *SupportActivities) getClient(ctx context.Context) (supportiface.SupportAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return support.New(sess), nil
}

func (a *SupportActivities) AddAttachmentsToSet(ctx context.Context, input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddAttachmentsToSetWithContext(ctx, input)
}

func (a *SupportActivities) AddCommunicationToCase(ctx context.Context, input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddCommunicationToCaseWithContext(ctx, input)
}

func (a *SupportActivities) CreateCase(ctx context.Context, input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCaseWithContext(ctx, input)
}

func (a *SupportActivities) DescribeAttachment(ctx context.Context, input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAttachmentWithContext(ctx, input)
}

func (a *SupportActivities) DescribeCases(ctx context.Context, input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCasesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeCommunications(ctx context.Context, input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCommunicationsWithContext(ctx, input)
}

func (a *SupportActivities) DescribeServices(ctx context.Context, input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServicesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeSeverityLevels(ctx context.Context, input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSeverityLevelsWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckRefreshStatuses(ctx context.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckResult(ctx context.Context, input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTrustedAdvisorCheckResultWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckSummaries(ctx context.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTrustedAdvisorCheckSummariesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorChecks(ctx context.Context, input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTrustedAdvisorChecksWithContext(ctx, input)
}

func (a *SupportActivities) RefreshTrustedAdvisorCheck(ctx context.Context, input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RefreshTrustedAdvisorCheckWithContext(ctx, input)
}

func (a *SupportActivities) ResolveCase(ctx context.Context, input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResolveCaseWithContext(ctx, input)
}
