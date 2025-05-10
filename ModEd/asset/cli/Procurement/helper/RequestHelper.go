package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

func DisplayRequestList(requests []model.InstrumentRequest) {
	if len(requests) == 0 {
		fmt.Println("No available requests.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\n--- Available Instrument Requests ---")
	for _, req := range requests {
		fmt.Printf("  ID: %d | Department ID: %d | Status: %s\n",
			req.InstrumentRequestID, req.DepartmentID, req.Status)
	}
}

func SelectInstrumentRequest(facade *controller.ProcurementControllerFacade) (*model.InstrumentRequest, error) {
	requests, err := facade.RequestedItem.ListAllInstrumentRequests()
	if err != nil {
		fmt.Println("Failed to list requests:", err)
		util.PressEnterToContinue()
		return nil, err
	}

	DisplayRequestList(*requests)
	requestID := util.GetUintInput("Enter Instrument Request ID: ")

	request, err := facade.RequestedItem.GetInstrumentRequestWithDetails(requestID)
	if err != nil {
		fmt.Println("Failed to retrieve request details:", err)
		util.PressEnterToContinue()
		return nil, err
	}
	return request, nil
}
