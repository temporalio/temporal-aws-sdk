package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoT1ClickDevicesServiceClient interface {
    ClaimDevicesByClaimCode(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error)
    ClaimDevicesByClaimCodeAsync(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) *Iot1clickdevicesserviceClaimDevicesByClaimCodeResult

    DescribeDevice(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error)
    DescribeDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) *Iot1clickdevicesserviceDescribeDeviceResult

    FinalizeDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error)
    FinalizeDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) *Iot1clickdevicesserviceFinalizeDeviceClaimResult

    GetDeviceMethods(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error)
    GetDeviceMethodsAsync(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) *Iot1clickdevicesserviceGetDeviceMethodsResult

    InitiateDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error)
    InitiateDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) *Iot1clickdevicesserviceInitiateDeviceClaimResult

    InvokeDeviceMethod(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error)
    InvokeDeviceMethodAsync(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) *Iot1clickdevicesserviceInvokeDeviceMethodResult

    ListDeviceEvents(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error)
    ListDeviceEventsAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) *Iot1clickdevicesserviceListDeviceEventsResult

    ListDevices(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error)
    ListDevicesAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) *Iot1clickdevicesserviceListDevicesResult

    ListTagsForResource(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) *Iot1clickdevicesserviceListTagsForResourceResult

    TagResource(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) *Iot1clickdevicesserviceTagResourceResult

    UnclaimDevice(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error)
    UnclaimDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) *Iot1clickdevicesserviceUnclaimDeviceResult

    UntagResource(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) *Iot1clickdevicesserviceUntagResourceResult

    UpdateDeviceState(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error)
    UpdateDeviceStateAsync(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) *Iot1clickdevicesserviceUpdateDeviceStateResult
}

type Iot1clickdevicesserviceClaimDevicesByClaimCodeResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceClaimDevicesByClaimCodeResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
    var output iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceDescribeDeviceResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceDescribeDeviceResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
    var output iot1clickdevicesservice.DescribeDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceFinalizeDeviceClaimResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceFinalizeDeviceClaimResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
    var output iot1clickdevicesservice.FinalizeDeviceClaimOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceGetDeviceMethodsResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceGetDeviceMethodsResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
    var output iot1clickdevicesservice.GetDeviceMethodsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceInitiateDeviceClaimResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceInitiateDeviceClaimResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
    var output iot1clickdevicesservice.InitiateDeviceClaimOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceInvokeDeviceMethodResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceInvokeDeviceMethodResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
    var output iot1clickdevicesservice.InvokeDeviceMethodOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceListDeviceEventsResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceListDeviceEventsResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
    var output iot1clickdevicesservice.ListDeviceEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceListDevicesResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceListDevicesResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListDevicesOutput, error) {
    var output iot1clickdevicesservice.ListDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceListTagsForResourceResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
    var output iot1clickdevicesservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceTagResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceTagResourceResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.TagResourceOutput, error) {
    var output iot1clickdevicesservice.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceUnclaimDeviceResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceUnclaimDeviceResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
    var output iot1clickdevicesservice.UnclaimDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceUntagResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceUntagResourceResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.UntagResourceOutput, error) {
    var output iot1clickdevicesservice.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickdevicesserviceUpdateDeviceStateResult struct {
	Result workflow.Future
}

func (r *Iot1clickdevicesserviceUpdateDeviceStateResult) Get(ctx workflow.Context) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
    var output iot1clickdevicesservice.UpdateDeviceStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IoT1ClickDevicesServiceStub struct {
    activities awsactivities.IoT1ClickDevicesServiceActivities
}

func NewIoT1ClickDevicesServiceStub() IoT1ClickDevicesServiceClient {
    return &IoT1ClickDevicesServiceStub{}
}

func (a *IoT1ClickDevicesServiceStub) ClaimDevicesByClaimCode(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
    var output iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ClaimDevicesByClaimCode, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ClaimDevicesByClaimCodeAsync(ctx workflow.Context, input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) *Iot1clickdevicesserviceClaimDevicesByClaimCodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ClaimDevicesByClaimCode, input)
    return &Iot1clickdevicesserviceClaimDevicesByClaimCodeResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) DescribeDevice(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
    var output iot1clickdevicesservice.DescribeDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) DescribeDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.DescribeDeviceInput) *Iot1clickdevicesserviceDescribeDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDevice, input)
    return &Iot1clickdevicesserviceDescribeDeviceResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) FinalizeDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
    var output iot1clickdevicesservice.FinalizeDeviceClaimOutput
    err := workflow.ExecuteActivity(ctx, a.activities.FinalizeDeviceClaim, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) FinalizeDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.FinalizeDeviceClaimInput) *Iot1clickdevicesserviceFinalizeDeviceClaimResult {
    future := workflow.ExecuteActivity(ctx, a.activities.FinalizeDeviceClaim, input)
    return &Iot1clickdevicesserviceFinalizeDeviceClaimResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) GetDeviceMethods(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
    var output iot1clickdevicesservice.GetDeviceMethodsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeviceMethods, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) GetDeviceMethodsAsync(ctx workflow.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput) *Iot1clickdevicesserviceGetDeviceMethodsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeviceMethods, input)
    return &Iot1clickdevicesserviceGetDeviceMethodsResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) InitiateDeviceClaim(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
    var output iot1clickdevicesservice.InitiateDeviceClaimOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateDeviceClaim, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) InitiateDeviceClaimAsync(ctx workflow.Context, input *iot1clickdevicesservice.InitiateDeviceClaimInput) *Iot1clickdevicesserviceInitiateDeviceClaimResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InitiateDeviceClaim, input)
    return &Iot1clickdevicesserviceInitiateDeviceClaimResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) InvokeDeviceMethod(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
    var output iot1clickdevicesservice.InvokeDeviceMethodOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InvokeDeviceMethod, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) InvokeDeviceMethodAsync(ctx workflow.Context, input *iot1clickdevicesservice.InvokeDeviceMethodInput) *Iot1clickdevicesserviceInvokeDeviceMethodResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InvokeDeviceMethod, input)
    return &Iot1clickdevicesserviceInvokeDeviceMethodResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) ListDeviceEvents(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
    var output iot1clickdevicesservice.ListDeviceEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeviceEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListDeviceEventsAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDeviceEventsInput) *Iot1clickdevicesserviceListDeviceEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeviceEvents, input)
    return &Iot1clickdevicesserviceListDeviceEventsResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) ListDevices(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error) {
    var output iot1clickdevicesservice.ListDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListDevicesAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListDevicesInput) *Iot1clickdevicesserviceListDevicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input)
    return &Iot1clickdevicesserviceListDevicesResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) ListTagsForResource(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
    var output iot1clickdevicesservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.ListTagsForResourceInput) *Iot1clickdevicesserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Iot1clickdevicesserviceListTagsForResourceResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) TagResource(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error) {
    var output iot1clickdevicesservice.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) TagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.TagResourceInput) *Iot1clickdevicesserviceTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Iot1clickdevicesserviceTagResourceResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) UnclaimDevice(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
    var output iot1clickdevicesservice.UnclaimDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnclaimDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UnclaimDeviceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UnclaimDeviceInput) *Iot1clickdevicesserviceUnclaimDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnclaimDevice, input)
    return &Iot1clickdevicesserviceUnclaimDeviceResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) UntagResource(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error) {
    var output iot1clickdevicesservice.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UntagResourceAsync(ctx workflow.Context, input *iot1clickdevicesservice.UntagResourceInput) *Iot1clickdevicesserviceUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Iot1clickdevicesserviceUntagResourceResult{Result: future}
}

func (a *IoT1ClickDevicesServiceStub) UpdateDeviceState(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
    var output iot1clickdevicesservice.UpdateDeviceStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceState, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickDevicesServiceStub) UpdateDeviceStateAsync(ctx workflow.Context, input *iot1clickdevicesservice.UpdateDeviceStateInput) *Iot1clickdevicesserviceUpdateDeviceStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceState, input)
    return &Iot1clickdevicesserviceUpdateDeviceStateResult{Result: future}
}