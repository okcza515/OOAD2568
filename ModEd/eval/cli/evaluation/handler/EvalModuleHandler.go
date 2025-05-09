package handler

import (
	"ModEd/asset/util"

	"ModEd/core/cli"

	"ModEd/eval/controller"

	"fmt"

	"gorm.io/gorm"
)

type EvalModuleMenuStateHandler struct {
	Manager *cli.CLIMenuStateManager
	wrapper *controller.EvalModuleWrapper

	ProgressMenuStateHandler *progressHandler.ProgressMenuStateHandler
}

func NewEvalModuleHandler(manager *cli.CLIMenuStateManager, db *gorm.DB) *EvalModuleMenuStateHandler {
	wrapper := controller.NewEvalModuleWrapper(db)
	evalModuleHandler := &EvalModuleMenuStateHandler{
		Manager: manager,
		wrapper: wrapper,
	}

	evalModuleHandler.ProgressMenuStateHandler = progressHandler.NewProgressMenuStateHandler(manager, wrapper, evalModuleHandler)

	evalModuleHandler.Manager.AddMenu("3", evalModuleHandler.ProgressMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("Exit", nil)

	return evalModuleHandler
}

func (handler *EvalModuleMenuStateHandler) Render() {
	fmt.Println("\nEvaluation Module Menu:")
	fmt.Println("3. Progress")
	fmt.Println("Exit the Evaluation Module")
}

func (handler *EvalModuleMenuStateHandler) HandleUserInput(input string) error {
	err := handler.Manager.GoToMenu(input)

	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		util.PressEnterToContinue()
	}

	return nil
}
