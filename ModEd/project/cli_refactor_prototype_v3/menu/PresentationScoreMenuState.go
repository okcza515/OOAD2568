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

type PresentationScoreMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewPresentationScoreMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *PresentationScoreMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Presentation Score Management")

	io := core.NewMenuIO()
	h := handlers.NewPresentationScoreHandler(storer)

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

	handlerContext.AddHandler("3", "Check Presentation Scores", handler.FuncStrategy{
		Action: func() error {
			h.CheckScore(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &PresentationScoreMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *PresentationScoreMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/presentation-score")
	fmt.Println()
	fmt.Println("Presentation Score Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *PresentationScoreMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
