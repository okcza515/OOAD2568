package menu

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type SupplyMenuState struct {
	manager          *cli.CLIMenuStateManager
	controllerFacade *controller.AssetControllerFacade

	assetMainMenu *AssetMenuState

	insertHandlerStrategy *handler.InsertHandlerStrategy[model.Supply]
	listHandlerStrategy   *handler.ListHandlerStrategy[model.Supply]
	deleteHandlerStrategy *handler.DeleteHandlerStrategy[model.Supply]
}

func NewSupplyMenuState(
	manager *cli.CLIMenuStateManager, controllerFacade *controller.AssetControllerFacade, assetMainMenu *AssetMenuState,
) *SupplyMenuState {
	return &SupplyMenuState{
		manager:               manager,
		controllerFacade:      controllerFacade,
		assetMainMenu:         assetMainMenu,
		insertHandlerStrategy: handler.NewInsertHandlerStrategy(controllerFacade.Supply),
		listHandlerStrategy:   handler.NewListHandlerStrategy(controllerFacade.Supply),
		deleteHandlerStrategy: handler.NewDeleteHandlerStrategy(controllerFacade.Supply),
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
	context := &handler.HandlerContext{}

	switch input {
	case "1":
		fmt.Println("Add New Supply")
		context.SetStrategy(menu.insertHandlerStrategy)
	case "2":
		fmt.Println("List all Supply")
		context.SetStrategy(menu.listHandlerStrategy)
	case "3":
		fmt.Println("Get detail of an Supply")
	case "4":
		fmt.Println("Update an Supply")
	case "5":
		fmt.Println("Delete an Supply")
		context.SetStrategy(menu.deleteHandlerStrategy)
	case "back":
		menu.manager.SetState(menu.assetMainMenu)
		return nil
	default:
		context.SetStrategy(handler.DoNothingHandlerStrategy{})
		fmt.Println("Invalid Command")
	}

	err := context.Execute()
	if err != nil {
		fmt.Println(err)
	}

	util.PressEnterToContinue()

	return nil
}
