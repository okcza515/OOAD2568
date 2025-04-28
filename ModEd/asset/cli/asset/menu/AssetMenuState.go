package menu

// MEP-1012 Asset

import (
	"ModEd/core/cli"
	"fmt"
)

type AssetMenuState struct {
	manager *cli.CLIMenuStateManager

	// Add more menu here
	instrumentMenu *InstrumentMenuState
	supplyMenu     *SupplyMenuState
}

func NewAssetMenuState(manager *cli.CLIMenuStateManager) *AssetMenuState {
	assetMenu := &AssetMenuState{
		manager: manager,
	}

	// Add more menu here
	assetMenu.instrumentMenu = NewInstrumentMenuState(manager, assetMenu)
	assetMenu.supplyMenu = NewSupplyMenuState(manager, assetMenu)

	return assetMenu
}

func (menu *AssetMenuState) Render() {
	fmt.Println()
	fmt.Println(":/asset/instrument")
	fmt.Println()
	fmt.Println("Welcome to ModEd Asset Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tCategory Page")
	fmt.Println("  2:\tInstrument Page")
	fmt.Println("  3:\tSupply Page")
	fmt.Println("  4:\tBorrow Page")
	fmt.Println("  5:\tInstrument Log Page")
	fmt.Println("  6:\tSupply Log Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssetMenuState) HandleUserInput(input string) error {
	switch input {
	case "1":
		fmt.Println("Not implemented yet...")
	case "2":
		menu.manager.SetState(menu.instrumentMenu)
	case "3":
		menu.manager.SetState(menu.supplyMenu)
	case "4":
		fmt.Println("Not implemented yet...")
	case "5":
		fmt.Println("Not implemented yet...")
	case "6":
		fmt.Println("Not implemented yet...")
	default:
		fmt.Println("invalid input")
	}

	return nil
}
