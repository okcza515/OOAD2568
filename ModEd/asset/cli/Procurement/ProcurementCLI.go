package main

import (
	handler "ModEd/asset/cli/Procurement/handler"
	controller "ModEd/asset/controller"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	facade := initProgram()

	inputBuffer := ""

	for inputBuffer != "exit" {
		util.ClearScreen()

		printOption()

		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			handler.InstrumentRequestHandler(facade)
		case "2":
			handler.ApprovalHandler(facade)
		case "3":
			handler.ProcurementHandler(facade)
		case "4":
			handler.TORHandler(facade)
		case "5":
			handler.QuotationSupplierHandler(facade)
		case "6":
			fmt.Println("Not implemented yet...")
		}
	}

	fmt.Println("Good Bye!")
}

func initProgram() *controller.ProcurementControllerFacade {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Printf("Init failed: %v\n", err)
		panic(err)
	}

	return facade
}

func printOption() {
	fmt.Println(":/Procurement")
	fmt.Println()
	fmt.Println("Welcome to ModEd Procurement Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tRequest Instrument Page")
	fmt.Println("  2:\tProcurement Page")
	fmt.Println("  3\tApproval Page")
	fmt.Println("  4:\tNot implemented yet... Page")
	fmt.Println("  5:\tNot implemented yet... Page")
	fmt.Println("  6:\tNot implemented yet... Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
