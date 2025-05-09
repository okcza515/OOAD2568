package evaluation

import (
	"ModEd/eval/cli/evaluation/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

func RunEvaluationModuleCLI(db *gorm.DB) {
	// Initialize controllers
	evaluationController := controller.NewEvaluationController([]*model.Evaluation{}, "")

	params := &handler.EvaluationCLIParams{
		EvaluationController: evaluationController,
	}

	mainState := handler.NewMainMenuState(params)
	stateManager := handler.NewMenuStateManager(mainState)

	// Run menu state manager
	err := stateManager.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
