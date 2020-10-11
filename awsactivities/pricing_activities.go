// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type PricingActivities struct {
	client pricingiface.PricingAPI

	sessionFactory SessionFactory
}

func NewPricingActivities(sess *session.Session, config ...*aws.Config) *PricingActivities {
	client := pricing.New(sess, config...)
	return &PricingActivities{client: client}
}

func NewPricingActivitiesWithSessionFactory(sessionFactory SessionFactory) *PricingActivities {
	return &PricingActivities{sessionFactory: sessionFactory}
}

func (a *PricingActivities) getClient(ctx context.Context) (pricingiface.PricingAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return pricing.New(sess), nil
}

func (a *PricingActivities) DescribeServices(ctx context.Context, input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServicesWithContext(ctx, input)
}

func (a *PricingActivities) GetAttributeValues(ctx context.Context, input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAttributeValuesWithContext(ctx, input)
}

func (a *PricingActivities) GetProducts(ctx context.Context, input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetProductsWithContext(ctx, input)
}
