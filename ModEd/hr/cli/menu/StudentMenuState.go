package menu

import (
	"ModEd/core/cli"
	coreHandler "ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/model"
	"fmt"
)

type StudentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *coreHandler.HandlerContext
}

// HandleUserInput implements cli.MenuState.
func (a *StudentMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Render implements cli.MenuState.
func (a *StudentMenuState) Render() {
	fmt.Println("=== Student Menu ===")
	a.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("exit !")
}

func NewStudentMenuState(manager *cli.CLIMenuStateManager) *StudentMenuState {
	handlerContext := coreHandler.NewHandlerContext()

	addStudentHandler := hrHandler.NewAddStudentStrategy[model.StudentInfo](nil)
	// listStudentHandler := handler.NewListHandlerStrategy[model.StudentInfo](nil)

	handlerContext.AddHandler("1", "Add new student", addStudentHandler)
	handlerContext.AddHandler("2", "List student", nil)

	return &StudentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
