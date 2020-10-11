// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package redshiftdataapiservicestub

import (
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelStatement(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error)
	CancelStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) *RedshiftDataAPIServiceCancelStatementFuture

	DescribeStatement(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error)
	DescribeStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) *RedshiftDataAPIServiceDescribeStatementFuture

	DescribeTable(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error)
	DescribeTableAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) *RedshiftDataAPIServiceDescribeTableFuture

	ExecuteStatement(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error)
	ExecuteStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) *RedshiftDataAPIServiceExecuteStatementFuture

	GetStatementResult(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error)
	GetStatementResultAsync(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) *RedshiftDataAPIServiceGetStatementResultFuture

	ListDatabases(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error)
	ListDatabasesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) *RedshiftDataAPIServiceListDatabasesFuture

	ListSchemas(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) *RedshiftDataAPIServiceListSchemasFuture

	ListStatements(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error)
	ListStatementsAsync(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) *RedshiftDataAPIServiceListStatementsFuture

	ListTables(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error)
	ListTablesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) *RedshiftDataAPIServiceListTablesFuture
}

func NewClient() Client {
	return &stub{}
}
