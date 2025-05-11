package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type InstrumentRequestMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewInstrumentRequestMenuState(manager *cli.CLIMenuStateManager) *InstrumentRequestMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &InstrumentRequestMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	handlerContext.AddHandler("1", "Create New Instrument Request", handler.FuncStrategy{
		Action: func() error {
			deptID := util.GetUintInput("Enter Department ID: ")

			newRequest := controller.NewInstrumentRequestBuilder().
				WithDepartmentID(deptID).
				WithStatus(model.InstrumentRequestStatusPending).
				Build()

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create request:", err)
				util.PressEnterToContinue()
				return nil
			}

			fmt.Println("Instrument Request created with ID:", newRequest.InstrumentRequestID)

			newBudgetApproval := controller.NewBudgetApprovalBuilder().
				WithInstrumentRequestID(newRequest.InstrumentRequestID).
				WithStatus(model.BudgetStatusPending).
				Build()

			err1 := facade.BudgetApproval.CreateBudgetRequest(newBudgetApproval)
			if err1 != nil {
				fmt.Println("Failed to create budget approval:", err1)
			} else {
				fmt.Println("Budget Approval created with ID:", newBudgetApproval.InstrumentRequestID)
			}

			addMore := util.GetStringInput("\nDo you want to add instruments to this request now? (y/n): ")
			if addMore == "y" || addMore == "Y" {
				helper.AddInstrumentsLoopToRequest(facade, newRequest.InstrumentRequestID)
			}

			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("2", "List All Instrument Requests", handler.FuncStrategy{
		Action: func() error {
			helper.ShowAllInstrumentRequests(facade)
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Instrument Request Details", handler.FuncStrategy{
		Action: func() error {
			helper.ShowAllInstrumentRequests(facade)

			requestID := util.GetUintInput("\nEnter Instrument Request ID to View Details: ")

			request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
			if err != nil {
				fmt.Println("Failed to retrieve request details:", err)
				util.PressEnterToContinue()
				return err
			}

			fmt.Printf("\n--- Instrument Request Details ---\n")
			fmt.Printf("Request ID: %d\nDepartment ID: %d\nStatus: %s\n",
				request.InstrumentRequestID, request.DepartmentID, request.Status)

			if len(request.Instruments) == 0 {
				fmt.Println("No instruments found for this request.")
			} else {
				for _, instrument := range request.Instruments {
					fmt.Printf("  - ID: %d | Label: %s | Qty: %d | Price: %.2f\n",
						instrument.InstrumentDetailID, instrument.InstrumentLabel, instrument.Quantity, instrument.EstimatedPrice)
				}
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("4", "Add Instrument to Existing Request", handler.FuncStrategy{
		Action: func() error {
			request, err := helper.SelectInstrumentRequest(facade)
			if err != nil {
				fmt.Println("Failed to select instrument request:", err)
				util.PressEnterToContinue()
				return err
			}

			fmt.Println()
			helper.AddInstrumentsLoopToRequest(facade, request.InstrumentRequestID)
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("5", "Edit Instrument in Request", handler.FuncStrategy{
		Action: func() error {
			request, err := helper.SelectInstrumentRequest(facade)
			if err != nil {
				fmt.Println("Failed to select instrument request:", err)
				util.PressEnterToContinue()
				return err
			}

			instrument, err := helper.SelectInstrument(request)
			if err != nil {
				fmt.Println("Failed to select instrument:", err)
				util.PressEnterToContinue()
				return err
			}

			fmt.Println("\n--- Edit Instrument Details ---")
			helper.EditInstrumentDetails(instrument)

			err = facade.RequestedItem.UpdateInstrumentDetail(instrument.InstrumentDetailID, instrument)
			if err != nil {
				fmt.Println("Failed to update instrument:", err)
			} else {
				fmt.Println("Instrument updated successfully!")
			}
			util.PressEnterToContinue()
			return err
		},
	})

	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))
	return menu
}

func (menu *InstrumentRequestMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/instrument-request")
	fmt.Println()
	fmt.Println("Instrument Request Management Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *InstrumentRequestMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
