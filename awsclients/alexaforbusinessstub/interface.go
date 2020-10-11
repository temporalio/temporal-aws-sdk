// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package alexaforbusinessstub

import (
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	ApproveSkill(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error)
	ApproveSkillAsync(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) *AlexaForBusinessApproveSkillFuture

	AssociateContactWithAddressBook(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateContactWithAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) *AlexaForBusinessAssociateContactWithAddressBookFuture

	AssociateDeviceWithNetworkProfile(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error)
	AssociateDeviceWithNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) *AlexaForBusinessAssociateDeviceWithNetworkProfileFuture

	AssociateDeviceWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateDeviceWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) *AlexaForBusinessAssociateDeviceWithRoomFuture

	AssociateSkillGroupWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillGroupWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) *AlexaForBusinessAssociateSkillGroupWithRoomFuture

	AssociateSkillWithSkillGroup(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) *AlexaForBusinessAssociateSkillWithSkillGroupFuture

	AssociateSkillWithUsers(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error)
	AssociateSkillWithUsersAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) *AlexaForBusinessAssociateSkillWithUsersFuture

	CreateAddressBook(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) *AlexaForBusinessCreateAddressBookFuture

	CreateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) *AlexaForBusinessCreateBusinessReportScheduleFuture

	CreateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) *AlexaForBusinessCreateConferenceProviderFuture

	CreateContact(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error)
	CreateContactAsync(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) *AlexaForBusinessCreateContactFuture

	CreateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error)
	CreateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) *AlexaForBusinessCreateGatewayGroupFuture

	CreateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error)
	CreateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) *AlexaForBusinessCreateNetworkProfileFuture

	CreateProfile(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error)
	CreateProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) *AlexaForBusinessCreateProfileFuture

	CreateRoom(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error)
	CreateRoomAsync(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) *AlexaForBusinessCreateRoomFuture

	CreateSkillGroup(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) *AlexaForBusinessCreateSkillGroupFuture

	CreateUser(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) *AlexaForBusinessCreateUserFuture

	DeleteAddressBook(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) *AlexaForBusinessDeleteAddressBookFuture

	DeleteBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) *AlexaForBusinessDeleteBusinessReportScheduleFuture

	DeleteConferenceProvider(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) *AlexaForBusinessDeleteConferenceProviderFuture

	DeleteContact(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteContactAsync(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) *AlexaForBusinessDeleteContactFuture

	DeleteDevice(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) *AlexaForBusinessDeleteDeviceFuture

	DeleteDeviceUsageData(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error)
	DeleteDeviceUsageDataAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) *AlexaForBusinessDeleteDeviceUsageDataFuture

	DeleteGatewayGroup(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error)
	DeleteGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) *AlexaForBusinessDeleteGatewayGroupFuture

	DeleteNetworkProfile(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error)
	DeleteNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) *AlexaForBusinessDeleteNetworkProfileFuture

	DeleteProfile(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) *AlexaForBusinessDeleteProfileFuture

	DeleteRoom(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) *AlexaForBusinessDeleteRoomFuture

	DeleteRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) *AlexaForBusinessDeleteRoomSkillParameterFuture

	DeleteSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) *AlexaForBusinessDeleteSkillAuthorizationFuture

	DeleteSkillGroup(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) *AlexaForBusinessDeleteSkillGroupFuture

	DeleteUser(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) *AlexaForBusinessDeleteUserFuture

	DisassociateContactFromAddressBook(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateContactFromAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) *AlexaForBusinessDisassociateContactFromAddressBookFuture

	DisassociateDeviceFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateDeviceFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) *AlexaForBusinessDisassociateDeviceFromRoomFuture

	DisassociateSkillFromSkillGroup(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) *AlexaForBusinessDisassociateSkillFromSkillGroupFuture

	DisassociateSkillFromUsers(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error)
	DisassociateSkillFromUsersAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) *AlexaForBusinessDisassociateSkillFromUsersFuture

	DisassociateSkillGroupFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	DisassociateSkillGroupFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) *AlexaForBusinessDisassociateSkillGroupFromRoomFuture

	ForgetSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	ForgetSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) *AlexaForBusinessForgetSmartHomeAppliancesFuture

	GetAddressBook(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error)
	GetAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) *AlexaForBusinessGetAddressBookFuture

	GetConferencePreference(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) *AlexaForBusinessGetConferencePreferenceFuture

	GetConferenceProvider(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) *AlexaForBusinessGetConferenceProviderFuture

	GetContact(ctx workflow.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error)
	GetContactAsync(ctx workflow.Context, input *alexaforbusiness.GetContactInput) *AlexaForBusinessGetContactFuture

	GetDevice(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) *AlexaForBusinessGetDeviceFuture

	GetGateway(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error)
	GetGatewayAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) *AlexaForBusinessGetGatewayFuture

	GetGatewayGroup(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error)
	GetGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) *AlexaForBusinessGetGatewayGroupFuture

	GetInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error)
	GetInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) *AlexaForBusinessGetInvitationConfigurationFuture

	GetNetworkProfile(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error)
	GetNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) *AlexaForBusinessGetNetworkProfileFuture

	GetProfile(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error)
	GetProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) *AlexaForBusinessGetProfileFuture

	GetRoom(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) *AlexaForBusinessGetRoomFuture

	GetRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) *AlexaForBusinessGetRoomSkillParameterFuture

	GetSkillGroup(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error)
	GetSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) *AlexaForBusinessGetSkillGroupFuture

	ListBusinessReportSchedules(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListBusinessReportSchedulesAsync(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) *AlexaForBusinessListBusinessReportSchedulesFuture

	ListConferenceProviders(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListConferenceProvidersAsync(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) *AlexaForBusinessListConferenceProvidersFuture

	ListDeviceEvents(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListDeviceEventsAsync(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) *AlexaForBusinessListDeviceEventsFuture

	ListGatewayGroups(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error)
	ListGatewayGroupsAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) *AlexaForBusinessListGatewayGroupsFuture

	ListGateways(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error)
	ListGatewaysAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) *AlexaForBusinessListGatewaysFuture

	ListSkills(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) *AlexaForBusinessListSkillsFuture

	ListSkillsStoreCategories(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreCategoriesAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) *AlexaForBusinessListSkillsStoreCategoriesFuture

	ListSkillsStoreSkillsByCategory(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSkillsStoreSkillsByCategoryAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) *AlexaForBusinessListSkillsStoreSkillsByCategoryFuture

	ListSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) *AlexaForBusinessListSmartHomeAppliancesFuture

	ListTags(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) *AlexaForBusinessListTagsFuture

	PutConferencePreference(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) *AlexaForBusinessPutConferencePreferenceFuture

	PutInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error)
	PutInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) *AlexaForBusinessPutInvitationConfigurationFuture

	PutRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) *AlexaForBusinessPutRoomSkillParameterFuture

	PutSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	PutSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) *AlexaForBusinessPutSkillAuthorizationFuture

	RegisterAVSDevice(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RegisterAVSDeviceAsync(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) *AlexaForBusinessRegisterAVSDeviceFuture

	RejectSkill(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error)
	RejectSkillAsync(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) *AlexaForBusinessRejectSkillFuture

	ResolveRoom(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error)
	ResolveRoomAsync(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) *AlexaForBusinessResolveRoomFuture

	RevokeInvitation(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error)
	RevokeInvitationAsync(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) *AlexaForBusinessRevokeInvitationFuture

	SearchAddressBooks(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchAddressBooksAsync(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) *AlexaForBusinessSearchAddressBooksFuture

	SearchContacts(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error)
	SearchContactsAsync(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) *AlexaForBusinessSearchContactsFuture

	SearchDevices(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchDevicesAsync(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) *AlexaForBusinessSearchDevicesFuture

	SearchNetworkProfiles(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error)
	SearchNetworkProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) *AlexaForBusinessSearchNetworkProfilesFuture

	SearchProfiles(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) *AlexaForBusinessSearchProfilesFuture

	SearchRooms(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchRoomsAsync(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) *AlexaForBusinessSearchRoomsFuture

	SearchSkillGroups(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchSkillGroupsAsync(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) *AlexaForBusinessSearchSkillGroupsFuture

	SearchUsers(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error)
	SearchUsersAsync(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) *AlexaForBusinessSearchUsersFuture

	SendAnnouncement(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error)
	SendAnnouncementAsync(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) *AlexaForBusinessSendAnnouncementFuture

	SendInvitation(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error)
	SendInvitationAsync(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) *AlexaForBusinessSendInvitationFuture

	StartDeviceSync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartDeviceSyncAsync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) *AlexaForBusinessStartDeviceSyncFuture

	StartSmartHomeApplianceDiscovery(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	StartSmartHomeApplianceDiscoveryAsync(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) *AlexaForBusinessStartSmartHomeApplianceDiscoveryFuture

	TagResource(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) *AlexaForBusinessTagResourceFuture

	UntagResource(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) *AlexaForBusinessUntagResourceFuture

	UpdateAddressBook(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) *AlexaForBusinessUpdateAddressBookFuture

	UpdateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) *AlexaForBusinessUpdateBusinessReportScheduleFuture

	UpdateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) *AlexaForBusinessUpdateConferenceProviderFuture

	UpdateContact(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateContactAsync(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) *AlexaForBusinessUpdateContactFuture

	UpdateDevice(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateDeviceAsync(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) *AlexaForBusinessUpdateDeviceFuture

	UpdateGateway(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error)
	UpdateGatewayAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) *AlexaForBusinessUpdateGatewayFuture

	UpdateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error)
	UpdateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) *AlexaForBusinessUpdateGatewayGroupFuture

	UpdateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error)
	UpdateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) *AlexaForBusinessUpdateNetworkProfileFuture

	UpdateProfile(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) *AlexaForBusinessUpdateProfileFuture

	UpdateRoom(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateRoomAsync(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) *AlexaForBusinessUpdateRoomFuture

	UpdateSkillGroup(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error)
	UpdateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) *AlexaForBusinessUpdateSkillGroupFuture
}

func NewClient() Client {
	return &stub{}
}

