package handler

import (
	"ModEd/asset/controller"
	util "ModEd/asset/util"
	"fmt"
)

func ApprovalHandler(facade *controller.ProcurementControllerFacade) {
	util.ClearScreen()
	fmt.Println("Select Approval Type:")
	fmt.Println("  1:\tBudget Approval")
	fmt.Println("  2:\tProcurement Approval")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println()

	input := util.GetCommandInput()

	switch input {
	case "1":
		printApprovalOption(&facade.BudgetApproval)
	case "2":
		printApprovalOption(&facade.Procurement)
	case "back":
		return
	default:
		fmt.Println("Invalid option, returning to menu...")
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
		fmt.Println("Available Budget Approvals:")
		for _, a := range approvals {
			fmt.Printf("  ID: %d | Status: %s\n", a.InstrumentRequestID, a.Status)
		}
		// case *controller.ProcurementApprovalController:
		// 	approvals, err := o.ListAllApprovals()
		// 	if err != nil {
		// 		fmt.Println("Failed to fetch procurement approvals:", err)
		// 		return
		// 	}
		// 	if len(approvals) == 0 {
		// 		fmt.Println("No procurement approvals found.")
		// 		return
		// 	}
		// 	fmt.Println("Available Procurement Approvals:")
		// 	for _, a := range approvals {
		// 		fmt.Printf("  ID: %d | Status: %s\n", a.ProcurementID, a.Status)
		// 	}
		// default:
		// 	fmt.Println("Unknown approval type.")
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
			observer.OnApproved(id)
			fmt.Println("Approved successfully.")
			WaitForEnter()
		case "2":
			id := util.GetUintInput("Enter ID to Reject: ")
			observer.OnRejected(id)
			fmt.Println("Rejected successfully.")
			WaitForEnter()
		case "back":
			return
		default:
			fmt.Println("Invalid command!")
			WaitForEnter()
		}
	}
}
