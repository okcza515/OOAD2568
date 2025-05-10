package menu

// MEP-1012 Asset

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type AssetMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAssetMenuState(manager *cli.CLIMenuStateManager) *AssetMenuState {
	handlerContext := handler.NewHandlerContext()
	assetMenu := &AssetMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	// Add more menu here
	manager.AddMenu(string(MENU_ASSET), assetMenu)
	manager.AddMenu(string(MENU_CATEGORY), NewCategoryMenuState(manager))
	manager.AddMenu(string(MENU_INSTRUMENT), NewInstrumentMenuState(manager))
	manager.AddMenu(string(MENU_SUPPLY), NewSupplyMenuState(manager))
	manager.AddMenu(string(MENU_BORROW), NewBorrowInstrumentMenuState(manager))
	manager.AddMenu(string(MENU_INSTRUMENT_LOG), NewInstrumentLogMenuState(manager))
	manager.AddMenu(string(MENU_SUPPLY_LOG), NewSupplyLogMenuState(manager))

	categoryHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_CATEGORY)))
	instrumentHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUMENT)))
	supplyHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_SUPPLY)))
	borrowHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_BORROW)))
	instrumentLogHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUMENT_LOG)))
	supplyLogHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_SUPPLY_LOG)))

	handlerContext.AddHandler("1", "Category Page", categoryHandler)
	handlerContext.AddHandler("2", "Instrument Page", instrumentHandler)
	handlerContext.AddHandler("3", "Supply Page", supplyHandler)
	handlerContext.AddHandler("4", "Borrow Page", borrowHandler)
	handlerContext.AddHandler("5", "Instrument Log Page", instrumentLogHandler)
	handlerContext.AddHandler("6", "Supply Log Page", supplyLogHandler)

	return assetMenu
}

func (menu *AssetMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset")
	fmt.Println()
	fmt.Println("Welcome to ModEd Asset Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssetMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("err: Invalid input, menu '" + input + "' doesn't exist")
		util.PressEnterToContinue()
	}

	return nil
}
