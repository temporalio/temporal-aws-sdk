// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package lightsailstub

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AllocateStaticIp(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpAsync(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) *LightsailAllocateStaticIpFuture

	AttachCertificateToDistribution(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error)
	AttachCertificateToDistributionAsync(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) *LightsailAttachCertificateToDistributionFuture

	AttachDisk(ctx workflow.Context, input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error)
	AttachDiskAsync(ctx workflow.Context, input *lightsail.AttachDiskInput) *LightsailAttachDiskFuture

	AttachInstancesToLoadBalancer(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
	AttachInstancesToLoadBalancerAsync(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) *LightsailAttachInstancesToLoadBalancerFuture

	AttachLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
	AttachLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) *LightsailAttachLoadBalancerTlsCertificateFuture

	AttachStaticIp(ctx workflow.Context, input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpAsync(ctx workflow.Context, input *lightsail.AttachStaticIpInput) *LightsailAttachStaticIpFuture

	CloseInstancePublicPorts(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) *LightsailCloseInstancePublicPortsFuture

	CopySnapshot(ctx workflow.Context, input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error)
	CopySnapshotAsync(ctx workflow.Context, input *lightsail.CopySnapshotInput) *LightsailCopySnapshotFuture

	CreateCertificate(ctx workflow.Context, input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error)
	CreateCertificateAsync(ctx workflow.Context, input *lightsail.CreateCertificateInput) *LightsailCreateCertificateFuture

	CreateCloudFormationStack(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error)
	CreateCloudFormationStackAsync(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) *LightsailCreateCloudFormationStackFuture

	CreateContactMethod(ctx workflow.Context, input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error)
	CreateContactMethodAsync(ctx workflow.Context, input *lightsail.CreateContactMethodInput) *LightsailCreateContactMethodFuture

	CreateDisk(ctx workflow.Context, input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error)
	CreateDiskAsync(ctx workflow.Context, input *lightsail.CreateDiskInput) *LightsailCreateDiskFuture

	CreateDiskFromSnapshot(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error)
	CreateDiskFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) *LightsailCreateDiskFromSnapshotFuture

	CreateDiskSnapshot(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error)
	CreateDiskSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) *LightsailCreateDiskSnapshotFuture

	CreateDistribution(ctx workflow.Context, input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error)
	CreateDistributionAsync(ctx workflow.Context, input *lightsail.CreateDistributionInput) *LightsailCreateDistributionFuture

	CreateDomain(ctx workflow.Context, input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *lightsail.CreateDomainInput) *LightsailCreateDomainFuture

	CreateDomainEntry(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryAsync(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) *LightsailCreateDomainEntryFuture

	CreateInstanceSnapshot(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) *LightsailCreateInstanceSnapshotFuture

	CreateInstances(ctx workflow.Context, input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesAsync(ctx workflow.Context, input *lightsail.CreateInstancesInput) *LightsailCreateInstancesFuture

	CreateInstancesFromSnapshot(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) *LightsailCreateInstancesFromSnapshotFuture

	CreateKeyPair(ctx workflow.Context, input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairAsync(ctx workflow.Context, input *lightsail.CreateKeyPairInput) *LightsailCreateKeyPairFuture

	CreateLoadBalancer(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) *LightsailCreateLoadBalancerFuture

	CreateLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
	CreateLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) *LightsailCreateLoadBalancerTlsCertificateFuture

	CreateRelationalDatabase(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error)
	CreateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) *LightsailCreateRelationalDatabaseFuture

	CreateRelationalDatabaseFromSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
	CreateRelationalDatabaseFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) *LightsailCreateRelationalDatabaseFromSnapshotFuture

	CreateRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
	CreateRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) *LightsailCreateRelationalDatabaseSnapshotFuture

	DeleteAlarm(ctx workflow.Context, input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error)
	DeleteAlarmAsync(ctx workflow.Context, input *lightsail.DeleteAlarmInput) *LightsailDeleteAlarmFuture

	DeleteAutoSnapshot(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error)
	DeleteAutoSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) *LightsailDeleteAutoSnapshotFuture

	DeleteCertificate(ctx workflow.Context, input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error)
	DeleteCertificateAsync(ctx workflow.Context, input *lightsail.DeleteCertificateInput) *LightsailDeleteCertificateFuture

	DeleteContactMethod(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error)
	DeleteContactMethodAsync(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) *LightsailDeleteContactMethodFuture

	DeleteDisk(ctx workflow.Context, input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error)
	DeleteDiskAsync(ctx workflow.Context, input *lightsail.DeleteDiskInput) *LightsailDeleteDiskFuture

	DeleteDiskSnapshot(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error)
	DeleteDiskSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) *LightsailDeleteDiskSnapshotFuture

	DeleteDistribution(ctx workflow.Context, input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error)
	DeleteDistributionAsync(ctx workflow.Context, input *lightsail.DeleteDistributionInput) *LightsailDeleteDistributionFuture

	DeleteDomain(ctx workflow.Context, input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *lightsail.DeleteDomainInput) *LightsailDeleteDomainFuture

	DeleteDomainEntry(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryAsync(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) *LightsailDeleteDomainEntryFuture

	DeleteInstance(ctx workflow.Context, input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceAsync(ctx workflow.Context, input *lightsail.DeleteInstanceInput) *LightsailDeleteInstanceFuture

	DeleteInstanceSnapshot(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) *LightsailDeleteInstanceSnapshotFuture

	DeleteKeyPair(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairAsync(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) *LightsailDeleteKeyPairFuture

	DeleteKnownHostKeys(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error)
	DeleteKnownHostKeysAsync(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) *LightsailDeleteKnownHostKeysFuture

	DeleteLoadBalancer(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) *LightsailDeleteLoadBalancerFuture

	DeleteLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
	DeleteLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) *LightsailDeleteLoadBalancerTlsCertificateFuture

	DeleteRelationalDatabase(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error)
	DeleteRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) *LightsailDeleteRelationalDatabaseFuture

	DeleteRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
	DeleteRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) *LightsailDeleteRelationalDatabaseSnapshotFuture

	DetachCertificateFromDistribution(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error)
	DetachCertificateFromDistributionAsync(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) *LightsailDetachCertificateFromDistributionFuture

	DetachDisk(ctx workflow.Context, input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error)
	DetachDiskAsync(ctx workflow.Context, input *lightsail.DetachDiskInput) *LightsailDetachDiskFuture

	DetachInstancesFromLoadBalancer(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
	DetachInstancesFromLoadBalancerAsync(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) *LightsailDetachInstancesFromLoadBalancerFuture

	DetachStaticIp(ctx workflow.Context, input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpAsync(ctx workflow.Context, input *lightsail.DetachStaticIpInput) *LightsailDetachStaticIpFuture

	DisableAddOn(ctx workflow.Context, input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error)
	DisableAddOnAsync(ctx workflow.Context, input *lightsail.DisableAddOnInput) *LightsailDisableAddOnFuture

	DownloadDefaultKeyPair(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairAsync(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) *LightsailDownloadDefaultKeyPairFuture

	EnableAddOn(ctx workflow.Context, input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error)
	EnableAddOnAsync(ctx workflow.Context, input *lightsail.EnableAddOnInput) *LightsailEnableAddOnFuture

	ExportSnapshot(ctx workflow.Context, input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error)
	ExportSnapshotAsync(ctx workflow.Context, input *lightsail.ExportSnapshotInput) *LightsailExportSnapshotFuture

	GetActiveNames(ctx workflow.Context, input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesAsync(ctx workflow.Context, input *lightsail.GetActiveNamesInput) *LightsailGetActiveNamesFuture

	GetAlarms(ctx workflow.Context, input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error)
	GetAlarmsAsync(ctx workflow.Context, input *lightsail.GetAlarmsInput) *LightsailGetAlarmsFuture

	GetAutoSnapshots(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error)
	GetAutoSnapshotsAsync(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) *LightsailGetAutoSnapshotsFuture

	GetBlueprints(ctx workflow.Context, input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsAsync(ctx workflow.Context, input *lightsail.GetBlueprintsInput) *LightsailGetBlueprintsFuture

	GetBundles(ctx workflow.Context, input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error)
	GetBundlesAsync(ctx workflow.Context, input *lightsail.GetBundlesInput) *LightsailGetBundlesFuture

	GetCertificates(ctx workflow.Context, input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error)
	GetCertificatesAsync(ctx workflow.Context, input *lightsail.GetCertificatesInput) *LightsailGetCertificatesFuture

	GetCloudFormationStackRecords(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetCloudFormationStackRecordsAsync(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) *LightsailGetCloudFormationStackRecordsFuture

	GetContactMethods(ctx workflow.Context, input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error)
	GetContactMethodsAsync(ctx workflow.Context, input *lightsail.GetContactMethodsInput) *LightsailGetContactMethodsFuture

	GetDisk(ctx workflow.Context, input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error)
	GetDiskAsync(ctx workflow.Context, input *lightsail.GetDiskInput) *LightsailGetDiskFuture

	GetDiskSnapshot(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshotAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) *LightsailGetDiskSnapshotFuture

	GetDiskSnapshots(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDiskSnapshotsAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) *LightsailGetDiskSnapshotsFuture

	GetDisks(ctx workflow.Context, input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error)
	GetDisksAsync(ctx workflow.Context, input *lightsail.GetDisksInput) *LightsailGetDisksFuture

	GetDistributionBundles(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error)
	GetDistributionBundlesAsync(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) *LightsailGetDistributionBundlesFuture

	GetDistributionLatestCacheReset(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error)
	GetDistributionLatestCacheResetAsync(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) *LightsailGetDistributionLatestCacheResetFuture

	GetDistributionMetricData(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error)
	GetDistributionMetricDataAsync(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) *LightsailGetDistributionMetricDataFuture

	GetDistributions(ctx workflow.Context, input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error)
	GetDistributionsAsync(ctx workflow.Context, input *lightsail.GetDistributionsInput) *LightsailGetDistributionsFuture

	GetDomain(ctx workflow.Context, input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error)
	GetDomainAsync(ctx workflow.Context, input *lightsail.GetDomainInput) *LightsailGetDomainFuture

	GetDomains(ctx workflow.Context, input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error)
	GetDomainsAsync(ctx workflow.Context, input *lightsail.GetDomainsInput) *LightsailGetDomainsFuture

	GetExportSnapshotRecords(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetExportSnapshotRecordsAsync(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) *LightsailGetExportSnapshotRecordsFuture

	GetInstance(ctx workflow.Context, input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error)
	GetInstanceAsync(ctx workflow.Context, input *lightsail.GetInstanceInput) *LightsailGetInstanceFuture

	GetInstanceAccessDetails(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsAsync(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) *LightsailGetInstanceAccessDetailsFuture

	GetInstanceMetricData(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataAsync(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) *LightsailGetInstanceMetricDataFuture

	GetInstancePortStates(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesAsync(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) *LightsailGetInstancePortStatesFuture

	GetInstanceSnapshot(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) *LightsailGetInstanceSnapshotFuture

	GetInstanceSnapshots(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) *LightsailGetInstanceSnapshotsFuture

	GetInstanceState(ctx workflow.Context, input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateAsync(ctx workflow.Context, input *lightsail.GetInstanceStateInput) *LightsailGetInstanceStateFuture

	GetInstances(ctx workflow.Context, input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error)
	GetInstancesAsync(ctx workflow.Context, input *lightsail.GetInstancesInput) *LightsailGetInstancesFuture

	GetKeyPair(ctx workflow.Context, input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairAsync(ctx workflow.Context, input *lightsail.GetKeyPairInput) *LightsailGetKeyPairFuture

	GetKeyPairs(ctx workflow.Context, input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsAsync(ctx workflow.Context, input *lightsail.GetKeyPairsInput) *LightsailGetKeyPairsFuture

	GetLoadBalancer(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) *LightsailGetLoadBalancerFuture

	GetLoadBalancerMetricData(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerMetricDataAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) *LightsailGetLoadBalancerMetricDataFuture

	GetLoadBalancerTlsCertificates(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsCertificatesAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) *LightsailGetLoadBalancerTlsCertificatesFuture

	GetLoadBalancers(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error)
	GetLoadBalancersAsync(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) *LightsailGetLoadBalancersFuture

	GetOperation(ctx workflow.Context, input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error)
	GetOperationAsync(ctx workflow.Context, input *lightsail.GetOperationInput) *LightsailGetOperationFuture

	GetOperations(ctx workflow.Context, input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error)
	GetOperationsAsync(ctx workflow.Context, input *lightsail.GetOperationsInput) *LightsailGetOperationsFuture

	GetOperationsForResource(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceAsync(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) *LightsailGetOperationsForResourceFuture

	GetRegions(ctx workflow.Context, input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error)
	GetRegionsAsync(ctx workflow.Context, input *lightsail.GetRegionsInput) *LightsailGetRegionsFuture

	GetRelationalDatabase(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) *LightsailGetRelationalDatabaseFuture

	GetRelationalDatabaseBlueprints(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBlueprintsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) *LightsailGetRelationalDatabaseBlueprintsFuture

	GetRelationalDatabaseBundles(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseBundlesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) *LightsailGetRelationalDatabaseBundlesFuture

	GetRelationalDatabaseEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) *LightsailGetRelationalDatabaseEventsFuture

	GetRelationalDatabaseLogEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) *LightsailGetRelationalDatabaseLogEventsFuture

	GetRelationalDatabaseLogStreams(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseLogStreamsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) *LightsailGetRelationalDatabaseLogStreamsFuture

	GetRelationalDatabaseMasterUserPassword(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMasterUserPasswordAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) *LightsailGetRelationalDatabaseMasterUserPasswordFuture

	GetRelationalDatabaseMetricData(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseMetricDataAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) *LightsailGetRelationalDatabaseMetricDataFuture

	GetRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) *LightsailGetRelationalDatabaseParametersFuture

	GetRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) *LightsailGetRelationalDatabaseSnapshotFuture

	GetRelationalDatabaseSnapshots(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabaseSnapshotsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) *LightsailGetRelationalDatabaseSnapshotsFuture

	GetRelationalDatabases(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error)
	GetRelationalDatabasesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) *LightsailGetRelationalDatabasesFuture

	GetStaticIp(ctx workflow.Context, input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpAsync(ctx workflow.Context, input *lightsail.GetStaticIpInput) *LightsailGetStaticIpFuture

	GetStaticIps(ctx workflow.Context, input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsAsync(ctx workflow.Context, input *lightsail.GetStaticIpsInput) *LightsailGetStaticIpsFuture

	ImportKeyPair(ctx workflow.Context, input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairAsync(ctx workflow.Context, input *lightsail.ImportKeyPairInput) *LightsailImportKeyPairFuture

	IsVpcPeered(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredAsync(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) *LightsailIsVpcPeeredFuture

	OpenInstancePublicPorts(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) *LightsailOpenInstancePublicPortsFuture

	PeerVpc(ctx workflow.Context, input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error)
	PeerVpcAsync(ctx workflow.Context, input *lightsail.PeerVpcInput) *LightsailPeerVpcFuture

	PutAlarm(ctx workflow.Context, input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error)
	PutAlarmAsync(ctx workflow.Context, input *lightsail.PutAlarmInput) *LightsailPutAlarmFuture

	PutInstancePublicPorts(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) *LightsailPutInstancePublicPortsFuture

	RebootInstance(ctx workflow.Context, input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceAsync(ctx workflow.Context, input *lightsail.RebootInstanceInput) *LightsailRebootInstanceFuture

	RebootRelationalDatabase(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error)
	RebootRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) *LightsailRebootRelationalDatabaseFuture

	ReleaseStaticIp(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpAsync(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) *LightsailReleaseStaticIpFuture

	ResetDistributionCache(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error)
	ResetDistributionCacheAsync(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) *LightsailResetDistributionCacheFuture

	SendContactMethodVerification(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error)
	SendContactMethodVerificationAsync(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) *LightsailSendContactMethodVerificationFuture

	StartInstance(ctx workflow.Context, input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error)
	StartInstanceAsync(ctx workflow.Context, input *lightsail.StartInstanceInput) *LightsailStartInstanceFuture

	StartRelationalDatabase(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error)
	StartRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) *LightsailStartRelationalDatabaseFuture

	StopInstance(ctx workflow.Context, input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error)
	StopInstanceAsync(ctx workflow.Context, input *lightsail.StopInstanceInput) *LightsailStopInstanceFuture

	StopRelationalDatabase(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error)
	StopRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) *LightsailStopRelationalDatabaseFuture

	TagResource(ctx workflow.Context, input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *lightsail.TagResourceInput) *LightsailTagResourceFuture

	TestAlarm(ctx workflow.Context, input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error)
	TestAlarmAsync(ctx workflow.Context, input *lightsail.TestAlarmInput) *LightsailTestAlarmFuture

	UnpeerVpc(ctx workflow.Context, input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcAsync(ctx workflow.Context, input *lightsail.UnpeerVpcInput) *LightsailUnpeerVpcFuture

	UntagResource(ctx workflow.Context, input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *lightsail.UntagResourceInput) *LightsailUntagResourceFuture

	UpdateDistribution(ctx workflow.Context, input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error)
	UpdateDistributionAsync(ctx workflow.Context, input *lightsail.UpdateDistributionInput) *LightsailUpdateDistributionFuture

	UpdateDistributionBundle(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error)
	UpdateDistributionBundleAsync(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) *LightsailUpdateDistributionBundleFuture

	UpdateDomainEntry(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryAsync(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) *LightsailUpdateDomainEntryFuture

	UpdateLoadBalancerAttribute(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
	UpdateLoadBalancerAttributeAsync(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) *LightsailUpdateLoadBalancerAttributeFuture

	UpdateRelationalDatabase(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error)
	UpdateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) *LightsailUpdateRelationalDatabaseFuture

	UpdateRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
	UpdateRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) *LightsailUpdateRelationalDatabaseParametersFuture
}

func NewClient() Client {
	return &stub{}
}