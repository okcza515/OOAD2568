package main

import (
	"ModEd/asset/cli/asset/handler"
	controller "ModEd/asset/controller/asset"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	facade, err := controller.CreateAssetControllerFacade()
	if err != nil {
		panic("err: initialize controllers failed")
	}

	inputBuffer := ""

	for inputBuffer != "exit" {
		util.ClearScreen()
		util.PrintBanner()

		printOption()

		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Not implemented yet...")
		case "2":
			handler.InstrumentHandler(facade)
		case "3":
			handler.SupplyHandler(facade)
		case "4":
			fmt.Println("Not implemented yet...")
		case "5":
			fmt.Println("Not implemented yet...")
		case "6":
			fmt.Println("Not implemented yet...")
		}
	}

	util.PrintByeBye()
}

func printOption() {
	fmt.Println(":/asset")
	fmt.Println()
	fmt.Println("Welcome to ModEd Asset Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tCategory Page")
	fmt.Println("  2:\tInstrument Page")
	fmt.Println("  3:\tSupply Page")
	fmt.Println("  4:\tBorrow Page")
	fmt.Println("  5:\tInstrument Log Page")
	fmt.Println("  6:\tSupply Log Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
