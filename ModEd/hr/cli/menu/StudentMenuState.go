package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

type StudentMenuState struct {
	manager           *cli.CLIMenuStateManager
	handlerContext    *coreHandler.HandlerContext
	studentController *controller.StudentHRController
}

// HandleUserInput implements cli.MenuState.
func (a *StudentMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

// Render implements cli.MenuState.
func (a *StudentMenuState) Render() {
	fmt.Println("=== Student Menu ===")
	a.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("back:\tBack to main menu")
}

func NewStudentMenuState(manager *cli.CLIMenuStateManager, studentCtrl *controller.StudentHRController) *StudentMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	// Pass the controller to your strategy/handler
	addStudentHandler := hrHandler.NewAddStudentStrategy(studentCtrl)

	handlerContext.AddHandler("1", "Add new student", addStudentHandler)
	handlerContext.AddHandler("2", "List student", nil) //

	backHandler := coreHandler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_HR)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &StudentMenuState{
		manager:           manager,
		handlerContext:    handlerContext,
		studentController: studentCtrl,
	}
}
