// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package smsstub

import (
	"github.com/aws/aws-sdk-go/service/sms"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SMSCreateAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSCreateAppFuture) Get(ctx workflow.Context) (*sms.CreateAppOutput, error) {
	var output sms.CreateAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSCreateReplicationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSCreateReplicationJobFuture) Get(ctx workflow.Context) (*sms.CreateReplicationJobOutput, error) {
	var output sms.CreateReplicationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteAppFuture) Get(ctx workflow.Context) (*sms.DeleteAppOutput, error) {
	var output sms.DeleteAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteAppLaunchConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteAppLaunchConfigurationFuture) Get(ctx workflow.Context) (*sms.DeleteAppLaunchConfigurationOutput, error) {
	var output sms.DeleteAppLaunchConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteAppReplicationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteAppReplicationConfigurationFuture) Get(ctx workflow.Context) (*sms.DeleteAppReplicationConfigurationOutput, error) {
	var output sms.DeleteAppReplicationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteAppValidationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteAppValidationConfigurationFuture) Get(ctx workflow.Context) (*sms.DeleteAppValidationConfigurationOutput, error) {
	var output sms.DeleteAppValidationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteReplicationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteReplicationJobFuture) Get(ctx workflow.Context) (*sms.DeleteReplicationJobOutput, error) {
	var output sms.DeleteReplicationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDeleteServerCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDeleteServerCatalogFuture) Get(ctx workflow.Context) (*sms.DeleteServerCatalogOutput, error) {
	var output sms.DeleteServerCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSDisassociateConnectorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSDisassociateConnectorFuture) Get(ctx workflow.Context) (*sms.DisassociateConnectorOutput, error) {
	var output sms.DisassociateConnectorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGenerateChangeSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGenerateChangeSetFuture) Get(ctx workflow.Context) (*sms.GenerateChangeSetOutput, error) {
	var output sms.GenerateChangeSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGenerateTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGenerateTemplateFuture) Get(ctx workflow.Context) (*sms.GenerateTemplateOutput, error) {
	var output sms.GenerateTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetAppFuture) Get(ctx workflow.Context) (*sms.GetAppOutput, error) {
	var output sms.GetAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetAppLaunchConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetAppLaunchConfigurationFuture) Get(ctx workflow.Context) (*sms.GetAppLaunchConfigurationOutput, error) {
	var output sms.GetAppLaunchConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetAppReplicationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetAppReplicationConfigurationFuture) Get(ctx workflow.Context) (*sms.GetAppReplicationConfigurationOutput, error) {
	var output sms.GetAppReplicationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetAppValidationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetAppValidationConfigurationFuture) Get(ctx workflow.Context) (*sms.GetAppValidationConfigurationOutput, error) {
	var output sms.GetAppValidationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetAppValidationOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetAppValidationOutputFuture) Get(ctx workflow.Context) (*sms.GetAppValidationOutputOutput, error) {
	var output sms.GetAppValidationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetConnectorsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetConnectorsFuture) Get(ctx workflow.Context) (*sms.GetConnectorsOutput, error) {
	var output sms.GetConnectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetReplicationJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetReplicationJobsFuture) Get(ctx workflow.Context) (*sms.GetReplicationJobsOutput, error) {
	var output sms.GetReplicationJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetReplicationRunsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetReplicationRunsFuture) Get(ctx workflow.Context) (*sms.GetReplicationRunsOutput, error) {
	var output sms.GetReplicationRunsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSGetServersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSGetServersFuture) Get(ctx workflow.Context) (*sms.GetServersOutput, error) {
	var output sms.GetServersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSImportAppCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSImportAppCatalogFuture) Get(ctx workflow.Context) (*sms.ImportAppCatalogOutput, error) {
	var output sms.ImportAppCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSImportServerCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSImportServerCatalogFuture) Get(ctx workflow.Context) (*sms.ImportServerCatalogOutput, error) {
	var output sms.ImportServerCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSLaunchAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSLaunchAppFuture) Get(ctx workflow.Context) (*sms.LaunchAppOutput, error) {
	var output sms.LaunchAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSListAppsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSListAppsFuture) Get(ctx workflow.Context) (*sms.ListAppsOutput, error) {
	var output sms.ListAppsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSNotifyAppValidationOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSNotifyAppValidationOutputFuture) Get(ctx workflow.Context) (*sms.NotifyAppValidationOutputOutput, error) {
	var output sms.NotifyAppValidationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSPutAppLaunchConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSPutAppLaunchConfigurationFuture) Get(ctx workflow.Context) (*sms.PutAppLaunchConfigurationOutput, error) {
	var output sms.PutAppLaunchConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSPutAppReplicationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSPutAppReplicationConfigurationFuture) Get(ctx workflow.Context) (*sms.PutAppReplicationConfigurationOutput, error) {
	var output sms.PutAppReplicationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSPutAppValidationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSPutAppValidationConfigurationFuture) Get(ctx workflow.Context) (*sms.PutAppValidationConfigurationOutput, error) {
	var output sms.PutAppValidationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSStartAppReplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSStartAppReplicationFuture) Get(ctx workflow.Context) (*sms.StartAppReplicationOutput, error) {
	var output sms.StartAppReplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSStartOnDemandAppReplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSStartOnDemandAppReplicationFuture) Get(ctx workflow.Context) (*sms.StartOnDemandAppReplicationOutput, error) {
	var output sms.StartOnDemandAppReplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSStartOnDemandReplicationRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSStartOnDemandReplicationRunFuture) Get(ctx workflow.Context) (*sms.StartOnDemandReplicationRunOutput, error) {
	var output sms.StartOnDemandReplicationRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSStopAppReplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSStopAppReplicationFuture) Get(ctx workflow.Context) (*sms.StopAppReplicationOutput, error) {
	var output sms.StopAppReplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSTerminateAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSTerminateAppFuture) Get(ctx workflow.Context) (*sms.TerminateAppOutput, error) {
	var output sms.TerminateAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSUpdateAppFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSUpdateAppFuture) Get(ctx workflow.Context) (*sms.UpdateAppOutput, error) {
	var output sms.UpdateAppOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SMSUpdateReplicationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SMSUpdateReplicationJobFuture) Get(ctx workflow.Context) (*sms.UpdateReplicationJobOutput, error) {
	var output sms.UpdateReplicationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApp(ctx workflow.Context, input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
	var output sms.CreateAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.CreateApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAppAsync(ctx workflow.Context, input *sms.CreateAppInput) *SMSCreateAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.CreateApp", input)
	return &SMSCreateAppFuture{Future: future}
}

func (a *stub) CreateReplicationJob(ctx workflow.Context, input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
	var output sms.CreateReplicationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.CreateReplicationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateReplicationJobAsync(ctx workflow.Context, input *sms.CreateReplicationJobInput) *SMSCreateReplicationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.CreateReplicationJob", input)
	return &SMSCreateReplicationJobFuture{Future: future}
}

func (a *stub) DeleteApp(ctx workflow.Context, input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error) {
	var output sms.DeleteAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppAsync(ctx workflow.Context, input *sms.DeleteAppInput) *SMSDeleteAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteApp", input)
	return &SMSDeleteAppFuture{Future: future}
}

func (a *stub) DeleteAppLaunchConfiguration(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error) {
	var output sms.DeleteAppLaunchConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppLaunchConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppLaunchConfigurationInput) *SMSDeleteAppLaunchConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppLaunchConfiguration", input)
	return &SMSDeleteAppLaunchConfigurationFuture{Future: future}
}

func (a *stub) DeleteAppReplicationConfiguration(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error) {
	var output sms.DeleteAppReplicationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppReplicationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppReplicationConfigurationInput) *SMSDeleteAppReplicationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppReplicationConfiguration", input)
	return &SMSDeleteAppReplicationConfigurationFuture{Future: future}
}

func (a *stub) DeleteAppValidationConfiguration(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error) {
	var output sms.DeleteAppValidationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppValidationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppValidationConfigurationAsync(ctx workflow.Context, input *sms.DeleteAppValidationConfigurationInput) *SMSDeleteAppValidationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteAppValidationConfiguration", input)
	return &SMSDeleteAppValidationConfigurationFuture{Future: future}
}

func (a *stub) DeleteReplicationJob(ctx workflow.Context, input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error) {
	var output sms.DeleteReplicationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteReplicationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteReplicationJobAsync(ctx workflow.Context, input *sms.DeleteReplicationJobInput) *SMSDeleteReplicationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteReplicationJob", input)
	return &SMSDeleteReplicationJobFuture{Future: future}
}

func (a *stub) DeleteServerCatalog(ctx workflow.Context, input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error) {
	var output sms.DeleteServerCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DeleteServerCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteServerCatalogAsync(ctx workflow.Context, input *sms.DeleteServerCatalogInput) *SMSDeleteServerCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DeleteServerCatalog", input)
	return &SMSDeleteServerCatalogFuture{Future: future}
}

func (a *stub) DisassociateConnector(ctx workflow.Context, input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error) {
	var output sms.DisassociateConnectorOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.DisassociateConnector", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateConnectorAsync(ctx workflow.Context, input *sms.DisassociateConnectorInput) *SMSDisassociateConnectorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.DisassociateConnector", input)
	return &SMSDisassociateConnectorFuture{Future: future}
}

func (a *stub) GenerateChangeSet(ctx workflow.Context, input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error) {
	var output sms.GenerateChangeSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GenerateChangeSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateChangeSetAsync(ctx workflow.Context, input *sms.GenerateChangeSetInput) *SMSGenerateChangeSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GenerateChangeSet", input)
	return &SMSGenerateChangeSetFuture{Future: future}
}

func (a *stub) GenerateTemplate(ctx workflow.Context, input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error) {
	var output sms.GenerateTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GenerateTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateTemplateAsync(ctx workflow.Context, input *sms.GenerateTemplateInput) *SMSGenerateTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GenerateTemplate", input)
	return &SMSGenerateTemplateFuture{Future: future}
}

func (a *stub) GetApp(ctx workflow.Context, input *sms.GetAppInput) (*sms.GetAppOutput, error) {
	var output sms.GetAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppAsync(ctx workflow.Context, input *sms.GetAppInput) *SMSGetAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetApp", input)
	return &SMSGetAppFuture{Future: future}
}

func (a *stub) GetAppLaunchConfiguration(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error) {
	var output sms.GetAppLaunchConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetAppLaunchConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.GetAppLaunchConfigurationInput) *SMSGetAppLaunchConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetAppLaunchConfiguration", input)
	return &SMSGetAppLaunchConfigurationFuture{Future: future}
}

func (a *stub) GetAppReplicationConfiguration(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error) {
	var output sms.GetAppReplicationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetAppReplicationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.GetAppReplicationConfigurationInput) *SMSGetAppReplicationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetAppReplicationConfiguration", input)
	return &SMSGetAppReplicationConfigurationFuture{Future: future}
}

func (a *stub) GetAppValidationConfiguration(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error) {
	var output sms.GetAppValidationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetAppValidationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppValidationConfigurationAsync(ctx workflow.Context, input *sms.GetAppValidationConfigurationInput) *SMSGetAppValidationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetAppValidationConfiguration", input)
	return &SMSGetAppValidationConfigurationFuture{Future: future}
}

func (a *stub) GetAppValidationOutput(ctx workflow.Context, input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error) {
	var output sms.GetAppValidationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetAppValidationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppValidationOutputAsync(ctx workflow.Context, input *sms.GetAppValidationOutputInput) *SMSGetAppValidationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetAppValidationOutput", input)
	return &SMSGetAppValidationOutputFuture{Future: future}
}

func (a *stub) GetConnectors(ctx workflow.Context, input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error) {
	var output sms.GetConnectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetConnectors", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConnectorsAsync(ctx workflow.Context, input *sms.GetConnectorsInput) *SMSGetConnectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetConnectors", input)
	return &SMSGetConnectorsFuture{Future: future}
}

func (a *stub) GetReplicationJobs(ctx workflow.Context, input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error) {
	var output sms.GetReplicationJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetReplicationJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetReplicationJobsAsync(ctx workflow.Context, input *sms.GetReplicationJobsInput) *SMSGetReplicationJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetReplicationJobs", input)
	return &SMSGetReplicationJobsFuture{Future: future}
}

func (a *stub) GetReplicationRuns(ctx workflow.Context, input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error) {
	var output sms.GetReplicationRunsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetReplicationRuns", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetReplicationRunsAsync(ctx workflow.Context, input *sms.GetReplicationRunsInput) *SMSGetReplicationRunsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetReplicationRuns", input)
	return &SMSGetReplicationRunsFuture{Future: future}
}

func (a *stub) GetServers(ctx workflow.Context, input *sms.GetServersInput) (*sms.GetServersOutput, error) {
	var output sms.GetServersOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.GetServers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetServersAsync(ctx workflow.Context, input *sms.GetServersInput) *SMSGetServersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.GetServers", input)
	return &SMSGetServersFuture{Future: future}
}

func (a *stub) ImportAppCatalog(ctx workflow.Context, input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error) {
	var output sms.ImportAppCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.ImportAppCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportAppCatalogAsync(ctx workflow.Context, input *sms.ImportAppCatalogInput) *SMSImportAppCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.ImportAppCatalog", input)
	return &SMSImportAppCatalogFuture{Future: future}
}

func (a *stub) ImportServerCatalog(ctx workflow.Context, input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error) {
	var output sms.ImportServerCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.ImportServerCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportServerCatalogAsync(ctx workflow.Context, input *sms.ImportServerCatalogInput) *SMSImportServerCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.ImportServerCatalog", input)
	return &SMSImportServerCatalogFuture{Future: future}
}

func (a *stub) LaunchApp(ctx workflow.Context, input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error) {
	var output sms.LaunchAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.LaunchApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) LaunchAppAsync(ctx workflow.Context, input *sms.LaunchAppInput) *SMSLaunchAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.LaunchApp", input)
	return &SMSLaunchAppFuture{Future: future}
}

func (a *stub) ListApps(ctx workflow.Context, input *sms.ListAppsInput) (*sms.ListAppsOutput, error) {
	var output sms.ListAppsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.ListApps", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAppsAsync(ctx workflow.Context, input *sms.ListAppsInput) *SMSListAppsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.ListApps", input)
	return &SMSListAppsFuture{Future: future}
}

func (a *stub) NotifyAppValidationOutput(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error) {
	var output sms.NotifyAppValidationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.NotifyAppValidationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) NotifyAppValidationOutputAsync(ctx workflow.Context, input *sms.NotifyAppValidationOutputInput) *SMSNotifyAppValidationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.NotifyAppValidationOutput", input)
	return &SMSNotifyAppValidationOutputFuture{Future: future}
}

func (a *stub) PutAppLaunchConfiguration(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error) {
	var output sms.PutAppLaunchConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.PutAppLaunchConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAppLaunchConfigurationAsync(ctx workflow.Context, input *sms.PutAppLaunchConfigurationInput) *SMSPutAppLaunchConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.PutAppLaunchConfiguration", input)
	return &SMSPutAppLaunchConfigurationFuture{Future: future}
}

func (a *stub) PutAppReplicationConfiguration(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error) {
	var output sms.PutAppReplicationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.PutAppReplicationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAppReplicationConfigurationAsync(ctx workflow.Context, input *sms.PutAppReplicationConfigurationInput) *SMSPutAppReplicationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.PutAppReplicationConfiguration", input)
	return &SMSPutAppReplicationConfigurationFuture{Future: future}
}

func (a *stub) PutAppValidationConfiguration(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error) {
	var output sms.PutAppValidationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.PutAppValidationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAppValidationConfigurationAsync(ctx workflow.Context, input *sms.PutAppValidationConfigurationInput) *SMSPutAppValidationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.PutAppValidationConfiguration", input)
	return &SMSPutAppValidationConfigurationFuture{Future: future}
}

func (a *stub) StartAppReplication(ctx workflow.Context, input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error) {
	var output sms.StartAppReplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.StartAppReplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartAppReplicationAsync(ctx workflow.Context, input *sms.StartAppReplicationInput) *SMSStartAppReplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.StartAppReplication", input)
	return &SMSStartAppReplicationFuture{Future: future}
}

func (a *stub) StartOnDemandAppReplication(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error) {
	var output sms.StartOnDemandAppReplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.StartOnDemandAppReplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartOnDemandAppReplicationAsync(ctx workflow.Context, input *sms.StartOnDemandAppReplicationInput) *SMSStartOnDemandAppReplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.StartOnDemandAppReplication", input)
	return &SMSStartOnDemandAppReplicationFuture{Future: future}
}

func (a *stub) StartOnDemandReplicationRun(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error) {
	var output sms.StartOnDemandReplicationRunOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.StartOnDemandReplicationRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartOnDemandReplicationRunAsync(ctx workflow.Context, input *sms.StartOnDemandReplicationRunInput) *SMSStartOnDemandReplicationRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.StartOnDemandReplicationRun", input)
	return &SMSStartOnDemandReplicationRunFuture{Future: future}
}

func (a *stub) StopAppReplication(ctx workflow.Context, input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error) {
	var output sms.StopAppReplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.StopAppReplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopAppReplicationAsync(ctx workflow.Context, input *sms.StopAppReplicationInput) *SMSStopAppReplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.StopAppReplication", input)
	return &SMSStopAppReplicationFuture{Future: future}
}

func (a *stub) TerminateApp(ctx workflow.Context, input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error) {
	var output sms.TerminateAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.TerminateApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TerminateAppAsync(ctx workflow.Context, input *sms.TerminateAppInput) *SMSTerminateAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.TerminateApp", input)
	return &SMSTerminateAppFuture{Future: future}
}

func (a *stub) UpdateApp(ctx workflow.Context, input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error) {
	var output sms.UpdateAppOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.UpdateApp", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAppAsync(ctx workflow.Context, input *sms.UpdateAppInput) *SMSUpdateAppFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.UpdateApp", input)
	return &SMSUpdateAppFuture{Future: future}
}

func (a *stub) UpdateReplicationJob(ctx workflow.Context, input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error) {
	var output sms.UpdateReplicationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.sms.UpdateReplicationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateReplicationJobAsync(ctx workflow.Context, input *sms.UpdateReplicationJobInput) *SMSUpdateReplicationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sms.UpdateReplicationJob", input)
	return &SMSUpdateReplicationJobFuture{Future: future}
}
