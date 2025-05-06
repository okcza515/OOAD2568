package menu

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type SupplyMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewSupplyMenuState(
	manager *cli.CLIMenuStateManager,
) *SupplyMenuState {
	controllerInstance := controller.GetAssetInstance().Supply
	handlerContext := handler.NewHandlerContext()

	insertHandler := handler.NewInsertHandlerStrategy(controllerInstance)
	listHandler := handler.NewListHandlerStrategy(controllerInstance)
	updateHandler := handler.NewUpdateHandlerStrategy(controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy(controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ASSET)))

	handlerContext.AddHandler("1", "Add New Supply", insertHandler)
	handlerContext.AddHandler("2", "List all Supply", listHandler)
	handlerContext.AddHandler("3", "Get detail of an Supply", insertHandler)
	handlerContext.AddHandler("4", "Update an Supply", updateHandler)
	handlerContext.AddHandler("5", "Delete an Supply", deleteHandler)
	handlerContext.AddHandler("back", "Go brrr", backHandler)

	return &SupplyMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *SupplyMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/Supply")
	fmt.Println()
	fmt.Println("Supply Management")
	fmt.Println("Your options are...")
	fmt.Println()
	fmt.Println("  1:\tAdd new Supply")
	fmt.Println("  2:\tList all Supply")
	fmt.Println("  3:\tGet detail of an Supply")
	fmt.Println("  4:\tUpdate an Supply")
	fmt.Println("  5:\tDelete an Supply")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *SupplyMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
