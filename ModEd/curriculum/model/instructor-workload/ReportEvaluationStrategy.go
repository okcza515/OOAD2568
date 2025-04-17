package model

import projectModel "ModEd/project/model"

type ReportEvaluationStrategy struct{}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) float64 {
	panic("unimplemented")
}
