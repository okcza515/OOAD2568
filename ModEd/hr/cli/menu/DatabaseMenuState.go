package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	hrHandler "ModEd/hr/cli/menu/handler"
	"ModEd/hr/controller"
	"fmt"
)

// var (
// 	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
// )

type DatabaseMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func (a *DatabaseMenuState) HandleUserInput(input string) error {
	err := a.handlerContext.HandleInput(input)
	if err != nil {
		return err
	}

	return nil
}

func (a *DatabaseMenuState) Render() {
	fmt.Println("=== Database Menu ===")
	a.handlerContext.ShowMenu()
	fmt.Println("exit:\tExit the program.")
}

func NewDatabaseMenuState(manager *cli.CLIMenuStateManager, studentCtrl *controller.StudentHRController, instructorCtrl *controller.InstructorHRController) *DatabaseMenuState {

	handlerContext := handler.NewHandlerContext()

	pullStudentHandler := hrHandler.NewPullStudentHandlerStrategy(studentCtrl)
	pullInstructorHandler := hrHandler.NewPullInstructorHandlerStrategy(instructorCtrl)

	handlerContext.AddHandler("1", "Pull student data", pullStudentHandler)
	handlerContext.AddHandler("2", "Pull instructor data", pullInstructorHandler)

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_HR)))
	handlerContext.AddHandler("0", "Back to main menu", backHandler)

	return &DatabaseMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}
