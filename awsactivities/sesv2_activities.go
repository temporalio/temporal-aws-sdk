// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SESV2Activities struct {
	client sesv2iface.SESV2API

	sessionFactory SessionFactory
}

func NewSESV2Activities(sess *session.Session, config ...*aws.Config) *SESV2Activities {
	client := sesv2.New(sess, config...)
	return &SESV2Activities{client: client}
}

func NewSESV2ActivitiesWithSessionFactory(sessionFactory SessionFactory) *SESV2Activities {
	return &SESV2Activities{sessionFactory: sessionFactory}
}

func (a *SESV2Activities) getClient(ctx context.Context) (sesv2iface.SESV2API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return sesv2.New(sess), nil
}

func (a *SESV2Activities) CreateConfigurationSet(ctx context.Context, input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) CreateConfigurationSetEventDestination(ctx context.Context, input *sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) CreateCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) CreateDedicatedIpPool(ctx context.Context, input *sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDedicatedIpPoolWithContext(ctx, input)
}

func (a *SESV2Activities) CreateDeliverabilityTestReport(ctx context.Context, input *sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeliverabilityTestReportWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailIdentity(ctx context.Context, input *sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailIdentityPolicy(ctx context.Context, input *sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) CreateEmailTemplate(ctx context.Context, input *sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) CreateImportJob(ctx context.Context, input *sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateImportJobWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteConfigurationSet(ctx context.Context, input *sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteConfigurationSetEventDestination(ctx context.Context, input *sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteDedicatedIpPool(ctx context.Context, input *sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDedicatedIpPoolWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailIdentity(ctx context.Context, input *sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailIdentityPolicy(ctx context.Context, input *sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteEmailTemplate(ctx context.Context, input *sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) DeleteSuppressedDestination(ctx context.Context, input *sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) GetAccount(ctx context.Context, input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccountWithContext(ctx, input)
}

func (a *SESV2Activities) GetBlacklistReports(ctx context.Context, input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBlacklistReportsWithContext(ctx, input)
}

func (a *SESV2Activities) GetConfigurationSet(ctx context.Context, input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationSetWithContext(ctx, input)
}

func (a *SESV2Activities) GetConfigurationSetEventDestinations(ctx context.Context, input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationSetEventDestinationsWithContext(ctx, input)
}

func (a *SESV2Activities) GetCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) GetDedicatedIp(ctx context.Context, input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDedicatedIpWithContext(ctx, input)
}

func (a *SESV2Activities) GetDedicatedIps(ctx context.Context, input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDedicatedIpsWithContext(ctx, input)
}

func (a *SESV2Activities) GetDeliverabilityDashboardOptions(ctx context.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeliverabilityDashboardOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) GetDeliverabilityTestReport(ctx context.Context, input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeliverabilityTestReportWithContext(ctx, input)
}

func (a *SESV2Activities) GetDomainDeliverabilityCampaign(ctx context.Context, input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainDeliverabilityCampaignWithContext(ctx, input)
}

func (a *SESV2Activities) GetDomainStatisticsReport(ctx context.Context, input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainStatisticsReportWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailIdentity(ctx context.Context, input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEmailIdentityWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailIdentityPolicies(ctx context.Context, input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEmailIdentityPoliciesWithContext(ctx, input)
}

func (a *SESV2Activities) GetEmailTemplate(ctx context.Context, input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) GetImportJob(ctx context.Context, input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImportJobWithContext(ctx, input)
}

func (a *SESV2Activities) GetSuppressedDestination(ctx context.Context, input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) ListConfigurationSets(ctx context.Context, input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *SESV2Activities) ListCustomVerificationEmailTemplates(ctx context.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCustomVerificationEmailTemplatesWithContext(ctx, input)
}

func (a *SESV2Activities) ListDedicatedIpPools(ctx context.Context, input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDedicatedIpPoolsWithContext(ctx, input)
}

func (a *SESV2Activities) ListDeliverabilityTestReports(ctx context.Context, input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeliverabilityTestReportsWithContext(ctx, input)
}

func (a *SESV2Activities) ListDomainDeliverabilityCampaigns(ctx context.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDomainDeliverabilityCampaignsWithContext(ctx, input)
}

func (a *SESV2Activities) ListEmailIdentities(ctx context.Context, input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEmailIdentitiesWithContext(ctx, input)
}

func (a *SESV2Activities) ListEmailTemplates(ctx context.Context, input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEmailTemplatesWithContext(ctx, input)
}

func (a *SESV2Activities) ListImportJobs(ctx context.Context, input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImportJobsWithContext(ctx, input)
}

func (a *SESV2Activities) ListSuppressedDestinations(ctx context.Context, input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSuppressedDestinationsWithContext(ctx, input)
}

func (a *SESV2Activities) ListTagsForResource(ctx context.Context, input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, input *sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountDetails(ctx context.Context, input *sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountDetailsWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountSendingAttributes(ctx context.Context, input *sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountSendingAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutAccountSuppressionAttributes(ctx context.Context, input *sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountSuppressionAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetDeliveryOptions(ctx context.Context, input *sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetDeliveryOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetReputationOptions(ctx context.Context, input *sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetReputationOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetSendingOptions(ctx context.Context, input *sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetSendingOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetSuppressionOptions(ctx context.Context, input *sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetSuppressionOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutConfigurationSetTrackingOptions(ctx context.Context, input *sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *SESV2Activities) PutDedicatedIpInPool(ctx context.Context, input *sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDedicatedIpInPoolWithContext(ctx, input)
}

func (a *SESV2Activities) PutDedicatedIpWarmupAttributes(ctx context.Context, input *sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutDeliverabilityDashboardOption(ctx context.Context, input *sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDeliverabilityDashboardOptionWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityDkimAttributes(ctx context.Context, input *sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityDkimAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityDkimSigningAttributes(ctx context.Context, input *sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityDkimSigningAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityFeedbackAttributes(ctx context.Context, input *sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityFeedbackAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutEmailIdentityMailFromAttributes(ctx context.Context, input *sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityMailFromAttributesWithContext(ctx, input)
}

func (a *SESV2Activities) PutSuppressedDestination(ctx context.Context, input *sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutSuppressedDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) SendBulkEmail(ctx context.Context, input *sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendBulkEmailWithContext(ctx, input)
}

func (a *SESV2Activities) SendCustomVerificationEmail(ctx context.Context, input *sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendCustomVerificationEmailWithContext(ctx, input)
}

func (a *SESV2Activities) SendEmail(ctx context.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendEmailWithContext(ctx, input)
}

func (a *SESV2Activities) TagResource(ctx context.Context, input *sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *SESV2Activities) TestRenderEmailTemplate(ctx context.Context, input *sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestRenderEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) UntagResource(ctx context.Context, input *sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateConfigurationSetEventDestination(ctx context.Context, input *sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateCustomVerificationEmailTemplate(ctx context.Context, input *sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateCustomVerificationEmailTemplateWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateEmailIdentityPolicy(ctx context.Context, input *sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateEmailIdentityPolicyWithContext(ctx, input)
}

func (a *SESV2Activities) UpdateEmailTemplate(ctx context.Context, input *sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateEmailTemplateWithContext(ctx, input)
}
