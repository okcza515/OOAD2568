package main

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/migration"
	"ModEd/hr/cli/menu"
	hrController "ModEd/hr/controller"
	"flag"
	"fmt"
)

var (
	databasePath = flag.String("database", "data/ModEd.bin", "Path to SQLite Database file")
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
	db, err := migration.
		GetInstance().
		MigrateModule(core.MODULE_HR).
		BuildDB()

	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return
	}

	manager := cli.NewCLIMenuManager()

	mainMenu := menu.NewHRMainMenuState(manager,
		hrController.NewStudentHRController(db),
		hrController.NewInstructorHRController(db))

	manager.SetState(mainMenu)

	app := NewHRMenuCLI(manager)

	app.Run()
}
