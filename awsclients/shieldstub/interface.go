// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package shieldstub

import (
	"github.com/aws/aws-sdk-go/service/shield"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssociateDRTLogBucket(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error)
	AssociateDRTLogBucketAsync(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) *ShieldAssociateDRTLogBucketFuture

	AssociateDRTRole(ctx workflow.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error)
	AssociateDRTRoleAsync(ctx workflow.Context, input *shield.AssociateDRTRoleInput) *ShieldAssociateDRTRoleFuture

	AssociateHealthCheck(ctx workflow.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error)
	AssociateHealthCheckAsync(ctx workflow.Context, input *shield.AssociateHealthCheckInput) *ShieldAssociateHealthCheckFuture

	AssociateProactiveEngagementDetails(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error)
	AssociateProactiveEngagementDetailsAsync(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) *ShieldAssociateProactiveEngagementDetailsFuture

	CreateProtection(ctx workflow.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error)
	CreateProtectionAsync(ctx workflow.Context, input *shield.CreateProtectionInput) *ShieldCreateProtectionFuture

	CreateSubscription(ctx workflow.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error)
	CreateSubscriptionAsync(ctx workflow.Context, input *shield.CreateSubscriptionInput) *ShieldCreateSubscriptionFuture

	DeleteProtection(ctx workflow.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error)
	DeleteProtectionAsync(ctx workflow.Context, input *shield.DeleteProtectionInput) *ShieldDeleteProtectionFuture

	DeleteSubscription(ctx workflow.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error)
	DeleteSubscriptionAsync(ctx workflow.Context, input *shield.DeleteSubscriptionInput) *ShieldDeleteSubscriptionFuture

	DescribeAttack(ctx workflow.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error)
	DescribeAttackAsync(ctx workflow.Context, input *shield.DescribeAttackInput) *ShieldDescribeAttackFuture

	DescribeDRTAccess(ctx workflow.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error)
	DescribeDRTAccessAsync(ctx workflow.Context, input *shield.DescribeDRTAccessInput) *ShieldDescribeDRTAccessFuture

	DescribeEmergencyContactSettings(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error)
	DescribeEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) *ShieldDescribeEmergencyContactSettingsFuture

	DescribeProtection(ctx workflow.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error)
	DescribeProtectionAsync(ctx workflow.Context, input *shield.DescribeProtectionInput) *ShieldDescribeProtectionFuture

	DescribeSubscription(ctx workflow.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error)
	DescribeSubscriptionAsync(ctx workflow.Context, input *shield.DescribeSubscriptionInput) *ShieldDescribeSubscriptionFuture

	DisableProactiveEngagement(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error)
	DisableProactiveEngagementAsync(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) *ShieldDisableProactiveEngagementFuture

	DisassociateDRTLogBucket(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error)
	DisassociateDRTLogBucketAsync(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) *ShieldDisassociateDRTLogBucketFuture

	DisassociateDRTRole(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error)
	DisassociateDRTRoleAsync(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) *ShieldDisassociateDRTRoleFuture

	DisassociateHealthCheck(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error)
	DisassociateHealthCheckAsync(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) *ShieldDisassociateHealthCheckFuture

	EnableProactiveEngagement(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error)
	EnableProactiveEngagementAsync(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) *ShieldEnableProactiveEngagementFuture

	GetSubscriptionState(ctx workflow.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error)
	GetSubscriptionStateAsync(ctx workflow.Context, input *shield.GetSubscriptionStateInput) *ShieldGetSubscriptionStateFuture

	ListAttacks(ctx workflow.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error)
	ListAttacksAsync(ctx workflow.Context, input *shield.ListAttacksInput) *ShieldListAttacksFuture

	ListProtections(ctx workflow.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error)
	ListProtectionsAsync(ctx workflow.Context, input *shield.ListProtectionsInput) *ShieldListProtectionsFuture

	UpdateEmergencyContactSettings(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error)
	UpdateEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) *ShieldUpdateEmergencyContactSettingsFuture

	UpdateSubscription(ctx workflow.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error)
	UpdateSubscriptionAsync(ctx workflow.Context, input *shield.UpdateSubscriptionInput) *ShieldUpdateSubscriptionFuture
}

func NewClient() Client {
	return &stub{}
}

