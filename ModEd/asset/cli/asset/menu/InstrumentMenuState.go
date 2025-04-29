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
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewInstrumentMenuState(
	manager *cli.CLIMenuStateManager, assetMainMenu *AssetMenuState,
) *InstrumentMenuState {
	controllerInstance := controller.GetAssetInstance().Instrument

	handlerContext := handler.NewHandlerContext()

	insertHandler := handler.NewInsertHandlerStrategy[model.Instrument](controllerInstance)
	listHandler := handler.NewListHandlerStrategy[model.Instrument](controllerInstance)
	updateHandler := handler.NewUpdateHandlerStrategy[model.Instrument](controllerInstance)
	deleteHandler := handler.NewDeleteHandlerStrategy[model.Instrument](controllerInstance)
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, assetMainMenu)

	handlerContext.AddHandler("1", "Add New Instrument", insertHandler)
	handlerContext.AddHandler("2", "List all Instrument", listHandler)
	handlerContext.AddHandler("3", "Get full detail of an Instrument", nil)
	handlerContext.AddHandler("4", "Update an Instrument", updateHandler)
	handlerContext.AddHandler("5", "Delete an Instrument", deleteHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &InstrumentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
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
	fmt.Println("  3:\tGet full detail of an Instrument")
	fmt.Println("  4:\tUpdate an Instrument")
	fmt.Println("  5:\tDelete an Instrument")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *InstrumentMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
