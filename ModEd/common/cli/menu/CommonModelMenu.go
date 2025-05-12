package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type CommonModelMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewCommonModelMenuState(manager *cli.CLIMenuStateManager) *CommonModelMenuState {
	handlerContext := handler.NewHandlerContext()
	CommonModelMenu := &CommonModelMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_COMMON)))

	handlerContext.AddHandler("1", "Student", readFileHandler)
	handlerContext.AddHandler("2", "Instructor", registerHandler)
	handlerContext.AddHandler("3", "Department", retrieveHandler)
	handlerContext.AddHandler("4", "Faculty", deleteHandler)
	handlerContext.AddHandler("5", "Back", backHandler)

	return CommonModelMenu
}

func (menu *CommonModelMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/instrument")
	fmt.Println()
	fmt.Println("Instrument Management")
	fmt.Println("Your options are...")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *CommonModelMenuState) HandleUserInput(input string) error {
	validInputs := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
	}

	if !validInputs[input] {
		return fmt.Errorf("invalid input: '%s' — please choose between 1 and 5", input)
	}

	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Handler error:", err)
		return err
	}

	return nil
}