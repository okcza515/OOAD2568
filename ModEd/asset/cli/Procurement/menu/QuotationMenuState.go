package menu

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/asset/cli/Procurement/helper"
	"fmt"
)

type QuotationMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewQuotationMenuState(manager *cli.CLIMenuStateManager) *QuotationMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &QuotationMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	handlerContext.AddHandler("1", "Import Quotations", handler.FuncStrategy{
		Action: func() error {
			filename := util.GetStringInput("Enter path to the JSON file (data/quotations.json): ")
			if err := helper.ImportQuotationsFromJSON(facade.GetDB(), filename); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Import successful.")
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("2", "Import Quotation Details", handler.FuncStrategy{
		Action: func() error {
			filename := util.GetStringInput("Enter path to the JSON file (data/quotationdetail.json): ")
			if err := helper.ImportQuotationDetailsFromJSON(facade.GetDB(), filename); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Import successful.")
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("3", "Show Quotations by TOR ID", handler.FuncStrategy{
		Action: func() error {
			helper.ShowQuotationsByTORID(facade.GetDB())
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("4", "Quotation Selection", handler.FuncStrategy{
		Action: func() error {
			helper.SelectQuotation(facade.GetDB())
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *QuotationMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/quotation")
	fmt.Println()
	fmt.Println("Quotation Management Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *QuotationMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
