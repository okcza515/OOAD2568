package controller

import (
	"gorm.io/gorm"
)

type EvalModuleWrapper struct {
	ProgressController   *ProgressController
	AssessmentController *AssessmentController
	EvaluationController *EvaluationController
}

func NewEvalModuleWrapper(db *gorm.DB) *EvalModuleWrapper {
	return &EvalModuleWrapper{
		ProgressController:   NewProgressController(db),
		AssessmentController: NewAssessmentController(db),
		EvaluationController: NewEvaluationController(db),
	}
}
