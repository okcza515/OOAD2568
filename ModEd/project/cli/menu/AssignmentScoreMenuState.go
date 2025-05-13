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

type AssignmentScoreMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAssignmentScoreMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *AssignmentScoreMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Assignment Score Management")

	io := core.NewMenuIO()
	h := handlers.NewAssignmentScoreHandler(storer)

	handlerContext.AddHandler("1", "Insert Advisor Score", handler.FuncStrategy{
		Action: func() error {
			h.InsertAdvisorScore(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Insert Committee Score", handler.FuncStrategy{
		Action: func() error {
			h.InsertCommitteeScore(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Check Assignment Scores", handler.FuncStrategy{
		Action: func() error {
			h.CheckScore(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &AssignmentScoreMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *AssignmentScoreMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/assignment-score")
	fmt.Println()
	fmt.Println("Assignment Score Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssignmentScoreMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
