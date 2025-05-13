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

type SeniorProjectMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewSeniorProjectMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *SeniorProjectMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Senior Project Management")

	io := core.NewMenuIO()
	h := handlers.NewSeniorProjectHandler(storer)

	handlerContext.AddHandler("1", "Create Senior Project", handler.FuncStrategy{
		Action: func() error {
			h.Create(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "List Senior Projects", handler.FuncStrategy{
		Action: func() error {
			h.List(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &SeniorProjectMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *SeniorProjectMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/senior-project")
	fmt.Println()
	fmt.Println("Senior Project Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *SeniorProjectMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
