// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type S3Activities struct {
	client s3iface.S3API

	sessionFactory SessionFactory
}

func NewS3Activities(sess *session.Session, config ...*aws.Config) *S3Activities {
	client := s3.New(sess, config...)
	return &S3Activities{client: client}
}

func NewS3ActivitiesWithSessionFactory(sessionFactory SessionFactory) *S3Activities {
	return &S3Activities{sessionFactory: sessionFactory}
}

func (a *S3Activities) getClient(ctx context.Context) (s3iface.S3API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

func (a *S3Activities) AbortMultipartUpload(ctx context.Context, input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AbortMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) CompleteMultipartUpload(ctx context.Context, input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CompleteMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) CopyObject(ctx context.Context, input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyObjectWithContext(ctx, input)
}

func (a *S3Activities) CreateBucket(ctx context.Context, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBucketWithContext(ctx, input)
}

func (a *S3Activities) CreateMultipartUpload(ctx context.Context, input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMultipartUploadWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucket(ctx context.Context, input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketAnalyticsConfiguration(ctx context.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketCors(ctx context.Context, input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketEncryption(ctx context.Context, input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketInventoryConfiguration(ctx context.Context, input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketLifecycle(ctx context.Context, input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketMetricsConfiguration(ctx context.Context, input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketPolicy(ctx context.Context, input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketReplication(ctx context.Context, input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketTagging(ctx context.Context, input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) DeleteBucketWebsite(ctx context.Context, input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) DeleteObject(ctx context.Context, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteObjectWithContext(ctx, input)
}

func (a *S3Activities) DeleteObjectTagging(ctx context.Context, input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) DeleteObjects(ctx context.Context, input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteObjectsWithContext(ctx, input)
}

func (a *S3Activities) DeletePublicAccessBlock(ctx context.Context, input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAccelerateConfiguration(ctx context.Context, input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketAccelerateConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAcl(ctx context.Context, input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketAclWithContext(ctx, input)
}

func (a *S3Activities) GetBucketAnalyticsConfiguration(ctx context.Context, input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketCors(ctx context.Context, input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) GetBucketEncryption(ctx context.Context, input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) GetBucketInventoryConfiguration(ctx context.Context, input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLifecycle(ctx context.Context, input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLifecycleConfiguration(ctx context.Context, input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketLifecycleConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLocation(ctx context.Context, input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketLocationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketLogging(ctx context.Context, input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketLoggingWithContext(ctx, input)
}

func (a *S3Activities) GetBucketMetricsConfiguration(ctx context.Context, input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketNotification(ctx context.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketNotificationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketNotificationConfiguration(ctx context.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketNotificationConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketPolicy(ctx context.Context, input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) GetBucketPolicyStatus(ctx context.Context, input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketPolicyStatusWithContext(ctx, input)
}

func (a *S3Activities) GetBucketReplication(ctx context.Context, input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) GetBucketRequestPayment(ctx context.Context, input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketRequestPaymentWithContext(ctx, input)
}

func (a *S3Activities) GetBucketTagging(ctx context.Context, input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) GetBucketVersioning(ctx context.Context, input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketVersioningWithContext(ctx, input)
}

func (a *S3Activities) GetBucketWebsite(ctx context.Context, input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) GetObject(ctx context.Context, input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectWithContext(ctx, input)
}

func (a *S3Activities) GetObjectAcl(ctx context.Context, input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectAclWithContext(ctx, input)
}

func (a *S3Activities) GetObjectLegalHold(ctx context.Context, input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectLegalHoldWithContext(ctx, input)
}

func (a *S3Activities) GetObjectLockConfiguration(ctx context.Context, input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectLockConfigurationWithContext(ctx, input)
}

func (a *S3Activities) GetObjectRetention(ctx context.Context, input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectRetentionWithContext(ctx, input)
}

func (a *S3Activities) GetObjectTagging(ctx context.Context, input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) GetObjectTorrent(ctx context.Context, input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetObjectTorrentWithContext(ctx, input)
}

func (a *S3Activities) GetPublicAccessBlock(ctx context.Context, input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) HeadBucket(ctx context.Context, input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.HeadBucketWithContext(ctx, input)
}

func (a *S3Activities) HeadObject(ctx context.Context, input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.HeadObjectWithContext(ctx, input)
}

func (a *S3Activities) ListBucketAnalyticsConfigurations(ctx context.Context, input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBucketAnalyticsConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBucketInventoryConfigurations(ctx context.Context, input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBucketInventoryConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBucketMetricsConfigurations(ctx context.Context, input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBucketMetricsConfigurationsWithContext(ctx, input)
}

func (a *S3Activities) ListBuckets(ctx context.Context, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBucketsWithContext(ctx, input)
}

func (a *S3Activities) ListMultipartUploads(ctx context.Context, input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMultipartUploadsWithContext(ctx, input)
}

func (a *S3Activities) ListObjectVersions(ctx context.Context, input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListObjectVersionsWithContext(ctx, input)
}

func (a *S3Activities) ListObjects(ctx context.Context, input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListObjectsWithContext(ctx, input)
}

func (a *S3Activities) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListObjectsV2WithContext(ctx, input)
}

func (a *S3Activities) ListParts(ctx context.Context, input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartsWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAccelerateConfiguration(ctx context.Context, input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketAccelerateConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAcl(ctx context.Context, input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketAclWithContext(ctx, input)
}

func (a *S3Activities) PutBucketAnalyticsConfiguration(ctx context.Context, input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketAnalyticsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketCors(ctx context.Context, input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketCorsWithContext(ctx, input)
}

func (a *S3Activities) PutBucketEncryption(ctx context.Context, input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketEncryptionWithContext(ctx, input)
}

func (a *S3Activities) PutBucketInventoryConfiguration(ctx context.Context, input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketInventoryConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLifecycle(ctx context.Context, input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketLifecycleWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLifecycleConfiguration(ctx context.Context, input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketLifecycleConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketLogging(ctx context.Context, input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketLoggingWithContext(ctx, input)
}

func (a *S3Activities) PutBucketMetricsConfiguration(ctx context.Context, input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketMetricsConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketNotification(ctx context.Context, input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketNotificationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketNotificationConfiguration(ctx context.Context, input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketNotificationConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketPolicy(ctx context.Context, input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketPolicyWithContext(ctx, input)
}

func (a *S3Activities) PutBucketReplication(ctx context.Context, input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketReplicationWithContext(ctx, input)
}

func (a *S3Activities) PutBucketRequestPayment(ctx context.Context, input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketRequestPaymentWithContext(ctx, input)
}

func (a *S3Activities) PutBucketTagging(ctx context.Context, input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketTaggingWithContext(ctx, input)
}

func (a *S3Activities) PutBucketVersioning(ctx context.Context, input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketVersioningWithContext(ctx, input)
}

func (a *S3Activities) PutBucketWebsite(ctx context.Context, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBucketWebsiteWithContext(ctx, input)
}

func (a *S3Activities) PutObject(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectWithContext(ctx, input)
}

func (a *S3Activities) PutObjectAcl(ctx context.Context, input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectAclWithContext(ctx, input)
}

func (a *S3Activities) PutObjectLegalHold(ctx context.Context, input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectLegalHoldWithContext(ctx, input)
}

func (a *S3Activities) PutObjectLockConfiguration(ctx context.Context, input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectLockConfigurationWithContext(ctx, input)
}

func (a *S3Activities) PutObjectRetention(ctx context.Context, input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectRetentionWithContext(ctx, input)
}

func (a *S3Activities) PutObjectTagging(ctx context.Context, input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutObjectTaggingWithContext(ctx, input)
}

func (a *S3Activities) PutPublicAccessBlock(ctx context.Context, input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPublicAccessBlockWithContext(ctx, input)
}

func (a *S3Activities) RestoreObject(ctx context.Context, input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreObjectWithContext(ctx, input)
}

func (a *S3Activities) SelectObjectContent(ctx context.Context, input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SelectObjectContentWithContext(ctx, input)
}

func (a *S3Activities) UploadPart(ctx context.Context, input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UploadPartWithContext(ctx, input)
}

func (a *S3Activities) UploadPartCopy(ctx context.Context, input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UploadPartCopyWithContext(ctx, input)
}

func (a *S3Activities) WaitUntilBucketExists(ctx context.Context, input *s3.HeadBucketInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilBucketExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilBucketNotExists(ctx context.Context, input *s3.HeadBucketInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilBucketNotExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilObjectExists(ctx context.Context, input *s3.HeadObjectInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilObjectExistsWithContext(ctx, input, options...)
	})
}

func (a *S3Activities) WaitUntilObjectNotExists(ctx context.Context, input *s3.HeadObjectInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilObjectNotExistsWithContext(ctx, input, options...)
	})
}
