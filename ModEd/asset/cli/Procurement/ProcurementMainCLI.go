// MEP-1014
package main

import (
	"ModEd/asset/cli/Procurement/menu"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"fmt"
)

func main() {
	util.ClearScreen()
	util.PrintBanner()

	manager := cli.NewCLIMenuManager()

	mainMenu := menu.NewProcurementMainMenuState(manager)
	manager.SetState(mainMenu)

	for {
		util.ClearScreen()
		util.PrintBanner()

		manager.Render()

		manager.UserInput = util.GetCommandInput()
		if manager.UserInput == "exit" {
			break
		}

		util.ClearScreen()

		err := manager.HandleUserInput()
		if err != nil {
			fmt.Println("Error handling user input:", err)
		}
	}

	util.PrintByeBye()
}
