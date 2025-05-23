// MEP-1014
package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
	"os"
)

type ProcurementMainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewProcurementMainMenuState(manager *cli.CLIMenuStateManager) *ProcurementMainMenuState {
	handlerContext := handler.NewHandlerContext()
	procurementMenu := &ProcurementMainMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
	manager.AddMenu(string(MENU_PROCUREMENT_MAIN), procurementMenu)

	manager.AddMenu(string(MENU_INSTRUMENT_REQUEST), NewInstrumentRequestMenuState(manager))
	manager.AddMenu(string(MENU_PROCUREMENT), NewProcurementMenuState(manager))
	manager.AddMenu(string(MENU_ACCEPTANCE), NewAcceptanceTestMenuState(manager))
	manager.AddMenu(string(MENU_APPROVAL), NewApprovalMenuState(manager))
	manager.AddMenu(string(MENU_ACCEPTEDINSTRUMENT), NewAcceptedInstrumentMenuState(manager))
	manager.AddMenu(string(MENU_IMPORTFILE), NewImportFileMenuState(manager))

	handlerContext.AddHandler("1", "Instrument Request Management", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_INSTRUMENT_REQUEST))))
	handlerContext.AddHandler("2", "Procurement Management", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT))))
	handlerContext.AddHandler("3", "Acceptance Test", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ACCEPTANCE))))
	handlerContext.AddHandler("4", "Approval Management", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_APPROVAL))))
	handlerContext.AddHandler("5", "Accepted Instrument Management", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_ACCEPTEDINSTRUMENT))))
	handlerContext.AddHandler("6", "Import File", handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_IMPORTFILE))))

	handlerContext.AddHandler("exit", "Exit the application", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("Goodbye!")
			os.Exit(0)
			return nil
		},
	})

	return procurementMenu
}

func (menu *ProcurementMainMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement")
	fmt.Println()
	fmt.Println("Procurement Management Main Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *ProcurementMainMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
