// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mturkstub

import (
	"github.com/aws/aws-sdk-go/service/mturk"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ApproveAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApproveAssignmentFuture) Get(ctx workflow.Context) (*mturk.ApproveAssignmentOutput, error) {
	var output mturk.ApproveAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateQualificationWithWorkerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateQualificationWithWorkerFuture) Get(ctx workflow.Context) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	var output mturk.AssociateQualificationWithWorkerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateAdditionalAssignmentsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateAdditionalAssignmentsForHITFuture) Get(ctx workflow.Context) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	var output mturk.CreateAdditionalAssignmentsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHITFuture) Get(ctx workflow.Context) (*mturk.CreateHITOutput, error) {
	var output mturk.CreateHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHITTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHITTypeFuture) Get(ctx workflow.Context) (*mturk.CreateHITTypeOutput, error) {
	var output mturk.CreateHITTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHITWithHITTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHITWithHITTypeFuture) Get(ctx workflow.Context) (*mturk.CreateHITWithHITTypeOutput, error) {
	var output mturk.CreateHITWithHITTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.CreateQualificationTypeOutput, error) {
	var output mturk.CreateQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateWorkerBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateWorkerBlockFuture) Get(ctx workflow.Context) (*mturk.CreateWorkerBlockOutput, error) {
	var output mturk.CreateWorkerBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteHITFuture) Get(ctx workflow.Context) (*mturk.DeleteHITOutput, error) {
	var output mturk.DeleteHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.DeleteQualificationTypeOutput, error) {
	var output mturk.DeleteQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteWorkerBlockFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteWorkerBlockFuture) Get(ctx workflow.Context) (*mturk.DeleteWorkerBlockOutput, error) {
	var output mturk.DeleteWorkerBlockOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateQualificationFromWorkerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateQualificationFromWorkerFuture) Get(ctx workflow.Context) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	var output mturk.DisassociateQualificationFromWorkerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAccountBalanceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAccountBalanceFuture) Get(ctx workflow.Context) (*mturk.GetAccountBalanceOutput, error) {
	var output mturk.GetAccountBalanceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAssignmentFuture) Get(ctx workflow.Context) (*mturk.GetAssignmentOutput, error) {
	var output mturk.GetAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetFileUploadURLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetFileUploadURLFuture) Get(ctx workflow.Context) (*mturk.GetFileUploadURLOutput, error) {
	var output mturk.GetFileUploadURLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetHITFuture) Get(ctx workflow.Context) (*mturk.GetHITOutput, error) {
	var output mturk.GetHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetQualificationScoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetQualificationScoreFuture) Get(ctx workflow.Context) (*mturk.GetQualificationScoreOutput, error) {
	var output mturk.GetQualificationScoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.GetQualificationTypeOutput, error) {
	var output mturk.GetQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAssignmentsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAssignmentsForHITFuture) Get(ctx workflow.Context) (*mturk.ListAssignmentsForHITOutput, error) {
	var output mturk.ListAssignmentsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBonusPaymentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBonusPaymentsFuture) Get(ctx workflow.Context) (*mturk.ListBonusPaymentsOutput, error) {
	var output mturk.ListBonusPaymentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHITsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHITsFuture) Get(ctx workflow.Context) (*mturk.ListHITsOutput, error) {
	var output mturk.ListHITsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHITsForQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHITsForQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.ListHITsForQualificationTypeOutput, error) {
	var output mturk.ListHITsForQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListQualificationRequestsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListQualificationRequestsFuture) Get(ctx workflow.Context) (*mturk.ListQualificationRequestsOutput, error) {
	var output mturk.ListQualificationRequestsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListQualificationTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListQualificationTypesFuture) Get(ctx workflow.Context) (*mturk.ListQualificationTypesOutput, error) {
	var output mturk.ListQualificationTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListReviewPolicyResultsForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListReviewPolicyResultsForHITFuture) Get(ctx workflow.Context) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	var output mturk.ListReviewPolicyResultsForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListReviewableHITsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListReviewableHITsFuture) Get(ctx workflow.Context) (*mturk.ListReviewableHITsOutput, error) {
	var output mturk.ListReviewableHITsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWorkerBlocksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWorkerBlocksFuture) Get(ctx workflow.Context) (*mturk.ListWorkerBlocksOutput, error) {
	var output mturk.ListWorkerBlocksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWorkersWithQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWorkersWithQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	var output mturk.ListWorkersWithQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NotifyWorkersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *NotifyWorkersFuture) Get(ctx workflow.Context) (*mturk.NotifyWorkersOutput, error) {
	var output mturk.NotifyWorkersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RejectAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RejectAssignmentFuture) Get(ctx workflow.Context) (*mturk.RejectAssignmentOutput, error) {
	var output mturk.RejectAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendBonusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendBonusFuture) Get(ctx workflow.Context) (*mturk.SendBonusOutput, error) {
	var output mturk.SendBonusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendTestEventNotificationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendTestEventNotificationFuture) Get(ctx workflow.Context) (*mturk.SendTestEventNotificationOutput, error) {
	var output mturk.SendTestEventNotificationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateExpirationForHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateExpirationForHITFuture) Get(ctx workflow.Context) (*mturk.UpdateExpirationForHITOutput, error) {
	var output mturk.UpdateExpirationForHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateHITReviewStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateHITReviewStatusFuture) Get(ctx workflow.Context) (*mturk.UpdateHITReviewStatusOutput, error) {
	var output mturk.UpdateHITReviewStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateHITTypeOfHITFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateHITTypeOfHITFuture) Get(ctx workflow.Context) (*mturk.UpdateHITTypeOfHITOutput, error) {
	var output mturk.UpdateHITTypeOfHITOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateNotificationSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateNotificationSettingsFuture) Get(ctx workflow.Context) (*mturk.UpdateNotificationSettingsOutput, error) {
	var output mturk.UpdateNotificationSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateQualificationTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateQualificationTypeFuture) Get(ctx workflow.Context) (*mturk.UpdateQualificationTypeOutput, error) {
	var output mturk.UpdateQualificationTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ApproveAssignment(ctx workflow.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
	var output mturk.ApproveAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ApproveAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ApproveAssignmentAsync(ctx workflow.Context, input *mturk.ApproveAssignmentInput) *ApproveAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ApproveAssignment", input)
	return &ApproveAssignmentFuture{Future: future}
}

func (a *stub) AssociateQualificationWithWorker(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	var output mturk.AssociateQualificationWithWorkerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.AssociateQualificationWithWorker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateQualificationWithWorkerAsync(ctx workflow.Context, input *mturk.AssociateQualificationWithWorkerInput) *AssociateQualificationWithWorkerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.AssociateQualificationWithWorker", input)
	return &AssociateQualificationWithWorkerFuture{Future: future}
}

func (a *stub) CreateAdditionalAssignmentsForHIT(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	var output mturk.CreateAdditionalAssignmentsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateAdditionalAssignmentsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAdditionalAssignmentsForHITAsync(ctx workflow.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) *CreateAdditionalAssignmentsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateAdditionalAssignmentsForHIT", input)
	return &CreateAdditionalAssignmentsForHITFuture{Future: future}
}

func (a *stub) CreateHIT(ctx workflow.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
	var output mturk.CreateHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITAsync(ctx workflow.Context, input *mturk.CreateHITInput) *CreateHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHIT", input)
	return &CreateHITFuture{Future: future}
}

func (a *stub) CreateHITType(ctx workflow.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
	var output mturk.CreateHITTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITTypeInput) *CreateHITTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITType", input)
	return &CreateHITTypeFuture{Future: future}
}

func (a *stub) CreateHITWithHITType(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
	var output mturk.CreateHITWithHITTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITWithHITType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHITWithHITTypeAsync(ctx workflow.Context, input *mturk.CreateHITWithHITTypeInput) *CreateHITWithHITTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateHITWithHITType", input)
	return &CreateHITWithHITTypeFuture{Future: future}
}

func (a *stub) CreateQualificationType(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
	var output mturk.CreateQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateQualificationTypeAsync(ctx workflow.Context, input *mturk.CreateQualificationTypeInput) *CreateQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateQualificationType", input)
	return &CreateQualificationTypeFuture{Future: future}
}

func (a *stub) CreateWorkerBlock(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
	var output mturk.CreateWorkerBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.CreateWorkerBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkerBlockAsync(ctx workflow.Context, input *mturk.CreateWorkerBlockInput) *CreateWorkerBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.CreateWorkerBlock", input)
	return &CreateWorkerBlockFuture{Future: future}
}

func (a *stub) DeleteHIT(ctx workflow.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
	var output mturk.DeleteHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHITAsync(ctx workflow.Context, input *mturk.DeleteHITInput) *DeleteHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteHIT", input)
	return &DeleteHITFuture{Future: future}
}

func (a *stub) DeleteQualificationType(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
	var output mturk.DeleteQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteQualificationTypeAsync(ctx workflow.Context, input *mturk.DeleteQualificationTypeInput) *DeleteQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteQualificationType", input)
	return &DeleteQualificationTypeFuture{Future: future}
}

func (a *stub) DeleteWorkerBlock(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
	var output mturk.DeleteWorkerBlockOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteWorkerBlock", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkerBlockAsync(ctx workflow.Context, input *mturk.DeleteWorkerBlockInput) *DeleteWorkerBlockFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DeleteWorkerBlock", input)
	return &DeleteWorkerBlockFuture{Future: future}
}

func (a *stub) DisassociateQualificationFromWorker(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	var output mturk.DisassociateQualificationFromWorkerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.DisassociateQualificationFromWorker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateQualificationFromWorkerAsync(ctx workflow.Context, input *mturk.DisassociateQualificationFromWorkerInput) *DisassociateQualificationFromWorkerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.DisassociateQualificationFromWorker", input)
	return &DisassociateQualificationFromWorkerFuture{Future: future}
}

func (a *stub) GetAccountBalance(ctx workflow.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
	var output mturk.GetAccountBalanceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetAccountBalance", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAccountBalanceAsync(ctx workflow.Context, input *mturk.GetAccountBalanceInput) *GetAccountBalanceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetAccountBalance", input)
	return &GetAccountBalanceFuture{Future: future}
}

func (a *stub) GetAssignment(ctx workflow.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
	var output mturk.GetAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAssignmentAsync(ctx workflow.Context, input *mturk.GetAssignmentInput) *GetAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetAssignment", input)
	return &GetAssignmentFuture{Future: future}
}

func (a *stub) GetFileUploadURL(ctx workflow.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
	var output mturk.GetFileUploadURLOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetFileUploadURL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFileUploadURLAsync(ctx workflow.Context, input *mturk.GetFileUploadURLInput) *GetFileUploadURLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetFileUploadURL", input)
	return &GetFileUploadURLFuture{Future: future}
}

func (a *stub) GetHIT(ctx workflow.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
	var output mturk.GetHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetHITAsync(ctx workflow.Context, input *mturk.GetHITInput) *GetHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetHIT", input)
	return &GetHITFuture{Future: future}
}

func (a *stub) GetQualificationScore(ctx workflow.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
	var output mturk.GetQualificationScoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationScore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQualificationScoreAsync(ctx workflow.Context, input *mturk.GetQualificationScoreInput) *GetQualificationScoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationScore", input)
	return &GetQualificationScoreFuture{Future: future}
}

func (a *stub) GetQualificationType(ctx workflow.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
	var output mturk.GetQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQualificationTypeAsync(ctx workflow.Context, input *mturk.GetQualificationTypeInput) *GetQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.GetQualificationType", input)
	return &GetQualificationTypeFuture{Future: future}
}

func (a *stub) ListAssignmentsForHIT(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
	var output mturk.ListAssignmentsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListAssignmentsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssignmentsForHITAsync(ctx workflow.Context, input *mturk.ListAssignmentsForHITInput) *ListAssignmentsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListAssignmentsForHIT", input)
	return &ListAssignmentsForHITFuture{Future: future}
}

func (a *stub) ListBonusPayments(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
	var output mturk.ListBonusPaymentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListBonusPayments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBonusPaymentsAsync(ctx workflow.Context, input *mturk.ListBonusPaymentsInput) *ListBonusPaymentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListBonusPayments", input)
	return &ListBonusPaymentsFuture{Future: future}
}

func (a *stub) ListHITs(ctx workflow.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
	var output mturk.ListHITsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHITsAsync(ctx workflow.Context, input *mturk.ListHITsInput) *ListHITsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITs", input)
	return &ListHITsFuture{Future: future}
}

func (a *stub) ListHITsForQualificationType(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
	var output mturk.ListHITsForQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITsForQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHITsForQualificationTypeAsync(ctx workflow.Context, input *mturk.ListHITsForQualificationTypeInput) *ListHITsForQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListHITsForQualificationType", input)
	return &ListHITsForQualificationTypeFuture{Future: future}
}

func (a *stub) ListQualificationRequests(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
	var output mturk.ListQualificationRequestsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationRequests", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQualificationRequestsAsync(ctx workflow.Context, input *mturk.ListQualificationRequestsInput) *ListQualificationRequestsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationRequests", input)
	return &ListQualificationRequestsFuture{Future: future}
}

func (a *stub) ListQualificationTypes(ctx workflow.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
	var output mturk.ListQualificationTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQualificationTypesAsync(ctx workflow.Context, input *mturk.ListQualificationTypesInput) *ListQualificationTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListQualificationTypes", input)
	return &ListQualificationTypesFuture{Future: future}
}

func (a *stub) ListReviewPolicyResultsForHIT(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	var output mturk.ListReviewPolicyResultsForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewPolicyResultsForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListReviewPolicyResultsForHITAsync(ctx workflow.Context, input *mturk.ListReviewPolicyResultsForHITInput) *ListReviewPolicyResultsForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewPolicyResultsForHIT", input)
	return &ListReviewPolicyResultsForHITFuture{Future: future}
}

func (a *stub) ListReviewableHITs(ctx workflow.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
	var output mturk.ListReviewableHITsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewableHITs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListReviewableHITsAsync(ctx workflow.Context, input *mturk.ListReviewableHITsInput) *ListReviewableHITsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListReviewableHITs", input)
	return &ListReviewableHITsFuture{Future: future}
}

func (a *stub) ListWorkerBlocks(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
	var output mturk.ListWorkerBlocksOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkerBlocks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkerBlocksAsync(ctx workflow.Context, input *mturk.ListWorkerBlocksInput) *ListWorkerBlocksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkerBlocks", input)
	return &ListWorkerBlocksFuture{Future: future}
}

func (a *stub) ListWorkersWithQualificationType(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	var output mturk.ListWorkersWithQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkersWithQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkersWithQualificationTypeAsync(ctx workflow.Context, input *mturk.ListWorkersWithQualificationTypeInput) *ListWorkersWithQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.ListWorkersWithQualificationType", input)
	return &ListWorkersWithQualificationTypeFuture{Future: future}
}

func (a *stub) NotifyWorkers(ctx workflow.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
	var output mturk.NotifyWorkersOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.NotifyWorkers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) NotifyWorkersAsync(ctx workflow.Context, input *mturk.NotifyWorkersInput) *NotifyWorkersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.NotifyWorkers", input)
	return &NotifyWorkersFuture{Future: future}
}

func (a *stub) RejectAssignment(ctx workflow.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
	var output mturk.RejectAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.RejectAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RejectAssignmentAsync(ctx workflow.Context, input *mturk.RejectAssignmentInput) *RejectAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.RejectAssignment", input)
	return &RejectAssignmentFuture{Future: future}
}

func (a *stub) SendBonus(ctx workflow.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
	var output mturk.SendBonusOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.SendBonus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendBonusAsync(ctx workflow.Context, input *mturk.SendBonusInput) *SendBonusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.SendBonus", input)
	return &SendBonusFuture{Future: future}
}

func (a *stub) SendTestEventNotification(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
	var output mturk.SendTestEventNotificationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.SendTestEventNotification", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendTestEventNotificationAsync(ctx workflow.Context, input *mturk.SendTestEventNotificationInput) *SendTestEventNotificationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.SendTestEventNotification", input)
	return &SendTestEventNotificationFuture{Future: future}
}

func (a *stub) UpdateExpirationForHIT(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
	var output mturk.UpdateExpirationForHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateExpirationForHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateExpirationForHITAsync(ctx workflow.Context, input *mturk.UpdateExpirationForHITInput) *UpdateExpirationForHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateExpirationForHIT", input)
	return &UpdateExpirationForHITFuture{Future: future}
}

func (a *stub) UpdateHITReviewStatus(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
	var output mturk.UpdateHITReviewStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITReviewStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateHITReviewStatusAsync(ctx workflow.Context, input *mturk.UpdateHITReviewStatusInput) *UpdateHITReviewStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITReviewStatus", input)
	return &UpdateHITReviewStatusFuture{Future: future}
}

func (a *stub) UpdateHITTypeOfHIT(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
	var output mturk.UpdateHITTypeOfHITOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITTypeOfHIT", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateHITTypeOfHITAsync(ctx workflow.Context, input *mturk.UpdateHITTypeOfHITInput) *UpdateHITTypeOfHITFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateHITTypeOfHIT", input)
	return &UpdateHITTypeOfHITFuture{Future: future}
}

func (a *stub) UpdateNotificationSettings(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
	var output mturk.UpdateNotificationSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateNotificationSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateNotificationSettingsAsync(ctx workflow.Context, input *mturk.UpdateNotificationSettingsInput) *UpdateNotificationSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateNotificationSettings", input)
	return &UpdateNotificationSettingsFuture{Future: future}
}

func (a *stub) UpdateQualificationType(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
	var output mturk.UpdateQualificationTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateQualificationType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateQualificationTypeAsync(ctx workflow.Context, input *mturk.UpdateQualificationTypeInput) *UpdateQualificationTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mturk.UpdateQualificationType", input)
	return &UpdateQualificationTypeFuture{Future: future}
}
