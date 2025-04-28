package main

// MEP-1012 Asset

import (
	"ModEd/asset/cli/asset/menu"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/migration"
	"flag"
)

func main() {
	err := initAssetCLI()
	if err != nil {
		panic(err)
	}

	manager := cli.NewCLIMenuManager()
	assetMenu := menu.NewAssetMenuState(manager)

	manager.SetState(assetMenu)

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
			panic(err)
		}
	}

	util.PrintByeBye()
}

func initAssetCLI() error {
	optionFlag := flag.String("option", "nothing here", "Reset database")

	flag.Parse()

	_, err := migration.GetInstance().MigrateModule(core.MODULE_ASSET).BuildDB()

	if err != nil {
		panic(err)
	}

	instance := controller.GetAssetInstance()

	switch *optionFlag {
	case "reset":
		err = instance.ResetAndLoadDB()
	case "blank":
		err = instance.ResetDB()
	default:
	}

	if err != nil {
		panic(err)
	}

	return nil
}
