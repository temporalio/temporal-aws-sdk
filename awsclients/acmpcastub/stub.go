// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package acmpcastub

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type ACMPCACreateCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCACreateCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCACreateCertificateAuthorityAuditReportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCACreateCertificateAuthorityAuditReportFuture) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCACreatePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCACreatePermissionFuture) Get(ctx workflow.Context) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCADeleteCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCADeleteCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCADeletePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCADeletePermissionFuture) Get(ctx workflow.Context) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCADeletePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCADeletePolicyFuture) Get(ctx workflow.Context) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCADescribeCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCADescribeCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCADescribeCertificateAuthorityAuditReportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCADescribeCertificateAuthorityAuditReportFuture) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAGetCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAGetCertificateFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAGetCertificateAuthorityCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAGetCertificateAuthorityCertificateFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAGetCertificateAuthorityCsrFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAGetCertificateAuthorityCsrFuture) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAGetPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAGetPolicyFuture) Get(ctx workflow.Context) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAImportCertificateAuthorityCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAImportCertificateAuthorityCertificateFuture) Get(ctx workflow.Context) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAIssueCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAIssueCertificateFuture) Get(ctx workflow.Context) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAListCertificateAuthoritiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAListCertificateAuthoritiesFuture) Get(ctx workflow.Context) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAListPermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAListPermissionsFuture) Get(ctx workflow.Context) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAListTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAListTagsFuture) Get(ctx workflow.Context) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAPutPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAPutPolicyFuture) Get(ctx workflow.Context) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCARestoreCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCARestoreCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCARevokeCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCARevokeCertificateFuture) Get(ctx workflow.Context) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCATagCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCATagCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAUntagCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAUntagCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ACMPCAUpdateCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ACMPCAUpdateCertificateAuthorityFuture) Get(ctx workflow.Context) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
	var output acmpca.CreateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *ACMPCACreateCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthority", input)
	return &ACMPCACreateCertificateAuthorityFuture{Future: future}
}

func (a *stub) CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.CreateCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *ACMPCACreateCertificateAuthorityAuditReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreateCertificateAuthorityAuditReport", input)
	return &ACMPCACreateCertificateAuthorityAuditReportFuture{Future: future}
}

func (a *stub) CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
	var output acmpca.CreatePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.CreatePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *ACMPCACreatePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.CreatePermission", input)
	return &ACMPCACreatePermissionFuture{Future: future}
}

func (a *stub) DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	var output acmpca.DeleteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeleteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *ACMPCADeleteCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeleteCertificateAuthority", input)
	return &ACMPCADeleteCertificateAuthorityFuture{Future: future}
}

func (a *stub) DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
	var output acmpca.DeletePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *ACMPCADeletePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePermission", input)
	return &ACMPCADeletePermissionFuture{Future: future}
}

func (a *stub) DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
	var output acmpca.DeletePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *ACMPCADeletePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DeletePolicy", input)
	return &ACMPCADeletePolicyFuture{Future: future}
}

func (a *stub) DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	var output acmpca.DescribeCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *ACMPCADescribeCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthority", input)
	return &ACMPCADescribeCertificateAuthorityFuture{Future: future}
}

func (a *stub) DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	var output acmpca.DescribeCertificateAuthorityAuditReportOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthorityAuditReport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *ACMPCADescribeCertificateAuthorityAuditReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.DescribeCertificateAuthorityAuditReport", input)
	return &ACMPCADescribeCertificateAuthorityAuditReportFuture{Future: future}
}

func (a *stub) GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	var output acmpca.GetCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *ACMPCAGetCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificate", input)
	return &ACMPCAGetCertificateFuture{Future: future}
}

func (a *stub) GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	var output acmpca.GetCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *ACMPCAGetCertificateAuthorityCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCertificate", input)
	return &ACMPCAGetCertificateAuthorityCertificateFuture{Future: future}
}

func (a *stub) GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	var output acmpca.GetCertificateAuthorityCsrOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCsr", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *ACMPCAGetCertificateAuthorityCsrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetCertificateAuthorityCsr", input)
	return &ACMPCAGetCertificateAuthorityCsrFuture{Future: future}
}

func (a *stub) GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	var output acmpca.GetPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.GetPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *ACMPCAGetPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.GetPolicy", input)
	return &ACMPCAGetPolicyFuture{Future: future}
}

func (a *stub) ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	var output acmpca.ImportCertificateAuthorityCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ImportCertificateAuthorityCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *ACMPCAImportCertificateAuthorityCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ImportCertificateAuthorityCertificate", input)
	return &ACMPCAImportCertificateAuthorityCertificateFuture{Future: future}
}

func (a *stub) IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
	var output acmpca.IssueCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.IssueCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *ACMPCAIssueCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.IssueCertificate", input)
	return &ACMPCAIssueCertificateFuture{Future: future}
}

func (a *stub) ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	var output acmpca.ListCertificateAuthoritiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListCertificateAuthorities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *ACMPCAListCertificateAuthoritiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListCertificateAuthorities", input)
	return &ACMPCAListCertificateAuthoritiesFuture{Future: future}
}

func (a *stub) ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	var output acmpca.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *ACMPCAListPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListPermissions", input)
	return &ACMPCAListPermissionsFuture{Future: future}
}

func (a *stub) ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	var output acmpca.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *ACMPCAListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.ListTags", input)
	return &ACMPCAListTagsFuture{Future: future}
}

func (a *stub) PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
	var output acmpca.PutPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.PutPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *ACMPCAPutPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.PutPolicy", input)
	return &ACMPCAPutPolicyFuture{Future: future}
}

func (a *stub) RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	var output acmpca.RestoreCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.RestoreCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *ACMPCARestoreCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.RestoreCertificateAuthority", input)
	return &ACMPCARestoreCertificateAuthorityFuture{Future: future}
}

func (a *stub) RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
	var output acmpca.RevokeCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.RevokeCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *ACMPCARevokeCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.RevokeCertificate", input)
	return &ACMPCARevokeCertificateFuture{Future: future}
}

func (a *stub) TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
	var output acmpca.TagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.TagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *ACMPCATagCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.TagCertificateAuthority", input)
	return &ACMPCATagCertificateAuthorityFuture{Future: future}
}

func (a *stub) UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
	var output acmpca.UntagCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.UntagCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *ACMPCAUntagCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.UntagCertificateAuthority", input)
	return &ACMPCAUntagCertificateAuthorityFuture{Future: future}
}

func (a *stub) UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	var output acmpca.UpdateCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.acmpca.UpdateCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *ACMPCAUpdateCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.UpdateCertificateAuthority", input)
	return &ACMPCAUpdateCertificateAuthorityFuture{Future: future}
}

func (a *stub) WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilAuditReportCreated", input).Get(ctx, nil)
}

func (a *stub) WaitUntilAuditReportCreatedAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *awsclients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilAuditReportCreated", input)
	return awsclients.NewVoidFuture(future)
}

func (a *stub) WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateAuthorityCSRCreated", input).Get(ctx, nil)
}

func (a *stub) WaitUntilCertificateAuthorityCSRCreatedAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *awsclients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateAuthorityCSRCreated", input)
	return awsclients.NewVoidFuture(future)
}

func (a *stub) WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error {
	return workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateIssued", input).Get(ctx, nil)
}

func (a *stub) WaitUntilCertificateIssuedAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *awsclients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.acmpca.WaitUntilCertificateIssued", input)
	return awsclients.NewVoidFuture(future)
}
