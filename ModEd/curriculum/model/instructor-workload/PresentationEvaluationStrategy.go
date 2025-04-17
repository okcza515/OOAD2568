package model

import projectModel "ModEd/project/model"

type PresentationEvaluationStrategy struct{}

func (a *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	panic("unimplemented")
}
