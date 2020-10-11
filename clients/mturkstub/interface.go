// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mturkstub

import (
	"github.com/aws/aws-sdk-go/service/mturk"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ApproveAssignment(ctx workflow.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error)
	ApproveAssignmentAsync(ctx workflow.Context, input *mturk.ApproveAssignmentInput) *MTurkApproveAssignmentFuture

	AssociateQualificationWithWorker(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error)
	AssociateQualificationWithWorkerAsync(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) *MTurkAssociateQualificationWithWorkerFuture

	CreateAdditionalAssignmentsForHIT(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error)
	CreateAdditionalAssignmentsForHITAsync(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) *MTurkCreateAdditionalAssignmentsForHITFuture

	CreateHIT(ctx workflow.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error)
	CreateHITAsync(ctx workflow.Context, input *mturk.CreateHITInput) *MTurkCreateHITFuture

	CreateHITType(ctx workflow.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error)
	CreateHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITTypeInput) *MTurkCreateHITTypeFuture

	CreateHITWithHITType(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error)
	CreateHITWithHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) *MTurkCreateHITWithHITTypeFuture

	CreateQualificationType(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error)
	CreateQualificationTypeAsync(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) *MTurkCreateQualificationTypeFuture

	CreateWorkerBlock(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error)
	CreateWorkerBlockAsync(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) *MTurkCreateWorkerBlockFuture

	DeleteHIT(ctx workflow.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error)
	DeleteHITAsync(ctx workflow.Context, input *mturk.DeleteHITInput) *MTurkDeleteHITFuture

	DeleteQualificationType(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error)
	DeleteQualificationTypeAsync(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) *MTurkDeleteQualificationTypeFuture

	DeleteWorkerBlock(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error)
	DeleteWorkerBlockAsync(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) *MTurkDeleteWorkerBlockFuture

	DisassociateQualificationFromWorker(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error)
	DisassociateQualificationFromWorkerAsync(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) *MTurkDisassociateQualificationFromWorkerFuture

	GetAccountBalance(ctx workflow.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error)
	GetAccountBalanceAsync(ctx workflow.Context, input *mturk.GetAccountBalanceInput) *MTurkGetAccountBalanceFuture

	GetAssignment(ctx workflow.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error)
	GetAssignmentAsync(ctx workflow.Context, input *mturk.GetAssignmentInput) *MTurkGetAssignmentFuture

	GetFileUploadURL(ctx workflow.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error)
	GetFileUploadURLAsync(ctx workflow.Context, input *mturk.GetFileUploadURLInput) *MTurkGetFileUploadURLFuture

	GetHIT(ctx workflow.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error)
	GetHITAsync(ctx workflow.Context, input *mturk.GetHITInput) *MTurkGetHITFuture

	GetQualificationScore(ctx workflow.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error)
	GetQualificationScoreAsync(ctx workflow.Context, input *mturk.GetQualificationScoreInput) *MTurkGetQualificationScoreFuture

	GetQualificationType(ctx workflow.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error)
	GetQualificationTypeAsync(ctx workflow.Context, input *mturk.GetQualificationTypeInput) *MTurkGetQualificationTypeFuture

	ListAssignmentsForHIT(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error)
	ListAssignmentsForHITAsync(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) *MTurkListAssignmentsForHITFuture

	ListBonusPayments(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error)
	ListBonusPaymentsAsync(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) *MTurkListBonusPaymentsFuture

	ListHITs(ctx workflow.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error)
	ListHITsAsync(ctx workflow.Context, input *mturk.ListHITsInput) *MTurkListHITsFuture

	ListHITsForQualificationType(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error)
	ListHITsForQualificationTypeAsync(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) *MTurkListHITsForQualificationTypeFuture

	ListQualificationRequests(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error)
	ListQualificationRequestsAsync(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) *MTurkListQualificationRequestsFuture

	ListQualificationTypes(ctx workflow.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error)
	ListQualificationTypesAsync(ctx workflow.Context, input *mturk.ListQualificationTypesInput) *MTurkListQualificationTypesFuture

	ListReviewPolicyResultsForHIT(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error)
	ListReviewPolicyResultsForHITAsync(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) *MTurkListReviewPolicyResultsForHITFuture

	ListReviewableHITs(ctx workflow.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error)
	ListReviewableHITsAsync(ctx workflow.Context, input *mturk.ListReviewableHITsInput) *MTurkListReviewableHITsFuture

	ListWorkerBlocks(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error)
	ListWorkerBlocksAsync(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) *MTurkListWorkerBlocksFuture

	ListWorkersWithQualificationType(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error)
	ListWorkersWithQualificationTypeAsync(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) *MTurkListWorkersWithQualificationTypeFuture

	NotifyWorkers(ctx workflow.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error)
	NotifyWorkersAsync(ctx workflow.Context, input *mturk.NotifyWorkersInput) *MTurkNotifyWorkersFuture

	RejectAssignment(ctx workflow.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error)
	RejectAssignmentAsync(ctx workflow.Context, input *mturk.RejectAssignmentInput) *MTurkRejectAssignmentFuture

	SendBonus(ctx workflow.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error)
	SendBonusAsync(ctx workflow.Context, input *mturk.SendBonusInput) *MTurkSendBonusFuture

	SendTestEventNotification(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error)
	SendTestEventNotificationAsync(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) *MTurkSendTestEventNotificationFuture

	UpdateExpirationForHIT(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error)
	UpdateExpirationForHITAsync(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) *MTurkUpdateExpirationForHITFuture

	UpdateHITReviewStatus(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error)
	UpdateHITReviewStatusAsync(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) *MTurkUpdateHITReviewStatusFuture

	UpdateHITTypeOfHIT(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error)
	UpdateHITTypeOfHITAsync(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) *MTurkUpdateHITTypeOfHITFuture

	UpdateNotificationSettings(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error)
	UpdateNotificationSettingsAsync(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) *MTurkUpdateNotificationSettingsFuture

	UpdateQualificationType(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error)
	UpdateQualificationTypeAsync(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) *MTurkUpdateQualificationTypeFuture
}

func NewClient() Client {
	return &stub{}
}
