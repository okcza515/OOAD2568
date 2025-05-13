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

type CommitteeMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewCommitteeMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *CommitteeMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Committee Management")

	io := core.NewMenuIO()
	h := handlers.NewCommitteeHandler(storer)

	handlerContext.AddHandler("1", "Add Committee Member", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "List Committee by Project", handler.FuncStrategy{
		Action: func() error {
			h.ListByProject(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "List Projects by Committee Member", handler.FuncStrategy{
		Action: func() error {
			h.ListProjectsByInstructor(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Remove Committee Member", handler.FuncStrategy{
		Action: func() error {
			h.Remove(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &CommitteeMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *CommitteeMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/committee")
	fmt.Println()
	fmt.Println("Committee Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *CommitteeMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
