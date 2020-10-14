// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client elbiface.ELBAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elb.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elbiface.ELBAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return elb.New(sess), nil
}

func (a *Activities) AddTags(ctx context.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ApplySecurityGroupsToLoadBalancer(ctx context.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApplySecurityGroupsToLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AttachLoadBalancerToSubnets(ctx context.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AttachLoadBalancerToSubnetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ConfigureHealthCheck(ctx context.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ConfigureHealthCheckWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateAppCookieStickinessPolicy(ctx context.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAppCookieStickinessPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLBCookieStickinessPolicy(ctx context.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLBCookieStickinessPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLoadBalancer(ctx context.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLoadBalancerListeners(ctx context.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLoadBalancerListenersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLoadBalancerPolicy(ctx context.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLoadBalancerPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLoadBalancer(ctx context.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLoadBalancerListeners(ctx context.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLoadBalancerListenersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLoadBalancerPolicy(ctx context.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLoadBalancerPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterInstancesFromLoadBalancer(ctx context.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterInstancesFromLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountLimits(ctx context.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstanceHealth(ctx context.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstanceHealthWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancerAttributes(ctx context.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancerAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancerPolicies(ctx context.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancerPoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancerPolicyTypes(ctx context.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancerPolicyTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoadBalancers(ctx context.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoadBalancersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeTags(ctx context.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetachLoadBalancerFromSubnets(ctx context.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetachLoadBalancerFromSubnetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableAvailabilityZonesForLoadBalancer(ctx context.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableAvailabilityZonesForLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableAvailabilityZonesForLoadBalancer(ctx context.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableAvailabilityZonesForLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyLoadBalancerAttributes(ctx context.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyLoadBalancerAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterInstancesWithLoadBalancer(ctx context.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterInstancesWithLoadBalancerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTags(ctx context.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetLoadBalancerListenerSSLCertificate(ctx context.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetLoadBalancerListenerSSLCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetLoadBalancerPoliciesForBackendServer(ctx context.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetLoadBalancerPoliciesForBackendServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SetLoadBalancerPoliciesOfListener(ctx context.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SetLoadBalancerPoliciesOfListenerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilAnyInstanceInService(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilAnyInstanceInServiceWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilInstanceDeregistered(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilInstanceDeregisteredWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilInstanceInService(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilInstanceInServiceWithContext(ctx, input, options...))
	})
}
