// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ramstub

import (
	"github.com/aws/aws-sdk-go/service/ram"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type RAMAcceptResourceShareInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMAcceptResourceShareInvitationFuture) Get(ctx workflow.Context) (*ram.AcceptResourceShareInvitationOutput, error) {
	var output ram.AcceptResourceShareInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMAssociateResourceShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMAssociateResourceShareFuture) Get(ctx workflow.Context) (*ram.AssociateResourceShareOutput, error) {
	var output ram.AssociateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMAssociateResourceSharePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMAssociateResourceSharePermissionFuture) Get(ctx workflow.Context) (*ram.AssociateResourceSharePermissionOutput, error) {
	var output ram.AssociateResourceSharePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMCreateResourceShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMCreateResourceShareFuture) Get(ctx workflow.Context) (*ram.CreateResourceShareOutput, error) {
	var output ram.CreateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMDeleteResourceShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMDeleteResourceShareFuture) Get(ctx workflow.Context) (*ram.DeleteResourceShareOutput, error) {
	var output ram.DeleteResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMDisassociateResourceShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMDisassociateResourceShareFuture) Get(ctx workflow.Context) (*ram.DisassociateResourceShareOutput, error) {
	var output ram.DisassociateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMDisassociateResourceSharePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMDisassociateResourceSharePermissionFuture) Get(ctx workflow.Context) (*ram.DisassociateResourceSharePermissionOutput, error) {
	var output ram.DisassociateResourceSharePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMEnableSharingWithAwsOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMEnableSharingWithAwsOrganizationFuture) Get(ctx workflow.Context) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
	var output ram.EnableSharingWithAwsOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMGetPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMGetPermissionFuture) Get(ctx workflow.Context) (*ram.GetPermissionOutput, error) {
	var output ram.GetPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMGetResourcePoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMGetResourcePoliciesFuture) Get(ctx workflow.Context) (*ram.GetResourcePoliciesOutput, error) {
	var output ram.GetResourcePoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMGetResourceShareAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMGetResourceShareAssociationsFuture) Get(ctx workflow.Context) (*ram.GetResourceShareAssociationsOutput, error) {
	var output ram.GetResourceShareAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMGetResourceShareInvitationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMGetResourceShareInvitationsFuture) Get(ctx workflow.Context) (*ram.GetResourceShareInvitationsOutput, error) {
	var output ram.GetResourceShareInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMGetResourceSharesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMGetResourceSharesFuture) Get(ctx workflow.Context) (*ram.GetResourceSharesOutput, error) {
	var output ram.GetResourceSharesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListPendingInvitationResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListPendingInvitationResourcesFuture) Get(ctx workflow.Context) (*ram.ListPendingInvitationResourcesOutput, error) {
	var output ram.ListPendingInvitationResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListPermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListPermissionsFuture) Get(ctx workflow.Context) (*ram.ListPermissionsOutput, error) {
	var output ram.ListPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListPrincipalsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListPrincipalsFuture) Get(ctx workflow.Context) (*ram.ListPrincipalsOutput, error) {
	var output ram.ListPrincipalsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListResourceSharePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListResourceSharePermissionsFuture) Get(ctx workflow.Context) (*ram.ListResourceSharePermissionsOutput, error) {
	var output ram.ListResourceSharePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListResourceTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListResourceTypesFuture) Get(ctx workflow.Context) (*ram.ListResourceTypesOutput, error) {
	var output ram.ListResourceTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMListResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMListResourcesFuture) Get(ctx workflow.Context) (*ram.ListResourcesOutput, error) {
	var output ram.ListResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMPromoteResourceShareCreatedFromPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMPromoteResourceShareCreatedFromPolicyFuture) Get(ctx workflow.Context) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
	var output ram.PromoteResourceShareCreatedFromPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMRejectResourceShareInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMRejectResourceShareInvitationFuture) Get(ctx workflow.Context) (*ram.RejectResourceShareInvitationOutput, error) {
	var output ram.RejectResourceShareInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMTagResourceFuture) Get(ctx workflow.Context) (*ram.TagResourceOutput, error) {
	var output ram.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMUntagResourceFuture) Get(ctx workflow.Context) (*ram.UntagResourceOutput, error) {
	var output ram.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RAMUpdateResourceShareFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RAMUpdateResourceShareFuture) Get(ctx workflow.Context) (*ram.UpdateResourceShareOutput, error) {
	var output ram.UpdateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
	var output ram.AcceptResourceShareInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AcceptResourceShareInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RAMAcceptResourceShareInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AcceptResourceShareInvitation", input)
	return &RAMAcceptResourceShareInvitationFuture{Future: future}
}

func (a *stub) AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error) {
	var output ram.AssociateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RAMAssociateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceShare", input)
	return &RAMAssociateResourceShareFuture{Future: future}
}

func (a *stub) AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error) {
	var output ram.AssociateResourceSharePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceSharePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RAMAssociateResourceSharePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceSharePermission", input)
	return &RAMAssociateResourceSharePermissionFuture{Future: future}
}

func (a *stub) CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error) {
	var output ram.CreateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.CreateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RAMCreateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.CreateResourceShare", input)
	return &RAMCreateResourceShareFuture{Future: future}
}

func (a *stub) DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error) {
	var output ram.DeleteResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DeleteResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RAMDeleteResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DeleteResourceShare", input)
	return &RAMDeleteResourceShareFuture{Future: future}
}

func (a *stub) DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error) {
	var output ram.DisassociateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RAMDisassociateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceShare", input)
	return &RAMDisassociateResourceShareFuture{Future: future}
}

func (a *stub) DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error) {
	var output ram.DisassociateResourceSharePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceSharePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RAMDisassociateResourceSharePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceSharePermission", input)
	return &RAMDisassociateResourceSharePermissionFuture{Future: future}
}

func (a *stub) EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
	var output ram.EnableSharingWithAwsOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.EnableSharingWithAwsOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RAMEnableSharingWithAwsOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.EnableSharingWithAwsOrganization", input)
	return &RAMEnableSharingWithAwsOrganizationFuture{Future: future}
}

func (a *stub) GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error) {
	var output ram.GetPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RAMGetPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetPermission", input)
	return &RAMGetPermissionFuture{Future: future}
}

func (a *stub) GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error) {
	var output ram.GetResourcePoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourcePolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RAMGetResourcePoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourcePolicies", input)
	return &RAMGetResourcePoliciesFuture{Future: future}
}

func (a *stub) GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error) {
	var output ram.GetResourceShareAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RAMGetResourceShareAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareAssociations", input)
	return &RAMGetResourceShareAssociationsFuture{Future: future}
}

func (a *stub) GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error) {
	var output ram.GetResourceShareInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RAMGetResourceShareInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareInvitations", input)
	return &RAMGetResourceShareInvitationsFuture{Future: future}
}

func (a *stub) GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error) {
	var output ram.GetResourceSharesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShares", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RAMGetResourceSharesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShares", input)
	return &RAMGetResourceSharesFuture{Future: future}
}

func (a *stub) ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error) {
	var output ram.ListPendingInvitationResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPendingInvitationResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RAMListPendingInvitationResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPendingInvitationResources", input)
	return &RAMListPendingInvitationResourcesFuture{Future: future}
}

func (a *stub) ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error) {
	var output ram.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RAMListPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPermissions", input)
	return &RAMListPermissionsFuture{Future: future}
}

func (a *stub) ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error) {
	var output ram.ListPrincipalsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPrincipals", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RAMListPrincipalsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPrincipals", input)
	return &RAMListPrincipalsFuture{Future: future}
}

func (a *stub) ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error) {
	var output ram.ListResourceSharePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceSharePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RAMListResourceSharePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceSharePermissions", input)
	return &RAMListResourceSharePermissionsFuture{Future: future}
}

func (a *stub) ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error) {
	var output ram.ListResourceTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RAMListResourceTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceTypes", input)
	return &RAMListResourceTypesFuture{Future: future}
}

func (a *stub) ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error) {
	var output ram.ListResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RAMListResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResources", input)
	return &RAMListResourcesFuture{Future: future}
}

func (a *stub) PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
	var output ram.PromoteResourceShareCreatedFromPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.PromoteResourceShareCreatedFromPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RAMPromoteResourceShareCreatedFromPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.PromoteResourceShareCreatedFromPolicy", input)
	return &RAMPromoteResourceShareCreatedFromPolicyFuture{Future: future}
}

func (a *stub) RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error) {
	var output ram.RejectResourceShareInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.RejectResourceShareInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RAMRejectResourceShareInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.RejectResourceShareInvitation", input)
	return &RAMRejectResourceShareInvitationFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error) {
	var output ram.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RAMTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.TagResource", input)
	return &RAMTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error) {
	var output ram.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RAMUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.UntagResource", input)
	return &RAMUntagResourceFuture{Future: future}
}

func (a *stub) UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error) {
	var output ram.UpdateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.UpdateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RAMUpdateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.UpdateResourceShare", input)
	return &RAMUpdateResourceShareFuture{Future: future}
}
