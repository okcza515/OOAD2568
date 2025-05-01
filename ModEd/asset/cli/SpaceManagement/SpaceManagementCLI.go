// MEP-1013
package main

import (
	"ModEd/asset/cli/spacemanagement/handler"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
)

func main() {

	input := ""
	for input != "exit" {
		util.ClearScreen()
		util.PrintSpaceManagementBanner()
		printOption()
		input = util.GetCommandInput()

		facade := controller.GetSpaceManagementInstance()
		switch input {
		case "1":
			handler.InstrumentManagementHandler(facade)
		case "2":
			handler.SupplyManagementHandler(facade)
		case "3":
			handler.BookingHandler(facade)
		case "4":
			handler.PermanentBookingHandler(facade)
		case "5":
			handler.RoomHandler(facade)
		}
	}
	if input == "exit" {
		util.PrintByeBye()
	}
}

func printOption() {
	fmt.Println("\n===============================")
	fmt.Println()
	fmt.Println("Welcome to ModEd Space Management Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tInstrument Management Page")
	fmt.Println("  2:\tSupply Management Page")
	fmt.Println("  3:\tBooking Page")
	fmt.Println("  4:\tPermanent Schedule Page")
	fmt.Println("  5:\tRoom Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
