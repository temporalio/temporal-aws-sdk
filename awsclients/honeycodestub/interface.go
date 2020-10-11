// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package honeycodestub

import (
	"github.com/aws/aws-sdk-go/service/honeycode"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	GetScreenData(ctx workflow.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error)
	GetScreenDataAsync(ctx workflow.Context, input *honeycode.GetScreenDataInput) *HoneycodeGetScreenDataFuture

	InvokeScreenAutomation(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error)
	InvokeScreenAutomationAsync(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) *HoneycodeInvokeScreenAutomationFuture
}

func NewClient() Client {
	return &stub{}
}