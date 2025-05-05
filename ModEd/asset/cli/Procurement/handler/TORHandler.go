package handler

import (
	procurement "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"fmt"
	"time"
)

func TORHandler(facade *procurement.ProcurementControllerFacade) {
	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printTOROptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			fmt.Println("Create TOR")
			fmt.Println("List All Instrument Requests")
			requests, err := facade.RequestedItem.ListAllInstrumentRequests()
			if !showApprovedRequests(requests, err) {
				break
			}
		
			instrumentRequestID := util.GetUintInput("Enter Instrument Request ID to create TOR for: ")
		
			torID := util.GetUintInput("Enter TOR ID: ")
			scope := util.GetStringInput("Enter TOR Scope: ")
			deliverables := util.GetStringInput("Enter TOR Deliverables: ")
		
			tor := &model.TOR{
				TORID:               torID,
				InstrumentRequestID: instrumentRequestID,
				Scope:               scope,
				Deliverables:        deliverables,
				CreatedAt:           time.Now(),
			}
		
			err = facade.TOR.CreateTOR(tor)
			if err != nil {
				fmt.Println("Failed to create TOR:", err)
			} else {
				fmt.Println("TOR created successfully.")
			}
		
			WaitForEnter()
		
		case "2":
			fmt.Println("List All TORs")
			ListAllTORs(facade)
			WaitForEnter()
		case "3":
			fmt.Println("View TOR by ID")
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
			fmt.Printf("Created At: %s\n", createdAt)
		
			WaitForEnter()
		
		case "4":
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
		}

		util.ClearScreen()
	}

	util.ClearScreen()

}

func printTOROptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--TOR Functions--")
	fmt.Println("  1:\tCreate TOR")
	fmt.Println("  2:\tList All TORs")
	fmt.Println("  3:\tView TOR by ID")
	fmt.Println("  4:\tDelete TOR")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit")
	fmt.Println()
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
		fmt.Printf("  TOR ID: %d | Instrument Request ID: %d | Scope: %s | Deliverables: %s | Status: %s | Created At: %s\n",
			tor.TORID, tor.InstrumentRequestID, tor.Scope, tor.Deliverables, tor.Status,createdAt)
	}

}
