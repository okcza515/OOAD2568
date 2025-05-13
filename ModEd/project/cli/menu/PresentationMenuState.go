package menu

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/project/cli_refactor_prototype_v3/handlers"
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

type PresentationMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewPresentationMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *PresentationMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Presentation Management")

	io := core.NewMenuIO()
	h := handlers.NewPresentationHandler(storer)

	handlerContext.AddHandler("1", "View All Presentations", handler.FuncStrategy{
		Action: func() error {
			h.ViewAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Add Presentation", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Presentation by ID", handler.FuncStrategy{
		Action: func() error {
			h.ViewByID(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Update Presentation", handler.FuncStrategy{
		Action: func() error {
			h.Update(io)
			return nil
		},
	})

	handlerContext.AddHandler("5", "Delete Presentation", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &PresentationMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *PresentationMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/presentation")
	fmt.Println()
	fmt.Println("Presentation Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *PresentationMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
