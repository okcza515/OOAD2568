package model

import projectModel "ModEd/project/model"

type EvaluationStrategy interface {
	Evaluate(criteria []projectModel.AssessmentCriteria) float64
}
