// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssooidcstub

import (
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error)
	CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *SSOOIDCCreateTokenFuture

	RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error)
	RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *SSOOIDCRegisterClientFuture

	StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error)
	StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *SSOOIDCStartDeviceAuthorizationFuture
}

func NewClient() Client {
	return &stub{}
}
