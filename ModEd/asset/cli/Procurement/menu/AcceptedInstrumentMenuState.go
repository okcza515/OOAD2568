// MEP-1014
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

	if err != nil || facade == nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	if _, err := facade.Instrument.ListAllInstruments(); err != nil {
		fmt.Println("Failed to initialize Instrument Controller. Database connection may be nil:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()

	menu := &AcceptedInstrumentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
		facade:         facade,
	}

	handlerContext.AddHandler("1", "List All Created Instruments", handler.FuncStrategy{
		Action: func() error {
			helper.HandleInstrumentOption(facade)
			return nil
		},
	})

	handlerContext.AddHandler("2", "View Instrument Details", handler.FuncStrategy{
		Action: func() error {
			helper.PrintInstrumentList(facade)

			fmt.Println()
			fmt.Println("Enter the Instrument ID to view details:")
			helper.HandleInstrumentDetails(facade)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Create Instruments from Acceptance", handler.FuncStrategy{
		Action: func() error {
			helper.HandleCreateInstrumentFromAcceptance(facade)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Import Instrument List", handler.FuncStrategy{
		Action: func() error {
			helper.HandleImportInstrument(facade)
			return nil
		},
	})

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
