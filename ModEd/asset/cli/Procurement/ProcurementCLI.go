package main

import (
	"ModEd/asset/cli/Procurement/handler"
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
	// resetFlag := flag.Bool("reset", false, "Reset database")
	// blankFlag := flag.Bool("blank", false, "Load seed data to database")

	// flag.Parse()

	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		panic(err)
	}

	// if *blankFlag {
	// 	err = facade.ResetDB()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// } else if *resetFlag {
	// 	err = facade.ResetAndLoadDB()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	return facade
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
