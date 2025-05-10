package menu

import (
	// "ModEd/asset/cli/Procurement/helper"
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
			fmt.Println("🔹 Create New Instrument Request")
			deptID := util.GetUintInput("Enter Department ID: ")

			newRequest := &model.InstrumentRequest{
				DepartmentID: deptID,
				Status:       model.InstrumentRequestStatusPending,
			}

			// Add Instruments to the Request
			for {
				fmt.Println("\n--- Add Instrument ---")
				label := util.GetStringInput("Enter Instrument Label: ")
				desc := util.GetStringInput("Enter Description: ")
				categoryID := util.GetUintInput("Enter Category ID: ")
				estimatedPrice := util.GetFloatInput("Enter Estimated Price: ")
				quantity := util.GetUintInput("Enter Quantity: ")

				detail := &model.InstrumentDetail{
					InstrumentLabel:     label,
					Description:         &desc,
					CategoryID:          categoryID,
					EstimatedPrice:      estimatedPrice,
					Quantity:            int(quantity),
					InstrumentRequestID: newRequest.InstrumentRequestID,
				}

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
			fmt.Println("List All Instrument Requests")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil {
				fmt.Println("Failed to list instrument requests:", err)
				return err
			}

			if len(*requests) == 0 {
				fmt.Println("No Instrument Requests found.")
				util.PressEnterToContinue()
				return nil
			}

			fmt.Println("Instrument Requests List:")
			for _, request := range *requests {
				fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n",
					request.InstrumentRequestID, request.DepartmentID, request.Status)
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddHandler("3", "View Instrument Request Details", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("View Instrument Request Details")
			requestID := util.GetUintInput("Enter Instrument Request ID: ")
			request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
			if err != nil {
				fmt.Println("Failed to retrieve request details:", err)
				util.PressEnterToContinue()
				return err
			}

			fmt.Printf("\nInstrument Request ID: %d\nDepartment ID: %d\nStatus: %s\n",
				request.InstrumentRequestID, request.DepartmentID, request.Status)

			if len(request.Instruments) == 0 {
				fmt.Println("No instruments found.")
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
			fmt.Println("Add Instrument to Existing Request")

			// List All Requests
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil || len(*requests) == 0 {
				fmt.Println("No available requests to add instruments.")
				util.PressEnterToContinue()
				return err
			}

			fmt.Println("Select an Instrument Request ID to add an instrument:")
			for _, req := range *requests {
				fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n",
					req.InstrumentRequestID, req.DepartmentID, req.Status)
			}

			requestID := util.GetUintInput("Enter Instrument Request ID: ")

			for {
				fmt.Println("\n--- Add Instrument ---")
				label := util.GetStringInput("Enter Instrument Label: ")
				desc := util.GetStringInput("Enter Description: ")
				categoryID := util.GetUintInput("Enter Category ID: ")
				estimatedPrice := util.GetFloatInput("Enter Estimated Price: ")
				quantity := util.GetUintInput("Enter Quantity: ")

				// Create Instrument Detail
				detail := &model.InstrumentDetail{
					InstrumentLabel:     label,
					Description:         &desc,
					CategoryID:          categoryID,
					EstimatedPrice:      estimatedPrice,
					Quantity:            int(quantity),
					InstrumentRequestID: requestID,
				}

				// Call the Facade to Add
				err := facade.RequestedItem.AddInstrumentToRequest(requestID, detail)
				if err != nil {
					fmt.Println("Failed to add instrument:", err)
				} else {
					fmt.Println("Instrument added to request successfully!")
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
	// handlerContext.AddHandler("5", "Edit Instrument in Request", handler.FuncStrategy{
	// 	Action: func() error {
	// 		fmt.Println("Edit Instrument in Request")

	// 		// Step 1: Select Instrument Request
	// 		request, err := helper.SelectInstrumentRequest(facade)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		// Step 2: Select Instrument
	// 		instrument, err := helper.SelectInstrument(request)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		// Step 3: Edit Instrument Details
	// 		helper.EditInstrumentDetails(instrument)

	// 		// Step 4: Submit the changes
	// 		err = facade.RequestedItem.UpdateInstrumentDetail(instrument.InstrumentDetailID, instrument)
	// 		if err != nil {
	// 			fmt.Println("Failed to update instrument:", err)
	// 		} else {
	// 			fmt.Println("Instrument updated successfully!")
	// 		}
	// 		util.PressEnterToContinue()
	// 		return err
	// 	},
	// })

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
