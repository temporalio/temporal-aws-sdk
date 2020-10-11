// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package servicequotasstub

import (
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error)
	AssociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) *ServiceQuotasAssociateServiceQuotaTemplateFuture

	DeleteServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error)
	DeleteServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) *ServiceQuotasDeleteServiceQuotaIncreaseRequestFromTemplateFuture

	DisassociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error)
	DisassociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) *ServiceQuotasDisassociateServiceQuotaTemplateFuture

	GetAWSDefaultServiceQuota(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
	GetAWSDefaultServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) *ServiceQuotasGetAWSDefaultServiceQuotaFuture

	GetAssociationForServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
	GetAssociationForServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) *ServiceQuotasGetAssociationForServiceQuotaTemplateFuture

	GetRequestedServiceQuotaChange(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
	GetRequestedServiceQuotaChangeAsync(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) *ServiceQuotasGetRequestedServiceQuotaChangeFuture

	GetServiceQuota(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error)
	GetServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) *ServiceQuotasGetServiceQuotaFuture

	GetServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
	GetServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) *ServiceQuotasGetServiceQuotaIncreaseRequestFromTemplateFuture

	ListAWSDefaultServiceQuotas(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
	ListAWSDefaultServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) *ServiceQuotasListAWSDefaultServiceQuotasFuture

	ListRequestedServiceQuotaChangeHistory(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
	ListRequestedServiceQuotaChangeHistoryAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) *ServiceQuotasListRequestedServiceQuotaChangeHistoryFuture

	ListRequestedServiceQuotaChangeHistoryByQuota(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
	ListRequestedServiceQuotaChangeHistoryByQuotaAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) *ServiceQuotasListRequestedServiceQuotaChangeHistoryByQuotaFuture

	ListServiceQuotaIncreaseRequestsInTemplate(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
	ListServiceQuotaIncreaseRequestsInTemplateAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) *ServiceQuotasListServiceQuotaIncreaseRequestsInTemplateFuture

	ListServiceQuotas(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error)
	ListServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) *ServiceQuotasListServiceQuotasFuture

	ListServices(ctx workflow.Context, input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error)
	ListServicesAsync(ctx workflow.Context, input *servicequotas.ListServicesInput) *ServiceQuotasListServicesFuture

	PutServiceQuotaIncreaseRequestIntoTemplate(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error)
	PutServiceQuotaIncreaseRequestIntoTemplateAsync(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) *ServiceQuotasPutServiceQuotaIncreaseRequestIntoTemplateFuture

	RequestServiceQuotaIncrease(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error)
	RequestServiceQuotaIncreaseAsync(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) *ServiceQuotasRequestServiceQuotaIncreaseFuture
}

func NewClient() Client {
	return &stub{}
}
