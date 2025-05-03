package menu

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"fmt"

	"gorm.io/gorm"
)

type SpaceManagementState struct {
	manager *cli.CLIMenuStateManager

	//Add more menu here
	bookingMenu *BookingMenuState
	roomMenu    *RoomMenuState
	//instrumentManagementMenu *InstrumentManagementMenuState
	//supplyManagementMenu *SupplyManagementMenuState
}

func NewSpaceManagementState(db *gorm.DB, manager *cli.CLIMenuStateManager) *SpaceManagementState {
	spaceManagementMenu := &SpaceManagementState{
		manager: manager,
	}

	spaceManagementMenu.roomMenu = NewRoomMenuState(db, manager, spaceManagementMenu)
	spaceManagementMenu.bookingMenu = NewBookingMenuState(db, manager, spaceManagementMenu)
	//spaceManagementMenu.instrumentManagementMenu = NewInstrumentManagementMenuState(db, manager, spaceManagementMenu)
	//spaceManagementMenu.supplyManagementMenu = NewSupplyManagementMenuState(db, manager, spaceManagementMenu)

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
		fmt.Println("Not implemented yet...")
	case "2":
		fmt.Println("Not implemented yet...")
	case "3":
		menu.manager.SetState(menu.bookingMenu)
	case "4":
		fmt.Println("Not implemented yet...")
	case "5":
		menu.manager.SetState(menu.roomMenu)
	case "exit":
		util.PrintByeBye()
	default:
		fmt.Println("Invalid input, please try again.")
	}
	return nil
}
