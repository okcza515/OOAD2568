package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	projectModel "ModEd/project/model"
)

type EvaluationContext struct {
	Strategy model.EvaluationStrategy
}

func NewEvaluationContext(strategyType string) *EvaluationContext {
	return &EvaluationContext{
		Strategy: GetEvaluationStrategy(strategyType),
	}
}

func (ctx *EvaluationContext) SetStrategy(strategy model.EvaluationStrategy) {
	ctx.Strategy = strategy
}

func GetEvaluationStrategy(strategyType string) model.EvaluationStrategy {
	switch strategyType {
	case "report":
		return &model.ReportEvaluationStrategy{}
	case "presentation":
		return &model.PresentationEvaluationStrategy{}
	case "assignment":
		return &model.AssignmentEvaluationStrategy{}
	default:
		return nil
	}
}

func (ctx *EvaluationContext) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	if ctx.Strategy == nil {
		return 0
	}
	return ctx.Strategy.Evaluate(criteria)
}
