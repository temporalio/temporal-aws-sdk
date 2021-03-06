// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client backupiface.BackupAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := backup.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (backupiface.BackupAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return backup.New(sess), nil
}

func (a *Activities) CreateBackupPlan(ctx context.Context, input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBackupPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBackupSelection(ctx context.Context, input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBackupSelectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBackupVault(ctx context.Context, input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBackupVaultWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackupPlan(ctx context.Context, input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackupSelection(ctx context.Context, input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupSelectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackupVault(ctx context.Context, input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupVaultWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackupVaultAccessPolicy(ctx context.Context, input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupVaultAccessPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackupVaultNotifications(ctx context.Context, input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupVaultNotificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRecoveryPoint(ctx context.Context, input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRecoveryPointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBackupJob(ctx context.Context, input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBackupJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBackupVault(ctx context.Context, input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBackupVaultWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCopyJob(ctx context.Context, input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCopyJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeProtectedResource(ctx context.Context, input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeProtectedResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRecoveryPoint(ctx context.Context, input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRecoveryPointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRegionSettings(ctx context.Context, input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRegionSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRestoreJob(ctx context.Context, input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRestoreJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExportBackupPlanTemplate(ctx context.Context, input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExportBackupPlanTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupPlan(ctx context.Context, input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupPlanFromJSON(ctx context.Context, input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupPlanFromJSONWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupPlanFromTemplate(ctx context.Context, input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupPlanFromTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupSelection(ctx context.Context, input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupSelectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupVaultAccessPolicy(ctx context.Context, input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupVaultAccessPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBackupVaultNotifications(ctx context.Context, input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBackupVaultNotificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRecoveryPointRestoreMetadata(ctx context.Context, input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRecoveryPointRestoreMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSupportedResourceTypes(ctx context.Context, input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSupportedResourceTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupJobs(ctx context.Context, input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupPlanTemplates(ctx context.Context, input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupPlanTemplatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupPlanVersions(ctx context.Context, input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupPlanVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupPlans(ctx context.Context, input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupPlansWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupSelections(ctx context.Context, input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupSelectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBackupVaults(ctx context.Context, input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBackupVaultsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCopyJobs(ctx context.Context, input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCopyJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProtectedResources(ctx context.Context, input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProtectedResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRecoveryPointsByBackupVault(ctx context.Context, input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRecoveryPointsByBackupVaultWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRecoveryPointsByResource(ctx context.Context, input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRecoveryPointsByResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRestoreJobs(ctx context.Context, input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRestoreJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTags(ctx context.Context, input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutBackupVaultAccessPolicy(ctx context.Context, input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutBackupVaultAccessPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutBackupVaultNotifications(ctx context.Context, input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutBackupVaultNotificationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartBackupJob(ctx context.Context, input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartBackupJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartCopyJob(ctx context.Context, input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartCopyJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartRestoreJob(ctx context.Context, input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartRestoreJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopBackupJob(ctx context.Context, input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopBackupJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *backup.TagResourceInput) (*backup.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateBackupPlan(ctx context.Context, input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateBackupPlanWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRecoveryPointLifecycle(ctx context.Context, input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRecoveryPointLifecycleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRegionSettings(ctx context.Context, input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRegionSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
