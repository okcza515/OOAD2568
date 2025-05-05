package handler

import (
	procurement "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"fmt"
	"time"
)

func ProcurementHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printProcurementOptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create Procurement")
			PID := util.GetUintInput("Enter Procurement ID: ")
			AID := util.GetUintPointerInput("Enter Approver ID: ")
			newProcurement := &model.Procurement{
				ProcurementID: PID,
				ApproverID:    AID,
				Status:        model.ProcurementStatusPending,
			}

			err := facade.Procurement.CreateProcurement(newProcurement)
			if err != nil {
				fmt.Println("Failed to create Procurement:", err)
				WaitForEnter()
				break
			}
			fmt.Println("Procurement created with ID:", newProcurement.ProcurementID)
			WaitForEnter()

		case "2":
			fmt.Println("List All Procurements")
			ListAllProcurements(facade)
			WaitForEnter()
		case "3":
			fmt.Println("View Procurement Detail by ID")
			ListAllProcurements(facade)
			id := util.GetUintInput("Enter procurement ID: ")
			procurement, err := facade.Procurement.GetProcurementByID(id)
			if err != nil {
				fmt.Printf("Failed to retrieve procurement with ID %d: %v\n", id, err)
				WaitForEnter()
				ProcurementHandler(facade)
			}
			approverID := "Not Assigned"
			if procurement.ApproverID != nil {
				approverID = fmt.Sprintf("%d", *procurement.ApproverID)
			}
			approvalTime := "-"
			if procurement.ApprovalTime != nil {
				approvalTime = procurement.ApprovalTime.Format("2006-01-02 15:04:05")
			}
			fmt.Println("Procurement Detail:")
			fmt.Printf("ID: %d\n", procurement.ProcurementID)
			fmt.Printf("ApproverID: %s\n", approverID)
			fmt.Printf("Status: %s\n", procurement.Status)
			fmt.Printf("ApprovalTime: %s\n", approvalTime)
			WaitForEnter()
		case "4": //NOT USE, DELETE LATER
			fmt.Println("Update Procurement Status")
			ListAllProcurements(facade)

			id := util.GetUintInput("Enter procurement ID: ")

			if _, err := facade.Procurement.GetProcurementByID(id); err != nil {
				fmt.Printf("Failed to retrieve procurement with ID %d: %v\n", id, err)
				WaitForEnter()
				break
			}

			fmt.Println("Choose new status:")
			fmt.Println("  1: Approve")
			fmt.Println("  2: Reject")
			statusChoice := util.GetCommandInput()

			now := time.Now()
			var updateErr error

			switch statusChoice {
			case "1":
				updateErr = facade.Procurement.Update(id, map[string]any{
					"status":        model.ProcurementStatusApproved,
					"approval_time": &now,
				})
			case "2":
				updateErr = facade.Procurement.Update(id, map[string]any{
					"status":        model.ProcurementStatusRejected,
					"approval_time": &now,
				})
			default:
				fmt.Println("Invalid status choice.")
				WaitForEnter()
			}

			if updateErr != nil {
				fmt.Printf("Failed to update status: %v\n", updateErr)
			} else {
				fmt.Println("Status updated successfully.")
			}
			WaitForEnter()
		case "5":
			fmt.Println("Delete Procurement")
			ListAllProcurements(facade)
			id := util.GetUintInput("Enter procurement ID to delete: ")
			err := facade.Procurement.Delete(id)
			if err != nil {
				fmt.Printf("Failed to delete procurement with ID %d: %v\n", id, err)
				return
			}
			fmt.Printf("Procurement with ID %d deleted successfully.\n", id)
			WaitForEnter()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func printProcurementOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--Procurement Functions--")
	fmt.Println("  1:\tCreate Procurement")
	fmt.Println("  2:\tList All Procurements")
	fmt.Println("  3:\tView Procurement Detail by ID")
	fmt.Println("  4:\tUpdate Procurement Status")
	fmt.Println("  5:\tDelete Procurement")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit ¯\\\\_(ツ)_/¯)")
	fmt.Println()
}

func ListAllProcurements(facade *procurement.ProcurementControllerFacade) {
	procurements, err := facade.Procurement.ListAllProcurement()
	if err != nil {
		fmt.Println("Failed to list procurements:", err)
		return
	} else {

		fmt.Println("Procurement List:")
		for _, procurement := range *procurements {
			approverID := "Not Assigned"
			if procurement.ApproverID != nil {
				approverID = fmt.Sprintf("%d", *procurement.ApproverID)
			}
			approvalTime := "-"
			if procurement.ApprovalTime != nil {
				approvalTime = procurement.ApprovalTime.Format("2006-01-02 15:04:05")
			}
			fmt.Printf("ID: %d | ApproverID: %s | Status: %s | ApprovalTime: %s\n", procurement.ProcurementID, approverID, procurement.Status, approvalTime)
		}
	}
}
