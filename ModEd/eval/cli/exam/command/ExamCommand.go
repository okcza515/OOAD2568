package command

//MEP-1007

import (
	"ModEd/eval/cli/exam/menu"
	controller "ModEd/eval/controller"

	"gorm.io/gorm"
)

type ExaminationCommand struct {
	DB                   *gorm.DB
	ExamController       *controller.ExamController
	QuestionController   *controller.QuestionController
	SubmissionController *controller.SubmissionController
}

func (e *ExaminationCommand) Execute() error {
	menu.RunExamModuleCLI(e.DB, e.ExamController, e.QuestionController, e.SubmissionController)
	return nil
}