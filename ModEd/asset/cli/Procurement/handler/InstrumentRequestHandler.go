package handler

import (
	procurement "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"fmt"
)

func InstrumentRequestHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printInstrumentRequestOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create New Instrument Request")
			deptID := util.GetUintInput("Enter Department ID: ")

			newRequest := &model.InstrumentRequest{
				DepartmentID: deptID,
				Status:       model.InstrumentRequestStatusPending,
			}

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create request:", err)
				WaitForEnter()
				break
			}

			fmt.Println("Instrument Request created with ID:", newRequest.InstrumentRequestID)

			newBudgetApproval := &model.BudgetApproval{
				InstrumentRequestID: newRequest.InstrumentRequestID,
				Status:              model.BudgetStatusPending,
			}
			err1 := facade.BudgetApproval.CreateBudgetRequest(newBudgetApproval)
			if err1 != nil {
				fmt.Println("Failed to create budget approval:", err1)
			} else {
				fmt.Println("Budget Approval created with ID:", newBudgetApproval.InstrumentRequestID)
			}

			addMore := util.GetStringInput("\nDo you want to add instruments to this request now? (y/n): ")

			for addMore == "y" || addMore == "Y" {
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

				err := facade.RequestedItem.AddInstrumentToRequest(newRequest.InstrumentRequestID, detail)
				if err != nil {
					fmt.Println("Failed to add instrument:", err)
				} else {
					fmt.Println("Instrument added to request!")
				}

				addMore = util.GetStringInput("\nAdd another instrument? (y/n): ")
			}

			WaitForEnter()
		case "2":
			fmt.Println("List All Instrument Requests")
			ListAllInstrumentRequest(facade)
			// requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			// if err != nil {
			// 	fmt.Println("Failed to list requests:", err)
			// } else {
			// 	fmt.Println("Instrument Requests List:")
			// 	for _, request := range *requests {
			// 		fmt.Printf("ID: %d, DepartmentID: %d, Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
			// 	}
			// }
			WaitForEnter()
		case "3":
			fmt.Println("Get Instrument Request by ID or Name")
			idOrName := util.GetStringInput("Enter Instrument Request ID or Name: ")

			var request *model.InstrumentRequest
			id, err := parseUint(idOrName)
			if err == nil {
				request, err = facade.RequestedItem.GetInstrumentRequestByID(id)
			} else {
				request, err = facade.RequestedItem.GetInstrumentRequestByName(idOrName)
			}

			if err != nil {
				fmt.Println("Failed to get request:", err)
			} else {
				fmt.Printf("Instrument Request found: ID: %d, DepartmentID: %d, Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
			}
			WaitForEnter()
		case "4":
			fmt.Println("Add Instrument to Existing Request")

			requestID := util.GetUintInput("Enter Instrument Request ID: ")
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
				InstrumentRequestID: requestID,
			}

			err := facade.RequestedItem.AddInstrumentToRequest(requestID, detail)
			if err != nil {
				fmt.Println("Failed to add instrument:", err)
			} else {
				fmt.Println("Instrument added to request!")
			}
			WaitForEnter()
		case "5":
			ListAllInstrumentRequest(facade)
			// requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			// if err != nil {
			// 	fmt.Println("Failed to retrieve requests:", err)
			// 	WaitForEnter()
			// 	break
			// }

			// if len(*requests) == 0 {
			// 	fmt.Println("No instrument requests available.")
			// 	WaitForEnter()
			// 	break
			// }

			// fmt.Println("Available Instrument Requests:")
			// for _, request := range *requests {
			// 	fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
			// }

			fmt.Println()
			fmt.Println("5: View Request with Instrument Details")
			requestID := util.GetUintInput("Enter Instrument Request ID: ")

			request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
			if err != nil {
				fmt.Println("Failed to get request with details:", err)
			} else {
				fmt.Printf("\nInstrument Request ID: %d\n", request.InstrumentRequestID)
				fmt.Printf("Department ID: %d\n", request.DepartmentID)
				fmt.Printf("Status: %s\n", request.Status)
				fmt.Println("Instruments:")

				if len(request.Instruments) == 0 {
					fmt.Println("  No instruments found for this request.")
				} else {
					for _, instrument := range request.Instruments {
						fmt.Printf("  - Instrument ID: %d\n", instrument.InstrumentDetailID)
						fmt.Printf("    Label: %s\n", instrument.InstrumentLabel)
						if instrument.Description != nil {
							fmt.Printf("    Description: %s\n", *instrument.Description)
						}
						fmt.Printf("    Category ID: %d\n", instrument.CategoryID)
						fmt.Printf("    Quantity: %d\n", instrument.Quantity)
						fmt.Printf("    Estimated Price: %.2f\n", instrument.EstimatedPrice)
						fmt.Println()
					}
				}
			}
			WaitForEnter()
		case "6":
			fmt.Println("Edit Instrument Items in Request")
			ListAllInstrumentRequest(facade)
			// requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			// if err != nil {
			// 	fmt.Println("Failed to retrieve requests:", err)
			// 	WaitForEnter()
			// 	break
			// }

			// if len(*requests) == 0 {
			// 	fmt.Println("No requests available.")
			// 	WaitForEnter()
			// 	break
			// }

			// fmt.Println("Available Requests:")
			// for _, r := range *requests {
			// 	fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n", r.InstrumentRequestID, r.DepartmentID, r.Status)
			// }

			requestID := util.GetUintInput("\nEnter Instrument Request ID: ")

			request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
			if err != nil {
				fmt.Println("Request not found:", err)
				WaitForEnter()
				break
			}

			if len(request.Instruments) == 0 {
				fmt.Println("No instrument items in this request.")
				WaitForEnter()
				break
			}

			fmt.Println("\nInstrument Items:")
			for _, inst := range request.Instruments {
				fmt.Printf("  ID: %d | Label: %s | Qty: %d | Price: %.2f\n",
					inst.InstrumentDetailID, inst.InstrumentLabel, inst.Quantity, inst.EstimatedPrice)
			}

			detailID := util.GetUintInput("\nEnter Instrument Detail ID to edit: ")

			var selected *model.InstrumentDetail
			for _, inst := range request.Instruments {
				if inst.InstrumentDetailID == detailID {
					selected = &inst
					break
				}
			}

			if selected == nil {
				fmt.Println("Instrument Detail not found in request.")
				WaitForEnter()
				break
			}

			fmt.Println("\nCurrent Values:")
			fmt.Printf("  Label: %s\n", selected.InstrumentLabel)
			fmt.Printf("  Description: %s\n", deref(selected.Description))
			fmt.Printf("  Category ID: %d\n", selected.CategoryID)
			fmt.Printf("  Quantity: %d\n", selected.Quantity)
			fmt.Printf("  Estimated Price: %.2f\n", selected.EstimatedPrice)

			newLabel := util.GetStringInput("New Label (Enter to keep): ")
			newDesc := util.GetStringInput("New Description (Enter to keep): ")
			newCategory := util.GetStringInput("New Category ID (Enter to keep): ")
			newQty := util.GetStringInput("New Quantity (Enter to keep): ")
			newPrice := util.GetStringInput("New Estimated Price (Enter to keep): ")

			// Update fields if provided
			if newLabel != "" {
				selected.InstrumentLabel = newLabel
			}
			if newDesc != "" {
				selected.Description = &newDesc
			}
			if newCategory != "" {
				if val, err := parseUint(newCategory); err == nil {
					selected.CategoryID = val
				}
			}
			if newQty != "" {
				if val, err := parseUint(newQty); err == nil {
					selected.Quantity = int(val)
				}
			}
			if newPrice != "" {
				var price float64
				_, err := fmt.Sscanf(newPrice, "%f", &price)
				if err == nil {
					selected.EstimatedPrice = price
				}
			}

			err = facade.RequestedItem.UpdateInstrumentDetail(selected.InstrumentDetailID, selected)
			if err != nil {
				fmt.Println("Failed to update item:", err)
			} else {
				fmt.Println("Instrument item updated successfully.")
			}

			WaitForEnter()
		}
		util.ClearScreen()
	}

	util.ClearScreen()
}

func printInstrumentRequestOption() {
	fmt.Println(":/Procurement/RequestItem")
	fmt.Println()
	fmt.Println("--RequestItem Function--")
	fmt.Println("  1:\tCreate Instrument Request")
	fmt.Println("  2:\tList All Instrument Requests")
	fmt.Println("  3:\tGet Instrument Request by ID")
	fmt.Println("  4:\tAdd Instrument to Request")
	fmt.Println("  5:\tView Request with Instrument Details")
	fmt.Println("  6:\tEdit Request by ID")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func WaitForEnter() {
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
}

func parseUint(input string) (uint, error) {
	var result uint
	_, err := fmt.Sscanf(input, "%d", &result)
	return result, err
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ListAllInstrumentRequest(facade *procurement.ProcurementControllerFacade) {
	requests, err := facade.RequestedItem.ListAllInstrumentRequests()
	if err != nil {
		fmt.Println("Failed to list requests:", err)
	} else {
		fmt.Println("Instrument Requests List:")
		for _, request := range *requests {
			fmt.Printf("ID: %d, DepartmentID: %d, Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
		}
	}
}
