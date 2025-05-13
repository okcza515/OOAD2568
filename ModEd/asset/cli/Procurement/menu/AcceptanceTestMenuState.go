// MEP-1014
package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type AcceptanceTestMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
	facade         *controller.ProcurementControllerFacade
}

func NewAcceptanceTestMenuState(manager *cli.CLIMenuStateManager) *AcceptanceTestMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &AcceptanceTestMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		facade:         facade,
	}

	handlerContext.AddHandler("1", "List All Acceptance Requests", handler.FuncStrategy{
		Action: func() error {
			helper.ListAllAcceptanceRequests(menu.facade)
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("2", "Acceptance Test", handler.FuncStrategy{
		Action: func() error {
			helper.ListAllAcceptanceRequests(menu.facade)
			fmt.Println("View Quotation Details by Acceptance ID")
			id := util.GetUintInput("Enter Acceptance ID: ")
			helper.PrintQuotationDetailsByAcceptance(menu.facade, id)
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("3", "Import Criteria", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Import Criteria from JSON")
			filename := util.GetStringInput("Enter path to the JSON file (data/criteria.json): ")
			err := controller.ImportCriteriaFromJSON(menu.facade.GetDB(), filename)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Import successful.")
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *AcceptanceTestMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/acceptance-test")
	fmt.Println()
	fmt.Println("Acceptance Test Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *AcceptanceTestMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
