package command

import (
	"ModEd/eval/cli/evaluation/menu"
	controller "ModEd/eval/controller"

	"gorm.io/gorm"
)

type EvaluationCommand struct {
	DB                   *gorm.DB
	EvaluationController *controller.EvaluationController
	ProgressController   *controller.ProgressController
	AssignmentController *controller.AssignmentController
}

func (e *EvaluationCommand) Execute() error {
	menu.RunEvalModuleCLI(e.DB, e.EvaluationController, e.ProgressController, e.AssignmentController)
	return nil
}
