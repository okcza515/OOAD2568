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
			// budgetapproverID := util.GetUintInput("Enter Approver ID: ")

			newRequest := &model.InstrumentRequest{
				DepartmentID: deptID,
				Status:       model.InstrumentRequestStatusPending,
			}

			err := facade.RequestedItem.CreateInstrumentRequest(newRequest)
			if err != nil {
				fmt.Println("Failed to create request:", err)
			} else {
				fmt.Println("Instrument Request created with ID:", newRequest.InstrumentRequestID)
			}

			newBudgetApproval := &model.BudgetApproval{
				// ApproverID: budgetapproverID,
				InstrumentRequestID: newRequest.InstrumentRequestID,
				Status:              model.BudgetStatusPending,
			}
			err1 := facade.BudgetApproval.CreateBudgetRequest(newBudgetApproval)

			if err1 != nil {
				fmt.Println("Failed to create request:", err1)
			} else {
				fmt.Println("Budget Approval created with ID:", newBudgetApproval.InstrumentRequestID)
			}
			WaitForEnter()
		case "2":
			fmt.Println("List All Instrument Requests")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if err != nil {
				fmt.Println("Failed to list requests:", err)
			} else {
				fmt.Println("Instrument Requests List:")
				for _, request := range *requests {
					fmt.Printf("ID: %d, DepartmentID: %d, Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
				}
			}
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
			fmt.Println("Show Instrument Request with Details")
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
	fmt.Println("  5:\tShow Request with Details")
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
