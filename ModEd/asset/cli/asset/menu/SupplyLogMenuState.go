package menu

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type SupplyLogMenuState struct {
	manager *cli.CLIMenuStateManager

	handlerContext *handler.HandlerContext
}

func NewSupplyLogMenuState(
	manager *cli.CLIMenuStateManager,
) *SupplyLogMenuState {
	controllerInstance := controller.GetAssetInstance().SupplyLog

	handlerContext := handler.NewHandlerContext()

	listHandler := handler.NewListHandlerStrategy(controllerInstance, "Supply")
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ASSET)))

	handlerContext.AddHandler("1", "Show Supply Log", listHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &SupplyLogMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *SupplyLogMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/SupplyLog")
	fmt.Println()
	fmt.Println("Supply Management")
	fmt.Println("Your options are...")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *SupplyLogMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
