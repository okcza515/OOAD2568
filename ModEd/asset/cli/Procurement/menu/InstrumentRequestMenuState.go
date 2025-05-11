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
	// Get the ProcurementControllerFacade instance
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

	// Register Handlers
	handlerContext.AddHandler("1", "Create New Instrument Request", handler.FuncStrategy{
		Action: func() error {
			deptID := util.GetUintInput("Enter Department ID: ")

			newRequest := controller.NewInstrumentRequestBuilder().
				WithDepartmentID(deptID).
				WithStatus(model.InstrumentRequestStatusPending).
				Build()

			for {
				fmt.Println("\n--- Add Instrument ---")
				label := util.GetStringInput("Enter Instrument Label: ")
				desc := util.GetStringInput("Enter Description: ")
				categoryID := util.GetUintInput("Enter Category ID: ")
				estimatedPrice := util.GetFloatInput("Enter Estimated Price: ")
				quantity := util.GetUintInput("Enter Quantity: ")

				detail := controller.NewInstrumentDetailBuilder().
					WithLabel(label).
					WithDescription(desc).
					WithCategoryID(categoryID).
					WithEstimatedPrice(estimatedPrice).
					WithQuantity(int(quantity)).
					WithRequestID(newRequest.InstrumentRequestID).
					Build()

				newRequest.Instruments = append(newRequest.Instruments, *detail)

				addMore := util.GetStringInput("\nDo you want to add another instrument? (y/n): ")
				if addMore != "y" && addMore != "Y" {
					break
				}
			}

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create Instrument Request:", err)
				util.PressEnterToContinue()
				return err
			}

			fmt.Println("\nInstrument Request created successfully with ID:", newRequest.InstrumentRequestID)

			fmt.Println("\n--- Request Summary ---")
			fmt.Printf("Department ID: %d\n", newRequest.DepartmentID)
			fmt.Println("Instruments:")
			for _, instrument := range newRequest.Instruments {
				fmt.Printf("  - Label: %s | Qty: %d | Price: %.2f\n",
					instrument.InstrumentLabel, instrument.Quantity, instrument.EstimatedPrice)
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("2", "List All Instrument Requests", handler.FuncStrategy{
		Action: func() error {
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil {
				fmt.Println("Failed to retrieve instrument requests:", err)
				util.PressEnterToContinue()
				return err
			}
			if len(*requests) == 0 {
				fmt.Println("No available instrument requests found.")
				util.PressEnterToContinue()
				return nil
			}

			helper.DisplayRequestList(*requests)
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Instrument Request Details", handler.FuncStrategy{
		Action: func() error {
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil {
				fmt.Println("Failed to retrieve instrument requests:", err)
				util.PressEnterToContinue()
				return err
			}
			if len(*requests) == 0 {
				fmt.Println("No available instrument requests found.")
				util.PressEnterToContinue()
				return nil
			}

			helper.DisplayRequestList(*requests)

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

			for {
				fmt.Println("\n--- Add Instrument ---")

				label := util.GetStringInput("Enter Instrument Label: ")
				desc := util.GetStringInput("Enter Description: ")
				categoryID := util.GetUintInput("Enter Category ID: ")
				estimatedPrice := util.GetFloatInput("Enter Estimated Price: ")
				quantity := util.GetUintInput("Enter Quantity: ")

				detail := controller.NewInstrumentDetailBuilder().
					WithLabel(label).
					WithDescription(desc).
					WithCategoryID(categoryID).
					WithEstimatedPrice(estimatedPrice).
					WithQuantity(int(quantity)).
					WithRequestID(request.InstrumentRequestID).
					Build()

				err := facade.RequestedItem.AddInstrumentToRequest(request.InstrumentRequestID, detail)
				if err != nil {
					fmt.Println("Failed to add instrument to the request:", err)
				} else {
					fmt.Println("Instrument added successfully!")
				}

				addMore := util.GetStringInput("\nDo you want to add another instrument? (y/n): ")
				if addMore != "y" && addMore != "Y" {
					break
				}
			}

			util.PressEnterToContinue()
			return nil
		},
	})

	// Register Handler for Editing
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

	// Register Back Handler
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
