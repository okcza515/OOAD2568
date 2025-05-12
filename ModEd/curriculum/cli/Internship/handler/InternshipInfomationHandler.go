package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternshipInformationHandler struct {
	manager    *cli.CLIMenuStateManager
	controller *controller.InternshipInformationController

	InternshipModule *InternShipModuleMenuStateHandler
}

func NewInternshipInformationHandler(manager *cli.CLIMenuStateManager, controller *controller.InternshipInformationController) *InternshipInformationHandler {
	return &InternshipInformationHandler{
		manager:    manager,
		controller: controller}
}

func (handler *InternshipInformationHandler) Render() {
	fmt.Println("\n==== Internship Information Menu ====")
	fmt.Println("1. Create Internship Information")
	fmt.Println("2. Retrieve Internship Information by ID")
	fmt.Println("3. Update Internship Information")
	fmt.Println("4. Delete Internship Information by ID")
	fmt.Println("5. List All Internship Information")
	fmt.Println("Type 'back' to return to the previous menu")
	fmt.Print("Enter your choice: ")
}

func (handler *InternshipInformationHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		return handler.createInternshipInformation()
	case "2":
		return handler.retrieveInternshipInformationByID()
	case "3":
		return handler.updateInternshipInformation()
	case "4":
		return handler.deleteInternshipInformationByID()
	case "5":
		return handler.listAllInternshipInformation()
	case "back":
		fmt.Println("Returning to the previous menu...")
		handler.manager.SetState(handler.InternshipModule)
		return nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil
	}
}

func (handler *InternshipInformationHandler) createInternshipInformation() error {
	studentCode := utils.GetUserInput("Enter Student Code: ")
	companyId := utils.GetUserInputUint("Enter Company ID: ")
	mentorId := utils.GetUserInputUint("Enter Mentor ID: ")

	info := &model.InternshipInformation{
		StudentCode:        studentCode,
		CompanyId:          companyId,
		InternshipMentorID: mentorId,
	}

	if err := handler.controller.Create(info); err != nil {
		return fmt.Errorf("failed to create internship information: %w", err)
	}

	fmt.Println("Internship information created successfully!")
	return nil
}

func (handler *InternshipInformationHandler) retrieveInternshipInformationByID() error {
	id := utils.GetUserInputUint("Enter Internship Information ID: ")

	info, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship information: %w", err)
	}

	fmt.Printf("Internship Information: %+v\n", info)
	return nil
}

func (handler *InternshipInformationHandler) updateInternshipInformation() error {
	id := utils.GetUserInputUint("Enter Internship Information ID to update: ")

	info, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship information: %w", err)
	}

	info.CompanyId = utils.GetUserInputUint("Enter new Company ID: ")
	info.InternshipMentorID = utils.GetUserInputUint("Enter new Mentor ID: ")

	if err := handler.controller.Update(info); err != nil {
		return fmt.Errorf("failed to update internship information: %w", err)
	}

	fmt.Println("Internship information updated successfully!")
	return nil
}

func (handler *InternshipInformationHandler) deleteInternshipInformationByID() error {
	id := utils.GetUserInputUint("Enter Internship Information ID to delete: ")

	if err := handler.controller.DeleteByID(id); err != nil {
		return fmt.Errorf("failed to delete internship information: %w", err)
	}

	fmt.Println("Internship information deleted successfully!")
	return nil
}

func (handler *InternshipInformationHandler) listAllInternshipInformation() error {
	infos, err := handler.controller.ListAll()
	if err != nil {
		return fmt.Errorf("failed to list internship information: %w", err)
	}

	fmt.Println("All Internship Information:")
	for _, info := range infos {
		fmt.Printf("ID: %d, StudentCode: %s, CompanyID: %d, MentorID: %d\n",
			info.ID, info.StudentCode, info.CompanyId, info.InternshipMentorID)
	}
	return nil
}
