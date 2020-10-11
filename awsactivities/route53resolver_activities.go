// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type Route53ResolverActivities struct {
	client route53resolveriface.Route53ResolverAPI

	sessionFactory SessionFactory
}

func NewRoute53ResolverActivities(sess *session.Session, config ...*aws.Config) *Route53ResolverActivities {
	client := route53resolver.New(sess, config...)
	return &Route53ResolverActivities{client: client}
}

func NewRoute53ResolverActivitiesWithSessionFactory(sessionFactory SessionFactory) *Route53ResolverActivities {
	return &Route53ResolverActivities{sessionFactory: sessionFactory}
}

func (a *Route53ResolverActivities) getClient(ctx context.Context) (route53resolveriface.Route53ResolverAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return route53resolver.New(sess), nil
}

func (a *Route53ResolverActivities) AssociateResolverEndpointIpAddress(ctx context.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateResolverEndpointIpAddressWithContext(ctx, input)
}

func (a *Route53ResolverActivities) AssociateResolverQueryLogConfig(ctx context.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) AssociateResolverRule(ctx context.Context, input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverEndpoint(ctx context.Context, input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverQueryLogConfig(ctx context.Context, input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverRule(ctx context.Context, input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverEndpoint(ctx context.Context, input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverQueryLogConfig(ctx context.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverRule(ctx context.Context, input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverEndpointIpAddress(ctx context.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateResolverEndpointIpAddressWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverQueryLogConfig(ctx context.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverRule(ctx context.Context, input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverEndpoint(ctx context.Context, input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfig(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigAssociation(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverQueryLogConfigAssociationWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigPolicy(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverQueryLogConfigPolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRule(ctx context.Context, input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRuleAssociation(ctx context.Context, input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverRuleAssociationWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRulePolicy(ctx context.Context, input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverRulePolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverEndpointIpAddresses(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverEndpointIpAddressesWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverEndpoints(ctx context.Context, input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverEndpointsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigAssociations(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverQueryLogConfigAssociationsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigs(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverQueryLogConfigsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverRuleAssociations(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverRuleAssociationsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverRules(ctx context.Context, input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolverRulesWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListTagsForResource(ctx context.Context, input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) PutResolverQueryLogConfigPolicy(ctx context.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutResolverQueryLogConfigPolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) PutResolverRulePolicy(ctx context.Context, input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutResolverRulePolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) TagResource(ctx context.Context, input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UntagResource(ctx context.Context, input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UpdateResolverEndpoint(ctx context.Context, input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UpdateResolverRule(ctx context.Context, input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResolverRuleWithContext(ctx, input)
}
