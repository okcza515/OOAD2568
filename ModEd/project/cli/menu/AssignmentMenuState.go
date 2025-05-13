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

type AssignmentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAssignmentMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *AssignmentMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Assignment Management")

	io := core.NewMenuIO()
	h := handlers.NewAssignmentHandler(storer)

	handlerContext.AddHandler("1", "View All Assignments", handler.FuncStrategy{
		Action: func() error {
			h.ViewAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Add New Assignment", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Delete Assignment", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &AssignmentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *AssignmentMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/assignment")
	fmt.Println()
	fmt.Println("Assignment Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssignmentMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
