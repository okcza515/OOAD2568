package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type AcceptedInstrumentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
	facade         *controller.ProcurementControllerFacade
}

func NewAcceptedInstrumentMenuState(manager *cli.CLIMenuStateManager) *AcceptedInstrumentMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &AcceptedInstrumentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		facade:         facade,
	}

	// List All Created Instruments
	handlerContext.AddHandler("1", "List All Created Instruments", handler.FuncStrategy{
		Action: func() error {
			helper.HandleInstrumentOption(facade)
			return nil
		},
	})

	// View Instrument Details
	handlerContext.AddHandler("2", "View Instrument Details", handler.FuncStrategy{
		Action: func() error {
			helper.HandleInstrumentDetails(facade)
			return nil
		},
	})

	// Create Instruments from Acceptance
	handlerContext.AddHandler("3", "Create Instruments from Acceptance", handler.FuncStrategy{
		Action: func() error {
			helper.HandleCreateInstrumentFromAcceptance(facade)
			return nil
		},
	})

	// Import Instrument List
	handlerContext.AddHandler("4", "Import Instrument List", handler.FuncStrategy{
		Action: func() error {
			helper.HandleImportInstrument(facade)
			return nil
		},
	})

	// Back to Main Menu
	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *AcceptedInstrumentMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/accepted-instrument")
	fmt.Println()
	fmt.Println("Accepted Instrument Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *AcceptedInstrumentMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
