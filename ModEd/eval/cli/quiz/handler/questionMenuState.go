package handler

import (
	"fmt"
)

// QuestionMenuState represents the question management menu for quizzes
type QuestionMenuState struct {
	params *QuizCLIParams
}

// NewQuestionMenuState creates a new question menu state
func NewQuestionMenuState(params *QuizCLIParams) *QuestionMenuState {
	return &QuestionMenuState{
		params: params,
	}
}

// Enter displays the question menu
func (s *QuestionMenuState) Enter() error {
	fmt.Println("\n===== Quiz Questions Management =====")
	fmt.Println("1. List Questions")
	fmt.Println("2. Add Question")
	fmt.Println("3. Update Question")
	fmt.Println("4. Delete Question")
	fmt.Println("back - Return to Quiz menu")
	return nil
}

// Exit handles exit from the question menu
func (s *QuestionMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the question menu
func (s *QuestionMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Quiz Questions =====")
		fmt.Println("Question listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Add Question =====")
		fmt.Println("Question addition will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Question =====")
		fmt.Println("Question update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Question =====")
		fmt.Println("Question deletion will be implemented here")
		return s, nil
	case "back":
		return NewMainMenuState(s.params), nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}
