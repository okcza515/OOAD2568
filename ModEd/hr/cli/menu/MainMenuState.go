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

// HandleUserInput implements cli.MenuState.
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

func NewHRMainMenuState(manager *cli.CLIMenuStateManager, studentCtrl *controller.StudentHRController) *HRMainMenuState {
	handlerContext := handler.NewHandlerContext()
	state := &HRMainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	manager.AddMenu(string(MENU_HR), state)
	manager.AddMenu(string(MENU_STUDENT), NewStudentMenuState(manager, studentCtrl))
	manager.AddMenu(string(MENU_INSTRUCTOR), NewInstructorMenuState(manager))

	studentHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_STUDENT)))
	instructorHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUCTOR)))

	handlerContext.AddHandler("1", "Student Menu", studentHandler)
	handlerContext.AddHandler("2", "Instructor Menu", instructorHandler)
	return state
}
