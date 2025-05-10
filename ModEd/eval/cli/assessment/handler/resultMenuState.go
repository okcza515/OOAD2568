package handler

import (
	"fmt"
)

// ResultMenuState represents the result management menu
type ResultMenuState struct {
	params *AssessmentCLIParams
}

// NewResultMenuState creates a new result menu state
func NewResultMenuState(params *AssessmentCLIParams) *ResultMenuState {
	return &ResultMenuState{
		params: params,
	}
}

// Enter displays the result menu
func (s *ResultMenuState) Enter() error {
	fmt.Println("\n===== Result Management =====")
	fmt.Println("1. List Results")
	fmt.Println("2. Record Result")
	fmt.Println("3. Update Result")
	fmt.Println("4. Delete Result")
	fmt.Println("back - Return to Assessment menu")
	return nil
}

// Exit handles exit from the result menu
func (s *ResultMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the result menu
func (s *ResultMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Results =====")
		fmt.Println("Result listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Record Result =====")
		fmt.Println("Result recording will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Result =====")
		fmt.Println("Result update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Result =====")
		fmt.Println("Result deletion will be implemented here")
		return s, nil
	case "back":
		return NewMainMenuState(s.params), nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
