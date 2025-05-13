//MEP-1006

package controller

import (
	"gorm.io/gorm"
)

type EvalModuleWrapper struct {
	ProgressController             *ProgressController
	AssignmentController           *AssignmentController
	AssignmentSubmissionController *AssignmentSubmissionController
	EvaluationController           *EvaluationController
	QuizController                 *QuizControllerAdapter
}

func NewEvalModuleWrapper(db *gorm.DB) *EvalModuleWrapper {
	return &EvalModuleWrapper{
		ProgressController:             NewProgressController(db),
		AssignmentController:           NewAssignmentController(db),
		AssignmentSubmissionController: NewAssignmentSubmissionController(db),
		EvaluationController:           NewEvaluationController(db),
		QuizController:                 NewQuizControllerAdapter(db),
	}
}
