// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ec2instanceconnectstub

import (
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	SendSSHPublicKey(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error)
	SendSSHPublicKeyAsync(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) *EC2InstanceConnectSendSSHPublicKeyFuture
}

func NewClient() Client {
	return &stub{}
}
