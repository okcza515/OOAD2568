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
	manager                     *cli.CLIMenuStateManager
	assetMainMenu               *AssetMenuState
	insertHandlerStrategy       *handler.InsertHandlerStrategy[model.InstrumentLog]
	listHandlerStrategy         *handler.ListHandlerStrategy[model.InstrumentLog]
	retrieveByIDHandlerStrategy *handler.RetrieveByIDHandlerStrategy[model.InstrumentLog]
}

func NewInstrumentLogMenuState(
	manager *cli.CLIMenuStateManager, assetMainMenu *AssetMenuState) *InstrumentLogMenuState {
	controllerInstance := controller.GetAssetInstance().InstrumentLog
	return &InstrumentLogMenuState{
		manager:                     manager,
		assetMainMenu:               assetMainMenu,
		insertHandlerStrategy:       handler.NewInsertHandlerStrategy[model.InstrumentLog](controllerInstance),
		listHandlerStrategy:         handler.NewListHandlerStrategy[model.InstrumentLog](controllerInstance, "Instrument"),
		retrieveByIDHandlerStrategy: handler.NewRetrieveByIDHandlerStrategy[model.InstrumentLog](controllerInstance, "Instrument"),
	}
}

func (menu *InstrumentLogMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/instrumentLog")
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
	context := &handler.HandlerContext{}

	switch input {
	case "1":
		fmt.Println("List all Instrument Log")
		context.SetStrategy(menu.listHandlerStrategy)
	case "2":
		fmt.Println("Get detail of an Instrument Log")
		context.SetStrategy(menu.retrieveByIDHandlerStrategy)
	case "back":
		menu.manager.SetState(menu.assetMainMenu)
		return nil
	case "exit":
		return nil
	}

	err := context.Execute()

	if err != nil {
		fmt.Println(err)
	}

	util.PressEnterToContinue()

	return nil
}
