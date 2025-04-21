package handler

import (
	procurement "ModEd/asset/controller/Procurement"
	model "ModEd/asset/model/Procurement"
	"ModEd/asset/util"
	"fmt"
)

func InstrumentRequestHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create New Instrument Request")
			deptID := GetUintInput("Enter Department ID: ")

			newRequest := &model.InstrumentRequest{
				DepartmentID: deptID,
				Status:       model.StatusDraft,
			}

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create request:", err)
			} else {
				fmt.Println("Instrument Request created with ID:", newRequest.InstrumentRequestID)
			}
			WaitForEnter()
			fmt.Println("\nPress Enter to continue...")
			WaitForEnter()
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
	fmt.Println("  1:\tCreate Instrument Request")
	fmt.Println("  2:\tList All Instrument Requests")
	fmt.Println("  3:\tGet Instrument Request by ID")
	fmt.Println("  4:\tAdd Instrument to Request")
	fmt.Println("  5:\tShow Request with Details")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func GetUintInput(prompt string) uint {
	var input uint
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func WaitForEnter() {
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
}
