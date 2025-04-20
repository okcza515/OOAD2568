package controller

import (
	model "ModEd/curriculum/model/instructor-workload"
	"fmt"

	projectModel "ModEd/project/model"
)

type EvaluationContext struct {
	Strategy model.EvaluationStrategy
}

type PresentationEvaluationStrategy struct{}
type AssignmentEvaluationStrategy struct{}
type ReportEvaluationStrategy struct{}

func NewEvaluation(strategyType string) *EvaluationContext {
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
		return &ReportEvaluationStrategy{}
	case "presentation":
		return &PresentationEvaluationStrategy{}
	case "assignment":
		return &AssignmentEvaluationStrategy{}
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

func (a *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating report criteria...")
	return 0.0
}

func (a *AssignmentEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating report criteria...")
	return 0.0
}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating report criteria...")
	return 0.0
}
