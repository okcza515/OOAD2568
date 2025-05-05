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

type InstrumentLogMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewInstrumentLogMenuState(
	manager *cli.CLIMenuStateManager,
) *InstrumentLogMenuState {
	controllerInstance := controller.GetAssetInstance().InstrumentLog
	handlerContext := handler.NewHandlerContext()

	listHandler := handler.NewListHandlerStrategy[model.InstrumentLog](controllerInstance, "Instrument")
	retrieveByIDHandler := handler.NewRetrieveByIDHandlerStrategy[model.InstrumentLog](controllerInstance, "Instrument")
	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ASSET)))

	handlerContext.AddHandler("1", "List all Instrument Log", listHandler)
	handlerContext.AddHandler("2", "Get detail of an Instrument Log", retrieveByIDHandler)
	handlerContext.AddHandler("back", "Back to main menu", backHandler)

	return &InstrumentLogMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *InstrumentLogMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/instrument-log")
	fmt.Println()
	fmt.Println("Instrument Log Management")
	fmt.Println("Your options are...")
	fmt.Println()
	fmt.Println("   1:\tList all Instrument Log")
	fmt.Println("   2:\tGet detail of an Instrument Log")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *InstrumentLogMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println(err)
	}

	if input != "back" {
		util.PressEnterToContinue()
	}

	return nil
}
