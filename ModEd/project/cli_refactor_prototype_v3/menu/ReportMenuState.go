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

type ReportMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewReportMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *ReportMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Report Management")

	io := core.NewMenuIO()
	h := handlers.NewReportHandler(storer)

	handlerContext.AddHandler("1", "View All Reports", handler.FuncStrategy{
		Action: func() error {
			h.ViewAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Add Report", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Report by ID", handler.FuncStrategy{
		Action: func() error {
			h.ViewByID(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Update Report Due Date", handler.FuncStrategy{
		Action: func() error {
			h.Update(io)
			return nil
		},
	})

	handlerContext.AddHandler("5", "Delete Report", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	handlerContext.AddHandler("6", "Submit Report", handler.FuncStrategy{
		Action: func() error {
			h.Submit(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &ReportMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *ReportMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/report")
	fmt.Println()
	fmt.Println("Report Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *ReportMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
