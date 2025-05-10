package evaluation

import (
	"ModEd/core/cli"

	"fmt"

	"ModEd/eval/cli/evaluation/handler"

	evalController "ModEd/eval/controller"

	"gorm.io/gorm"
)

func RunEvalModuleCLI(db *gorm.DB,
	evaluationController *evalController.EvaluationController,
	progressController *evalController.ProgressController,
	assessmentController *evalController.AssessmentController,
) {
	Manager := cli.NewCLIMenuManager()
	wrapper := evalController.NewEvalModuleWrapper(db)
	evalModuleState := handler.NewEvalModuleHandler(Manager, wrapper)
	Manager.SetState(evalModuleState)

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
