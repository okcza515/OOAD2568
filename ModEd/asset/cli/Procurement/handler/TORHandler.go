package handler

import (
	procurement "ModEd/asset/controller"
	util "ModEd/asset/util"
	"fmt"
)

func TORHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printTOROptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create Procurement")
			WaitForEnter()

		case "2":
			fmt.Println("List All Procurements")
			WaitForEnter()
		case "3":
			fmt.Println("View Procurement by ID")
			WaitForEnter()
		case "4":
			fmt.Println("Update Procurement Status")
			WaitForEnter()
		}

		util.ClearScreen()
	}

	util.ClearScreen()

}

func printTOROptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--TOR Functions--")
	fmt.Println("  1:\tCreate TOR")
	fmt.Println("  2:\tList All TORs")
	fmt.Println("  3:\tView TOR by ID")
	fmt.Println("  4:\tDelete TOR")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit")
	fmt.Println()
}
