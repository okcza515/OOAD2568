package handler

import (
	"ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"fmt"
)

func InstrumentRequestHandler(facade *controller.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printInstrumentRequestOption()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create New Instrument Request")
			deptID := util.GetUintInput("Enter Department ID: ")

			newRequest := controller.NewInstrumentRequestBuilder().
				WithDepartmentID(deptID).
				WithStatus(model.InstrumentRequestStatusPending).
				Build()

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create request:", err)
				WaitForEnter()
				break
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

			for addMore == "y" || addMore == "Y" {
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

				err := facade.RequestedItem.AddInstrumentToRequest(newRequest.InstrumentRequestID, detail)
				if err != nil {
					fmt.Println("Failed to add instrument:", err)
				} else {
					fmt.Println("Instrument added to request!")
					err = facade.RequestedItem.UpdateTotalEstimatedPrice(newRequest.InstrumentRequestID)
					if err != nil {
						fmt.Println("Failed to update total estimated price:", err)
					}
				}

				addMore = util.GetStringInput("\nAdd another instrument? (y/n): ")
			}

			WaitForEnter()
		case "2":
			fmt.Println("List All Instrument Requests")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showAvailableRequests(requests, err) {
				break
			}

			WaitForEnter()
		case "3":
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showAvailableRequests(requests, err) {
				break
			}

			fmt.Println()
			fmt.Println("Add Instrument to Existing Request")

			requestID := util.GetUintInput("Enter Instrument Request ID: ")
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
				WithRequestID(requestID).
				Build()

			err = facade.RequestedItem.AddInstrumentToRequest(requestID, detail)
			if err != nil {
				fmt.Println("Failed to add instrument:", err)
			} else {
				fmt.Println("Instrument added to request!")
				err = facade.RequestedItem.UpdateTotalEstimatedPrice(requestID)
				if err != nil {
					fmt.Println("Failed to update total estimated price:", err)
				}
			}
			WaitForEnter()
		case "4":
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showAvailableRequests(requests, err) {
				break
			}

			fmt.Println()
			fmt.Println("View Request with Instrument Details")
			requestID := util.GetUintInput("Enter Instrument Request ID: ")

			request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
			if err != nil {
				fmt.Println("Failed to get request with details:", err)
			} else {
				fmt.Printf("\nInstrument Request ID: %d\n", request.InstrumentRequestID)
				fmt.Printf("Department ID: %d\n", request.DepartmentID)
				fmt.Printf("Status: %s\n", request.Status)
				fmt.Printf("Total Estimated Price: %.2f\n", request.TotalEstimatedPrice)
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
		case "5":
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showAvailableRequests(requests, err) {
				break
			}
			fmt.Println()
			fmt.Println("Edit Instrument in Request")
			requestID := util.GetUintInput("Enter Instrument Request ID: ")

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
				err = facade.RequestedItem.UpdateTotalEstimatedPrice(requestID)
				if err != nil {
					fmt.Println("Failed to update total estimated price:", err)
				}
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
	fmt.Println("  3:\tAdd Instrument to Existing Request")
	fmt.Println("  4:\tView Request with Instrument Details")
	fmt.Println("  5:\tEdit Instrument in Request")
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

func showAvailableRequests(requests *[]model.InstrumentRequest, err error) bool {
	if err != nil {
		fmt.Println("Failed to retrieve requests:", err)
		return false
	}

	if len(*requests) == 0 {
		fmt.Println("No instrument requests available.")
		return false
	}

	fmt.Println("Available Instrument Requests:")
	for _, request := range *requests {
		fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
	}

	return true
}
