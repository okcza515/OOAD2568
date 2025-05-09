package handler

import (
	"ModEd/eval/controller"
	"fmt"
)

// EvaluationCLIParams contains controllers needed for Evaluation CLI
type EvaluationCLIParams struct {
	EvaluationController *controller.EvaluationController
}

// MainMenuState represents the main menu of the Evaluation CLI
type MainMenuState struct {
	params *EvaluationCLIParams
}

// NewMainMenuState creates a new main menu state
func NewMainMenuState(params *EvaluationCLIParams) *MainMenuState {
	return &MainMenuState{
		params: params,
	}
}

// Enter displays the main menu
func (s *MainMenuState) Enter() error {
	fmt.Println("\n===== Evaluation Management =====")
	fmt.Println("1. List Evaluations")
	fmt.Println("2. Create Evaluation")
	fmt.Println("3. Update Evaluation")
	fmt.Println("4. Delete Evaluation")
	fmt.Println("back - Return to previous menu")
	return nil
}

// Exit handles exit from the main menu
func (s *MainMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the main menu
func (s *MainMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Evaluations =====")
		fmt.Println("Evaluation listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Create Evaluation =====")
		fmt.Println("Evaluation creation will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Evaluation =====")
		fmt.Println("Evaluation update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Evaluation =====")
		fmt.Println("Evaluation deletion will be implemented here")
		return s, nil
	case "back":
		return nil, nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
