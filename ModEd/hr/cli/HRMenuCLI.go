package main

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/migration"
	"ModEd/hr/cli/menu"
	hrController "ModEd/hr/controller"
	"fmt"
)

type HRMenuCLI struct {
	manager   *cli.CLIMenuStateManager
	lastError error
}

func NewHRMenuCLI(manager *cli.CLIMenuStateManager) *HRMenuCLI {
	return &HRMenuCLI{
		manager:   manager,
		lastError: nil,
	}
}

func (app *HRMenuCLI) Run() {
	for {
		// core.ClearScreen()

		if app.lastError != nil {
			fmt.Printf("Error: %v\nPlease try again.\n\n", app.lastError)
			app.lastError = nil
		}

		app.manager.Render()

		userInput := core.GetUserInput("Enter your choice: ")
		if userInput == "exit" {
			fmt.Println("Exiting HR Menu CLI.")
			break
		}

		app.manager.UserInput = userInput
		err := app.manager.HandleUserInput()

		if err != nil {
			app.lastError = err
		}
	}
}

func main() {
	_, err := migration.
		GetInstance().
		MigrateModule(core.MODULE_HR).
		BuildDB()

	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return
	}

	cliManager := cli.NewCLIMenuManager()

	hrManager := hrController.GetHRInstance()

    mainMenu := menu.NewHRMainMenuState(cliManager, hrManager)


	cliManager.SetState(mainMenu)

	app := NewHRMenuCLI(cliManager)

	app.Run()
}
