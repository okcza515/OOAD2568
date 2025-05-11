package handler

import (
	"ModEd/asset/controller"
	procurement "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"ModEd/core/validation"
	"fmt"
	"time"
)

func ProcurementHandler(facade *controller.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printProcurementOptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create TOR and Procurement")
			fmt.Println("List All Instrument Requests")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showApprovedRequests(requests, err) {
				break
			}

			instrumentRequestID := util.GetUintInput("Enter Instrument Request ID to create TOR for: ")
			scope := util.GetStringInput("Enter TOR Scope: ")
			deliverables := util.GetStringInput("Enter TOR Deliverables: ")
			timeline := util.GetStringInput("Enter TOR Timeline: ")
			committee := util.GetStringInput("Enter TOR Committee: ")

			tor := controller.NewTORBuilder().
				WithInstrumentRequestID(instrumentRequestID).
				WithScope(scope).
				WithDeliverables(deliverables).
				WithTimeline(timeline).
				WithCommittee(committee).
				WithCreatedAt(time.Now()).
				Build()

			validator := validation.NewModelValidator()
			if err := validator.ModelValidate(tor); err != nil {
				fmt.Println("Validation failed:", err)
				WaitForEnter()
				break
			}

			err = facade.TOR.CreateTOR(tor)
			if err != nil {
				fmt.Println("Failed to create TOR:", err)
			} else {
				fmt.Println("TOR created successfully with ID:", tor.TORID)
			}
			newProcurement := controller.NewProcurementBuilder().
				WithTOR(tor).
				WithStatus(model.ProcurementStatusPending).
				WithCreatedAt(time.Now()).
				Build()
			err = facade.Procurement.CreateProcurement(newProcurement)
			if err != nil {
				fmt.Println("Failed to create Procurement:", err)
			} else {
				fmt.Println("Procurement created with ID:", newProcurement.ProcurementID)
			}

			err = facade.RequestedItem.MarkAsUsed(instrumentRequestID)
			if err != nil {
				fmt.Println("Warning: Failed to mark Instrument Request as used:", err)
			}
			WaitForEnter()
		case "2":
			fmt.Println("List All Procurements")
			ListAllProcurements(facade)
			WaitForEnter()
		case "3":
			fmt.Println("List All TORs")
			ListAllTORs(facade)
			WaitForEnter()
		case "4":
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
		case "5":
			fmt.Println("View TOR Detail by ID")
			ListAllTORs(facade)
			id := util.GetUintInput("Enter TOR ID: ")
			tor, err := facade.TOR.GetTORByID(id)
			if err != nil {
				fmt.Printf("Failed to retrieve TOR with ID %d: %v\n", id, err)
				WaitForEnter()
				break
			}

			createdAt := "-"
			if !tor.CreatedAt.IsZero() {
				createdAt = tor.CreatedAt.Format("2006-01-02 15:04:05")
			}

			fmt.Printf("TOR ID: %d\n", tor.TORID)
			fmt.Printf("Instrument Request ID: %d\n", tor.InstrumentRequestID)
			fmt.Printf("Scope: %s\n", tor.Scope)
			fmt.Printf("Deliverables: %s\n", tor.Deliverables)
			fmt.Printf("Status: %s\n", tor.Status)
			fmt.Printf("Committee: %s\n", tor.Committee)
			fmt.Printf("Created At: %s\n", createdAt)

			WaitForEnter()
		case "6":
			fmt.Println("Delete Procurement")
			ListAllProcurements(facade)
			id := util.GetUintInput("Enter procurement ID to delete: ")
			err := facade.Procurement.DeleteProcurement(id)
			if err != nil {
				fmt.Printf("Failed to delete procurement with ID %d: %v\n", id, err)
				return
			}
			fmt.Printf("Procurement with ID %d deleted successfully.\n", id)
			WaitForEnter()
		case "7":
			fmt.Println("Delete TOR")
			ListAllTORs(facade)
			id := util.GetUintInput("Enter TOR ID to delete: ")
			err := facade.TOR.DeleteTOR(id)
			if err != nil {
				fmt.Printf("Failed to delete TOR with ID %d: %v\n", id, err)
				return
			}
			fmt.Printf("TOR with ID %d deleted successfully.\n", id)
			WaitForEnter()
		case "8":
			QuotationHandler(facade)
		}
	}

	util.ClearScreen()
}

func printProcurementOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--Procurement Functions--")
	fmt.Println("  1:\tCreate TOR and Procurement")
	fmt.Println("  2:\tList All Procurements")
	fmt.Println("  3:\tList All TORs")
	fmt.Println("  4:\tView Procurement Detail by ID")
	fmt.Println("  5:\tView TOR Detail by ID")
	fmt.Println("  6:\tDelete Procurement")
	fmt.Println("  7:\tDelete TOR")
	fmt.Println("  8:\tQuotation Page")
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
			createdAt := "-"
			if !procurement.CreatedAt.IsZero() {
				createdAt = procurement.CreatedAt.Format("2006-01-02 15:04:05")
			}
			fmt.Printf("ID: %d | ApproverID: %s | Status: %s | CreatedAt: %s\n",
				procurement.ProcurementID, approverID, procurement.Status, createdAt)
		}
	}
}

func showApprovedRequests(requests *[]model.InstrumentRequest, err error) bool {
	if err != nil {
		fmt.Println("Failed to retrieve requests:", err)
		return false
	}

	approved := []model.InstrumentRequest{}
	for _, request := range *requests {
		if request.Status == "approved" {
			approved = append(approved, request)
		}
	}

	if len(approved) == 0 {
		fmt.Println("No approved instrument requests available.")
		return false
	}

	fmt.Println("Approved Instrument Requests:")
	for _, request := range approved {
		fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n", request.InstrumentRequestID, request.DepartmentID, request.Status)
	}

	return true
}

func ListAllTORs(facade *procurement.ProcurementControllerFacade) {
	tors, err := facade.TOR.GetAllTORs()
	if err != nil {
		fmt.Println("Failed to retrieve TORs:", err)
		return
	}

	if len(tors) == 0 {
		fmt.Println("No TOR records found.")
		return
	}

	fmt.Println("TOR List:")
	for _, tor := range tors {
		createdAt := tor.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Printf("  TOR ID: %d | Instrument Request ID: %d | Status: %s | Created At: %s\n",
			tor.TORID, tor.InstrumentRequestID, tor.Status, createdAt)
	}

}
