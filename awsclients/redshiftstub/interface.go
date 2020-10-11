// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package redshiftstub

import (
	"github.com/aws/aws-sdk-go/service/redshift"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AcceptReservedNodeExchange(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error)
	AcceptReservedNodeExchangeAsync(ctx workflow.Context, input *redshift.AcceptReservedNodeExchangeInput) *RedshiftAcceptReservedNodeExchangeFuture

	AuthorizeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) (*redshift.AuthorizeClusterSecurityGroupIngressOutput, error)
	AuthorizeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.AuthorizeClusterSecurityGroupIngressInput) *RedshiftAuthorizeClusterSecurityGroupIngressFuture

	AuthorizeSnapshotAccess(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) (*redshift.AuthorizeSnapshotAccessOutput, error)
	AuthorizeSnapshotAccessAsync(ctx workflow.Context, input *redshift.AuthorizeSnapshotAccessInput) *RedshiftAuthorizeSnapshotAccessFuture

	BatchDeleteClusterSnapshots(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) (*redshift.BatchDeleteClusterSnapshotsOutput, error)
	BatchDeleteClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchDeleteClusterSnapshotsInput) *RedshiftBatchDeleteClusterSnapshotsFuture

	BatchModifyClusterSnapshots(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) (*redshift.BatchModifyClusterSnapshotsOutput, error)
	BatchModifyClusterSnapshotsAsync(ctx workflow.Context, input *redshift.BatchModifyClusterSnapshotsInput) *RedshiftBatchModifyClusterSnapshotsFuture

	CancelResize(ctx workflow.Context, input *redshift.CancelResizeInput) (*redshift.CancelResizeOutput, error)
	CancelResizeAsync(ctx workflow.Context, input *redshift.CancelResizeInput) *RedshiftCancelResizeFuture

	CopyClusterSnapshot(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) (*redshift.CopyClusterSnapshotOutput, error)
	CopyClusterSnapshotAsync(ctx workflow.Context, input *redshift.CopyClusterSnapshotInput) *RedshiftCopyClusterSnapshotFuture

	CreateCluster(ctx workflow.Context, input *redshift.CreateClusterInput) (*redshift.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *redshift.CreateClusterInput) *RedshiftCreateClusterFuture

	CreateClusterParameterGroup(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) (*redshift.CreateClusterParameterGroupOutput, error)
	CreateClusterParameterGroupAsync(ctx workflow.Context, input *redshift.CreateClusterParameterGroupInput) *RedshiftCreateClusterParameterGroupFuture

	CreateClusterSecurityGroup(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) (*redshift.CreateClusterSecurityGroupOutput, error)
	CreateClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSecurityGroupInput) *RedshiftCreateClusterSecurityGroupFuture

	CreateClusterSnapshot(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) (*redshift.CreateClusterSnapshotOutput, error)
	CreateClusterSnapshotAsync(ctx workflow.Context, input *redshift.CreateClusterSnapshotInput) *RedshiftCreateClusterSnapshotFuture

	CreateClusterSubnetGroup(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) (*redshift.CreateClusterSubnetGroupOutput, error)
	CreateClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.CreateClusterSubnetGroupInput) *RedshiftCreateClusterSubnetGroupFuture

	CreateEventSubscription(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) (*redshift.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionAsync(ctx workflow.Context, input *redshift.CreateEventSubscriptionInput) *RedshiftCreateEventSubscriptionFuture

	CreateHsmClientCertificate(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) (*redshift.CreateHsmClientCertificateOutput, error)
	CreateHsmClientCertificateAsync(ctx workflow.Context, input *redshift.CreateHsmClientCertificateInput) *RedshiftCreateHsmClientCertificateFuture

	CreateHsmConfiguration(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) (*redshift.CreateHsmConfigurationOutput, error)
	CreateHsmConfigurationAsync(ctx workflow.Context, input *redshift.CreateHsmConfigurationInput) *RedshiftCreateHsmConfigurationFuture

	CreateScheduledAction(ctx workflow.Context, input *redshift.CreateScheduledActionInput) (*redshift.CreateScheduledActionOutput, error)
	CreateScheduledActionAsync(ctx workflow.Context, input *redshift.CreateScheduledActionInput) *RedshiftCreateScheduledActionFuture

	CreateSnapshotCopyGrant(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) (*redshift.CreateSnapshotCopyGrantOutput, error)
	CreateSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.CreateSnapshotCopyGrantInput) *RedshiftCreateSnapshotCopyGrantFuture

	CreateSnapshotSchedule(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) (*redshift.CreateSnapshotScheduleOutput, error)
	CreateSnapshotScheduleAsync(ctx workflow.Context, input *redshift.CreateSnapshotScheduleInput) *RedshiftCreateSnapshotScheduleFuture

	CreateTags(ctx workflow.Context, input *redshift.CreateTagsInput) (*redshift.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *redshift.CreateTagsInput) *RedshiftCreateTagsFuture

	CreateUsageLimit(ctx workflow.Context, input *redshift.CreateUsageLimitInput) (*redshift.CreateUsageLimitOutput, error)
	CreateUsageLimitAsync(ctx workflow.Context, input *redshift.CreateUsageLimitInput) *RedshiftCreateUsageLimitFuture

	DeleteCluster(ctx workflow.Context, input *redshift.DeleteClusterInput) (*redshift.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *redshift.DeleteClusterInput) *RedshiftDeleteClusterFuture

	DeleteClusterParameterGroup(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) (*redshift.DeleteClusterParameterGroupOutput, error)
	DeleteClusterParameterGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterParameterGroupInput) *RedshiftDeleteClusterParameterGroupFuture

	DeleteClusterSecurityGroup(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) (*redshift.DeleteClusterSecurityGroupOutput, error)
	DeleteClusterSecurityGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSecurityGroupInput) *RedshiftDeleteClusterSecurityGroupFuture

	DeleteClusterSnapshot(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) (*redshift.DeleteClusterSnapshotOutput, error)
	DeleteClusterSnapshotAsync(ctx workflow.Context, input *redshift.DeleteClusterSnapshotInput) *RedshiftDeleteClusterSnapshotFuture

	DeleteClusterSubnetGroup(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) (*redshift.DeleteClusterSubnetGroupOutput, error)
	DeleteClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.DeleteClusterSubnetGroupInput) *RedshiftDeleteClusterSubnetGroupFuture

	DeleteEventSubscription(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) (*redshift.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionAsync(ctx workflow.Context, input *redshift.DeleteEventSubscriptionInput) *RedshiftDeleteEventSubscriptionFuture

	DeleteHsmClientCertificate(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) (*redshift.DeleteHsmClientCertificateOutput, error)
	DeleteHsmClientCertificateAsync(ctx workflow.Context, input *redshift.DeleteHsmClientCertificateInput) *RedshiftDeleteHsmClientCertificateFuture

	DeleteHsmConfiguration(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) (*redshift.DeleteHsmConfigurationOutput, error)
	DeleteHsmConfigurationAsync(ctx workflow.Context, input *redshift.DeleteHsmConfigurationInput) *RedshiftDeleteHsmConfigurationFuture

	DeleteScheduledAction(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) (*redshift.DeleteScheduledActionOutput, error)
	DeleteScheduledActionAsync(ctx workflow.Context, input *redshift.DeleteScheduledActionInput) *RedshiftDeleteScheduledActionFuture

	DeleteSnapshotCopyGrant(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) (*redshift.DeleteSnapshotCopyGrantOutput, error)
	DeleteSnapshotCopyGrantAsync(ctx workflow.Context, input *redshift.DeleteSnapshotCopyGrantInput) *RedshiftDeleteSnapshotCopyGrantFuture

	DeleteSnapshotSchedule(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) (*redshift.DeleteSnapshotScheduleOutput, error)
	DeleteSnapshotScheduleAsync(ctx workflow.Context, input *redshift.DeleteSnapshotScheduleInput) *RedshiftDeleteSnapshotScheduleFuture

	DeleteTags(ctx workflow.Context, input *redshift.DeleteTagsInput) (*redshift.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *redshift.DeleteTagsInput) *RedshiftDeleteTagsFuture

	DeleteUsageLimit(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) (*redshift.DeleteUsageLimitOutput, error)
	DeleteUsageLimitAsync(ctx workflow.Context, input *redshift.DeleteUsageLimitInput) *RedshiftDeleteUsageLimitFuture

	DescribeAccountAttributes(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) (*redshift.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *redshift.DescribeAccountAttributesInput) *RedshiftDescribeAccountAttributesFuture

	DescribeClusterDbRevisions(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) (*redshift.DescribeClusterDbRevisionsOutput, error)
	DescribeClusterDbRevisionsAsync(ctx workflow.Context, input *redshift.DescribeClusterDbRevisionsInput) *RedshiftDescribeClusterDbRevisionsFuture

	DescribeClusterParameterGroups(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) (*redshift.DescribeClusterParameterGroupsOutput, error)
	DescribeClusterParameterGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterParameterGroupsInput) *RedshiftDescribeClusterParameterGroupsFuture

	DescribeClusterParameters(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) (*redshift.DescribeClusterParametersOutput, error)
	DescribeClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeClusterParametersInput) *RedshiftDescribeClusterParametersFuture

	DescribeClusterSecurityGroups(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) (*redshift.DescribeClusterSecurityGroupsOutput, error)
	DescribeClusterSecurityGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSecurityGroupsInput) *RedshiftDescribeClusterSecurityGroupsFuture

	DescribeClusterSnapshots(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) (*redshift.DescribeClusterSnapshotsOutput, error)
	DescribeClusterSnapshotsAsync(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) *RedshiftDescribeClusterSnapshotsFuture

	DescribeClusterSubnetGroups(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) (*redshift.DescribeClusterSubnetGroupsOutput, error)
	DescribeClusterSubnetGroupsAsync(ctx workflow.Context, input *redshift.DescribeClusterSubnetGroupsInput) *RedshiftDescribeClusterSubnetGroupsFuture

	DescribeClusterTracks(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) (*redshift.DescribeClusterTracksOutput, error)
	DescribeClusterTracksAsync(ctx workflow.Context, input *redshift.DescribeClusterTracksInput) *RedshiftDescribeClusterTracksFuture

	DescribeClusterVersions(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) (*redshift.DescribeClusterVersionsOutput, error)
	DescribeClusterVersionsAsync(ctx workflow.Context, input *redshift.DescribeClusterVersionsInput) *RedshiftDescribeClusterVersionsFuture

	DescribeClusters(ctx workflow.Context, input *redshift.DescribeClustersInput) (*redshift.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *RedshiftDescribeClustersFuture

	DescribeDefaultClusterParameters(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) (*redshift.DescribeDefaultClusterParametersOutput, error)
	DescribeDefaultClusterParametersAsync(ctx workflow.Context, input *redshift.DescribeDefaultClusterParametersInput) *RedshiftDescribeDefaultClusterParametersFuture

	DescribeEventCategories(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) (*redshift.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesAsync(ctx workflow.Context, input *redshift.DescribeEventCategoriesInput) *RedshiftDescribeEventCategoriesFuture

	DescribeEventSubscriptions(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) (*redshift.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsAsync(ctx workflow.Context, input *redshift.DescribeEventSubscriptionsInput) *RedshiftDescribeEventSubscriptionsFuture

	DescribeEvents(ctx workflow.Context, input *redshift.DescribeEventsInput) (*redshift.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *redshift.DescribeEventsInput) *RedshiftDescribeEventsFuture

	DescribeHsmClientCertificates(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) (*redshift.DescribeHsmClientCertificatesOutput, error)
	DescribeHsmClientCertificatesAsync(ctx workflow.Context, input *redshift.DescribeHsmClientCertificatesInput) *RedshiftDescribeHsmClientCertificatesFuture

	DescribeHsmConfigurations(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) (*redshift.DescribeHsmConfigurationsOutput, error)
	DescribeHsmConfigurationsAsync(ctx workflow.Context, input *redshift.DescribeHsmConfigurationsInput) *RedshiftDescribeHsmConfigurationsFuture

	DescribeLoggingStatus(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) (*redshift.LoggingStatus, error)
	DescribeLoggingStatusAsync(ctx workflow.Context, input *redshift.DescribeLoggingStatusInput) *RedshiftDescribeLoggingStatusFuture

	DescribeNodeConfigurationOptions(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) (*redshift.DescribeNodeConfigurationOptionsOutput, error)
	DescribeNodeConfigurationOptionsAsync(ctx workflow.Context, input *redshift.DescribeNodeConfigurationOptionsInput) *RedshiftDescribeNodeConfigurationOptionsFuture

	DescribeOrderableClusterOptions(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) (*redshift.DescribeOrderableClusterOptionsOutput, error)
	DescribeOrderableClusterOptionsAsync(ctx workflow.Context, input *redshift.DescribeOrderableClusterOptionsInput) *RedshiftDescribeOrderableClusterOptionsFuture

	DescribeReservedNodeOfferings(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) (*redshift.DescribeReservedNodeOfferingsOutput, error)
	DescribeReservedNodeOfferingsAsync(ctx workflow.Context, input *redshift.DescribeReservedNodeOfferingsInput) *RedshiftDescribeReservedNodeOfferingsFuture

	DescribeReservedNodes(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) (*redshift.DescribeReservedNodesOutput, error)
	DescribeReservedNodesAsync(ctx workflow.Context, input *redshift.DescribeReservedNodesInput) *RedshiftDescribeReservedNodesFuture

	DescribeResize(ctx workflow.Context, input *redshift.DescribeResizeInput) (*redshift.DescribeResizeOutput, error)
	DescribeResizeAsync(ctx workflow.Context, input *redshift.DescribeResizeInput) *RedshiftDescribeResizeFuture

	DescribeScheduledActions(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) (*redshift.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsAsync(ctx workflow.Context, input *redshift.DescribeScheduledActionsInput) *RedshiftDescribeScheduledActionsFuture

	DescribeSnapshotCopyGrants(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) (*redshift.DescribeSnapshotCopyGrantsOutput, error)
	DescribeSnapshotCopyGrantsAsync(ctx workflow.Context, input *redshift.DescribeSnapshotCopyGrantsInput) *RedshiftDescribeSnapshotCopyGrantsFuture

	DescribeSnapshotSchedules(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) (*redshift.DescribeSnapshotSchedulesOutput, error)
	DescribeSnapshotSchedulesAsync(ctx workflow.Context, input *redshift.DescribeSnapshotSchedulesInput) *RedshiftDescribeSnapshotSchedulesFuture

	DescribeStorage(ctx workflow.Context, input *redshift.DescribeStorageInput) (*redshift.DescribeStorageOutput, error)
	DescribeStorageAsync(ctx workflow.Context, input *redshift.DescribeStorageInput) *RedshiftDescribeStorageFuture

	DescribeTableRestoreStatus(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) (*redshift.DescribeTableRestoreStatusOutput, error)
	DescribeTableRestoreStatusAsync(ctx workflow.Context, input *redshift.DescribeTableRestoreStatusInput) *RedshiftDescribeTableRestoreStatusFuture

	DescribeTags(ctx workflow.Context, input *redshift.DescribeTagsInput) (*redshift.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *redshift.DescribeTagsInput) *RedshiftDescribeTagsFuture

	DescribeUsageLimits(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) (*redshift.DescribeUsageLimitsOutput, error)
	DescribeUsageLimitsAsync(ctx workflow.Context, input *redshift.DescribeUsageLimitsInput) *RedshiftDescribeUsageLimitsFuture

	DisableLogging(ctx workflow.Context, input *redshift.DisableLoggingInput) (*redshift.LoggingStatus, error)
	DisableLoggingAsync(ctx workflow.Context, input *redshift.DisableLoggingInput) *RedshiftDisableLoggingFuture

	DisableSnapshotCopy(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) (*redshift.DisableSnapshotCopyOutput, error)
	DisableSnapshotCopyAsync(ctx workflow.Context, input *redshift.DisableSnapshotCopyInput) *RedshiftDisableSnapshotCopyFuture

	EnableLogging(ctx workflow.Context, input *redshift.EnableLoggingInput) (*redshift.LoggingStatus, error)
	EnableLoggingAsync(ctx workflow.Context, input *redshift.EnableLoggingInput) *RedshiftEnableLoggingFuture

	EnableSnapshotCopy(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) (*redshift.EnableSnapshotCopyOutput, error)
	EnableSnapshotCopyAsync(ctx workflow.Context, input *redshift.EnableSnapshotCopyInput) *RedshiftEnableSnapshotCopyFuture

	GetClusterCredentials(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) (*redshift.GetClusterCredentialsOutput, error)
	GetClusterCredentialsAsync(ctx workflow.Context, input *redshift.GetClusterCredentialsInput) *RedshiftGetClusterCredentialsFuture

	GetReservedNodeExchangeOfferings(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) (*redshift.GetReservedNodeExchangeOfferingsOutput, error)
	GetReservedNodeExchangeOfferingsAsync(ctx workflow.Context, input *redshift.GetReservedNodeExchangeOfferingsInput) *RedshiftGetReservedNodeExchangeOfferingsFuture

	ModifyCluster(ctx workflow.Context, input *redshift.ModifyClusterInput) (*redshift.ModifyClusterOutput, error)
	ModifyClusterAsync(ctx workflow.Context, input *redshift.ModifyClusterInput) *RedshiftModifyClusterFuture

	ModifyClusterDbRevision(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) (*redshift.ModifyClusterDbRevisionOutput, error)
	ModifyClusterDbRevisionAsync(ctx workflow.Context, input *redshift.ModifyClusterDbRevisionInput) *RedshiftModifyClusterDbRevisionFuture

	ModifyClusterIamRoles(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) (*redshift.ModifyClusterIamRolesOutput, error)
	ModifyClusterIamRolesAsync(ctx workflow.Context, input *redshift.ModifyClusterIamRolesInput) *RedshiftModifyClusterIamRolesFuture

	ModifyClusterMaintenance(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) (*redshift.ModifyClusterMaintenanceOutput, error)
	ModifyClusterMaintenanceAsync(ctx workflow.Context, input *redshift.ModifyClusterMaintenanceInput) *RedshiftModifyClusterMaintenanceFuture

	ModifyClusterParameterGroup(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)
	ModifyClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterParameterGroupInput) *RedshiftModifyClusterParameterGroupFuture

	ModifyClusterSnapshot(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) (*redshift.ModifyClusterSnapshotOutput, error)
	ModifyClusterSnapshotAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotInput) *RedshiftModifyClusterSnapshotFuture

	ModifyClusterSnapshotSchedule(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) (*redshift.ModifyClusterSnapshotScheduleOutput, error)
	ModifyClusterSnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifyClusterSnapshotScheduleInput) *RedshiftModifyClusterSnapshotScheduleFuture

	ModifyClusterSubnetGroup(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) (*redshift.ModifyClusterSubnetGroupOutput, error)
	ModifyClusterSubnetGroupAsync(ctx workflow.Context, input *redshift.ModifyClusterSubnetGroupInput) *RedshiftModifyClusterSubnetGroupFuture

	ModifyEventSubscription(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) (*redshift.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionAsync(ctx workflow.Context, input *redshift.ModifyEventSubscriptionInput) *RedshiftModifyEventSubscriptionFuture

	ModifyScheduledAction(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) (*redshift.ModifyScheduledActionOutput, error)
	ModifyScheduledActionAsync(ctx workflow.Context, input *redshift.ModifyScheduledActionInput) *RedshiftModifyScheduledActionFuture

	ModifySnapshotCopyRetentionPeriod(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) (*redshift.ModifySnapshotCopyRetentionPeriodOutput, error)
	ModifySnapshotCopyRetentionPeriodAsync(ctx workflow.Context, input *redshift.ModifySnapshotCopyRetentionPeriodInput) *RedshiftModifySnapshotCopyRetentionPeriodFuture

	ModifySnapshotSchedule(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) (*redshift.ModifySnapshotScheduleOutput, error)
	ModifySnapshotScheduleAsync(ctx workflow.Context, input *redshift.ModifySnapshotScheduleInput) *RedshiftModifySnapshotScheduleFuture

	ModifyUsageLimit(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) (*redshift.ModifyUsageLimitOutput, error)
	ModifyUsageLimitAsync(ctx workflow.Context, input *redshift.ModifyUsageLimitInput) *RedshiftModifyUsageLimitFuture

	PauseCluster(ctx workflow.Context, input *redshift.PauseClusterInput) (*redshift.PauseClusterOutput, error)
	PauseClusterAsync(ctx workflow.Context, input *redshift.PauseClusterInput) *RedshiftPauseClusterFuture

	PurchaseReservedNodeOffering(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) (*redshift.PurchaseReservedNodeOfferingOutput, error)
	PurchaseReservedNodeOfferingAsync(ctx workflow.Context, input *redshift.PurchaseReservedNodeOfferingInput) *RedshiftPurchaseReservedNodeOfferingFuture

	RebootCluster(ctx workflow.Context, input *redshift.RebootClusterInput) (*redshift.RebootClusterOutput, error)
	RebootClusterAsync(ctx workflow.Context, input *redshift.RebootClusterInput) *RedshiftRebootClusterFuture

	ResetClusterParameterGroup(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) (*redshift.ClusterParameterGroupNameMessage, error)
	ResetClusterParameterGroupAsync(ctx workflow.Context, input *redshift.ResetClusterParameterGroupInput) *RedshiftResetClusterParameterGroupFuture

	ResizeCluster(ctx workflow.Context, input *redshift.ResizeClusterInput) (*redshift.ResizeClusterOutput, error)
	ResizeClusterAsync(ctx workflow.Context, input *redshift.ResizeClusterInput) *RedshiftResizeClusterFuture

	RestoreFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) (*redshift.RestoreFromClusterSnapshotOutput, error)
	RestoreFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreFromClusterSnapshotInput) *RedshiftRestoreFromClusterSnapshotFuture

	RestoreTableFromClusterSnapshot(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) (*redshift.RestoreTableFromClusterSnapshotOutput, error)
	RestoreTableFromClusterSnapshotAsync(ctx workflow.Context, input *redshift.RestoreTableFromClusterSnapshotInput) *RedshiftRestoreTableFromClusterSnapshotFuture

	ResumeCluster(ctx workflow.Context, input *redshift.ResumeClusterInput) (*redshift.ResumeClusterOutput, error)
	ResumeClusterAsync(ctx workflow.Context, input *redshift.ResumeClusterInput) *RedshiftResumeClusterFuture

	RevokeClusterSecurityGroupIngress(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) (*redshift.RevokeClusterSecurityGroupIngressOutput, error)
	RevokeClusterSecurityGroupIngressAsync(ctx workflow.Context, input *redshift.RevokeClusterSecurityGroupIngressInput) *RedshiftRevokeClusterSecurityGroupIngressFuture

	RevokeSnapshotAccess(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) (*redshift.RevokeSnapshotAccessOutput, error)
	RevokeSnapshotAccessAsync(ctx workflow.Context, input *redshift.RevokeSnapshotAccessInput) *RedshiftRevokeSnapshotAccessFuture

	RotateEncryptionKey(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) (*redshift.RotateEncryptionKeyOutput, error)
	RotateEncryptionKeyAsync(ctx workflow.Context, input *redshift.RotateEncryptionKeyInput) *RedshiftRotateEncryptionKeyFuture

	WaitUntilClusterAvailable(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilClusterAvailableAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *awsclients.VoidFuture

	WaitUntilClusterDeleted(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilClusterDeletedAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *awsclients.VoidFuture

	WaitUntilClusterRestored(ctx workflow.Context, input *redshift.DescribeClustersInput) error
	WaitUntilClusterRestoredAsync(ctx workflow.Context, input *redshift.DescribeClustersInput) *awsclients.VoidFuture

	WaitUntilSnapshotAvailable(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) error
	WaitUntilSnapshotAvailableAsync(ctx workflow.Context, input *redshift.DescribeClusterSnapshotsInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

