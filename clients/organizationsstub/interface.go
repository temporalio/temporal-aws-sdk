// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package organizationsstub

import (
	"github.com/aws/aws-sdk-go/service/organizations"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptHandshake(ctx workflow.Context, input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error)
	AcceptHandshakeAsync(ctx workflow.Context, input *organizations.AcceptHandshakeInput) *AcceptHandshakeFuture

	AttachPolicy(ctx workflow.Context, input *organizations.AttachPolicyInput) (*organizations.AttachPolicyOutput, error)
	AttachPolicyAsync(ctx workflow.Context, input *organizations.AttachPolicyInput) *AttachPolicyFuture

	CancelHandshake(ctx workflow.Context, input *organizations.CancelHandshakeInput) (*organizations.CancelHandshakeOutput, error)
	CancelHandshakeAsync(ctx workflow.Context, input *organizations.CancelHandshakeInput) *CancelHandshakeFuture

	CreateAccount(ctx workflow.Context, input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error)
	CreateAccountAsync(ctx workflow.Context, input *organizations.CreateAccountInput) *CreateAccountFuture

	CreateGovCloudAccount(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) (*organizations.CreateGovCloudAccountOutput, error)
	CreateGovCloudAccountAsync(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) *CreateGovCloudAccountFuture

	CreateOrganization(ctx workflow.Context, input *organizations.CreateOrganizationInput) (*organizations.CreateOrganizationOutput, error)
	CreateOrganizationAsync(ctx workflow.Context, input *organizations.CreateOrganizationInput) *CreateOrganizationFuture

	CreateOrganizationalUnit(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) (*organizations.CreateOrganizationalUnitOutput, error)
	CreateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) *CreateOrganizationalUnitFuture

	CreatePolicy(ctx workflow.Context, input *organizations.CreatePolicyInput) (*organizations.CreatePolicyOutput, error)
	CreatePolicyAsync(ctx workflow.Context, input *organizations.CreatePolicyInput) *CreatePolicyFuture

	DeclineHandshake(ctx workflow.Context, input *organizations.DeclineHandshakeInput) (*organizations.DeclineHandshakeOutput, error)
	DeclineHandshakeAsync(ctx workflow.Context, input *organizations.DeclineHandshakeInput) *DeclineHandshakeFuture

	DeleteOrganization(ctx workflow.Context, input *organizations.DeleteOrganizationInput) (*organizations.DeleteOrganizationOutput, error)
	DeleteOrganizationAsync(ctx workflow.Context, input *organizations.DeleteOrganizationInput) *DeleteOrganizationFuture

	DeleteOrganizationalUnit(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) (*organizations.DeleteOrganizationalUnitOutput, error)
	DeleteOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) *DeleteOrganizationalUnitFuture

	DeletePolicy(ctx workflow.Context, input *organizations.DeletePolicyInput) (*organizations.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *organizations.DeletePolicyInput) *DeletePolicyFuture

	DeregisterDelegatedAdministrator(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) (*organizations.DeregisterDelegatedAdministratorOutput, error)
	DeregisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) *DeregisterDelegatedAdministratorFuture

	DescribeAccount(ctx workflow.Context, input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error)
	DescribeAccountAsync(ctx workflow.Context, input *organizations.DescribeAccountInput) *DescribeAccountFuture

	DescribeCreateAccountStatus(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error)
	DescribeCreateAccountStatusAsync(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) *DescribeCreateAccountStatusFuture

	DescribeEffectivePolicy(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error)
	DescribeEffectivePolicyAsync(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) *DescribeEffectivePolicyFuture

	DescribeHandshake(ctx workflow.Context, input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error)
	DescribeHandshakeAsync(ctx workflow.Context, input *organizations.DescribeHandshakeInput) *DescribeHandshakeFuture

	DescribeOrganization(ctx workflow.Context, input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error)
	DescribeOrganizationAsync(ctx workflow.Context, input *organizations.DescribeOrganizationInput) *DescribeOrganizationFuture

	DescribeOrganizationalUnit(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error)
	DescribeOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) *DescribeOrganizationalUnitFuture

	DescribePolicy(ctx workflow.Context, input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error)
	DescribePolicyAsync(ctx workflow.Context, input *organizations.DescribePolicyInput) *DescribePolicyFuture

	DetachPolicy(ctx workflow.Context, input *organizations.DetachPolicyInput) (*organizations.DetachPolicyOutput, error)
	DetachPolicyAsync(ctx workflow.Context, input *organizations.DetachPolicyInput) *DetachPolicyFuture

	DisableAWSServiceAccess(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) (*organizations.DisableAWSServiceAccessOutput, error)
	DisableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) *DisableAWSServiceAccessFuture

	DisablePolicyType(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) (*organizations.DisablePolicyTypeOutput, error)
	DisablePolicyTypeAsync(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) *DisablePolicyTypeFuture

	EnableAWSServiceAccess(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) (*organizations.EnableAWSServiceAccessOutput, error)
	EnableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) *EnableAWSServiceAccessFuture

	EnableAllFeatures(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) (*organizations.EnableAllFeaturesOutput, error)
	EnableAllFeaturesAsync(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) *EnableAllFeaturesFuture

	EnablePolicyType(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) (*organizations.EnablePolicyTypeOutput, error)
	EnablePolicyTypeAsync(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) *EnablePolicyTypeFuture

	InviteAccountToOrganization(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) (*organizations.InviteAccountToOrganizationOutput, error)
	InviteAccountToOrganizationAsync(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) *InviteAccountToOrganizationFuture

	LeaveOrganization(ctx workflow.Context, input *organizations.LeaveOrganizationInput) (*organizations.LeaveOrganizationOutput, error)
	LeaveOrganizationAsync(ctx workflow.Context, input *organizations.LeaveOrganizationInput) *LeaveOrganizationFuture

	ListAWSServiceAccessForOrganization(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error)
	ListAWSServiceAccessForOrganizationAsync(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) *ListAWSServiceAccessForOrganizationFuture

	ListAccounts(ctx workflow.Context, input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error)
	ListAccountsAsync(ctx workflow.Context, input *organizations.ListAccountsInput) *ListAccountsFuture

	ListAccountsForParent(ctx workflow.Context, input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error)
	ListAccountsForParentAsync(ctx workflow.Context, input *organizations.ListAccountsForParentInput) *ListAccountsForParentFuture

	ListChildren(ctx workflow.Context, input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error)
	ListChildrenAsync(ctx workflow.Context, input *organizations.ListChildrenInput) *ListChildrenFuture

	ListCreateAccountStatus(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error)
	ListCreateAccountStatusAsync(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) *ListCreateAccountStatusFuture

	ListDelegatedAdministrators(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error)
	ListDelegatedAdministratorsAsync(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) *ListDelegatedAdministratorsFuture

	ListDelegatedServicesForAccount(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error)
	ListDelegatedServicesForAccountAsync(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) *ListDelegatedServicesForAccountFuture

	ListHandshakesForAccount(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error)
	ListHandshakesForAccountAsync(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) *ListHandshakesForAccountFuture

	ListHandshakesForOrganization(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error)
	ListHandshakesForOrganizationAsync(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) *ListHandshakesForOrganizationFuture

	ListOrganizationalUnitsForParent(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error)
	ListOrganizationalUnitsForParentAsync(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) *ListOrganizationalUnitsForParentFuture

	ListParents(ctx workflow.Context, input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error)
	ListParentsAsync(ctx workflow.Context, input *organizations.ListParentsInput) *ListParentsFuture

	ListPolicies(ctx workflow.Context, input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error)
	ListPoliciesAsync(ctx workflow.Context, input *organizations.ListPoliciesInput) *ListPoliciesFuture

	ListPoliciesForTarget(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error)
	ListPoliciesForTargetAsync(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) *ListPoliciesForTargetFuture

	ListRoots(ctx workflow.Context, input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error)
	ListRootsAsync(ctx workflow.Context, input *organizations.ListRootsInput) *ListRootsFuture

	ListTagsForResource(ctx workflow.Context, input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *organizations.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTargetsForPolicy(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error)
	ListTargetsForPolicyAsync(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) *ListTargetsForPolicyFuture

	MoveAccount(ctx workflow.Context, input *organizations.MoveAccountInput) (*organizations.MoveAccountOutput, error)
	MoveAccountAsync(ctx workflow.Context, input *organizations.MoveAccountInput) *MoveAccountFuture

	RegisterDelegatedAdministrator(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) (*organizations.RegisterDelegatedAdministratorOutput, error)
	RegisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) *RegisterDelegatedAdministratorFuture

	RemoveAccountFromOrganization(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) (*organizations.RemoveAccountFromOrganizationOutput, error)
	RemoveAccountFromOrganizationAsync(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) *RemoveAccountFromOrganizationFuture

	TagResource(ctx workflow.Context, input *organizations.TagResourceInput) (*organizations.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *organizations.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *organizations.UntagResourceInput) (*organizations.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *organizations.UntagResourceInput) *UntagResourceFuture

	UpdateOrganizationalUnit(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) (*organizations.UpdateOrganizationalUnitOutput, error)
	UpdateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) *UpdateOrganizationalUnitFuture

	UpdatePolicy(ctx workflow.Context, input *organizations.UpdatePolicyInput) (*organizations.UpdatePolicyOutput, error)
	UpdatePolicyAsync(ctx workflow.Context, input *organizations.UpdatePolicyInput) *UpdatePolicyFuture
}

func NewClient() Client {
	return &stub{}
}
