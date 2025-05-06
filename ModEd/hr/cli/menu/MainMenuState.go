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

// HandleUserInput implements cli.MenuState.
func (state *HRMainMenuState) HandleUserInput(input string) error {
	err := state.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (state *HRMainMenuState) Render() {
	fmt.Println("=== HR Menu ===")
    fmt.Println()
	state.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("exit !")
	fmt.Println("Enter your choice: ")
}

func NewHRMainMenuState(manager *cli.CLIMenuStateManager) *HRMainMenuState {
	handlerContext := handler.NewHandlerContext()
	state := &HRMainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	manager.AddMenu(string(MENU_HR), state)
	manager.AddMenu(string(MENU_ADD), NewAddMenuState(manager))

    addHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ADD)))

    handlerContext.AddHandler("1", "Add New Student/Instructor", addHandler)
	return state
}
