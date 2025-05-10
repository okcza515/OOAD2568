package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type InstructorMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

// HandleUserInput implements cli.MenuState.
func (a *InstructorMenuState) HandleUserInput(input string) error {
	panic("unimplemented")
}

// Render implements cli.MenuState.
func (a *InstructorMenuState) Render() {
	fmt.Println("=== Instructor Menu ===")
	a.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("back:\tBack to main menu")
}

func NewInstructorMenuState(manager *cli.CLIMenuStateManager) *InstructorMenuState {
	handlerContext := handler.NewHandlerContext()
	addMenu := &InstructorMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	// implement

	return addMenu
}
