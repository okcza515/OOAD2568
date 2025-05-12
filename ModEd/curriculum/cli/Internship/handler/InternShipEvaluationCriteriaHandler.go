package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternShipEvaluationCriteriaHandler struct {
	manager          *cli.CLIMenuStateManager
	controller       *controller.InternshipCriteriaController
	InternshipModule *InternShipModuleMenuStateHandler
}

func NewInternShipEvaluationCriteriaHandler(manager *cli.CLIMenuStateManager, controller *controller.InternshipCriteriaController) *InternShipEvaluationCriteriaHandler {
	return &InternShipEvaluationCriteriaHandler{
		manager:    manager,
		controller: controller,
	}
}

func (handler *InternShipEvaluationCriteriaHandler) Render() {
	fmt.Println("\n==== Internship Evaluation System ====")
	fmt.Println("1. Create Evaluation Criteria")
	fmt.Println("2. Retrieve Evaluation Criteria by ID")
	fmt.Println("3. Update Evaluation Criteria")
	fmt.Println("4. Delete Evaluation Criteria by ID")
	fmt.Println("5. List All Evaluation Criteria")
	fmt.Println("Type 'exit' to quit")
	fmt.Print("Enter your choice: ")
}

func (handler *InternShipEvaluationCriteriaHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		return handler.createEvaluationCriteria()
	case "2":
		return handler.retrieveEvaluationCriteriaByID()
	case "3":
		return handler.updateEvaluationCriteria()
	case "4":
		return handler.deleteEvaluationCriteriaByID()
	case "5":
		return handler.listAllEvaluationCriteria()
	case "exit":
		fmt.Println("Exiting the Internship Evaluation System...")
		handler.manager.SetState(handler.InternshipModule)
		return nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil
	}
}

func (handler *InternShipEvaluationCriteriaHandler) createEvaluationCriteria() error {
	title := utils.GetUserInput("Enter Evaluation Criteria Title: ")
	if title == "" {
		fmt.Println("Error: Title cannot be empty.")
		return nil
	}

	description := utils.GetUserInput("Enter Evaluation Criteria Description: ")
	if description == "" {
		fmt.Println("Error: Description cannot be empty.")
		return nil
	}

	applicationID := utils.GetUserInputUint("Enter Internship Application ID: ")
	if applicationID == 0 {
		fmt.Println("Error: Internship Application ID must be valid.")
		return nil
	}

	criteria := &model.InternshipCriteria{
		Title:                   title,
		Description:             description,
		Score:                   0,
		InternshipApplicationId: applicationID,
	}

	if err := handler.controller.Create(criteria); err != nil {
		return fmt.Errorf("failed to create evaluation criteria: %w", err)
	}

	fmt.Println("Evaluation criteria created successfully!")
	return nil
}

func (handler *InternShipEvaluationCriteriaHandler) retrieveEvaluationCriteriaByID() error {
	id := utils.GetUserInputUint("Enter Evaluation Criteria ID: ")
	if id == 0 {
		fmt.Println("Error: Invalid ID.")
		return nil
	}

	criteria, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve evaluation criteria: %w", err)
	}

	fmt.Printf("Evaluation Criteria: ID: %d, Title: %s, Description: %s, Score: %d, Application ID: %d\n",
		criteria.ID, criteria.Title, criteria.Description, criteria.Score, criteria.InternshipApplicationId)
	return nil
}

func (handler *InternShipEvaluationCriteriaHandler) updateEvaluationCriteria() error {
	id := utils.GetUserInputUint("Enter Evaluation Criteria ID to update: ")
	if id == 0 {
		fmt.Println("Error: Invalid ID.")
		return nil
	}

	criteria, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve evaluation criteria: %w", err)
	}

	title := utils.GetUserInput("Enter new Evaluation Criteria Title (leave blank to keep current): ")
	if title != "" {
		criteria.Title = title
	}

	description := utils.GetUserInput("Enter new Evaluation Criteria Description (leave blank to keep current): ")
	if description != "" {
		criteria.Description = description
	}

	score := utils.GetUserInputUint("Enter new Evaluation Score (e.g., 100, leave blank to keep current): ")
	if score > 0 {
		criteria.Score = score
	}

	if err := handler.controller.Update(criteria); err != nil {
		return fmt.Errorf("failed to update evaluation criteria: %w", err)
	}

	fmt.Println("Evaluation criteria updated successfully!")
	return nil
}

func (handler *InternShipEvaluationCriteriaHandler) deleteEvaluationCriteriaByID() error {
	id := utils.GetUserInputUint("Enter Evaluation Criteria ID to delete: ")
	if id == 0 {
		fmt.Println("Error: Invalid ID.")
		return nil
	}

	if err := handler.controller.DeleteByID(id); err != nil {
		return fmt.Errorf("failed to delete evaluation criteria: %w", err)
	}

	fmt.Println("Evaluation criteria deleted successfully!")
	return nil
}

func (handler *InternShipEvaluationCriteriaHandler) listAllEvaluationCriteria() error {
	criteriaList, err := handler.controller.ListAll()
	if err != nil {
		return fmt.Errorf("failed to list evaluation criteria: %w", err)
	}

	if len(criteriaList) == 0 {
		fmt.Println("No evaluation criteria found.")
		return nil
	}

	fmt.Println("All Evaluation Criteria:")
	for _, criteria := range criteriaList {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Score: %d, Application ID: %d\n",
			criteria.ID, criteria.Title, criteria.Description, criteria.Score, criteria.InternshipApplicationId)
	}
	return nil
}
