// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssostub

import (
	"github.com/aws/aws-sdk-go/service/sso"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	GetRoleCredentials(ctx workflow.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error)
	GetRoleCredentialsAsync(ctx workflow.Context, input *sso.GetRoleCredentialsInput) *SSOGetRoleCredentialsFuture

	ListAccountRoles(ctx workflow.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error)
	ListAccountRolesAsync(ctx workflow.Context, input *sso.ListAccountRolesInput) *SSOListAccountRolesFuture

	ListAccounts(ctx workflow.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error)
	ListAccountsAsync(ctx workflow.Context, input *sso.ListAccountsInput) *SSOListAccountsFuture

	Logout(ctx workflow.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error)
	LogoutAsync(ctx workflow.Context, input *sso.LogoutInput) *SSOLogoutFuture
}

func NewClient() Client {
	return &stub{}
}
