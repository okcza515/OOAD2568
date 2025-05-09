package handler

import (
	"ModEd/eval/controller"
	"fmt"
)

// ProgressCLIParams contains controllers needed for Progress CLI
type ProgressCLIParams struct {
	ProgressController *controller.ProgressController
}

// MainMenuState represents the main menu of the Progress CLI
type MainMenuState struct {
	params *ProgressCLIParams
}

// NewMainMenuState creates a new main menu state
func NewMainMenuState(params *ProgressCLIParams) *MainMenuState {
	return &MainMenuState{
		params: params,
	}
}

// Enter displays the main menu
func (s *MainMenuState) Enter() error {
	fmt.Println("\n===== Progress Management =====")
	fmt.Println("1. List Progress Records")
	fmt.Println("2. Create Progress Record")
	fmt.Println("3. Update Progress Record")
	fmt.Println("4. Delete Progress Record")
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
		fmt.Println("\n===== Progress Records =====")
		fmt.Println("Progress listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Create Progress Record =====")
		fmt.Println("Progress creation will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Progress Record =====")
		fmt.Println("Progress update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Progress Record =====")
		fmt.Println("Progress deletion will be implemented here")
		return s, nil
	case "back":
		return nil, nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
