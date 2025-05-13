package main

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/project/cli_refactor_prototype_v3/menu"
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

func main() {
	db := utils.OpenDatabase("project.db")
	db.Exec("PRAGMA foreign_keys = ON;")

	instance := controller.CreateInstance(db)

	manager := cli.NewCLIMenuManager()
	mainMenuState := menu.NewProjectMainMenuState(manager, instance)
	manager.SetState(mainMenuState)

	for {
		manager.Render()

		manager.UserInput = util.GetCommandInput()
		if manager.UserInput == "exit" {
			break
		}

		err := manager.HandleUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			util.PressEnterToContinue()
		}
	}
}
