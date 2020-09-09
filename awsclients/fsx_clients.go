package awsclients

import (
	"github.com/aws/aws-sdk-go/service/fsx"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type FSxClient interface {
    CancelDataRepositoryTask(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error)
    CancelDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) *FsxCancelDataRepositoryTaskResult

    CreateBackup(ctx workflow.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error)
    CreateBackupAsync(ctx workflow.Context, input *fsx.CreateBackupInput) *FsxCreateBackupResult

    CreateDataRepositoryTask(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error)
    CreateDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) *FsxCreateDataRepositoryTaskResult

    CreateFileSystem(ctx workflow.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error)
    CreateFileSystemAsync(ctx workflow.Context, input *fsx.CreateFileSystemInput) *FsxCreateFileSystemResult

    CreateFileSystemFromBackup(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error)
    CreateFileSystemFromBackupAsync(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) *FsxCreateFileSystemFromBackupResult

    DeleteBackup(ctx workflow.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error)
    DeleteBackupAsync(ctx workflow.Context, input *fsx.DeleteBackupInput) *FsxDeleteBackupResult

    DeleteFileSystem(ctx workflow.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error)
    DeleteFileSystemAsync(ctx workflow.Context, input *fsx.DeleteFileSystemInput) *FsxDeleteFileSystemResult

    DescribeBackups(ctx workflow.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error)
    DescribeBackupsAsync(ctx workflow.Context, input *fsx.DescribeBackupsInput) *FsxDescribeBackupsResult

    DescribeDataRepositoryTasks(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error)
    DescribeDataRepositoryTasksAsync(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) *FsxDescribeDataRepositoryTasksResult

    DescribeFileSystems(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error)
    DescribeFileSystemsAsync(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) *FsxDescribeFileSystemsResult

    ListTagsForResource(ctx workflow.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *fsx.ListTagsForResourceInput) *FsxListTagsForResourceResult

    TagResource(ctx workflow.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *fsx.TagResourceInput) *FsxTagResourceResult

    UntagResource(ctx workflow.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *fsx.UntagResourceInput) *FsxUntagResourceResult

    UpdateFileSystem(ctx workflow.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error)
    UpdateFileSystemAsync(ctx workflow.Context, input *fsx.UpdateFileSystemInput) *FsxUpdateFileSystemResult
}

type FsxCancelDataRepositoryTaskResult struct {
	Result workflow.Future
}

func (r *FsxCancelDataRepositoryTaskResult) Get(ctx workflow.Context) (*fsx.CancelDataRepositoryTaskOutput, error) {
    var output fsx.CancelDataRepositoryTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxCreateBackupResult struct {
	Result workflow.Future
}

func (r *FsxCreateBackupResult) Get(ctx workflow.Context) (*fsx.CreateBackupOutput, error) {
    var output fsx.CreateBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxCreateDataRepositoryTaskResult struct {
	Result workflow.Future
}

func (r *FsxCreateDataRepositoryTaskResult) Get(ctx workflow.Context) (*fsx.CreateDataRepositoryTaskOutput, error) {
    var output fsx.CreateDataRepositoryTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxCreateFileSystemResult struct {
	Result workflow.Future
}

func (r *FsxCreateFileSystemResult) Get(ctx workflow.Context) (*fsx.CreateFileSystemOutput, error) {
    var output fsx.CreateFileSystemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxCreateFileSystemFromBackupResult struct {
	Result workflow.Future
}

func (r *FsxCreateFileSystemFromBackupResult) Get(ctx workflow.Context) (*fsx.CreateFileSystemFromBackupOutput, error) {
    var output fsx.CreateFileSystemFromBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxDeleteBackupResult struct {
	Result workflow.Future
}

func (r *FsxDeleteBackupResult) Get(ctx workflow.Context) (*fsx.DeleteBackupOutput, error) {
    var output fsx.DeleteBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxDeleteFileSystemResult struct {
	Result workflow.Future
}

func (r *FsxDeleteFileSystemResult) Get(ctx workflow.Context) (*fsx.DeleteFileSystemOutput, error) {
    var output fsx.DeleteFileSystemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxDescribeBackupsResult struct {
	Result workflow.Future
}

func (r *FsxDescribeBackupsResult) Get(ctx workflow.Context) (*fsx.DescribeBackupsOutput, error) {
    var output fsx.DescribeBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxDescribeDataRepositoryTasksResult struct {
	Result workflow.Future
}

func (r *FsxDescribeDataRepositoryTasksResult) Get(ctx workflow.Context) (*fsx.DescribeDataRepositoryTasksOutput, error) {
    var output fsx.DescribeDataRepositoryTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxDescribeFileSystemsResult struct {
	Result workflow.Future
}

func (r *FsxDescribeFileSystemsResult) Get(ctx workflow.Context) (*fsx.DescribeFileSystemsOutput, error) {
    var output fsx.DescribeFileSystemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *FsxListTagsForResourceResult) Get(ctx workflow.Context) (*fsx.ListTagsForResourceOutput, error) {
    var output fsx.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxTagResourceResult struct {
	Result workflow.Future
}

func (r *FsxTagResourceResult) Get(ctx workflow.Context) (*fsx.TagResourceOutput, error) {
    var output fsx.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxUntagResourceResult struct {
	Result workflow.Future
}

func (r *FsxUntagResourceResult) Get(ctx workflow.Context) (*fsx.UntagResourceOutput, error) {
    var output fsx.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FsxUpdateFileSystemResult struct {
	Result workflow.Future
}

func (r *FsxUpdateFileSystemResult) Get(ctx workflow.Context) (*fsx.UpdateFileSystemOutput, error) {
    var output fsx.UpdateFileSystemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type FSxStub struct {
    activities awsactivities.FSxActivities
}

func NewFSxStub() FSxClient {
    return &FSxStub{}
}

func (a *FSxStub) CancelDataRepositoryTask(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error) {
    var output fsx.CancelDataRepositoryTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelDataRepositoryTask, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) CancelDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) *FsxCancelDataRepositoryTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelDataRepositoryTask, input)
    return &FsxCancelDataRepositoryTaskResult{Result: future}
}

func (a *FSxStub) CreateBackup(ctx workflow.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
    var output fsx.CreateBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) CreateBackupAsync(ctx workflow.Context, input *fsx.CreateBackupInput) *FsxCreateBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackup, input)
    return &FsxCreateBackupResult{Result: future}
}

func (a *FSxStub) CreateDataRepositoryTask(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error) {
    var output fsx.CreateDataRepositoryTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataRepositoryTask, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) CreateDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) *FsxCreateDataRepositoryTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataRepositoryTask, input)
    return &FsxCreateDataRepositoryTaskResult{Result: future}
}

func (a *FSxStub) CreateFileSystem(ctx workflow.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error) {
    var output fsx.CreateFileSystemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) CreateFileSystemAsync(ctx workflow.Context, input *fsx.CreateFileSystemInput) *FsxCreateFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystem, input)
    return &FsxCreateFileSystemResult{Result: future}
}

func (a *FSxStub) CreateFileSystemFromBackup(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error) {
    var output fsx.CreateFileSystemFromBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystemFromBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) CreateFileSystemFromBackupAsync(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) *FsxCreateFileSystemFromBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystemFromBackup, input)
    return &FsxCreateFileSystemFromBackupResult{Result: future}
}

func (a *FSxStub) DeleteBackup(ctx workflow.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error) {
    var output fsx.DeleteBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) DeleteBackupAsync(ctx workflow.Context, input *fsx.DeleteBackupInput) *FsxDeleteBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input)
    return &FsxDeleteBackupResult{Result: future}
}

func (a *FSxStub) DeleteFileSystem(ctx workflow.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error) {
    var output fsx.DeleteFileSystemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) DeleteFileSystemAsync(ctx workflow.Context, input *fsx.DeleteFileSystemInput) *FsxDeleteFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystem, input)
    return &FsxDeleteFileSystemResult{Result: future}
}

func (a *FSxStub) DescribeBackups(ctx workflow.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
    var output fsx.DescribeBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) DescribeBackupsAsync(ctx workflow.Context, input *fsx.DescribeBackupsInput) *FsxDescribeBackupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackups, input)
    return &FsxDescribeBackupsResult{Result: future}
}

func (a *FSxStub) DescribeDataRepositoryTasks(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
    var output fsx.DescribeDataRepositoryTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataRepositoryTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) DescribeDataRepositoryTasksAsync(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) *FsxDescribeDataRepositoryTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataRepositoryTasks, input)
    return &FsxDescribeDataRepositoryTasksResult{Result: future}
}

func (a *FSxStub) DescribeFileSystems(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
    var output fsx.DescribeFileSystemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystems, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) DescribeFileSystemsAsync(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) *FsxDescribeFileSystemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystems, input)
    return &FsxDescribeFileSystemsResult{Result: future}
}

func (a *FSxStub) ListTagsForResource(ctx workflow.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
    var output fsx.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) ListTagsForResourceAsync(ctx workflow.Context, input *fsx.ListTagsForResourceInput) *FsxListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &FsxListTagsForResourceResult{Result: future}
}

func (a *FSxStub) TagResource(ctx workflow.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error) {
    var output fsx.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) TagResourceAsync(ctx workflow.Context, input *fsx.TagResourceInput) *FsxTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &FsxTagResourceResult{Result: future}
}

func (a *FSxStub) UntagResource(ctx workflow.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error) {
    var output fsx.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) UntagResourceAsync(ctx workflow.Context, input *fsx.UntagResourceInput) *FsxUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &FsxUntagResourceResult{Result: future}
}

func (a *FSxStub) UpdateFileSystem(ctx workflow.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error) {
    var output fsx.UpdateFileSystemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *FSxStub) UpdateFileSystemAsync(ctx workflow.Context, input *fsx.UpdateFileSystemInput) *FsxUpdateFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFileSystem, input)
    return &FsxUpdateFileSystemResult{Result: future}
}
