// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WorkLinkActivities struct {
	client worklinkiface.WorkLinkAPI

	sessionFactory SessionFactory
}

func NewWorkLinkActivities(sess *session.Session, config ...*aws.Config) *WorkLinkActivities {
	client := worklink.New(sess, config...)
	return &WorkLinkActivities{client: client}
}

func NewWorkLinkActivitiesWithSessionFactory(sessionFactory SessionFactory) *WorkLinkActivities {
	return &WorkLinkActivities{sessionFactory: sessionFactory}
}

func (a *WorkLinkActivities) getClient(ctx context.Context) (worklinkiface.WorkLinkAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return worklink.New(sess), nil
}

func (a *WorkLinkActivities) AssociateDomain(ctx context.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) AssociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateWebsiteAuthorizationProviderWithContext(ctx, input)
}

func (a *WorkLinkActivities) AssociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) CreateFleet(ctx context.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFleetWithContext(ctx, input)
}

func (a *WorkLinkActivities) DeleteFleet(ctx context.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFleetWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeAuditStreamConfiguration(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAuditStreamConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeCompanyNetworkConfiguration(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCompanyNetworkConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDevice(ctx context.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDeviceWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDevicePolicyConfiguration(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDevicePolicyConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDomain(ctx context.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeFleetMetadata(ctx context.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeFleetMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeIdentityProviderConfiguration(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityProviderConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeWebsiteCertificateAuthority(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateDomain(ctx context.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateWebsiteAuthorizationProviderWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListDevices(ctx context.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDevicesWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListDomains(ctx context.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDomainsWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListFleets(ctx context.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFleetsWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListTagsForResource(ctx context.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListWebsiteAuthorizationProviders(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListWebsiteAuthorizationProvidersWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListWebsiteCertificateAuthorities(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListWebsiteCertificateAuthoritiesWithContext(ctx, input)
}

func (a *WorkLinkActivities) RestoreDomainAccess(ctx context.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreDomainAccessWithContext(ctx, input)
}

func (a *WorkLinkActivities) RevokeDomainAccess(ctx context.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RevokeDomainAccessWithContext(ctx, input)
}

func (a *WorkLinkActivities) SignOutUser(ctx context.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SignOutUserWithContext(ctx, input)
}

func (a *WorkLinkActivities) TagResource(ctx context.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) UntagResource(ctx context.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateAuditStreamConfiguration(ctx context.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAuditStreamConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateCompanyNetworkConfiguration(ctx context.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateCompanyNetworkConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateDevicePolicyConfiguration(ctx context.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDevicePolicyConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateDomainMetadata(ctx context.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDomainMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateFleetMetadata(ctx context.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFleetMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateIdentityProviderConfiguration(ctx context.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIdentityProviderConfigurationWithContext(ctx, input)
}
