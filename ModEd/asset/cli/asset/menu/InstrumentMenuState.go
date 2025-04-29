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

type InstrumentMenuState struct {
	manager *cli.CLIMenuStateManager

	assetMainMenu *AssetMenuState

	insertHandlerStrategy *handler.InsertHandlerStrategy[model.Instrument]
	listHandlerStrategy   *handler.ListHandlerStrategy[model.Instrument]
	updateHandlerStrategy *handler.UpdateHandlerStrategy[model.Instrument]
}

func NewInstrumentMenuState(
	manager *cli.CLIMenuStateManager, assetMainMenu *AssetMenuState,
) *InstrumentMenuState {

	controllerInstance := controller.GetAssetInstance().Instrument

	return &InstrumentMenuState{
		manager:               manager,
		assetMainMenu:         assetMainMenu,
		insertHandlerStrategy: handler.NewInsertHandlerStrategy[model.Instrument](controllerInstance),
		listHandlerStrategy:   handler.NewListHandlerStrategy[model.Instrument](controllerInstance),
		updateHandlerStrategy: handler.NewUpdateHandlerStrategy[model.Instrument](controllerInstance),
	}
}

func (menu *InstrumentMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/instrument")
	fmt.Println()
	fmt.Println("Instrument Management")
	fmt.Println("Your options are...")
	fmt.Println()
	fmt.Println("  1:\tAdd new Instrument")
	fmt.Println("  2:\tList all Instrument")
	fmt.Println("  3:\tGet detail of an Instrument")
	fmt.Println("  4:\tUpdate an Instrument")
	fmt.Println("  5:\tDelete an Instrument")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *InstrumentMenuState) HandleUserInput(input string) error {
	context := &handler.HandlerContext{}

	switch input {
	case "1":
		fmt.Println("Add New Instrument")
		context.SetStrategy(menu.insertHandlerStrategy)
	case "2":
		fmt.Println("List all Instrument")
		context.SetStrategy(menu.listHandlerStrategy)
	case "3":
		fmt.Println("Get detail of an Instrument")
	case "4":
		fmt.Println("Update an Instrument")
		context.SetStrategy(menu.updateHandlerStrategy)
	case "5":
		fmt.Println("Delete an Instrument")
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
