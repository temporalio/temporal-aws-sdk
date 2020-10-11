// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iamstub

import (
	"github.com/aws/aws-sdk-go/service/iam"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AddClientIDToOpenIDConnectProvider(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddClientIDToOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) *IAMAddClientIDToOpenIDConnectProviderFuture

	AddRoleToInstanceProfile(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)
	AddRoleToInstanceProfileAsync(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) *IAMAddRoleToInstanceProfileFuture

	AddUserToGroup(ctx workflow.Context, input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)
	AddUserToGroupAsync(ctx workflow.Context, input *iam.AddUserToGroupInput) *IAMAddUserToGroupFuture

	AttachGroupPolicy(ctx workflow.Context, input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)
	AttachGroupPolicyAsync(ctx workflow.Context, input *iam.AttachGroupPolicyInput) *IAMAttachGroupPolicyFuture

	AttachRolePolicy(ctx workflow.Context, input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)
	AttachRolePolicyAsync(ctx workflow.Context, input *iam.AttachRolePolicyInput) *IAMAttachRolePolicyFuture

	AttachUserPolicy(ctx workflow.Context, input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)
	AttachUserPolicyAsync(ctx workflow.Context, input *iam.AttachUserPolicyInput) *IAMAttachUserPolicyFuture

	ChangePassword(ctx workflow.Context, input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)
	ChangePasswordAsync(ctx workflow.Context, input *iam.ChangePasswordInput) *IAMChangePasswordFuture

	CreateAccessKey(ctx workflow.Context, input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)
	CreateAccessKeyAsync(ctx workflow.Context, input *iam.CreateAccessKeyInput) *IAMCreateAccessKeyFuture

	CreateAccountAlias(ctx workflow.Context, input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)
	CreateAccountAliasAsync(ctx workflow.Context, input *iam.CreateAccountAliasInput) *IAMCreateAccountAliasFuture

	CreateGroup(ctx workflow.Context, input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *iam.CreateGroupInput) *IAMCreateGroupFuture

	CreateInstanceProfile(ctx workflow.Context, input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)
	CreateInstanceProfileAsync(ctx workflow.Context, input *iam.CreateInstanceProfileInput) *IAMCreateInstanceProfileFuture

	CreateLoginProfile(ctx workflow.Context, input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)
	CreateLoginProfileAsync(ctx workflow.Context, input *iam.CreateLoginProfileInput) *IAMCreateLoginProfileFuture

	CreateOpenIDConnectProvider(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreateOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) *IAMCreateOpenIDConnectProviderFuture

	CreatePolicy(ctx workflow.Context, input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
	CreatePolicyAsync(ctx workflow.Context, input *iam.CreatePolicyInput) *IAMCreatePolicyFuture

	CreatePolicyVersion(ctx workflow.Context, input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyVersionAsync(ctx workflow.Context, input *iam.CreatePolicyVersionInput) *IAMCreatePolicyVersionFuture

	CreateRole(ctx workflow.Context, input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error)
	CreateRoleAsync(ctx workflow.Context, input *iam.CreateRoleInput) *IAMCreateRoleFuture

	CreateSAMLProvider(ctx workflow.Context, input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)
	CreateSAMLProviderAsync(ctx workflow.Context, input *iam.CreateSAMLProviderInput) *IAMCreateSAMLProviderFuture

	CreateServiceLinkedRole(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleAsync(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) *IAMCreateServiceLinkedRoleFuture

	CreateServiceSpecificCredential(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) *IAMCreateServiceSpecificCredentialFuture

	CreateUser(ctx workflow.Context, input *iam.CreateUserInput) (*iam.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *iam.CreateUserInput) *IAMCreateUserFuture

	CreateVirtualMFADevice(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)
	CreateVirtualMFADeviceAsync(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) *IAMCreateVirtualMFADeviceFuture

	DeactivateMFADevice(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)
	DeactivateMFADeviceAsync(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) *IAMDeactivateMFADeviceFuture

	DeleteAccessKey(ctx workflow.Context, input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccessKeyAsync(ctx workflow.Context, input *iam.DeleteAccessKeyInput) *IAMDeleteAccessKeyFuture

	DeleteAccountAlias(ctx workflow.Context, input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountAliasAsync(ctx workflow.Context, input *iam.DeleteAccountAliasInput) *IAMDeleteAccountAliasFuture

	DeleteAccountPasswordPolicy(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) *IAMDeleteAccountPasswordPolicyFuture

	DeleteGroup(ctx workflow.Context, input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *iam.DeleteGroupInput) *IAMDeleteGroupFuture

	DeleteGroupPolicy(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupPolicyAsync(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) *IAMDeleteGroupPolicyFuture

	DeleteInstanceProfile(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileAsync(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) *IAMDeleteInstanceProfileFuture

	DeleteLoginProfile(ctx workflow.Context, input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)
	DeleteLoginProfileAsync(ctx workflow.Context, input *iam.DeleteLoginProfileInput) *IAMDeleteLoginProfileFuture

	DeleteOpenIDConnectProvider(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeleteOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) *IAMDeleteOpenIDConnectProviderFuture

	DeletePolicy(ctx workflow.Context, input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *iam.DeletePolicyInput) *IAMDeletePolicyFuture

	DeletePolicyVersion(ctx workflow.Context, input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyVersionAsync(ctx workflow.Context, input *iam.DeletePolicyVersionInput) *IAMDeletePolicyVersionFuture

	DeleteRole(ctx workflow.Context, input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
	DeleteRoleAsync(ctx workflow.Context, input *iam.DeleteRoleInput) *IAMDeleteRoleFuture

	DeleteRolePermissionsBoundary(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error)
	DeleteRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) *IAMDeleteRolePermissionsBoundaryFuture

	DeleteRolePolicy(ctx workflow.Context, input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)
	DeleteRolePolicyAsync(ctx workflow.Context, input *iam.DeleteRolePolicyInput) *IAMDeleteRolePolicyFuture

	DeleteSAMLProvider(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderAsync(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) *IAMDeleteSAMLProviderFuture

	DeleteSSHPublicKey(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteSSHPublicKeyAsync(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) *IAMDeleteSSHPublicKeyFuture

	DeleteServerCertificate(ctx workflow.Context, input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)
	DeleteServerCertificateAsync(ctx workflow.Context, input *iam.DeleteServerCertificateInput) *IAMDeleteServerCertificateFuture

	DeleteServiceLinkedRole(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleAsync(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) *IAMDeleteServiceLinkedRoleFuture

	DeleteServiceSpecificCredential(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) *IAMDeleteServiceSpecificCredentialFuture

	DeleteSigningCertificate(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)
	DeleteSigningCertificateAsync(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) *IAMDeleteSigningCertificateFuture

	DeleteUser(ctx workflow.Context, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *iam.DeleteUserInput) *IAMDeleteUserFuture

	DeleteUserPermissionsBoundary(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error)
	DeleteUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) *IAMDeleteUserPermissionsBoundaryFuture

	DeleteUserPolicy(ctx workflow.Context, input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserPolicyAsync(ctx workflow.Context, input *iam.DeleteUserPolicyInput) *IAMDeleteUserPolicyFuture

	DeleteVirtualMFADevice(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)
	DeleteVirtualMFADeviceAsync(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) *IAMDeleteVirtualMFADeviceFuture

	DetachGroupPolicy(ctx workflow.Context, input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)
	DetachGroupPolicyAsync(ctx workflow.Context, input *iam.DetachGroupPolicyInput) *IAMDetachGroupPolicyFuture

	DetachRolePolicy(ctx workflow.Context, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)
	DetachRolePolicyAsync(ctx workflow.Context, input *iam.DetachRolePolicyInput) *IAMDetachRolePolicyFuture

	DetachUserPolicy(ctx workflow.Context, input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)
	DetachUserPolicyAsync(ctx workflow.Context, input *iam.DetachUserPolicyInput) *IAMDetachUserPolicyFuture

	EnableMFADevice(ctx workflow.Context, input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)
	EnableMFADeviceAsync(ctx workflow.Context, input *iam.EnableMFADeviceInput) *IAMEnableMFADeviceFuture

	GenerateCredentialReport(ctx workflow.Context, input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)
	GenerateCredentialReportAsync(ctx workflow.Context, input *iam.GenerateCredentialReportInput) *IAMGenerateCredentialReportFuture

	GenerateOrganizationsAccessReport(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error)
	GenerateOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) *IAMGenerateOrganizationsAccessReportFuture

	GenerateServiceLastAccessedDetails(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error)
	GenerateServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) *IAMGenerateServiceLastAccessedDetailsFuture

	GetAccessKeyLastUsed(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedAsync(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) *IAMGetAccessKeyLastUsedFuture

	GetAccountAuthorizationDetails(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountAuthorizationDetailsAsync(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) *IAMGetAccountAuthorizationDetailsFuture

	GetAccountPasswordPolicy(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) *IAMGetAccountPasswordPolicyFuture

	GetAccountSummary(ctx workflow.Context, input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)
	GetAccountSummaryAsync(ctx workflow.Context, input *iam.GetAccountSummaryInput) *IAMGetAccountSummaryFuture

	GetContextKeysForCustomPolicy(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForCustomPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) *IAMGetContextKeysForCustomPolicyFuture

	GetContextKeysForPrincipalPolicy(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) *IAMGetContextKeysForPrincipalPolicyFuture

	GetCredentialReport(ctx workflow.Context, input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)
	GetCredentialReportAsync(ctx workflow.Context, input *iam.GetCredentialReportInput) *IAMGetCredentialReportFuture

	GetGroup(ctx workflow.Context, input *iam.GetGroupInput) (*iam.GetGroupOutput, error)
	GetGroupAsync(ctx workflow.Context, input *iam.GetGroupInput) *IAMGetGroupFuture

	GetGroupPolicy(ctx workflow.Context, input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)
	GetGroupPolicyAsync(ctx workflow.Context, input *iam.GetGroupPolicyInput) *IAMGetGroupPolicyFuture

	GetInstanceProfile(ctx workflow.Context, input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)
	GetInstanceProfileAsync(ctx workflow.Context, input *iam.GetInstanceProfileInput) *IAMGetInstanceProfileFuture

	GetLoginProfile(ctx workflow.Context, input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)
	GetLoginProfileAsync(ctx workflow.Context, input *iam.GetLoginProfileInput) *IAMGetLoginProfileFuture

	GetOpenIDConnectProvider(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) *IAMGetOpenIDConnectProviderFuture

	GetOrganizationsAccessReport(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error)
	GetOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) *IAMGetOrganizationsAccessReportFuture

	GetPolicy(ctx workflow.Context, input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *iam.GetPolicyInput) *IAMGetPolicyFuture

	GetPolicyVersion(ctx workflow.Context, input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)
	GetPolicyVersionAsync(ctx workflow.Context, input *iam.GetPolicyVersionInput) *IAMGetPolicyVersionFuture

	GetRole(ctx workflow.Context, input *iam.GetRoleInput) (*iam.GetRoleOutput, error)
	GetRoleAsync(ctx workflow.Context, input *iam.GetRoleInput) *IAMGetRoleFuture

	GetRolePolicy(ctx workflow.Context, input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)
	GetRolePolicyAsync(ctx workflow.Context, input *iam.GetRolePolicyInput) *IAMGetRolePolicyFuture

	GetSAMLProvider(ctx workflow.Context, input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)
	GetSAMLProviderAsync(ctx workflow.Context, input *iam.GetSAMLProviderInput) *IAMGetSAMLProviderFuture

	GetSSHPublicKey(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)
	GetSSHPublicKeyAsync(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) *IAMGetSSHPublicKeyFuture

	GetServerCertificate(ctx workflow.Context, input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)
	GetServerCertificateAsync(ctx workflow.Context, input *iam.GetServerCertificateInput) *IAMGetServerCertificateFuture

	GetServiceLastAccessedDetails(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error)
	GetServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) *IAMGetServiceLastAccessedDetailsFuture

	GetServiceLastAccessedDetailsWithEntities(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
	GetServiceLastAccessedDetailsWithEntitiesAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) *IAMGetServiceLastAccessedDetailsWithEntitiesFuture

	GetServiceLinkedRoleDeletionStatus(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetServiceLinkedRoleDeletionStatusAsync(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) *IAMGetServiceLinkedRoleDeletionStatusFuture

	GetUser(ctx workflow.Context, input *iam.GetUserInput) (*iam.GetUserOutput, error)
	GetUserAsync(ctx workflow.Context, input *iam.GetUserInput) *IAMGetUserFuture

	GetUserPolicy(ctx workflow.Context, input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)
	GetUserPolicyAsync(ctx workflow.Context, input *iam.GetUserPolicyInput) *IAMGetUserPolicyFuture

	ListAccessKeys(ctx workflow.Context, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
	ListAccessKeysAsync(ctx workflow.Context, input *iam.ListAccessKeysInput) *IAMListAccessKeysFuture

	ListAccountAliases(ctx workflow.Context, input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)
	ListAccountAliasesAsync(ctx workflow.Context, input *iam.ListAccountAliasesInput) *IAMListAccountAliasesFuture

	ListAttachedGroupPolicies(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedGroupPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) *IAMListAttachedGroupPoliciesFuture

	ListAttachedRolePolicies(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesAsync(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) *IAMListAttachedRolePoliciesFuture

	ListAttachedUserPolicies(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) *IAMListAttachedUserPoliciesFuture

	ListEntitiesForPolicy(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyAsync(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) *IAMListEntitiesForPolicyFuture

	ListGroupPolicies(ctx workflow.Context, input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)
	ListGroupPoliciesAsync(ctx workflow.Context, input *iam.ListGroupPoliciesInput) *IAMListGroupPoliciesFuture

	ListGroups(ctx workflow.Context, input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *iam.ListGroupsInput) *IAMListGroupsFuture

	ListGroupsForUser(ctx workflow.Context, input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)
	ListGroupsForUserAsync(ctx workflow.Context, input *iam.ListGroupsForUserInput) *IAMListGroupsForUserFuture

	ListInstanceProfiles(ctx workflow.Context, input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesAsync(ctx workflow.Context, input *iam.ListInstanceProfilesInput) *IAMListInstanceProfilesFuture

	ListInstanceProfilesForRole(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesForRoleAsync(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) *IAMListInstanceProfilesForRoleFuture

	ListMFADevices(ctx workflow.Context, input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)
	ListMFADevicesAsync(ctx workflow.Context, input *iam.ListMFADevicesInput) *IAMListMFADevicesFuture

	ListOpenIDConnectProviders(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListOpenIDConnectProvidersAsync(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) *IAMListOpenIDConnectProvidersFuture

	ListPolicies(ctx workflow.Context, input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)
	ListPoliciesAsync(ctx workflow.Context, input *iam.ListPoliciesInput) *IAMListPoliciesFuture

	ListPoliciesGrantingServiceAccess(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
	ListPoliciesGrantingServiceAccessAsync(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) *IAMListPoliciesGrantingServiceAccessFuture

	ListPolicyVersions(ctx workflow.Context, input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)
	ListPolicyVersionsAsync(ctx workflow.Context, input *iam.ListPolicyVersionsInput) *IAMListPolicyVersionsFuture

	ListRolePolicies(ctx workflow.Context, input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)
	ListRolePoliciesAsync(ctx workflow.Context, input *iam.ListRolePoliciesInput) *IAMListRolePoliciesFuture

	ListRoleTags(ctx workflow.Context, input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error)
	ListRoleTagsAsync(ctx workflow.Context, input *iam.ListRoleTagsInput) *IAMListRoleTagsFuture

	ListRoles(ctx workflow.Context, input *iam.ListRolesInput) (*iam.ListRolesOutput, error)
	ListRolesAsync(ctx workflow.Context, input *iam.ListRolesInput) *IAMListRolesFuture

	ListSAMLProviders(ctx workflow.Context, input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)
	ListSAMLProvidersAsync(ctx workflow.Context, input *iam.ListSAMLProvidersInput) *IAMListSAMLProvidersFuture

	ListSSHPublicKeys(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)
	ListSSHPublicKeysAsync(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) *IAMListSSHPublicKeysFuture

	ListServerCertificates(ctx workflow.Context, input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
	ListServerCertificatesAsync(ctx workflow.Context, input *iam.ListServerCertificatesInput) *IAMListServerCertificatesFuture

	ListServiceSpecificCredentials(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListServiceSpecificCredentialsAsync(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) *IAMListServiceSpecificCredentialsFuture

	ListSigningCertificates(ctx workflow.Context, input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)
	ListSigningCertificatesAsync(ctx workflow.Context, input *iam.ListSigningCertificatesInput) *IAMListSigningCertificatesFuture

	ListUserPolicies(ctx workflow.Context, input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)
	ListUserPoliciesAsync(ctx workflow.Context, input *iam.ListUserPoliciesInput) *IAMListUserPoliciesFuture

	ListUserTags(ctx workflow.Context, input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error)
	ListUserTagsAsync(ctx workflow.Context, input *iam.ListUserTagsInput) *IAMListUserTagsFuture

	ListUsers(ctx workflow.Context, input *iam.ListUsersInput) (*iam.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *iam.ListUsersInput) *IAMListUsersFuture

	ListVirtualMFADevices(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)
	ListVirtualMFADevicesAsync(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) *IAMListVirtualMFADevicesFuture

	PutGroupPolicy(ctx workflow.Context, input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)
	PutGroupPolicyAsync(ctx workflow.Context, input *iam.PutGroupPolicyInput) *IAMPutGroupPolicyFuture

	PutRolePermissionsBoundary(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error)
	PutRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) *IAMPutRolePermissionsBoundaryFuture

	PutRolePolicy(ctx workflow.Context, input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)
	PutRolePolicyAsync(ctx workflow.Context, input *iam.PutRolePolicyInput) *IAMPutRolePolicyFuture

	PutUserPermissionsBoundary(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error)
	PutUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) *IAMPutUserPermissionsBoundaryFuture

	PutUserPolicy(ctx workflow.Context, input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)
	PutUserPolicyAsync(ctx workflow.Context, input *iam.PutUserPolicyInput) *IAMPutUserPolicyFuture

	RemoveClientIDFromOpenIDConnectProvider(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveClientIDFromOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) *IAMRemoveClientIDFromOpenIDConnectProviderFuture

	RemoveRoleFromInstanceProfile(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfileAsync(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) *IAMRemoveRoleFromInstanceProfileFuture

	RemoveUserFromGroup(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupAsync(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) *IAMRemoveUserFromGroupFuture

	ResetServiceSpecificCredential(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResetServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) *IAMResetServiceSpecificCredentialFuture

	ResyncMFADevice(ctx workflow.Context, input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)
	ResyncMFADeviceAsync(ctx workflow.Context, input *iam.ResyncMFADeviceInput) *IAMResyncMFADeviceFuture

	SetDefaultPolicyVersion(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionAsync(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) *IAMSetDefaultPolicyVersionFuture

	SetSecurityTokenServicePreferences(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error)
	SetSecurityTokenServicePreferencesAsync(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) *IAMSetSecurityTokenServicePreferencesFuture

	SimulateCustomPolicy(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulateCustomPolicyAsync(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) *IAMSimulateCustomPolicyFuture

	SimulatePrincipalPolicy(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyAsync(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) *IAMSimulatePrincipalPolicyFuture

	TagRole(ctx workflow.Context, input *iam.TagRoleInput) (*iam.TagRoleOutput, error)
	TagRoleAsync(ctx workflow.Context, input *iam.TagRoleInput) *IAMTagRoleFuture

	TagUser(ctx workflow.Context, input *iam.TagUserInput) (*iam.TagUserOutput, error)
	TagUserAsync(ctx workflow.Context, input *iam.TagUserInput) *IAMTagUserFuture

	UntagRole(ctx workflow.Context, input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error)
	UntagRoleAsync(ctx workflow.Context, input *iam.UntagRoleInput) *IAMUntagRoleFuture

	UntagUser(ctx workflow.Context, input *iam.UntagUserInput) (*iam.UntagUserOutput, error)
	UntagUserAsync(ctx workflow.Context, input *iam.UntagUserInput) *IAMUntagUserFuture

	UpdateAccessKey(ctx workflow.Context, input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccessKeyAsync(ctx workflow.Context, input *iam.UpdateAccessKeyInput) *IAMUpdateAccessKeyFuture

	UpdateAccountPasswordPolicy(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) *IAMUpdateAccountPasswordPolicyFuture

	UpdateAssumeRolePolicy(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateAssumeRolePolicyAsync(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) *IAMUpdateAssumeRolePolicyFuture

	UpdateGroup(ctx workflow.Context, input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *iam.UpdateGroupInput) *IAMUpdateGroupFuture

	UpdateLoginProfile(ctx workflow.Context, input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)
	UpdateLoginProfileAsync(ctx workflow.Context, input *iam.UpdateLoginProfileInput) *IAMUpdateLoginProfileFuture

	UpdateOpenIDConnectProviderThumbprint(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateOpenIDConnectProviderThumbprintAsync(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) *IAMUpdateOpenIDConnectProviderThumbprintFuture

	UpdateRole(ctx workflow.Context, input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error)
	UpdateRoleAsync(ctx workflow.Context, input *iam.UpdateRoleInput) *IAMUpdateRoleFuture

	UpdateRoleDescription(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateRoleDescriptionAsync(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) *IAMUpdateRoleDescriptionFuture

	UpdateSAMLProvider(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderAsync(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) *IAMUpdateSAMLProviderFuture

	UpdateSSHPublicKey(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateSSHPublicKeyAsync(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) *IAMUpdateSSHPublicKeyFuture

	UpdateServerCertificate(ctx workflow.Context, input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)
	UpdateServerCertificateAsync(ctx workflow.Context, input *iam.UpdateServerCertificateInput) *IAMUpdateServerCertificateFuture

	UpdateServiceSpecificCredential(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) *IAMUpdateServiceSpecificCredentialFuture

	UpdateSigningCertificate(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)
	UpdateSigningCertificateAsync(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) *IAMUpdateSigningCertificateFuture

	UpdateUser(ctx workflow.Context, input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *iam.UpdateUserInput) *IAMUpdateUserFuture

	UploadSSHPublicKey(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)
	UploadSSHPublicKeyAsync(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) *IAMUploadSSHPublicKeyFuture

	UploadServerCertificate(ctx workflow.Context, input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)
	UploadServerCertificateAsync(ctx workflow.Context, input *iam.UploadServerCertificateInput) *IAMUploadServerCertificateFuture

	UploadSigningCertificate(ctx workflow.Context, input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
	UploadSigningCertificateAsync(ctx workflow.Context, input *iam.UploadSigningCertificateInput) *IAMUploadSigningCertificateFuture

	WaitUntilInstanceProfileExists(ctx workflow.Context, input *iam.GetInstanceProfileInput) error
	WaitUntilInstanceProfileExistsAsync(ctx workflow.Context, input *iam.GetInstanceProfileInput) *awsclients.VoidFuture

	WaitUntilPolicyExists(ctx workflow.Context, input *iam.GetPolicyInput) error
	WaitUntilPolicyExistsAsync(ctx workflow.Context, input *iam.GetPolicyInput) *awsclients.VoidFuture

	WaitUntilRoleExists(ctx workflow.Context, input *iam.GetRoleInput) error
	WaitUntilRoleExistsAsync(ctx workflow.Context, input *iam.GetRoleInput) *awsclients.VoidFuture

	WaitUntilUserExists(ctx workflow.Context, input *iam.GetUserInput) error
	WaitUntilUserExistsAsync(ctx workflow.Context, input *iam.GetUserInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

