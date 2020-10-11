// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mturkstub

import (
	"github.com/aws/aws-sdk-go/service/mturk"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type MTurkApproveAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkApproveAssignmentFuture) Get(ctx workflow.Context) (*mturk.ApproveAssignmentOutput, error) {
	var output mturk.ApproveAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkAssociateQualificationWithWorkerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkAssociateQualificationWithWorkerFuture) Get(ctx workflow.Context) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	var output mturk.AssociateQualificationWithWorkerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateAdditionalAssignmentsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateAdditionalAssignmentsForHITFuture) Get(ctx workflow.Context) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	var output mturk.CreateAdditionalAssignmentsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateHITFuture) Get(ctx workflow.Context) (*mturk.CreateHITOutput, error) {
	var output mturk.CreateHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateHITTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateHITTypeFuture) Get(ctx workflow.Context) (*mturk.CreateHITTypeOutput, error) {
	var output mturk.CreateHITTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateHITWithHITTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateHITWithHITTypeFuture) Get(ctx workflow.Context) (*mturk.CreateHITWithHITTypeOutput, error) {
	var output mturk.CreateHITWithHITTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.CreateQualificationTypeOutput, error) {
	var output mturk.CreateQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkCreateWorkerBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkCreateWorkerBlockFuture) Get(ctx workflow.Context) (*mturk.CreateWorkerBlockOutput, error) {
	var output mturk.CreateWorkerBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkDeleteHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkDeleteHITFuture) Get(ctx workflow.Context) (*mturk.DeleteHITOutput, error) {
	var output mturk.DeleteHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkDeleteQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkDeleteQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.DeleteQualificationTypeOutput, error) {
	var output mturk.DeleteQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkDeleteWorkerBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkDeleteWorkerBlockFuture) Get(ctx workflow.Context) (*mturk.DeleteWorkerBlockOutput, error) {
	var output mturk.DeleteWorkerBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkDisassociateQualificationFromWorkerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkDisassociateQualificationFromWorkerFuture) Get(ctx workflow.Context) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	var output mturk.DisassociateQualificationFromWorkerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetAccountBalanceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetAccountBalanceFuture) Get(ctx workflow.Context) (*mturk.GetAccountBalanceOutput, error) {
	var output mturk.GetAccountBalanceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetAssignmentFuture) Get(ctx workflow.Context) (*mturk.GetAssignmentOutput, error) {
	var output mturk.GetAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetFileUploadURLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetFileUploadURLFuture) Get(ctx workflow.Context) (*mturk.GetFileUploadURLOutput, error) {
	var output mturk.GetFileUploadURLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetHITFuture) Get(ctx workflow.Context) (*mturk.GetHITOutput, error) {
	var output mturk.GetHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetQualificationScoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetQualificationScoreFuture) Get(ctx workflow.Context) (*mturk.GetQualificationScoreOutput, error) {
	var output mturk.GetQualificationScoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkGetQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkGetQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.GetQualificationTypeOutput, error) {
	var output mturk.GetQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListAssignmentsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListAssignmentsForHITFuture) Get(ctx workflow.Context) (*mturk.ListAssignmentsForHITOutput, error) {
	var output mturk.ListAssignmentsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListBonusPaymentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListBonusPaymentsFuture) Get(ctx workflow.Context) (*mturk.ListBonusPaymentsOutput, error) {
	var output mturk.ListBonusPaymentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListHITsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListHITsFuture) Get(ctx workflow.Context) (*mturk.ListHITsOutput, error) {
	var output mturk.ListHITsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListHITsForQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListHITsForQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.ListHITsForQualificationTypeOutput, error) {
	var output mturk.ListHITsForQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListQualificationRequestsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListQualificationRequestsFuture) Get(ctx workflow.Context) (*mturk.ListQualificationRequestsOutput, error) {
	var output mturk.ListQualificationRequestsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListQualificationTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListQualificationTypesFuture) Get(ctx workflow.Context) (*mturk.ListQualificationTypesOutput, error) {
	var output mturk.ListQualificationTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListReviewPolicyResultsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListReviewPolicyResultsForHITFuture) Get(ctx workflow.Context) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	var output mturk.ListReviewPolicyResultsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListReviewableHITsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListReviewableHITsFuture) Get(ctx workflow.Context) (*mturk.ListReviewableHITsOutput, error) {
	var output mturk.ListReviewableHITsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListWorkerBlocksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListWorkerBlocksFuture) Get(ctx workflow.Context) (*mturk.ListWorkerBlocksOutput, error) {
	var output mturk.ListWorkerBlocksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkListWorkersWithQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkListWorkersWithQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	var output mturk.ListWorkersWithQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkNotifyWorkersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkNotifyWorkersFuture) Get(ctx workflow.Context) (*mturk.NotifyWorkersOutput, error) {
	var output mturk.NotifyWorkersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkRejectAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkRejectAssignmentFuture) Get(ctx workflow.Context) (*mturk.RejectAssignmentOutput, error) {
	var output mturk.RejectAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkSendBonusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkSendBonusFuture) Get(ctx workflow.Context) (*mturk.SendBonusOutput, error) {
	var output mturk.SendBonusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkSendTestEventNotificationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkSendTestEventNotificationFuture) Get(ctx workflow.Context) (*mturk.SendTestEventNotificationOutput, error) {
	var output mturk.SendTestEventNotificationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkUpdateExpirationForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkUpdateExpirationForHITFuture) Get(ctx workflow.Context) (*mturk.UpdateExpirationForHITOutput, error) {
	var output mturk.UpdateExpirationForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkUpdateHITReviewStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkUpdateHITReviewStatusFuture) Get(ctx workflow.Context) (*mturk.UpdateHITReviewStatusOutput, error) {
	var output mturk.UpdateHITReviewStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkUpdateHITTypeOfHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkUpdateHITTypeOfHITFuture) Get(ctx workflow.Context) (*mturk.UpdateHITTypeOfHITOutput, error) {
	var output mturk.UpdateHITTypeOfHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkUpdateNotificationSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkUpdateNotificationSettingsFuture) Get(ctx workflow.Context) (*mturk.UpdateNotificationSettingsOutput, error) {
	var output mturk.UpdateNotificationSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MTurkUpdateQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MTurkUpdateQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.UpdateQualificationTypeOutput, error) {
	var output mturk.UpdateQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ApproveAssignment(ctx workflow.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
	var output mturk.ApproveAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ApproveAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ApproveAssignmentAsync(ctx workflow.Context, input *mturk.ApproveAssignmentInput) *MTurkApproveAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ApproveAssignment", input)
	return &MTurkApproveAssignmentFuture{Future: future}
}

func (a *stub) AssociateQualificationWithWorker(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	var output mturk.AssociateQualificationWithWorkerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.AssociateQualificationWithWorker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateQualificationWithWorkerAsync(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) *MTurkAssociateQualificationWithWorkerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.AssociateQualificationWithWorker", input)
	return &MTurkAssociateQualificationWithWorkerFuture{Future: future}
}

func (a *stub) CreateAdditionalAssignmentsForHIT(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	var output mturk.CreateAdditionalAssignmentsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateAdditionalAssignmentsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAdditionalAssignmentsForHITAsync(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) *MTurkCreateAdditionalAssignmentsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateAdditionalAssignmentsForHIT", input)
	return &MTurkCreateAdditionalAssignmentsForHITFuture{Future: future}
}

func (a *stub) CreateHIT(ctx workflow.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
	var output mturk.CreateHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITAsync(ctx workflow.Context, input *mturk.CreateHITInput) *MTurkCreateHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHIT", input)
	return &MTurkCreateHITFuture{Future: future}
}

func (a *stub) CreateHITType(ctx workflow.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
	var output mturk.CreateHITTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITTypeInput) *MTurkCreateHITTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITType", input)
	return &MTurkCreateHITTypeFuture{Future: future}
}

func (a *stub) CreateHITWithHITType(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
	var output mturk.CreateHITWithHITTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITWithHITType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITWithHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) *MTurkCreateHITWithHITTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITWithHITType", input)
	return &MTurkCreateHITWithHITTypeFuture{Future: future}
}

func (a *stub) CreateQualificationType(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
	var output mturk.CreateQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateQualificationTypeAsync(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) *MTurkCreateQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateQualificationType", input)
	return &MTurkCreateQualificationTypeFuture{Future: future}
}

func (a *stub) CreateWorkerBlock(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
	var output mturk.CreateWorkerBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateWorkerBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkerBlockAsync(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) *MTurkCreateWorkerBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateWorkerBlock", input)
	return &MTurkCreateWorkerBlockFuture{Future: future}
}

func (a *stub) DeleteHIT(ctx workflow.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
	var output mturk.DeleteHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHITAsync(ctx workflow.Context, input *mturk.DeleteHITInput) *MTurkDeleteHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteHIT", input)
	return &MTurkDeleteHITFuture{Future: future}
}

func (a *stub) DeleteQualificationType(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
	var output mturk.DeleteQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteQualificationTypeAsync(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) *MTurkDeleteQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteQualificationType", input)
	return &MTurkDeleteQualificationTypeFuture{Future: future}
}

func (a *stub) DeleteWorkerBlock(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
	var output mturk.DeleteWorkerBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteWorkerBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkerBlockAsync(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) *MTurkDeleteWorkerBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteWorkerBlock", input)
	return &MTurkDeleteWorkerBlockFuture{Future: future}
}

func (a *stub) DisassociateQualificationFromWorker(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	var output mturk.DisassociateQualificationFromWorkerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DisassociateQualificationFromWorker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateQualificationFromWorkerAsync(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) *MTurkDisassociateQualificationFromWorkerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DisassociateQualificationFromWorker", input)
	return &MTurkDisassociateQualificationFromWorkerFuture{Future: future}
}

func (a *stub) GetAccountBalance(ctx workflow.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
	var output mturk.GetAccountBalanceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetAccountBalance", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAccountBalanceAsync(ctx workflow.Context, input *mturk.GetAccountBalanceInput) *MTurkGetAccountBalanceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetAccountBalance", input)
	return &MTurkGetAccountBalanceFuture{Future: future}
}

func (a *stub) GetAssignment(ctx workflow.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
	var output mturk.GetAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAssignmentAsync(ctx workflow.Context, input *mturk.GetAssignmentInput) *MTurkGetAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetAssignment", input)
	return &MTurkGetAssignmentFuture{Future: future}
}

func (a *stub) GetFileUploadURL(ctx workflow.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
	var output mturk.GetFileUploadURLOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetFileUploadURL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFileUploadURLAsync(ctx workflow.Context, input *mturk.GetFileUploadURLInput) *MTurkGetFileUploadURLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetFileUploadURL", input)
	return &MTurkGetFileUploadURLFuture{Future: future}
}

func (a *stub) GetHIT(ctx workflow.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
	var output mturk.GetHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetHITAsync(ctx workflow.Context, input *mturk.GetHITInput) *MTurkGetHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetHIT", input)
	return &MTurkGetHITFuture{Future: future}
}

func (a *stub) GetQualificationScore(ctx workflow.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
	var output mturk.GetQualificationScoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationScore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQualificationScoreAsync(ctx workflow.Context, input *mturk.GetQualificationScoreInput) *MTurkGetQualificationScoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationScore", input)
	return &MTurkGetQualificationScoreFuture{Future: future}
}

func (a *stub) GetQualificationType(ctx workflow.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
	var output mturk.GetQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQualificationTypeAsync(ctx workflow.Context, input *mturk.GetQualificationTypeInput) *MTurkGetQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationType", input)
	return &MTurkGetQualificationTypeFuture{Future: future}
}

func (a *stub) ListAssignmentsForHIT(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
	var output mturk.ListAssignmentsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListAssignmentsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssignmentsForHITAsync(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) *MTurkListAssignmentsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListAssignmentsForHIT", input)
	return &MTurkListAssignmentsForHITFuture{Future: future}
}

func (a *stub) ListBonusPayments(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
	var output mturk.ListBonusPaymentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListBonusPayments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBonusPaymentsAsync(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) *MTurkListBonusPaymentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListBonusPayments", input)
	return &MTurkListBonusPaymentsFuture{Future: future}
}

func (a *stub) ListHITs(ctx workflow.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
	var output mturk.ListHITsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHITsAsync(ctx workflow.Context, input *mturk.ListHITsInput) *MTurkListHITsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITs", input)
	return &MTurkListHITsFuture{Future: future}
}

func (a *stub) ListHITsForQualificationType(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
	var output mturk.ListHITsForQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITsForQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHITsForQualificationTypeAsync(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) *MTurkListHITsForQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITsForQualificationType", input)
	return &MTurkListHITsForQualificationTypeFuture{Future: future}
}

func (a *stub) ListQualificationRequests(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
	var output mturk.ListQualificationRequestsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationRequests", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQualificationRequestsAsync(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) *MTurkListQualificationRequestsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationRequests", input)
	return &MTurkListQualificationRequestsFuture{Future: future}
}

func (a *stub) ListQualificationTypes(ctx workflow.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
	var output mturk.ListQualificationTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQualificationTypesAsync(ctx workflow.Context, input *mturk.ListQualificationTypesInput) *MTurkListQualificationTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationTypes", input)
	return &MTurkListQualificationTypesFuture{Future: future}
}

func (a *stub) ListReviewPolicyResultsForHIT(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	var output mturk.ListReviewPolicyResultsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewPolicyResultsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListReviewPolicyResultsForHITAsync(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) *MTurkListReviewPolicyResultsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewPolicyResultsForHIT", input)
	return &MTurkListReviewPolicyResultsForHITFuture{Future: future}
}

func (a *stub) ListReviewableHITs(ctx workflow.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
	var output mturk.ListReviewableHITsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewableHITs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListReviewableHITsAsync(ctx workflow.Context, input *mturk.ListReviewableHITsInput) *MTurkListReviewableHITsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewableHITs", input)
	return &MTurkListReviewableHITsFuture{Future: future}
}

func (a *stub) ListWorkerBlocks(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
	var output mturk.ListWorkerBlocksOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkerBlocks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkerBlocksAsync(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) *MTurkListWorkerBlocksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkerBlocks", input)
	return &MTurkListWorkerBlocksFuture{Future: future}
}

func (a *stub) ListWorkersWithQualificationType(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	var output mturk.ListWorkersWithQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkersWithQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkersWithQualificationTypeAsync(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) *MTurkListWorkersWithQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkersWithQualificationType", input)
	return &MTurkListWorkersWithQualificationTypeFuture{Future: future}
}

func (a *stub) NotifyWorkers(ctx workflow.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
	var output mturk.NotifyWorkersOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.NotifyWorkers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) NotifyWorkersAsync(ctx workflow.Context, input *mturk.NotifyWorkersInput) *MTurkNotifyWorkersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.NotifyWorkers", input)
	return &MTurkNotifyWorkersFuture{Future: future}
}

func (a *stub) RejectAssignment(ctx workflow.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
	var output mturk.RejectAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.RejectAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RejectAssignmentAsync(ctx workflow.Context, input *mturk.RejectAssignmentInput) *MTurkRejectAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.RejectAssignment", input)
	return &MTurkRejectAssignmentFuture{Future: future}
}

func (a *stub) SendBonus(ctx workflow.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
	var output mturk.SendBonusOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.SendBonus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendBonusAsync(ctx workflow.Context, input *mturk.SendBonusInput) *MTurkSendBonusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.SendBonus", input)
	return &MTurkSendBonusFuture{Future: future}
}

func (a *stub) SendTestEventNotification(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
	var output mturk.SendTestEventNotificationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.SendTestEventNotification", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTestEventNotificationAsync(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) *MTurkSendTestEventNotificationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.SendTestEventNotification", input)
	return &MTurkSendTestEventNotificationFuture{Future: future}
}

func (a *stub) UpdateExpirationForHIT(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
	var output mturk.UpdateExpirationForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateExpirationForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateExpirationForHITAsync(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) *MTurkUpdateExpirationForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateExpirationForHIT", input)
	return &MTurkUpdateExpirationForHITFuture{Future: future}
}

func (a *stub) UpdateHITReviewStatus(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
	var output mturk.UpdateHITReviewStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITReviewStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateHITReviewStatusAsync(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) *MTurkUpdateHITReviewStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITReviewStatus", input)
	return &MTurkUpdateHITReviewStatusFuture{Future: future}
}

func (a *stub) UpdateHITTypeOfHIT(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
	var output mturk.UpdateHITTypeOfHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITTypeOfHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateHITTypeOfHITAsync(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) *MTurkUpdateHITTypeOfHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITTypeOfHIT", input)
	return &MTurkUpdateHITTypeOfHITFuture{Future: future}
}

func (a *stub) UpdateNotificationSettings(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
	var output mturk.UpdateNotificationSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateNotificationSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateNotificationSettingsAsync(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) *MTurkUpdateNotificationSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateNotificationSettings", input)
	return &MTurkUpdateNotificationSettingsFuture{Future: future}
}

func (a *stub) UpdateQualificationType(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
	var output mturk.UpdateQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateQualificationTypeAsync(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) *MTurkUpdateQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateQualificationType", input)
	return &MTurkUpdateQualificationTypeFuture{Future: future}
}
