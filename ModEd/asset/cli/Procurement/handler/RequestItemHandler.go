package handler

import (
	"ModEd/asset/controller/Procurement"
	"ModEd/asset/util"
	"fmt"
)

func RequestItemHandler(facade *Procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Not implemented yet...")
		case "2":
			fmt.Println("Not implemented yet...")
		case "3":
			fmt.Println("Not implemented yet...")
		case "4":
			fmt.Println("Not implemented yet...")
		case "5":
			fmt.Println("Not implemented yet...")
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func printOption() {
	fmt.Println(":/asset/instrument")
	fmt.Println()
	fmt.Println("Instrument Management")
	fmt.Println("Your options are...")
	fmt.Println()
	fmt.Println("  1:\tAdd new Instrument")
	fmt.Println("  2:\tList all Instrument")
	fmt.Println("  3:\tGet detail of an Instrument")
	fmt.Println("  4:\tUpdate an Instrument")
	fmt.Println("  5:\tDelete an Instrument")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
