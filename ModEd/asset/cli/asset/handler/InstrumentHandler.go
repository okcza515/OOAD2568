package handler

import (
	"ModEd/asset/controller/asset"
	"ModEd/asset/util"
	"fmt"
)

func InstrumentHandler(facade *asset.AssetControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Instrument")
		case "2":
			util.ClearScreen()
			fmt.Println("List all Instrument")

			// ok wavie this is for testing you can remove this in the morning it's 4:00 am and I need to sleep
			logList, err := facade.InstrumentLog.ListAll()
			if err != nil {
				panic(err)
			}

			for _, log := range logList {
				fmt.Println(log)
			}

			util.PressEnterToContinue()

		case "3":
			fmt.Println("Get detail of an Instrument")
		case "4":
			fmt.Println("Update an Instrument")
		case "5":
			fmt.Println("Delete an Instrument")
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
