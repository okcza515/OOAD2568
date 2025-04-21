package model

import (
	projectModel "ModEd/project/model"
	"fmt"
)

type EvaluationStrategy interface {
	Evaluate(criteria []projectModel.AssessmentCriteria) float64
}

type PresentationEvaluationStrategy struct{}
type AssignmentEvaluationStrategy struct{}
type ReportEvaluationStrategy struct{}

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
