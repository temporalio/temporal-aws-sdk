package awsclients

import (
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type StorageGatewayClient interface {
    ActivateGateway(ctx workflow.Context, input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error)
    ActivateGatewayAsync(ctx workflow.Context, input *storagegateway.ActivateGatewayInput) *StoragegatewayActivateGatewayResult

    AddCache(ctx workflow.Context, input *storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error)
    AddCacheAsync(ctx workflow.Context, input *storagegateway.AddCacheInput) *StoragegatewayAddCacheResult

    AddTagsToResource(ctx workflow.Context, input *storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error)
    AddTagsToResourceAsync(ctx workflow.Context, input *storagegateway.AddTagsToResourceInput) *StoragegatewayAddTagsToResourceResult

    AddUploadBuffer(ctx workflow.Context, input *storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error)
    AddUploadBufferAsync(ctx workflow.Context, input *storagegateway.AddUploadBufferInput) *StoragegatewayAddUploadBufferResult

    AddWorkingStorage(ctx workflow.Context, input *storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error)
    AddWorkingStorageAsync(ctx workflow.Context, input *storagegateway.AddWorkingStorageInput) *StoragegatewayAddWorkingStorageResult

    AssignTapePool(ctx workflow.Context, input *storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error)
    AssignTapePoolAsync(ctx workflow.Context, input *storagegateway.AssignTapePoolInput) *StoragegatewayAssignTapePoolResult

    AttachVolume(ctx workflow.Context, input *storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error)
    AttachVolumeAsync(ctx workflow.Context, input *storagegateway.AttachVolumeInput) *StoragegatewayAttachVolumeResult

    CancelArchival(ctx workflow.Context, input *storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error)
    CancelArchivalAsync(ctx workflow.Context, input *storagegateway.CancelArchivalInput) *StoragegatewayCancelArchivalResult

    CancelRetrieval(ctx workflow.Context, input *storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error)
    CancelRetrievalAsync(ctx workflow.Context, input *storagegateway.CancelRetrievalInput) *StoragegatewayCancelRetrievalResult

    CreateCachediSCSIVolume(ctx workflow.Context, input *storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error)
    CreateCachediSCSIVolumeAsync(ctx workflow.Context, input *storagegateway.CreateCachediSCSIVolumeInput) *StoragegatewayCreateCachediSCSIVolumeResult

    CreateNFSFileShare(ctx workflow.Context, input *storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error)
    CreateNFSFileShareAsync(ctx workflow.Context, input *storagegateway.CreateNFSFileShareInput) *StoragegatewayCreateNFSFileShareResult

    CreateSMBFileShare(ctx workflow.Context, input *storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error)
    CreateSMBFileShareAsync(ctx workflow.Context, input *storagegateway.CreateSMBFileShareInput) *StoragegatewayCreateSMBFileShareResult

    CreateSnapshot(ctx workflow.Context, input *storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error)
    CreateSnapshotAsync(ctx workflow.Context, input *storagegateway.CreateSnapshotInput) *StoragegatewayCreateSnapshotResult

    CreateSnapshotFromVolumeRecoveryPoint(ctx workflow.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)
    CreateSnapshotFromVolumeRecoveryPointAsync(ctx workflow.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) *StoragegatewayCreateSnapshotFromVolumeRecoveryPointResult

    CreateStorediSCSIVolume(ctx workflow.Context, input *storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error)
    CreateStorediSCSIVolumeAsync(ctx workflow.Context, input *storagegateway.CreateStorediSCSIVolumeInput) *StoragegatewayCreateStorediSCSIVolumeResult

    CreateTapePool(ctx workflow.Context, input *storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error)
    CreateTapePoolAsync(ctx workflow.Context, input *storagegateway.CreateTapePoolInput) *StoragegatewayCreateTapePoolResult

    CreateTapeWithBarcode(ctx workflow.Context, input *storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error)
    CreateTapeWithBarcodeAsync(ctx workflow.Context, input *storagegateway.CreateTapeWithBarcodeInput) *StoragegatewayCreateTapeWithBarcodeResult

    CreateTapes(ctx workflow.Context, input *storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error)
    CreateTapesAsync(ctx workflow.Context, input *storagegateway.CreateTapesInput) *StoragegatewayCreateTapesResult

    DeleteAutomaticTapeCreationPolicy(ctx workflow.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error)
    DeleteAutomaticTapeCreationPolicyAsync(ctx workflow.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) *StoragegatewayDeleteAutomaticTapeCreationPolicyResult

    DeleteBandwidthRateLimit(ctx workflow.Context, input *storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error)
    DeleteBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.DeleteBandwidthRateLimitInput) *StoragegatewayDeleteBandwidthRateLimitResult

    DeleteChapCredentials(ctx workflow.Context, input *storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error)
    DeleteChapCredentialsAsync(ctx workflow.Context, input *storagegateway.DeleteChapCredentialsInput) *StoragegatewayDeleteChapCredentialsResult

    DeleteFileShare(ctx workflow.Context, input *storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error)
    DeleteFileShareAsync(ctx workflow.Context, input *storagegateway.DeleteFileShareInput) *StoragegatewayDeleteFileShareResult

    DeleteGateway(ctx workflow.Context, input *storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error)
    DeleteGatewayAsync(ctx workflow.Context, input *storagegateway.DeleteGatewayInput) *StoragegatewayDeleteGatewayResult

    DeleteSnapshotSchedule(ctx workflow.Context, input *storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error)
    DeleteSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.DeleteSnapshotScheduleInput) *StoragegatewayDeleteSnapshotScheduleResult

    DeleteTape(ctx workflow.Context, input *storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error)
    DeleteTapeAsync(ctx workflow.Context, input *storagegateway.DeleteTapeInput) *StoragegatewayDeleteTapeResult

    DeleteTapeArchive(ctx workflow.Context, input *storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error)
    DeleteTapeArchiveAsync(ctx workflow.Context, input *storagegateway.DeleteTapeArchiveInput) *StoragegatewayDeleteTapeArchiveResult

    DeleteTapePool(ctx workflow.Context, input *storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error)
    DeleteTapePoolAsync(ctx workflow.Context, input *storagegateway.DeleteTapePoolInput) *StoragegatewayDeleteTapePoolResult

    DeleteVolume(ctx workflow.Context, input *storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error)
    DeleteVolumeAsync(ctx workflow.Context, input *storagegateway.DeleteVolumeInput) *StoragegatewayDeleteVolumeResult

    DescribeAvailabilityMonitorTest(ctx workflow.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error)
    DescribeAvailabilityMonitorTestAsync(ctx workflow.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) *StoragegatewayDescribeAvailabilityMonitorTestResult

    DescribeBandwidthRateLimit(ctx workflow.Context, input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error)
    DescribeBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.DescribeBandwidthRateLimitInput) *StoragegatewayDescribeBandwidthRateLimitResult

    DescribeCache(ctx workflow.Context, input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error)
    DescribeCacheAsync(ctx workflow.Context, input *storagegateway.DescribeCacheInput) *StoragegatewayDescribeCacheResult

    DescribeCachediSCSIVolumes(ctx workflow.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)
    DescribeCachediSCSIVolumesAsync(ctx workflow.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) *StoragegatewayDescribeCachediSCSIVolumesResult

    DescribeChapCredentials(ctx workflow.Context, input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error)
    DescribeChapCredentialsAsync(ctx workflow.Context, input *storagegateway.DescribeChapCredentialsInput) *StoragegatewayDescribeChapCredentialsResult

    DescribeGatewayInformation(ctx workflow.Context, input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error)
    DescribeGatewayInformationAsync(ctx workflow.Context, input *storagegateway.DescribeGatewayInformationInput) *StoragegatewayDescribeGatewayInformationResult

    DescribeMaintenanceStartTime(ctx workflow.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)
    DescribeMaintenanceStartTimeAsync(ctx workflow.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) *StoragegatewayDescribeMaintenanceStartTimeResult

    DescribeNFSFileShares(ctx workflow.Context, input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error)
    DescribeNFSFileSharesAsync(ctx workflow.Context, input *storagegateway.DescribeNFSFileSharesInput) *StoragegatewayDescribeNFSFileSharesResult

    DescribeSMBFileShares(ctx workflow.Context, input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error)
    DescribeSMBFileSharesAsync(ctx workflow.Context, input *storagegateway.DescribeSMBFileSharesInput) *StoragegatewayDescribeSMBFileSharesResult

    DescribeSMBSettings(ctx workflow.Context, input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error)
    DescribeSMBSettingsAsync(ctx workflow.Context, input *storagegateway.DescribeSMBSettingsInput) *StoragegatewayDescribeSMBSettingsResult

    DescribeSnapshotSchedule(ctx workflow.Context, input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error)
    DescribeSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.DescribeSnapshotScheduleInput) *StoragegatewayDescribeSnapshotScheduleResult

    DescribeStorediSCSIVolumes(ctx workflow.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)
    DescribeStorediSCSIVolumesAsync(ctx workflow.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) *StoragegatewayDescribeStorediSCSIVolumesResult

    DescribeTapeArchives(ctx workflow.Context, input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error)
    DescribeTapeArchivesAsync(ctx workflow.Context, input *storagegateway.DescribeTapeArchivesInput) *StoragegatewayDescribeTapeArchivesResult

    DescribeTapeRecoveryPoints(ctx workflow.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)
    DescribeTapeRecoveryPointsAsync(ctx workflow.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) *StoragegatewayDescribeTapeRecoveryPointsResult

    DescribeTapes(ctx workflow.Context, input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error)
    DescribeTapesAsync(ctx workflow.Context, input *storagegateway.DescribeTapesInput) *StoragegatewayDescribeTapesResult

    DescribeUploadBuffer(ctx workflow.Context, input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error)
    DescribeUploadBufferAsync(ctx workflow.Context, input *storagegateway.DescribeUploadBufferInput) *StoragegatewayDescribeUploadBufferResult

    DescribeVTLDevices(ctx workflow.Context, input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error)
    DescribeVTLDevicesAsync(ctx workflow.Context, input *storagegateway.DescribeVTLDevicesInput) *StoragegatewayDescribeVTLDevicesResult

    DescribeWorkingStorage(ctx workflow.Context, input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error)
    DescribeWorkingStorageAsync(ctx workflow.Context, input *storagegateway.DescribeWorkingStorageInput) *StoragegatewayDescribeWorkingStorageResult

    DetachVolume(ctx workflow.Context, input *storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error)
    DetachVolumeAsync(ctx workflow.Context, input *storagegateway.DetachVolumeInput) *StoragegatewayDetachVolumeResult

    DisableGateway(ctx workflow.Context, input *storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error)
    DisableGatewayAsync(ctx workflow.Context, input *storagegateway.DisableGatewayInput) *StoragegatewayDisableGatewayResult

    JoinDomain(ctx workflow.Context, input *storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error)
    JoinDomainAsync(ctx workflow.Context, input *storagegateway.JoinDomainInput) *StoragegatewayJoinDomainResult

    ListAutomaticTapeCreationPolicies(ctx workflow.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error)
    ListAutomaticTapeCreationPoliciesAsync(ctx workflow.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) *StoragegatewayListAutomaticTapeCreationPoliciesResult

    ListFileShares(ctx workflow.Context, input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error)
    ListFileSharesAsync(ctx workflow.Context, input *storagegateway.ListFileSharesInput) *StoragegatewayListFileSharesResult

    ListGateways(ctx workflow.Context, input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error)
    ListGatewaysAsync(ctx workflow.Context, input *storagegateway.ListGatewaysInput) *StoragegatewayListGatewaysResult

    ListLocalDisks(ctx workflow.Context, input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error)
    ListLocalDisksAsync(ctx workflow.Context, input *storagegateway.ListLocalDisksInput) *StoragegatewayListLocalDisksResult

    ListTagsForResource(ctx workflow.Context, input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *storagegateway.ListTagsForResourceInput) *StoragegatewayListTagsForResourceResult

    ListTapePools(ctx workflow.Context, input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error)
    ListTapePoolsAsync(ctx workflow.Context, input *storagegateway.ListTapePoolsInput) *StoragegatewayListTapePoolsResult

    ListTapes(ctx workflow.Context, input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error)
    ListTapesAsync(ctx workflow.Context, input *storagegateway.ListTapesInput) *StoragegatewayListTapesResult

    ListVolumeInitiators(ctx workflow.Context, input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error)
    ListVolumeInitiatorsAsync(ctx workflow.Context, input *storagegateway.ListVolumeInitiatorsInput) *StoragegatewayListVolumeInitiatorsResult

    ListVolumeRecoveryPoints(ctx workflow.Context, input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error)
    ListVolumeRecoveryPointsAsync(ctx workflow.Context, input *storagegateway.ListVolumeRecoveryPointsInput) *StoragegatewayListVolumeRecoveryPointsResult

    ListVolumes(ctx workflow.Context, input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error)
    ListVolumesAsync(ctx workflow.Context, input *storagegateway.ListVolumesInput) *StoragegatewayListVolumesResult

    NotifyWhenUploaded(ctx workflow.Context, input *storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error)
    NotifyWhenUploadedAsync(ctx workflow.Context, input *storagegateway.NotifyWhenUploadedInput) *StoragegatewayNotifyWhenUploadedResult

    RefreshCache(ctx workflow.Context, input *storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error)
    RefreshCacheAsync(ctx workflow.Context, input *storagegateway.RefreshCacheInput) *StoragegatewayRefreshCacheResult

    RemoveTagsFromResource(ctx workflow.Context, input *storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error)
    RemoveTagsFromResourceAsync(ctx workflow.Context, input *storagegateway.RemoveTagsFromResourceInput) *StoragegatewayRemoveTagsFromResourceResult

    ResetCache(ctx workflow.Context, input *storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error)
    ResetCacheAsync(ctx workflow.Context, input *storagegateway.ResetCacheInput) *StoragegatewayResetCacheResult

    RetrieveTapeArchive(ctx workflow.Context, input *storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error)
    RetrieveTapeArchiveAsync(ctx workflow.Context, input *storagegateway.RetrieveTapeArchiveInput) *StoragegatewayRetrieveTapeArchiveResult

    RetrieveTapeRecoveryPoint(ctx workflow.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)
    RetrieveTapeRecoveryPointAsync(ctx workflow.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) *StoragegatewayRetrieveTapeRecoveryPointResult

    SetLocalConsolePassword(ctx workflow.Context, input *storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error)
    SetLocalConsolePasswordAsync(ctx workflow.Context, input *storagegateway.SetLocalConsolePasswordInput) *StoragegatewaySetLocalConsolePasswordResult

    SetSMBGuestPassword(ctx workflow.Context, input *storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error)
    SetSMBGuestPasswordAsync(ctx workflow.Context, input *storagegateway.SetSMBGuestPasswordInput) *StoragegatewaySetSMBGuestPasswordResult

    ShutdownGateway(ctx workflow.Context, input *storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error)
    ShutdownGatewayAsync(ctx workflow.Context, input *storagegateway.ShutdownGatewayInput) *StoragegatewayShutdownGatewayResult

    StartAvailabilityMonitorTest(ctx workflow.Context, input *storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error)
    StartAvailabilityMonitorTestAsync(ctx workflow.Context, input *storagegateway.StartAvailabilityMonitorTestInput) *StoragegatewayStartAvailabilityMonitorTestResult

    StartGateway(ctx workflow.Context, input *storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error)
    StartGatewayAsync(ctx workflow.Context, input *storagegateway.StartGatewayInput) *StoragegatewayStartGatewayResult

    UpdateAutomaticTapeCreationPolicy(ctx workflow.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error)
    UpdateAutomaticTapeCreationPolicyAsync(ctx workflow.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) *StoragegatewayUpdateAutomaticTapeCreationPolicyResult

    UpdateBandwidthRateLimit(ctx workflow.Context, input *storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error)
    UpdateBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.UpdateBandwidthRateLimitInput) *StoragegatewayUpdateBandwidthRateLimitResult

    UpdateChapCredentials(ctx workflow.Context, input *storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error)
    UpdateChapCredentialsAsync(ctx workflow.Context, input *storagegateway.UpdateChapCredentialsInput) *StoragegatewayUpdateChapCredentialsResult

    UpdateGatewayInformation(ctx workflow.Context, input *storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error)
    UpdateGatewayInformationAsync(ctx workflow.Context, input *storagegateway.UpdateGatewayInformationInput) *StoragegatewayUpdateGatewayInformationResult

    UpdateGatewaySoftwareNow(ctx workflow.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)
    UpdateGatewaySoftwareNowAsync(ctx workflow.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) *StoragegatewayUpdateGatewaySoftwareNowResult

    UpdateMaintenanceStartTime(ctx workflow.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)
    UpdateMaintenanceStartTimeAsync(ctx workflow.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) *StoragegatewayUpdateMaintenanceStartTimeResult

    UpdateNFSFileShare(ctx workflow.Context, input *storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error)
    UpdateNFSFileShareAsync(ctx workflow.Context, input *storagegateway.UpdateNFSFileShareInput) *StoragegatewayUpdateNFSFileShareResult

    UpdateSMBFileShare(ctx workflow.Context, input *storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error)
    UpdateSMBFileShareAsync(ctx workflow.Context, input *storagegateway.UpdateSMBFileShareInput) *StoragegatewayUpdateSMBFileShareResult

    UpdateSMBSecurityStrategy(ctx workflow.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error)
    UpdateSMBSecurityStrategyAsync(ctx workflow.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) *StoragegatewayUpdateSMBSecurityStrategyResult

    UpdateSnapshotSchedule(ctx workflow.Context, input *storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error)
    UpdateSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.UpdateSnapshotScheduleInput) *StoragegatewayUpdateSnapshotScheduleResult

    UpdateVTLDeviceType(ctx workflow.Context, input *storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
    UpdateVTLDeviceTypeAsync(ctx workflow.Context, input *storagegateway.UpdateVTLDeviceTypeInput) *StoragegatewayUpdateVTLDeviceTypeResult
}

type StoragegatewayActivateGatewayResult struct {
	Result workflow.Future
}

func (r *StoragegatewayActivateGatewayResult) Get(ctx workflow.Context) (*storagegateway.ActivateGatewayOutput, error) {
    var output storagegateway.ActivateGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAddCacheResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAddCacheResult) Get(ctx workflow.Context) (*storagegateway.AddCacheOutput, error) {
    var output storagegateway.AddCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAddTagsToResourceResult) Get(ctx workflow.Context) (*storagegateway.AddTagsToResourceOutput, error) {
    var output storagegateway.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAddUploadBufferResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAddUploadBufferResult) Get(ctx workflow.Context) (*storagegateway.AddUploadBufferOutput, error) {
    var output storagegateway.AddUploadBufferOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAddWorkingStorageResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAddWorkingStorageResult) Get(ctx workflow.Context) (*storagegateway.AddWorkingStorageOutput, error) {
    var output storagegateway.AddWorkingStorageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAssignTapePoolResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAssignTapePoolResult) Get(ctx workflow.Context) (*storagegateway.AssignTapePoolOutput, error) {
    var output storagegateway.AssignTapePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayAttachVolumeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayAttachVolumeResult) Get(ctx workflow.Context) (*storagegateway.AttachVolumeOutput, error) {
    var output storagegateway.AttachVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCancelArchivalResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCancelArchivalResult) Get(ctx workflow.Context) (*storagegateway.CancelArchivalOutput, error) {
    var output storagegateway.CancelArchivalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCancelRetrievalResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCancelRetrievalResult) Get(ctx workflow.Context) (*storagegateway.CancelRetrievalOutput, error) {
    var output storagegateway.CancelRetrievalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateCachediSCSIVolumeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateCachediSCSIVolumeResult) Get(ctx workflow.Context) (*storagegateway.CreateCachediSCSIVolumeOutput, error) {
    var output storagegateway.CreateCachediSCSIVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateNFSFileShareResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateNFSFileShareResult) Get(ctx workflow.Context) (*storagegateway.CreateNFSFileShareOutput, error) {
    var output storagegateway.CreateNFSFileShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateSMBFileShareResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateSMBFileShareResult) Get(ctx workflow.Context) (*storagegateway.CreateSMBFileShareOutput, error) {
    var output storagegateway.CreateSMBFileShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateSnapshotResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateSnapshotResult) Get(ctx workflow.Context) (*storagegateway.CreateSnapshotOutput, error) {
    var output storagegateway.CreateSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateSnapshotFromVolumeRecoveryPointResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateSnapshotFromVolumeRecoveryPointResult) Get(ctx workflow.Context) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error) {
    var output storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateStorediSCSIVolumeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateStorediSCSIVolumeResult) Get(ctx workflow.Context) (*storagegateway.CreateStorediSCSIVolumeOutput, error) {
    var output storagegateway.CreateStorediSCSIVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateTapePoolResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateTapePoolResult) Get(ctx workflow.Context) (*storagegateway.CreateTapePoolOutput, error) {
    var output storagegateway.CreateTapePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateTapeWithBarcodeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateTapeWithBarcodeResult) Get(ctx workflow.Context) (*storagegateway.CreateTapeWithBarcodeOutput, error) {
    var output storagegateway.CreateTapeWithBarcodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayCreateTapesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayCreateTapesResult) Get(ctx workflow.Context) (*storagegateway.CreateTapesOutput, error) {
    var output storagegateway.CreateTapesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteAutomaticTapeCreationPolicyResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteAutomaticTapeCreationPolicyResult) Get(ctx workflow.Context) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error) {
    var output storagegateway.DeleteAutomaticTapeCreationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteBandwidthRateLimitResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteBandwidthRateLimitResult) Get(ctx workflow.Context) (*storagegateway.DeleteBandwidthRateLimitOutput, error) {
    var output storagegateway.DeleteBandwidthRateLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteChapCredentialsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteChapCredentialsResult) Get(ctx workflow.Context) (*storagegateway.DeleteChapCredentialsOutput, error) {
    var output storagegateway.DeleteChapCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteFileShareResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteFileShareResult) Get(ctx workflow.Context) (*storagegateway.DeleteFileShareOutput, error) {
    var output storagegateway.DeleteFileShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteGatewayResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteGatewayResult) Get(ctx workflow.Context) (*storagegateway.DeleteGatewayOutput, error) {
    var output storagegateway.DeleteGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteSnapshotScheduleResult) Get(ctx workflow.Context) (*storagegateway.DeleteSnapshotScheduleOutput, error) {
    var output storagegateway.DeleteSnapshotScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteTapeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteTapeResult) Get(ctx workflow.Context) (*storagegateway.DeleteTapeOutput, error) {
    var output storagegateway.DeleteTapeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteTapeArchiveResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteTapeArchiveResult) Get(ctx workflow.Context) (*storagegateway.DeleteTapeArchiveOutput, error) {
    var output storagegateway.DeleteTapeArchiveOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteTapePoolResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteTapePoolResult) Get(ctx workflow.Context) (*storagegateway.DeleteTapePoolOutput, error) {
    var output storagegateway.DeleteTapePoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDeleteVolumeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDeleteVolumeResult) Get(ctx workflow.Context) (*storagegateway.DeleteVolumeOutput, error) {
    var output storagegateway.DeleteVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeAvailabilityMonitorTestResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeAvailabilityMonitorTestResult) Get(ctx workflow.Context) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
    var output storagegateway.DescribeAvailabilityMonitorTestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeBandwidthRateLimitResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeBandwidthRateLimitResult) Get(ctx workflow.Context) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
    var output storagegateway.DescribeBandwidthRateLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeCacheResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeCacheResult) Get(ctx workflow.Context) (*storagegateway.DescribeCacheOutput, error) {
    var output storagegateway.DescribeCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeCachediSCSIVolumesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeCachediSCSIVolumesResult) Get(ctx workflow.Context) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
    var output storagegateway.DescribeCachediSCSIVolumesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeChapCredentialsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeChapCredentialsResult) Get(ctx workflow.Context) (*storagegateway.DescribeChapCredentialsOutput, error) {
    var output storagegateway.DescribeChapCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeGatewayInformationResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeGatewayInformationResult) Get(ctx workflow.Context) (*storagegateway.DescribeGatewayInformationOutput, error) {
    var output storagegateway.DescribeGatewayInformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeMaintenanceStartTimeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeMaintenanceStartTimeResult) Get(ctx workflow.Context) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
    var output storagegateway.DescribeMaintenanceStartTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeNFSFileSharesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeNFSFileSharesResult) Get(ctx workflow.Context) (*storagegateway.DescribeNFSFileSharesOutput, error) {
    var output storagegateway.DescribeNFSFileSharesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeSMBFileSharesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeSMBFileSharesResult) Get(ctx workflow.Context) (*storagegateway.DescribeSMBFileSharesOutput, error) {
    var output storagegateway.DescribeSMBFileSharesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeSMBSettingsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeSMBSettingsResult) Get(ctx workflow.Context) (*storagegateway.DescribeSMBSettingsOutput, error) {
    var output storagegateway.DescribeSMBSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeSnapshotScheduleResult) Get(ctx workflow.Context) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
    var output storagegateway.DescribeSnapshotScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeStorediSCSIVolumesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeStorediSCSIVolumesResult) Get(ctx workflow.Context) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
    var output storagegateway.DescribeStorediSCSIVolumesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeTapeArchivesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeTapeArchivesResult) Get(ctx workflow.Context) (*storagegateway.DescribeTapeArchivesOutput, error) {
    var output storagegateway.DescribeTapeArchivesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeTapeRecoveryPointsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeTapeRecoveryPointsResult) Get(ctx workflow.Context) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
    var output storagegateway.DescribeTapeRecoveryPointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeTapesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeTapesResult) Get(ctx workflow.Context) (*storagegateway.DescribeTapesOutput, error) {
    var output storagegateway.DescribeTapesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeUploadBufferResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeUploadBufferResult) Get(ctx workflow.Context) (*storagegateway.DescribeUploadBufferOutput, error) {
    var output storagegateway.DescribeUploadBufferOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeVTLDevicesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeVTLDevicesResult) Get(ctx workflow.Context) (*storagegateway.DescribeVTLDevicesOutput, error) {
    var output storagegateway.DescribeVTLDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDescribeWorkingStorageResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDescribeWorkingStorageResult) Get(ctx workflow.Context) (*storagegateway.DescribeWorkingStorageOutput, error) {
    var output storagegateway.DescribeWorkingStorageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDetachVolumeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDetachVolumeResult) Get(ctx workflow.Context) (*storagegateway.DetachVolumeOutput, error) {
    var output storagegateway.DetachVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayDisableGatewayResult struct {
	Result workflow.Future
}

func (r *StoragegatewayDisableGatewayResult) Get(ctx workflow.Context) (*storagegateway.DisableGatewayOutput, error) {
    var output storagegateway.DisableGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayJoinDomainResult struct {
	Result workflow.Future
}

func (r *StoragegatewayJoinDomainResult) Get(ctx workflow.Context) (*storagegateway.JoinDomainOutput, error) {
    var output storagegateway.JoinDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListAutomaticTapeCreationPoliciesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListAutomaticTapeCreationPoliciesResult) Get(ctx workflow.Context) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
    var output storagegateway.ListAutomaticTapeCreationPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListFileSharesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListFileSharesResult) Get(ctx workflow.Context) (*storagegateway.ListFileSharesOutput, error) {
    var output storagegateway.ListFileSharesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListGatewaysResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListGatewaysResult) Get(ctx workflow.Context) (*storagegateway.ListGatewaysOutput, error) {
    var output storagegateway.ListGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListLocalDisksResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListLocalDisksResult) Get(ctx workflow.Context) (*storagegateway.ListLocalDisksOutput, error) {
    var output storagegateway.ListLocalDisksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListTagsForResourceResult) Get(ctx workflow.Context) (*storagegateway.ListTagsForResourceOutput, error) {
    var output storagegateway.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListTapePoolsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListTapePoolsResult) Get(ctx workflow.Context) (*storagegateway.ListTapePoolsOutput, error) {
    var output storagegateway.ListTapePoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListTapesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListTapesResult) Get(ctx workflow.Context) (*storagegateway.ListTapesOutput, error) {
    var output storagegateway.ListTapesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListVolumeInitiatorsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListVolumeInitiatorsResult) Get(ctx workflow.Context) (*storagegateway.ListVolumeInitiatorsOutput, error) {
    var output storagegateway.ListVolumeInitiatorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListVolumeRecoveryPointsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListVolumeRecoveryPointsResult) Get(ctx workflow.Context) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
    var output storagegateway.ListVolumeRecoveryPointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayListVolumesResult struct {
	Result workflow.Future
}

func (r *StoragegatewayListVolumesResult) Get(ctx workflow.Context) (*storagegateway.ListVolumesOutput, error) {
    var output storagegateway.ListVolumesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayNotifyWhenUploadedResult struct {
	Result workflow.Future
}

func (r *StoragegatewayNotifyWhenUploadedResult) Get(ctx workflow.Context) (*storagegateway.NotifyWhenUploadedOutput, error) {
    var output storagegateway.NotifyWhenUploadedOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayRefreshCacheResult struct {
	Result workflow.Future
}

func (r *StoragegatewayRefreshCacheResult) Get(ctx workflow.Context) (*storagegateway.RefreshCacheOutput, error) {
    var output storagegateway.RefreshCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *StoragegatewayRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*storagegateway.RemoveTagsFromResourceOutput, error) {
    var output storagegateway.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayResetCacheResult struct {
	Result workflow.Future
}

func (r *StoragegatewayResetCacheResult) Get(ctx workflow.Context) (*storagegateway.ResetCacheOutput, error) {
    var output storagegateway.ResetCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayRetrieveTapeArchiveResult struct {
	Result workflow.Future
}

func (r *StoragegatewayRetrieveTapeArchiveResult) Get(ctx workflow.Context) (*storagegateway.RetrieveTapeArchiveOutput, error) {
    var output storagegateway.RetrieveTapeArchiveOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayRetrieveTapeRecoveryPointResult struct {
	Result workflow.Future
}

func (r *StoragegatewayRetrieveTapeRecoveryPointResult) Get(ctx workflow.Context) (*storagegateway.RetrieveTapeRecoveryPointOutput, error) {
    var output storagegateway.RetrieveTapeRecoveryPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewaySetLocalConsolePasswordResult struct {
	Result workflow.Future
}

func (r *StoragegatewaySetLocalConsolePasswordResult) Get(ctx workflow.Context) (*storagegateway.SetLocalConsolePasswordOutput, error) {
    var output storagegateway.SetLocalConsolePasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewaySetSMBGuestPasswordResult struct {
	Result workflow.Future
}

func (r *StoragegatewaySetSMBGuestPasswordResult) Get(ctx workflow.Context) (*storagegateway.SetSMBGuestPasswordOutput, error) {
    var output storagegateway.SetSMBGuestPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayShutdownGatewayResult struct {
	Result workflow.Future
}

func (r *StoragegatewayShutdownGatewayResult) Get(ctx workflow.Context) (*storagegateway.ShutdownGatewayOutput, error) {
    var output storagegateway.ShutdownGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayStartAvailabilityMonitorTestResult struct {
	Result workflow.Future
}

func (r *StoragegatewayStartAvailabilityMonitorTestResult) Get(ctx workflow.Context) (*storagegateway.StartAvailabilityMonitorTestOutput, error) {
    var output storagegateway.StartAvailabilityMonitorTestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayStartGatewayResult struct {
	Result workflow.Future
}

func (r *StoragegatewayStartGatewayResult) Get(ctx workflow.Context) (*storagegateway.StartGatewayOutput, error) {
    var output storagegateway.StartGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateAutomaticTapeCreationPolicyResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateAutomaticTapeCreationPolicyResult) Get(ctx workflow.Context) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error) {
    var output storagegateway.UpdateAutomaticTapeCreationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateBandwidthRateLimitResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateBandwidthRateLimitResult) Get(ctx workflow.Context) (*storagegateway.UpdateBandwidthRateLimitOutput, error) {
    var output storagegateway.UpdateBandwidthRateLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateChapCredentialsResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateChapCredentialsResult) Get(ctx workflow.Context) (*storagegateway.UpdateChapCredentialsOutput, error) {
    var output storagegateway.UpdateChapCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateGatewayInformationResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateGatewayInformationResult) Get(ctx workflow.Context) (*storagegateway.UpdateGatewayInformationOutput, error) {
    var output storagegateway.UpdateGatewayInformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateGatewaySoftwareNowResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateGatewaySoftwareNowResult) Get(ctx workflow.Context) (*storagegateway.UpdateGatewaySoftwareNowOutput, error) {
    var output storagegateway.UpdateGatewaySoftwareNowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateMaintenanceStartTimeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateMaintenanceStartTimeResult) Get(ctx workflow.Context) (*storagegateway.UpdateMaintenanceStartTimeOutput, error) {
    var output storagegateway.UpdateMaintenanceStartTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateNFSFileShareResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateNFSFileShareResult) Get(ctx workflow.Context) (*storagegateway.UpdateNFSFileShareOutput, error) {
    var output storagegateway.UpdateNFSFileShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateSMBFileShareResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateSMBFileShareResult) Get(ctx workflow.Context) (*storagegateway.UpdateSMBFileShareOutput, error) {
    var output storagegateway.UpdateSMBFileShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateSMBSecurityStrategyResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateSMBSecurityStrategyResult) Get(ctx workflow.Context) (*storagegateway.UpdateSMBSecurityStrategyOutput, error) {
    var output storagegateway.UpdateSMBSecurityStrategyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateSnapshotScheduleResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateSnapshotScheduleResult) Get(ctx workflow.Context) (*storagegateway.UpdateSnapshotScheduleOutput, error) {
    var output storagegateway.UpdateSnapshotScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StoragegatewayUpdateVTLDeviceTypeResult struct {
	Result workflow.Future
}

func (r *StoragegatewayUpdateVTLDeviceTypeResult) Get(ctx workflow.Context) (*storagegateway.UpdateVTLDeviceTypeOutput, error) {
    var output storagegateway.UpdateVTLDeviceTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type StorageGatewayStub struct {
    activities awsactivities.StorageGatewayActivities
}

func NewStorageGatewayStub() StorageGatewayClient {
    return &StorageGatewayStub{}
}

func (a *StorageGatewayStub) ActivateGateway(ctx workflow.Context, input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error) {
    var output storagegateway.ActivateGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ActivateGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ActivateGatewayAsync(ctx workflow.Context, input *storagegateway.ActivateGatewayInput) *StoragegatewayActivateGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ActivateGateway, input)
    return &StoragegatewayActivateGatewayResult{Result: future}
}

func (a *StorageGatewayStub) AddCache(ctx workflow.Context, input *storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error) {
    var output storagegateway.AddCacheOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddCache, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AddCacheAsync(ctx workflow.Context, input *storagegateway.AddCacheInput) *StoragegatewayAddCacheResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddCache, input)
    return &StoragegatewayAddCacheResult{Result: future}
}

func (a *StorageGatewayStub) AddTagsToResource(ctx workflow.Context, input *storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error) {
    var output storagegateway.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AddTagsToResourceAsync(ctx workflow.Context, input *storagegateway.AddTagsToResourceInput) *StoragegatewayAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input)
    return &StoragegatewayAddTagsToResourceResult{Result: future}
}

func (a *StorageGatewayStub) AddUploadBuffer(ctx workflow.Context, input *storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error) {
    var output storagegateway.AddUploadBufferOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddUploadBuffer, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AddUploadBufferAsync(ctx workflow.Context, input *storagegateway.AddUploadBufferInput) *StoragegatewayAddUploadBufferResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddUploadBuffer, input)
    return &StoragegatewayAddUploadBufferResult{Result: future}
}

func (a *StorageGatewayStub) AddWorkingStorage(ctx workflow.Context, input *storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error) {
    var output storagegateway.AddWorkingStorageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddWorkingStorage, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AddWorkingStorageAsync(ctx workflow.Context, input *storagegateway.AddWorkingStorageInput) *StoragegatewayAddWorkingStorageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddWorkingStorage, input)
    return &StoragegatewayAddWorkingStorageResult{Result: future}
}

func (a *StorageGatewayStub) AssignTapePool(ctx workflow.Context, input *storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error) {
    var output storagegateway.AssignTapePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssignTapePool, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AssignTapePoolAsync(ctx workflow.Context, input *storagegateway.AssignTapePoolInput) *StoragegatewayAssignTapePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssignTapePool, input)
    return &StoragegatewayAssignTapePoolResult{Result: future}
}

func (a *StorageGatewayStub) AttachVolume(ctx workflow.Context, input *storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error) {
    var output storagegateway.AttachVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) AttachVolumeAsync(ctx workflow.Context, input *storagegateway.AttachVolumeInput) *StoragegatewayAttachVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachVolume, input)
    return &StoragegatewayAttachVolumeResult{Result: future}
}

func (a *StorageGatewayStub) CancelArchival(ctx workflow.Context, input *storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error) {
    var output storagegateway.CancelArchivalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelArchival, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CancelArchivalAsync(ctx workflow.Context, input *storagegateway.CancelArchivalInput) *StoragegatewayCancelArchivalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelArchival, input)
    return &StoragegatewayCancelArchivalResult{Result: future}
}

func (a *StorageGatewayStub) CancelRetrieval(ctx workflow.Context, input *storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error) {
    var output storagegateway.CancelRetrievalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelRetrieval, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CancelRetrievalAsync(ctx workflow.Context, input *storagegateway.CancelRetrievalInput) *StoragegatewayCancelRetrievalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelRetrieval, input)
    return &StoragegatewayCancelRetrievalResult{Result: future}
}

func (a *StorageGatewayStub) CreateCachediSCSIVolume(ctx workflow.Context, input *storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error) {
    var output storagegateway.CreateCachediSCSIVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCachediSCSIVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateCachediSCSIVolumeAsync(ctx workflow.Context, input *storagegateway.CreateCachediSCSIVolumeInput) *StoragegatewayCreateCachediSCSIVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCachediSCSIVolume, input)
    return &StoragegatewayCreateCachediSCSIVolumeResult{Result: future}
}

func (a *StorageGatewayStub) CreateNFSFileShare(ctx workflow.Context, input *storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error) {
    var output storagegateway.CreateNFSFileShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNFSFileShare, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateNFSFileShareAsync(ctx workflow.Context, input *storagegateway.CreateNFSFileShareInput) *StoragegatewayCreateNFSFileShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNFSFileShare, input)
    return &StoragegatewayCreateNFSFileShareResult{Result: future}
}

func (a *StorageGatewayStub) CreateSMBFileShare(ctx workflow.Context, input *storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error) {
    var output storagegateway.CreateSMBFileShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSMBFileShare, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateSMBFileShareAsync(ctx workflow.Context, input *storagegateway.CreateSMBFileShareInput) *StoragegatewayCreateSMBFileShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSMBFileShare, input)
    return &StoragegatewayCreateSMBFileShareResult{Result: future}
}

func (a *StorageGatewayStub) CreateSnapshot(ctx workflow.Context, input *storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error) {
    var output storagegateway.CreateSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateSnapshotAsync(ctx workflow.Context, input *storagegateway.CreateSnapshotInput) *StoragegatewayCreateSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshot, input)
    return &StoragegatewayCreateSnapshotResult{Result: future}
}

func (a *StorageGatewayStub) CreateSnapshotFromVolumeRecoveryPoint(ctx workflow.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error) {
    var output storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotFromVolumeRecoveryPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateSnapshotFromVolumeRecoveryPointAsync(ctx workflow.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) *StoragegatewayCreateSnapshotFromVolumeRecoveryPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshotFromVolumeRecoveryPoint, input)
    return &StoragegatewayCreateSnapshotFromVolumeRecoveryPointResult{Result: future}
}

func (a *StorageGatewayStub) CreateStorediSCSIVolume(ctx workflow.Context, input *storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error) {
    var output storagegateway.CreateStorediSCSIVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStorediSCSIVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateStorediSCSIVolumeAsync(ctx workflow.Context, input *storagegateway.CreateStorediSCSIVolumeInput) *StoragegatewayCreateStorediSCSIVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStorediSCSIVolume, input)
    return &StoragegatewayCreateStorediSCSIVolumeResult{Result: future}
}

func (a *StorageGatewayStub) CreateTapePool(ctx workflow.Context, input *storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error) {
    var output storagegateway.CreateTapePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTapePool, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateTapePoolAsync(ctx workflow.Context, input *storagegateway.CreateTapePoolInput) *StoragegatewayCreateTapePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTapePool, input)
    return &StoragegatewayCreateTapePoolResult{Result: future}
}

func (a *StorageGatewayStub) CreateTapeWithBarcode(ctx workflow.Context, input *storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error) {
    var output storagegateway.CreateTapeWithBarcodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTapeWithBarcode, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateTapeWithBarcodeAsync(ctx workflow.Context, input *storagegateway.CreateTapeWithBarcodeInput) *StoragegatewayCreateTapeWithBarcodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTapeWithBarcode, input)
    return &StoragegatewayCreateTapeWithBarcodeResult{Result: future}
}

func (a *StorageGatewayStub) CreateTapes(ctx workflow.Context, input *storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error) {
    var output storagegateway.CreateTapesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTapes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) CreateTapesAsync(ctx workflow.Context, input *storagegateway.CreateTapesInput) *StoragegatewayCreateTapesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTapes, input)
    return &StoragegatewayCreateTapesResult{Result: future}
}

func (a *StorageGatewayStub) DeleteAutomaticTapeCreationPolicy(ctx workflow.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error) {
    var output storagegateway.DeleteAutomaticTapeCreationPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAutomaticTapeCreationPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteAutomaticTapeCreationPolicyAsync(ctx workflow.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) *StoragegatewayDeleteAutomaticTapeCreationPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAutomaticTapeCreationPolicy, input)
    return &StoragegatewayDeleteAutomaticTapeCreationPolicyResult{Result: future}
}

func (a *StorageGatewayStub) DeleteBandwidthRateLimit(ctx workflow.Context, input *storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error) {
    var output storagegateway.DeleteBandwidthRateLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBandwidthRateLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.DeleteBandwidthRateLimitInput) *StoragegatewayDeleteBandwidthRateLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBandwidthRateLimit, input)
    return &StoragegatewayDeleteBandwidthRateLimitResult{Result: future}
}

func (a *StorageGatewayStub) DeleteChapCredentials(ctx workflow.Context, input *storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error) {
    var output storagegateway.DeleteChapCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteChapCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteChapCredentialsAsync(ctx workflow.Context, input *storagegateway.DeleteChapCredentialsInput) *StoragegatewayDeleteChapCredentialsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteChapCredentials, input)
    return &StoragegatewayDeleteChapCredentialsResult{Result: future}
}

func (a *StorageGatewayStub) DeleteFileShare(ctx workflow.Context, input *storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error) {
    var output storagegateway.DeleteFileShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFileShare, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteFileShareAsync(ctx workflow.Context, input *storagegateway.DeleteFileShareInput) *StoragegatewayDeleteFileShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFileShare, input)
    return &StoragegatewayDeleteFileShareResult{Result: future}
}

func (a *StorageGatewayStub) DeleteGateway(ctx workflow.Context, input *storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error) {
    var output storagegateway.DeleteGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteGatewayAsync(ctx workflow.Context, input *storagegateway.DeleteGatewayInput) *StoragegatewayDeleteGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGateway, input)
    return &StoragegatewayDeleteGatewayResult{Result: future}
}

func (a *StorageGatewayStub) DeleteSnapshotSchedule(ctx workflow.Context, input *storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error) {
    var output storagegateway.DeleteSnapshotScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.DeleteSnapshotScheduleInput) *StoragegatewayDeleteSnapshotScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshotSchedule, input)
    return &StoragegatewayDeleteSnapshotScheduleResult{Result: future}
}

func (a *StorageGatewayStub) DeleteTape(ctx workflow.Context, input *storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error) {
    var output storagegateway.DeleteTapeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTape, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteTapeAsync(ctx workflow.Context, input *storagegateway.DeleteTapeInput) *StoragegatewayDeleteTapeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTape, input)
    return &StoragegatewayDeleteTapeResult{Result: future}
}

func (a *StorageGatewayStub) DeleteTapeArchive(ctx workflow.Context, input *storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error) {
    var output storagegateway.DeleteTapeArchiveOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTapeArchive, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteTapeArchiveAsync(ctx workflow.Context, input *storagegateway.DeleteTapeArchiveInput) *StoragegatewayDeleteTapeArchiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTapeArchive, input)
    return &StoragegatewayDeleteTapeArchiveResult{Result: future}
}

func (a *StorageGatewayStub) DeleteTapePool(ctx workflow.Context, input *storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error) {
    var output storagegateway.DeleteTapePoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTapePool, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteTapePoolAsync(ctx workflow.Context, input *storagegateway.DeleteTapePoolInput) *StoragegatewayDeleteTapePoolResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTapePool, input)
    return &StoragegatewayDeleteTapePoolResult{Result: future}
}

func (a *StorageGatewayStub) DeleteVolume(ctx workflow.Context, input *storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error) {
    var output storagegateway.DeleteVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DeleteVolumeAsync(ctx workflow.Context, input *storagegateway.DeleteVolumeInput) *StoragegatewayDeleteVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVolume, input)
    return &StoragegatewayDeleteVolumeResult{Result: future}
}

func (a *StorageGatewayStub) DescribeAvailabilityMonitorTest(ctx workflow.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
    var output storagegateway.DescribeAvailabilityMonitorTestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailabilityMonitorTest, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeAvailabilityMonitorTestAsync(ctx workflow.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) *StoragegatewayDescribeAvailabilityMonitorTestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailabilityMonitorTest, input)
    return &StoragegatewayDescribeAvailabilityMonitorTestResult{Result: future}
}

func (a *StorageGatewayStub) DescribeBandwidthRateLimit(ctx workflow.Context, input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
    var output storagegateway.DescribeBandwidthRateLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBandwidthRateLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.DescribeBandwidthRateLimitInput) *StoragegatewayDescribeBandwidthRateLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBandwidthRateLimit, input)
    return &StoragegatewayDescribeBandwidthRateLimitResult{Result: future}
}

func (a *StorageGatewayStub) DescribeCache(ctx workflow.Context, input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error) {
    var output storagegateway.DescribeCacheOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCache, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeCacheAsync(ctx workflow.Context, input *storagegateway.DescribeCacheInput) *StoragegatewayDescribeCacheResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCache, input)
    return &StoragegatewayDescribeCacheResult{Result: future}
}

func (a *StorageGatewayStub) DescribeCachediSCSIVolumes(ctx workflow.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
    var output storagegateway.DescribeCachediSCSIVolumesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCachediSCSIVolumes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeCachediSCSIVolumesAsync(ctx workflow.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) *StoragegatewayDescribeCachediSCSIVolumesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCachediSCSIVolumes, input)
    return &StoragegatewayDescribeCachediSCSIVolumesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeChapCredentials(ctx workflow.Context, input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error) {
    var output storagegateway.DescribeChapCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeChapCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeChapCredentialsAsync(ctx workflow.Context, input *storagegateway.DescribeChapCredentialsInput) *StoragegatewayDescribeChapCredentialsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeChapCredentials, input)
    return &StoragegatewayDescribeChapCredentialsResult{Result: future}
}

func (a *StorageGatewayStub) DescribeGatewayInformation(ctx workflow.Context, input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error) {
    var output storagegateway.DescribeGatewayInformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGatewayInformation, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeGatewayInformationAsync(ctx workflow.Context, input *storagegateway.DescribeGatewayInformationInput) *StoragegatewayDescribeGatewayInformationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGatewayInformation, input)
    return &StoragegatewayDescribeGatewayInformationResult{Result: future}
}

func (a *StorageGatewayStub) DescribeMaintenanceStartTime(ctx workflow.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
    var output storagegateway.DescribeMaintenanceStartTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceStartTime, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeMaintenanceStartTimeAsync(ctx workflow.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) *StoragegatewayDescribeMaintenanceStartTimeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceStartTime, input)
    return &StoragegatewayDescribeMaintenanceStartTimeResult{Result: future}
}

func (a *StorageGatewayStub) DescribeNFSFileShares(ctx workflow.Context, input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error) {
    var output storagegateway.DescribeNFSFileSharesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNFSFileShares, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeNFSFileSharesAsync(ctx workflow.Context, input *storagegateway.DescribeNFSFileSharesInput) *StoragegatewayDescribeNFSFileSharesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNFSFileShares, input)
    return &StoragegatewayDescribeNFSFileSharesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeSMBFileShares(ctx workflow.Context, input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error) {
    var output storagegateway.DescribeSMBFileSharesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSMBFileShares, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeSMBFileSharesAsync(ctx workflow.Context, input *storagegateway.DescribeSMBFileSharesInput) *StoragegatewayDescribeSMBFileSharesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSMBFileShares, input)
    return &StoragegatewayDescribeSMBFileSharesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeSMBSettings(ctx workflow.Context, input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error) {
    var output storagegateway.DescribeSMBSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSMBSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeSMBSettingsAsync(ctx workflow.Context, input *storagegateway.DescribeSMBSettingsInput) *StoragegatewayDescribeSMBSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSMBSettings, input)
    return &StoragegatewayDescribeSMBSettingsResult{Result: future}
}

func (a *StorageGatewayStub) DescribeSnapshotSchedule(ctx workflow.Context, input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
    var output storagegateway.DescribeSnapshotScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.DescribeSnapshotScheduleInput) *StoragegatewayDescribeSnapshotScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotSchedule, input)
    return &StoragegatewayDescribeSnapshotScheduleResult{Result: future}
}

func (a *StorageGatewayStub) DescribeStorediSCSIVolumes(ctx workflow.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
    var output storagegateway.DescribeStorediSCSIVolumesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStorediSCSIVolumes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeStorediSCSIVolumesAsync(ctx workflow.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) *StoragegatewayDescribeStorediSCSIVolumesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStorediSCSIVolumes, input)
    return &StoragegatewayDescribeStorediSCSIVolumesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeTapeArchives(ctx workflow.Context, input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error) {
    var output storagegateway.DescribeTapeArchivesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTapeArchives, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeTapeArchivesAsync(ctx workflow.Context, input *storagegateway.DescribeTapeArchivesInput) *StoragegatewayDescribeTapeArchivesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTapeArchives, input)
    return &StoragegatewayDescribeTapeArchivesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeTapeRecoveryPoints(ctx workflow.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
    var output storagegateway.DescribeTapeRecoveryPointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTapeRecoveryPoints, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeTapeRecoveryPointsAsync(ctx workflow.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) *StoragegatewayDescribeTapeRecoveryPointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTapeRecoveryPoints, input)
    return &StoragegatewayDescribeTapeRecoveryPointsResult{Result: future}
}

func (a *StorageGatewayStub) DescribeTapes(ctx workflow.Context, input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error) {
    var output storagegateway.DescribeTapesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTapes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeTapesAsync(ctx workflow.Context, input *storagegateway.DescribeTapesInput) *StoragegatewayDescribeTapesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTapes, input)
    return &StoragegatewayDescribeTapesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeUploadBuffer(ctx workflow.Context, input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error) {
    var output storagegateway.DescribeUploadBufferOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUploadBuffer, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeUploadBufferAsync(ctx workflow.Context, input *storagegateway.DescribeUploadBufferInput) *StoragegatewayDescribeUploadBufferResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUploadBuffer, input)
    return &StoragegatewayDescribeUploadBufferResult{Result: future}
}

func (a *StorageGatewayStub) DescribeVTLDevices(ctx workflow.Context, input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error) {
    var output storagegateway.DescribeVTLDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVTLDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeVTLDevicesAsync(ctx workflow.Context, input *storagegateway.DescribeVTLDevicesInput) *StoragegatewayDescribeVTLDevicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVTLDevices, input)
    return &StoragegatewayDescribeVTLDevicesResult{Result: future}
}

func (a *StorageGatewayStub) DescribeWorkingStorage(ctx workflow.Context, input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error) {
    var output storagegateway.DescribeWorkingStorageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkingStorage, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DescribeWorkingStorageAsync(ctx workflow.Context, input *storagegateway.DescribeWorkingStorageInput) *StoragegatewayDescribeWorkingStorageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkingStorage, input)
    return &StoragegatewayDescribeWorkingStorageResult{Result: future}
}

func (a *StorageGatewayStub) DetachVolume(ctx workflow.Context, input *storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error) {
    var output storagegateway.DetachVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DetachVolumeAsync(ctx workflow.Context, input *storagegateway.DetachVolumeInput) *StoragegatewayDetachVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachVolume, input)
    return &StoragegatewayDetachVolumeResult{Result: future}
}

func (a *StorageGatewayStub) DisableGateway(ctx workflow.Context, input *storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error) {
    var output storagegateway.DisableGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) DisableGatewayAsync(ctx workflow.Context, input *storagegateway.DisableGatewayInput) *StoragegatewayDisableGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableGateway, input)
    return &StoragegatewayDisableGatewayResult{Result: future}
}

func (a *StorageGatewayStub) JoinDomain(ctx workflow.Context, input *storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error) {
    var output storagegateway.JoinDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.JoinDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) JoinDomainAsync(ctx workflow.Context, input *storagegateway.JoinDomainInput) *StoragegatewayJoinDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.JoinDomain, input)
    return &StoragegatewayJoinDomainResult{Result: future}
}

func (a *StorageGatewayStub) ListAutomaticTapeCreationPolicies(ctx workflow.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
    var output storagegateway.ListAutomaticTapeCreationPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAutomaticTapeCreationPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListAutomaticTapeCreationPoliciesAsync(ctx workflow.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) *StoragegatewayListAutomaticTapeCreationPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAutomaticTapeCreationPolicies, input)
    return &StoragegatewayListAutomaticTapeCreationPoliciesResult{Result: future}
}

func (a *StorageGatewayStub) ListFileShares(ctx workflow.Context, input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error) {
    var output storagegateway.ListFileSharesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFileShares, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListFileSharesAsync(ctx workflow.Context, input *storagegateway.ListFileSharesInput) *StoragegatewayListFileSharesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFileShares, input)
    return &StoragegatewayListFileSharesResult{Result: future}
}

func (a *StorageGatewayStub) ListGateways(ctx workflow.Context, input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error) {
    var output storagegateway.ListGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListGatewaysAsync(ctx workflow.Context, input *storagegateway.ListGatewaysInput) *StoragegatewayListGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input)
    return &StoragegatewayListGatewaysResult{Result: future}
}

func (a *StorageGatewayStub) ListLocalDisks(ctx workflow.Context, input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error) {
    var output storagegateway.ListLocalDisksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLocalDisks, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListLocalDisksAsync(ctx workflow.Context, input *storagegateway.ListLocalDisksInput) *StoragegatewayListLocalDisksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLocalDisks, input)
    return &StoragegatewayListLocalDisksResult{Result: future}
}

func (a *StorageGatewayStub) ListTagsForResource(ctx workflow.Context, input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error) {
    var output storagegateway.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListTagsForResourceAsync(ctx workflow.Context, input *storagegateway.ListTagsForResourceInput) *StoragegatewayListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &StoragegatewayListTagsForResourceResult{Result: future}
}

func (a *StorageGatewayStub) ListTapePools(ctx workflow.Context, input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error) {
    var output storagegateway.ListTapePoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTapePools, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListTapePoolsAsync(ctx workflow.Context, input *storagegateway.ListTapePoolsInput) *StoragegatewayListTapePoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTapePools, input)
    return &StoragegatewayListTapePoolsResult{Result: future}
}

func (a *StorageGatewayStub) ListTapes(ctx workflow.Context, input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error) {
    var output storagegateway.ListTapesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTapes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListTapesAsync(ctx workflow.Context, input *storagegateway.ListTapesInput) *StoragegatewayListTapesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTapes, input)
    return &StoragegatewayListTapesResult{Result: future}
}

func (a *StorageGatewayStub) ListVolumeInitiators(ctx workflow.Context, input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error) {
    var output storagegateway.ListVolumeInitiatorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVolumeInitiators, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListVolumeInitiatorsAsync(ctx workflow.Context, input *storagegateway.ListVolumeInitiatorsInput) *StoragegatewayListVolumeInitiatorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVolumeInitiators, input)
    return &StoragegatewayListVolumeInitiatorsResult{Result: future}
}

func (a *StorageGatewayStub) ListVolumeRecoveryPoints(ctx workflow.Context, input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
    var output storagegateway.ListVolumeRecoveryPointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVolumeRecoveryPoints, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListVolumeRecoveryPointsAsync(ctx workflow.Context, input *storagegateway.ListVolumeRecoveryPointsInput) *StoragegatewayListVolumeRecoveryPointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVolumeRecoveryPoints, input)
    return &StoragegatewayListVolumeRecoveryPointsResult{Result: future}
}

func (a *StorageGatewayStub) ListVolumes(ctx workflow.Context, input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error) {
    var output storagegateway.ListVolumesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVolumes, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ListVolumesAsync(ctx workflow.Context, input *storagegateway.ListVolumesInput) *StoragegatewayListVolumesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVolumes, input)
    return &StoragegatewayListVolumesResult{Result: future}
}

func (a *StorageGatewayStub) NotifyWhenUploaded(ctx workflow.Context, input *storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error) {
    var output storagegateway.NotifyWhenUploadedOutput
    err := workflow.ExecuteActivity(ctx, a.activities.NotifyWhenUploaded, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) NotifyWhenUploadedAsync(ctx workflow.Context, input *storagegateway.NotifyWhenUploadedInput) *StoragegatewayNotifyWhenUploadedResult {
    future := workflow.ExecuteActivity(ctx, a.activities.NotifyWhenUploaded, input)
    return &StoragegatewayNotifyWhenUploadedResult{Result: future}
}

func (a *StorageGatewayStub) RefreshCache(ctx workflow.Context, input *storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error) {
    var output storagegateway.RefreshCacheOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RefreshCache, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) RefreshCacheAsync(ctx workflow.Context, input *storagegateway.RefreshCacheInput) *StoragegatewayRefreshCacheResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RefreshCache, input)
    return &StoragegatewayRefreshCacheResult{Result: future}
}

func (a *StorageGatewayStub) RemoveTagsFromResource(ctx workflow.Context, input *storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error) {
    var output storagegateway.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *storagegateway.RemoveTagsFromResourceInput) *StoragegatewayRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input)
    return &StoragegatewayRemoveTagsFromResourceResult{Result: future}
}

func (a *StorageGatewayStub) ResetCache(ctx workflow.Context, input *storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error) {
    var output storagegateway.ResetCacheOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetCache, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ResetCacheAsync(ctx workflow.Context, input *storagegateway.ResetCacheInput) *StoragegatewayResetCacheResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetCache, input)
    return &StoragegatewayResetCacheResult{Result: future}
}

func (a *StorageGatewayStub) RetrieveTapeArchive(ctx workflow.Context, input *storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error) {
    var output storagegateway.RetrieveTapeArchiveOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RetrieveTapeArchive, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) RetrieveTapeArchiveAsync(ctx workflow.Context, input *storagegateway.RetrieveTapeArchiveInput) *StoragegatewayRetrieveTapeArchiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RetrieveTapeArchive, input)
    return &StoragegatewayRetrieveTapeArchiveResult{Result: future}
}

func (a *StorageGatewayStub) RetrieveTapeRecoveryPoint(ctx workflow.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error) {
    var output storagegateway.RetrieveTapeRecoveryPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RetrieveTapeRecoveryPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) RetrieveTapeRecoveryPointAsync(ctx workflow.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) *StoragegatewayRetrieveTapeRecoveryPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RetrieveTapeRecoveryPoint, input)
    return &StoragegatewayRetrieveTapeRecoveryPointResult{Result: future}
}

func (a *StorageGatewayStub) SetLocalConsolePassword(ctx workflow.Context, input *storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error) {
    var output storagegateway.SetLocalConsolePasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetLocalConsolePassword, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) SetLocalConsolePasswordAsync(ctx workflow.Context, input *storagegateway.SetLocalConsolePasswordInput) *StoragegatewaySetLocalConsolePasswordResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetLocalConsolePassword, input)
    return &StoragegatewaySetLocalConsolePasswordResult{Result: future}
}

func (a *StorageGatewayStub) SetSMBGuestPassword(ctx workflow.Context, input *storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error) {
    var output storagegateway.SetSMBGuestPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetSMBGuestPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) SetSMBGuestPasswordAsync(ctx workflow.Context, input *storagegateway.SetSMBGuestPasswordInput) *StoragegatewaySetSMBGuestPasswordResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetSMBGuestPassword, input)
    return &StoragegatewaySetSMBGuestPasswordResult{Result: future}
}

func (a *StorageGatewayStub) ShutdownGateway(ctx workflow.Context, input *storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error) {
    var output storagegateway.ShutdownGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ShutdownGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) ShutdownGatewayAsync(ctx workflow.Context, input *storagegateway.ShutdownGatewayInput) *StoragegatewayShutdownGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ShutdownGateway, input)
    return &StoragegatewayShutdownGatewayResult{Result: future}
}

func (a *StorageGatewayStub) StartAvailabilityMonitorTest(ctx workflow.Context, input *storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error) {
    var output storagegateway.StartAvailabilityMonitorTestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAvailabilityMonitorTest, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) StartAvailabilityMonitorTestAsync(ctx workflow.Context, input *storagegateway.StartAvailabilityMonitorTestInput) *StoragegatewayStartAvailabilityMonitorTestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartAvailabilityMonitorTest, input)
    return &StoragegatewayStartAvailabilityMonitorTestResult{Result: future}
}

func (a *StorageGatewayStub) StartGateway(ctx workflow.Context, input *storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error) {
    var output storagegateway.StartGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) StartGatewayAsync(ctx workflow.Context, input *storagegateway.StartGatewayInput) *StoragegatewayStartGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartGateway, input)
    return &StoragegatewayStartGatewayResult{Result: future}
}

func (a *StorageGatewayStub) UpdateAutomaticTapeCreationPolicy(ctx workflow.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error) {
    var output storagegateway.UpdateAutomaticTapeCreationPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAutomaticTapeCreationPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateAutomaticTapeCreationPolicyAsync(ctx workflow.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) *StoragegatewayUpdateAutomaticTapeCreationPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAutomaticTapeCreationPolicy, input)
    return &StoragegatewayUpdateAutomaticTapeCreationPolicyResult{Result: future}
}

func (a *StorageGatewayStub) UpdateBandwidthRateLimit(ctx workflow.Context, input *storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error) {
    var output storagegateway.UpdateBandwidthRateLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBandwidthRateLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateBandwidthRateLimitAsync(ctx workflow.Context, input *storagegateway.UpdateBandwidthRateLimitInput) *StoragegatewayUpdateBandwidthRateLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBandwidthRateLimit, input)
    return &StoragegatewayUpdateBandwidthRateLimitResult{Result: future}
}

func (a *StorageGatewayStub) UpdateChapCredentials(ctx workflow.Context, input *storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error) {
    var output storagegateway.UpdateChapCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateChapCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateChapCredentialsAsync(ctx workflow.Context, input *storagegateway.UpdateChapCredentialsInput) *StoragegatewayUpdateChapCredentialsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateChapCredentials, input)
    return &StoragegatewayUpdateChapCredentialsResult{Result: future}
}

func (a *StorageGatewayStub) UpdateGatewayInformation(ctx workflow.Context, input *storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error) {
    var output storagegateway.UpdateGatewayInformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayInformation, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateGatewayInformationAsync(ctx workflow.Context, input *storagegateway.UpdateGatewayInformationInput) *StoragegatewayUpdateGatewayInformationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayInformation, input)
    return &StoragegatewayUpdateGatewayInformationResult{Result: future}
}

func (a *StorageGatewayStub) UpdateGatewaySoftwareNow(ctx workflow.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error) {
    var output storagegateway.UpdateGatewaySoftwareNowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewaySoftwareNow, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateGatewaySoftwareNowAsync(ctx workflow.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) *StoragegatewayUpdateGatewaySoftwareNowResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewaySoftwareNow, input)
    return &StoragegatewayUpdateGatewaySoftwareNowResult{Result: future}
}

func (a *StorageGatewayStub) UpdateMaintenanceStartTime(ctx workflow.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error) {
    var output storagegateway.UpdateMaintenanceStartTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMaintenanceStartTime, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateMaintenanceStartTimeAsync(ctx workflow.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) *StoragegatewayUpdateMaintenanceStartTimeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMaintenanceStartTime, input)
    return &StoragegatewayUpdateMaintenanceStartTimeResult{Result: future}
}

func (a *StorageGatewayStub) UpdateNFSFileShare(ctx workflow.Context, input *storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error) {
    var output storagegateway.UpdateNFSFileShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNFSFileShare, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateNFSFileShareAsync(ctx workflow.Context, input *storagegateway.UpdateNFSFileShareInput) *StoragegatewayUpdateNFSFileShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateNFSFileShare, input)
    return &StoragegatewayUpdateNFSFileShareResult{Result: future}
}

func (a *StorageGatewayStub) UpdateSMBFileShare(ctx workflow.Context, input *storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error) {
    var output storagegateway.UpdateSMBFileShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSMBFileShare, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateSMBFileShareAsync(ctx workflow.Context, input *storagegateway.UpdateSMBFileShareInput) *StoragegatewayUpdateSMBFileShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSMBFileShare, input)
    return &StoragegatewayUpdateSMBFileShareResult{Result: future}
}

func (a *StorageGatewayStub) UpdateSMBSecurityStrategy(ctx workflow.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error) {
    var output storagegateway.UpdateSMBSecurityStrategyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSMBSecurityStrategy, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateSMBSecurityStrategyAsync(ctx workflow.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) *StoragegatewayUpdateSMBSecurityStrategyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSMBSecurityStrategy, input)
    return &StoragegatewayUpdateSMBSecurityStrategyResult{Result: future}
}

func (a *StorageGatewayStub) UpdateSnapshotSchedule(ctx workflow.Context, input *storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error) {
    var output storagegateway.UpdateSnapshotScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSnapshotSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateSnapshotScheduleAsync(ctx workflow.Context, input *storagegateway.UpdateSnapshotScheduleInput) *StoragegatewayUpdateSnapshotScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSnapshotSchedule, input)
    return &StoragegatewayUpdateSnapshotScheduleResult{Result: future}
}

func (a *StorageGatewayStub) UpdateVTLDeviceType(ctx workflow.Context, input *storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error) {
    var output storagegateway.UpdateVTLDeviceTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVTLDeviceType, input).Get(ctx, &output)
    return &output, err
}

func (a *StorageGatewayStub) UpdateVTLDeviceTypeAsync(ctx workflow.Context, input *storagegateway.UpdateVTLDeviceTypeInput) *StoragegatewayUpdateVTLDeviceTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVTLDeviceType, input)
    return &StoragegatewayUpdateVTLDeviceTypeResult{Result: future}
}