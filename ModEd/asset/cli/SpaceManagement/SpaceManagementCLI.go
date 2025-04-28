package main

import (
	handler "ModEd/asset/cli/spacemanagement/handler"
	controller "ModEd/asset/controller"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	facade, err := controller.NewSpaceManagementControllerFacade()
	if err != nil {
		panic("err: initialize controllers failed")
	}

	input := ""
	for input != "exit" {
		util.ClearScreen()
		util.PrintSpaceManagementBanner()
		printOption()
		input = util.GetCommandInput()

		switch input {
		case "1":
			handler.AssetManagementHandler(facade)
		case "2":
			handler.BookingHandler(facade)
		case "3":
			handler.PermanentBookingHandler(facade)
		case "4":
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
	fmt.Println("  1:\tAsset Management Page")
	fmt.Println("  2:\tBooking Page")
	fmt.Println("  3:\tPermanent Schedule Page")
	fmt.Println("  4:\tRoom Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
