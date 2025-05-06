package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type HRMainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewHRMainMenuState(manager *cli.CLIMenuStateManager) *HRMainMenuState {
	handlerContext := handler.NewHandlerContext()
	state := &HRMainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

    manager.AddMenu(string(MENU_ADD), NewAddMenuState(manager))
	return state
}

func (state *HRMainMenuState) Render() {
	fmt.Println("=== HR Main Menu ===")
	fmt.Println("1. Add Personnel")
	fmt.Println("2. List Personnel")
	// implement the remaining menu options
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}
