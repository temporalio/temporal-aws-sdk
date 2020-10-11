// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kmsstub

import (
	"github.com/aws/aws-sdk-go/service/kms"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type KMSCancelKeyDeletionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSCancelKeyDeletionFuture) Get(ctx workflow.Context) (*kms.CancelKeyDeletionOutput, error) {
	var output kms.CancelKeyDeletionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSConnectCustomKeyStoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSConnectCustomKeyStoreFuture) Get(ctx workflow.Context) (*kms.ConnectCustomKeyStoreOutput, error) {
	var output kms.ConnectCustomKeyStoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSCreateAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSCreateAliasFuture) Get(ctx workflow.Context) (*kms.CreateAliasOutput, error) {
	var output kms.CreateAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSCreateCustomKeyStoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSCreateCustomKeyStoreFuture) Get(ctx workflow.Context) (*kms.CreateCustomKeyStoreOutput, error) {
	var output kms.CreateCustomKeyStoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSCreateGrantFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSCreateGrantFuture) Get(ctx workflow.Context) (*kms.CreateGrantOutput, error) {
	var output kms.CreateGrantOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSCreateKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSCreateKeyFuture) Get(ctx workflow.Context) (*kms.CreateKeyOutput, error) {
	var output kms.CreateKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDecryptFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDecryptFuture) Get(ctx workflow.Context) (*kms.DecryptOutput, error) {
	var output kms.DecryptOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDeleteAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDeleteAliasFuture) Get(ctx workflow.Context) (*kms.DeleteAliasOutput, error) {
	var output kms.DeleteAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDeleteCustomKeyStoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDeleteCustomKeyStoreFuture) Get(ctx workflow.Context) (*kms.DeleteCustomKeyStoreOutput, error) {
	var output kms.DeleteCustomKeyStoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDeleteImportedKeyMaterialFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDeleteImportedKeyMaterialFuture) Get(ctx workflow.Context) (*kms.DeleteImportedKeyMaterialOutput, error) {
	var output kms.DeleteImportedKeyMaterialOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDescribeCustomKeyStoresFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDescribeCustomKeyStoresFuture) Get(ctx workflow.Context) (*kms.DescribeCustomKeyStoresOutput, error) {
	var output kms.DescribeCustomKeyStoresOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDescribeKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDescribeKeyFuture) Get(ctx workflow.Context) (*kms.DescribeKeyOutput, error) {
	var output kms.DescribeKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDisableKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDisableKeyFuture) Get(ctx workflow.Context) (*kms.DisableKeyOutput, error) {
	var output kms.DisableKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDisableKeyRotationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDisableKeyRotationFuture) Get(ctx workflow.Context) (*kms.DisableKeyRotationOutput, error) {
	var output kms.DisableKeyRotationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSDisconnectCustomKeyStoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSDisconnectCustomKeyStoreFuture) Get(ctx workflow.Context) (*kms.DisconnectCustomKeyStoreOutput, error) {
	var output kms.DisconnectCustomKeyStoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSEnableKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSEnableKeyFuture) Get(ctx workflow.Context) (*kms.EnableKeyOutput, error) {
	var output kms.EnableKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSEnableKeyRotationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSEnableKeyRotationFuture) Get(ctx workflow.Context) (*kms.EnableKeyRotationOutput, error) {
	var output kms.EnableKeyRotationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSEncryptFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSEncryptFuture) Get(ctx workflow.Context) (*kms.EncryptOutput, error) {
	var output kms.EncryptOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGenerateDataKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGenerateDataKeyFuture) Get(ctx workflow.Context) (*kms.GenerateDataKeyOutput, error) {
	var output kms.GenerateDataKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGenerateDataKeyPairFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGenerateDataKeyPairFuture) Get(ctx workflow.Context) (*kms.GenerateDataKeyPairOutput, error) {
	var output kms.GenerateDataKeyPairOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGenerateDataKeyPairWithoutPlaintextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGenerateDataKeyPairWithoutPlaintextFuture) Get(ctx workflow.Context) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
	var output kms.GenerateDataKeyPairWithoutPlaintextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGenerateDataKeyWithoutPlaintextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGenerateDataKeyWithoutPlaintextFuture) Get(ctx workflow.Context) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
	var output kms.GenerateDataKeyWithoutPlaintextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGenerateRandomFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGenerateRandomFuture) Get(ctx workflow.Context) (*kms.GenerateRandomOutput, error) {
	var output kms.GenerateRandomOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGetKeyPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGetKeyPolicyFuture) Get(ctx workflow.Context) (*kms.GetKeyPolicyOutput, error) {
	var output kms.GetKeyPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGetKeyRotationStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGetKeyRotationStatusFuture) Get(ctx workflow.Context) (*kms.GetKeyRotationStatusOutput, error) {
	var output kms.GetKeyRotationStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGetParametersForImportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGetParametersForImportFuture) Get(ctx workflow.Context) (*kms.GetParametersForImportOutput, error) {
	var output kms.GetParametersForImportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSGetPublicKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSGetPublicKeyFuture) Get(ctx workflow.Context) (*kms.GetPublicKeyOutput, error) {
	var output kms.GetPublicKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSImportKeyMaterialFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSImportKeyMaterialFuture) Get(ctx workflow.Context) (*kms.ImportKeyMaterialOutput, error) {
	var output kms.ImportKeyMaterialOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListAliasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListAliasesFuture) Get(ctx workflow.Context) (*kms.ListAliasesOutput, error) {
	var output kms.ListAliasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListGrantsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListGrantsFuture) Get(ctx workflow.Context) (*kms.ListGrantsResponse, error) {
	var output kms.ListGrantsResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListKeyPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListKeyPoliciesFuture) Get(ctx workflow.Context) (*kms.ListKeyPoliciesOutput, error) {
	var output kms.ListKeyPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListKeysFuture) Get(ctx workflow.Context) (*kms.ListKeysOutput, error) {
	var output kms.ListKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListResourceTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListResourceTagsFuture) Get(ctx workflow.Context) (*kms.ListResourceTagsOutput, error) {
	var output kms.ListResourceTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSListRetirableGrantsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSListRetirableGrantsFuture) Get(ctx workflow.Context) (*kms.ListGrantsResponse, error) {
	var output kms.ListGrantsResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSPutKeyPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSPutKeyPolicyFuture) Get(ctx workflow.Context) (*kms.PutKeyPolicyOutput, error) {
	var output kms.PutKeyPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSReEncryptFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSReEncryptFuture) Get(ctx workflow.Context) (*kms.ReEncryptOutput, error) {
	var output kms.ReEncryptOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSRetireGrantFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSRetireGrantFuture) Get(ctx workflow.Context) (*kms.RetireGrantOutput, error) {
	var output kms.RetireGrantOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSRevokeGrantFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSRevokeGrantFuture) Get(ctx workflow.Context) (*kms.RevokeGrantOutput, error) {
	var output kms.RevokeGrantOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSScheduleKeyDeletionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSScheduleKeyDeletionFuture) Get(ctx workflow.Context) (*kms.ScheduleKeyDeletionOutput, error) {
	var output kms.ScheduleKeyDeletionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSSignFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSSignFuture) Get(ctx workflow.Context) (*kms.SignOutput, error) {
	var output kms.SignOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSTagResourceFuture) Get(ctx workflow.Context) (*kms.TagResourceOutput, error) {
	var output kms.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSUntagResourceFuture) Get(ctx workflow.Context) (*kms.UntagResourceOutput, error) {
	var output kms.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSUpdateAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSUpdateAliasFuture) Get(ctx workflow.Context) (*kms.UpdateAliasOutput, error) {
	var output kms.UpdateAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSUpdateCustomKeyStoreFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSUpdateCustomKeyStoreFuture) Get(ctx workflow.Context) (*kms.UpdateCustomKeyStoreOutput, error) {
	var output kms.UpdateCustomKeyStoreOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSUpdateKeyDescriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSUpdateKeyDescriptionFuture) Get(ctx workflow.Context) (*kms.UpdateKeyDescriptionOutput, error) {
	var output kms.UpdateKeyDescriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KMSVerifyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KMSVerifyFuture) Get(ctx workflow.Context) (*kms.VerifyOutput, error) {
	var output kms.VerifyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelKeyDeletion(ctx workflow.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
	var output kms.CancelKeyDeletionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.CancelKeyDeletion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelKeyDeletionAsync(ctx workflow.Context, input *kms.CancelKeyDeletionInput) *KMSCancelKeyDeletionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.CancelKeyDeletion", input)
	return &KMSCancelKeyDeletionFuture{Future: future}
}

func (a *stub) ConnectCustomKeyStore(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error) {
	var output kms.ConnectCustomKeyStoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ConnectCustomKeyStore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) *KMSConnectCustomKeyStoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ConnectCustomKeyStore", input)
	return &KMSConnectCustomKeyStoreFuture{Future: future}
}

func (a *stub) CreateAlias(ctx workflow.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
	var output kms.CreateAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.CreateAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAliasAsync(ctx workflow.Context, input *kms.CreateAliasInput) *KMSCreateAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.CreateAlias", input)
	return &KMSCreateAliasFuture{Future: future}
}

func (a *stub) CreateCustomKeyStore(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error) {
	var output kms.CreateCustomKeyStoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.CreateCustomKeyStore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCustomKeyStoreAsync(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) *KMSCreateCustomKeyStoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.CreateCustomKeyStore", input)
	return &KMSCreateCustomKeyStoreFuture{Future: future}
}

func (a *stub) CreateGrant(ctx workflow.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error) {
	var output kms.CreateGrantOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.CreateGrant", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGrantAsync(ctx workflow.Context, input *kms.CreateGrantInput) *KMSCreateGrantFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.CreateGrant", input)
	return &KMSCreateGrantFuture{Future: future}
}

func (a *stub) CreateKey(ctx workflow.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
	var output kms.CreateKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.CreateKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateKeyAsync(ctx workflow.Context, input *kms.CreateKeyInput) *KMSCreateKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.CreateKey", input)
	return &KMSCreateKeyFuture{Future: future}
}

func (a *stub) Decrypt(ctx workflow.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	var output kms.DecryptOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.Decrypt", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DecryptAsync(ctx workflow.Context, input *kms.DecryptInput) *KMSDecryptFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.Decrypt", input)
	return &KMSDecryptFuture{Future: future}
}

func (a *stub) DeleteAlias(ctx workflow.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
	var output kms.DeleteAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DeleteAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAliasAsync(ctx workflow.Context, input *kms.DeleteAliasInput) *KMSDeleteAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DeleteAlias", input)
	return &KMSDeleteAliasFuture{Future: future}
}

func (a *stub) DeleteCustomKeyStore(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error) {
	var output kms.DeleteCustomKeyStoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DeleteCustomKeyStore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteCustomKeyStoreAsync(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) *KMSDeleteCustomKeyStoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DeleteCustomKeyStore", input)
	return &KMSDeleteCustomKeyStoreFuture{Future: future}
}

func (a *stub) DeleteImportedKeyMaterial(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error) {
	var output kms.DeleteImportedKeyMaterialOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DeleteImportedKeyMaterial", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteImportedKeyMaterialAsync(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) *KMSDeleteImportedKeyMaterialFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DeleteImportedKeyMaterial", input)
	return &KMSDeleteImportedKeyMaterialFuture{Future: future}
}

func (a *stub) DescribeCustomKeyStores(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error) {
	var output kms.DescribeCustomKeyStoresOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DescribeCustomKeyStores", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCustomKeyStoresAsync(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) *KMSDescribeCustomKeyStoresFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DescribeCustomKeyStores", input)
	return &KMSDescribeCustomKeyStoresFuture{Future: future}
}

func (a *stub) DescribeKey(ctx workflow.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
	var output kms.DescribeKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DescribeKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeKeyAsync(ctx workflow.Context, input *kms.DescribeKeyInput) *KMSDescribeKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DescribeKey", input)
	return &KMSDescribeKeyFuture{Future: future}
}

func (a *stub) DisableKey(ctx workflow.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
	var output kms.DisableKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DisableKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableKeyAsync(ctx workflow.Context, input *kms.DisableKeyInput) *KMSDisableKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DisableKey", input)
	return &KMSDisableKeyFuture{Future: future}
}

func (a *stub) DisableKeyRotation(ctx workflow.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error) {
	var output kms.DisableKeyRotationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DisableKeyRotation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableKeyRotationAsync(ctx workflow.Context, input *kms.DisableKeyRotationInput) *KMSDisableKeyRotationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DisableKeyRotation", input)
	return &KMSDisableKeyRotationFuture{Future: future}
}

func (a *stub) DisconnectCustomKeyStore(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error) {
	var output kms.DisconnectCustomKeyStoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.DisconnectCustomKeyStore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisconnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) *KMSDisconnectCustomKeyStoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.DisconnectCustomKeyStore", input)
	return &KMSDisconnectCustomKeyStoreFuture{Future: future}
}

func (a *stub) EnableKey(ctx workflow.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error) {
	var output kms.EnableKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.EnableKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableKeyAsync(ctx workflow.Context, input *kms.EnableKeyInput) *KMSEnableKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.EnableKey", input)
	return &KMSEnableKeyFuture{Future: future}
}

func (a *stub) EnableKeyRotation(ctx workflow.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error) {
	var output kms.EnableKeyRotationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.EnableKeyRotation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableKeyRotationAsync(ctx workflow.Context, input *kms.EnableKeyRotationInput) *KMSEnableKeyRotationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.EnableKeyRotation", input)
	return &KMSEnableKeyRotationFuture{Future: future}
}

func (a *stub) Encrypt(ctx workflow.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
	var output kms.EncryptOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.Encrypt", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EncryptAsync(ctx workflow.Context, input *kms.EncryptInput) *KMSEncryptFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.Encrypt", input)
	return &KMSEncryptFuture{Future: future}
}

func (a *stub) GenerateDataKey(ctx workflow.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error) {
	var output kms.GenerateDataKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataKeyAsync(ctx workflow.Context, input *kms.GenerateDataKeyInput) *KMSGenerateDataKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKey", input)
	return &KMSGenerateDataKeyFuture{Future: future}
}

func (a *stub) GenerateDataKeyPair(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error) {
	var output kms.GenerateDataKeyPairOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyPair", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataKeyPairAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) *KMSGenerateDataKeyPairFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyPair", input)
	return &KMSGenerateDataKeyPairFuture{Future: future}
}

func (a *stub) GenerateDataKeyPairWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
	var output kms.GenerateDataKeyPairWithoutPlaintextOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyPairWithoutPlaintext", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataKeyPairWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) *KMSGenerateDataKeyPairWithoutPlaintextFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyPairWithoutPlaintext", input)
	return &KMSGenerateDataKeyPairWithoutPlaintextFuture{Future: future}
}

func (a *stub) GenerateDataKeyWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
	var output kms.GenerateDataKeyWithoutPlaintextOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyWithoutPlaintext", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataKeyWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) *KMSGenerateDataKeyWithoutPlaintextFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GenerateDataKeyWithoutPlaintext", input)
	return &KMSGenerateDataKeyWithoutPlaintextFuture{Future: future}
}

func (a *stub) GenerateRandom(ctx workflow.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error) {
	var output kms.GenerateRandomOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GenerateRandom", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateRandomAsync(ctx workflow.Context, input *kms.GenerateRandomInput) *KMSGenerateRandomFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GenerateRandom", input)
	return &KMSGenerateRandomFuture{Future: future}
}

func (a *stub) GetKeyPolicy(ctx workflow.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error) {
	var output kms.GetKeyPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GetKeyPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetKeyPolicyAsync(ctx workflow.Context, input *kms.GetKeyPolicyInput) *KMSGetKeyPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GetKeyPolicy", input)
	return &KMSGetKeyPolicyFuture{Future: future}
}

func (a *stub) GetKeyRotationStatus(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error) {
	var output kms.GetKeyRotationStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GetKeyRotationStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetKeyRotationStatusAsync(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) *KMSGetKeyRotationStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GetKeyRotationStatus", input)
	return &KMSGetKeyRotationStatusFuture{Future: future}
}

func (a *stub) GetParametersForImport(ctx workflow.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error) {
	var output kms.GetParametersForImportOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GetParametersForImport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetParametersForImportAsync(ctx workflow.Context, input *kms.GetParametersForImportInput) *KMSGetParametersForImportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GetParametersForImport", input)
	return &KMSGetParametersForImportFuture{Future: future}
}

func (a *stub) GetPublicKey(ctx workflow.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error) {
	var output kms.GetPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.GetPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPublicKeyAsync(ctx workflow.Context, input *kms.GetPublicKeyInput) *KMSGetPublicKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.GetPublicKey", input)
	return &KMSGetPublicKeyFuture{Future: future}
}

func (a *stub) ImportKeyMaterial(ctx workflow.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error) {
	var output kms.ImportKeyMaterialOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ImportKeyMaterial", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportKeyMaterialAsync(ctx workflow.Context, input *kms.ImportKeyMaterialInput) *KMSImportKeyMaterialFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ImportKeyMaterial", input)
	return &KMSImportKeyMaterialFuture{Future: future}
}

func (a *stub) ListAliases(ctx workflow.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
	var output kms.ListAliasesOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListAliases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAliasesAsync(ctx workflow.Context, input *kms.ListAliasesInput) *KMSListAliasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListAliases", input)
	return &KMSListAliasesFuture{Future: future}
}

func (a *stub) ListGrants(ctx workflow.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error) {
	var output kms.ListGrantsResponse
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListGrants", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGrantsAsync(ctx workflow.Context, input *kms.ListGrantsInput) *KMSListGrantsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListGrants", input)
	return &KMSListGrantsFuture{Future: future}
}

func (a *stub) ListKeyPolicies(ctx workflow.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error) {
	var output kms.ListKeyPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListKeyPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListKeyPoliciesAsync(ctx workflow.Context, input *kms.ListKeyPoliciesInput) *KMSListKeyPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListKeyPolicies", input)
	return &KMSListKeyPoliciesFuture{Future: future}
}

func (a *stub) ListKeys(ctx workflow.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
	var output kms.ListKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListKeysAsync(ctx workflow.Context, input *kms.ListKeysInput) *KMSListKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListKeys", input)
	return &KMSListKeysFuture{Future: future}
}

func (a *stub) ListResourceTags(ctx workflow.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error) {
	var output kms.ListResourceTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListResourceTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourceTagsAsync(ctx workflow.Context, input *kms.ListResourceTagsInput) *KMSListResourceTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListResourceTags", input)
	return &KMSListResourceTagsFuture{Future: future}
}

func (a *stub) ListRetirableGrants(ctx workflow.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error) {
	var output kms.ListGrantsResponse
	err := workflow.ExecuteActivity(ctx, "aws.kms.ListRetirableGrants", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRetirableGrantsAsync(ctx workflow.Context, input *kms.ListRetirableGrantsInput) *KMSListRetirableGrantsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ListRetirableGrants", input)
	return &KMSListRetirableGrantsFuture{Future: future}
}

func (a *stub) PutKeyPolicy(ctx workflow.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error) {
	var output kms.PutKeyPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.PutKeyPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutKeyPolicyAsync(ctx workflow.Context, input *kms.PutKeyPolicyInput) *KMSPutKeyPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.PutKeyPolicy", input)
	return &KMSPutKeyPolicyFuture{Future: future}
}

func (a *stub) ReEncrypt(ctx workflow.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error) {
	var output kms.ReEncryptOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ReEncrypt", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReEncryptAsync(ctx workflow.Context, input *kms.ReEncryptInput) *KMSReEncryptFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ReEncrypt", input)
	return &KMSReEncryptFuture{Future: future}
}

func (a *stub) RetireGrant(ctx workflow.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error) {
	var output kms.RetireGrantOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.RetireGrant", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RetireGrantAsync(ctx workflow.Context, input *kms.RetireGrantInput) *KMSRetireGrantFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.RetireGrant", input)
	return &KMSRetireGrantFuture{Future: future}
}

func (a *stub) RevokeGrant(ctx workflow.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error) {
	var output kms.RevokeGrantOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.RevokeGrant", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RevokeGrantAsync(ctx workflow.Context, input *kms.RevokeGrantInput) *KMSRevokeGrantFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.RevokeGrant", input)
	return &KMSRevokeGrantFuture{Future: future}
}

func (a *stub) ScheduleKeyDeletion(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
	var output kms.ScheduleKeyDeletionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.ScheduleKeyDeletion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ScheduleKeyDeletionAsync(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) *KMSScheduleKeyDeletionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.ScheduleKeyDeletion", input)
	return &KMSScheduleKeyDeletionFuture{Future: future}
}

func (a *stub) Sign(ctx workflow.Context, input *kms.SignInput) (*kms.SignOutput, error) {
	var output kms.SignOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.Sign", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SignAsync(ctx workflow.Context, input *kms.SignInput) *KMSSignFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.Sign", input)
	return &KMSSignFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error) {
	var output kms.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *kms.TagResourceInput) *KMSTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.TagResource", input)
	return &KMSTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error) {
	var output kms.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *kms.UntagResourceInput) *KMSUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.UntagResource", input)
	return &KMSUntagResourceFuture{Future: future}
}

func (a *stub) UpdateAlias(ctx workflow.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error) {
	var output kms.UpdateAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.UpdateAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAliasAsync(ctx workflow.Context, input *kms.UpdateAliasInput) *KMSUpdateAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.UpdateAlias", input)
	return &KMSUpdateAliasFuture{Future: future}
}

func (a *stub) UpdateCustomKeyStore(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error) {
	var output kms.UpdateCustomKeyStoreOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.UpdateCustomKeyStore", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateCustomKeyStoreAsync(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) *KMSUpdateCustomKeyStoreFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.UpdateCustomKeyStore", input)
	return &KMSUpdateCustomKeyStoreFuture{Future: future}
}

func (a *stub) UpdateKeyDescription(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error) {
	var output kms.UpdateKeyDescriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.UpdateKeyDescription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateKeyDescriptionAsync(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) *KMSUpdateKeyDescriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.UpdateKeyDescription", input)
	return &KMSUpdateKeyDescriptionFuture{Future: future}
}

func (a *stub) Verify(ctx workflow.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error) {
	var output kms.VerifyOutput
	err := workflow.ExecuteActivity(ctx, "aws.kms.Verify", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) VerifyAsync(ctx workflow.Context, input *kms.VerifyInput) *KMSVerifyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kms.Verify", input)
	return &KMSVerifyFuture{Future: future}
}
