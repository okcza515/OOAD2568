package main

// MEP-1012 Asset

import (
	menu "ModEd/asset/cli/asset/menu"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"flag"
)

func main() {
	assetControllerFacade := newAssetControllerFacade()
	manager := cli.NewCLIMenuManager()
	assetMenu := menu.NewAssetMenuState(manager, assetControllerFacade)

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

func newAssetControllerFacade() *controller.AssetControllerFacade {
	resetFlag := flag.Bool("reset", false, "Reset database")
	blankFlag := flag.Bool("blank", false, "Load seed data to database")

	flag.Parse()

	facade, err := controller.CreateAssetControllerFacade()
	if err != nil {
		panic(err)
	}

	if *blankFlag {
		err = facade.ResetDB()
		if err != nil {
			panic(err)
		}
	} else if *resetFlag {
		err = facade.ResetAndLoadDB()
		if err != nil {
			panic(err)
		}
	}

	return facade
}
