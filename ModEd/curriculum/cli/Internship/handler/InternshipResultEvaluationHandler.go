package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternshipResultEvaluationHandler struct {
	controller *controller.InternshipResultEvaluationController
}

func NewInternshipResultEvaluationHandler(controller *controller.InternshipResultEvaluationController) *InternshipResultEvaluationHandler {
	return &InternshipResultEvaluationHandler{controller: controller}
}

func (handler *InternshipResultEvaluationHandler) Render() {
	fmt.Println("\n==== Internship Result Evaluation Menu ====")
	fmt.Println("1. Create Result Evaluation")
	fmt.Println("2. Retrieve Result Evaluation by ID")
	fmt.Println("3. Update Result Evaluation")
	fmt.Println("4. Delete Result Evaluation by ID")
	fmt.Println("5. List All Result Evaluations")
	fmt.Println("Type 'back' to return to the previous menu")
	fmt.Print("Enter your choice: ")
}

func (handler *InternshipResultEvaluationHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		return handler.createResultEvaluation()
	case "2":
		return handler.retrieveResultEvaluationByID()
	case "3":
		return handler.updateResultEvaluation()
	case "4":
		return handler.deleteResultEvaluationByID()
	case "5":
		return handler.listAllResultEvaluations()
	case "back":
		fmt.Println("Returning to the previous menu...")
		return nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil
	}
}

func (handler *InternshipResultEvaluationHandler) createResultEvaluation() error {
	comment := utils.GetUserInput("Enter Comment: ")
	score := utils.GetUserInputUint("Enter Score (0-100): ")
	internshipInfoID := utils.GetUserInputUint("Enter Internship Information ID: ")

	evaluation := &model.InternshipResultEvaluation{
		Comment:                 comment,
		Score:                   score,
		InternshipInformationId: internshipInfoID,
	}

	if err := handler.controller.Create(evaluation); err != nil {
		return fmt.Errorf("failed to create result evaluation: %w", err)
	}

	fmt.Println("Result evaluation created successfully!")
	return nil
}

func (handler *InternshipResultEvaluationHandler) retrieveResultEvaluationByID() error {
	id := utils.GetUserInputUint("Enter Result Evaluation ID: ")

	evaluation, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve result evaluation: %w", err)
	}

	fmt.Printf("Result Evaluation: ID: %d, Comment: %s, Score: %d, Internship Information ID: %d\n",
		evaluation.ID, evaluation.Comment, evaluation.Score, evaluation.InternshipInformationId)
	return nil
}

func (handler *InternshipResultEvaluationHandler) updateResultEvaluation() error {
	id := utils.GetUserInputUint("Enter Result Evaluation ID to update: ")

	evaluation, err := handler.controller.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve result evaluation: %w", err)
	}

	evaluation.Comment = utils.GetUserInput("Enter new Comment: ")
	evaluation.Score = utils.GetUserInputUint("Enter new Score (0-100): ")

	if err := handler.controller.Update(evaluation); err != nil {
		return fmt.Errorf("failed to update result evaluation: %w", err)
	}

	fmt.Println("Result evaluation updated successfully!")
	return nil
}

func (handler *InternshipResultEvaluationHandler) deleteResultEvaluationByID() error {
	id := utils.GetUserInputUint("Enter Result Evaluation ID to delete: ")

	if err := handler.controller.DeleteByID(id); err != nil {
		return fmt.Errorf("failed to delete result evaluation: %w", err)
	}

	fmt.Println("Result evaluation deleted successfully!")
	return nil
}

func (handler *InternshipResultEvaluationHandler) listAllResultEvaluations() error {
	evaluations, err := handler.controller.ListAll()
	if err != nil {
		return fmt.Errorf("failed to list result evaluations: %w", err)
	}

	fmt.Println("All Result Evaluations:")
	for _, evaluation := range evaluations {
		fmt.Printf("ID: %d, Comment: %s, Score: %d, Internship Information ID: %d\n",
			evaluation.ID, evaluation.Comment, evaluation.Score, evaluation.InternshipInformationId)
	}
	return nil
}
