// MEP-1008
package model

import (
	projectModel "ModEd/project/model"
)

type EvaluationContext struct {
	Strategy EvaluationStrategy
}

func NewEvaluationContext(strategyType string) *EvaluationContext {
	return &EvaluationContext{
		Strategy: GetEvaluationStrategy(strategyType),
	}
}

func (ctx *EvaluationContext) SetStrategy(strategy EvaluationStrategy) {
	ctx.Strategy = strategy
}

func (ctx *EvaluationContext) Evaluate(evaluation *ProjectEvaluation, criteria []projectModel.AssessmentCriteria) (float64, string) {
	if ctx.Strategy == nil {
		return 0, "error: no strategy set"
	}
	return ctx.Strategy.Evaluate(criteria)
}

func GetEvaluationStrategy(strategyType string) EvaluationStrategy {
	switch strategyType {
	case "report":
		return &ReportEvaluationStrategy{}
	case "presentation":
		return &PresentationEvaluationStrategy{}
	case "assignment":
		return &AssignmentEvaluationStrategy{}
	default:
		return nil
	}
}
