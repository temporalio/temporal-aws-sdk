// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package chimestub

import (
	"github.com/aws/aws-sdk-go/service/chime"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociatePhoneNumberWithUser(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error)
	AssociatePhoneNumberWithUserAsync(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) *ChimeAssociatePhoneNumberWithUserFuture

	AssociatePhoneNumbersWithVoiceConnector(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) *ChimeAssociatePhoneNumbersWithVoiceConnectorFuture

	AssociatePhoneNumbersWithVoiceConnectorGroup(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) *ChimeAssociatePhoneNumbersWithVoiceConnectorGroupFuture

	AssociateSigninDelegateGroupsWithAccount(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error)
	AssociateSigninDelegateGroupsWithAccountAsync(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) *ChimeAssociateSigninDelegateGroupsWithAccountFuture

	BatchCreateAttendee(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error)
	BatchCreateAttendeeAsync(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) *ChimeBatchCreateAttendeeFuture

	BatchCreateRoomMembership(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error)
	BatchCreateRoomMembershipAsync(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) *ChimeBatchCreateRoomMembershipFuture

	BatchDeletePhoneNumber(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error)
	BatchDeletePhoneNumberAsync(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) *ChimeBatchDeletePhoneNumberFuture

	BatchSuspendUser(ctx workflow.Context, input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserAsync(ctx workflow.Context, input *chime.BatchSuspendUserInput) *ChimeBatchSuspendUserFuture

	BatchUnsuspendUser(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserAsync(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) *ChimeBatchUnsuspendUserFuture

	BatchUpdatePhoneNumber(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error)
	BatchUpdatePhoneNumberAsync(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) *ChimeBatchUpdatePhoneNumberFuture

	BatchUpdateUser(ctx workflow.Context, input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserAsync(ctx workflow.Context, input *chime.BatchUpdateUserInput) *ChimeBatchUpdateUserFuture

	CreateAccount(ctx workflow.Context, input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error)
	CreateAccountAsync(ctx workflow.Context, input *chime.CreateAccountInput) *ChimeCreateAccountFuture

	CreateAttendee(ctx workflow.Context, input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error)
	CreateAttendeeAsync(ctx workflow.Context, input *chime.CreateAttendeeInput) *ChimeCreateAttendeeFuture

	CreateBot(ctx workflow.Context, input *chime.CreateBotInput) (*chime.CreateBotOutput, error)
	CreateBotAsync(ctx workflow.Context, input *chime.CreateBotInput) *ChimeCreateBotFuture

	CreateMeeting(ctx workflow.Context, input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error)
	CreateMeetingAsync(ctx workflow.Context, input *chime.CreateMeetingInput) *ChimeCreateMeetingFuture

	CreateMeetingWithAttendees(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error)
	CreateMeetingWithAttendeesAsync(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) *ChimeCreateMeetingWithAttendeesFuture

	CreatePhoneNumberOrder(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error)
	CreatePhoneNumberOrderAsync(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) *ChimeCreatePhoneNumberOrderFuture

	CreateProxySession(ctx workflow.Context, input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error)
	CreateProxySessionAsync(ctx workflow.Context, input *chime.CreateProxySessionInput) *ChimeCreateProxySessionFuture

	CreateRoom(ctx workflow.Context, input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error)
	CreateRoomAsync(ctx workflow.Context, input *chime.CreateRoomInput) *ChimeCreateRoomFuture

	CreateRoomMembership(ctx workflow.Context, input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error)
	CreateRoomMembershipAsync(ctx workflow.Context, input *chime.CreateRoomMembershipInput) *ChimeCreateRoomMembershipFuture

	CreateUser(ctx workflow.Context, input *chime.CreateUserInput) (*chime.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *chime.CreateUserInput) *ChimeCreateUserFuture

	CreateVoiceConnector(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error)
	CreateVoiceConnectorAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) *ChimeCreateVoiceConnectorFuture

	CreateVoiceConnectorGroup(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error)
	CreateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) *ChimeCreateVoiceConnectorGroupFuture

	DeleteAccount(ctx workflow.Context, input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error)
	DeleteAccountAsync(ctx workflow.Context, input *chime.DeleteAccountInput) *ChimeDeleteAccountFuture

	DeleteAttendee(ctx workflow.Context, input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error)
	DeleteAttendeeAsync(ctx workflow.Context, input *chime.DeleteAttendeeInput) *ChimeDeleteAttendeeFuture

	DeleteEventsConfiguration(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error)
	DeleteEventsConfigurationAsync(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) *ChimeDeleteEventsConfigurationFuture

	DeleteMeeting(ctx workflow.Context, input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error)
	DeleteMeetingAsync(ctx workflow.Context, input *chime.DeleteMeetingInput) *ChimeDeleteMeetingFuture

	DeletePhoneNumber(ctx workflow.Context, input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error)
	DeletePhoneNumberAsync(ctx workflow.Context, input *chime.DeletePhoneNumberInput) *ChimeDeletePhoneNumberFuture

	DeleteProxySession(ctx workflow.Context, input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error)
	DeleteProxySessionAsync(ctx workflow.Context, input *chime.DeleteProxySessionInput) *ChimeDeleteProxySessionFuture

	DeleteRoom(ctx workflow.Context, input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error)
	DeleteRoomAsync(ctx workflow.Context, input *chime.DeleteRoomInput) *ChimeDeleteRoomFuture

	DeleteRoomMembership(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error)
	DeleteRoomMembershipAsync(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) *ChimeDeleteRoomMembershipFuture

	DeleteVoiceConnector(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error)
	DeleteVoiceConnectorAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) *ChimeDeleteVoiceConnectorFuture

	DeleteVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error)
	DeleteVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) *ChimeDeleteVoiceConnectorEmergencyCallingConfigurationFuture

	DeleteVoiceConnectorGroup(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error)
	DeleteVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) *ChimeDeleteVoiceConnectorGroupFuture

	DeleteVoiceConnectorOrigination(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error)
	DeleteVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) *ChimeDeleteVoiceConnectorOriginationFuture

	DeleteVoiceConnectorProxy(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error)
	DeleteVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) *ChimeDeleteVoiceConnectorProxyFuture

	DeleteVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error)
	DeleteVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) *ChimeDeleteVoiceConnectorStreamingConfigurationFuture

	DeleteVoiceConnectorTermination(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error)
	DeleteVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) *ChimeDeleteVoiceConnectorTerminationFuture

	DeleteVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
	DeleteVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) *ChimeDeleteVoiceConnectorTerminationCredentialsFuture

	DisassociatePhoneNumberFromUser(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error)
	DisassociatePhoneNumberFromUserAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) *ChimeDisassociatePhoneNumberFromUserFuture

	DisassociatePhoneNumbersFromVoiceConnector(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) *ChimeDisassociatePhoneNumbersFromVoiceConnectorFuture

	DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) *ChimeDisassociatePhoneNumbersFromVoiceConnectorGroupFuture

	DisassociateSigninDelegateGroupsFromAccount(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error)
	DisassociateSigninDelegateGroupsFromAccountAsync(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) *ChimeDisassociateSigninDelegateGroupsFromAccountFuture

	GetAccount(ctx workflow.Context, input *chime.GetAccountInput) (*chime.GetAccountOutput, error)
	GetAccountAsync(ctx workflow.Context, input *chime.GetAccountInput) *ChimeGetAccountFuture

	GetAccountSettings(ctx workflow.Context, input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsAsync(ctx workflow.Context, input *chime.GetAccountSettingsInput) *ChimeGetAccountSettingsFuture

	GetAttendee(ctx workflow.Context, input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error)
	GetAttendeeAsync(ctx workflow.Context, input *chime.GetAttendeeInput) *ChimeGetAttendeeFuture

	GetBot(ctx workflow.Context, input *chime.GetBotInput) (*chime.GetBotOutput, error)
	GetBotAsync(ctx workflow.Context, input *chime.GetBotInput) *ChimeGetBotFuture

	GetEventsConfiguration(ctx workflow.Context, input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error)
	GetEventsConfigurationAsync(ctx workflow.Context, input *chime.GetEventsConfigurationInput) *ChimeGetEventsConfigurationFuture

	GetGlobalSettings(ctx workflow.Context, input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error)
	GetGlobalSettingsAsync(ctx workflow.Context, input *chime.GetGlobalSettingsInput) *ChimeGetGlobalSettingsFuture

	GetMeeting(ctx workflow.Context, input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error)
	GetMeetingAsync(ctx workflow.Context, input *chime.GetMeetingInput) *ChimeGetMeetingFuture

	GetPhoneNumber(ctx workflow.Context, input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error)
	GetPhoneNumberAsync(ctx workflow.Context, input *chime.GetPhoneNumberInput) *ChimeGetPhoneNumberFuture

	GetPhoneNumberOrder(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error)
	GetPhoneNumberOrderAsync(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) *ChimeGetPhoneNumberOrderFuture

	GetPhoneNumberSettings(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error)
	GetPhoneNumberSettingsAsync(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) *ChimeGetPhoneNumberSettingsFuture

	GetProxySession(ctx workflow.Context, input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error)
	GetProxySessionAsync(ctx workflow.Context, input *chime.GetProxySessionInput) *ChimeGetProxySessionFuture

	GetRetentionSettings(ctx workflow.Context, input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error)
	GetRetentionSettingsAsync(ctx workflow.Context, input *chime.GetRetentionSettingsInput) *ChimeGetRetentionSettingsFuture

	GetRoom(ctx workflow.Context, input *chime.GetRoomInput) (*chime.GetRoomOutput, error)
	GetRoomAsync(ctx workflow.Context, input *chime.GetRoomInput) *ChimeGetRoomFuture

	GetUser(ctx workflow.Context, input *chime.GetUserInput) (*chime.GetUserOutput, error)
	GetUserAsync(ctx workflow.Context, input *chime.GetUserInput) *ChimeGetUserFuture

	GetUserSettings(ctx workflow.Context, input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error)
	GetUserSettingsAsync(ctx workflow.Context, input *chime.GetUserSettingsInput) *ChimeGetUserSettingsFuture

	GetVoiceConnector(ctx workflow.Context, input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error)
	GetVoiceConnectorAsync(ctx workflow.Context, input *chime.GetVoiceConnectorInput) *ChimeGetVoiceConnectorFuture

	GetVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error)
	GetVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) *ChimeGetVoiceConnectorEmergencyCallingConfigurationFuture

	GetVoiceConnectorGroup(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error)
	GetVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) *ChimeGetVoiceConnectorGroupFuture

	GetVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error)
	GetVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) *ChimeGetVoiceConnectorLoggingConfigurationFuture

	GetVoiceConnectorOrigination(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error)
	GetVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) *ChimeGetVoiceConnectorOriginationFuture

	GetVoiceConnectorProxy(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error)
	GetVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) *ChimeGetVoiceConnectorProxyFuture

	GetVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error)
	GetVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) *ChimeGetVoiceConnectorStreamingConfigurationFuture

	GetVoiceConnectorTermination(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error)
	GetVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) *ChimeGetVoiceConnectorTerminationFuture

	GetVoiceConnectorTerminationHealth(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
	GetVoiceConnectorTerminationHealthAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) *ChimeGetVoiceConnectorTerminationHealthFuture

	InviteUsers(ctx workflow.Context, input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error)
	InviteUsersAsync(ctx workflow.Context, input *chime.InviteUsersInput) *ChimeInviteUsersFuture

	ListAccounts(ctx workflow.Context, input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error)
	ListAccountsAsync(ctx workflow.Context, input *chime.ListAccountsInput) *ChimeListAccountsFuture

	ListAttendeeTags(ctx workflow.Context, input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error)
	ListAttendeeTagsAsync(ctx workflow.Context, input *chime.ListAttendeeTagsInput) *ChimeListAttendeeTagsFuture

	ListAttendees(ctx workflow.Context, input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error)
	ListAttendeesAsync(ctx workflow.Context, input *chime.ListAttendeesInput) *ChimeListAttendeesFuture

	ListBots(ctx workflow.Context, input *chime.ListBotsInput) (*chime.ListBotsOutput, error)
	ListBotsAsync(ctx workflow.Context, input *chime.ListBotsInput) *ChimeListBotsFuture

	ListMeetingTags(ctx workflow.Context, input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error)
	ListMeetingTagsAsync(ctx workflow.Context, input *chime.ListMeetingTagsInput) *ChimeListMeetingTagsFuture

	ListMeetings(ctx workflow.Context, input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error)
	ListMeetingsAsync(ctx workflow.Context, input *chime.ListMeetingsInput) *ChimeListMeetingsFuture

	ListPhoneNumberOrders(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error)
	ListPhoneNumberOrdersAsync(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) *ChimeListPhoneNumberOrdersFuture

	ListPhoneNumbers(ctx workflow.Context, input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error)
	ListPhoneNumbersAsync(ctx workflow.Context, input *chime.ListPhoneNumbersInput) *ChimeListPhoneNumbersFuture

	ListProxySessions(ctx workflow.Context, input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error)
	ListProxySessionsAsync(ctx workflow.Context, input *chime.ListProxySessionsInput) *ChimeListProxySessionsFuture

	ListRoomMemberships(ctx workflow.Context, input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error)
	ListRoomMembershipsAsync(ctx workflow.Context, input *chime.ListRoomMembershipsInput) *ChimeListRoomMembershipsFuture

	ListRooms(ctx workflow.Context, input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error)
	ListRoomsAsync(ctx workflow.Context, input *chime.ListRoomsInput) *ChimeListRoomsFuture

	ListTagsForResource(ctx workflow.Context, input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *chime.ListTagsForResourceInput) *ChimeListTagsForResourceFuture

	ListUsers(ctx workflow.Context, input *chime.ListUsersInput) (*chime.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *chime.ListUsersInput) *ChimeListUsersFuture

	ListVoiceConnectorGroups(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error)
	ListVoiceConnectorGroupsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) *ChimeListVoiceConnectorGroupsFuture

	ListVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
	ListVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) *ChimeListVoiceConnectorTerminationCredentialsFuture

	ListVoiceConnectors(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error)
	ListVoiceConnectorsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) *ChimeListVoiceConnectorsFuture

	LogoutUser(ctx workflow.Context, input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error)
	LogoutUserAsync(ctx workflow.Context, input *chime.LogoutUserInput) *ChimeLogoutUserFuture

	PutEventsConfiguration(ctx workflow.Context, input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error)
	PutEventsConfigurationAsync(ctx workflow.Context, input *chime.PutEventsConfigurationInput) *ChimePutEventsConfigurationFuture

	PutRetentionSettings(ctx workflow.Context, input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error)
	PutRetentionSettingsAsync(ctx workflow.Context, input *chime.PutRetentionSettingsInput) *ChimePutRetentionSettingsFuture

	PutVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error)
	PutVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) *ChimePutVoiceConnectorEmergencyCallingConfigurationFuture

	PutVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error)
	PutVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) *ChimePutVoiceConnectorLoggingConfigurationFuture

	PutVoiceConnectorOrigination(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error)
	PutVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) *ChimePutVoiceConnectorOriginationFuture

	PutVoiceConnectorProxy(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error)
	PutVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) *ChimePutVoiceConnectorProxyFuture

	PutVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error)
	PutVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) *ChimePutVoiceConnectorStreamingConfigurationFuture

	PutVoiceConnectorTermination(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error)
	PutVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) *ChimePutVoiceConnectorTerminationFuture

	PutVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
	PutVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) *ChimePutVoiceConnectorTerminationCredentialsFuture

	RedactConversationMessage(ctx workflow.Context, input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error)
	RedactConversationMessageAsync(ctx workflow.Context, input *chime.RedactConversationMessageInput) *ChimeRedactConversationMessageFuture

	RedactRoomMessage(ctx workflow.Context, input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error)
	RedactRoomMessageAsync(ctx workflow.Context, input *chime.RedactRoomMessageInput) *ChimeRedactRoomMessageFuture

	RegenerateSecurityToken(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error)
	RegenerateSecurityTokenAsync(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) *ChimeRegenerateSecurityTokenFuture

	ResetPersonalPIN(ctx workflow.Context, input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINAsync(ctx workflow.Context, input *chime.ResetPersonalPINInput) *ChimeResetPersonalPINFuture

	RestorePhoneNumber(ctx workflow.Context, input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error)
	RestorePhoneNumberAsync(ctx workflow.Context, input *chime.RestorePhoneNumberInput) *ChimeRestorePhoneNumberFuture

	SearchAvailablePhoneNumbers(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error)
	SearchAvailablePhoneNumbersAsync(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) *ChimeSearchAvailablePhoneNumbersFuture

	TagAttendee(ctx workflow.Context, input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error)
	TagAttendeeAsync(ctx workflow.Context, input *chime.TagAttendeeInput) *ChimeTagAttendeeFuture

	TagMeeting(ctx workflow.Context, input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error)
	TagMeetingAsync(ctx workflow.Context, input *chime.TagMeetingInput) *ChimeTagMeetingFuture

	TagResource(ctx workflow.Context, input *chime.TagResourceInput) (*chime.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *chime.TagResourceInput) *ChimeTagResourceFuture

	UntagAttendee(ctx workflow.Context, input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error)
	UntagAttendeeAsync(ctx workflow.Context, input *chime.UntagAttendeeInput) *ChimeUntagAttendeeFuture

	UntagMeeting(ctx workflow.Context, input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error)
	UntagMeetingAsync(ctx workflow.Context, input *chime.UntagMeetingInput) *ChimeUntagMeetingFuture

	UntagResource(ctx workflow.Context, input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *chime.UntagResourceInput) *ChimeUntagResourceFuture

	UpdateAccount(ctx workflow.Context, input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error)
	UpdateAccountAsync(ctx workflow.Context, input *chime.UpdateAccountInput) *ChimeUpdateAccountFuture

	UpdateAccountSettings(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsAsync(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) *ChimeUpdateAccountSettingsFuture

	UpdateBot(ctx workflow.Context, input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error)
	UpdateBotAsync(ctx workflow.Context, input *chime.UpdateBotInput) *ChimeUpdateBotFuture

	UpdateGlobalSettings(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error)
	UpdateGlobalSettingsAsync(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) *ChimeUpdateGlobalSettingsFuture

	UpdatePhoneNumber(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) *ChimeUpdatePhoneNumberFuture

	UpdatePhoneNumberSettings(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error)
	UpdatePhoneNumberSettingsAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) *ChimeUpdatePhoneNumberSettingsFuture

	UpdateProxySession(ctx workflow.Context, input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error)
	UpdateProxySessionAsync(ctx workflow.Context, input *chime.UpdateProxySessionInput) *ChimeUpdateProxySessionFuture

	UpdateRoom(ctx workflow.Context, input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error)
	UpdateRoomAsync(ctx workflow.Context, input *chime.UpdateRoomInput) *ChimeUpdateRoomFuture

	UpdateRoomMembership(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error)
	UpdateRoomMembershipAsync(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) *ChimeUpdateRoomMembershipFuture

	UpdateUser(ctx workflow.Context, input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *chime.UpdateUserInput) *ChimeUpdateUserFuture

	UpdateUserSettings(ctx workflow.Context, input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error)
	UpdateUserSettingsAsync(ctx workflow.Context, input *chime.UpdateUserSettingsInput) *ChimeUpdateUserSettingsFuture

	UpdateVoiceConnector(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error)
	UpdateVoiceConnectorAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) *ChimeUpdateVoiceConnectorFuture

	UpdateVoiceConnectorGroup(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error)
	UpdateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) *ChimeUpdateVoiceConnectorGroupFuture
}

func NewClient() Client {
	return &stub{}
}
