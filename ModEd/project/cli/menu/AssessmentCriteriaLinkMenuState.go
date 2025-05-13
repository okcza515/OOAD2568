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

type AssessmentCriteriaLinkMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewAssessmentCriteriaLinkMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *AssessmentCriteriaLinkMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Assessment Criteria Link Management")

	io := core.NewMenuIO()
	h := handlers.NewAssessmentCriteriaHandler(storer)
	linkHandler := handlers.NewAssessmentCriteriaLinkHandler(storer)

	handlerContext.AddHandler("1", "Define New Criteria", handler.FuncStrategy{
		Action: func() error {
			h.Define(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "List All Criteria", handler.FuncStrategy{
		Action: func() error {
			h.ListAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Update Criteria", handler.FuncStrategy{
		Action: func() error {
			h.Update(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Delete Criteria", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	handlerContext.AddHandler("5", "List Linked Criteria to Assessment", handler.FuncStrategy{
		Action: func() error {
			linkHandler.ListCriteriaLinkedToAssessment(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &AssessmentCriteriaLinkMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *AssessmentCriteriaLinkMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/assessment-criteria-link")
	fmt.Println()
	fmt.Println("Assessment Criteria Link Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *AssessmentCriteriaLinkMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
