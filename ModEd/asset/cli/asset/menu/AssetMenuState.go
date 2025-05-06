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
	manager.AddMenu(string(MENU_INSTRUMENT), NewInstrumentMenuState(manager))
	manager.AddMenu(string(MENU_SUPPLY), NewSupplyMenuState(manager))
	manager.AddMenu(string(MENU_CATEGORY), NewCategoryMenuState(manager))
	manager.AddMenu(string(MENU_INSTRUMENT_LOG), NewInstrumentLogMenuState(manager))
	manager.AddMenu(string(MENU_SUPPLY_LOG), NewSupplyLogMenuState(manager))

	categoryHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_CATEGORY)))
	instrumentHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUMENT)))
	supplyHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_SUPPLY)))
	instrumentLogHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUMENT_LOG)))
	supplyLogHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_SUPPLY_LOG)))

	handlerContext.AddHandler("1", "", categoryHandler)
	handlerContext.AddHandler("2", "", instrumentHandler)
	handlerContext.AddHandler("3", "", supplyHandler)
	handlerContext.AddHandler("4", "", nil)
	handlerContext.AddHandler("5", "", instrumentLogHandler)
	handlerContext.AddHandler("6", "", supplyLogHandler)

	return assetMenu
}

func (menu *AssetMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset")
	fmt.Println()
	fmt.Println("Welcome to ModEd Asset Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println()
	fmt.Println("  1:\tCategory Page")
	fmt.Println("  2:\tInstrument Page")
	fmt.Println("  3:\tSupply Page")
	fmt.Println("  4:\tBorrow Page")
	fmt.Println("  5:\tInstrument Log Page")
	fmt.Println("  6:\tSupply Log Page")
	//  fmt.Println("  7:\tDetail Report")
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
