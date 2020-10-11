// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package worklinkstub

import (
	"github.com/aws/aws-sdk-go/service/worklink"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssociateDomain(ctx workflow.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error)
	AssociateDomainAsync(ctx workflow.Context, input *worklink.AssociateDomainInput) *WorkLinkAssociateDomainFuture

	AssociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error)
	AssociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) *WorkLinkAssociateWebsiteAuthorizationProviderFuture

	AssociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error)
	AssociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) *WorkLinkAssociateWebsiteCertificateAuthorityFuture

	CreateFleet(ctx workflow.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error)
	CreateFleetAsync(ctx workflow.Context, input *worklink.CreateFleetInput) *WorkLinkCreateFleetFuture

	DeleteFleet(ctx workflow.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error)
	DeleteFleetAsync(ctx workflow.Context, input *worklink.DeleteFleetInput) *WorkLinkDeleteFleetFuture

	DescribeAuditStreamConfiguration(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error)
	DescribeAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) *WorkLinkDescribeAuditStreamConfigurationFuture

	DescribeCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error)
	DescribeCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) *WorkLinkDescribeCompanyNetworkConfigurationFuture

	DescribeDevice(ctx workflow.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error)
	DescribeDeviceAsync(ctx workflow.Context, input *worklink.DescribeDeviceInput) *WorkLinkDescribeDeviceFuture

	DescribeDevicePolicyConfiguration(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error)
	DescribeDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) *WorkLinkDescribeDevicePolicyConfigurationFuture

	DescribeDomain(ctx workflow.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error)
	DescribeDomainAsync(ctx workflow.Context, input *worklink.DescribeDomainInput) *WorkLinkDescribeDomainFuture

	DescribeFleetMetadata(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error)
	DescribeFleetMetadataAsync(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) *WorkLinkDescribeFleetMetadataFuture

	DescribeIdentityProviderConfiguration(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error)
	DescribeIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) *WorkLinkDescribeIdentityProviderConfigurationFuture

	DescribeWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error)
	DescribeWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) *WorkLinkDescribeWebsiteCertificateAuthorityFuture

	DisassociateDomain(ctx workflow.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error)
	DisassociateDomainAsync(ctx workflow.Context, input *worklink.DisassociateDomainInput) *WorkLinkDisassociateDomainFuture

	DisassociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error)
	DisassociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) *WorkLinkDisassociateWebsiteAuthorizationProviderFuture

	DisassociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error)
	DisassociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) *WorkLinkDisassociateWebsiteCertificateAuthorityFuture

	ListDevices(ctx workflow.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error)
	ListDevicesAsync(ctx workflow.Context, input *worklink.ListDevicesInput) *WorkLinkListDevicesFuture

	ListDomains(ctx workflow.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *worklink.ListDomainsInput) *WorkLinkListDomainsFuture

	ListFleets(ctx workflow.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error)
	ListFleetsAsync(ctx workflow.Context, input *worklink.ListFleetsInput) *WorkLinkListFleetsFuture

	ListTagsForResource(ctx workflow.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *worklink.ListTagsForResourceInput) *WorkLinkListTagsForResourceFuture

	ListWebsiteAuthorizationProviders(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error)
	ListWebsiteAuthorizationProvidersAsync(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) *WorkLinkListWebsiteAuthorizationProvidersFuture

	ListWebsiteCertificateAuthorities(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error)
	ListWebsiteCertificateAuthoritiesAsync(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) *WorkLinkListWebsiteCertificateAuthoritiesFuture

	RestoreDomainAccess(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error)
	RestoreDomainAccessAsync(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) *WorkLinkRestoreDomainAccessFuture

	RevokeDomainAccess(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error)
	RevokeDomainAccessAsync(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) *WorkLinkRevokeDomainAccessFuture

	SignOutUser(ctx workflow.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error)
	SignOutUserAsync(ctx workflow.Context, input *worklink.SignOutUserInput) *WorkLinkSignOutUserFuture

	TagResource(ctx workflow.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *worklink.TagResourceInput) *WorkLinkTagResourceFuture

	UntagResource(ctx workflow.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *worklink.UntagResourceInput) *WorkLinkUntagResourceFuture

	UpdateAuditStreamConfiguration(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error)
	UpdateAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) *WorkLinkUpdateAuditStreamConfigurationFuture

	UpdateCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error)
	UpdateCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) *WorkLinkUpdateCompanyNetworkConfigurationFuture

	UpdateDevicePolicyConfiguration(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error)
	UpdateDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) *WorkLinkUpdateDevicePolicyConfigurationFuture

	UpdateDomainMetadata(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error)
	UpdateDomainMetadataAsync(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) *WorkLinkUpdateDomainMetadataFuture

	UpdateFleetMetadata(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error)
	UpdateFleetMetadataAsync(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) *WorkLinkUpdateFleetMetadataFuture

	UpdateIdentityProviderConfiguration(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error)
	UpdateIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) *WorkLinkUpdateIdentityProviderConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}