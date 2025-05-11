package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	coreHandler "ModEd/core/handler"
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

	handlerContext.AddHandler("1", "Add new instructor", nil)
	handlerContext.AddHandler("2", "List instructor", nil)
	handlerContext.AddHandler("3", "Update instructor Info", nil)
	handlerContext.AddHandler("4", "Delete instructor", nil)
	handlerContext.AddHandler("5", "Request leave", nil)
	handlerContext.AddHandler("6", "Request resignation", nil)
	handlerContext.AddHandler("7", "Request raise", nil)
	handlerContext.AddHandler("8", "Review leave", nil)
	handlerContext.AddHandler("9", "Review resignation", nil)
	handlerContext.AddHandler("10", "Review raise", nil)

	backHandler := coreHandler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_HR)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &InstructorMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
