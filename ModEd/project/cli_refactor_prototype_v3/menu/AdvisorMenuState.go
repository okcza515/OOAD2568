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

type AdvisorMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAdvisorMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *AdvisorMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Advisor Management")

	io := core.NewMenuIO()
	advisor := handlers.NewAdvisorHandler(storer)

	handlerContext.AddHandler("1", "Assign Advisor to Project", handler.FuncStrategy{
		Action: func() error {
			advisor.AssignAdvisor(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Update Advisor Role", handler.FuncStrategy{
		Action: func() error {
			advisor.UpdateAdvisorRole(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Remove Advisor", handler.FuncStrategy{
		Action: func() error {
			advisor.RemoveAdvisor(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "List Advisors by Project", handler.FuncStrategy{
		Action: func() error {
			advisor.ListAdvisorsByProject(io)
			return nil
		},
	})

	handlerContext.AddHandler("5", "List Projects by Instructor", handler.FuncStrategy{
		Action: func() error {
			advisor.ListProjectsByInstructor(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &AdvisorMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *AdvisorMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/advisor")
	fmt.Println()
	fmt.Println("Advisor Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AdvisorMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
