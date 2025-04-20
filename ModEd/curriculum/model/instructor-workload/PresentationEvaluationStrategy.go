package model

import (
	projectModel "ModEd/project/model"
	"fmt"
)

type PresentationEvaluationStrategy struct{}

func (a *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	fmt.Println("Evaluating report criteria...")
	return 0.0
}
