// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client ecsiface.ECSAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := ecs.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (ecsiface.ECSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return ecs.New(sess), nil
}

func (a *Activities) CreateCapacityProvider(ctx context.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCapacityProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCluster(ctx context.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateService(ctx context.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateTaskSet(ctx context.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateTaskSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAccountSetting(ctx context.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAccountSettingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAttributes(ctx context.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCapacityProvider(ctx context.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCapacityProviderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCluster(ctx context.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteService(ctx context.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTaskSet(ctx context.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTaskSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterContainerInstance(ctx context.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterContainerInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterTaskDefinition(ctx context.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterTaskDefinitionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCapacityProviders(ctx context.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCapacityProvidersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeClusters(ctx context.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContainerInstances(ctx context.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContainerInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeServices(ctx context.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeServicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTaskDefinition(ctx context.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTaskDefinitionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTaskSets(ctx context.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTaskSetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTasks(ctx context.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTasksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DiscoverPollEndpoint(ctx context.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DiscoverPollEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAccountSettings(ctx context.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAccountSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAttributes(ctx context.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListClusters(ctx context.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListContainerInstances(ctx context.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListContainerInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListServices(ctx context.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListServicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTaskDefinitionFamilies(ctx context.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTaskDefinitionFamiliesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTaskDefinitions(ctx context.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTaskDefinitionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTasks(ctx context.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTasksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAccountSetting(ctx context.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAccountSettingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAccountSettingDefault(ctx context.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAccountSettingDefaultWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAttributes(ctx context.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutClusterCapacityProviders(ctx context.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutClusterCapacityProvidersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterContainerInstance(ctx context.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterContainerInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterTaskDefinition(ctx context.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterTaskDefinitionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RunTask(ctx context.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RunTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartTask(ctx context.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopTask(ctx context.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SubmitAttachmentStateChanges(ctx context.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SubmitAttachmentStateChangesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SubmitContainerStateChange(ctx context.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SubmitContainerStateChangeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SubmitTaskStateChange(ctx context.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SubmitTaskStateChangeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateClusterSettings(ctx context.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateClusterSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContainerAgent(ctx context.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContainerAgentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContainerInstancesState(ctx context.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContainerInstancesStateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateService(ctx context.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateServiceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateServicePrimaryTaskSet(ctx context.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateServicePrimaryTaskSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTaskSet(ctx context.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTaskSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilServicesInactive(ctx context.Context, input *ecs.DescribeServicesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilServicesInactiveWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilServicesStable(ctx context.Context, input *ecs.DescribeServicesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilServicesStableWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTasksRunning(ctx context.Context, input *ecs.DescribeTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTasksRunningWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilTasksStopped(ctx context.Context, input *ecs.DescribeTasksInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilTasksStoppedWithContext(ctx, input, options...))
	})
}
