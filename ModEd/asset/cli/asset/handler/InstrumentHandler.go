package handler

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/utils/deserializer"
	"fmt"
)

func InstrumentHandler(facade *controller.AssetControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printOption()
		inputBuffer = util.GetCommandInput()

		util.ClearScreen()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Instrument")

			path := ""
			fmt.Println("Please enter the path of the instrument file (csv or json): ")
			_, _ = fmt.Scanln(&path)

			var insModel []model.Instrument

			fd, err := deserializer.NewFileDeserializer(path)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}
			err = fd.Deserialize(&insModel)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}

			err = facade.Instrument.InsertMany(insModel)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}

			fmt.Println("Instrument successfully inserted!")
			util.PressEnterToContinue()

		case "2":
			fmt.Println("List all Instrument")

			instList, err := facade.Instrument.ListAll()
			if err != nil {
				panic(err)
			}

			for _, inst := range instList {
				fmt.Println(inst)
			}

			util.PressEnterToContinue()

		case "3":
			fmt.Println("Get detail of an Instrument")
		case "4":
			fmt.Println("Update an Instrument")
		case "5":
			fmt.Println("Delete an Instrument")
		default:
			fmt.Println("Invalid Command")
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
