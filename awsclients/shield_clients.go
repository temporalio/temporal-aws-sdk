package awsclients

import (
	"github.com/aws/aws-sdk-go/service/shield"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ShieldClient interface {
    AssociateDRTLogBucket(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error)
    AssociateDRTLogBucketAsync(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) *ShieldAssociateDRTLogBucketResult

    AssociateDRTRole(ctx workflow.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error)
    AssociateDRTRoleAsync(ctx workflow.Context, input *shield.AssociateDRTRoleInput) *ShieldAssociateDRTRoleResult

    AssociateHealthCheck(ctx workflow.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error)
    AssociateHealthCheckAsync(ctx workflow.Context, input *shield.AssociateHealthCheckInput) *ShieldAssociateHealthCheckResult

    AssociateProactiveEngagementDetails(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error)
    AssociateProactiveEngagementDetailsAsync(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) *ShieldAssociateProactiveEngagementDetailsResult

    CreateProtection(ctx workflow.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error)
    CreateProtectionAsync(ctx workflow.Context, input *shield.CreateProtectionInput) *ShieldCreateProtectionResult

    CreateSubscription(ctx workflow.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error)
    CreateSubscriptionAsync(ctx workflow.Context, input *shield.CreateSubscriptionInput) *ShieldCreateSubscriptionResult

    DeleteProtection(ctx workflow.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error)
    DeleteProtectionAsync(ctx workflow.Context, input *shield.DeleteProtectionInput) *ShieldDeleteProtectionResult

    DeleteSubscription(ctx workflow.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error)
    DeleteSubscriptionAsync(ctx workflow.Context, input *shield.DeleteSubscriptionInput) *ShieldDeleteSubscriptionResult

    DescribeAttack(ctx workflow.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error)
    DescribeAttackAsync(ctx workflow.Context, input *shield.DescribeAttackInput) *ShieldDescribeAttackResult

    DescribeDRTAccess(ctx workflow.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error)
    DescribeDRTAccessAsync(ctx workflow.Context, input *shield.DescribeDRTAccessInput) *ShieldDescribeDRTAccessResult

    DescribeEmergencyContactSettings(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error)
    DescribeEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) *ShieldDescribeEmergencyContactSettingsResult

    DescribeProtection(ctx workflow.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error)
    DescribeProtectionAsync(ctx workflow.Context, input *shield.DescribeProtectionInput) *ShieldDescribeProtectionResult

    DescribeSubscription(ctx workflow.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error)
    DescribeSubscriptionAsync(ctx workflow.Context, input *shield.DescribeSubscriptionInput) *ShieldDescribeSubscriptionResult

    DisableProactiveEngagement(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error)
    DisableProactiveEngagementAsync(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) *ShieldDisableProactiveEngagementResult

    DisassociateDRTLogBucket(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error)
    DisassociateDRTLogBucketAsync(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) *ShieldDisassociateDRTLogBucketResult

    DisassociateDRTRole(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error)
    DisassociateDRTRoleAsync(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) *ShieldDisassociateDRTRoleResult

    DisassociateHealthCheck(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error)
    DisassociateHealthCheckAsync(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) *ShieldDisassociateHealthCheckResult

    EnableProactiveEngagement(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error)
    EnableProactiveEngagementAsync(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) *ShieldEnableProactiveEngagementResult

    GetSubscriptionState(ctx workflow.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error)
    GetSubscriptionStateAsync(ctx workflow.Context, input *shield.GetSubscriptionStateInput) *ShieldGetSubscriptionStateResult

    ListAttacks(ctx workflow.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error)
    ListAttacksAsync(ctx workflow.Context, input *shield.ListAttacksInput) *ShieldListAttacksResult

    ListProtections(ctx workflow.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error)
    ListProtectionsAsync(ctx workflow.Context, input *shield.ListProtectionsInput) *ShieldListProtectionsResult

    UpdateEmergencyContactSettings(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error)
    UpdateEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) *ShieldUpdateEmergencyContactSettingsResult

    UpdateSubscription(ctx workflow.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error)
    UpdateSubscriptionAsync(ctx workflow.Context, input *shield.UpdateSubscriptionInput) *ShieldUpdateSubscriptionResult
}

type ShieldAssociateDRTLogBucketResult struct {
	Result workflow.Future
}

func (r *ShieldAssociateDRTLogBucketResult) Get(ctx workflow.Context) (*shield.AssociateDRTLogBucketOutput, error) {
    var output shield.AssociateDRTLogBucketOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldAssociateDRTRoleResult struct {
	Result workflow.Future
}

func (r *ShieldAssociateDRTRoleResult) Get(ctx workflow.Context) (*shield.AssociateDRTRoleOutput, error) {
    var output shield.AssociateDRTRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldAssociateHealthCheckResult struct {
	Result workflow.Future
}

func (r *ShieldAssociateHealthCheckResult) Get(ctx workflow.Context) (*shield.AssociateHealthCheckOutput, error) {
    var output shield.AssociateHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldAssociateProactiveEngagementDetailsResult struct {
	Result workflow.Future
}

func (r *ShieldAssociateProactiveEngagementDetailsResult) Get(ctx workflow.Context) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
    var output shield.AssociateProactiveEngagementDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldCreateProtectionResult struct {
	Result workflow.Future
}

func (r *ShieldCreateProtectionResult) Get(ctx workflow.Context) (*shield.CreateProtectionOutput, error) {
    var output shield.CreateProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldCreateSubscriptionResult struct {
	Result workflow.Future
}

func (r *ShieldCreateSubscriptionResult) Get(ctx workflow.Context) (*shield.CreateSubscriptionOutput, error) {
    var output shield.CreateSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDeleteProtectionResult struct {
	Result workflow.Future
}

func (r *ShieldDeleteProtectionResult) Get(ctx workflow.Context) (*shield.DeleteProtectionOutput, error) {
    var output shield.DeleteProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDeleteSubscriptionResult struct {
	Result workflow.Future
}

func (r *ShieldDeleteSubscriptionResult) Get(ctx workflow.Context) (*shield.DeleteSubscriptionOutput, error) {
    var output shield.DeleteSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDescribeAttackResult struct {
	Result workflow.Future
}

func (r *ShieldDescribeAttackResult) Get(ctx workflow.Context) (*shield.DescribeAttackOutput, error) {
    var output shield.DescribeAttackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDescribeDRTAccessResult struct {
	Result workflow.Future
}

func (r *ShieldDescribeDRTAccessResult) Get(ctx workflow.Context) (*shield.DescribeDRTAccessOutput, error) {
    var output shield.DescribeDRTAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDescribeEmergencyContactSettingsResult struct {
	Result workflow.Future
}

func (r *ShieldDescribeEmergencyContactSettingsResult) Get(ctx workflow.Context) (*shield.DescribeEmergencyContactSettingsOutput, error) {
    var output shield.DescribeEmergencyContactSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDescribeProtectionResult struct {
	Result workflow.Future
}

func (r *ShieldDescribeProtectionResult) Get(ctx workflow.Context) (*shield.DescribeProtectionOutput, error) {
    var output shield.DescribeProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDescribeSubscriptionResult struct {
	Result workflow.Future
}

func (r *ShieldDescribeSubscriptionResult) Get(ctx workflow.Context) (*shield.DescribeSubscriptionOutput, error) {
    var output shield.DescribeSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDisableProactiveEngagementResult struct {
	Result workflow.Future
}

func (r *ShieldDisableProactiveEngagementResult) Get(ctx workflow.Context) (*shield.DisableProactiveEngagementOutput, error) {
    var output shield.DisableProactiveEngagementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDisassociateDRTLogBucketResult struct {
	Result workflow.Future
}

func (r *ShieldDisassociateDRTLogBucketResult) Get(ctx workflow.Context) (*shield.DisassociateDRTLogBucketOutput, error) {
    var output shield.DisassociateDRTLogBucketOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDisassociateDRTRoleResult struct {
	Result workflow.Future
}

func (r *ShieldDisassociateDRTRoleResult) Get(ctx workflow.Context) (*shield.DisassociateDRTRoleOutput, error) {
    var output shield.DisassociateDRTRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldDisassociateHealthCheckResult struct {
	Result workflow.Future
}

func (r *ShieldDisassociateHealthCheckResult) Get(ctx workflow.Context) (*shield.DisassociateHealthCheckOutput, error) {
    var output shield.DisassociateHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldEnableProactiveEngagementResult struct {
	Result workflow.Future
}

func (r *ShieldEnableProactiveEngagementResult) Get(ctx workflow.Context) (*shield.EnableProactiveEngagementOutput, error) {
    var output shield.EnableProactiveEngagementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldGetSubscriptionStateResult struct {
	Result workflow.Future
}

func (r *ShieldGetSubscriptionStateResult) Get(ctx workflow.Context) (*shield.GetSubscriptionStateOutput, error) {
    var output shield.GetSubscriptionStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldListAttacksResult struct {
	Result workflow.Future
}

func (r *ShieldListAttacksResult) Get(ctx workflow.Context) (*shield.ListAttacksOutput, error) {
    var output shield.ListAttacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldListProtectionsResult struct {
	Result workflow.Future
}

func (r *ShieldListProtectionsResult) Get(ctx workflow.Context) (*shield.ListProtectionsOutput, error) {
    var output shield.ListProtectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldUpdateEmergencyContactSettingsResult struct {
	Result workflow.Future
}

func (r *ShieldUpdateEmergencyContactSettingsResult) Get(ctx workflow.Context) (*shield.UpdateEmergencyContactSettingsOutput, error) {
    var output shield.UpdateEmergencyContactSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldUpdateSubscriptionResult struct {
	Result workflow.Future
}

func (r *ShieldUpdateSubscriptionResult) Get(ctx workflow.Context) (*shield.UpdateSubscriptionOutput, error) {
    var output shield.UpdateSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ShieldStub struct {
    activities awsactivities.ShieldActivities
}

func NewShieldStub() ShieldClient {
    return &ShieldStub{}
}

func (a *ShieldStub) AssociateDRTLogBucket(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
    var output shield.AssociateDRTLogBucketOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateDRTLogBucket, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) AssociateDRTLogBucketAsync(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) *ShieldAssociateDRTLogBucketResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateDRTLogBucket, input)
    return &ShieldAssociateDRTLogBucketResult{Result: future}
}

func (a *ShieldStub) AssociateDRTRole(ctx workflow.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error) {
    var output shield.AssociateDRTRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateDRTRole, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) AssociateDRTRoleAsync(ctx workflow.Context, input *shield.AssociateDRTRoleInput) *ShieldAssociateDRTRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateDRTRole, input)
    return &ShieldAssociateDRTRoleResult{Result: future}
}

func (a *ShieldStub) AssociateHealthCheck(ctx workflow.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error) {
    var output shield.AssociateHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) AssociateHealthCheckAsync(ctx workflow.Context, input *shield.AssociateHealthCheckInput) *ShieldAssociateHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateHealthCheck, input)
    return &ShieldAssociateHealthCheckResult{Result: future}
}

func (a *ShieldStub) AssociateProactiveEngagementDetails(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
    var output shield.AssociateProactiveEngagementDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateProactiveEngagementDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) AssociateProactiveEngagementDetailsAsync(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) *ShieldAssociateProactiveEngagementDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateProactiveEngagementDetails, input)
    return &ShieldAssociateProactiveEngagementDetailsResult{Result: future}
}

func (a *ShieldStub) CreateProtection(ctx workflow.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
    var output shield.CreateProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) CreateProtectionAsync(ctx workflow.Context, input *shield.CreateProtectionInput) *ShieldCreateProtectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProtection, input)
    return &ShieldCreateProtectionResult{Result: future}
}

func (a *ShieldStub) CreateSubscription(ctx workflow.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error) {
    var output shield.CreateSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) CreateSubscriptionAsync(ctx workflow.Context, input *shield.CreateSubscriptionInput) *ShieldCreateSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSubscription, input)
    return &ShieldCreateSubscriptionResult{Result: future}
}

func (a *ShieldStub) DeleteProtection(ctx workflow.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error) {
    var output shield.DeleteProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DeleteProtectionAsync(ctx workflow.Context, input *shield.DeleteProtectionInput) *ShieldDeleteProtectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProtection, input)
    return &ShieldDeleteProtectionResult{Result: future}
}

func (a *ShieldStub) DeleteSubscription(ctx workflow.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error) {
    var output shield.DeleteSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DeleteSubscriptionAsync(ctx workflow.Context, input *shield.DeleteSubscriptionInput) *ShieldDeleteSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscription, input)
    return &ShieldDeleteSubscriptionResult{Result: future}
}

func (a *ShieldStub) DescribeAttack(ctx workflow.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
    var output shield.DescribeAttackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAttack, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DescribeAttackAsync(ctx workflow.Context, input *shield.DescribeAttackInput) *ShieldDescribeAttackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAttack, input)
    return &ShieldDescribeAttackResult{Result: future}
}

func (a *ShieldStub) DescribeDRTAccess(ctx workflow.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
    var output shield.DescribeDRTAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDRTAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DescribeDRTAccessAsync(ctx workflow.Context, input *shield.DescribeDRTAccessInput) *ShieldDescribeDRTAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDRTAccess, input)
    return &ShieldDescribeDRTAccessResult{Result: future}
}

func (a *ShieldStub) DescribeEmergencyContactSettings(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
    var output shield.DescribeEmergencyContactSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEmergencyContactSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DescribeEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) *ShieldDescribeEmergencyContactSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEmergencyContactSettings, input)
    return &ShieldDescribeEmergencyContactSettingsResult{Result: future}
}

func (a *ShieldStub) DescribeProtection(ctx workflow.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
    var output shield.DescribeProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DescribeProtectionAsync(ctx workflow.Context, input *shield.DescribeProtectionInput) *ShieldDescribeProtectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProtection, input)
    return &ShieldDescribeProtectionResult{Result: future}
}

func (a *ShieldStub) DescribeSubscription(ctx workflow.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
    var output shield.DescribeSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DescribeSubscriptionAsync(ctx workflow.Context, input *shield.DescribeSubscriptionInput) *ShieldDescribeSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscription, input)
    return &ShieldDescribeSubscriptionResult{Result: future}
}

func (a *ShieldStub) DisableProactiveEngagement(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error) {
    var output shield.DisableProactiveEngagementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableProactiveEngagement, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DisableProactiveEngagementAsync(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) *ShieldDisableProactiveEngagementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableProactiveEngagement, input)
    return &ShieldDisableProactiveEngagementResult{Result: future}
}

func (a *ShieldStub) DisassociateDRTLogBucket(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error) {
    var output shield.DisassociateDRTLogBucketOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateDRTLogBucket, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DisassociateDRTLogBucketAsync(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) *ShieldDisassociateDRTLogBucketResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateDRTLogBucket, input)
    return &ShieldDisassociateDRTLogBucketResult{Result: future}
}

func (a *ShieldStub) DisassociateDRTRole(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error) {
    var output shield.DisassociateDRTRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateDRTRole, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DisassociateDRTRoleAsync(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) *ShieldDisassociateDRTRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateDRTRole, input)
    return &ShieldDisassociateDRTRoleResult{Result: future}
}

func (a *ShieldStub) DisassociateHealthCheck(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error) {
    var output shield.DisassociateHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) DisassociateHealthCheckAsync(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) *ShieldDisassociateHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateHealthCheck, input)
    return &ShieldDisassociateHealthCheckResult{Result: future}
}

func (a *ShieldStub) EnableProactiveEngagement(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error) {
    var output shield.EnableProactiveEngagementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableProactiveEngagement, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) EnableProactiveEngagementAsync(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) *ShieldEnableProactiveEngagementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableProactiveEngagement, input)
    return &ShieldEnableProactiveEngagementResult{Result: future}
}

func (a *ShieldStub) GetSubscriptionState(ctx workflow.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
    var output shield.GetSubscriptionStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSubscriptionState, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) GetSubscriptionStateAsync(ctx workflow.Context, input *shield.GetSubscriptionStateInput) *ShieldGetSubscriptionStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSubscriptionState, input)
    return &ShieldGetSubscriptionStateResult{Result: future}
}

func (a *ShieldStub) ListAttacks(ctx workflow.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
    var output shield.ListAttacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttacks, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) ListAttacksAsync(ctx workflow.Context, input *shield.ListAttacksInput) *ShieldListAttacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAttacks, input)
    return &ShieldListAttacksResult{Result: future}
}

func (a *ShieldStub) ListProtections(ctx workflow.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
    var output shield.ListProtectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProtections, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) ListProtectionsAsync(ctx workflow.Context, input *shield.ListProtectionsInput) *ShieldListProtectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProtections, input)
    return &ShieldListProtectionsResult{Result: future}
}

func (a *ShieldStub) UpdateEmergencyContactSettings(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error) {
    var output shield.UpdateEmergencyContactSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEmergencyContactSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) UpdateEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) *ShieldUpdateEmergencyContactSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEmergencyContactSettings, input)
    return &ShieldUpdateEmergencyContactSettingsResult{Result: future}
}

func (a *ShieldStub) UpdateSubscription(ctx workflow.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error) {
    var output shield.UpdateSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *ShieldStub) UpdateSubscriptionAsync(ctx workflow.Context, input *shield.UpdateSubscriptionInput) *ShieldUpdateSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSubscription, input)
    return &ShieldUpdateSubscriptionResult{Result: future}
}
