// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elasticachestub

import (
	"github.com/aws/aws-sdk-go/service/elasticache"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToResource(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) (*elasticache.TagListMessage, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *elasticache.AddTagsToResourceInput) *ElastiCacheAddTagsToResourceFuture

	AuthorizeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) (*elasticache.AuthorizeCacheSecurityGroupIngressOutput, error)
	AuthorizeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.AuthorizeCacheSecurityGroupIngressInput) *ElastiCacheAuthorizeCacheSecurityGroupIngressFuture

	BatchApplyUpdateAction(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error)
	BatchApplyUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchApplyUpdateActionInput) *ElastiCacheBatchApplyUpdateActionFuture

	BatchStopUpdateAction(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error)
	BatchStopUpdateActionAsync(ctx workflow.Context, input *elasticache.BatchStopUpdateActionInput) *ElastiCacheBatchStopUpdateActionFuture

	CompleteMigration(ctx workflow.Context, input *elasticache.CompleteMigrationInput) (*elasticache.CompleteMigrationOutput, error)
	CompleteMigrationAsync(ctx workflow.Context, input *elasticache.CompleteMigrationInput) *ElastiCacheCompleteMigrationFuture

	CopySnapshot(ctx workflow.Context, input *elasticache.CopySnapshotInput) (*elasticache.CopySnapshotOutput, error)
	CopySnapshotAsync(ctx workflow.Context, input *elasticache.CopySnapshotInput) *ElastiCacheCopySnapshotFuture

	CreateCacheCluster(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) (*elasticache.CreateCacheClusterOutput, error)
	CreateCacheClusterAsync(ctx workflow.Context, input *elasticache.CreateCacheClusterInput) *ElastiCacheCreateCacheClusterFuture

	CreateCacheParameterGroup(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) (*elasticache.CreateCacheParameterGroupOutput, error)
	CreateCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheParameterGroupInput) *ElastiCacheCreateCacheParameterGroupFuture

	CreateCacheSecurityGroup(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) (*elasticache.CreateCacheSecurityGroupOutput, error)
	CreateCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSecurityGroupInput) *ElastiCacheCreateCacheSecurityGroupFuture

	CreateCacheSubnetGroup(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) (*elasticache.CreateCacheSubnetGroupOutput, error)
	CreateCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.CreateCacheSubnetGroupInput) *ElastiCacheCreateCacheSubnetGroupFuture

	CreateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) (*elasticache.CreateGlobalReplicationGroupOutput, error)
	CreateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateGlobalReplicationGroupInput) *ElastiCacheCreateGlobalReplicationGroupFuture

	CreateReplicationGroup(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) (*elasticache.CreateReplicationGroupOutput, error)
	CreateReplicationGroupAsync(ctx workflow.Context, input *elasticache.CreateReplicationGroupInput) *ElastiCacheCreateReplicationGroupFuture

	CreateSnapshot(ctx workflow.Context, input *elasticache.CreateSnapshotInput) (*elasticache.CreateSnapshotOutput, error)
	CreateSnapshotAsync(ctx workflow.Context, input *elasticache.CreateSnapshotInput) *ElastiCacheCreateSnapshotFuture

	CreateUser(ctx workflow.Context, input *elasticache.CreateUserInput) (*elasticache.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *elasticache.CreateUserInput) *ElastiCacheCreateUserFuture

	CreateUserGroup(ctx workflow.Context, input *elasticache.CreateUserGroupInput) (*elasticache.CreateUserGroupOutput, error)
	CreateUserGroupAsync(ctx workflow.Context, input *elasticache.CreateUserGroupInput) *ElastiCacheCreateUserGroupFuture

	DecreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.DecreaseNodeGroupsInGlobalReplicationGroupOutput, error)
	DecreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DecreaseNodeGroupsInGlobalReplicationGroupInput) *ElastiCacheDecreaseNodeGroupsInGlobalReplicationGroupFuture

	DecreaseReplicaCount(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) (*elasticache.DecreaseReplicaCountOutput, error)
	DecreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.DecreaseReplicaCountInput) *ElastiCacheDecreaseReplicaCountFuture

	DeleteCacheCluster(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) (*elasticache.DeleteCacheClusterOutput, error)
	DeleteCacheClusterAsync(ctx workflow.Context, input *elasticache.DeleteCacheClusterInput) *ElastiCacheDeleteCacheClusterFuture

	DeleteCacheParameterGroup(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) (*elasticache.DeleteCacheParameterGroupOutput, error)
	DeleteCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheParameterGroupInput) *ElastiCacheDeleteCacheParameterGroupFuture

	DeleteCacheSecurityGroup(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) (*elasticache.DeleteCacheSecurityGroupOutput, error)
	DeleteCacheSecurityGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSecurityGroupInput) *ElastiCacheDeleteCacheSecurityGroupFuture

	DeleteCacheSubnetGroup(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) (*elasticache.DeleteCacheSubnetGroupOutput, error)
	DeleteCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.DeleteCacheSubnetGroupInput) *ElastiCacheDeleteCacheSubnetGroupFuture

	DeleteGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) (*elasticache.DeleteGlobalReplicationGroupOutput, error)
	DeleteGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteGlobalReplicationGroupInput) *ElastiCacheDeleteGlobalReplicationGroupFuture

	DeleteReplicationGroup(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) (*elasticache.DeleteReplicationGroupOutput, error)
	DeleteReplicationGroupAsync(ctx workflow.Context, input *elasticache.DeleteReplicationGroupInput) *ElastiCacheDeleteReplicationGroupFuture

	DeleteSnapshot(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) (*elasticache.DeleteSnapshotOutput, error)
	DeleteSnapshotAsync(ctx workflow.Context, input *elasticache.DeleteSnapshotInput) *ElastiCacheDeleteSnapshotFuture

	DeleteUser(ctx workflow.Context, input *elasticache.DeleteUserInput) (*elasticache.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *elasticache.DeleteUserInput) *ElastiCacheDeleteUserFuture

	DeleteUserGroup(ctx workflow.Context, input *elasticache.DeleteUserGroupInput) (*elasticache.DeleteUserGroupOutput, error)
	DeleteUserGroupAsync(ctx workflow.Context, input *elasticache.DeleteUserGroupInput) *ElastiCacheDeleteUserGroupFuture

	DescribeCacheClusters(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheClustersAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) *ElastiCacheDescribeCacheClustersFuture

	DescribeCacheEngineVersions(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheEngineVersionsAsync(ctx workflow.Context, input *elasticache.DescribeCacheEngineVersionsInput) *ElastiCacheDescribeCacheEngineVersionsFuture

	DescribeCacheParameterGroups(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameterGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheParameterGroupsInput) *ElastiCacheDescribeCacheParameterGroupsFuture

	DescribeCacheParameters(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheParametersAsync(ctx workflow.Context, input *elasticache.DescribeCacheParametersInput) *ElastiCacheDescribeCacheParametersFuture

	DescribeCacheSecurityGroups(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
	DescribeCacheSecurityGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSecurityGroupsInput) *ElastiCacheDescribeCacheSecurityGroupsFuture

	DescribeCacheSubnetGroups(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeCacheSubnetGroupsAsync(ctx workflow.Context, input *elasticache.DescribeCacheSubnetGroupsInput) *ElastiCacheDescribeCacheSubnetGroupsFuture

	DescribeEngineDefaultParameters(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *elasticache.DescribeEngineDefaultParametersInput) *ElastiCacheDescribeEngineDefaultParametersFuture

	DescribeEvents(ctx workflow.Context, input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *elasticache.DescribeEventsInput) *ElastiCacheDescribeEventsFuture

	DescribeGlobalReplicationGroups(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error)
	DescribeGlobalReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeGlobalReplicationGroupsInput) *ElastiCacheDescribeGlobalReplicationGroupsFuture

	DescribeReplicationGroups(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReplicationGroupsAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) *ElastiCacheDescribeReplicationGroupsFuture

	DescribeReservedCacheNodes(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesInput) *ElastiCacheDescribeReservedCacheNodesFuture

	DescribeReservedCacheNodesOfferings(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeReservedCacheNodesOfferingsAsync(ctx workflow.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput) *ElastiCacheDescribeReservedCacheNodesOfferingsFuture

	DescribeServiceUpdates(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error)
	DescribeServiceUpdatesAsync(ctx workflow.Context, input *elasticache.DescribeServiceUpdatesInput) *ElastiCacheDescribeServiceUpdatesFuture

	DescribeSnapshots(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeSnapshotsAsync(ctx workflow.Context, input *elasticache.DescribeSnapshotsInput) *ElastiCacheDescribeSnapshotsFuture

	DescribeUpdateActions(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error)
	DescribeUpdateActionsAsync(ctx workflow.Context, input *elasticache.DescribeUpdateActionsInput) *ElastiCacheDescribeUpdateActionsFuture

	DescribeUserGroups(ctx workflow.Context, input *elasticache.DescribeUserGroupsInput) (*elasticache.DescribeUserGroupsOutput, error)
	DescribeUserGroupsAsync(ctx workflow.Context, input *elasticache.DescribeUserGroupsInput) *ElastiCacheDescribeUserGroupsFuture

	DescribeUsers(ctx workflow.Context, input *elasticache.DescribeUsersInput) (*elasticache.DescribeUsersOutput, error)
	DescribeUsersAsync(ctx workflow.Context, input *elasticache.DescribeUsersInput) *ElastiCacheDescribeUsersFuture

	DisassociateGlobalReplicationGroup(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) (*elasticache.DisassociateGlobalReplicationGroupOutput, error)
	DisassociateGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.DisassociateGlobalReplicationGroupInput) *ElastiCacheDisassociateGlobalReplicationGroupFuture

	FailoverGlobalReplicationGroup(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) (*elasticache.FailoverGlobalReplicationGroupOutput, error)
	FailoverGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.FailoverGlobalReplicationGroupInput) *ElastiCacheFailoverGlobalReplicationGroupFuture

	IncreaseNodeGroupsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) (*elasticache.IncreaseNodeGroupsInGlobalReplicationGroupOutput, error)
	IncreaseNodeGroupsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.IncreaseNodeGroupsInGlobalReplicationGroupInput) *ElastiCacheIncreaseNodeGroupsInGlobalReplicationGroupFuture

	IncreaseReplicaCount(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) (*elasticache.IncreaseReplicaCountOutput, error)
	IncreaseReplicaCountAsync(ctx workflow.Context, input *elasticache.IncreaseReplicaCountInput) *ElastiCacheIncreaseReplicaCountFuture

	ListAllowedNodeTypeModifications(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
	ListAllowedNodeTypeModificationsAsync(ctx workflow.Context, input *elasticache.ListAllowedNodeTypeModificationsInput) *ElastiCacheListAllowedNodeTypeModificationsFuture

	ListTagsForResource(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) (*elasticache.TagListMessage, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *elasticache.ListTagsForResourceInput) *ElastiCacheListTagsForResourceFuture

	ModifyCacheCluster(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) (*elasticache.ModifyCacheClusterOutput, error)
	ModifyCacheClusterAsync(ctx workflow.Context, input *elasticache.ModifyCacheClusterInput) *ElastiCacheModifyCacheClusterFuture

	ModifyCacheParameterGroup(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
	ModifyCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheParameterGroupInput) *ElastiCacheModifyCacheParameterGroupFuture

	ModifyCacheSubnetGroup(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) (*elasticache.ModifyCacheSubnetGroupOutput, error)
	ModifyCacheSubnetGroupAsync(ctx workflow.Context, input *elasticache.ModifyCacheSubnetGroupInput) *ElastiCacheModifyCacheSubnetGroupFuture

	ModifyGlobalReplicationGroup(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) (*elasticache.ModifyGlobalReplicationGroupOutput, error)
	ModifyGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyGlobalReplicationGroupInput) *ElastiCacheModifyGlobalReplicationGroupFuture

	ModifyReplicationGroup(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) (*elasticache.ModifyReplicationGroupOutput, error)
	ModifyReplicationGroupAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupInput) *ElastiCacheModifyReplicationGroupFuture

	ModifyReplicationGroupShardConfiguration(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) (*elasticache.ModifyReplicationGroupShardConfigurationOutput, error)
	ModifyReplicationGroupShardConfigurationAsync(ctx workflow.Context, input *elasticache.ModifyReplicationGroupShardConfigurationInput) *ElastiCacheModifyReplicationGroupShardConfigurationFuture

	ModifyUser(ctx workflow.Context, input *elasticache.ModifyUserInput) (*elasticache.ModifyUserOutput, error)
	ModifyUserAsync(ctx workflow.Context, input *elasticache.ModifyUserInput) *ElastiCacheModifyUserFuture

	ModifyUserGroup(ctx workflow.Context, input *elasticache.ModifyUserGroupInput) (*elasticache.ModifyUserGroupOutput, error)
	ModifyUserGroupAsync(ctx workflow.Context, input *elasticache.ModifyUserGroupInput) *ElastiCacheModifyUserGroupFuture

	PurchaseReservedCacheNodesOffering(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) (*elasticache.PurchaseReservedCacheNodesOfferingOutput, error)
	PurchaseReservedCacheNodesOfferingAsync(ctx workflow.Context, input *elasticache.PurchaseReservedCacheNodesOfferingInput) *ElastiCachePurchaseReservedCacheNodesOfferingFuture

	RebalanceSlotsInGlobalReplicationGroup(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) (*elasticache.RebalanceSlotsInGlobalReplicationGroupOutput, error)
	RebalanceSlotsInGlobalReplicationGroupAsync(ctx workflow.Context, input *elasticache.RebalanceSlotsInGlobalReplicationGroupInput) *ElastiCacheRebalanceSlotsInGlobalReplicationGroupFuture

	RebootCacheCluster(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) (*elasticache.RebootCacheClusterOutput, error)
	RebootCacheClusterAsync(ctx workflow.Context, input *elasticache.RebootCacheClusterInput) *ElastiCacheRebootCacheClusterFuture

	RemoveTagsFromResource(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) (*elasticache.TagListMessage, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *elasticache.RemoveTagsFromResourceInput) *ElastiCacheRemoveTagsFromResourceFuture

	ResetCacheParameterGroup(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) (*elasticache.CacheParameterGroupNameMessage, error)
	ResetCacheParameterGroupAsync(ctx workflow.Context, input *elasticache.ResetCacheParameterGroupInput) *ElastiCacheResetCacheParameterGroupFuture

	RevokeCacheSecurityGroupIngress(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) (*elasticache.RevokeCacheSecurityGroupIngressOutput, error)
	RevokeCacheSecurityGroupIngressAsync(ctx workflow.Context, input *elasticache.RevokeCacheSecurityGroupIngressInput) *ElastiCacheRevokeCacheSecurityGroupIngressFuture

	StartMigration(ctx workflow.Context, input *elasticache.StartMigrationInput) (*elasticache.StartMigrationOutput, error)
	StartMigrationAsync(ctx workflow.Context, input *elasticache.StartMigrationInput) *ElastiCacheStartMigrationFuture

	TestFailover(ctx workflow.Context, input *elasticache.TestFailoverInput) (*elasticache.TestFailoverOutput, error)
	TestFailoverAsync(ctx workflow.Context, input *elasticache.TestFailoverInput) *ElastiCacheTestFailoverFuture

	WaitUntilCacheClusterAvailable(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterAvailableAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) *clients.VoidFuture

	WaitUntilCacheClusterDeleted(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) error
	WaitUntilCacheClusterDeletedAsync(ctx workflow.Context, input *elasticache.DescribeCacheClustersInput) *clients.VoidFuture

	WaitUntilReplicationGroupAvailable(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupAvailableAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) *clients.VoidFuture

	WaitUntilReplicationGroupDeleted(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) error
	WaitUntilReplicationGroupDeletedAsync(ctx workflow.Context, input *elasticache.DescribeReplicationGroupsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
