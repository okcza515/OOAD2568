package handler

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"fmt"
)

func ApprovalHandler(approval controller.ApprovalObserver) {
	for {
		util.ClearScreen()
		printApprovalOptions()

		cmd := util.GetCommandInput()

		switch cmd {
		case "1":
			id := GetUintInput("Enter ID to Approve: ")
			approval.OnApproved(id)
			fmt.Println("Approved successfully.")
			WaitForEnter()
		case "2":
			id := GetUintInput("Enter ID to Reject: ")
			approval.OnRejected(id)
			fmt.Println("Rejected successfully.")
			WaitForEnter()
		case "back":
			return
		default:
			fmt.Println("Invalid command!")
		}
	}
}

func printApprovalOptions() {
	fmt.Println(":/Approval Menu")
	fmt.Println()
	fmt.Println("--Approval Function--")
	fmt.Println("  1:\tApprove by ID")
	fmt.Println("  2:\tReject by ID")
	fmt.Println("  back:\tBack to main menu")
	fmt.Println()
}
