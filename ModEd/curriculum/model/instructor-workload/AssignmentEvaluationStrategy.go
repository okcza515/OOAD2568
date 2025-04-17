package model

import projectModel "ModEd/project/model"

type AssignmentEvaluationStrategy struct{}

func (a *AssignmentEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	panic("unimplemented")
}
