package menu

import (
	"ModEd/core/cli"
	"ModEd/eval/cli/exam/handler"
	"ModEd/eval/controller"

	"fmt"
	"gorm.io/gorm"
)

func RunExamModuleCLI(
	db *gorm.DB,
	examController *controller.ExamController,
	questionController *controller.QuestionController,
	submissionController *controller.SubmissionController,
) {
	Manager := cli.NewCLIMenuManager()
	wrapper := controller.NewExamModuleWrapper(db)
	examModuleState := handler.NewExamModuleHandler(Manager, wrapper)
	Manager.SetState(examModuleState)

	for {
		Manager.Render()
		Manager.UserInput = GetUserChoice()

		if Manager.UserInput == "Exit" {
			break
		}

		err := Manager.HandleUserInput()
		if err != nil {
			return
		}
	}
}

func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}
