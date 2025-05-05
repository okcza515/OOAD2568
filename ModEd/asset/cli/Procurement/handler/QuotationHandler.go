package handler

import (
	procurement "ModEd/asset/controller"
	model "ModEd/asset/model"
	util "ModEd/asset/util"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func QuotationHandler(facade *procurement.ProcurementControllerFacade) {

	inputBuffer := ""

	for inputBuffer != "back" {
		util.ClearScreen()
		printQuotationSupplierOptions()
		inputBuffer = util.GetCommandInput()

		switch inputBuffer {
		case "1":
			// fmt.Println("Import Quotations")
			// QID := util.GetUintInput("Enter Quotation ID: ")
			// Quotation := &model.Quotation{
			// 	QuotationID: QID,
			// 	Status:   model.QuotationStatusPending,
			// }
			// jsonPath := "path/to/Quotation.json"
			// if err := ImportQuotationByID(db, jsonPath, QID); err != nil {
			// 	fmt.Println("Error:", err)
			// }

			WaitForEnter()
		case "2":
			fmt.Println("List by TOR ID")

			WaitForEnter()
		case "3":
			fmt.Println("Quotation Selection")

			WaitForEnter()
		}

		util.ClearScreen()
	}

	util.ClearScreen()
}

func printQuotationSupplierOptions() {
	fmt.Println(":/Procurement/Main")
	fmt.Println()
	fmt.Println("--Quotation Functions--")
	fmt.Println("  1:\tImport Quotations")
	fmt.Println("  2:\tList by TOR ID")
	fmt.Println("  3:\tQuotation Selection")
	fmt.Println("  back:\tBack to main menu (or Ctrl+C to exit")
	fmt.Println()
}

func ImportQuotationByID(db *gorm.DB, jsonPath string, targetID uint) error {
	file, err := os.Open(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var quotations []model.Quotation
	if err := json.NewDecoder(file).Decode(&quotations); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	var found *model.Quotation
	for _, q := range quotations {
		if q.QuotationID == targetID {
			found = &q
			break
		}
	}

	if found == nil {
		return fmt.Errorf("quotation with ID %d not found", targetID)
	}

	found.Status = model.QuotationStatusPending

	if err := db.Create(found).Error; err != nil {
		return fmt.Errorf("failed to insert quotation: %w", err)
	}

	fmt.Println("Quotation imported successfully.")
	return nil
}
