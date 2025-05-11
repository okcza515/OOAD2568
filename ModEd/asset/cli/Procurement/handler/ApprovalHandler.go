package handler

import (
	"ModEd/asset/controller"
	util "ModEd/asset/util"
	model "ModEd/asset/model"
	"fmt"
)

func ApprovalHandler(facade *controller.ProcurementControllerFacade) {
	util.ClearScreen()
	fmt.Println("Select Approval Type:")
	fmt.Println("  1:\tBudget Approval")
	fmt.Println("  2:\tProcurement Approval")
	fmt.Println("  3:\tAcceptance Approval")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println()

	input := util.GetCommandInput()

	switch input {
	case "1":
		printApprovalOption(&facade.BudgetApproval)
	case "2":
		printApprovalOption(&facade.Procurement)
	case "3":
		printApprovalOption(&facade.Acceptance)
	case "back":
		return
	default:
		fmt.Println("Invalid option, returning to menu...")
	}
}

func printApprovalOption(observer controller.ApprovalObserver) {
	for {
		util.ClearScreen()
		printApprovalList(observer)
		fmt.Println(":/Approval Menu")
		fmt.Println("  1:\tApprove by ID")
		fmt.Println("  2:\tReject by ID")
		fmt.Println("  back:\tBack to previous menu")
		fmt.Println()

		cmd := util.GetCommandInput()

		switch cmd {

		case "1":
			id := util.GetUintInput("Enter ID to Approve: ")
			approverID := util.GetUintInput("Enter Your Instructor ID (Approver): ")
			err := observer.OnApproved(id, approverID)
			if err != nil {
				fmt.Println("Failed to approve:", err)
			} else {
				fmt.Println("Approved successfully.")
			}
			WaitForEnter()
		case "2":
			id := util.GetUintInput("Enter ID to Reject: ")
			approverID := util.GetUintInput("Enter Approver ID: ")
			err := observer.OnRejected(id, approverID)
			if err != nil {
				fmt.Println("Rejection failed:", err)
			} else {
				fmt.Println("Rejected successfully.")
			}
			WaitForEnter()
		case "back":
			return
		default:
			fmt.Println("Invalid command!")
			WaitForEnter()
		}
	}
}

func printApprovalList(observer controller.ApprovalObserver) {
	switch o := observer.(type) {

	case *controller.BudgetApprovalController:
		approvals, err := o.ListAllApprovals()
		if err != nil {
			fmt.Println("Failed to fetch budget approvals:", err)
			return
		}
		if len(approvals) == 0 {
			fmt.Println("No budget approvals found.")
			return
		}
		printFormattedApprovals("Budget Approval", approvals, "InstrumentRequestID")

	case *controller.ProcurementController:
		approvals, err := o.ListSelectedProcurements()
		if err != nil {
			fmt.Println("Failed to fetch procurement approvals:", err)
			return
		}
		if approvals == nil || len(*approvals) == 0 {
			fmt.Println("No procurement approvals found.")
			return
		}
		printFormattedApprovals("Procurement Approval", *approvals, "ProcurementID")

	case *controller.AcceptanceApprovalController:
		approvals, err := o.ListAllApprovals()
		if err != nil {
			fmt.Println("Failed to fetch acceptance approvals:", err)
			return
		}
		if len(approvals) == 0 {
			fmt.Println("No acceptance approvals found.")
			return
		}
		printFormattedApprovals("Acceptance Approval", approvals, "ProcurementID")
	}
}


func printFormattedApprovals(title string, approvals interface{}, idField string) {
	fmt.Printf("Available %s:\n", title)
	
	switch items := approvals.(type) {
	case []model.BudgetApproval:
		for _, a := range items {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | RequestID: %d | Status: %s | Approver ID: %s\n",
				a.BudgetApprovalID, a.InstrumentRequestID, a.Status, approverID)
		}
	case []model.Procurement:
		for _, a := range items {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | ProcurementID: %d | TOR ID: %d | Status: %s | Approver ID: %s\n",
				a.ProcurementID, a.ProcurementID, a.TORID, a.Status, approverID)
		}
	case []model.AcceptanceApproval:
		for _, a := range items {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | ProcurementID: %d | Status: %s | Approver ID: %s\n",
				a.AcceptanceApprovalID, a.ProcurementID, a.Status, approverID)
		}
	}
}
