package main

import (
	handler "ModEd/asset/cli/Procurement/handler"
	controller "ModEd/asset/controller/Procurement"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	facade := initProgram()

	inputBuffer := ""

	for inputBuffer != "exit" {
		util.ClearScreen()
		util.PrintBanner()

		printOption()

		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			handler.RequestItemHandler(facade)
		case "2":
			fmt.Println("Not implemented yet...")
		case "3":
			fmt.Println("Not implemented yet...")
		case "4":
			fmt.Println("Not implemented yet...")
		case "5":
			fmt.Println("Not implemented yet...")
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
	fmt.Println(":/RequestItem")
	fmt.Println()
	fmt.Println("Welcome to ModEd Procurement Service CLI!")
	fmt.Println("Here is the list of page you can use, choose wisely!")
	fmt.Println("  1:\tRequest Item Page")
	fmt.Println("  2:\t<null> Page")
	fmt.Println("  3:\t<null> Page")
	fmt.Println("  4:\t<null> Page")
	fmt.Println("  5:\t<null> Page")
	fmt.Println("  6:\t<null> Page")
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
