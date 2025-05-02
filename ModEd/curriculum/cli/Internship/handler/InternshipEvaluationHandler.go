package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type InternshipEvaluationHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InternshipModuleWrapper
}

func NewIInternshipEvaluationHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternshipEvaluationHandler {
	return &InternshipEvaluationHandler{
		manager: manager,
		wrapper: wrapper,
	}
}

func (handler *InternshipEvaluationHandler) Render() {
	fmt.Println("\n==== Internship Evaluation System ====")
	fmt.Println("1. Review Evaluation")
	fmt.Println("2. Report Evaluation")
	fmt.Println("3. Search Review and Report Scores")
	fmt.Println("4. Update Review Scores")
	fmt.Println("5. Update Report Scores")
	fmt.Println("Type 'exit' to quit")
	fmt.Print("Enter your choice: ")
}
