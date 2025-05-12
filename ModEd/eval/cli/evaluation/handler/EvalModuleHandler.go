package handler

import (
	"ModEd/asset/util"

	"ModEd/core/cli"

	"ModEd/eval/controller"

	"fmt"
)

type EvalModuleMenuStateHandler struct {
	Manager *cli.CLIMenuStateManager
	wrapper *controller.EvalModuleWrapper

	AssessmentMenuStateHandler *AssessmentMenuStateHandler
	ProgressMenuStateHandler   *ProgressMenuStateHandler
	EvaluationMenuStateHandler *EvaluationMenuStateHandler
}

func NewEvalModuleHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper) *EvalModuleMenuStateHandler {
	evalModuleHandler := &EvalModuleMenuStateHandler{
		Manager: manager,
		wrapper: wrapper,
	}

	evalModuleHandler.ProgressMenuStateHandler = NewProgressMenuStateHandler(manager, wrapper, evalModuleHandler)

	evalModuleHandler.Manager.AddMenu("1", evalModuleHandler.AssessmentMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("2", evalModuleHandler.AssessmentMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("3", evalModuleHandler.ProgressMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("4", evalModuleHandler.EvaluationMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("Exit", nil)

	return evalModuleHandler
}

func (handler *EvalModuleMenuStateHandler) Render() {
	fmt.Println("\nEvaluation Module Menu:")
	fmt.Println("1. Assessment")
	fmt.Println("2. Quiz")
	fmt.Println("3. Progress")
	fmt.Println("4. Evaluation")
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
