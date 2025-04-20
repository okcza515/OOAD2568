package handler

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
)

func SupplyHandler(facade *controller.AssetControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printSupplyOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Supply")
		case "2":
			supplies, err := facade.Supply.GetAll()
			if err != nil {
				panic(err)
			}
			for _, supply := range *supplies {
				fmt.Println(supply)
			}
		case "3":
			fmt.Println("Get detail of an Supply")
		case "4":
			fmt.Println("Update an Supply Information")
		case "5":
			fmt.Println("Update an Supply Stock")
		case "6":
			fmt.Println("Delete an Supply")
		case "7":
			fmt.Println("Miagration Supply")
		default:
			fmt.Println("Invalid Command")
		}
		util.PressEnterToContinue()
		util.ClearScreen()
	}

	util.ClearScreen()
}

func printSupplyOption() {
	fmt.Println()
	fmt.Println("  1:\tAdd new Supply")
	fmt.Println("  2:\tList all Supply")
	fmt.Println("  3:\tGet detail of an Supply")
	fmt.Println("  4:\tUpdate an Supply Information")
	fmt.Println("  5:\tUpdate an Supply Stock")
	fmt.Println("  6:\tDelete an Supply")
	fmt.Println("  6:\tMiagration Supply")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}
