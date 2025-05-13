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

type ProgressMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewProgressMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *ProgressMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Progress Management")

	io := core.NewMenuIO()
	h := handlers.NewProgressHandler(storer)

	handlerContext.AddHandler("1", "View All Progress", handler.FuncStrategy{
		Action: func() error {
			h.ViewAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Add New Progress", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Progress by ID", handler.FuncStrategy{
		Action: func() error {
			h.ViewByID(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Update Progress Name", handler.FuncStrategy{
		Action: func() error {
			h.UpdateName(io)
			return nil
		},
	})

	handlerContext.AddHandler("5", "Delete Progress", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	handlerContext.AddHandler("6", "Mark as Completed", handler.FuncStrategy{
		Action: func() error {
			h.MarkCompleted(io)
			return nil
		},
	})

	handlerContext.AddHandler("7", "Mark as Incomplete", handler.FuncStrategy{
		Action: func() error {
			h.MarkIncomplete(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &ProgressMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *ProgressMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/progress")
	fmt.Println()
	fmt.Println("Progress Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *ProgressMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
