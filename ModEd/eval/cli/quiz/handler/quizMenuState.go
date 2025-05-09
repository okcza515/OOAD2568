package handler

import (
	"ModEd/eval/controller"
	"fmt"
)

// QuizCLIParams contains controllers needed for Quiz CLI
type QuizCLIParams struct {
	QuizController     controller.QuizController
	QuestionController *controller.QuestionController
}

// QuizMenuState represents the main menu of the Quiz CLI
type QuizMenuState struct {
	params *QuizCLIParams
}

// NewMainMenuState creates a new main menu state
func NewMainMenuState(params *QuizCLIParams) *QuizMenuState {
	return &QuizMenuState{
		params: params,
	}
}

// Enter displays the main menu
func (s *QuizMenuState) Enter() error {
	fmt.Println("\n===== Quiz Management =====")
	fmt.Println("1. List Quizzes")
	fmt.Println("2. Create Quiz")
	fmt.Println("3. Update Quiz")
	fmt.Println("4. Delete Quiz")
	fmt.Println("5. Manage Quiz Questions")
	fmt.Println("back - Return to previous menu")
	return nil
}

// Exit handles exit from the main menu
func (s *QuizMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the main menu
func (s *QuizMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Quizzes =====")
		fmt.Println("Quiz listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Create Quiz =====")
		fmt.Println("Quiz creation will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Quiz =====")
		fmt.Println("Quiz update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Quiz =====")
		fmt.Println("Quiz deletion will be implemented here")
		return s, nil
	case "5":
		return NewQuestionMenuState(s.params), nil
	case "back":
		return nil, nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
