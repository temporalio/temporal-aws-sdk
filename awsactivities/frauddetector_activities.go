// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type FraudDetectorActivities struct {
	client frauddetectoriface.FraudDetectorAPI

	sessionFactory SessionFactory
}

func NewFraudDetectorActivities(sess *session.Session, config ...*aws.Config) *FraudDetectorActivities {
	client := frauddetector.New(sess, config...)
	return &FraudDetectorActivities{client: client}
}

func NewFraudDetectorActivitiesWithSessionFactory(sessionFactory SessionFactory) *FraudDetectorActivities {
	return &FraudDetectorActivities{sessionFactory: sessionFactory}
}

func (a *FraudDetectorActivities) getClient(ctx context.Context) (frauddetectoriface.FraudDetectorAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return frauddetector.New(sess), nil
}

func (a *FraudDetectorActivities) BatchCreateVariable(ctx context.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchCreateVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) BatchGetVariable(ctx context.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateDetectorVersion(ctx context.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateModel(ctx context.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateModelVersion(ctx context.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateRule(ctx context.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRuleWithContext(ctx, input)
}

func (a *FraudDetectorActivities) CreateVariable(ctx context.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVariableWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteDetector(ctx context.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteDetectorVersion(ctx context.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteEvent(ctx context.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DeleteRule(ctx context.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRuleWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DescribeDetector(ctx context.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) DescribeModelVersions(ctx context.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeModelVersionsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetDetectorVersion(ctx context.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetDetectors(ctx context.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDetectorsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEntityTypes(ctx context.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEntityTypesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEventPrediction(ctx context.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEventPredictionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetEventTypes(ctx context.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEventTypesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetExternalModels(ctx context.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetExternalModelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetKMSEncryptionKey(ctx context.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetKMSEncryptionKeyWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetLabels(ctx context.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLabelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetModelVersion(ctx context.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetModels(ctx context.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetModelsWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetOutcomes(ctx context.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOutcomesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetRules(ctx context.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRulesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) GetVariables(ctx context.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVariablesWithContext(ctx, input)
}

func (a *FraudDetectorActivities) ListTagsForResource(ctx context.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutDetector(ctx context.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDetectorWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutEntityType(ctx context.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEntityTypeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutEventType(ctx context.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEventTypeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutExternalModel(ctx context.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutExternalModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutKMSEncryptionKey(ctx context.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutKMSEncryptionKeyWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutLabel(ctx context.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutLabelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) PutOutcome(ctx context.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutOutcomeWithContext(ctx, input)
}

func (a *FraudDetectorActivities) TagResource(ctx context.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UntagResource(ctx context.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersion(ctx context.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDetectorVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionMetadata(ctx context.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDetectorVersionMetadataWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateDetectorVersionStatus(ctx context.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDetectorVersionStatusWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModel(ctx context.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateModelWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModelVersion(ctx context.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateModelVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateModelVersionStatus(ctx context.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateModelVersionStatusWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateRuleMetadata(ctx context.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRuleMetadataWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateRuleVersion(ctx context.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRuleVersionWithContext(ctx, input)
}

func (a *FraudDetectorActivities) UpdateVariable(ctx context.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVariableWithContext(ctx, input)
}
