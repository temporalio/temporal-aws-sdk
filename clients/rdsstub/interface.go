// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package rdsstub

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddRoleToDBCluster(ctx workflow.Context, input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error)
	AddRoleToDBClusterAsync(ctx workflow.Context, input *rds.AddRoleToDBClusterInput) *RDSAddRoleToDBClusterFuture

	AddRoleToDBInstance(ctx workflow.Context, input *rds.AddRoleToDBInstanceInput) (*rds.AddRoleToDBInstanceOutput, error)
	AddRoleToDBInstanceAsync(ctx workflow.Context, input *rds.AddRoleToDBInstanceInput) *RDSAddRoleToDBInstanceFuture

	AddSourceIdentifierToSubscription(ctx workflow.Context, input *rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error)
	AddSourceIdentifierToSubscriptionAsync(ctx workflow.Context, input *rds.AddSourceIdentifierToSubscriptionInput) *RDSAddSourceIdentifierToSubscriptionFuture

	AddTagsToResource(ctx workflow.Context, input *rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *rds.AddTagsToResourceInput) *RDSAddTagsToResourceFuture

	ApplyPendingMaintenanceAction(ctx workflow.Context, input *rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error)
	ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *rds.ApplyPendingMaintenanceActionInput) *RDSApplyPendingMaintenanceActionFuture

	AuthorizeDBSecurityGroupIngress(ctx workflow.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error)
	AuthorizeDBSecurityGroupIngressAsync(ctx workflow.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) *RDSAuthorizeDBSecurityGroupIngressFuture

	BacktrackDBCluster(ctx workflow.Context, input *rds.BacktrackDBClusterInput) (*rds.BacktrackDBClusterOutput, error)
	BacktrackDBClusterAsync(ctx workflow.Context, input *rds.BacktrackDBClusterInput) *RDSBacktrackDBClusterFuture

	CancelExportTask(ctx workflow.Context, input *rds.CancelExportTaskInput) (*rds.CancelExportTaskOutput, error)
	CancelExportTaskAsync(ctx workflow.Context, input *rds.CancelExportTaskInput) *RDSCancelExportTaskFuture

	CopyDBClusterParameterGroup(ctx workflow.Context, input *rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error)
	CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.CopyDBClusterParameterGroupInput) *RDSCopyDBClusterParameterGroupFuture

	CopyDBClusterSnapshot(ctx workflow.Context, input *rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error)
	CopyDBClusterSnapshotAsync(ctx workflow.Context, input *rds.CopyDBClusterSnapshotInput) *RDSCopyDBClusterSnapshotFuture

	CopyDBParameterGroup(ctx workflow.Context, input *rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error)
	CopyDBParameterGroupAsync(ctx workflow.Context, input *rds.CopyDBParameterGroupInput) *RDSCopyDBParameterGroupFuture

	CopyDBSnapshot(ctx workflow.Context, input *rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error)
	CopyDBSnapshotAsync(ctx workflow.Context, input *rds.CopyDBSnapshotInput) *RDSCopyDBSnapshotFuture

	CopyOptionGroup(ctx workflow.Context, input *rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error)
	CopyOptionGroupAsync(ctx workflow.Context, input *rds.CopyOptionGroupInput) *RDSCopyOptionGroupFuture

	CreateCustomAvailabilityZone(ctx workflow.Context, input *rds.CreateCustomAvailabilityZoneInput) (*rds.CreateCustomAvailabilityZoneOutput, error)
	CreateCustomAvailabilityZoneAsync(ctx workflow.Context, input *rds.CreateCustomAvailabilityZoneInput) *RDSCreateCustomAvailabilityZoneFuture

	CreateDBCluster(ctx workflow.Context, input *rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error)
	CreateDBClusterAsync(ctx workflow.Context, input *rds.CreateDBClusterInput) *RDSCreateDBClusterFuture

	CreateDBClusterEndpoint(ctx workflow.Context, input *rds.CreateDBClusterEndpointInput) (*rds.CreateDBClusterEndpointOutput, error)
	CreateDBClusterEndpointAsync(ctx workflow.Context, input *rds.CreateDBClusterEndpointInput) *RDSCreateDBClusterEndpointFuture

	CreateDBClusterParameterGroup(ctx workflow.Context, input *rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error)
	CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.CreateDBClusterParameterGroupInput) *RDSCreateDBClusterParameterGroupFuture

	CreateDBClusterSnapshot(ctx workflow.Context, input *rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error)
	CreateDBClusterSnapshotAsync(ctx workflow.Context, input *rds.CreateDBClusterSnapshotInput) *RDSCreateDBClusterSnapshotFuture

	CreateDBInstance(ctx workflow.Context, input *rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error)
	CreateDBInstanceAsync(ctx workflow.Context, input *rds.CreateDBInstanceInput) *RDSCreateDBInstanceFuture

	CreateDBInstanceReadReplica(ctx workflow.Context, input *rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error)
	CreateDBInstanceReadReplicaAsync(ctx workflow.Context, input *rds.CreateDBInstanceReadReplicaInput) *RDSCreateDBInstanceReadReplicaFuture

	CreateDBParameterGroup(ctx workflow.Context, input *rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error)
	CreateDBParameterGroupAsync(ctx workflow.Context, input *rds.CreateDBParameterGroupInput) *RDSCreateDBParameterGroupFuture

	CreateDBProxy(ctx workflow.Context, input *rds.CreateDBProxyInput) (*rds.CreateDBProxyOutput, error)
	CreateDBProxyAsync(ctx workflow.Context, input *rds.CreateDBProxyInput) *RDSCreateDBProxyFuture

	CreateDBSecurityGroup(ctx workflow.Context, input *rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error)
	CreateDBSecurityGroupAsync(ctx workflow.Context, input *rds.CreateDBSecurityGroupInput) *RDSCreateDBSecurityGroupFuture

	CreateDBSnapshot(ctx workflow.Context, input *rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error)
	CreateDBSnapshotAsync(ctx workflow.Context, input *rds.CreateDBSnapshotInput) *RDSCreateDBSnapshotFuture

	CreateDBSubnetGroup(ctx workflow.Context, input *rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error)
	CreateDBSubnetGroupAsync(ctx workflow.Context, input *rds.CreateDBSubnetGroupInput) *RDSCreateDBSubnetGroupFuture

	CreateEventSubscription(ctx workflow.Context, input *rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionAsync(ctx workflow.Context, input *rds.CreateEventSubscriptionInput) *RDSCreateEventSubscriptionFuture

	CreateGlobalCluster(ctx workflow.Context, input *rds.CreateGlobalClusterInput) (*rds.CreateGlobalClusterOutput, error)
	CreateGlobalClusterAsync(ctx workflow.Context, input *rds.CreateGlobalClusterInput) *RDSCreateGlobalClusterFuture

	CreateOptionGroup(ctx workflow.Context, input *rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error)
	CreateOptionGroupAsync(ctx workflow.Context, input *rds.CreateOptionGroupInput) *RDSCreateOptionGroupFuture

	DeleteCustomAvailabilityZone(ctx workflow.Context, input *rds.DeleteCustomAvailabilityZoneInput) (*rds.DeleteCustomAvailabilityZoneOutput, error)
	DeleteCustomAvailabilityZoneAsync(ctx workflow.Context, input *rds.DeleteCustomAvailabilityZoneInput) *RDSDeleteCustomAvailabilityZoneFuture

	DeleteDBCluster(ctx workflow.Context, input *rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error)
	DeleteDBClusterAsync(ctx workflow.Context, input *rds.DeleteDBClusterInput) *RDSDeleteDBClusterFuture

	DeleteDBClusterEndpoint(ctx workflow.Context, input *rds.DeleteDBClusterEndpointInput) (*rds.DeleteDBClusterEndpointOutput, error)
	DeleteDBClusterEndpointAsync(ctx workflow.Context, input *rds.DeleteDBClusterEndpointInput) *RDSDeleteDBClusterEndpointFuture

	DeleteDBClusterParameterGroup(ctx workflow.Context, input *rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error)
	DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.DeleteDBClusterParameterGroupInput) *RDSDeleteDBClusterParameterGroupFuture

	DeleteDBClusterSnapshot(ctx workflow.Context, input *rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error)
	DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *rds.DeleteDBClusterSnapshotInput) *RDSDeleteDBClusterSnapshotFuture

	DeleteDBInstance(ctx workflow.Context, input *rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error)
	DeleteDBInstanceAsync(ctx workflow.Context, input *rds.DeleteDBInstanceInput) *RDSDeleteDBInstanceFuture

	DeleteDBInstanceAutomatedBackup(ctx workflow.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) (*rds.DeleteDBInstanceAutomatedBackupOutput, error)
	DeleteDBInstanceAutomatedBackupAsync(ctx workflow.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) *RDSDeleteDBInstanceAutomatedBackupFuture

	DeleteDBParameterGroup(ctx workflow.Context, input *rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error)
	DeleteDBParameterGroupAsync(ctx workflow.Context, input *rds.DeleteDBParameterGroupInput) *RDSDeleteDBParameterGroupFuture

	DeleteDBProxy(ctx workflow.Context, input *rds.DeleteDBProxyInput) (*rds.DeleteDBProxyOutput, error)
	DeleteDBProxyAsync(ctx workflow.Context, input *rds.DeleteDBProxyInput) *RDSDeleteDBProxyFuture

	DeleteDBSecurityGroup(ctx workflow.Context, input *rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error)
	DeleteDBSecurityGroupAsync(ctx workflow.Context, input *rds.DeleteDBSecurityGroupInput) *RDSDeleteDBSecurityGroupFuture

	DeleteDBSnapshot(ctx workflow.Context, input *rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error)
	DeleteDBSnapshotAsync(ctx workflow.Context, input *rds.DeleteDBSnapshotInput) *RDSDeleteDBSnapshotFuture

	DeleteDBSubnetGroup(ctx workflow.Context, input *rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error)
	DeleteDBSubnetGroupAsync(ctx workflow.Context, input *rds.DeleteDBSubnetGroupInput) *RDSDeleteDBSubnetGroupFuture

	DeleteEventSubscription(ctx workflow.Context, input *rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionAsync(ctx workflow.Context, input *rds.DeleteEventSubscriptionInput) *RDSDeleteEventSubscriptionFuture

	DeleteGlobalCluster(ctx workflow.Context, input *rds.DeleteGlobalClusterInput) (*rds.DeleteGlobalClusterOutput, error)
	DeleteGlobalClusterAsync(ctx workflow.Context, input *rds.DeleteGlobalClusterInput) *RDSDeleteGlobalClusterFuture

	DeleteInstallationMedia(ctx workflow.Context, input *rds.DeleteInstallationMediaInput) (*rds.DeleteInstallationMediaOutput, error)
	DeleteInstallationMediaAsync(ctx workflow.Context, input *rds.DeleteInstallationMediaInput) *RDSDeleteInstallationMediaFuture

	DeleteOptionGroup(ctx workflow.Context, input *rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error)
	DeleteOptionGroupAsync(ctx workflow.Context, input *rds.DeleteOptionGroupInput) *RDSDeleteOptionGroupFuture

	DeregisterDBProxyTargets(ctx workflow.Context, input *rds.DeregisterDBProxyTargetsInput) (*rds.DeregisterDBProxyTargetsOutput, error)
	DeregisterDBProxyTargetsAsync(ctx workflow.Context, input *rds.DeregisterDBProxyTargetsInput) *RDSDeregisterDBProxyTargetsFuture

	DescribeAccountAttributes(ctx workflow.Context, input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *rds.DescribeAccountAttributesInput) *RDSDescribeAccountAttributesFuture

	DescribeCertificates(ctx workflow.Context, input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error)
	DescribeCertificatesAsync(ctx workflow.Context, input *rds.DescribeCertificatesInput) *RDSDescribeCertificatesFuture

	DescribeCustomAvailabilityZones(ctx workflow.Context, input *rds.DescribeCustomAvailabilityZonesInput) (*rds.DescribeCustomAvailabilityZonesOutput, error)
	DescribeCustomAvailabilityZonesAsync(ctx workflow.Context, input *rds.DescribeCustomAvailabilityZonesInput) *RDSDescribeCustomAvailabilityZonesFuture

	DescribeDBClusterBacktracks(ctx workflow.Context, input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error)
	DescribeDBClusterBacktracksAsync(ctx workflow.Context, input *rds.DescribeDBClusterBacktracksInput) *RDSDescribeDBClusterBacktracksFuture

	DescribeDBClusterEndpoints(ctx workflow.Context, input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error)
	DescribeDBClusterEndpointsAsync(ctx workflow.Context, input *rds.DescribeDBClusterEndpointsInput) *RDSDescribeDBClusterEndpointsFuture

	DescribeDBClusterParameterGroups(ctx workflow.Context, input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *rds.DescribeDBClusterParameterGroupsInput) *RDSDescribeDBClusterParameterGroupsFuture

	DescribeDBClusterParameters(ctx workflow.Context, input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterParametersAsync(ctx workflow.Context, input *rds.DescribeDBClusterParametersInput) *RDSDescribeDBClusterParametersFuture

	DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) *RDSDescribeDBClusterSnapshotAttributesFuture

	DescribeDBClusterSnapshots(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) *RDSDescribeDBClusterSnapshotsFuture

	DescribeDBClusters(ctx workflow.Context, input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error)
	DescribeDBClustersAsync(ctx workflow.Context, input *rds.DescribeDBClustersInput) *RDSDescribeDBClustersFuture

	DescribeDBEngineVersions(ctx workflow.Context, input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error)
	DescribeDBEngineVersionsAsync(ctx workflow.Context, input *rds.DescribeDBEngineVersionsInput) *RDSDescribeDBEngineVersionsFuture

	DescribeDBInstanceAutomatedBackups(ctx workflow.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error)
	DescribeDBInstanceAutomatedBackupsAsync(ctx workflow.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) *RDSDescribeDBInstanceAutomatedBackupsFuture

	DescribeDBInstances(ctx workflow.Context, input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error)
	DescribeDBInstancesAsync(ctx workflow.Context, input *rds.DescribeDBInstancesInput) *RDSDescribeDBInstancesFuture

	DescribeDBLogFiles(ctx workflow.Context, input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error)
	DescribeDBLogFilesAsync(ctx workflow.Context, input *rds.DescribeDBLogFilesInput) *RDSDescribeDBLogFilesFuture

	DescribeDBParameterGroups(ctx workflow.Context, input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameterGroupsAsync(ctx workflow.Context, input *rds.DescribeDBParameterGroupsInput) *RDSDescribeDBParameterGroupsFuture

	DescribeDBParameters(ctx workflow.Context, input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error)
	DescribeDBParametersAsync(ctx workflow.Context, input *rds.DescribeDBParametersInput) *RDSDescribeDBParametersFuture

	DescribeDBProxies(ctx workflow.Context, input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error)
	DescribeDBProxiesAsync(ctx workflow.Context, input *rds.DescribeDBProxiesInput) *RDSDescribeDBProxiesFuture

	DescribeDBProxyTargetGroups(ctx workflow.Context, input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error)
	DescribeDBProxyTargetGroupsAsync(ctx workflow.Context, input *rds.DescribeDBProxyTargetGroupsInput) *RDSDescribeDBProxyTargetGroupsFuture

	DescribeDBProxyTargets(ctx workflow.Context, input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error)
	DescribeDBProxyTargetsAsync(ctx workflow.Context, input *rds.DescribeDBProxyTargetsInput) *RDSDescribeDBProxyTargetsFuture

	DescribeDBSecurityGroups(ctx workflow.Context, input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSecurityGroupsAsync(ctx workflow.Context, input *rds.DescribeDBSecurityGroupsInput) *RDSDescribeDBSecurityGroupsFuture

	DescribeDBSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error)
	DescribeDBSnapshotAttributesAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotAttributesInput) *RDSDescribeDBSnapshotAttributesFuture

	DescribeDBSnapshots(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error)
	DescribeDBSnapshotsAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) *RDSDescribeDBSnapshotsFuture

	DescribeDBSubnetGroups(ctx workflow.Context, input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error)
	DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *rds.DescribeDBSubnetGroupsInput) *RDSDescribeDBSubnetGroupsFuture

	DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error)
	DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *rds.DescribeEngineDefaultClusterParametersInput) *RDSDescribeEngineDefaultClusterParametersFuture

	DescribeEngineDefaultParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *rds.DescribeEngineDefaultParametersInput) *RDSDescribeEngineDefaultParametersFuture

	DescribeEventCategories(ctx workflow.Context, input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesAsync(ctx workflow.Context, input *rds.DescribeEventCategoriesInput) *RDSDescribeEventCategoriesFuture

	DescribeEventSubscriptions(ctx workflow.Context, input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsAsync(ctx workflow.Context, input *rds.DescribeEventSubscriptionsInput) *RDSDescribeEventSubscriptionsFuture

	DescribeEvents(ctx workflow.Context, input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *rds.DescribeEventsInput) *RDSDescribeEventsFuture

	DescribeExportTasks(ctx workflow.Context, input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error)
	DescribeExportTasksAsync(ctx workflow.Context, input *rds.DescribeExportTasksInput) *RDSDescribeExportTasksFuture

	DescribeGlobalClusters(ctx workflow.Context, input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error)
	DescribeGlobalClustersAsync(ctx workflow.Context, input *rds.DescribeGlobalClustersInput) *RDSDescribeGlobalClustersFuture

	DescribeInstallationMedia(ctx workflow.Context, input *rds.DescribeInstallationMediaInput) (*rds.DescribeInstallationMediaOutput, error)
	DescribeInstallationMediaAsync(ctx workflow.Context, input *rds.DescribeInstallationMediaInput) *RDSDescribeInstallationMediaFuture

	DescribeOptionGroupOptions(ctx workflow.Context, input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error)
	DescribeOptionGroupOptionsAsync(ctx workflow.Context, input *rds.DescribeOptionGroupOptionsInput) *RDSDescribeOptionGroupOptionsFuture

	DescribeOptionGroups(ctx workflow.Context, input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error)
	DescribeOptionGroupsAsync(ctx workflow.Context, input *rds.DescribeOptionGroupsInput) *RDSDescribeOptionGroupsFuture

	DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) *RDSDescribeOrderableDBInstanceOptionsFuture

	DescribePendingMaintenanceActions(ctx workflow.Context, input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error)
	DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *rds.DescribePendingMaintenanceActionsInput) *RDSDescribePendingMaintenanceActionsFuture

	DescribeReservedDBInstances(ctx workflow.Context, input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error)
	DescribeReservedDBInstancesAsync(ctx workflow.Context, input *rds.DescribeReservedDBInstancesInput) *RDSDescribeReservedDBInstancesFuture

	DescribeReservedDBInstancesOfferings(ctx workflow.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)
	DescribeReservedDBInstancesOfferingsAsync(ctx workflow.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) *RDSDescribeReservedDBInstancesOfferingsFuture

	DescribeSourceRegions(ctx workflow.Context, input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error)
	DescribeSourceRegionsAsync(ctx workflow.Context, input *rds.DescribeSourceRegionsInput) *RDSDescribeSourceRegionsFuture

	DescribeValidDBInstanceModifications(ctx workflow.Context, input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error)
	DescribeValidDBInstanceModificationsAsync(ctx workflow.Context, input *rds.DescribeValidDBInstanceModificationsInput) *RDSDescribeValidDBInstanceModificationsFuture

	DownloadDBLogFilePortion(ctx workflow.Context, input *rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error)
	DownloadDBLogFilePortionAsync(ctx workflow.Context, input *rds.DownloadDBLogFilePortionInput) *RDSDownloadDBLogFilePortionFuture

	FailoverDBCluster(ctx workflow.Context, input *rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error)
	FailoverDBClusterAsync(ctx workflow.Context, input *rds.FailoverDBClusterInput) *RDSFailoverDBClusterFuture

	ImportInstallationMedia(ctx workflow.Context, input *rds.ImportInstallationMediaInput) (*rds.ImportInstallationMediaOutput, error)
	ImportInstallationMediaAsync(ctx workflow.Context, input *rds.ImportInstallationMediaInput) *RDSImportInstallationMediaFuture

	ListTagsForResource(ctx workflow.Context, input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *rds.ListTagsForResourceInput) *RDSListTagsForResourceFuture

	ModifyCertificates(ctx workflow.Context, input *rds.ModifyCertificatesInput) (*rds.ModifyCertificatesOutput, error)
	ModifyCertificatesAsync(ctx workflow.Context, input *rds.ModifyCertificatesInput) *RDSModifyCertificatesFuture

	ModifyCurrentDBClusterCapacity(ctx workflow.Context, input *rds.ModifyCurrentDBClusterCapacityInput) (*rds.ModifyCurrentDBClusterCapacityOutput, error)
	ModifyCurrentDBClusterCapacityAsync(ctx workflow.Context, input *rds.ModifyCurrentDBClusterCapacityInput) *RDSModifyCurrentDBClusterCapacityFuture

	ModifyDBCluster(ctx workflow.Context, input *rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error)
	ModifyDBClusterAsync(ctx workflow.Context, input *rds.ModifyDBClusterInput) *RDSModifyDBClusterFuture

	ModifyDBClusterEndpoint(ctx workflow.Context, input *rds.ModifyDBClusterEndpointInput) (*rds.ModifyDBClusterEndpointOutput, error)
	ModifyDBClusterEndpointAsync(ctx workflow.Context, input *rds.ModifyDBClusterEndpointInput) *RDSModifyDBClusterEndpointFuture

	ModifyDBClusterParameterGroup(ctx workflow.Context, input *rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)
	ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.ModifyDBClusterParameterGroupInput) *RDSModifyDBClusterParameterGroupFuture

	ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error)
	ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) *RDSModifyDBClusterSnapshotAttributeFuture

	ModifyDBInstance(ctx workflow.Context, input *rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error)
	ModifyDBInstanceAsync(ctx workflow.Context, input *rds.ModifyDBInstanceInput) *RDSModifyDBInstanceFuture

	ModifyDBParameterGroup(ctx workflow.Context, input *rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)
	ModifyDBParameterGroupAsync(ctx workflow.Context, input *rds.ModifyDBParameterGroupInput) *RDSModifyDBParameterGroupFuture

	ModifyDBProxy(ctx workflow.Context, input *rds.ModifyDBProxyInput) (*rds.ModifyDBProxyOutput, error)
	ModifyDBProxyAsync(ctx workflow.Context, input *rds.ModifyDBProxyInput) *RDSModifyDBProxyFuture

	ModifyDBProxyTargetGroup(ctx workflow.Context, input *rds.ModifyDBProxyTargetGroupInput) (*rds.ModifyDBProxyTargetGroupOutput, error)
	ModifyDBProxyTargetGroupAsync(ctx workflow.Context, input *rds.ModifyDBProxyTargetGroupInput) *RDSModifyDBProxyTargetGroupFuture

	ModifyDBSnapshot(ctx workflow.Context, input *rds.ModifyDBSnapshotInput) (*rds.ModifyDBSnapshotOutput, error)
	ModifyDBSnapshotAsync(ctx workflow.Context, input *rds.ModifyDBSnapshotInput) *RDSModifyDBSnapshotFuture

	ModifyDBSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error)
	ModifyDBSnapshotAttributeAsync(ctx workflow.Context, input *rds.ModifyDBSnapshotAttributeInput) *RDSModifyDBSnapshotAttributeFuture

	ModifyDBSubnetGroup(ctx workflow.Context, input *rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error)
	ModifyDBSubnetGroupAsync(ctx workflow.Context, input *rds.ModifyDBSubnetGroupInput) *RDSModifyDBSubnetGroupFuture

	ModifyEventSubscription(ctx workflow.Context, input *rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionAsync(ctx workflow.Context, input *rds.ModifyEventSubscriptionInput) *RDSModifyEventSubscriptionFuture

	ModifyGlobalCluster(ctx workflow.Context, input *rds.ModifyGlobalClusterInput) (*rds.ModifyGlobalClusterOutput, error)
	ModifyGlobalClusterAsync(ctx workflow.Context, input *rds.ModifyGlobalClusterInput) *RDSModifyGlobalClusterFuture

	ModifyOptionGroup(ctx workflow.Context, input *rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error)
	ModifyOptionGroupAsync(ctx workflow.Context, input *rds.ModifyOptionGroupInput) *RDSModifyOptionGroupFuture

	PromoteReadReplica(ctx workflow.Context, input *rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error)
	PromoteReadReplicaAsync(ctx workflow.Context, input *rds.PromoteReadReplicaInput) *RDSPromoteReadReplicaFuture

	PromoteReadReplicaDBCluster(ctx workflow.Context, input *rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error)
	PromoteReadReplicaDBClusterAsync(ctx workflow.Context, input *rds.PromoteReadReplicaDBClusterInput) *RDSPromoteReadReplicaDBClusterFuture

	PurchaseReservedDBInstancesOffering(ctx workflow.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error)
	PurchaseReservedDBInstancesOfferingAsync(ctx workflow.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) *RDSPurchaseReservedDBInstancesOfferingFuture

	RebootDBInstance(ctx workflow.Context, input *rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error)
	RebootDBInstanceAsync(ctx workflow.Context, input *rds.RebootDBInstanceInput) *RDSRebootDBInstanceFuture

	RegisterDBProxyTargets(ctx workflow.Context, input *rds.RegisterDBProxyTargetsInput) (*rds.RegisterDBProxyTargetsOutput, error)
	RegisterDBProxyTargetsAsync(ctx workflow.Context, input *rds.RegisterDBProxyTargetsInput) *RDSRegisterDBProxyTargetsFuture

	RemoveFromGlobalCluster(ctx workflow.Context, input *rds.RemoveFromGlobalClusterInput) (*rds.RemoveFromGlobalClusterOutput, error)
	RemoveFromGlobalClusterAsync(ctx workflow.Context, input *rds.RemoveFromGlobalClusterInput) *RDSRemoveFromGlobalClusterFuture

	RemoveRoleFromDBCluster(ctx workflow.Context, input *rds.RemoveRoleFromDBClusterInput) (*rds.RemoveRoleFromDBClusterOutput, error)
	RemoveRoleFromDBClusterAsync(ctx workflow.Context, input *rds.RemoveRoleFromDBClusterInput) *RDSRemoveRoleFromDBClusterFuture

	RemoveRoleFromDBInstance(ctx workflow.Context, input *rds.RemoveRoleFromDBInstanceInput) (*rds.RemoveRoleFromDBInstanceOutput, error)
	RemoveRoleFromDBInstanceAsync(ctx workflow.Context, input *rds.RemoveRoleFromDBInstanceInput) *RDSRemoveRoleFromDBInstanceFuture

	RemoveSourceIdentifierFromSubscription(ctx workflow.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error)
	RemoveSourceIdentifierFromSubscriptionAsync(ctx workflow.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) *RDSRemoveSourceIdentifierFromSubscriptionFuture

	RemoveTagsFromResource(ctx workflow.Context, input *rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *rds.RemoveTagsFromResourceInput) *RDSRemoveTagsFromResourceFuture

	ResetDBClusterParameterGroup(ctx workflow.Context, input *rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)
	ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.ResetDBClusterParameterGroupInput) *RDSResetDBClusterParameterGroupFuture

	ResetDBParameterGroup(ctx workflow.Context, input *rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)
	ResetDBParameterGroupAsync(ctx workflow.Context, input *rds.ResetDBParameterGroupInput) *RDSResetDBParameterGroupFuture

	RestoreDBClusterFromS3(ctx workflow.Context, input *rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error)
	RestoreDBClusterFromS3Async(ctx workflow.Context, input *rds.RestoreDBClusterFromS3Input) *RDSRestoreDBClusterFromS3Future

	RestoreDBClusterFromSnapshot(ctx workflow.Context, input *rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error)
	RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *rds.RestoreDBClusterFromSnapshotInput) *RDSRestoreDBClusterFromSnapshotFuture

	RestoreDBClusterToPointInTime(ctx workflow.Context, input *rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error)
	RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *rds.RestoreDBClusterToPointInTimeInput) *RDSRestoreDBClusterToPointInTimeFuture

	RestoreDBInstanceFromDBSnapshot(ctx workflow.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error)
	RestoreDBInstanceFromDBSnapshotAsync(ctx workflow.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) *RDSRestoreDBInstanceFromDBSnapshotFuture

	RestoreDBInstanceFromS3(ctx workflow.Context, input *rds.RestoreDBInstanceFromS3Input) (*rds.RestoreDBInstanceFromS3Output, error)
	RestoreDBInstanceFromS3Async(ctx workflow.Context, input *rds.RestoreDBInstanceFromS3Input) *RDSRestoreDBInstanceFromS3Future

	RestoreDBInstanceToPointInTime(ctx workflow.Context, input *rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error)
	RestoreDBInstanceToPointInTimeAsync(ctx workflow.Context, input *rds.RestoreDBInstanceToPointInTimeInput) *RDSRestoreDBInstanceToPointInTimeFuture

	RevokeDBSecurityGroupIngress(ctx workflow.Context, input *rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error)
	RevokeDBSecurityGroupIngressAsync(ctx workflow.Context, input *rds.RevokeDBSecurityGroupIngressInput) *RDSRevokeDBSecurityGroupIngressFuture

	StartActivityStream(ctx workflow.Context, input *rds.StartActivityStreamInput) (*rds.StartActivityStreamOutput, error)
	StartActivityStreamAsync(ctx workflow.Context, input *rds.StartActivityStreamInput) *RDSStartActivityStreamFuture

	StartDBCluster(ctx workflow.Context, input *rds.StartDBClusterInput) (*rds.StartDBClusterOutput, error)
	StartDBClusterAsync(ctx workflow.Context, input *rds.StartDBClusterInput) *RDSStartDBClusterFuture

	StartDBInstance(ctx workflow.Context, input *rds.StartDBInstanceInput) (*rds.StartDBInstanceOutput, error)
	StartDBInstanceAsync(ctx workflow.Context, input *rds.StartDBInstanceInput) *RDSStartDBInstanceFuture

	StartExportTask(ctx workflow.Context, input *rds.StartExportTaskInput) (*rds.StartExportTaskOutput, error)
	StartExportTaskAsync(ctx workflow.Context, input *rds.StartExportTaskInput) *RDSStartExportTaskFuture

	StopActivityStream(ctx workflow.Context, input *rds.StopActivityStreamInput) (*rds.StopActivityStreamOutput, error)
	StopActivityStreamAsync(ctx workflow.Context, input *rds.StopActivityStreamInput) *RDSStopActivityStreamFuture

	StopDBCluster(ctx workflow.Context, input *rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error)
	StopDBClusterAsync(ctx workflow.Context, input *rds.StopDBClusterInput) *RDSStopDBClusterFuture

	StopDBInstance(ctx workflow.Context, input *rds.StopDBInstanceInput) (*rds.StopDBInstanceOutput, error)
	StopDBInstanceAsync(ctx workflow.Context, input *rds.StopDBInstanceInput) *RDSStopDBInstanceFuture

	WaitUntilDBClusterSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error
	WaitUntilDBClusterSnapshotAvailableAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) *clients.VoidFuture

	WaitUntilDBClusterSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error
	WaitUntilDBClusterSnapshotDeletedAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) *clients.VoidFuture

	WaitUntilDBInstanceAvailable(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error
	WaitUntilDBInstanceAvailableAsync(ctx workflow.Context, input *rds.DescribeDBInstancesInput) *clients.VoidFuture

	WaitUntilDBInstanceDeleted(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error
	WaitUntilDBInstanceDeletedAsync(ctx workflow.Context, input *rds.DescribeDBInstancesInput) *clients.VoidFuture

	WaitUntilDBSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error
	WaitUntilDBSnapshotAvailableAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) *clients.VoidFuture

	WaitUntilDBSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error
	WaitUntilDBSnapshotDeletedAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
