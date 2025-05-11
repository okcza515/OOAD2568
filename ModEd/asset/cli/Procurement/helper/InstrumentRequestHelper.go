package helper

import (
	"ModEd/asset/controller"
	"ModEd/asset/model"
	"ModEd/asset/util"
	"ModEd/core/validation"
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

func ShowAllInstrumentRequests(facade *controller.ProcurementControllerFacade) {
	requests, err := facade.RequestedItem.ListAllInstrumentRequests()
	if err != nil {
		fmt.Println("Failed to retrieve instrument requests:", err)
		util.PressEnterToContinue()
		return
	}
	if len(*requests) == 0 {
		fmt.Println("No available instrument requests found.")
		util.PressEnterToContinue()
		return
	}

	DisplayRequestList(*requests)
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

func AddInstrumentsLoopToRequest(facade *controller.ProcurementControllerFacade, requestID uint) {
	for {
		fmt.Println("\n--- Add Instrument ---")

		label := util.GetStringInput("Enter Instrument Label: ")
		desc := util.GetStringInput("Enter Description: ")
		categoryID := util.GetUintInput("Enter Category ID: ")
		estimatedPrice := util.GetFloatInput("Enter Estimated Price: ")
		quantity := util.GetUintInput("Enter Quantity: ")

		detail := controller.NewInstrumentDetailBuilder().
			WithLabel(label).
			WithDescription(desc).
			WithCategoryID(categoryID).
			WithEstimatedPrice(estimatedPrice).
			WithQuantity(int(quantity)).
			WithRequestID(requestID).
			Build()

		validator := validation.NewModelValidator()
		if err := validator.ModelValidate(detail); err != nil {
			fmt.Println("Validation failed:", err)
			util.PressEnterToContinue()
			return
		}

		err := facade.RequestedItem.AddInstrumentToRequest(requestID, detail)
		if err != nil {
			fmt.Println("Failed to add instrument:", err)
		} else {
			fmt.Println("Instrument added to request!")
		}

		addMore := util.GetStringInput("\nAdd another instrument? (y/n): ")
		if addMore != "y" && addMore != "Y" {
			break
		}
	}
}

func DisplayInstrumentList(instruments []model.InstrumentDetail) {
	if len(instruments) == 0 {
		fmt.Println("No instruments available in this request.")
		util.PressEnterToContinue()
		return
	}

	fmt.Println("\n--- Available Instruments ---")
	for _, instrument := range instruments {
		fmt.Printf("  - ID: %d | Label: %s | Qty: %d | Price: %.2f\n",
			instrument.InstrumentDetailID, instrument.InstrumentLabel, instrument.Quantity, instrument.EstimatedPrice)
	}
}

func SelectInstrument(request *model.InstrumentRequest) (*model.InstrumentDetail, error) {
	DisplayInstrumentList(request.Instruments)
	instrumentID := util.GetUintInput("\nEnter Instrument Detail ID to edit: ")

	for _, inst := range request.Instruments {
		if inst.InstrumentDetailID == instrumentID {
			return &inst, nil
		}
	}

	fmt.Println("Instrument Detail not found.")
	util.PressEnterToContinue()
	return nil, fmt.Errorf("instrument not found")
}

func EditInstrumentDetails(selected *model.InstrumentDetail) {
	fmt.Println("\n--- Current Values ---")
	fmt.Printf("Label: %s\n", selected.InstrumentLabel)
	fmt.Printf("Description: %s\n", util.DereferenceString(selected.Description))
	fmt.Printf("Category ID: %d\n", selected.CategoryID)
	fmt.Printf("Quantity: %d\n", selected.Quantity)
	fmt.Printf("Estimated Price: %.2f\n", selected.EstimatedPrice)

	fmt.Println("\n--- Enter New Values (Press Enter to keep current) ---")
	newLabel := util.GetOptionalStringInput("New Label", selected.InstrumentLabel)
	newDesc := util.GetOptionalStringInput("New Description", util.DereferenceString(selected.Description))
	newCategory := util.GetOptionalUintInput("New Category ID", selected.CategoryID)
	newQty := util.GetOptionalUintInput("New Quantity", uint(selected.Quantity))
	newPrice := util.GetOptionalFloatInput("New Estimated Price", selected.EstimatedPrice)

	selected.InstrumentLabel = newLabel
	selected.Description = &newDesc
	selected.CategoryID = newCategory
	selected.Quantity = int(newQty)
	selected.EstimatedPrice = newPrice
}
