// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elbstub

import (
	"github.com/aws/aws-sdk-go/service/elb"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTags(ctx workflow.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elb.AddTagsInput) *ELBAddTagsFuture

	ApplySecurityGroupsToLoadBalancer(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	ApplySecurityGroupsToLoadBalancerAsync(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) *ELBApplySecurityGroupsToLoadBalancerFuture

	AttachLoadBalancerToSubnets(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	AttachLoadBalancerToSubnetsAsync(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) *ELBAttachLoadBalancerToSubnetsFuture

	ConfigureHealthCheck(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)
	ConfigureHealthCheckAsync(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) *ELBConfigureHealthCheckFuture

	CreateAppCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error)
	CreateAppCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) *ELBCreateAppCookieStickinessPolicyFuture

	CreateLBCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error)
	CreateLBCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) *ELBCreateLBCookieStickinessPolicyFuture

	CreateLoadBalancer(ctx workflow.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *elb.CreateLoadBalancerInput) *ELBCreateLoadBalancerFuture

	CreateLoadBalancerListeners(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)
	CreateLoadBalancerListenersAsync(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) *ELBCreateLoadBalancerListenersFuture

	CreateLoadBalancerPolicy(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)
	CreateLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) *ELBCreateLoadBalancerPolicyFuture

	DeleteLoadBalancer(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) *ELBDeleteLoadBalancerFuture

	DeleteLoadBalancerListeners(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)
	DeleteLoadBalancerListenersAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) *ELBDeleteLoadBalancerListenersFuture

	DeleteLoadBalancerPolicy(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error)
	DeleteLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) *ELBDeleteLoadBalancerPolicyFuture

	DeregisterInstancesFromLoadBalancer(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	DeregisterInstancesFromLoadBalancerAsync(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) *ELBDeregisterInstancesFromLoadBalancerFuture

	DescribeAccountLimits(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) *ELBDescribeAccountLimitsFuture

	DescribeInstanceHealth(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error)
	DescribeInstanceHealthAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *ELBDescribeInstanceHealthFuture

	DescribeLoadBalancerAttributes(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) *ELBDescribeLoadBalancerAttributesFuture

	DescribeLoadBalancerPolicies(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPoliciesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) *ELBDescribeLoadBalancerPoliciesFuture

	DescribeLoadBalancerPolicyTypes(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancerPolicyTypesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) *ELBDescribeLoadBalancerPolicyTypesFuture

	DescribeLoadBalancers(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) *ELBDescribeLoadBalancersFuture

	DescribeTags(ctx workflow.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *elb.DescribeTagsInput) *ELBDescribeTagsFuture

	DetachLoadBalancerFromSubnets(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	DetachLoadBalancerFromSubnetsAsync(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) *ELBDetachLoadBalancerFromSubnetsFuture

	DisableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)
	DisableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) *ELBDisableAvailabilityZonesForLoadBalancerFuture

	EnableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)
	EnableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) *ELBEnableAvailabilityZonesForLoadBalancerFuture

	ModifyLoadBalancerAttributes(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) *ELBModifyLoadBalancerAttributesFuture

	RegisterInstancesWithLoadBalancer(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	RegisterInstancesWithLoadBalancerAsync(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) *ELBRegisterInstancesWithLoadBalancerFuture

	RemoveTags(ctx workflow.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elb.RemoveTagsInput) *ELBRemoveTagsFuture

	SetLoadBalancerListenerSSLCertificate(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)
	SetLoadBalancerListenerSSLCertificateAsync(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) *ELBSetLoadBalancerListenerSSLCertificateFuture

	SetLoadBalancerPoliciesForBackendServer(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesForBackendServerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) *ELBSetLoadBalancerPoliciesForBackendServerFuture

	SetLoadBalancerPoliciesOfListener(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
	SetLoadBalancerPoliciesOfListenerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) *ELBSetLoadBalancerPoliciesOfListenerFuture

	WaitUntilAnyInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error
	WaitUntilAnyInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture

	WaitUntilInstanceDeregistered(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceDeregisteredAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture

	WaitUntilInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
