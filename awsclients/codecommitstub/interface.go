// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package codecommitstub

import (
	"github.com/aws/aws-sdk-go/service/codecommit"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssociateApprovalRuleTemplateWithRepository(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) (*codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput, error)
	AssociateApprovalRuleTemplateWithRepositoryAsync(ctx workflow.Context, input *codecommit.AssociateApprovalRuleTemplateWithRepositoryInput) *CodeCommitAssociateApprovalRuleTemplateWithRepositoryFuture

	BatchAssociateApprovalRuleTemplateWithRepositories(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) (*codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput, error)
	BatchAssociateApprovalRuleTemplateWithRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput) *CodeCommitBatchAssociateApprovalRuleTemplateWithRepositoriesFuture

	BatchDescribeMergeConflicts(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) (*codecommit.BatchDescribeMergeConflictsOutput, error)
	BatchDescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.BatchDescribeMergeConflictsInput) *CodeCommitBatchDescribeMergeConflictsFuture

	BatchDisassociateApprovalRuleTemplateFromRepositories(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) (*codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error)
	BatchDisassociateApprovalRuleTemplateFromRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) *CodeCommitBatchDisassociateApprovalRuleTemplateFromRepositoriesFuture

	BatchGetCommits(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) (*codecommit.BatchGetCommitsOutput, error)
	BatchGetCommitsAsync(ctx workflow.Context, input *codecommit.BatchGetCommitsInput) *CodeCommitBatchGetCommitsFuture

	BatchGetRepositories(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error)
	BatchGetRepositoriesAsync(ctx workflow.Context, input *codecommit.BatchGetRepositoriesInput) *CodeCommitBatchGetRepositoriesFuture

	CreateApprovalRuleTemplate(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) (*codecommit.CreateApprovalRuleTemplateOutput, error)
	CreateApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.CreateApprovalRuleTemplateInput) *CodeCommitCreateApprovalRuleTemplateFuture

	CreateBranch(ctx workflow.Context, input *codecommit.CreateBranchInput) (*codecommit.CreateBranchOutput, error)
	CreateBranchAsync(ctx workflow.Context, input *codecommit.CreateBranchInput) *CodeCommitCreateBranchFuture

	CreateCommit(ctx workflow.Context, input *codecommit.CreateCommitInput) (*codecommit.CreateCommitOutput, error)
	CreateCommitAsync(ctx workflow.Context, input *codecommit.CreateCommitInput) *CodeCommitCreateCommitFuture

	CreatePullRequestApprovalRule(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) (*codecommit.CreatePullRequestApprovalRuleOutput, error)
	CreatePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.CreatePullRequestApprovalRuleInput) *CodeCommitCreatePullRequestApprovalRuleFuture

	CreateRepository(ctx workflow.Context, input *codecommit.CreateRepositoryInput) (*codecommit.CreateRepositoryOutput, error)
	CreateRepositoryAsync(ctx workflow.Context, input *codecommit.CreateRepositoryInput) *CodeCommitCreateRepositoryFuture

	CreateUnreferencedMergeCommit(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) (*codecommit.CreateUnreferencedMergeCommitOutput, error)
	CreateUnreferencedMergeCommitAsync(ctx workflow.Context, input *codecommit.CreateUnreferencedMergeCommitInput) *CodeCommitCreateUnreferencedMergeCommitFuture

	DeleteApprovalRuleTemplate(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) (*codecommit.DeleteApprovalRuleTemplateOutput, error)
	DeleteApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.DeleteApprovalRuleTemplateInput) *CodeCommitDeleteApprovalRuleTemplateFuture

	DeleteBranch(ctx workflow.Context, input *codecommit.DeleteBranchInput) (*codecommit.DeleteBranchOutput, error)
	DeleteBranchAsync(ctx workflow.Context, input *codecommit.DeleteBranchInput) *CodeCommitDeleteBranchFuture

	DeleteCommentContent(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) (*codecommit.DeleteCommentContentOutput, error)
	DeleteCommentContentAsync(ctx workflow.Context, input *codecommit.DeleteCommentContentInput) *CodeCommitDeleteCommentContentFuture

	DeleteFile(ctx workflow.Context, input *codecommit.DeleteFileInput) (*codecommit.DeleteFileOutput, error)
	DeleteFileAsync(ctx workflow.Context, input *codecommit.DeleteFileInput) *CodeCommitDeleteFileFuture

	DeletePullRequestApprovalRule(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) (*codecommit.DeletePullRequestApprovalRuleOutput, error)
	DeletePullRequestApprovalRuleAsync(ctx workflow.Context, input *codecommit.DeletePullRequestApprovalRuleInput) *CodeCommitDeletePullRequestApprovalRuleFuture

	DeleteRepository(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) (*codecommit.DeleteRepositoryOutput, error)
	DeleteRepositoryAsync(ctx workflow.Context, input *codecommit.DeleteRepositoryInput) *CodeCommitDeleteRepositoryFuture

	DescribeMergeConflicts(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) (*codecommit.DescribeMergeConflictsOutput, error)
	DescribeMergeConflictsAsync(ctx workflow.Context, input *codecommit.DescribeMergeConflictsInput) *CodeCommitDescribeMergeConflictsFuture

	DescribePullRequestEvents(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) (*codecommit.DescribePullRequestEventsOutput, error)
	DescribePullRequestEventsAsync(ctx workflow.Context, input *codecommit.DescribePullRequestEventsInput) *CodeCommitDescribePullRequestEventsFuture

	DisassociateApprovalRuleTemplateFromRepository(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) (*codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput, error)
	DisassociateApprovalRuleTemplateFromRepositoryAsync(ctx workflow.Context, input *codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput) *CodeCommitDisassociateApprovalRuleTemplateFromRepositoryFuture

	EvaluatePullRequestApprovalRules(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) (*codecommit.EvaluatePullRequestApprovalRulesOutput, error)
	EvaluatePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.EvaluatePullRequestApprovalRulesInput) *CodeCommitEvaluatePullRequestApprovalRulesFuture

	GetApprovalRuleTemplate(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) (*codecommit.GetApprovalRuleTemplateOutput, error)
	GetApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.GetApprovalRuleTemplateInput) *CodeCommitGetApprovalRuleTemplateFuture

	GetBlob(ctx workflow.Context, input *codecommit.GetBlobInput) (*codecommit.GetBlobOutput, error)
	GetBlobAsync(ctx workflow.Context, input *codecommit.GetBlobInput) *CodeCommitGetBlobFuture

	GetBranch(ctx workflow.Context, input *codecommit.GetBranchInput) (*codecommit.GetBranchOutput, error)
	GetBranchAsync(ctx workflow.Context, input *codecommit.GetBranchInput) *CodeCommitGetBranchFuture

	GetComment(ctx workflow.Context, input *codecommit.GetCommentInput) (*codecommit.GetCommentOutput, error)
	GetCommentAsync(ctx workflow.Context, input *codecommit.GetCommentInput) *CodeCommitGetCommentFuture

	GetCommentReactions(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) (*codecommit.GetCommentReactionsOutput, error)
	GetCommentReactionsAsync(ctx workflow.Context, input *codecommit.GetCommentReactionsInput) *CodeCommitGetCommentReactionsFuture

	GetCommentsForComparedCommit(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) (*codecommit.GetCommentsForComparedCommitOutput, error)
	GetCommentsForComparedCommitAsync(ctx workflow.Context, input *codecommit.GetCommentsForComparedCommitInput) *CodeCommitGetCommentsForComparedCommitFuture

	GetCommit(ctx workflow.Context, input *codecommit.GetCommitInput) (*codecommit.GetCommitOutput, error)
	GetCommitAsync(ctx workflow.Context, input *codecommit.GetCommitInput) *CodeCommitGetCommitFuture

	GetDifferences(ctx workflow.Context, input *codecommit.GetDifferencesInput) (*codecommit.GetDifferencesOutput, error)
	GetDifferencesAsync(ctx workflow.Context, input *codecommit.GetDifferencesInput) *CodeCommitGetDifferencesFuture

	GetFile(ctx workflow.Context, input *codecommit.GetFileInput) (*codecommit.GetFileOutput, error)
	GetFileAsync(ctx workflow.Context, input *codecommit.GetFileInput) *CodeCommitGetFileFuture

	GetFolder(ctx workflow.Context, input *codecommit.GetFolderInput) (*codecommit.GetFolderOutput, error)
	GetFolderAsync(ctx workflow.Context, input *codecommit.GetFolderInput) *CodeCommitGetFolderFuture

	GetMergeCommit(ctx workflow.Context, input *codecommit.GetMergeCommitInput) (*codecommit.GetMergeCommitOutput, error)
	GetMergeCommitAsync(ctx workflow.Context, input *codecommit.GetMergeCommitInput) *CodeCommitGetMergeCommitFuture

	GetMergeConflicts(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) (*codecommit.GetMergeConflictsOutput, error)
	GetMergeConflictsAsync(ctx workflow.Context, input *codecommit.GetMergeConflictsInput) *CodeCommitGetMergeConflictsFuture

	GetMergeOptions(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) (*codecommit.GetMergeOptionsOutput, error)
	GetMergeOptionsAsync(ctx workflow.Context, input *codecommit.GetMergeOptionsInput) *CodeCommitGetMergeOptionsFuture

	GetPullRequestApprovalStates(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) (*codecommit.GetPullRequestApprovalStatesOutput, error)
	GetPullRequestApprovalStatesAsync(ctx workflow.Context, input *codecommit.GetPullRequestApprovalStatesInput) *CodeCommitGetPullRequestApprovalStatesFuture

	GetPullRequestOverrideState(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) (*codecommit.GetPullRequestOverrideStateOutput, error)
	GetPullRequestOverrideStateAsync(ctx workflow.Context, input *codecommit.GetPullRequestOverrideStateInput) *CodeCommitGetPullRequestOverrideStateFuture

	GetRepository(ctx workflow.Context, input *codecommit.GetRepositoryInput) (*codecommit.GetRepositoryOutput, error)
	GetRepositoryAsync(ctx workflow.Context, input *codecommit.GetRepositoryInput) *CodeCommitGetRepositoryFuture

	GetRepositoryTriggers(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) (*codecommit.GetRepositoryTriggersOutput, error)
	GetRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.GetRepositoryTriggersInput) *CodeCommitGetRepositoryTriggersFuture

	ListApprovalRuleTemplates(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) (*codecommit.ListApprovalRuleTemplatesOutput, error)
	ListApprovalRuleTemplatesAsync(ctx workflow.Context, input *codecommit.ListApprovalRuleTemplatesInput) *CodeCommitListApprovalRuleTemplatesFuture

	ListAssociatedApprovalRuleTemplatesForRepository(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) (*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error)
	ListAssociatedApprovalRuleTemplatesForRepositoryAsync(ctx workflow.Context, input *codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput) *CodeCommitListAssociatedApprovalRuleTemplatesForRepositoryFuture

	ListBranches(ctx workflow.Context, input *codecommit.ListBranchesInput) (*codecommit.ListBranchesOutput, error)
	ListBranchesAsync(ctx workflow.Context, input *codecommit.ListBranchesInput) *CodeCommitListBranchesFuture

	ListPullRequests(ctx workflow.Context, input *codecommit.ListPullRequestsInput) (*codecommit.ListPullRequestsOutput, error)
	ListPullRequestsAsync(ctx workflow.Context, input *codecommit.ListPullRequestsInput) *CodeCommitListPullRequestsFuture

	ListRepositories(ctx workflow.Context, input *codecommit.ListRepositoriesInput) (*codecommit.ListRepositoriesOutput, error)
	ListRepositoriesAsync(ctx workflow.Context, input *codecommit.ListRepositoriesInput) *CodeCommitListRepositoriesFuture

	ListRepositoriesForApprovalRuleTemplate(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) (*codecommit.ListRepositoriesForApprovalRuleTemplateOutput, error)
	ListRepositoriesForApprovalRuleTemplateAsync(ctx workflow.Context, input *codecommit.ListRepositoriesForApprovalRuleTemplateInput) *CodeCommitListRepositoriesForApprovalRuleTemplateFuture

	ListTagsForResource(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) (*codecommit.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codecommit.ListTagsForResourceInput) *CodeCommitListTagsForResourceFuture

	MergeBranchesByFastForward(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) (*codecommit.MergeBranchesByFastForwardOutput, error)
	MergeBranchesByFastForwardAsync(ctx workflow.Context, input *codecommit.MergeBranchesByFastForwardInput) *CodeCommitMergeBranchesByFastForwardFuture

	MergeBranchesBySquash(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) (*codecommit.MergeBranchesBySquashOutput, error)
	MergeBranchesBySquashAsync(ctx workflow.Context, input *codecommit.MergeBranchesBySquashInput) *CodeCommitMergeBranchesBySquashFuture

	MergeBranchesByThreeWay(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) (*codecommit.MergeBranchesByThreeWayOutput, error)
	MergeBranchesByThreeWayAsync(ctx workflow.Context, input *codecommit.MergeBranchesByThreeWayInput) *CodeCommitMergeBranchesByThreeWayFuture

	MergePullRequestByFastForward(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) (*codecommit.MergePullRequestByFastForwardOutput, error)
	MergePullRequestByFastForwardAsync(ctx workflow.Context, input *codecommit.MergePullRequestByFastForwardInput) *CodeCommitMergePullRequestByFastForwardFuture

	MergePullRequestBySquash(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) (*codecommit.MergePullRequestBySquashOutput, error)
	MergePullRequestBySquashAsync(ctx workflow.Context, input *codecommit.MergePullRequestBySquashInput) *CodeCommitMergePullRequestBySquashFuture

	MergePullRequestByThreeWay(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) (*codecommit.MergePullRequestByThreeWayOutput, error)
	MergePullRequestByThreeWayAsync(ctx workflow.Context, input *codecommit.MergePullRequestByThreeWayInput) *CodeCommitMergePullRequestByThreeWayFuture

	OverridePullRequestApprovalRules(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) (*codecommit.OverridePullRequestApprovalRulesOutput, error)
	OverridePullRequestApprovalRulesAsync(ctx workflow.Context, input *codecommit.OverridePullRequestApprovalRulesInput) *CodeCommitOverridePullRequestApprovalRulesFuture

	PostCommentForComparedCommit(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) (*codecommit.PostCommentForComparedCommitOutput, error)
	PostCommentForComparedCommitAsync(ctx workflow.Context, input *codecommit.PostCommentForComparedCommitInput) *CodeCommitPostCommentForComparedCommitFuture

	PostCommentReply(ctx workflow.Context, input *codecommit.PostCommentReplyInput) (*codecommit.PostCommentReplyOutput, error)
	PostCommentReplyAsync(ctx workflow.Context, input *codecommit.PostCommentReplyInput) *CodeCommitPostCommentReplyFuture

	PutCommentReaction(ctx workflow.Context, input *codecommit.PutCommentReactionInput) (*codecommit.PutCommentReactionOutput, error)
	PutCommentReactionAsync(ctx workflow.Context, input *codecommit.PutCommentReactionInput) *CodeCommitPutCommentReactionFuture

	PutFile(ctx workflow.Context, input *codecommit.PutFileInput) (*codecommit.PutFileOutput, error)
	PutFileAsync(ctx workflow.Context, input *codecommit.PutFileInput) *CodeCommitPutFileFuture

	PutRepositoryTriggers(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) (*codecommit.PutRepositoryTriggersOutput, error)
	PutRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.PutRepositoryTriggersInput) *CodeCommitPutRepositoryTriggersFuture

	TagResource(ctx workflow.Context, input *codecommit.TagResourceInput) (*codecommit.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codecommit.TagResourceInput) *CodeCommitTagResourceFuture

	TestRepositoryTriggers(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) (*codecommit.TestRepositoryTriggersOutput, error)
	TestRepositoryTriggersAsync(ctx workflow.Context, input *codecommit.TestRepositoryTriggersInput) *CodeCommitTestRepositoryTriggersFuture

	UntagResource(ctx workflow.Context, input *codecommit.UntagResourceInput) (*codecommit.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codecommit.UntagResourceInput) *CodeCommitUntagResourceFuture

	UpdateApprovalRuleTemplateContent(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) (*codecommit.UpdateApprovalRuleTemplateContentOutput, error)
	UpdateApprovalRuleTemplateContentAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateContentInput) *CodeCommitUpdateApprovalRuleTemplateContentFuture

	UpdateApprovalRuleTemplateDescription(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) (*codecommit.UpdateApprovalRuleTemplateDescriptionOutput, error)
	UpdateApprovalRuleTemplateDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateDescriptionInput) *CodeCommitUpdateApprovalRuleTemplateDescriptionFuture

	UpdateApprovalRuleTemplateName(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) (*codecommit.UpdateApprovalRuleTemplateNameOutput, error)
	UpdateApprovalRuleTemplateNameAsync(ctx workflow.Context, input *codecommit.UpdateApprovalRuleTemplateNameInput) *CodeCommitUpdateApprovalRuleTemplateNameFuture

	UpdateComment(ctx workflow.Context, input *codecommit.UpdateCommentInput) (*codecommit.UpdateCommentOutput, error)
	UpdateCommentAsync(ctx workflow.Context, input *codecommit.UpdateCommentInput) *CodeCommitUpdateCommentFuture

	UpdateDefaultBranch(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) (*codecommit.UpdateDefaultBranchOutput, error)
	UpdateDefaultBranchAsync(ctx workflow.Context, input *codecommit.UpdateDefaultBranchInput) *CodeCommitUpdateDefaultBranchFuture

	UpdatePullRequestApprovalRuleContent(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) (*codecommit.UpdatePullRequestApprovalRuleContentOutput, error)
	UpdatePullRequestApprovalRuleContentAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalRuleContentInput) *CodeCommitUpdatePullRequestApprovalRuleContentFuture

	UpdatePullRequestApprovalState(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) (*codecommit.UpdatePullRequestApprovalStateOutput, error)
	UpdatePullRequestApprovalStateAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestApprovalStateInput) *CodeCommitUpdatePullRequestApprovalStateFuture

	UpdatePullRequestDescription(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) (*codecommit.UpdatePullRequestDescriptionOutput, error)
	UpdatePullRequestDescriptionAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestDescriptionInput) *CodeCommitUpdatePullRequestDescriptionFuture

	UpdatePullRequestStatus(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) (*codecommit.UpdatePullRequestStatusOutput, error)
	UpdatePullRequestStatusAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestStatusInput) *CodeCommitUpdatePullRequestStatusFuture

	UpdatePullRequestTitle(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) (*codecommit.UpdatePullRequestTitleOutput, error)
	UpdatePullRequestTitleAsync(ctx workflow.Context, input *codecommit.UpdatePullRequestTitleInput) *CodeCommitUpdatePullRequestTitleFuture

	UpdateRepositoryDescription(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) (*codecommit.UpdateRepositoryDescriptionOutput, error)
	UpdateRepositoryDescriptionAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryDescriptionInput) *CodeCommitUpdateRepositoryDescriptionFuture

	UpdateRepositoryName(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) (*codecommit.UpdateRepositoryNameOutput, error)
	UpdateRepositoryNameAsync(ctx workflow.Context, input *codecommit.UpdateRepositoryNameInput) *CodeCommitUpdateRepositoryNameFuture
}

func NewClient() Client {
	return &stub{}
}

