// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appconfigstub

import (
	"github.com/aws/aws-sdk-go/service/appconfig"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateApplication(ctx workflow.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *appconfig.CreateApplicationInput) *AppConfigCreateApplicationFuture

	CreateConfigurationProfile(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error)
	CreateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) *AppConfigCreateConfigurationProfileFuture

	CreateDeploymentStrategy(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error)
	CreateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) *AppConfigCreateDeploymentStrategyFuture

	CreateEnvironment(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error)
	CreateEnvironmentAsync(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) *AppConfigCreateEnvironmentFuture

	CreateHostedConfigurationVersion(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error)
	CreateHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) *AppConfigCreateHostedConfigurationVersionFuture

	DeleteApplication(ctx workflow.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *appconfig.DeleteApplicationInput) *AppConfigDeleteApplicationFuture

	DeleteConfigurationProfile(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error)
	DeleteConfigurationProfileAsync(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) *AppConfigDeleteConfigurationProfileFuture

	DeleteDeploymentStrategy(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error)
	DeleteDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) *AppConfigDeleteDeploymentStrategyFuture

	DeleteEnvironment(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error)
	DeleteEnvironmentAsync(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) *AppConfigDeleteEnvironmentFuture

	DeleteHostedConfigurationVersion(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error)
	DeleteHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) *AppConfigDeleteHostedConfigurationVersionFuture

	GetApplication(ctx workflow.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error)
	GetApplicationAsync(ctx workflow.Context, input *appconfig.GetApplicationInput) *AppConfigGetApplicationFuture

	GetConfiguration(ctx workflow.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error)
	GetConfigurationAsync(ctx workflow.Context, input *appconfig.GetConfigurationInput) *AppConfigGetConfigurationFuture

	GetConfigurationProfile(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error)
	GetConfigurationProfileAsync(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) *AppConfigGetConfigurationProfileFuture

	GetDeployment(ctx workflow.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error)
	GetDeploymentAsync(ctx workflow.Context, input *appconfig.GetDeploymentInput) *AppConfigGetDeploymentFuture

	GetDeploymentStrategy(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error)
	GetDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) *AppConfigGetDeploymentStrategyFuture

	GetEnvironment(ctx workflow.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error)
	GetEnvironmentAsync(ctx workflow.Context, input *appconfig.GetEnvironmentInput) *AppConfigGetEnvironmentFuture

	GetHostedConfigurationVersion(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error)
	GetHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) *AppConfigGetHostedConfigurationVersionFuture

	ListApplications(ctx workflow.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *appconfig.ListApplicationsInput) *AppConfigListApplicationsFuture

	ListConfigurationProfiles(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error)
	ListConfigurationProfilesAsync(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) *AppConfigListConfigurationProfilesFuture

	ListDeploymentStrategies(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error)
	ListDeploymentStrategiesAsync(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) *AppConfigListDeploymentStrategiesFuture

	ListDeployments(ctx workflow.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error)
	ListDeploymentsAsync(ctx workflow.Context, input *appconfig.ListDeploymentsInput) *AppConfigListDeploymentsFuture

	ListEnvironments(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error)
	ListEnvironmentsAsync(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) *AppConfigListEnvironmentsFuture

	ListHostedConfigurationVersions(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error)
	ListHostedConfigurationVersionsAsync(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) *AppConfigListHostedConfigurationVersionsFuture

	ListTagsForResource(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) *AppConfigListTagsForResourceFuture

	StartDeployment(ctx workflow.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error)
	StartDeploymentAsync(ctx workflow.Context, input *appconfig.StartDeploymentInput) *AppConfigStartDeploymentFuture

	StopDeployment(ctx workflow.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error)
	StopDeploymentAsync(ctx workflow.Context, input *appconfig.StopDeploymentInput) *AppConfigStopDeploymentFuture

	TagResource(ctx workflow.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *appconfig.TagResourceInput) *AppConfigTagResourceFuture

	UntagResource(ctx workflow.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *appconfig.UntagResourceInput) *AppConfigUntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *appconfig.UpdateApplicationInput) *AppConfigUpdateApplicationFuture

	UpdateConfigurationProfile(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error)
	UpdateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) *AppConfigUpdateConfigurationProfileFuture

	UpdateDeploymentStrategy(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error)
	UpdateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) *AppConfigUpdateDeploymentStrategyFuture

	UpdateEnvironment(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error)
	UpdateEnvironmentAsync(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) *AppConfigUpdateEnvironmentFuture

	ValidateConfiguration(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error)
	ValidateConfigurationAsync(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) *AppConfigValidateConfigurationFuture
}

func NewClient() Client {
	return &stub{}
}

