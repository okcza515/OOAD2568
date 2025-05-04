package handler

import (
	"ModEd/curriculum/utils"
	"errors"
	"fmt"
)

var (
	ExitCommand = errors.New("exit")
	BackCommand = errors.New("back")
)

// MenuState defines a state in the menu system
type MenuState interface {
	// Display shows the menu
	Display()

	// HandleInput processes user input and returns the next state or error
	HandleInput(input string) (MenuState, error)

	// GetName returns the name of this menu state
	GetName() string
}

// MenuStateManager manages transitions between menu states
type MenuStateManager struct {
	CurrentState MenuState
}

// NewMenuStateManager creates a new menu state manager
func NewMenuStateManager(initialState MenuState) *MenuStateManager {
	return &MenuStateManager{
		CurrentState: initialState,
	}
}

// Run runs the menu state manager
func (m *MenuStateManager) Run() error {
	for {
		m.CurrentState.Display()

		input := utils.GetUserChoice()

		nextState, err := m.CurrentState.HandleInput(input)
		if err != nil {
			if errors.Is(err, ExitCommand) {
				return nil
			}
			if errors.Is(err, BackCommand) {
				return err
			}
			fmt.Println("Error:", err)
			continue
		}

		if nextState != nil {
			m.CurrentState = nextState
		}
	}
}
