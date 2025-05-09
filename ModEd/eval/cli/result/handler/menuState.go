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
