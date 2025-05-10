package helper

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"fmt"
)

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
