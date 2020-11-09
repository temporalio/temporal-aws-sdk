// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package migrationhubconfigstub

import (
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateHomeRegionControl(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error)
	CreateHomeRegionControlAsync(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) *CreateHomeRegionControlFuture

	DescribeHomeRegionControls(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error)
	DescribeHomeRegionControlsAsync(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) *DescribeHomeRegionControlsFuture

	GetHomeRegion(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error)
	GetHomeRegionAsync(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) *GetHomeRegionFuture
}

func NewClient() Client {
	return &stub{}
}
