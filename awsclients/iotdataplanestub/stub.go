// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotdataplanestub

import (
	"github.com/aws/aws-sdk-go/service/iotdataplane"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type IoTDataPlaneDeleteThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTDataPlaneDeleteThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTDataPlaneGetThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTDataPlaneGetThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTDataPlaneListNamedShadowsForThingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTDataPlaneListNamedShadowsForThingFuture) Get(ctx workflow.Context) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTDataPlanePublishFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTDataPlanePublishFuture) Get(ctx workflow.Context) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTDataPlaneUpdateThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTDataPlaneUpdateThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteThingShadow(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.DeleteThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteThingShadowAsync(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) *IoTDataPlaneDeleteThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.DeleteThingShadow", input)
	return &IoTDataPlaneDeleteThingShadowFuture{Future: future}
}

func (a *stub) GetThingShadow(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.GetThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetThingShadowAsync(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) *IoTDataPlaneGetThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.GetThingShadow", input)
	return &IoTDataPlaneGetThingShadowFuture{Future: future}
}

func (a *stub) ListNamedShadowsForThing(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.ListNamedShadowsForThing", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNamedShadowsForThingAsync(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) *IoTDataPlaneListNamedShadowsForThingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.ListNamedShadowsForThing", input)
	return &IoTDataPlaneListNamedShadowsForThingFuture{Future: future}
}

func (a *stub) Publish(ctx workflow.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.Publish", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PublishAsync(ctx workflow.Context, input *iotdataplane.PublishInput) *IoTDataPlanePublishFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.Publish", input)
	return &IoTDataPlanePublishFuture{Future: future}
}

func (a *stub) UpdateThingShadow(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.UpdateThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateThingShadowAsync(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) *IoTDataPlaneUpdateThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.UpdateThingShadow", input)
	return &IoTDataPlaneUpdateThingShadowFuture{Future: future}
}
