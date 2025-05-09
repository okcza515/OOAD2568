package handler

import (
	"fmt"
)

// MenuState interface for managing menu states
type MenuState interface {
	Enter() error
	Exit() error
	HandleInput(input string) (MenuState, error)
}

// MenuStateManager manages menu state transitions
type MenuStateManager struct {
	currentState MenuState
}

// NewMenuStateManager creates a new menu state manager
func NewMenuStateManager(initialState MenuState) *MenuStateManager {
	return &MenuStateManager{
		currentState: initialState,
	}
}

// Run executes the menu state manager loop
func (m *MenuStateManager) Run() error {
	if err := m.currentState.Enter(); err != nil {
		return err
	}

	for {
		var input string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		nextState, err := m.currentState.HandleInput(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if nextState == nil {
			return m.currentState.Exit()
		}

		if err := m.currentState.Exit(); err != nil {
			return err
		}

		m.currentState = nextState

		if err := m.currentState.Enter(); err != nil {
			return err
		}
	}
}

// BaseMenuState provides common functionality for menu states
type BaseMenuState struct {
	title    string
	options  map[string]string
	handlers map[string]func() (MenuState, error)
}

// NewBaseMenuState creates a new base menu state
func NewBaseMenuState(title string) *BaseMenuState {
	return &BaseMenuState{
		title:    title,
		options:  make(map[string]string),
		handlers: make(map[string]func() (MenuState, error)),
	}
}

// AddOption adds a menu option
func (s *BaseMenuState) AddOption(key, description string, handler func() (MenuState, error)) {
	s.options[key] = description
	s.handlers[key] = handler
}

// Enter displays the menu
func (s *BaseMenuState) Enter() error {
	fmt.Printf("\n===== %s =====\n", s.title)
	for key, description := range s.options {
		fmt.Printf("%s. %s\n", key, description)
	}
	fmt.Println("back - Return to previous menu")
	return nil
}

// Exit handles exit from the menu
func (s *BaseMenuState) Exit() error {
	return nil
}

// HandleInput processes user input
func (s *BaseMenuState) HandleInput(input string) (MenuState, error) {
	if input == "back" {
		return nil, nil
	}

	if handler, exists := s.handlers[input]; exists {
		return handler()
	}

	return s, fmt.Errorf("invalid choice: %s", input)
}
