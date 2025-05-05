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
			fmt.Println("Import Quotations")
			filename := util.GetStringInput("Enter path to the JSON file (e.g., data/quotations.json): ")
		
			err := ImportQuotationsFromJSON(facade.GetDB(), filename)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Import successful.")
			}
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

func ImportQuotationsFromJSON(db *gorm.DB, filename string) error {
	var quotations []model.Quotation

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&quotations); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	for _, q := range quotations {
		// Avoid inserting associations if not needed (just flat data)
		err := db.Omit("TOR", "Supplier", "Details").Create(&q).Error
		if err != nil {
			return fmt.Errorf("failed to insert quotation ID %d: %v", q.QuotationID, err)
		}
	}

	fmt.Println("Quotations imported successfully.")
	return nil
}