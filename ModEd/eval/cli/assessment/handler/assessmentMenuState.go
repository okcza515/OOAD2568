package handler

import (
	"ModEd/eval/controller"
	"fmt"
)

// AssessmentCLIParams contains controllers needed for Assessment CLI
type AssessmentCLIParams struct {
	AssessmentController controller.AssessmentController
	SubmissionController controller.SubmissionController
	ResultController     controller.ResultController
}

// AssessmentMenuState represents the main menu of the Assessment CLI
type AssessmentMenuState struct {
	params *AssessmentCLIParams
}

// NewMainMenuState creates a new main menu state
func NewMainMenuState(params *AssessmentCLIParams) *AssessmentMenuState {
	return &AssessmentMenuState{
		params: params,
	}
}

// Enter displays the main menu
func (s *AssessmentMenuState) Enter() error {
	fmt.Println("\n===== Assessment Management =====")
	fmt.Println("1. List Assessments")
	fmt.Println("2. Create Assessment")
	fmt.Println("3. Update Assessment")
	fmt.Println("4. Delete Assessment")
	fmt.Println("5. Manage Submissions")
	fmt.Println("6. Manage Results")
	fmt.Println("back - Return to previous menu")
	return nil
}

// Exit handles exit from the main menu
func (s *AssessmentMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the main menu
func (s *AssessmentMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Assessments =====")
		fmt.Println("Assessment listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Create Assessment =====")
		fmt.Println("Assessment creation will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Assessment =====")
		fmt.Println("Assessment update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Assessment =====")
		fmt.Println("Assessment deletion will be implemented here")
		return s, nil
	case "5":
		return NewSubmissionMenuState(s.params), nil
	case "6":
		return NewResultMenuState(s.params), nil
	case "back":
		return nil, nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
