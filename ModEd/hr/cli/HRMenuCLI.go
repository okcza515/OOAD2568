package main

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/hr/cli/menu"
)

func main() {

	manager := cli.NewCLIMenuManager()
	assetMenu := menu.NewHRMainMenuState(manager)

	manager.SetState(assetMenu)

	for {
		// core.ClearScreen()

		manager.Render()

		manager.UserInput = core.GetUserInput("Enter your choice: ")
		if manager.UserInput == "exit" {
			break
		}

		// core.ClearScreen()

		err := manager.HandleUserInput()
		if err != nil {
			panic(err)
		}
	}

}
