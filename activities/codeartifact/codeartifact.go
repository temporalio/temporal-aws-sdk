// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client codeartifactiface.CodeArtifactAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := codeartifact.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (codeartifactiface.CodeArtifactAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return codeartifact.New(sess), nil
}

func (a *Activities) AssociateExternalConnection(ctx context.Context, input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateExternalConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyPackageVersions(ctx context.Context, input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyPackageVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDomain(ctx context.Context, input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRepository(ctx context.Context, input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDomain(ctx context.Context, input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDomainPermissionsPolicy(ctx context.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDomainPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePackageVersions(ctx context.Context, input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePackageVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRepository(ctx context.Context, input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRepositoryPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDomain(ctx context.Context, input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePackageVersion(ctx context.Context, input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePackageVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRepository(ctx context.Context, input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateExternalConnection(ctx context.Context, input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateExternalConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisposePackageVersions(ctx context.Context, input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisposePackageVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAuthorizationToken(ctx context.Context, input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAuthorizationTokenWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDomainPermissionsPolicy(ctx context.Context, input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDomainPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPackageVersionAsset(ctx context.Context, input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPackageVersionAssetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPackageVersionReadme(ctx context.Context, input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPackageVersionReadmeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRepositoryEndpoint(ctx context.Context, input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRepositoryEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRepositoryPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDomains(ctx context.Context, input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDomainsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPackageVersionAssets(ctx context.Context, input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPackageVersionAssetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPackageVersionDependencies(ctx context.Context, input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPackageVersionDependenciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPackageVersions(ctx context.Context, input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPackageVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPackages(ctx context.Context, input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPackagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRepositories(ctx context.Context, input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRepositoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRepositoriesInDomain(ctx context.Context, input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRepositoriesInDomainWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutDomainPermissionsPolicy(ctx context.Context, input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutDomainPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRepositoryPermissionsPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePackageVersionsStatus(ctx context.Context, input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePackageVersionsStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRepository(ctx context.Context, input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRepositoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
