// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package pinpointstub

import (
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApp(ctx workflow.Context, input *pinpoint.CreateAppInput) (*pinpoint.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *pinpoint.CreateAppInput) *PinpointCreateAppFuture

	CreateCampaign(ctx workflow.Context, input *pinpoint.CreateCampaignInput) (*pinpoint.CreateCampaignOutput, error)
	CreateCampaignAsync(ctx workflow.Context, input *pinpoint.CreateCampaignInput) *PinpointCreateCampaignFuture

	CreateEmailTemplate(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) (*pinpoint.CreateEmailTemplateOutput, error)
	CreateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.CreateEmailTemplateInput) *PinpointCreateEmailTemplateFuture

	CreateExportJob(ctx workflow.Context, input *pinpoint.CreateExportJobInput) (*pinpoint.CreateExportJobOutput, error)
	CreateExportJobAsync(ctx workflow.Context, input *pinpoint.CreateExportJobInput) *PinpointCreateExportJobFuture

	CreateImportJob(ctx workflow.Context, input *pinpoint.CreateImportJobInput) (*pinpoint.CreateImportJobOutput, error)
	CreateImportJobAsync(ctx workflow.Context, input *pinpoint.CreateImportJobInput) *PinpointCreateImportJobFuture

	CreateJourney(ctx workflow.Context, input *pinpoint.CreateJourneyInput) (*pinpoint.CreateJourneyOutput, error)
	CreateJourneyAsync(ctx workflow.Context, input *pinpoint.CreateJourneyInput) *PinpointCreateJourneyFuture

	CreatePushTemplate(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) (*pinpoint.CreatePushTemplateOutput, error)
	CreatePushTemplateAsync(ctx workflow.Context, input *pinpoint.CreatePushTemplateInput) *PinpointCreatePushTemplateFuture

	CreateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) (*pinpoint.CreateRecommenderConfigurationOutput, error)
	CreateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.CreateRecommenderConfigurationInput) *PinpointCreateRecommenderConfigurationFuture

	CreateSegment(ctx workflow.Context, input *pinpoint.CreateSegmentInput) (*pinpoint.CreateSegmentOutput, error)
	CreateSegmentAsync(ctx workflow.Context, input *pinpoint.CreateSegmentInput) *PinpointCreateSegmentFuture

	CreateSmsTemplate(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) (*pinpoint.CreateSmsTemplateOutput, error)
	CreateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.CreateSmsTemplateInput) *PinpointCreateSmsTemplateFuture

	CreateVoiceTemplate(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) (*pinpoint.CreateVoiceTemplateOutput, error)
	CreateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.CreateVoiceTemplateInput) *PinpointCreateVoiceTemplateFuture

	DeleteAdmChannel(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) (*pinpoint.DeleteAdmChannelOutput, error)
	DeleteAdmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteAdmChannelInput) *PinpointDeleteAdmChannelFuture

	DeleteApnsChannel(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) (*pinpoint.DeleteApnsChannelOutput, error)
	DeleteApnsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsChannelInput) *PinpointDeleteApnsChannelFuture

	DeleteApnsSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) (*pinpoint.DeleteApnsSandboxChannelOutput, error)
	DeleteApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsSandboxChannelInput) *PinpointDeleteApnsSandboxChannelFuture

	DeleteApnsVoipChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) (*pinpoint.DeleteApnsVoipChannelOutput, error)
	DeleteApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipChannelInput) *PinpointDeleteApnsVoipChannelFuture

	DeleteApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) (*pinpoint.DeleteApnsVoipSandboxChannelOutput, error)
	DeleteApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.DeleteApnsVoipSandboxChannelInput) *PinpointDeleteApnsVoipSandboxChannelFuture

	DeleteApp(ctx workflow.Context, input *pinpoint.DeleteAppInput) (*pinpoint.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *pinpoint.DeleteAppInput) *PinpointDeleteAppFuture

	DeleteBaiduChannel(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) (*pinpoint.DeleteBaiduChannelOutput, error)
	DeleteBaiduChannelAsync(ctx workflow.Context, input *pinpoint.DeleteBaiduChannelInput) *PinpointDeleteBaiduChannelFuture

	DeleteCampaign(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) (*pinpoint.DeleteCampaignOutput, error)
	DeleteCampaignAsync(ctx workflow.Context, input *pinpoint.DeleteCampaignInput) *PinpointDeleteCampaignFuture

	DeleteEmailChannel(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) (*pinpoint.DeleteEmailChannelOutput, error)
	DeleteEmailChannelAsync(ctx workflow.Context, input *pinpoint.DeleteEmailChannelInput) *PinpointDeleteEmailChannelFuture

	DeleteEmailTemplate(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) (*pinpoint.DeleteEmailTemplateOutput, error)
	DeleteEmailTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteEmailTemplateInput) *PinpointDeleteEmailTemplateFuture

	DeleteEndpoint(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) (*pinpoint.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *pinpoint.DeleteEndpointInput) *PinpointDeleteEndpointFuture

	DeleteEventStream(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) (*pinpoint.DeleteEventStreamOutput, error)
	DeleteEventStreamAsync(ctx workflow.Context, input *pinpoint.DeleteEventStreamInput) *PinpointDeleteEventStreamFuture

	DeleteGcmChannel(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) (*pinpoint.DeleteGcmChannelOutput, error)
	DeleteGcmChannelAsync(ctx workflow.Context, input *pinpoint.DeleteGcmChannelInput) *PinpointDeleteGcmChannelFuture

	DeleteJourney(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) (*pinpoint.DeleteJourneyOutput, error)
	DeleteJourneyAsync(ctx workflow.Context, input *pinpoint.DeleteJourneyInput) *PinpointDeleteJourneyFuture

	DeletePushTemplate(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) (*pinpoint.DeletePushTemplateOutput, error)
	DeletePushTemplateAsync(ctx workflow.Context, input *pinpoint.DeletePushTemplateInput) *PinpointDeletePushTemplateFuture

	DeleteRecommenderConfiguration(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) (*pinpoint.DeleteRecommenderConfigurationOutput, error)
	DeleteRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.DeleteRecommenderConfigurationInput) *PinpointDeleteRecommenderConfigurationFuture

	DeleteSegment(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) (*pinpoint.DeleteSegmentOutput, error)
	DeleteSegmentAsync(ctx workflow.Context, input *pinpoint.DeleteSegmentInput) *PinpointDeleteSegmentFuture

	DeleteSmsChannel(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) (*pinpoint.DeleteSmsChannelOutput, error)
	DeleteSmsChannelAsync(ctx workflow.Context, input *pinpoint.DeleteSmsChannelInput) *PinpointDeleteSmsChannelFuture

	DeleteSmsTemplate(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) (*pinpoint.DeleteSmsTemplateOutput, error)
	DeleteSmsTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteSmsTemplateInput) *PinpointDeleteSmsTemplateFuture

	DeleteUserEndpoints(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) (*pinpoint.DeleteUserEndpointsOutput, error)
	DeleteUserEndpointsAsync(ctx workflow.Context, input *pinpoint.DeleteUserEndpointsInput) *PinpointDeleteUserEndpointsFuture

	DeleteVoiceChannel(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) (*pinpoint.DeleteVoiceChannelOutput, error)
	DeleteVoiceChannelAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceChannelInput) *PinpointDeleteVoiceChannelFuture

	DeleteVoiceTemplate(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) (*pinpoint.DeleteVoiceTemplateOutput, error)
	DeleteVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.DeleteVoiceTemplateInput) *PinpointDeleteVoiceTemplateFuture

	GetAdmChannel(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error)
	GetAdmChannelAsync(ctx workflow.Context, input *pinpoint.GetAdmChannelInput) *PinpointGetAdmChannelFuture

	GetApnsChannel(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error)
	GetApnsChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsChannelInput) *PinpointGetApnsChannelFuture

	GetApnsSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error)
	GetApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsSandboxChannelInput) *PinpointGetApnsSandboxChannelFuture

	GetApnsVoipChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error)
	GetApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipChannelInput) *PinpointGetApnsVoipChannelFuture

	GetApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error)
	GetApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.GetApnsVoipSandboxChannelInput) *PinpointGetApnsVoipSandboxChannelFuture

	GetApp(ctx workflow.Context, input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error)
	GetAppAsync(ctx workflow.Context, input *pinpoint.GetAppInput) *PinpointGetAppFuture

	GetApplicationDateRangeKpi(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error)
	GetApplicationDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetApplicationDateRangeKpiInput) *PinpointGetApplicationDateRangeKpiFuture

	GetApplicationSettings(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error)
	GetApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.GetApplicationSettingsInput) *PinpointGetApplicationSettingsFuture

	GetApps(ctx workflow.Context, input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error)
	GetAppsAsync(ctx workflow.Context, input *pinpoint.GetAppsInput) *PinpointGetAppsFuture

	GetBaiduChannel(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error)
	GetBaiduChannelAsync(ctx workflow.Context, input *pinpoint.GetBaiduChannelInput) *PinpointGetBaiduChannelFuture

	GetCampaign(ctx workflow.Context, input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error)
	GetCampaignAsync(ctx workflow.Context, input *pinpoint.GetCampaignInput) *PinpointGetCampaignFuture

	GetCampaignActivities(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error)
	GetCampaignActivitiesAsync(ctx workflow.Context, input *pinpoint.GetCampaignActivitiesInput) *PinpointGetCampaignActivitiesFuture

	GetCampaignDateRangeKpi(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error)
	GetCampaignDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetCampaignDateRangeKpiInput) *PinpointGetCampaignDateRangeKpiFuture

	GetCampaignVersion(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error)
	GetCampaignVersionAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionInput) *PinpointGetCampaignVersionFuture

	GetCampaignVersions(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error)
	GetCampaignVersionsAsync(ctx workflow.Context, input *pinpoint.GetCampaignVersionsInput) *PinpointGetCampaignVersionsFuture

	GetCampaigns(ctx workflow.Context, input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error)
	GetCampaignsAsync(ctx workflow.Context, input *pinpoint.GetCampaignsInput) *PinpointGetCampaignsFuture

	GetChannels(ctx workflow.Context, input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error)
	GetChannelsAsync(ctx workflow.Context, input *pinpoint.GetChannelsInput) *PinpointGetChannelsFuture

	GetEmailChannel(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error)
	GetEmailChannelAsync(ctx workflow.Context, input *pinpoint.GetEmailChannelInput) *PinpointGetEmailChannelFuture

	GetEmailTemplate(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error)
	GetEmailTemplateAsync(ctx workflow.Context, input *pinpoint.GetEmailTemplateInput) *PinpointGetEmailTemplateFuture

	GetEndpoint(ctx workflow.Context, input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error)
	GetEndpointAsync(ctx workflow.Context, input *pinpoint.GetEndpointInput) *PinpointGetEndpointFuture

	GetEventStream(ctx workflow.Context, input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error)
	GetEventStreamAsync(ctx workflow.Context, input *pinpoint.GetEventStreamInput) *PinpointGetEventStreamFuture

	GetExportJob(ctx workflow.Context, input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error)
	GetExportJobAsync(ctx workflow.Context, input *pinpoint.GetExportJobInput) *PinpointGetExportJobFuture

	GetExportJobs(ctx workflow.Context, input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error)
	GetExportJobsAsync(ctx workflow.Context, input *pinpoint.GetExportJobsInput) *PinpointGetExportJobsFuture

	GetGcmChannel(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error)
	GetGcmChannelAsync(ctx workflow.Context, input *pinpoint.GetGcmChannelInput) *PinpointGetGcmChannelFuture

	GetImportJob(ctx workflow.Context, input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error)
	GetImportJobAsync(ctx workflow.Context, input *pinpoint.GetImportJobInput) *PinpointGetImportJobFuture

	GetImportJobs(ctx workflow.Context, input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error)
	GetImportJobsAsync(ctx workflow.Context, input *pinpoint.GetImportJobsInput) *PinpointGetImportJobsFuture

	GetJourney(ctx workflow.Context, input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error)
	GetJourneyAsync(ctx workflow.Context, input *pinpoint.GetJourneyInput) *PinpointGetJourneyFuture

	GetJourneyDateRangeKpi(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error)
	GetJourneyDateRangeKpiAsync(ctx workflow.Context, input *pinpoint.GetJourneyDateRangeKpiInput) *PinpointGetJourneyDateRangeKpiFuture

	GetJourneyExecutionActivityMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error)
	GetJourneyExecutionActivityMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput) *PinpointGetJourneyExecutionActivityMetricsFuture

	GetJourneyExecutionMetrics(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error)
	GetJourneyExecutionMetricsAsync(ctx workflow.Context, input *pinpoint.GetJourneyExecutionMetricsInput) *PinpointGetJourneyExecutionMetricsFuture

	GetPushTemplate(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error)
	GetPushTemplateAsync(ctx workflow.Context, input *pinpoint.GetPushTemplateInput) *PinpointGetPushTemplateFuture

	GetRecommenderConfiguration(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error)
	GetRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationInput) *PinpointGetRecommenderConfigurationFuture

	GetRecommenderConfigurations(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error)
	GetRecommenderConfigurationsAsync(ctx workflow.Context, input *pinpoint.GetRecommenderConfigurationsInput) *PinpointGetRecommenderConfigurationsFuture

	GetSegment(ctx workflow.Context, input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error)
	GetSegmentAsync(ctx workflow.Context, input *pinpoint.GetSegmentInput) *PinpointGetSegmentFuture

	GetSegmentExportJobs(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error)
	GetSegmentExportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentExportJobsInput) *PinpointGetSegmentExportJobsFuture

	GetSegmentImportJobs(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error)
	GetSegmentImportJobsAsync(ctx workflow.Context, input *pinpoint.GetSegmentImportJobsInput) *PinpointGetSegmentImportJobsFuture

	GetSegmentVersion(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error)
	GetSegmentVersionAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionInput) *PinpointGetSegmentVersionFuture

	GetSegmentVersions(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error)
	GetSegmentVersionsAsync(ctx workflow.Context, input *pinpoint.GetSegmentVersionsInput) *PinpointGetSegmentVersionsFuture

	GetSegments(ctx workflow.Context, input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error)
	GetSegmentsAsync(ctx workflow.Context, input *pinpoint.GetSegmentsInput) *PinpointGetSegmentsFuture

	GetSmsChannel(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error)
	GetSmsChannelAsync(ctx workflow.Context, input *pinpoint.GetSmsChannelInput) *PinpointGetSmsChannelFuture

	GetSmsTemplate(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error)
	GetSmsTemplateAsync(ctx workflow.Context, input *pinpoint.GetSmsTemplateInput) *PinpointGetSmsTemplateFuture

	GetUserEndpoints(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error)
	GetUserEndpointsAsync(ctx workflow.Context, input *pinpoint.GetUserEndpointsInput) *PinpointGetUserEndpointsFuture

	GetVoiceChannel(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error)
	GetVoiceChannelAsync(ctx workflow.Context, input *pinpoint.GetVoiceChannelInput) *PinpointGetVoiceChannelFuture

	GetVoiceTemplate(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error)
	GetVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.GetVoiceTemplateInput) *PinpointGetVoiceTemplateFuture

	ListJourneys(ctx workflow.Context, input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error)
	ListJourneysAsync(ctx workflow.Context, input *pinpoint.ListJourneysInput) *PinpointListJourneysFuture

	ListTagsForResource(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *pinpoint.ListTagsForResourceInput) *PinpointListTagsForResourceFuture

	ListTemplateVersions(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error)
	ListTemplateVersionsAsync(ctx workflow.Context, input *pinpoint.ListTemplateVersionsInput) *PinpointListTemplateVersionsFuture

	ListTemplates(ctx workflow.Context, input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error)
	ListTemplatesAsync(ctx workflow.Context, input *pinpoint.ListTemplatesInput) *PinpointListTemplatesFuture

	PhoneNumberValidate(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) (*pinpoint.PhoneNumberValidateOutput, error)
	PhoneNumberValidateAsync(ctx workflow.Context, input *pinpoint.PhoneNumberValidateInput) *PinpointPhoneNumberValidateFuture

	PutEventStream(ctx workflow.Context, input *pinpoint.PutEventStreamInput) (*pinpoint.PutEventStreamOutput, error)
	PutEventStreamAsync(ctx workflow.Context, input *pinpoint.PutEventStreamInput) *PinpointPutEventStreamFuture

	PutEvents(ctx workflow.Context, input *pinpoint.PutEventsInput) (*pinpoint.PutEventsOutput, error)
	PutEventsAsync(ctx workflow.Context, input *pinpoint.PutEventsInput) *PinpointPutEventsFuture

	RemoveAttributes(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) (*pinpoint.RemoveAttributesOutput, error)
	RemoveAttributesAsync(ctx workflow.Context, input *pinpoint.RemoveAttributesInput) *PinpointRemoveAttributesFuture

	SendMessages(ctx workflow.Context, input *pinpoint.SendMessagesInput) (*pinpoint.SendMessagesOutput, error)
	SendMessagesAsync(ctx workflow.Context, input *pinpoint.SendMessagesInput) *PinpointSendMessagesFuture

	SendUsersMessages(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) (*pinpoint.SendUsersMessagesOutput, error)
	SendUsersMessagesAsync(ctx workflow.Context, input *pinpoint.SendUsersMessagesInput) *PinpointSendUsersMessagesFuture

	TagResource(ctx workflow.Context, input *pinpoint.TagResourceInput) (*pinpoint.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *pinpoint.TagResourceInput) *PinpointTagResourceFuture

	UntagResource(ctx workflow.Context, input *pinpoint.UntagResourceInput) (*pinpoint.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *pinpoint.UntagResourceInput) *PinpointUntagResourceFuture

	UpdateAdmChannel(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) (*pinpoint.UpdateAdmChannelOutput, error)
	UpdateAdmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateAdmChannelInput) *PinpointUpdateAdmChannelFuture

	UpdateApnsChannel(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) (*pinpoint.UpdateApnsChannelOutput, error)
	UpdateApnsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsChannelInput) *PinpointUpdateApnsChannelFuture

	UpdateApnsSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) (*pinpoint.UpdateApnsSandboxChannelOutput, error)
	UpdateApnsSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsSandboxChannelInput) *PinpointUpdateApnsSandboxChannelFuture

	UpdateApnsVoipChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) (*pinpoint.UpdateApnsVoipChannelOutput, error)
	UpdateApnsVoipChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipChannelInput) *PinpointUpdateApnsVoipChannelFuture

	UpdateApnsVoipSandboxChannel(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) (*pinpoint.UpdateApnsVoipSandboxChannelOutput, error)
	UpdateApnsVoipSandboxChannelAsync(ctx workflow.Context, input *pinpoint.UpdateApnsVoipSandboxChannelInput) *PinpointUpdateApnsVoipSandboxChannelFuture

	UpdateApplicationSettings(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) (*pinpoint.UpdateApplicationSettingsOutput, error)
	UpdateApplicationSettingsAsync(ctx workflow.Context, input *pinpoint.UpdateApplicationSettingsInput) *PinpointUpdateApplicationSettingsFuture

	UpdateBaiduChannel(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) (*pinpoint.UpdateBaiduChannelOutput, error)
	UpdateBaiduChannelAsync(ctx workflow.Context, input *pinpoint.UpdateBaiduChannelInput) *PinpointUpdateBaiduChannelFuture

	UpdateCampaign(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) (*pinpoint.UpdateCampaignOutput, error)
	UpdateCampaignAsync(ctx workflow.Context, input *pinpoint.UpdateCampaignInput) *PinpointUpdateCampaignFuture

	UpdateEmailChannel(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) (*pinpoint.UpdateEmailChannelOutput, error)
	UpdateEmailChannelAsync(ctx workflow.Context, input *pinpoint.UpdateEmailChannelInput) *PinpointUpdateEmailChannelFuture

	UpdateEmailTemplate(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) (*pinpoint.UpdateEmailTemplateOutput, error)
	UpdateEmailTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateEmailTemplateInput) *PinpointUpdateEmailTemplateFuture

	UpdateEndpoint(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) (*pinpoint.UpdateEndpointOutput, error)
	UpdateEndpointAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointInput) *PinpointUpdateEndpointFuture

	UpdateEndpointsBatch(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) (*pinpoint.UpdateEndpointsBatchOutput, error)
	UpdateEndpointsBatchAsync(ctx workflow.Context, input *pinpoint.UpdateEndpointsBatchInput) *PinpointUpdateEndpointsBatchFuture

	UpdateGcmChannel(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) (*pinpoint.UpdateGcmChannelOutput, error)
	UpdateGcmChannelAsync(ctx workflow.Context, input *pinpoint.UpdateGcmChannelInput) *PinpointUpdateGcmChannelFuture

	UpdateJourney(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) (*pinpoint.UpdateJourneyOutput, error)
	UpdateJourneyAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyInput) *PinpointUpdateJourneyFuture

	UpdateJourneyState(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) (*pinpoint.UpdateJourneyStateOutput, error)
	UpdateJourneyStateAsync(ctx workflow.Context, input *pinpoint.UpdateJourneyStateInput) *PinpointUpdateJourneyStateFuture

	UpdatePushTemplate(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) (*pinpoint.UpdatePushTemplateOutput, error)
	UpdatePushTemplateAsync(ctx workflow.Context, input *pinpoint.UpdatePushTemplateInput) *PinpointUpdatePushTemplateFuture

	UpdateRecommenderConfiguration(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) (*pinpoint.UpdateRecommenderConfigurationOutput, error)
	UpdateRecommenderConfigurationAsync(ctx workflow.Context, input *pinpoint.UpdateRecommenderConfigurationInput) *PinpointUpdateRecommenderConfigurationFuture

	UpdateSegment(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) (*pinpoint.UpdateSegmentOutput, error)
	UpdateSegmentAsync(ctx workflow.Context, input *pinpoint.UpdateSegmentInput) *PinpointUpdateSegmentFuture

	UpdateSmsChannel(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) (*pinpoint.UpdateSmsChannelOutput, error)
	UpdateSmsChannelAsync(ctx workflow.Context, input *pinpoint.UpdateSmsChannelInput) *PinpointUpdateSmsChannelFuture

	UpdateSmsTemplate(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) (*pinpoint.UpdateSmsTemplateOutput, error)
	UpdateSmsTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateSmsTemplateInput) *PinpointUpdateSmsTemplateFuture

	UpdateTemplateActiveVersion(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) (*pinpoint.UpdateTemplateActiveVersionOutput, error)
	UpdateTemplateActiveVersionAsync(ctx workflow.Context, input *pinpoint.UpdateTemplateActiveVersionInput) *PinpointUpdateTemplateActiveVersionFuture

	UpdateVoiceChannel(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) (*pinpoint.UpdateVoiceChannelOutput, error)
	UpdateVoiceChannelAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceChannelInput) *PinpointUpdateVoiceChannelFuture

	UpdateVoiceTemplate(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) (*pinpoint.UpdateVoiceTemplateOutput, error)
	UpdateVoiceTemplateAsync(ctx workflow.Context, input *pinpoint.UpdateVoiceTemplateInput) *PinpointUpdateVoiceTemplateFuture
}

func NewClient() Client {
	return &stub{}
}
