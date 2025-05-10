package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

// Helper to Select Instrument Request
func SelectInstrumentRequest(facade *controller.ProcurementControllerFacade) (*model.InstrumentRequest, error) {
	requests, err := facade.RequestedItem.ListAllInstrumentRequests()
	if err != nil || len(*requests) == 0 {
		fmt.Println("No available requests to edit.")
		util.PressEnterToContinue()
		return nil, err
	}

	fmt.Println("\nSelect an Instrument Request ID to edit:")
	for _, req := range *requests {
		fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n",
			req.InstrumentRequestID, req.DepartmentID, req.Status)
	}

	requestID := util.GetUintInput("Enter Instrument Request ID: ")
	request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
	if err != nil {
		fmt.Println("Failed to retrieve request details:", err)
		util.PressEnterToContinue()
		return nil, err
	}
	return request, nil
}
