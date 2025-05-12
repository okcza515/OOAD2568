package menu

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"

	"fmt"
)

var MENU_COMMON = "common"
var READ_FILE = "readfile"
var REGISTER = "register"
var RETRIEVE = "retrieve"
var DELETE = "delete"
var CLEAR_DB = "cleardb"
var TEST = "test"

type CommonMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewCommonMenuState(manager *cli.CLIMenuStateManager) *CommonMenuState {
	handlerContext := handler.NewHandlerContext()
	CommonMenu := &CommonMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
	
	manager.AddMenu(string(MENU_COMMON), CommonMenu)
	manager.AddMenu(string(READ_FILE), 	ReadFileMenuState(manager))
	manager.AddMenu(string(REGISTER), 	NewCommonModelMenuState(manager))
	manager.AddMenu(string(RETRIEVE), 	NewCommonModelMenuState(manager))
	manager.AddMenu(string(DELETE), 	NewCommonModelMenuState(manager))
	manager.AddMenu(string(CLEAR_DB), 	NewCommonModelMenuState(manager))
	manager.AddMenu(string(TEST), 		NewCommonModelMenuState(manager))

	readFileHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(READ_FILE)))
	registerHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(REGISTER)))
	retrieveHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(RETRIEVE)))
	deleteHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(DELETE)))
	cleardbHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(CLEAR_DB)))
	testHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(TEST)))

	handlerContext.AddHandler("1", "Read file", readFileHandler)
	handlerContext.AddHandler("2", "Register", registerHandler)
	handlerContext.AddHandler("3", "Retrieve", retrieveHandler)
	handlerContext.AddHandler("4", "Delete", deleteHandler)
	handlerContext.AddHandler("5", "Clear Database", cleardbHandler)
	handlerContext.AddHandler("6", "test", testHandler)

	return CommonMenu
}

func (menu *CommonMenuState) Render() {
	fmt.Println()
	fmt.Println("Common CLI menu")
	fmt.Println("Choose your action")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *CommonMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("err: Invalid input, menu '" + input + "' doesn't exist")
		util.PressEnterToContinue()
	}

	return nil
}