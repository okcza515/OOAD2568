package menu

// MEP-1012 Asset

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/utils/deserializer"
	"fmt"
)

func SupplyHandler(facade *controller.AssetControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		util.PrintBanner()
		printSupplyOption()
		inputBuffer = util.GetCommandInput()

		util.ClearScreen()

		switch inputBuffer {
		case "1":
			fmt.Println("Add new Supply")

			path := ""
			fmt.Println("Please enter the path of the supply file (csv or json): ")
			_, _ = fmt.Scanln(&path)

			var supModel []model.Supply

			fd, err := deserializer.NewFileDeserializer(path)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}
			err = fd.Deserialize(&supModel)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}

			err = facade.Supply.InsertMany(supModel)
			if err != nil {
				fmt.Println(err)
				util.PressEnterToContinue()
				break
			}

			fmt.Println("Supply successfully inserted!")
			util.PressEnterToContinue()
		case "2":
			util.ClearScreen()
			fmt.Println("List all Supply")

			suppList, err := facade.Supply.ListAll()
			if err != nil {
				panic(err)
			}

			for _, supp := range suppList {
				fmt.Println(supp)
			}

			util.PressEnterToContinue()
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
