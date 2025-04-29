package handler

import (
	"ModEd/asset/controller"
	"ModEd/asset/util"
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
        printApprovalOption(&facade.ProcurementApproval)
    case "back":
        return
    default:
        fmt.Println("Invalid option, returning to menu...")
    }
}

func printApprovalOption(observer controller.ApprovalObserver) {
    for {
        util.ClearScreen()
        fmt.Println(":/Approval Menu")
        fmt.Println("  1:\tApprove by ID")
        fmt.Println("  2:\tReject by ID")
        fmt.Println("  back:\tBack to previous menu")
        fmt.Println()

        cmd := util.GetCommandInput()

        switch cmd {
        case "1":
            id := GetUintInput("Enter ID to Approve: ")
            observer.OnApproved(id)
            fmt.Println("Approved successfully.")
            WaitForEnter()
        case "2":
            id := GetUintInput("Enter ID to Reject: ")
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

