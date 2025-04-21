package handler

import (
	procurement "ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
)

func ProcurementHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printProcurementOptions()
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
		case "5":
			fmt.Println("Delete Procurement")
			WaitForEnter()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func printProcurementOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--Procurement Functions--")
	fmt.Println("  1:\tCreate Procurement")
	fmt.Println("  2:\tList All Procurements")
	fmt.Println("  3:\tView Procurement by ID")
	fmt.Println("  4:\tUpdate Procurement Status")
	fmt.Println("  5:\tDelete Procurement")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
