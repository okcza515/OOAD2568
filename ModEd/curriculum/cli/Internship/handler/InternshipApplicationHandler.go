package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	// "strconv"
	"time"
)

type InternshipApplicationHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InternshipModuleWrapper
}

func NewInternshipApplicationHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternshipApplicationHandler {
	return &InternshipApplicationHandler{
		manager: manager,
		wrapper: wrapper,
	}
}

func (handler *InternshipApplicationHandler) handleCreateInternshipApplication() error {

	studentCode := utils.GetUserInput("Enter Student Code: ")
	if len(studentCode) != 11 {
		fmt.Println("Student Code cannot be empty.")
		return fmt.Errorf("invalid input: student code is empty")
	}

	companyName := utils.GetUserInput("Enter Company Name: ")
	if companyName == "" {
		fmt.Println("Company Name cannot be empty.")
		return fmt.Errorf("invalid input: company name is empty")
	}

	company, err := handler.wrapper.Company.GetCompanyByName(companyName)
	if err != nil {
		fmt.Printf("Error finding company with name '%s': %v\n", companyName, err)
		return fmt.Errorf("failed to find company: %w", err)
	}

	// advisorCodeStr := utils.GetUserInput("Enter Advisor Code: ")
	// advisorCode, err := strconv.Atoi(advisorCodeStr)
	// if err != nil || advisorCode <= 0 {
	// 	fmt.Println("Invalid Advisor Code. Please enter a positive integer.")
	// 	return fmt.Errorf("invalid input: advisor code must be a positive integer")
	// }

	application := &model.InternshipApplication{
		TurninDate:            time.Now(),
		ApprovalAdvisorStatus: model.WAIT,
		ApprovalCompanyStatus: model.WAIT,
		// AdvisorCode:           uint(advisorCode),
		CompanyId:             company.ID,
		StudentCode:           studentCode,
	}

	err = handler.wrapper.InternshipApplication.RegisterInternshipApplications([]*model.InternshipApplication{application})
	if err != nil {
		fmt.Println("Error creating internship application:", err)
		return fmt.Errorf("failed to register internship application: %w", err)
	}

	fmt.Println("Internship application created successfully!")
	return nil
}

func (handler *InternshipApplicationHandler) ListApplications() error {
	applications, err := handler.wrapper.InternshipApplication.GetAllInternshipApplications()
	if err != nil {
		return fmt.Errorf("failed to retrieve internship applications: %w", err)
	}

	fmt.Println("Internship Applications:")
	for _, app := range applications {
		fmt.Printf("ID: %d, StudentCode: %s, AdvisorStatus: %s, CompanyStatus: %s\n",
			app.ID, app.StudentCode, app.ApprovalAdvisorStatus, app.ApprovalCompanyStatus)
	}
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
		return handler.handleCreateInternshipApplication()
	case "2":
		return handler.ListApplications()
	case "3":
		return handler.GetApplicationStatus()
	case "4":
		return handler.DeleteApplication()
	case "back":
		fmt.Println("Returning to the previous menu...")
		return nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil
	}
}

func (handler *InternshipApplicationHandler) Render() {
	fmt.Println("\n==== Internship Application Menu ====")
	fmt.Println("1. Create Internship Application")
	fmt.Println("2. List All Internship Applications")
	fmt.Println("3. Get Application Status")
	fmt.Println("4. Delete Internship Application")
	fmt.Println("Type 'back' to return to the previous menu")
	fmt.Print("Enter your choice: ")
}
