// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudWatchActivities struct {
	client cloudwatchiface.CloudWatchAPI

	sessionFactory SessionFactory
}

func NewCloudWatchActivities(sess *session.Session, config ...*aws.Config) *CloudWatchActivities {
	client := cloudwatch.New(sess, config...)
	return &CloudWatchActivities{client: client}
}

func NewCloudWatchActivitiesWithSessionFactory(sessionFactory SessionFactory) *CloudWatchActivities {
	return &CloudWatchActivities{sessionFactory: sessionFactory}
}

func (a *CloudWatchActivities) getClient(ctx context.Context) (cloudwatchiface.CloudWatchAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cloudwatch.New(sess), nil
}

func (a *CloudWatchActivities) DeleteAlarms(ctx context.Context, input *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAlarmsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteAnomalyDetector(ctx context.Context, input *cloudwatch.DeleteAnomalyDetectorInput) (*cloudwatch.DeleteAnomalyDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAnomalyDetectorWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteDashboards(ctx context.Context, input *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDashboardsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DeleteInsightRules(ctx context.Context, input *cloudwatch.DeleteInsightRulesInput) (*cloudwatch.DeleteInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarmHistory(ctx context.Context, input *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAlarmHistoryWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarms(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAlarmsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAlarmsForMetric(ctx context.Context, input *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAlarmsForMetricWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeAnomalyDetectors(ctx context.Context, input *cloudwatch.DescribeAnomalyDetectorsInput) (*cloudwatch.DescribeAnomalyDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAnomalyDetectorsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DescribeInsightRules(ctx context.Context, input *cloudwatch.DescribeInsightRulesInput) (*cloudwatch.DescribeInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) DisableAlarmActions(ctx context.Context, input *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableAlarmActionsWithContext(ctx, input)
}

func (a *CloudWatchActivities) DisableInsightRules(ctx context.Context, input *cloudwatch.DisableInsightRulesInput) (*cloudwatch.DisableInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) EnableAlarmActions(ctx context.Context, input *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableAlarmActionsWithContext(ctx, input)
}

func (a *CloudWatchActivities) EnableInsightRules(ctx context.Context, input *cloudwatch.EnableInsightRulesInput) (*cloudwatch.EnableInsightRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableInsightRulesWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetDashboard(ctx context.Context, input *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDashboardWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetInsightRuleReport(ctx context.Context, input *cloudwatch.GetInsightRuleReportInput) (*cloudwatch.GetInsightRuleReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInsightRuleReportWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricData(ctx context.Context, input *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMetricDataWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricStatistics(ctx context.Context, input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMetricStatisticsWithContext(ctx, input)
}

func (a *CloudWatchActivities) GetMetricWidgetImage(ctx context.Context, input *cloudwatch.GetMetricWidgetImageInput) (*cloudwatch.GetMetricWidgetImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMetricWidgetImageWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListDashboards(ctx context.Context, input *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDashboardsWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListMetrics(ctx context.Context, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMetricsWithContext(ctx, input)
}

func (a *CloudWatchActivities) ListTagsForResource(ctx context.Context, input *cloudwatch.ListTagsForResourceInput) (*cloudwatch.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutAnomalyDetector(ctx context.Context, input *cloudwatch.PutAnomalyDetectorInput) (*cloudwatch.PutAnomalyDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAnomalyDetectorWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutCompositeAlarm(ctx context.Context, input *cloudwatch.PutCompositeAlarmInput) (*cloudwatch.PutCompositeAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutCompositeAlarmWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutDashboard(ctx context.Context, input *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDashboardWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutInsightRule(ctx context.Context, input *cloudwatch.PutInsightRuleInput) (*cloudwatch.PutInsightRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutInsightRuleWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutMetricAlarm(ctx context.Context, input *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutMetricAlarmWithContext(ctx, input)
}

func (a *CloudWatchActivities) PutMetricData(ctx context.Context, input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutMetricDataWithContext(ctx, input)
}

func (a *CloudWatchActivities) SetAlarmState(ctx context.Context, input *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetAlarmStateWithContext(ctx, input)
}

func (a *CloudWatchActivities) TagResource(ctx context.Context, input *cloudwatch.TagResourceInput) (*cloudwatch.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) UntagResource(ctx context.Context, input *cloudwatch.UntagResourceInput) (*cloudwatch.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *CloudWatchActivities) WaitUntilAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilAlarmExistsWithContext(ctx, input, options...)
	})
}

func (a *CloudWatchActivities) WaitUntilCompositeAlarmExists(ctx context.Context, input *cloudwatch.DescribeAlarmsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilCompositeAlarmExistsWithContext(ctx, input, options...)
	})
}
