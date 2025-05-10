package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternShipModuleMenuStateHandler struct {
	menuManager *cli.CLIMenuStateManager
	wrapper     *controller.InternshipModuleWrapper

	InternshipApplicationMenuStateHandler *InternshipApplicationHandler
}

func NewInternShipModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternShipModuleMenuStateHandler {
	InternshipModule := &InternShipModuleMenuStateHandler{
		menuManager: manager,
		wrapper:     wrapper,
	}
	InternshipModule.InternshipApplicationMenuStateHandler = NewInternshipApplicationHandler(manager, wrapper)

	return InternshipModule
}

func (handler *InternShipModuleMenuStateHandler) Render() {
	fmt.Println("\n==== Internship Application System ====")
	fmt.Println("1. Load csv data")
	fmt.Println("2. Application Management")
	fmt.Println("3. Evaluate Student Performance")
	fmt.Println("4. Update Approval Status")
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
		return nil
	case "4":
		err := handler.handleUpdateApprovalStatus()
		if err != nil {
			fmt.Println("Error updating approval status:", err)
		}
		return err
	case "exit":
		fmt.Println("Exiting...")
		return nil
	default:
		fmt.Println("Invalid input")
		return nil
	}
}

func (handler *InternShipModuleMenuStateHandler) handleUpdateApprovalStatus() error {
	studentCode := utils.GetUserInput("Enter Student Code: ")
	if studentCode == "" {
		return fmt.Errorf("error: student code cannot be empty")
	}

	advisorStatusStr := utils.GetUserInput("Enter Advisor Approval Status (APPROVED/REJECT): ")
	if advisorStatusStr != string(model.APPROVED) && advisorStatusStr != string(model.REJECT) {
		return fmt.Errorf("error: invalid advisor approval status, must be 'APPROVED' or 'REJECT'")
	}
	advisorStatus := model.ApprovedStatus(advisorStatusStr)

	companyStatusStr := utils.GetUserInput("Enter Company Approval Status (APPROVED/REJECT): ")
	if companyStatusStr != string(model.APPROVED) && companyStatusStr != string(model.REJECT) {
		return fmt.Errorf("error: invalid company approval status, must be 'APPROVED' or 'REJECT'")
	}
	companyStatus := model.ApprovedStatus(companyStatusStr)

	err := handler.wrapper.Approved.UpdateApprovalStatuses(studentCode, advisorStatus, companyStatus)
	if err != nil {
		return fmt.Errorf("error updating approval statuses: %w", err)
	}

	fmt.Println("Approval statuses updated successfully!")
	return nil
}
