// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kafkastub

import (
	"github.com/aws/aws-sdk-go/service/kafka"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	BatchAssociateScramSecret(ctx workflow.Context, input *kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error)
	BatchAssociateScramSecretAsync(ctx workflow.Context, input *kafka.BatchAssociateScramSecretInput) *KafkaBatchAssociateScramSecretFuture

	BatchDisassociateScramSecret(ctx workflow.Context, input *kafka.BatchDisassociateScramSecretInput) (*kafka.BatchDisassociateScramSecretOutput, error)
	BatchDisassociateScramSecretAsync(ctx workflow.Context, input *kafka.BatchDisassociateScramSecretInput) *KafkaBatchDisassociateScramSecretFuture

	CreateCluster(ctx workflow.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *kafka.CreateClusterInput) *KafkaCreateClusterFuture

	CreateConfiguration(ctx workflow.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error)
	CreateConfigurationAsync(ctx workflow.Context, input *kafka.CreateConfigurationInput) *KafkaCreateConfigurationFuture

	DeleteCluster(ctx workflow.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *kafka.DeleteClusterInput) *KafkaDeleteClusterFuture

	DeleteConfiguration(ctx workflow.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error)
	DeleteConfigurationAsync(ctx workflow.Context, input *kafka.DeleteConfigurationInput) *KafkaDeleteConfigurationFuture

	DescribeCluster(ctx workflow.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *kafka.DescribeClusterInput) *KafkaDescribeClusterFuture

	DescribeClusterOperation(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterOperationAsync(ctx workflow.Context, input *kafka.DescribeClusterOperationInput) *KafkaDescribeClusterOperationFuture

	DescribeConfiguration(ctx workflow.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationAsync(ctx workflow.Context, input *kafka.DescribeConfigurationInput) *KafkaDescribeConfigurationFuture

	DescribeConfigurationRevision(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error)
	DescribeConfigurationRevisionAsync(ctx workflow.Context, input *kafka.DescribeConfigurationRevisionInput) *KafkaDescribeConfigurationRevisionFuture

	GetBootstrapBrokers(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersAsync(ctx workflow.Context, input *kafka.GetBootstrapBrokersInput) *KafkaGetBootstrapBrokersFuture

	GetCompatibleKafkaVersions(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	GetCompatibleKafkaVersionsAsync(ctx workflow.Context, input *kafka.GetCompatibleKafkaVersionsInput) *KafkaGetCompatibleKafkaVersionsFuture

	ListClusterOperations(ctx workflow.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsAsync(ctx workflow.Context, input *kafka.ListClusterOperationsInput) *KafkaListClusterOperationsFuture

	ListClusters(ctx workflow.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *kafka.ListClustersInput) *KafkaListClustersFuture

	ListConfigurationRevisions(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsAsync(ctx workflow.Context, input *kafka.ListConfigurationRevisionsInput) *KafkaListConfigurationRevisionsFuture

	ListConfigurations(ctx workflow.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsAsync(ctx workflow.Context, input *kafka.ListConfigurationsInput) *KafkaListConfigurationsFuture

	ListKafkaVersions(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsAsync(ctx workflow.Context, input *kafka.ListKafkaVersionsInput) *KafkaListKafkaVersionsFuture

	ListNodes(ctx workflow.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error)
	ListNodesAsync(ctx workflow.Context, input *kafka.ListNodesInput) *KafkaListNodesFuture

	ListScramSecrets(ctx workflow.Context, input *kafka.ListScramSecretsInput) (*kafka.ListScramSecretsOutput, error)
	ListScramSecretsAsync(ctx workflow.Context, input *kafka.ListScramSecretsInput) *KafkaListScramSecretsFuture

	ListTagsForResource(ctx workflow.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kafka.ListTagsForResourceInput) *KafkaListTagsForResourceFuture

	RebootBroker(ctx workflow.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error)
	RebootBrokerAsync(ctx workflow.Context, input *kafka.RebootBrokerInput) *KafkaRebootBrokerFuture

	TagResource(ctx workflow.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kafka.TagResourceInput) *KafkaTagResourceFuture

	UntagResource(ctx workflow.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kafka.UntagResourceInput) *KafkaUntagResourceFuture

	UpdateBrokerCount(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerCountAsync(ctx workflow.Context, input *kafka.UpdateBrokerCountInput) *KafkaUpdateBrokerCountFuture

	UpdateBrokerStorage(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateBrokerStorageAsync(ctx workflow.Context, input *kafka.UpdateBrokerStorageInput) *KafkaUpdateBrokerStorageFuture

	UpdateClusterConfiguration(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterConfigurationAsync(ctx workflow.Context, input *kafka.UpdateClusterConfigurationInput) *KafkaUpdateClusterConfigurationFuture

	UpdateClusterKafkaVersion(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateClusterKafkaVersionAsync(ctx workflow.Context, input *kafka.UpdateClusterKafkaVersionInput) *KafkaUpdateClusterKafkaVersionFuture

	UpdateConfiguration(ctx workflow.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error)
	UpdateConfigurationAsync(ctx workflow.Context, input *kafka.UpdateConfigurationInput) *KafkaUpdateConfigurationFuture

	UpdateMonitoring(ctx workflow.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error)
	UpdateMonitoringAsync(ctx workflow.Context, input *kafka.UpdateMonitoringInput) *KafkaUpdateMonitoringFuture
}

func NewClient() Client {
	return &stub{}
}

