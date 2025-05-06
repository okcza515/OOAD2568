package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
)

type AddMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

// HandleUserInput implements cli.MenuState.
func (a *AddMenuState) HandleUserInput(input string) error {
	panic("unimplemented")
}

// Render implements cli.MenuState.
func (a *AddMenuState) Render() {
	panic("unimplemented")
}

func NewAddMenuState(manager *cli.CLIMenuStateManager) *AddMenuState {
	handlerContext := handler.NewHandlerContext()
	addMenu := &AddMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	// implement

	return addMenu
}
