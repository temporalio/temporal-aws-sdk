package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MediaConvertClient interface {
    AssociateCertificate(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error)
    AssociateCertificateAsync(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) *MediaconvertAssociateCertificateResult

    CancelJob(ctx workflow.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error)
    CancelJobAsync(ctx workflow.Context, input *mediaconvert.CancelJobInput) *MediaconvertCancelJobResult

    CreateJob(ctx workflow.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error)
    CreateJobAsync(ctx workflow.Context, input *mediaconvert.CreateJobInput) *MediaconvertCreateJobResult

    CreateJobTemplate(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error)
    CreateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) *MediaconvertCreateJobTemplateResult

    CreatePreset(ctx workflow.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error)
    CreatePresetAsync(ctx workflow.Context, input *mediaconvert.CreatePresetInput) *MediaconvertCreatePresetResult

    CreateQueue(ctx workflow.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error)
    CreateQueueAsync(ctx workflow.Context, input *mediaconvert.CreateQueueInput) *MediaconvertCreateQueueResult

    DeleteJobTemplate(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error)
    DeleteJobTemplateAsync(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) *MediaconvertDeleteJobTemplateResult

    DeletePreset(ctx workflow.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error)
    DeletePresetAsync(ctx workflow.Context, input *mediaconvert.DeletePresetInput) *MediaconvertDeletePresetResult

    DeleteQueue(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error)
    DeleteQueueAsync(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) *MediaconvertDeleteQueueResult

    DescribeEndpoints(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error)
    DescribeEndpointsAsync(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) *MediaconvertDescribeEndpointsResult

    DisassociateCertificate(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error)
    DisassociateCertificateAsync(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) *MediaconvertDisassociateCertificateResult

    GetJob(ctx workflow.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error)
    GetJobAsync(ctx workflow.Context, input *mediaconvert.GetJobInput) *MediaconvertGetJobResult

    GetJobTemplate(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error)
    GetJobTemplateAsync(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) *MediaconvertGetJobTemplateResult

    GetPreset(ctx workflow.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error)
    GetPresetAsync(ctx workflow.Context, input *mediaconvert.GetPresetInput) *MediaconvertGetPresetResult

    GetQueue(ctx workflow.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error)
    GetQueueAsync(ctx workflow.Context, input *mediaconvert.GetQueueInput) *MediaconvertGetQueueResult

    ListJobTemplates(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error)
    ListJobTemplatesAsync(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) *MediaconvertListJobTemplatesResult

    ListJobs(ctx workflow.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *mediaconvert.ListJobsInput) *MediaconvertListJobsResult

    ListPresets(ctx workflow.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error)
    ListPresetsAsync(ctx workflow.Context, input *mediaconvert.ListPresetsInput) *MediaconvertListPresetsResult

    ListQueues(ctx workflow.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error)
    ListQueuesAsync(ctx workflow.Context, input *mediaconvert.ListQueuesInput) *MediaconvertListQueuesResult

    ListTagsForResource(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) *MediaconvertListTagsForResourceResult

    TagResource(ctx workflow.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *mediaconvert.TagResourceInput) *MediaconvertTagResourceResult

    UntagResource(ctx workflow.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *mediaconvert.UntagResourceInput) *MediaconvertUntagResourceResult

    UpdateJobTemplate(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error)
    UpdateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) *MediaconvertUpdateJobTemplateResult

    UpdatePreset(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error)
    UpdatePresetAsync(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) *MediaconvertUpdatePresetResult

    UpdateQueue(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error)
    UpdateQueueAsync(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) *MediaconvertUpdateQueueResult
}

type MediaconvertAssociateCertificateResult struct {
	Result workflow.Future
}

func (r *MediaconvertAssociateCertificateResult) Get(ctx workflow.Context) (*mediaconvert.AssociateCertificateOutput, error) {
    var output mediaconvert.AssociateCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertCancelJobResult struct {
	Result workflow.Future
}

func (r *MediaconvertCancelJobResult) Get(ctx workflow.Context) (*mediaconvert.CancelJobOutput, error) {
    var output mediaconvert.CancelJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertCreateJobResult struct {
	Result workflow.Future
}

func (r *MediaconvertCreateJobResult) Get(ctx workflow.Context) (*mediaconvert.CreateJobOutput, error) {
    var output mediaconvert.CreateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertCreateJobTemplateResult struct {
	Result workflow.Future
}

func (r *MediaconvertCreateJobTemplateResult) Get(ctx workflow.Context) (*mediaconvert.CreateJobTemplateOutput, error) {
    var output mediaconvert.CreateJobTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertCreatePresetResult struct {
	Result workflow.Future
}

func (r *MediaconvertCreatePresetResult) Get(ctx workflow.Context) (*mediaconvert.CreatePresetOutput, error) {
    var output mediaconvert.CreatePresetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertCreateQueueResult struct {
	Result workflow.Future
}

func (r *MediaconvertCreateQueueResult) Get(ctx workflow.Context) (*mediaconvert.CreateQueueOutput, error) {
    var output mediaconvert.CreateQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertDeleteJobTemplateResult struct {
	Result workflow.Future
}

func (r *MediaconvertDeleteJobTemplateResult) Get(ctx workflow.Context) (*mediaconvert.DeleteJobTemplateOutput, error) {
    var output mediaconvert.DeleteJobTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertDeletePresetResult struct {
	Result workflow.Future
}

func (r *MediaconvertDeletePresetResult) Get(ctx workflow.Context) (*mediaconvert.DeletePresetOutput, error) {
    var output mediaconvert.DeletePresetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertDeleteQueueResult struct {
	Result workflow.Future
}

func (r *MediaconvertDeleteQueueResult) Get(ctx workflow.Context) (*mediaconvert.DeleteQueueOutput, error) {
    var output mediaconvert.DeleteQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertDescribeEndpointsResult struct {
	Result workflow.Future
}

func (r *MediaconvertDescribeEndpointsResult) Get(ctx workflow.Context) (*mediaconvert.DescribeEndpointsOutput, error) {
    var output mediaconvert.DescribeEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertDisassociateCertificateResult struct {
	Result workflow.Future
}

func (r *MediaconvertDisassociateCertificateResult) Get(ctx workflow.Context) (*mediaconvert.DisassociateCertificateOutput, error) {
    var output mediaconvert.DisassociateCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertGetJobResult struct {
	Result workflow.Future
}

func (r *MediaconvertGetJobResult) Get(ctx workflow.Context) (*mediaconvert.GetJobOutput, error) {
    var output mediaconvert.GetJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertGetJobTemplateResult struct {
	Result workflow.Future
}

func (r *MediaconvertGetJobTemplateResult) Get(ctx workflow.Context) (*mediaconvert.GetJobTemplateOutput, error) {
    var output mediaconvert.GetJobTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertGetPresetResult struct {
	Result workflow.Future
}

func (r *MediaconvertGetPresetResult) Get(ctx workflow.Context) (*mediaconvert.GetPresetOutput, error) {
    var output mediaconvert.GetPresetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertGetQueueResult struct {
	Result workflow.Future
}

func (r *MediaconvertGetQueueResult) Get(ctx workflow.Context) (*mediaconvert.GetQueueOutput, error) {
    var output mediaconvert.GetQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertListJobTemplatesResult struct {
	Result workflow.Future
}

func (r *MediaconvertListJobTemplatesResult) Get(ctx workflow.Context) (*mediaconvert.ListJobTemplatesOutput, error) {
    var output mediaconvert.ListJobTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertListJobsResult struct {
	Result workflow.Future
}

func (r *MediaconvertListJobsResult) Get(ctx workflow.Context) (*mediaconvert.ListJobsOutput, error) {
    var output mediaconvert.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertListPresetsResult struct {
	Result workflow.Future
}

func (r *MediaconvertListPresetsResult) Get(ctx workflow.Context) (*mediaconvert.ListPresetsOutput, error) {
    var output mediaconvert.ListPresetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertListQueuesResult struct {
	Result workflow.Future
}

func (r *MediaconvertListQueuesResult) Get(ctx workflow.Context) (*mediaconvert.ListQueuesOutput, error) {
    var output mediaconvert.ListQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediaconvertListTagsForResourceResult) Get(ctx workflow.Context) (*mediaconvert.ListTagsForResourceOutput, error) {
    var output mediaconvert.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertTagResourceResult struct {
	Result workflow.Future
}

func (r *MediaconvertTagResourceResult) Get(ctx workflow.Context) (*mediaconvert.TagResourceOutput, error) {
    var output mediaconvert.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediaconvertUntagResourceResult) Get(ctx workflow.Context) (*mediaconvert.UntagResourceOutput, error) {
    var output mediaconvert.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertUpdateJobTemplateResult struct {
	Result workflow.Future
}

func (r *MediaconvertUpdateJobTemplateResult) Get(ctx workflow.Context) (*mediaconvert.UpdateJobTemplateOutput, error) {
    var output mediaconvert.UpdateJobTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertUpdatePresetResult struct {
	Result workflow.Future
}

func (r *MediaconvertUpdatePresetResult) Get(ctx workflow.Context) (*mediaconvert.UpdatePresetOutput, error) {
    var output mediaconvert.UpdatePresetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconvertUpdateQueueResult struct {
	Result workflow.Future
}

func (r *MediaconvertUpdateQueueResult) Get(ctx workflow.Context) (*mediaconvert.UpdateQueueOutput, error) {
    var output mediaconvert.UpdateQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaConvertStub struct {
    activities awsactivities.MediaConvertActivities
}

func NewMediaConvertStub() MediaConvertClient {
    return &MediaConvertStub{}
}

func (a *MediaConvertStub) AssociateCertificate(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error) {
    var output mediaconvert.AssociateCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) AssociateCertificateAsync(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) *MediaconvertAssociateCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateCertificate, input)
    return &MediaconvertAssociateCertificateResult{Result: future}
}

func (a *MediaConvertStub) CancelJob(ctx workflow.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
    var output mediaconvert.CancelJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) CancelJobAsync(ctx workflow.Context, input *mediaconvert.CancelJobInput) *MediaconvertCancelJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
    return &MediaconvertCancelJobResult{Result: future}
}

func (a *MediaConvertStub) CreateJob(ctx workflow.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error) {
    var output mediaconvert.CreateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) CreateJobAsync(ctx workflow.Context, input *mediaconvert.CreateJobInput) *MediaconvertCreateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
    return &MediaconvertCreateJobResult{Result: future}
}

func (a *MediaConvertStub) CreateJobTemplate(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error) {
    var output mediaconvert.CreateJobTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJobTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) CreateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) *MediaconvertCreateJobTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJobTemplate, input)
    return &MediaconvertCreateJobTemplateResult{Result: future}
}

func (a *MediaConvertStub) CreatePreset(ctx workflow.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error) {
    var output mediaconvert.CreatePresetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePreset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) CreatePresetAsync(ctx workflow.Context, input *mediaconvert.CreatePresetInput) *MediaconvertCreatePresetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePreset, input)
    return &MediaconvertCreatePresetResult{Result: future}
}

func (a *MediaConvertStub) CreateQueue(ctx workflow.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error) {
    var output mediaconvert.CreateQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) CreateQueueAsync(ctx workflow.Context, input *mediaconvert.CreateQueueInput) *MediaconvertCreateQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateQueue, input)
    return &MediaconvertCreateQueueResult{Result: future}
}

func (a *MediaConvertStub) DeleteJobTemplate(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error) {
    var output mediaconvert.DeleteJobTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJobTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) DeleteJobTemplateAsync(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) *MediaconvertDeleteJobTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJobTemplate, input)
    return &MediaconvertDeleteJobTemplateResult{Result: future}
}

func (a *MediaConvertStub) DeletePreset(ctx workflow.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error) {
    var output mediaconvert.DeletePresetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePreset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) DeletePresetAsync(ctx workflow.Context, input *mediaconvert.DeletePresetInput) *MediaconvertDeletePresetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePreset, input)
    return &MediaconvertDeletePresetResult{Result: future}
}

func (a *MediaConvertStub) DeleteQueue(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error) {
    var output mediaconvert.DeleteQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) DeleteQueueAsync(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) *MediaconvertDeleteQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQueue, input)
    return &MediaconvertDeleteQueueResult{Result: future}
}

func (a *MediaConvertStub) DescribeEndpoints(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
    var output mediaconvert.DescribeEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) DescribeEndpointsAsync(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) *MediaconvertDescribeEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input)
    return &MediaconvertDescribeEndpointsResult{Result: future}
}

func (a *MediaConvertStub) DisassociateCertificate(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error) {
    var output mediaconvert.DisassociateCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) DisassociateCertificateAsync(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) *MediaconvertDisassociateCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateCertificate, input)
    return &MediaconvertDisassociateCertificateResult{Result: future}
}

func (a *MediaConvertStub) GetJob(ctx workflow.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
    var output mediaconvert.GetJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJob, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) GetJobAsync(ctx workflow.Context, input *mediaconvert.GetJobInput) *MediaconvertGetJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJob, input)
    return &MediaconvertGetJobResult{Result: future}
}

func (a *MediaConvertStub) GetJobTemplate(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
    var output mediaconvert.GetJobTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) GetJobTemplateAsync(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) *MediaconvertGetJobTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobTemplate, input)
    return &MediaconvertGetJobTemplateResult{Result: future}
}

func (a *MediaConvertStub) GetPreset(ctx workflow.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
    var output mediaconvert.GetPresetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPreset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) GetPresetAsync(ctx workflow.Context, input *mediaconvert.GetPresetInput) *MediaconvertGetPresetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPreset, input)
    return &MediaconvertGetPresetResult{Result: future}
}

func (a *MediaConvertStub) GetQueue(ctx workflow.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
    var output mediaconvert.GetQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) GetQueueAsync(ctx workflow.Context, input *mediaconvert.GetQueueInput) *MediaconvertGetQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueue, input)
    return &MediaconvertGetQueueResult{Result: future}
}

func (a *MediaConvertStub) ListJobTemplates(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
    var output mediaconvert.ListJobTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) ListJobTemplatesAsync(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) *MediaconvertListJobTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobTemplates, input)
    return &MediaconvertListJobTemplatesResult{Result: future}
}

func (a *MediaConvertStub) ListJobs(ctx workflow.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
    var output mediaconvert.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) ListJobsAsync(ctx workflow.Context, input *mediaconvert.ListJobsInput) *MediaconvertListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &MediaconvertListJobsResult{Result: future}
}

func (a *MediaConvertStub) ListPresets(ctx workflow.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
    var output mediaconvert.ListPresetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPresets, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) ListPresetsAsync(ctx workflow.Context, input *mediaconvert.ListPresetsInput) *MediaconvertListPresetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPresets, input)
    return &MediaconvertListPresetsResult{Result: future}
}

func (a *MediaConvertStub) ListQueues(ctx workflow.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
    var output mediaconvert.ListQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) ListQueuesAsync(ctx workflow.Context, input *mediaconvert.ListQueuesInput) *MediaconvertListQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input)
    return &MediaconvertListQueuesResult{Result: future}
}

func (a *MediaConvertStub) ListTagsForResource(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
    var output mediaconvert.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) *MediaconvertListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &MediaconvertListTagsForResourceResult{Result: future}
}

func (a *MediaConvertStub) TagResource(ctx workflow.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error) {
    var output mediaconvert.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) TagResourceAsync(ctx workflow.Context, input *mediaconvert.TagResourceInput) *MediaconvertTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &MediaconvertTagResourceResult{Result: future}
}

func (a *MediaConvertStub) UntagResource(ctx workflow.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error) {
    var output mediaconvert.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) UntagResourceAsync(ctx workflow.Context, input *mediaconvert.UntagResourceInput) *MediaconvertUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &MediaconvertUntagResourceResult{Result: future}
}

func (a *MediaConvertStub) UpdateJobTemplate(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error) {
    var output mediaconvert.UpdateJobTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJobTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) UpdateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) *MediaconvertUpdateJobTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJobTemplate, input)
    return &MediaconvertUpdateJobTemplateResult{Result: future}
}

func (a *MediaConvertStub) UpdatePreset(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error) {
    var output mediaconvert.UpdatePresetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePreset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) UpdatePresetAsync(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) *MediaconvertUpdatePresetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePreset, input)
    return &MediaconvertUpdatePresetResult{Result: future}
}

func (a *MediaConvertStub) UpdateQueue(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error) {
    var output mediaconvert.UpdateQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConvertStub) UpdateQueueAsync(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) *MediaconvertUpdateQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateQueue, input)
    return &MediaconvertUpdateQueueResult{Result: future}
}
