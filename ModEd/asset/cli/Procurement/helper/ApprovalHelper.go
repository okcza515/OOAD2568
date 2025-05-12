package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
)

func HandleApprovalOption(observer controller.ApprovalObserver) {
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
			util.PressEnterToContinue()
		case "2":
			id := util.GetUintInput("Enter ID to Reject: ")
			approverID := util.GetUintInput("Enter Approver ID: ")
			err := observer.OnRejected(id, approverID)
			if err != nil {
				fmt.Println("Rejection failed:", err)
			} else {
				fmt.Println("Rejected successfully.")
			}
			util.PressEnterToContinue()
		case "back":
			return
		default:
			fmt.Println("Invalid command!")
			util.PressEnterToContinue()
		}
	}
}

func printApprovalList(observer controller.ApprovalObserver) {
	switch o := observer.(type) {
	case *controller.BudgetApprovalController:
		approvals, err := o.ListAllPendingApprovals()
		if err != nil {
			fmt.Println("Failed to fetch budget approvals:", err)
			return
		}
		if len(approvals) == 0 {
			fmt.Println("No budget approvals found.")
			return
		}
		fmt.Println("Available Budget Approvals:")
		for _, a := range approvals {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | RequestID: %d | Status: %s | Approver ID: %s\n",
				a.BudgetApprovalID, a.InstrumentRequestID, a.Status, approverID)
		}

	case *controller.ProcurementController:
		approvals, err := o.ListAllPendingProcurement()
		if err != nil {
			fmt.Println("Failed to fetch procurement approvals:", err)
			return
		}
		if len(*approvals) == 0 {
			fmt.Println("No procurement approvals found.")
			return
		}
		fmt.Println("Available Procurement Approvals:")
		for _, a := range *approvals {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | ProcurementID: %d | Status: %s | Approver ID: %s\n",
				a.ProcurementID, a.ProcurementID, a.Status, approverID)
		}

	case *controller.AcceptanceApprovalController:
		approvals, err := o.ListAllPendingApprovals()
		if err != nil {
			fmt.Println("Failed to fetch acceptance approvals:", err)
			return
		}
		if len(approvals) == 0 {
			fmt.Println("No acceptance approvals found.")
			return
		}
		fmt.Println("Available Acceptance Approvals:")
		for _, a := range approvals {
			approverID := "waiting"
			if a.ApproverID != nil && *a.ApproverID != 0 {
				approverID = fmt.Sprintf("%d", *a.ApproverID)
			}
			fmt.Printf("  ApprovalID: %d | ProcurementID: %d | Status: %s | Approver ID: %s\n",
				a.AcceptanceApprovalID, a.ProcurementID, a.Status, approverID)
		}
	}
}
