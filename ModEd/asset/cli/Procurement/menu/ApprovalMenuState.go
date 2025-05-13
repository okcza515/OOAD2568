// MEP-1014
package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type ApprovalMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewApprovalMenuState(manager *cli.CLIMenuStateManager) *ApprovalMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &ApprovalMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	handlerContext.AddHandler("1", "Budget Approval", handler.FuncStrategy{
		Action: func() error {
			helper.HandleApprovalOption(&facade.BudgetApproval)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Procurement Approval", handler.FuncStrategy{
		Action: func() error {
			helper.HandleApprovalOption(&facade.Procurement)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Acceptance Approval", handler.FuncStrategy{
		Action: func() error {
			helper.HandleApprovalOption(&facade.Acceptance)
			return nil
		},
	})

	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *ApprovalMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/approval")
	fmt.Println()
	fmt.Println("Approval Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *ApprovalMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
