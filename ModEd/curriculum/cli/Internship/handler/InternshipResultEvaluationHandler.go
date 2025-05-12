package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternshipResultEvaluationHandler struct {
	manager                    *cli.CLIMenuStateManager
	controller                 *controller.InternshipResultEvaluationController
	InternshipInformation      *controller.InternshipInformationController
	InternshipCriteria         *controller.InternshipCriteriaController
	InternshipResultEvaluation *controller.InternshipResultEvaluationController
}

func NewInternshipResultEvaluationHandler(
	manager *cli.CLIMenuStateManager,
	resultCtrl *controller.InternshipResultEvaluationController,
	infoCtrl *controller.InternshipInformationController,
	criteriaCtrl *controller.InternshipCriteriaController,
) *InternshipResultEvaluationHandler {
	return &InternshipResultEvaluationHandler{
		manager:                    manager,
		controller:                 resultCtrl,
		InternshipInformation:      infoCtrl,
		InternshipCriteria:         criteriaCtrl,
		InternshipResultEvaluation: resultCtrl,
	}
}

func (handler *InternshipResultEvaluationHandler) Render() {
	fmt.Println("\n==== Internship Result Evaluation Menu ====")
	fmt.Println("1. Create Result Evaluation")
	fmt.Println("2. Retrieve Result Evaluation by ID")
	fmt.Println("3. Update Result Evaluation")
	fmt.Println("4. Delete Result Evaluation by ID")
	fmt.Println("5. List All Result Evaluations")
	fmt.Println("6. Evaluate Student Internship")
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
	case "6":
		return handler.EvaluateStudentInternship()
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

func (handler *InternshipResultEvaluationHandler) EvaluateStudentInternship() error {
	studentCode := utils.GetUserInput("Enter Student Code: ")
	if studentCode == "" {
		fmt.Println("Error: Student Code cannot be empty.")
		return nil
	}

	if handler.InternshipInformation == nil || handler.InternshipCriteria == nil || handler.InternshipResultEvaluation == nil {
		fmt.Println("Error: Missing required controllers for evaluation.")
		return fmt.Errorf("missing required controllers")
	}

	criteriaList, err := handler.InternshipCriteria.ListAllByStudentCode(studentCode)
	if err != nil {
		fmt.Printf("Failed to retrieve criteria for student %s: %v\n", studentCode, err)
		return err
	}

	if len(criteriaList) == 0 {
		fmt.Printf("No criteria found for student %s.\n", studentCode)
		return nil
	}

	fmt.Println("Criteria for the student:")
	for _, criteria := range criteriaList {
		fmt.Printf("Criteria ID: %d, Title: %s\n", criteria.ID, criteria.Title)
	}

	criteriaScores := map[uint]uint{}
	for _, criteria := range criteriaList {
		score := utils.GetUserInputUint(fmt.Sprintf("Enter Score for Criteria ID %d (1-5): ", criteria.ID))
		if score < 1 || score > 5 {
			fmt.Println("Invalid score. Please enter a value between 1 and 5.")
			continue
		}
		criteriaScores[criteria.ID] = score
	}

	comment := utils.GetUserInput("Enter Evaluation Comment: ")
	if comment == "" {
		fmt.Println("Error: Comment cannot be empty.")
		return nil
	}

	facade := controller.NewInternshipEvaluationFacade(
		*handler.InternshipInformation,
		*handler.InternshipCriteria,
		*handler.InternshipResultEvaluation,
	)

	err = facade.EvaluateInternship(studentCode, criteriaScores, comment)
	if err != nil {
		fmt.Printf("Evaluation failed: %v\n", err)
		return err
	}

	fmt.Println("Evaluation completed successfully!")
	return nil
}
