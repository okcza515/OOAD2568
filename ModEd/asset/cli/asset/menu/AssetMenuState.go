package menu

// MEP-1012 Asset

import (
	"ModEd/core/cli"
	"fmt"
)

type AssetMenuState struct {
	manager *cli.CLIMenuStateManager
}

func NewAssetMenuState(manager *cli.CLIMenuStateManager) *AssetMenuState {
	assetMenu := &AssetMenuState{
		manager: manager,
	}

	// Add more menu here
	manager.AddMenu(string(MENU_ASSET), assetMenu)
	manager.AddMenu(string(MENU_INSTRUMENT), NewInstrumentMenuState(manager))
	manager.AddMenu(string(MENU_SUPPLY), NewSupplyMenuState(manager))
	manager.AddMenu(string(MENU_INSTRUMENT_LOG), NewInstrumentLogMenuState(manager))

	return assetMenu
}

func (menu *AssetMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset")
	fmt.Println()
	fmt.Println("Welcome to ModEd Asset Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tCategory Page")
	fmt.Println("  2:\tInstrument Page")
	fmt.Println("  3:\tSupply Page")
	fmt.Println("  4:\tBorrow Page")
	fmt.Println("  5:\tInstrument Log Page")
	fmt.Println("  6:\tSupply Log Page")
	fmt.Println("  7:\tDetail Report")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssetMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		fmt.Println("Not implemented yet...")
	case "2":
		menu.manager.GoToMenu(string(MENU_INSTRUMENT))
	case "3":
		menu.manager.GoToMenu(string(MENU_SUPPLY))
	case "4":
		fmt.Println("Not implemented yet...")
	case "5":
		menu.manager.GoToMenu(string(MENU_INSTRUMENT_LOG))
	case "6":
		fmt.Println("Not implemented yet...")
	default:
		fmt.Println("invalid input")
	}

	return nil
}
