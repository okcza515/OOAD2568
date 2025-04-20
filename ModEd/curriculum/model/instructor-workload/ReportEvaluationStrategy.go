package model

import (
	projectModel "ModEd/project/model"
	"fmt"
)

type ReportEvaluationStrategy struct{}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating report criteria...")
	return 0.0
}
