package progress

import (
	"ModEd/core/cli"
	"ModEd/eval/cli/progress/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

func RunProgressModuleCLI(db *gorm.DB) {
	evaluations := make([]*model.Evaluation, 0)
	wrapper := &controller.EvalModuleWrapper{
		ProgressController:   controller.NewProgressController(db),
		AssessmentController: controller.NewAssessmentController(),
		EvaluationController: controller.NewEvaluationController(evaluations, ""),
	}

	manager := cli.NewCLIMenuManager()

	progressHandler := handler.NewProgressMenuStateHandler(manager, wrapper, nil)

	manager.SetState(progressHandler)

	for {
		progressHandler.Render()
		var input string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		err := progressHandler.HandleUserInput(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
