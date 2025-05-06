package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/hr/model"
	"fmt"
)

type StudentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

// HandleUserInput implements cli.MenuState.
func (a *StudentMenuState) HandleUserInput(input string) error {
	panic("unimplemented")
}

// Render implements cli.MenuState.
func (a *StudentMenuState) Render() {
	fmt.Println("=== Student Menu ===")
	fmt.Println()
	a.handlerContext.ShowMenu()
	// implement the remaining menu options
	fmt.Println("exit !")
}

func NewStudentMenuState(manager *cli.CLIMenuStateManager) *StudentMenuState {
	handlerContext := handler.NewHandlerContext()

	addStudentHandler := handler.NewInsertHandlerStrategy[model.StudentInfo](nil)
	listStudentHandler := handler.NewListHandlerStrategy[model.StudentInfo](nil)

	handlerContext.AddHandler("1", "Add new student", addStudentHandler)
	handlerContext.AddHandler("2", "List student", listStudentHandler)

	return &StudentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
