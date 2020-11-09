// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package guarddutystub

import (
	"github.com/aws/aws-sdk-go/service/guardduty"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptInvitation(ctx workflow.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error)
	AcceptInvitationAsync(ctx workflow.Context, input *guardduty.AcceptInvitationInput) *AcceptInvitationFuture

	ArchiveFindings(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error)
	ArchiveFindingsAsync(ctx workflow.Context, input *guardduty.ArchiveFindingsInput) *ArchiveFindingsFuture

	CreateDetector(ctx workflow.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error)
	CreateDetectorAsync(ctx workflow.Context, input *guardduty.CreateDetectorInput) *CreateDetectorFuture

	CreateFilter(ctx workflow.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error)
	CreateFilterAsync(ctx workflow.Context, input *guardduty.CreateFilterInput) *CreateFilterFuture

	CreateIPSet(ctx workflow.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error)
	CreateIPSetAsync(ctx workflow.Context, input *guardduty.CreateIPSetInput) *CreateIPSetFuture

	CreateMembers(ctx workflow.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error)
	CreateMembersAsync(ctx workflow.Context, input *guardduty.CreateMembersInput) *CreateMembersFuture

	CreatePublishingDestination(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error)
	CreatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.CreatePublishingDestinationInput) *CreatePublishingDestinationFuture

	CreateSampleFindings(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error)
	CreateSampleFindingsAsync(ctx workflow.Context, input *guardduty.CreateSampleFindingsInput) *CreateSampleFindingsFuture

	CreateThreatIntelSet(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error)
	CreateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.CreateThreatIntelSetInput) *CreateThreatIntelSetFuture

	DeclineInvitations(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error)
	DeclineInvitationsAsync(ctx workflow.Context, input *guardduty.DeclineInvitationsInput) *DeclineInvitationsFuture

	DeleteDetector(ctx workflow.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error)
	DeleteDetectorAsync(ctx workflow.Context, input *guardduty.DeleteDetectorInput) *DeleteDetectorFuture

	DeleteFilter(ctx workflow.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error)
	DeleteFilterAsync(ctx workflow.Context, input *guardduty.DeleteFilterInput) *DeleteFilterFuture

	DeleteIPSet(ctx workflow.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error)
	DeleteIPSetAsync(ctx workflow.Context, input *guardduty.DeleteIPSetInput) *DeleteIPSetFuture

	DeleteInvitations(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error)
	DeleteInvitationsAsync(ctx workflow.Context, input *guardduty.DeleteInvitationsInput) *DeleteInvitationsFuture

	DeleteMembers(ctx workflow.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error)
	DeleteMembersAsync(ctx workflow.Context, input *guardduty.DeleteMembersInput) *DeleteMembersFuture

	DeletePublishingDestination(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error)
	DeletePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DeletePublishingDestinationInput) *DeletePublishingDestinationFuture

	DeleteThreatIntelSet(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error)
	DeleteThreatIntelSetAsync(ctx workflow.Context, input *guardduty.DeleteThreatIntelSetInput) *DeleteThreatIntelSetFuture

	DescribeOrganizationConfiguration(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.DescribeOrganizationConfigurationInput) *DescribeOrganizationConfigurationFuture

	DescribePublishingDestination(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error)
	DescribePublishingDestinationAsync(ctx workflow.Context, input *guardduty.DescribePublishingDestinationInput) *DescribePublishingDestinationFuture

	DisableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.DisableOrganizationAdminAccountInput) *DisableOrganizationAdminAccountFuture

	DisassociateFromMasterAccount(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountAsync(ctx workflow.Context, input *guardduty.DisassociateFromMasterAccountInput) *DisassociateFromMasterAccountFuture

	DisassociateMembers(ctx workflow.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error)
	DisassociateMembersAsync(ctx workflow.Context, input *guardduty.DisassociateMembersInput) *DisassociateMembersFuture

	EnableOrganizationAdminAccount(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountAsync(ctx workflow.Context, input *guardduty.EnableOrganizationAdminAccountInput) *EnableOrganizationAdminAccountFuture

	GetDetector(ctx workflow.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error)
	GetDetectorAsync(ctx workflow.Context, input *guardduty.GetDetectorInput) *GetDetectorFuture

	GetFilter(ctx workflow.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error)
	GetFilterAsync(ctx workflow.Context, input *guardduty.GetFilterInput) *GetFilterFuture

	GetFindings(ctx workflow.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error)
	GetFindingsAsync(ctx workflow.Context, input *guardduty.GetFindingsInput) *GetFindingsFuture

	GetFindingsStatistics(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error)
	GetFindingsStatisticsAsync(ctx workflow.Context, input *guardduty.GetFindingsStatisticsInput) *GetFindingsStatisticsFuture

	GetIPSet(ctx workflow.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error)
	GetIPSetAsync(ctx workflow.Context, input *guardduty.GetIPSetInput) *GetIPSetFuture

	GetInvitationsCount(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error)
	GetInvitationsCountAsync(ctx workflow.Context, input *guardduty.GetInvitationsCountInput) *GetInvitationsCountFuture

	GetMasterAccount(ctx workflow.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error)
	GetMasterAccountAsync(ctx workflow.Context, input *guardduty.GetMasterAccountInput) *GetMasterAccountFuture

	GetMemberDetectors(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error)
	GetMemberDetectorsAsync(ctx workflow.Context, input *guardduty.GetMemberDetectorsInput) *GetMemberDetectorsFuture

	GetMembers(ctx workflow.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error)
	GetMembersAsync(ctx workflow.Context, input *guardduty.GetMembersInput) *GetMembersFuture

	GetThreatIntelSet(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error)
	GetThreatIntelSetAsync(ctx workflow.Context, input *guardduty.GetThreatIntelSetInput) *GetThreatIntelSetFuture

	GetUsageStatistics(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error)
	GetUsageStatisticsAsync(ctx workflow.Context, input *guardduty.GetUsageStatisticsInput) *GetUsageStatisticsFuture

	InviteMembers(ctx workflow.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error)
	InviteMembersAsync(ctx workflow.Context, input *guardduty.InviteMembersInput) *InviteMembersFuture

	ListDetectors(ctx workflow.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error)
	ListDetectorsAsync(ctx workflow.Context, input *guardduty.ListDetectorsInput) *ListDetectorsFuture

	ListFilters(ctx workflow.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error)
	ListFiltersAsync(ctx workflow.Context, input *guardduty.ListFiltersInput) *ListFiltersFuture

	ListFindings(ctx workflow.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error)
	ListFindingsAsync(ctx workflow.Context, input *guardduty.ListFindingsInput) *ListFindingsFuture

	ListIPSets(ctx workflow.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error)
	ListIPSetsAsync(ctx workflow.Context, input *guardduty.ListIPSetsInput) *ListIPSetsFuture

	ListInvitations(ctx workflow.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *guardduty.ListInvitationsInput) *ListInvitationsFuture

	ListMembers(ctx workflow.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *guardduty.ListMembersInput) *ListMembersFuture

	ListOrganizationAdminAccounts(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsAsync(ctx workflow.Context, input *guardduty.ListOrganizationAdminAccountsInput) *ListOrganizationAdminAccountsFuture

	ListPublishingDestinations(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error)
	ListPublishingDestinationsAsync(ctx workflow.Context, input *guardduty.ListPublishingDestinationsInput) *ListPublishingDestinationsFuture

	ListTagsForResource(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *guardduty.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListThreatIntelSets(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error)
	ListThreatIntelSetsAsync(ctx workflow.Context, input *guardduty.ListThreatIntelSetsInput) *ListThreatIntelSetsFuture

	StartMonitoringMembers(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error)
	StartMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StartMonitoringMembersInput) *StartMonitoringMembersFuture

	StopMonitoringMembers(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error)
	StopMonitoringMembersAsync(ctx workflow.Context, input *guardduty.StopMonitoringMembersInput) *StopMonitoringMembersFuture

	TagResource(ctx workflow.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *guardduty.TagResourceInput) *TagResourceFuture

	UnarchiveFindings(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error)
	UnarchiveFindingsAsync(ctx workflow.Context, input *guardduty.UnarchiveFindingsInput) *UnarchiveFindingsFuture

	UntagResource(ctx workflow.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *guardduty.UntagResourceInput) *UntagResourceFuture

	UpdateDetector(ctx workflow.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error)
	UpdateDetectorAsync(ctx workflow.Context, input *guardduty.UpdateDetectorInput) *UpdateDetectorFuture

	UpdateFilter(ctx workflow.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error)
	UpdateFilterAsync(ctx workflow.Context, input *guardduty.UpdateFilterInput) *UpdateFilterFuture

	UpdateFindingsFeedback(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error)
	UpdateFindingsFeedbackAsync(ctx workflow.Context, input *guardduty.UpdateFindingsFeedbackInput) *UpdateFindingsFeedbackFuture

	UpdateIPSet(ctx workflow.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error)
	UpdateIPSetAsync(ctx workflow.Context, input *guardduty.UpdateIPSetInput) *UpdateIPSetFuture

	UpdateMemberDetectors(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error)
	UpdateMemberDetectorsAsync(ctx workflow.Context, input *guardduty.UpdateMemberDetectorsInput) *UpdateMemberDetectorsFuture

	UpdateOrganizationConfiguration(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationAsync(ctx workflow.Context, input *guardduty.UpdateOrganizationConfigurationInput) *UpdateOrganizationConfigurationFuture

	UpdatePublishingDestination(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error)
	UpdatePublishingDestinationAsync(ctx workflow.Context, input *guardduty.UpdatePublishingDestinationInput) *UpdatePublishingDestinationFuture

	UpdateThreatIntelSet(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error)
	UpdateThreatIntelSetAsync(ctx workflow.Context, input *guardduty.UpdateThreatIntelSetInput) *UpdateThreatIntelSetFuture
}

func NewClient() Client {
	return &stub{}
}
