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

type AssessmentScoreMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAssessmentScoreMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *AssessmentScoreMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Assessment Score Management")

	io := core.NewMenuIO()
	scoreHandler := handlers.NewAssessmentScoreHandler(storer)

	handlerContext.AddHandler("1", "View Project Scores", handler.FuncStrategy{
		Action: func() error {
			scoreHandler.ViewProjectScores(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Submit Advisor Score", handler.FuncStrategy{
		Action: func() error {
			scoreHandler.SubmitAdvisorScore(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Submit Committee Score", handler.FuncStrategy{
		Action: func() error {
			scoreHandler.SubmitCommitteeScore(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &AssessmentScoreMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *AssessmentScoreMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/assessment-score")
	fmt.Println()
	fmt.Println("Assessment Score Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssessmentScoreMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
