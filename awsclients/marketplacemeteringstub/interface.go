// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package marketplacemeteringstub

import (
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	BatchMeterUsage(ctx workflow.Context, input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error)
	BatchMeterUsageAsync(ctx workflow.Context, input *marketplacemetering.BatchMeterUsageInput) *MarketplaceMeteringBatchMeterUsageFuture

	MeterUsage(ctx workflow.Context, input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error)
	MeterUsageAsync(ctx workflow.Context, input *marketplacemetering.MeterUsageInput) *MarketplaceMeteringMeterUsageFuture

	RegisterUsage(ctx workflow.Context, input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error)
	RegisterUsageAsync(ctx workflow.Context, input *marketplacemetering.RegisterUsageInput) *MarketplaceMeteringRegisterUsageFuture

	ResolveCustomer(ctx workflow.Context, input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error)
	ResolveCustomerAsync(ctx workflow.Context, input *marketplacemetering.ResolveCustomerInput) *MarketplaceMeteringResolveCustomerFuture
}

func NewClient() Client {
	return &stub{}
}