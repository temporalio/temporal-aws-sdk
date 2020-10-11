// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CognitoIdentityProviderActivities struct {
	client cognitoidentityprovideriface.CognitoIdentityProviderAPI

	sessionFactory SessionFactory
}

func NewCognitoIdentityProviderActivities(sess *session.Session, config ...*aws.Config) *CognitoIdentityProviderActivities {
	client := cognitoidentityprovider.New(sess, config...)
	return &CognitoIdentityProviderActivities{client: client}
}

func NewCognitoIdentityProviderActivitiesWithSessionFactory(sessionFactory SessionFactory) *CognitoIdentityProviderActivities {
	return &CognitoIdentityProviderActivities{sessionFactory: sessionFactory}
}

func (a *CognitoIdentityProviderActivities) getClient(ctx context.Context) (cognitoidentityprovideriface.CognitoIdentityProviderAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cognitoidentityprovider.New(sess), nil
}

func (a *CognitoIdentityProviderActivities) AddCustomAttributes(ctx context.Context, input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddCustomAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminAddUserToGroup(ctx context.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminAddUserToGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminConfirmSignUp(ctx context.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminConfirmSignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminCreateUser(ctx context.Context, input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminCreateUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUser(ctx context.Context, input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminDeleteUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUserAttributes(ctx context.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminDeleteUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableProviderForUser(ctx context.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminDisableProviderForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableUser(ctx context.Context, input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminDisableUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminEnableUser(ctx context.Context, input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminEnableUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminForgetDevice(ctx context.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminForgetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminGetDevice(ctx context.Context, input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminGetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminGetUser(ctx context.Context, input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminGetUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminInitiateAuth(ctx context.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminInitiateAuthWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminLinkProviderForUser(ctx context.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminLinkProviderForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListDevices(ctx context.Context, input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminListDevicesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListGroupsForUser(ctx context.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminListGroupsForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListUserAuthEvents(ctx context.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminListUserAuthEventsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminRemoveUserFromGroup(ctx context.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminRemoveUserFromGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminResetUserPassword(ctx context.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminResetUserPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminRespondToAuthChallenge(ctx context.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminRespondToAuthChallengeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserMFAPreference(ctx context.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminSetUserMFAPreferenceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserPassword(ctx context.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminSetUserPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserSettings(ctx context.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminSetUserSettingsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateAuthEventFeedback(ctx context.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminUpdateAuthEventFeedbackWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateDeviceStatus(ctx context.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminUpdateDeviceStatusWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateUserAttributes(ctx context.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminUpdateUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUserGlobalSignOut(ctx context.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AdminUserGlobalSignOutWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AssociateSoftwareToken(ctx context.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateSoftwareTokenWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ChangePassword(ctx context.Context, input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ChangePasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmDevice(ctx context.Context, input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmForgotPassword(ctx context.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmForgotPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmSignUp(ctx context.Context, input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmSignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateGroup(ctx context.Context, input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateIdentityProvider(ctx context.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateResourceServer(ctx context.Context, input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserImportJob(ctx context.Context, input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPool(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolClient(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteGroup(ctx context.Context, input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteIdentityProvider(ctx context.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteResourceServer(ctx context.Context, input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUser(ctx context.Context, input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserAttributes(ctx context.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPool(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolClient(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeIdentityProvider(ctx context.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeResourceServer(ctx context.Context, input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeRiskConfiguration(ctx context.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRiskConfigurationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserImportJob(ctx context.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPool(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolClient(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ForgetDevice(ctx context.Context, input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ForgetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ForgotPassword(ctx context.Context, input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ForgotPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetCSVHeader(ctx context.Context, input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCSVHeaderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetDevice(ctx context.Context, input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetGroup(ctx context.Context, input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetIdentityProviderByIdentifier(ctx context.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIdentityProviderByIdentifierWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetSigningCertificate(ctx context.Context, input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSigningCertificateWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUICustomization(ctx context.Context, input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUICustomizationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUser(ctx context.Context, input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUserAttributeVerificationCode(ctx context.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUserAttributeVerificationCodeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUserPoolMfaConfig(ctx context.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUserPoolMfaConfigWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GlobalSignOut(ctx context.Context, input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GlobalSignOutWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) InitiateAuth(ctx context.Context, input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InitiateAuthWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListDevices(ctx context.Context, input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDevicesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListGroups(ctx context.Context, input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGroupsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListIdentityProviders(ctx context.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIdentityProvidersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListResourceServers(ctx context.Context, input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceServersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListTagsForResource(ctx context.Context, input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserImportJobs(ctx context.Context, input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUserImportJobsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserPoolClients(ctx context.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUserPoolClientsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserPools(ctx context.Context, input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUserPoolsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUsers(ctx context.Context, input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUsersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUsersInGroup(ctx context.Context, input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUsersInGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ResendConfirmationCode(ctx context.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResendConfirmationCodeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) RespondToAuthChallenge(ctx context.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RespondToAuthChallengeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetRiskConfiguration(ctx context.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetRiskConfigurationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUICustomization(ctx context.Context, input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetUICustomizationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserMFAPreference(ctx context.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetUserMFAPreferenceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserPoolMfaConfig(ctx context.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetUserPoolMfaConfigWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserSettings(ctx context.Context, input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetUserSettingsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SignUp(ctx context.Context, input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) StartUserImportJob(ctx context.Context, input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) StopUserImportJob(ctx context.Context, input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) TagResource(ctx context.Context, input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UntagResource(ctx context.Context, input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateAuthEventFeedback(ctx context.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAuthEventFeedbackWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateDeviceStatus(ctx context.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDeviceStatusWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateGroup(ctx context.Context, input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateIdentityProvider(ctx context.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateResourceServer(ctx context.Context, input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserAttributes(ctx context.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPool(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolClient(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) VerifySoftwareToken(ctx context.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.VerifySoftwareTokenWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) VerifyUserAttribute(ctx context.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.VerifyUserAttributeWithContext(ctx, input)
}
