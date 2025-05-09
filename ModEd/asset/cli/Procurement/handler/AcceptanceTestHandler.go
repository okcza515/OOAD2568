package handler

import (
	"ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"fmt"
)

func AcceptanceTestHandler(facade *controller.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		facade.Acceptance.ListAllApprovals()
		printAcceptanceTestOptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("List All Acceptance Requests")
			acceptanceRequests, err := facade.Acceptance.ListAllApprovals()
			if err != nil {
				fmt.Println("Failed to fetch acceptance requests:", err)
				WaitForEnter()
				break
			}

			if len(acceptanceRequests) == 0 {
				fmt.Println("No acceptance requests found.")
				WaitForEnter()
				break
			}

			fmt.Println("\nAvailable Acceptance Requests:")
			found := false
			for _, req := range acceptanceRequests {
				if req.Procurement != nil && req.Procurement.Status == model.ProcurementStatusApproved {
					found = true
					approverID := "Not Assigned"
					if req.ApproverID != nil {
						approverID = fmt.Sprintf("%d", *req.ApproverID)
					}

					approvalTime := "-"
					if !req.ApprovalTime.IsZero() {
						approvalTime = req.ApprovalTime.Format("2006-01-02 15:04:05")
					}

					fmt.Printf("  - Acceptance ID: %d | Procurement ID: %d | Status: %s | Approver ID: %s | Approval Time: %s\n",
						req.AcceptanceApprovalID, req.ProcurementID, req.Status, approverID, approvalTime)
				}
			}
			if !found {
				fmt.Println("No acceptance requests found with approved procurements.")
			}

			WaitForEnter()
		case "2":
			// caid := util.GetUintInput("Enter category ID: ")
			// facade.AcceptanceTest.PrintCategoriesByIDs(caid)

			WaitForEnter()
		}
	}
}

func printAcceptanceTestOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--Acceptance Functions--")
	fmt.Println("  1:\tList All Acceptance Requests")
	fmt.Println("  2:\tAcceptance test")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

// func PrintCategoriesByTOR(torID uint, facade *controller.ProcurementControllerFacade) error {
// 	details, err := facade.AcceptanceTest.GetQuotationDetailsByTOR(torID)
// 	if err != nil {
// 		return err
// 	}

// 	categorySet := make(map[uint]bool)
// 	for _, d := range details {
// 		categorySet[d.CategoryID] = true
// 	}

// 	var categoryIDs []uint
// 	for id := range categorySet {
// 		categoryIDs = append(categoryIDs, id)
// 	}

// 	categories, err := facade.AcceptanceTest.GetCategoriesByIDs(categoryIDs)
// 	if err != nil {
// 		return err
// 	}

// 	for _, category := range categories {
// 		fmt.Println("Category Name:", category.CategoryName)
// 		if category.Description != nil {
// 			fmt.Println("Description:", *category.Description)
// 		} else {
// 			fmt.Println("Description: (none)")
// 		}
// 		fmt.Println("-----")
// 	}

// 	return nil
// }
