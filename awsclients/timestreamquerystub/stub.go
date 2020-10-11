// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package timestreamquerystub

import (
	"github.com/aws/aws-sdk-go/service/timestreamquery"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type TimestreamQueryCancelQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamQueryCancelQueryFuture) Get(ctx workflow.Context) (*timestreamquery.CancelQueryOutput, error) {
	var output timestreamquery.CancelQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamQueryDescribeEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamQueryDescribeEndpointsFuture) Get(ctx workflow.Context) (*timestreamquery.DescribeEndpointsOutput, error) {
	var output timestreamquery.DescribeEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamQueryQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamQueryQueryFuture) Get(ctx workflow.Context) (*timestreamquery.QueryOutput, error) {
	var output timestreamquery.QueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelQuery(ctx workflow.Context, input *timestreamquery.CancelQueryInput) (*timestreamquery.CancelQueryOutput, error) {
	var output timestreamquery.CancelQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamquery.CancelQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelQueryAsync(ctx workflow.Context, input *timestreamquery.CancelQueryInput) *TimestreamQueryCancelQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamquery.CancelQuery", input)
	return &TimestreamQueryCancelQueryFuture{Future: future}
}

func (a *stub) DescribeEndpoints(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) (*timestreamquery.DescribeEndpointsOutput, error) {
	var output timestreamquery.DescribeEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamquery.DescribeEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEndpointsAsync(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) *TimestreamQueryDescribeEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamquery.DescribeEndpoints", input)
	return &TimestreamQueryDescribeEndpointsFuture{Future: future}
}

func (a *stub) Query(ctx workflow.Context, input *timestreamquery.QueryInput) (*timestreamquery.QueryOutput, error) {
	var output timestreamquery.QueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamquery.Query", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryAsync(ctx workflow.Context, input *timestreamquery.QueryInput) *TimestreamQueryQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamquery.Query", input)
	return &TimestreamQueryQueryFuture{Future: future}
}
