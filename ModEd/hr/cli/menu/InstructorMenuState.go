package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type InstructorMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *coreHandler.HandlerContext
}

// HandleUserInput implements cli.MenuState.
func (a *InstructorMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

// Render implements cli.MenuState.
func (a *InstructorMenuState) Render() {
	fmt.Println("=== Instructor Menu ===")
	a.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("exit:\tExit the program.")
}

func NewInstructorMenuState(manager *cli.CLIMenuStateManager, instructorCtrl *controller.InstructorHRController) *InstructorMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	addInstructorHandler := hrHandler.NewAddInstructorStrategy(instructorCtrl)
	listInstructorHandler := hrHandler.NewListInstructorStrategy(instructorCtrl)
	updateInstructorHandler := hrHandler.NewUpdateInstructorInfoStrategy(instructorCtrl)
	// requestInstructorHandler := hrHandler.NewRequestInstructorLeaveStrategy(instructorCtrl)

	handlerContext.AddHandler("1", "Add new instructor", addInstructorHandler)
	handlerContext.AddHandler("2", "List instructor", listInstructorHandler)
	handlerContext.AddHandler("3", "Update instructor Info", updateInstructorHandler)
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
