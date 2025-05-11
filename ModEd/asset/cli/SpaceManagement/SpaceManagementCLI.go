package main

import (
	"ModEd/asset/cli/spacemanagement/menu"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/migration"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	db, err := initialSpaceManagementCLI()
	if err != nil {
		panic(err)
	}

	manager := cli.NewCLIMenuManager()
	spaceManagementMenu := menu.NewSpaceManagementState(db, manager)

	manager.SetState(spaceManagementMenu)

	for {
		manager.Render()

		manager.UserInput = util.GetCommandInput()
		if manager.UserInput == "exit" {
			break
		}
		err := manager.HandleUserInput()
		if err != nil {
			panic(err)
		}
	}
}

func initialSpaceManagementCLI() (db *gorm.DB, err error) {
	util.ClearScreen()
	optionFlag := flag.String("option", "", "seed")
	flag.Parse()
	db, err = migration.GetInstance().MigrateModule(core.MODULE_SPACEMANAGEMENT).BuildDB()

	if err != nil {
		panic(err)
	}

	if *optionFlag != "" {
		instance := controller.GetSpaceManagementInstance(db)
		switch *optionFlag {
		case "reset":
			fmt.Println("Resetting database...")
			err = instance.ResetDatabase()
		case "seed":
			fmt.Println("Seeding database...")
			err = instance.LoadSeedData()
		default:
			fmt.Println("Invalid option, please use 'reset' or 'seed'")
		}
	}

	return db, err
}
