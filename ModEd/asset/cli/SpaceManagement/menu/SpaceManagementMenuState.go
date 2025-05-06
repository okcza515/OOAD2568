// MEP-1013
package menu

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"fmt"

	"gorm.io/gorm"
)

type SpaceManagementState struct {
	manager *cli.CLIMenuStateManager
}

func NewSpaceManagementState(db *gorm.DB, manager *cli.CLIMenuStateManager) *SpaceManagementState {
	spaceManagementMenu := &SpaceManagementState{
		manager: manager,
	}

	//facade
	//Add more menu here (state)
	manager.AddMenu("Instrument Management", NewInstrumentMenuState(db, manager, spaceManagementMenu))
	manager.AddMenu("Supply Management", NewSupplyMenuState(db, manager, spaceManagementMenu))
	manager.AddMenu("room", NewRoomMenuState(db, manager, spaceManagementMenu))
	manager.AddMenu("booking", NewBookingMenuState(db, manager, spaceManagementMenu))
	manager.AddMenu("Schedule", NewPermanentScheduleState(db, manager, spaceManagementMenu))

	return spaceManagementMenu
}

func (menu *SpaceManagementState) Render() {
	util.PrintSpaceManagementBanner()
	fmt.Println("\n===============================")
	fmt.Println()
	fmt.Println("Welcome to ModEd Space Management Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tInstrument Management Page")
	fmt.Println("  2:\tSupply Management Page")
	fmt.Println("  3:\tBooking Page")
	fmt.Println("  4:\tPermanent Schedule Page")
	fmt.Println("  5:\tRoom Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C")
	fmt.Println()
}

func (menu *SpaceManagementState) HandleUserInput(input string) error {
	switch input {
	case "1":
		menu.manager.GoToMenu("Instrument Management")
	case "2":
		menu.manager.GoToMenu("Supply Management")
	case "3":
		menu.manager.GoToMenu("booking")
	case "4":
		menu.manager.GoToMenu("Schedule")
	case "5":
		menu.manager.GoToMenu("room")
	case "exit":
		util.PrintByeBye()
	default:
		fmt.Println("Invalid input, please try again.")
	}
	return nil
}
