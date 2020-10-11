// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mobileanalyticsstub

import (
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	PutEvents(ctx workflow.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error)
	PutEventsAsync(ctx workflow.Context, input *mobileanalytics.PutEventsInput) *MobileAnalyticsPutEventsFuture
}

func NewClient() Client {
	return &stub{}
}
