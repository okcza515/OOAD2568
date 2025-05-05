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
			fmt.Println("Show Quotations by TOR ID")
			ListAllTORs(facade)
			ShowQuotationsByTORID(facade.GetDB())
			WaitForEnter()
		case "3":
			fmt.Println("Quotation Selection")
			ListAllTORs(facade)
			SelectQuotation(facade.GetDB())
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
	fmt.Println("  2:\tShow Quotations by TOR ID")
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
		err := db.Omit("TOR", "Supplier", "Details").Create(&q).Error
		if err != nil {
			return fmt.Errorf("failed to insert quotation ID %d: %v", q.QuotationID, err)
		}
	}

	fmt.Println("Quotations imported successfully.")
	return nil
}

func ShowQuotationsByTORID(db *gorm.DB) (uint, []model.Quotation, error) {
	torID := util.GetUintInput("Enter TOR ID: ")

	var quotations []model.Quotation
	if err := db.Where("tor_id = ?", torID).Find(&quotations).Error; err != nil {
		fmt.Println("Failed to retrieve quotations:", err)
		return 0, nil, err
	}

	if len(quotations) == 0 {
		fmt.Println("No quotations found for this TOR ID.")
		return 0, nil, nil
	}

	fmt.Printf("Quotations for TOR ID %d:\n", torID)
	for _, q := range quotations {
		fmt.Printf("  QuotationID: %d | SupplierID: %d | Status: %s | Total Price: %.2f\n",
			q.QuotationID, q.SupplierID, q.Status, q.TotalOfferedPrice)
	}

	return torID, quotations, nil
}

func SelectQuotation(db *gorm.DB) {
	torID, quotations, err := ShowQuotationsByTORID(db)
	if err != nil || len(quotations) == 0 {
		return
	}

	selectedID := util.GetUintInput("Enter Quotation ID to APPROVE: ")

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, q := range quotations {
			newStatus := "rejected"
			if q.QuotationID == selectedID {
				newStatus = "approved"
			}
			if err := tx.Model(&model.Quotation{}).
				Where("quotation_id = ?", q.QuotationID).
				Update("status", newStatus).Error; err != nil {
				return err
			}
		}

		return tx.Model(&model.TOR{}).
			Where("tor_id = ?", torID).
			Update("status", "selected").Error
	})

	if err != nil {
		fmt.Println("Failed to select quotation:", err)
	} else {
		fmt.Println("Quotation selected and TOR updated successfully.")
	}
}