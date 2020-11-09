// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cognitoidentitystub

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateIdentityPoolFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateIdentityPoolFuture) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIdentitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIdentitiesFuture) Get(ctx workflow.Context) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	var output cognitoidentity.DeleteIdentitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIdentityPoolFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIdentityPoolFuture) Get(ctx workflow.Context) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	var output cognitoidentity.DeleteIdentityPoolOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.IdentityDescription, error) {
	var output cognitoidentity.IdentityDescription
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeIdentityPoolFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeIdentityPoolFuture) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetCredentialsForIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetCredentialsForIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	var output cognitoidentity.GetCredentialsForIdentityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetIdFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetIdFuture) Get(ctx workflow.Context) (*cognitoidentity.GetIdOutput, error) {
	var output cognitoidentity.GetIdOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetIdentityPoolRolesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetIdentityPoolRolesFuture) Get(ctx workflow.Context) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.GetIdentityPoolRolesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetOpenIdTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetOpenIdTokenFuture) Get(ctx workflow.Context) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	var output cognitoidentity.GetOpenIdTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetOpenIdTokenForDeveloperIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetOpenIdTokenForDeveloperIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	var output cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListIdentitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListIdentitiesFuture) Get(ctx workflow.Context) (*cognitoidentity.ListIdentitiesOutput, error) {
	var output cognitoidentity.ListIdentitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListIdentityPoolsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListIdentityPoolsFuture) Get(ctx workflow.Context) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	var output cognitoidentity.ListIdentityPoolsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*cognitoidentity.ListTagsForResourceOutput, error) {
	var output cognitoidentity.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LookupDeveloperIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LookupDeveloperIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	var output cognitoidentity.LookupDeveloperIdentityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MergeDeveloperIdentitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MergeDeveloperIdentitiesFuture) Get(ctx workflow.Context) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	var output cognitoidentity.MergeDeveloperIdentitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetIdentityPoolRolesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetIdentityPoolRolesFuture) Get(ctx workflow.Context) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.SetIdentityPoolRolesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*cognitoidentity.TagResourceOutput, error) {
	var output cognitoidentity.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UnlinkDeveloperIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UnlinkDeveloperIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	var output cognitoidentity.UnlinkDeveloperIdentityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UnlinkIdentityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UnlinkIdentityFuture) Get(ctx workflow.Context) (*cognitoidentity.UnlinkIdentityOutput, error) {
	var output cognitoidentity.UnlinkIdentityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*cognitoidentity.UntagResourceOutput, error) {
	var output cognitoidentity.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateIdentityPoolFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateIdentityPoolFuture) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIdentityPool(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.CreateIdentityPool", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) *CreateIdentityPoolFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.CreateIdentityPool", input)
	return &CreateIdentityPoolFuture{Future: future}
}

func (a *stub) DeleteIdentities(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	var output cognitoidentity.DeleteIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DeleteIdentities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) *DeleteIdentitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DeleteIdentities", input)
	return &DeleteIdentitiesFuture{Future: future}
}

func (a *stub) DeleteIdentityPool(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	var output cognitoidentity.DeleteIdentityPoolOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DeleteIdentityPool", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) *DeleteIdentityPoolFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DeleteIdentityPool", input)
	return &DeleteIdentityPoolFuture{Future: future}
}

func (a *stub) DescribeIdentity(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error) {
	var output cognitoidentity.IdentityDescription
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DescribeIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIdentityAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) *DescribeIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DescribeIdentity", input)
	return &DescribeIdentityFuture{Future: future}
}

func (a *stub) DescribeIdentityPool(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DescribeIdentityPool", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) *DescribeIdentityPoolFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.DescribeIdentityPool", input)
	return &DescribeIdentityPoolFuture{Future: future}
}

func (a *stub) GetCredentialsForIdentity(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	var output cognitoidentity.GetCredentialsForIdentityOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetCredentialsForIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCredentialsForIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) *GetCredentialsForIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetCredentialsForIdentity", input)
	return &GetCredentialsForIdentityFuture{Future: future}
}

func (a *stub) GetId(ctx workflow.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
	var output cognitoidentity.GetIdOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetId", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIdAsync(ctx workflow.Context, input *cognitoidentity.GetIdInput) *GetIdFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetId", input)
	return &GetIdFuture{Future: future}
}

func (a *stub) GetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.GetIdentityPoolRolesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetIdentityPoolRoles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) *GetIdentityPoolRolesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetIdentityPoolRoles", input)
	return &GetIdentityPoolRolesFuture{Future: future}
}

func (a *stub) GetOpenIdToken(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	var output cognitoidentity.GetOpenIdTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetOpenIdToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetOpenIdTokenAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) *GetOpenIdTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetOpenIdToken", input)
	return &GetOpenIdTokenFuture{Future: future}
}

func (a *stub) GetOpenIdTokenForDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	var output cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetOpenIdTokenForDeveloperIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetOpenIdTokenForDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) *GetOpenIdTokenForDeveloperIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.GetOpenIdTokenForDeveloperIdentity", input)
	return &GetOpenIdTokenForDeveloperIdentityFuture{Future: future}
}

func (a *stub) ListIdentities(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
	var output cognitoidentity.ListIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListIdentities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) *ListIdentitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListIdentities", input)
	return &ListIdentitiesFuture{Future: future}
}

func (a *stub) ListIdentityPools(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	var output cognitoidentity.ListIdentityPoolsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListIdentityPools", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIdentityPoolsAsync(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) *ListIdentityPoolsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListIdentityPools", input)
	return &ListIdentityPoolsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
	var output cognitoidentity.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) LookupDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	var output cognitoidentity.LookupDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.LookupDeveloperIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) LookupDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) *LookupDeveloperIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.LookupDeveloperIdentity", input)
	return &LookupDeveloperIdentityFuture{Future: future}
}

func (a *stub) MergeDeveloperIdentities(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	var output cognitoidentity.MergeDeveloperIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.MergeDeveloperIdentities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) MergeDeveloperIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) *MergeDeveloperIdentitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.MergeDeveloperIdentities", input)
	return &MergeDeveloperIdentitiesFuture{Future: future}
}

func (a *stub) SetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.SetIdentityPoolRolesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.SetIdentityPoolRoles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) *SetIdentityPoolRolesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.SetIdentityPoolRoles", input)
	return &SetIdentityPoolRolesFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error) {
	var output cognitoidentity.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *cognitoidentity.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UnlinkDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	var output cognitoidentity.UnlinkDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UnlinkDeveloperIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnlinkDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) *UnlinkDeveloperIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UnlinkDeveloperIdentity", input)
	return &UnlinkDeveloperIdentityFuture{Future: future}
}

func (a *stub) UnlinkIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error) {
	var output cognitoidentity.UnlinkIdentityOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UnlinkIdentity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnlinkIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) *UnlinkIdentityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UnlinkIdentity", input)
	return &UnlinkIdentityFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error) {
	var output cognitoidentity.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateIdentityPool(ctx workflow.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UpdateIdentityPool", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.IdentityPool) *UpdateIdentityPoolFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitoidentity.UpdateIdentityPool", input)
	return &UpdateIdentityPoolFuture{Future: future}
}
