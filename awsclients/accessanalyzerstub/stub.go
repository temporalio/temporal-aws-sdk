// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package accessanalyzerstub

import (
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"

)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type AccessAnalyzerCreateAnalyzerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerCreateAnalyzerFuture) Get(ctx workflow.Context) (*accessanalyzer.CreateAnalyzerOutput, error) {
	var output accessanalyzer.CreateAnalyzerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerCreateArchiveRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerCreateArchiveRuleFuture) Get(ctx workflow.Context) (*accessanalyzer.CreateArchiveRuleOutput, error) {
	var output accessanalyzer.CreateArchiveRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerDeleteAnalyzerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerDeleteAnalyzerFuture) Get(ctx workflow.Context) (*accessanalyzer.DeleteAnalyzerOutput, error) {
	var output accessanalyzer.DeleteAnalyzerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerDeleteArchiveRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerDeleteArchiveRuleFuture) Get(ctx workflow.Context) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
	var output accessanalyzer.DeleteArchiveRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerGetAnalyzedResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerGetAnalyzedResourceFuture) Get(ctx workflow.Context) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	var output accessanalyzer.GetAnalyzedResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerGetAnalyzerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerGetAnalyzerFuture) Get(ctx workflow.Context) (*accessanalyzer.GetAnalyzerOutput, error) {
	var output accessanalyzer.GetAnalyzerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerGetArchiveRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerGetArchiveRuleFuture) Get(ctx workflow.Context) (*accessanalyzer.GetArchiveRuleOutput, error) {
	var output accessanalyzer.GetArchiveRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerGetFindingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerGetFindingFuture) Get(ctx workflow.Context) (*accessanalyzer.GetFindingOutput, error) {
	var output accessanalyzer.GetFindingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerListAnalyzedResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerListAnalyzedResourcesFuture) Get(ctx workflow.Context) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	var output accessanalyzer.ListAnalyzedResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerListAnalyzersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerListAnalyzersFuture) Get(ctx workflow.Context) (*accessanalyzer.ListAnalyzersOutput, error) {
	var output accessanalyzer.ListAnalyzersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerListArchiveRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerListArchiveRulesFuture) Get(ctx workflow.Context) (*accessanalyzer.ListArchiveRulesOutput, error) {
	var output accessanalyzer.ListArchiveRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerListFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerListFindingsFuture) Get(ctx workflow.Context) (*accessanalyzer.ListFindingsOutput, error) {
	var output accessanalyzer.ListFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerListTagsForResourceFuture) Get(ctx workflow.Context) (*accessanalyzer.ListTagsForResourceOutput, error) {
	var output accessanalyzer.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerStartResourceScanFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerStartResourceScanFuture) Get(ctx workflow.Context) (*accessanalyzer.StartResourceScanOutput, error) {
	var output accessanalyzer.StartResourceScanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerTagResourceFuture) Get(ctx workflow.Context) (*accessanalyzer.TagResourceOutput, error) {
	var output accessanalyzer.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerUntagResourceFuture) Get(ctx workflow.Context) (*accessanalyzer.UntagResourceOutput, error) {
	var output accessanalyzer.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerUpdateArchiveRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerUpdateArchiveRuleFuture) Get(ctx workflow.Context) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
	var output accessanalyzer.UpdateArchiveRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AccessAnalyzerUpdateFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AccessAnalyzerUpdateFindingsFuture) Get(ctx workflow.Context) (*accessanalyzer.UpdateFindingsOutput, error) {
	var output accessanalyzer.UpdateFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAnalyzer(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error) {
	var output accessanalyzer.CreateAnalyzerOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.CreateAnalyzer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) *AccessAnalyzerCreateAnalyzerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.CreateAnalyzer", input)
	return &AccessAnalyzerCreateAnalyzerFuture{Future: future}
}

func (a *stub) CreateArchiveRule(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error) {
	var output accessanalyzer.CreateArchiveRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.CreateArchiveRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) *AccessAnalyzerCreateArchiveRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.CreateArchiveRule", input)
	return &AccessAnalyzerCreateArchiveRuleFuture{Future: future}
}

func (a *stub) DeleteAnalyzer(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error) {
	var output accessanalyzer.DeleteAnalyzerOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.DeleteAnalyzer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) *AccessAnalyzerDeleteAnalyzerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.DeleteAnalyzer", input)
	return &AccessAnalyzerDeleteAnalyzerFuture{Future: future}
}

func (a *stub) DeleteArchiveRule(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
	var output accessanalyzer.DeleteArchiveRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.DeleteArchiveRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) *AccessAnalyzerDeleteArchiveRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.DeleteArchiveRule", input)
	return &AccessAnalyzerDeleteArchiveRuleFuture{Future: future}
}

func (a *stub) GetAnalyzedResource(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	var output accessanalyzer.GetAnalyzedResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetAnalyzedResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAnalyzedResourceAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) *AccessAnalyzerGetAnalyzedResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetAnalyzedResource", input)
	return &AccessAnalyzerGetAnalyzedResourceFuture{Future: future}
}

func (a *stub) GetAnalyzer(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
	var output accessanalyzer.GetAnalyzerOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetAnalyzer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) *AccessAnalyzerGetAnalyzerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetAnalyzer", input)
	return &AccessAnalyzerGetAnalyzerFuture{Future: future}
}

func (a *stub) GetArchiveRule(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
	var output accessanalyzer.GetArchiveRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetArchiveRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) *AccessAnalyzerGetArchiveRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetArchiveRule", input)
	return &AccessAnalyzerGetArchiveRuleFuture{Future: future}
}

func (a *stub) GetFinding(ctx workflow.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
	var output accessanalyzer.GetFindingOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetFinding", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFindingAsync(ctx workflow.Context, input *accessanalyzer.GetFindingInput) *AccessAnalyzerGetFindingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.GetFinding", input)
	return &AccessAnalyzerGetFindingFuture{Future: future}
}

func (a *stub) ListAnalyzedResources(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	var output accessanalyzer.ListAnalyzedResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListAnalyzedResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAnalyzedResourcesAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) *AccessAnalyzerListAnalyzedResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListAnalyzedResources", input)
	return &AccessAnalyzerListAnalyzedResourcesFuture{Future: future}
}

func (a *stub) ListAnalyzers(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
	var output accessanalyzer.ListAnalyzersOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListAnalyzers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAnalyzersAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) *AccessAnalyzerListAnalyzersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListAnalyzers", input)
	return &AccessAnalyzerListAnalyzersFuture{Future: future}
}

func (a *stub) ListArchiveRules(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
	var output accessanalyzer.ListArchiveRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListArchiveRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListArchiveRulesAsync(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) *AccessAnalyzerListArchiveRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListArchiveRules", input)
	return &AccessAnalyzerListArchiveRulesFuture{Future: future}
}

func (a *stub) ListFindings(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
	var output accessanalyzer.ListFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFindingsAsync(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) *AccessAnalyzerListFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListFindings", input)
	return &AccessAnalyzerListFindingsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
	var output accessanalyzer.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) *AccessAnalyzerListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.ListTagsForResource", input)
	return &AccessAnalyzerListTagsForResourceFuture{Future: future}
}

func (a *stub) StartResourceScan(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error) {
	var output accessanalyzer.StartResourceScanOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.StartResourceScan", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartResourceScanAsync(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) *AccessAnalyzerStartResourceScanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.StartResourceScan", input)
	return &AccessAnalyzerStartResourceScanFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error) {
	var output accessanalyzer.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *accessanalyzer.TagResourceInput) *AccessAnalyzerTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.TagResource", input)
	return &AccessAnalyzerTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error) {
	var output accessanalyzer.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) *AccessAnalyzerUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UntagResource", input)
	return &AccessAnalyzerUntagResourceFuture{Future: future}
}

func (a *stub) UpdateArchiveRule(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
	var output accessanalyzer.UpdateArchiveRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UpdateArchiveRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) *AccessAnalyzerUpdateArchiveRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UpdateArchiveRule", input)
	return &AccessAnalyzerUpdateArchiveRuleFuture{Future: future}
}

func (a *stub) UpdateFindings(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error) {
	var output accessanalyzer.UpdateFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UpdateFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFindingsAsync(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) *AccessAnalyzerUpdateFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.accessanalyzer.UpdateFindings", input)
	return &AccessAnalyzerUpdateFindingsFuture{Future: future}
}
