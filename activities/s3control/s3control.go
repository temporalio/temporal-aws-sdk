// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client s3controliface.S3ControlAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := s3control.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (s3controliface.S3ControlAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return s3control.New(sess), nil
}

func (a *Activities) CreateAccessPoint(ctx context.Context, input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAccessPointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBucket(ctx context.Context, input *s3control.CreateBucketInput) (*s3control.CreateBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBucketWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateJob(ctx context.Context, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAccessPoint(ctx context.Context, input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAccessPointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAccessPointPolicy(ctx context.Context, input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAccessPointPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBucket(ctx context.Context, input *s3control.DeleteBucketInput) (*s3control.DeleteBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBucketWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBucketLifecycleConfiguration(ctx context.Context, input *s3control.DeleteBucketLifecycleConfigurationInput) (*s3control.DeleteBucketLifecycleConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBucketLifecycleConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBucketPolicy(ctx context.Context, input *s3control.DeleteBucketPolicyInput) (*s3control.DeleteBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBucketPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBucketTagging(ctx context.Context, input *s3control.DeleteBucketTaggingInput) (*s3control.DeleteBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBucketTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteJobTagging(ctx context.Context, input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteJobTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePublicAccessBlock(ctx context.Context, input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePublicAccessBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeJob(ctx context.Context, input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAccessPoint(ctx context.Context, input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAccessPointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAccessPointPolicy(ctx context.Context, input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAccessPointPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAccessPointPolicyStatus(ctx context.Context, input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAccessPointPolicyStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBucket(ctx context.Context, input *s3control.GetBucketInput) (*s3control.GetBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBucketWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBucketLifecycleConfiguration(ctx context.Context, input *s3control.GetBucketLifecycleConfigurationInput) (*s3control.GetBucketLifecycleConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBucketLifecycleConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBucketPolicy(ctx context.Context, input *s3control.GetBucketPolicyInput) (*s3control.GetBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBucketPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBucketTagging(ctx context.Context, input *s3control.GetBucketTaggingInput) (*s3control.GetBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBucketTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetJobTagging(ctx context.Context, input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetJobTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPublicAccessBlock(ctx context.Context, input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPublicAccessBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAccessPoints(ctx context.Context, input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAccessPointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobs(ctx context.Context, input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRegionalBuckets(ctx context.Context, input *s3control.ListRegionalBucketsInput) (*s3control.ListRegionalBucketsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRegionalBucketsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAccessPointPolicy(ctx context.Context, input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAccessPointPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutBucketLifecycleConfiguration(ctx context.Context, input *s3control.PutBucketLifecycleConfigurationInput) (*s3control.PutBucketLifecycleConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutBucketLifecycleConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutBucketPolicy(ctx context.Context, input *s3control.PutBucketPolicyInput) (*s3control.PutBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutBucketPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutBucketTagging(ctx context.Context, input *s3control.PutBucketTaggingInput) (*s3control.PutBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutBucketTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutJobTagging(ctx context.Context, input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutJobTaggingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutPublicAccessBlock(ctx context.Context, input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutPublicAccessBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateJobPriority(ctx context.Context, input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateJobPriorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateJobStatus(ctx context.Context, input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateJobStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
