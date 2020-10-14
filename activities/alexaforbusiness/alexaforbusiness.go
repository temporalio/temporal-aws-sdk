// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness/alexaforbusinessiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client alexaforbusinessiface.AlexaForBusinessAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := alexaforbusiness.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (alexaforbusinessiface.AlexaForBusinessAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return alexaforbusiness.New(sess), nil
}

func (a *Activities) ApproveSkill(ctx context.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApproveSkillWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateContactWithAddressBook(ctx context.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateContactWithAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateDeviceWithNetworkProfile(ctx context.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateDeviceWithNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateDeviceWithRoom(ctx context.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateDeviceWithRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateSkillGroupWithRoom(ctx context.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateSkillGroupWithRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateSkillWithSkillGroup(ctx context.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateSkillWithSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateSkillWithUsers(ctx context.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateSkillWithUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateAddressBook(ctx context.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBusinessReportScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConferenceProvider(ctx context.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConferenceProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateContact(ctx context.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGatewayGroup(ctx context.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGatewayGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateNetworkProfile(ctx context.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProfile(ctx context.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRoom(ctx context.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSkillGroup(ctx context.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUser(ctx context.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAddressBook(ctx context.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBusinessReportScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConferenceProvider(ctx context.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConferenceProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteContact(ctx context.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDevice(ctx context.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDeviceUsageData(ctx context.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDeviceUsageDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGatewayGroup(ctx context.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGatewayGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNetworkProfile(ctx context.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProfile(ctx context.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRoom(ctx context.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRoomSkillParameter(ctx context.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRoomSkillParameterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSkillAuthorization(ctx context.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSkillAuthorizationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSkillGroup(ctx context.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUser(ctx context.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateContactFromAddressBook(ctx context.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateContactFromAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateDeviceFromRoom(ctx context.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateDeviceFromRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateSkillFromSkillGroup(ctx context.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateSkillFromSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateSkillFromUsers(ctx context.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateSkillFromUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateSkillGroupFromRoom(ctx context.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateSkillGroupFromRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ForgetSmartHomeAppliances(ctx context.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ForgetSmartHomeAppliancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAddressBook(ctx context.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConferencePreference(ctx context.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConferencePreferenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConferenceProvider(ctx context.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConferenceProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetContact(ctx context.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevice(ctx context.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGateway(ctx context.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGatewayGroup(ctx context.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGatewayGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetInvitationConfiguration(ctx context.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetInvitationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNetworkProfile(ctx context.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetProfile(ctx context.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRoom(ctx context.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRoomSkillParameter(ctx context.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRoomSkillParameterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSkillGroup(ctx context.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBusinessReportSchedules(ctx context.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBusinessReportSchedulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListConferenceProviders(ctx context.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListConferenceProvidersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDeviceEvents(ctx context.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDeviceEventsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGatewayGroups(ctx context.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGatewayGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGateways(ctx context.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGatewaysWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSkills(ctx context.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSkillsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSkillsStoreCategories(ctx context.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSkillsStoreCategoriesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSkillsStoreSkillsByCategory(ctx context.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSkillsStoreSkillsByCategoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSmartHomeAppliances(ctx context.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSmartHomeAppliancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTags(ctx context.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutConferencePreference(ctx context.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutConferencePreferenceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutInvitationConfiguration(ctx context.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutInvitationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRoomSkillParameter(ctx context.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRoomSkillParameterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutSkillAuthorization(ctx context.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutSkillAuthorizationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterAVSDevice(ctx context.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterAVSDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectSkill(ctx context.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectSkillWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResolveRoom(ctx context.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResolveRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RevokeInvitation(ctx context.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RevokeInvitationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchAddressBooks(ctx context.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchAddressBooksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchContacts(ctx context.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchContactsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchDevices(ctx context.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchDevicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchNetworkProfiles(ctx context.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchNetworkProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchProfiles(ctx context.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchRooms(ctx context.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchRoomsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchSkillGroups(ctx context.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchSkillGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchUsers(ctx context.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendAnnouncement(ctx context.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendAnnouncementWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendInvitation(ctx context.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendInvitationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartDeviceSync(ctx context.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartDeviceSyncWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartSmartHomeApplianceDiscovery(ctx context.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartSmartHomeApplianceDiscoveryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateAddressBook(ctx context.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAddressBookWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateBusinessReportScheduleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConferenceProvider(ctx context.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConferenceProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContact(ctx context.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDevice(ctx context.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGateway(ctx context.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGatewayGroup(ctx context.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGatewayGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNetworkProfile(ctx context.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNetworkProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateProfile(ctx context.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoom(ctx context.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoomWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateSkillGroup(ctx context.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateSkillGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
