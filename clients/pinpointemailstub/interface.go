// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package pinpointemailstub

import (
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateConfigurationSet(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error)
	CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetInput) *PinpointEmailCreateConfigurationSetFuture

	CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) *PinpointEmailCreateConfigurationSetEventDestinationFuture

	CreateDedicatedIpPool(ctx workflow.Context, input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error)
	CreateDedicatedIpPoolAsync(ctx workflow.Context, input *pinpointemail.CreateDedicatedIpPoolInput) *PinpointEmailCreateDedicatedIpPoolFuture

	CreateDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error)
	CreateDeliverabilityTestReportAsync(ctx workflow.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) *PinpointEmailCreateDeliverabilityTestReportFuture

	CreateEmailIdentity(ctx workflow.Context, input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error)
	CreateEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.CreateEmailIdentityInput) *PinpointEmailCreateEmailIdentityFuture

	DeleteConfigurationSet(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetInput) *PinpointEmailDeleteConfigurationSetFuture

	DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) *PinpointEmailDeleteConfigurationSetEventDestinationFuture

	DeleteDedicatedIpPool(ctx workflow.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error)
	DeleteDedicatedIpPoolAsync(ctx workflow.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) *PinpointEmailDeleteDedicatedIpPoolFuture

	DeleteEmailIdentity(ctx workflow.Context, input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error)
	DeleteEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.DeleteEmailIdentityInput) *PinpointEmailDeleteEmailIdentityFuture

	GetAccount(ctx workflow.Context, input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error)
	GetAccountAsync(ctx workflow.Context, input *pinpointemail.GetAccountInput) *PinpointEmailGetAccountFuture

	GetBlacklistReports(ctx workflow.Context, input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error)
	GetBlacklistReportsAsync(ctx workflow.Context, input *pinpointemail.GetBlacklistReportsInput) *PinpointEmailGetBlacklistReportsFuture

	GetConfigurationSet(ctx workflow.Context, input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error)
	GetConfigurationSetAsync(ctx workflow.Context, input *pinpointemail.GetConfigurationSetInput) *PinpointEmailGetConfigurationSetFuture

	GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) *PinpointEmailGetConfigurationSetEventDestinationsFuture

	GetDedicatedIp(ctx workflow.Context, input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error)
	GetDedicatedIpAsync(ctx workflow.Context, input *pinpointemail.GetDedicatedIpInput) *PinpointEmailGetDedicatedIpFuture

	GetDedicatedIps(ctx workflow.Context, input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error)
	GetDedicatedIpsAsync(ctx workflow.Context, input *pinpointemail.GetDedicatedIpsInput) *PinpointEmailGetDedicatedIpsFuture

	GetDeliverabilityDashboardOptions(ctx workflow.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityDashboardOptionsAsync(ctx workflow.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) *PinpointEmailGetDeliverabilityDashboardOptionsFuture

	GetDeliverabilityTestReport(ctx workflow.Context, input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error)
	GetDeliverabilityTestReportAsync(ctx workflow.Context, input *pinpointemail.GetDeliverabilityTestReportInput) *PinpointEmailGetDeliverabilityTestReportFuture

	GetDomainDeliverabilityCampaign(ctx workflow.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainDeliverabilityCampaignAsync(ctx workflow.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) *PinpointEmailGetDomainDeliverabilityCampaignFuture

	GetDomainStatisticsReport(ctx workflow.Context, input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error)
	GetDomainStatisticsReportAsync(ctx workflow.Context, input *pinpointemail.GetDomainStatisticsReportInput) *PinpointEmailGetDomainStatisticsReportFuture

	GetEmailIdentity(ctx workflow.Context, input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error)
	GetEmailIdentityAsync(ctx workflow.Context, input *pinpointemail.GetEmailIdentityInput) *PinpointEmailGetEmailIdentityFuture

	ListConfigurationSets(ctx workflow.Context, input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error)
	ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointemail.ListConfigurationSetsInput) *PinpointEmailListConfigurationSetsFuture

	ListDedicatedIpPools(ctx workflow.Context, input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error)
	ListDedicatedIpPoolsAsync(ctx workflow.Context, input *pinpointemail.ListDedicatedIpPoolsInput) *PinpointEmailListDedicatedIpPoolsFuture

	ListDeliverabilityTestReports(ctx workflow.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error)
	ListDeliverabilityTestReportsAsync(ctx workflow.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) *PinpointEmailListDeliverabilityTestReportsFuture

	ListDomainDeliverabilityCampaigns(ctx workflow.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error)
	ListDomainDeliverabilityCampaignsAsync(ctx workflow.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) *PinpointEmailListDomainDeliverabilityCampaignsFuture

	ListEmailIdentities(ctx workflow.Context, input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error)
	ListEmailIdentitiesAsync(ctx workflow.Context, input *pinpointemail.ListEmailIdentitiesInput) *PinpointEmailListEmailIdentitiesFuture

	ListTagsForResource(ctx workflow.Context, input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *pinpointemail.ListTagsForResourceInput) *PinpointEmailListTagsForResourceFuture

	PutAccountDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error)
	PutAccountDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) *PinpointEmailPutAccountDedicatedIpWarmupAttributesFuture

	PutAccountSendingAttributes(ctx workflow.Context, input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error)
	PutAccountSendingAttributesAsync(ctx workflow.Context, input *pinpointemail.PutAccountSendingAttributesInput) *PinpointEmailPutAccountSendingAttributesFuture

	PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) *PinpointEmailPutConfigurationSetDeliveryOptionsFuture

	PutConfigurationSetReputationOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error)
	PutConfigurationSetReputationOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) *PinpointEmailPutConfigurationSetReputationOptionsFuture

	PutConfigurationSetSendingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error)
	PutConfigurationSetSendingOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) *PinpointEmailPutConfigurationSetSendingOptionsFuture

	PutConfigurationSetTrackingOptions(ctx workflow.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error)
	PutConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) *PinpointEmailPutConfigurationSetTrackingOptionsFuture

	PutDedicatedIpInPool(ctx workflow.Context, input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error)
	PutDedicatedIpInPoolAsync(ctx workflow.Context, input *pinpointemail.PutDedicatedIpInPoolInput) *PinpointEmailPutDedicatedIpInPoolFuture

	PutDedicatedIpWarmupAttributes(ctx workflow.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error)
	PutDedicatedIpWarmupAttributesAsync(ctx workflow.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) *PinpointEmailPutDedicatedIpWarmupAttributesFuture

	PutDeliverabilityDashboardOption(ctx workflow.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error)
	PutDeliverabilityDashboardOptionAsync(ctx workflow.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) *PinpointEmailPutDeliverabilityDashboardOptionFuture

	PutEmailIdentityDkimAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error)
	PutEmailIdentityDkimAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) *PinpointEmailPutEmailIdentityDkimAttributesFuture

	PutEmailIdentityFeedbackAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error)
	PutEmailIdentityFeedbackAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) *PinpointEmailPutEmailIdentityFeedbackAttributesFuture

	PutEmailIdentityMailFromAttributes(ctx workflow.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error)
	PutEmailIdentityMailFromAttributesAsync(ctx workflow.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) *PinpointEmailPutEmailIdentityMailFromAttributesFuture

	SendEmail(ctx workflow.Context, input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error)
	SendEmailAsync(ctx workflow.Context, input *pinpointemail.SendEmailInput) *PinpointEmailSendEmailFuture

	TagResource(ctx workflow.Context, input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *pinpointemail.TagResourceInput) *PinpointEmailTagResourceFuture

	UntagResource(ctx workflow.Context, input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *pinpointemail.UntagResourceInput) *PinpointEmailUntagResourceFuture

	UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) *PinpointEmailUpdateConfigurationSetEventDestinationFuture
}

func NewClient() Client {
	return &stub{}
}
