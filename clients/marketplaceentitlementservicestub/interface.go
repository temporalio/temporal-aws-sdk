// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package marketplaceentitlementservicestub

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
	GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *GetEntitlementsFuture
}

func NewClient() Client {
	return &stub{}
}
