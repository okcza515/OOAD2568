package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type EvalModuleWrapper struct {
	ProgressController   *ProgressController
	AssessmentController *AssessmentController
	EvaluationController *EvaluationController
}

func NewEvalModuleWrapper(db *gorm.DB) *EvalModuleWrapper {
	evaluations := make([]*model.Evaluation, 0)
	csvPath := ""

	return &EvalModuleWrapper{
		ProgressController:   NewProgressController(db),
		AssessmentController: NewAssessmentController(db),
		EvaluationController: NewEvaluationController(evaluations, csvPath),
	}
}
