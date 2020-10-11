// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package workdocsstub

import (
	"github.com/aws/aws-sdk-go/service/workdocs"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type WorkDocsAbortDocumentVersionUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsAbortDocumentVersionUploadFuture) Get(ctx workflow.Context) (*workdocs.AbortDocumentVersionUploadOutput, error) {
	var output workdocs.AbortDocumentVersionUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsActivateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsActivateUserFuture) Get(ctx workflow.Context) (*workdocs.ActivateUserOutput, error) {
	var output workdocs.ActivateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsAddResourcePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsAddResourcePermissionsFuture) Get(ctx workflow.Context) (*workdocs.AddResourcePermissionsOutput, error) {
	var output workdocs.AddResourcePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateCommentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateCommentFuture) Get(ctx workflow.Context) (*workdocs.CreateCommentOutput, error) {
	var output workdocs.CreateCommentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateCustomMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateCustomMetadataFuture) Get(ctx workflow.Context) (*workdocs.CreateCustomMetadataOutput, error) {
	var output workdocs.CreateCustomMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateFolderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateFolderFuture) Get(ctx workflow.Context) (*workdocs.CreateFolderOutput, error) {
	var output workdocs.CreateFolderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateLabelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateLabelsFuture) Get(ctx workflow.Context) (*workdocs.CreateLabelsOutput, error) {
	var output workdocs.CreateLabelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateNotificationSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateNotificationSubscriptionFuture) Get(ctx workflow.Context) (*workdocs.CreateNotificationSubscriptionOutput, error) {
	var output workdocs.CreateNotificationSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsCreateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsCreateUserFuture) Get(ctx workflow.Context) (*workdocs.CreateUserOutput, error) {
	var output workdocs.CreateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeactivateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeactivateUserFuture) Get(ctx workflow.Context) (*workdocs.DeactivateUserOutput, error) {
	var output workdocs.DeactivateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteCommentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteCommentFuture) Get(ctx workflow.Context) (*workdocs.DeleteCommentOutput, error) {
	var output workdocs.DeleteCommentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteCustomMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteCustomMetadataFuture) Get(ctx workflow.Context) (*workdocs.DeleteCustomMetadataOutput, error) {
	var output workdocs.DeleteCustomMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteDocumentFuture) Get(ctx workflow.Context) (*workdocs.DeleteDocumentOutput, error) {
	var output workdocs.DeleteDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteFolderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteFolderFuture) Get(ctx workflow.Context) (*workdocs.DeleteFolderOutput, error) {
	var output workdocs.DeleteFolderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteFolderContentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteFolderContentsFuture) Get(ctx workflow.Context) (*workdocs.DeleteFolderContentsOutput, error) {
	var output workdocs.DeleteFolderContentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteLabelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteLabelsFuture) Get(ctx workflow.Context) (*workdocs.DeleteLabelsOutput, error) {
	var output workdocs.DeleteLabelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteNotificationSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteNotificationSubscriptionFuture) Get(ctx workflow.Context) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
	var output workdocs.DeleteNotificationSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDeleteUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDeleteUserFuture) Get(ctx workflow.Context) (*workdocs.DeleteUserOutput, error) {
	var output workdocs.DeleteUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeActivitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeActivitiesFuture) Get(ctx workflow.Context) (*workdocs.DescribeActivitiesOutput, error) {
	var output workdocs.DescribeActivitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeCommentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeCommentsFuture) Get(ctx workflow.Context) (*workdocs.DescribeCommentsOutput, error) {
	var output workdocs.DescribeCommentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeDocumentVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeDocumentVersionsFuture) Get(ctx workflow.Context) (*workdocs.DescribeDocumentVersionsOutput, error) {
	var output workdocs.DescribeDocumentVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeFolderContentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeFolderContentsFuture) Get(ctx workflow.Context) (*workdocs.DescribeFolderContentsOutput, error) {
	var output workdocs.DescribeFolderContentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeGroupsFuture) Get(ctx workflow.Context) (*workdocs.DescribeGroupsOutput, error) {
	var output workdocs.DescribeGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeNotificationSubscriptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeNotificationSubscriptionsFuture) Get(ctx workflow.Context) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
	var output workdocs.DescribeNotificationSubscriptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeResourcePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeResourcePermissionsFuture) Get(ctx workflow.Context) (*workdocs.DescribeResourcePermissionsOutput, error) {
	var output workdocs.DescribeResourcePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeRootFoldersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeRootFoldersFuture) Get(ctx workflow.Context) (*workdocs.DescribeRootFoldersOutput, error) {
	var output workdocs.DescribeRootFoldersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsDescribeUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsDescribeUsersFuture) Get(ctx workflow.Context) (*workdocs.DescribeUsersOutput, error) {
	var output workdocs.DescribeUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetCurrentUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetCurrentUserFuture) Get(ctx workflow.Context) (*workdocs.GetCurrentUserOutput, error) {
	var output workdocs.GetCurrentUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetDocumentFuture) Get(ctx workflow.Context) (*workdocs.GetDocumentOutput, error) {
	var output workdocs.GetDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetDocumentPathFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetDocumentPathFuture) Get(ctx workflow.Context) (*workdocs.GetDocumentPathOutput, error) {
	var output workdocs.GetDocumentPathOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetDocumentVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetDocumentVersionFuture) Get(ctx workflow.Context) (*workdocs.GetDocumentVersionOutput, error) {
	var output workdocs.GetDocumentVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetFolderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetFolderFuture) Get(ctx workflow.Context) (*workdocs.GetFolderOutput, error) {
	var output workdocs.GetFolderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetFolderPathFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetFolderPathFuture) Get(ctx workflow.Context) (*workdocs.GetFolderPathOutput, error) {
	var output workdocs.GetFolderPathOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsGetResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsGetResourcesFuture) Get(ctx workflow.Context) (*workdocs.GetResourcesOutput, error) {
	var output workdocs.GetResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsInitiateDocumentVersionUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsInitiateDocumentVersionUploadFuture) Get(ctx workflow.Context) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
	var output workdocs.InitiateDocumentVersionUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsRemoveAllResourcePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsRemoveAllResourcePermissionsFuture) Get(ctx workflow.Context) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
	var output workdocs.RemoveAllResourcePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsRemoveResourcePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsRemoveResourcePermissionFuture) Get(ctx workflow.Context) (*workdocs.RemoveResourcePermissionOutput, error) {
	var output workdocs.RemoveResourcePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsUpdateDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsUpdateDocumentFuture) Get(ctx workflow.Context) (*workdocs.UpdateDocumentOutput, error) {
	var output workdocs.UpdateDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsUpdateDocumentVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsUpdateDocumentVersionFuture) Get(ctx workflow.Context) (*workdocs.UpdateDocumentVersionOutput, error) {
	var output workdocs.UpdateDocumentVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsUpdateFolderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsUpdateFolderFuture) Get(ctx workflow.Context) (*workdocs.UpdateFolderOutput, error) {
	var output workdocs.UpdateFolderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkDocsUpdateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkDocsUpdateUserFuture) Get(ctx workflow.Context) (*workdocs.UpdateUserOutput, error) {
	var output workdocs.UpdateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AbortDocumentVersionUpload(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error) {
	var output workdocs.AbortDocumentVersionUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.AbortDocumentVersionUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AbortDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) *WorkDocsAbortDocumentVersionUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.AbortDocumentVersionUpload", input)
	return &WorkDocsAbortDocumentVersionUploadFuture{Future: future}
}

func (a *stub) ActivateUser(ctx workflow.Context, input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error) {
	var output workdocs.ActivateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.ActivateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateUserAsync(ctx workflow.Context, input *workdocs.ActivateUserInput) *WorkDocsActivateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.ActivateUser", input)
	return &WorkDocsActivateUserFuture{Future: future}
}

func (a *stub) AddResourcePermissions(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error) {
	var output workdocs.AddResourcePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.AddResourcePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddResourcePermissionsAsync(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) *WorkDocsAddResourcePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.AddResourcePermissions", input)
	return &WorkDocsAddResourcePermissionsFuture{Future: future}
}

func (a *stub) CreateComment(ctx workflow.Context, input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error) {
	var output workdocs.CreateCommentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateComment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCommentAsync(ctx workflow.Context, input *workdocs.CreateCommentInput) *WorkDocsCreateCommentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateComment", input)
	return &WorkDocsCreateCommentFuture{Future: future}
}

func (a *stub) CreateCustomMetadata(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error) {
	var output workdocs.CreateCustomMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateCustomMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCustomMetadataAsync(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) *WorkDocsCreateCustomMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateCustomMetadata", input)
	return &WorkDocsCreateCustomMetadataFuture{Future: future}
}

func (a *stub) CreateFolder(ctx workflow.Context, input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error) {
	var output workdocs.CreateFolderOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateFolder", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFolderAsync(ctx workflow.Context, input *workdocs.CreateFolderInput) *WorkDocsCreateFolderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateFolder", input)
	return &WorkDocsCreateFolderFuture{Future: future}
}

func (a *stub) CreateLabels(ctx workflow.Context, input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error) {
	var output workdocs.CreateLabelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateLabels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLabelsAsync(ctx workflow.Context, input *workdocs.CreateLabelsInput) *WorkDocsCreateLabelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateLabels", input)
	return &WorkDocsCreateLabelsFuture{Future: future}
}

func (a *stub) CreateNotificationSubscription(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error) {
	var output workdocs.CreateNotificationSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateNotificationSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) *WorkDocsCreateNotificationSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateNotificationSubscription", input)
	return &WorkDocsCreateNotificationSubscriptionFuture{Future: future}
}

func (a *stub) CreateUser(ctx workflow.Context, input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error) {
	var output workdocs.CreateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateUserAsync(ctx workflow.Context, input *workdocs.CreateUserInput) *WorkDocsCreateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.CreateUser", input)
	return &WorkDocsCreateUserFuture{Future: future}
}

func (a *stub) DeactivateUser(ctx workflow.Context, input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error) {
	var output workdocs.DeactivateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeactivateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivateUserAsync(ctx workflow.Context, input *workdocs.DeactivateUserInput) *WorkDocsDeactivateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeactivateUser", input)
	return &WorkDocsDeactivateUserFuture{Future: future}
}

func (a *stub) DeleteComment(ctx workflow.Context, input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error) {
	var output workdocs.DeleteCommentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteComment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteCommentAsync(ctx workflow.Context, input *workdocs.DeleteCommentInput) *WorkDocsDeleteCommentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteComment", input)
	return &WorkDocsDeleteCommentFuture{Future: future}
}

func (a *stub) DeleteCustomMetadata(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error) {
	var output workdocs.DeleteCustomMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteCustomMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteCustomMetadataAsync(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) *WorkDocsDeleteCustomMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteCustomMetadata", input)
	return &WorkDocsDeleteCustomMetadataFuture{Future: future}
}

func (a *stub) DeleteDocument(ctx workflow.Context, input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error) {
	var output workdocs.DeleteDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDocumentAsync(ctx workflow.Context, input *workdocs.DeleteDocumentInput) *WorkDocsDeleteDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteDocument", input)
	return &WorkDocsDeleteDocumentFuture{Future: future}
}

func (a *stub) DeleteFolder(ctx workflow.Context, input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error) {
	var output workdocs.DeleteFolderOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteFolder", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFolderAsync(ctx workflow.Context, input *workdocs.DeleteFolderInput) *WorkDocsDeleteFolderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteFolder", input)
	return &WorkDocsDeleteFolderFuture{Future: future}
}

func (a *stub) DeleteFolderContents(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error) {
	var output workdocs.DeleteFolderContentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteFolderContents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFolderContentsAsync(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) *WorkDocsDeleteFolderContentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteFolderContents", input)
	return &WorkDocsDeleteFolderContentsFuture{Future: future}
}

func (a *stub) DeleteLabels(ctx workflow.Context, input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error) {
	var output workdocs.DeleteLabelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteLabels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLabelsAsync(ctx workflow.Context, input *workdocs.DeleteLabelsInput) *WorkDocsDeleteLabelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteLabels", input)
	return &WorkDocsDeleteLabelsFuture{Future: future}
}

func (a *stub) DeleteNotificationSubscription(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
	var output workdocs.DeleteNotificationSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteNotificationSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) *WorkDocsDeleteNotificationSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteNotificationSubscription", input)
	return &WorkDocsDeleteNotificationSubscriptionFuture{Future: future}
}

func (a *stub) DeleteUser(ctx workflow.Context, input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error) {
	var output workdocs.DeleteUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteUserAsync(ctx workflow.Context, input *workdocs.DeleteUserInput) *WorkDocsDeleteUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DeleteUser", input)
	return &WorkDocsDeleteUserFuture{Future: future}
}

func (a *stub) DescribeActivities(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error) {
	var output workdocs.DescribeActivitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeActivities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeActivitiesAsync(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) *WorkDocsDescribeActivitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeActivities", input)
	return &WorkDocsDescribeActivitiesFuture{Future: future}
}

func (a *stub) DescribeComments(ctx workflow.Context, input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error) {
	var output workdocs.DescribeCommentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeComments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCommentsAsync(ctx workflow.Context, input *workdocs.DescribeCommentsInput) *WorkDocsDescribeCommentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeComments", input)
	return &WorkDocsDescribeCommentsFuture{Future: future}
}

func (a *stub) DescribeDocumentVersions(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error) {
	var output workdocs.DescribeDocumentVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeDocumentVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDocumentVersionsAsync(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) *WorkDocsDescribeDocumentVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeDocumentVersions", input)
	return &WorkDocsDescribeDocumentVersionsFuture{Future: future}
}

func (a *stub) DescribeFolderContents(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error) {
	var output workdocs.DescribeFolderContentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeFolderContents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFolderContentsAsync(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) *WorkDocsDescribeFolderContentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeFolderContents", input)
	return &WorkDocsDescribeFolderContentsFuture{Future: future}
}

func (a *stub) DescribeGroups(ctx workflow.Context, input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error) {
	var output workdocs.DescribeGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGroupsAsync(ctx workflow.Context, input *workdocs.DescribeGroupsInput) *WorkDocsDescribeGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeGroups", input)
	return &WorkDocsDescribeGroupsFuture{Future: future}
}

func (a *stub) DescribeNotificationSubscriptions(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
	var output workdocs.DescribeNotificationSubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeNotificationSubscriptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeNotificationSubscriptionsAsync(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) *WorkDocsDescribeNotificationSubscriptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeNotificationSubscriptions", input)
	return &WorkDocsDescribeNotificationSubscriptionsFuture{Future: future}
}

func (a *stub) DescribeResourcePermissions(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error) {
	var output workdocs.DescribeResourcePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeResourcePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeResourcePermissionsAsync(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) *WorkDocsDescribeResourcePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeResourcePermissions", input)
	return &WorkDocsDescribeResourcePermissionsFuture{Future: future}
}

func (a *stub) DescribeRootFolders(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error) {
	var output workdocs.DescribeRootFoldersOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeRootFolders", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRootFoldersAsync(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) *WorkDocsDescribeRootFoldersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeRootFolders", input)
	return &WorkDocsDescribeRootFoldersFuture{Future: future}
}

func (a *stub) DescribeUsers(ctx workflow.Context, input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error) {
	var output workdocs.DescribeUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeUsersAsync(ctx workflow.Context, input *workdocs.DescribeUsersInput) *WorkDocsDescribeUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.DescribeUsers", input)
	return &WorkDocsDescribeUsersFuture{Future: future}
}

func (a *stub) GetCurrentUser(ctx workflow.Context, input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error) {
	var output workdocs.GetCurrentUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetCurrentUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCurrentUserAsync(ctx workflow.Context, input *workdocs.GetCurrentUserInput) *WorkDocsGetCurrentUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetCurrentUser", input)
	return &WorkDocsGetCurrentUserFuture{Future: future}
}

func (a *stub) GetDocument(ctx workflow.Context, input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error) {
	var output workdocs.GetDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDocumentAsync(ctx workflow.Context, input *workdocs.GetDocumentInput) *WorkDocsGetDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocument", input)
	return &WorkDocsGetDocumentFuture{Future: future}
}

func (a *stub) GetDocumentPath(ctx workflow.Context, input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error) {
	var output workdocs.GetDocumentPathOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocumentPath", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDocumentPathAsync(ctx workflow.Context, input *workdocs.GetDocumentPathInput) *WorkDocsGetDocumentPathFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocumentPath", input)
	return &WorkDocsGetDocumentPathFuture{Future: future}
}

func (a *stub) GetDocumentVersion(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error) {
	var output workdocs.GetDocumentVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocumentVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDocumentVersionAsync(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) *WorkDocsGetDocumentVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetDocumentVersion", input)
	return &WorkDocsGetDocumentVersionFuture{Future: future}
}

func (a *stub) GetFolder(ctx workflow.Context, input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error) {
	var output workdocs.GetFolderOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetFolder", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFolderAsync(ctx workflow.Context, input *workdocs.GetFolderInput) *WorkDocsGetFolderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetFolder", input)
	return &WorkDocsGetFolderFuture{Future: future}
}

func (a *stub) GetFolderPath(ctx workflow.Context, input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error) {
	var output workdocs.GetFolderPathOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetFolderPath", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFolderPathAsync(ctx workflow.Context, input *workdocs.GetFolderPathInput) *WorkDocsGetFolderPathFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetFolderPath", input)
	return &WorkDocsGetFolderPathFuture{Future: future}
}

func (a *stub) GetResources(ctx workflow.Context, input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error) {
	var output workdocs.GetResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.GetResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResourcesAsync(ctx workflow.Context, input *workdocs.GetResourcesInput) *WorkDocsGetResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.GetResources", input)
	return &WorkDocsGetResourcesFuture{Future: future}
}

func (a *stub) InitiateDocumentVersionUpload(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
	var output workdocs.InitiateDocumentVersionUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.InitiateDocumentVersionUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitiateDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) *WorkDocsInitiateDocumentVersionUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.InitiateDocumentVersionUpload", input)
	return &WorkDocsInitiateDocumentVersionUploadFuture{Future: future}
}

func (a *stub) RemoveAllResourcePermissions(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
	var output workdocs.RemoveAllResourcePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.RemoveAllResourcePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveAllResourcePermissionsAsync(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) *WorkDocsRemoveAllResourcePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.RemoveAllResourcePermissions", input)
	return &WorkDocsRemoveAllResourcePermissionsFuture{Future: future}
}

func (a *stub) RemoveResourcePermission(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error) {
	var output workdocs.RemoveResourcePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.RemoveResourcePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveResourcePermissionAsync(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) *WorkDocsRemoveResourcePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.RemoveResourcePermission", input)
	return &WorkDocsRemoveResourcePermissionFuture{Future: future}
}

func (a *stub) UpdateDocument(ctx workflow.Context, input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error) {
	var output workdocs.UpdateDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDocumentAsync(ctx workflow.Context, input *workdocs.UpdateDocumentInput) *WorkDocsUpdateDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateDocument", input)
	return &WorkDocsUpdateDocumentFuture{Future: future}
}

func (a *stub) UpdateDocumentVersion(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error) {
	var output workdocs.UpdateDocumentVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateDocumentVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDocumentVersionAsync(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) *WorkDocsUpdateDocumentVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateDocumentVersion", input)
	return &WorkDocsUpdateDocumentVersionFuture{Future: future}
}

func (a *stub) UpdateFolder(ctx workflow.Context, input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error) {
	var output workdocs.UpdateFolderOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateFolder", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFolderAsync(ctx workflow.Context, input *workdocs.UpdateFolderInput) *WorkDocsUpdateFolderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateFolder", input)
	return &WorkDocsUpdateFolderFuture{Future: future}
}

func (a *stub) UpdateUser(ctx workflow.Context, input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error) {
	var output workdocs.UpdateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateUserAsync(ctx workflow.Context, input *workdocs.UpdateUserInput) *WorkDocsUpdateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workdocs.UpdateUser", input)
	return &WorkDocsUpdateUserFuture{Future: future}
}