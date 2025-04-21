package handler

import (
	procurement "ModEd/asset/controller/Procurement"
	"ModEd/asset/util"
	"fmt"
)

func RequestItemHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":

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
	fmt.Println(":/Procurement/RequestItem")
	fmt.Println()
	fmt.Println("--RequestItem Function--")
	fmt.Println("  1:\tUpdate")
	fmt.Println("  2:\t<null>")
	fmt.Println("  3:\t<null>")
	fmt.Println("  4:\t<null>")
	fmt.Println("  5:\t<null>")
	fmt.Println("  6:\t<null>")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
