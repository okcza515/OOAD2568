package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/hr/controller"
	"fmt"
)

type HRMainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (state *HRMainMenuState) HandleUserInput(input string) error {
	err := state.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error invalid input, please try again")
		return err
	}

	return nil
}

func (state *HRMainMenuState) Render() {
	fmt.Println("=== HR Menu ===")
	state.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
	fmt.Println()
}

func NewHRMainMenuState(
	cliManager *cli.CLIMenuStateManager,
	hrManager *controller.HRControllerManager,
) *HRMainMenuState {
	handlerContext := handler.NewHandlerContext()
	state := &HRMainMenuState{
		manager:        cliManager,
		handlerContext: handlerContext,
	}

	cliManager.AddMenu(string(MENU_HR), state)

	studentMenu := NewStudentMenuState(cliManager, hrManager)
	cliManager.AddMenu(string(MENU_STUDENT), studentMenu)

	instructorMenu := NewInstructorMenuState(cliManager, hrManager)
	cliManager.AddMenu(string(MENU_INSTRUCTOR), instructorMenu)

	databaseMenu := NewDatabaseMenuState(cliManager, hrManager)
	cliManager.AddMenu(string(MENU_DATABASE), databaseMenu)

	studentHandler := handler.NewChangeMenuHandlerStrategy(cliManager, studentMenu)
	instructorHandler := handler.NewChangeMenuHandlerStrategy(cliManager, instructorMenu)
	databaseHandler := handler.NewChangeMenuHandlerStrategy(cliManager, databaseMenu)

	handlerContext.AddHandler("1", "Student Menu", studentHandler)
	handlerContext.AddHandler("2", "Instructor Menu", instructorHandler)
	handlerContext.AddHandler("3", "Database Menu", databaseHandler)

	return state
}
