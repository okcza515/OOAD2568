// MEP-1013
package main

import (
	"ModEd/asset/cli/spacemanagement/menu"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/migration"
	"flag"

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

	util.PrintByeBye()
}

func initialSpaceManagementCLI() (db *gorm.DB, err error) {
	optionFlag := flag.String("Reset Database", "", "")
	flag.Parse()
	db, err = migration.GetInstance().MigrateModule(core.MODULE_SPACEMANAGEMENT).BuildDB()

	if err != nil {
		panic(err)
	}
	instance := controller.GetSpaceManagementInstance(db)
	if *optionFlag == "Reset Database" {
		err = instance.ResetDatabase()
		if err != nil {
			panic(err)
		}
	}
	return db, err
}
