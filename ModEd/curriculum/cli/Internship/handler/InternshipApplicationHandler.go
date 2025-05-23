package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternshipApplicationHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InternshipModuleWrapper

	InternshipModule *InternShipModuleMenuStateHandler
}

func NewInternshipApplicationHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternshipApplicationHandler {
	return &InternshipApplicationHandler{
		manager: manager,
		wrapper: wrapper,
	}
}

func (handler *InternshipApplicationHandler) handleApplicationByRole() error {
	application := &model.InternshipApplication{}
	role := utils.GetUserInput("Role: (student/university/company): ")

	err := controller.SubmitApplication(application, role, handler.wrapper.InternshipApplication.Connector)
	if err != nil {
			return fmt.Errorf("failed to submit application: %w", err)
	}

	fmt.Println("Application submitted successfully!")
	return nil
}

func (handler *InternshipApplicationHandler) GetApplicationStatus() error {
	id := utils.GetUserInput("\nWhat is the ID of the internship application you want to check? ")

	status, err := handler.wrapper.InternshipApplication.GetApplicationStatusByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve application status: %w", err)
	}

	fmt.Printf("Application Status for ID %s: %s\n", id, status)
	return nil
}

func (handler *InternshipApplicationHandler) DeleteApplication() error {
	id := utils.GetUserInput("Enter the ID of the internship application to delete: ")

	err := handler.wrapper.InternshipApplication.DeleteApplicationByID(id)
	if err != nil {
		return fmt.Errorf("failed to delete internship application: %w", err)
	}

	fmt.Printf("Internship application with ID %s deleted successfully!\n", id)
	return nil
}

func (handler *InternshipApplicationHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		return handler.handleApplicationByRole()
	case "2":
		return handler.GetApplicationStatus()
	case "3":
		return handler.DeleteApplication()
	case "0":
		handler.manager.SetState(handler.InternshipModule)
		return nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil
	}
}

func (handler *InternshipApplicationHandler) Render() {
	fmt.Println("\n==== Internship Application Menu ====")
	fmt.Println("1. Create Internship Application")
	fmt.Println("2. Get Application Status")
	fmt.Println("3. Delete Internship Application")
	fmt.Println("0. return to the previous menu")
	fmt.Print("Enter your choice: ")
}
