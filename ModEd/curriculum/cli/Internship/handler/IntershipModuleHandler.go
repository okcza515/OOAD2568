package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"fmt"
)

type InternShipModuleMenuStateHandler struct {
	menuManager *cli.CLIMenuStateManager
	wrapper     *controller.InternshipModuleWrapper

	InternshipApplicationMenuStateHandler *InternshipApplicationHandler
	InternshipEvaluationCriteriaHandler   *InternShipEvaluationCriteriaHandler
	InternshipResultEvaluationHandler     *InternshipResultEvaluationHandler
	InternshipInformationHandler          *InternshipInformationHandler
}

func NewInternShipModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternShipModuleMenuStateHandler {
	InternshipModule := &InternShipModuleMenuStateHandler{
		menuManager: manager,
		wrapper:     wrapper,
	}
	InternshipModule.InternshipApplicationMenuStateHandler = NewInternshipApplicationHandler(manager, wrapper)
	InternshipModule.InternshipEvaluationCriteriaHandler = NewInternShipEvaluationCriteriaHandler(manager, wrapper.InternshipCriteriaController)
	InternshipModule.InternshipResultEvaluationHandler = NewInternshipResultEvaluationHandler(manager, wrapper.InternshipResultEvaluationController)
	InternshipModule.InternshipInformationHandler = NewInternshipInformationHandler(manager, wrapper.InformationController)
	return InternshipModule
}

func (handler *InternShipModuleMenuStateHandler) Render() {
	fmt.Println("\n==== Internship Application System ====")
	fmt.Println("1. Load csv data")
	fmt.Println("2. Application Management")
	fmt.Println("3. Evaluate Student Performance")
	fmt.Println("5. Manage Evaluation Criteria")
	fmt.Println("6. Manage Result Evaluations")
	fmt.Println("7. Manage Internship Information")
	fmt.Println("Type 'exit' to quit")
	fmt.Print("Enter your choice: ")
}

func (handler *InternShipModuleMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		handler.wrapper.GenericImport.ImportCompanies("")
		handler.wrapper.GenericImport.ImportInternStudents("")
		return nil
	case "2":
		handler.menuManager.SetState(handler.InternshipApplicationMenuStateHandler)
		return nil
	case "3":
		handler.menuManager.SetState(handler.InternshipResultEvaluationHandler)
		return nil
	case "5":
		handler.menuManager.SetState(handler.InternshipEvaluationCriteriaHandler)
		return nil
	case "6":
		handler.menuManager.SetState(handler.InternshipResultEvaluationHandler)
		return nil
	case "7":
		handler.menuManager.SetState(handler.InternshipInformationHandler)
		return nil
	case "exit":
		fmt.Println("Exiting...")
		return nil
	default:
		fmt.Println("Invalid input")
		return nil
	}
}
