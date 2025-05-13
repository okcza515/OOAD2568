// MEP-1014
package helper

import (
	"ModEd/asset/model"
	"ModEd/asset/util"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/gorm"
)

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
	return nil
}

func ImportQuotationDetailsFromJSON(db *gorm.DB, filename string) error {
	var details []model.QuotationDetail

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&details); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	for _, d := range details {
		err := db.Create(&d).Error
		if err != nil {
			return fmt.Errorf("failed to insert detail (QuotationID %d, InstrumentLabel: %s): %v", d.QuotationID, d.InstrumentLabel, err)
		}
	}
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
		return torID, nil, nil
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

	selectedID := util.GetUintInput("Enter Quotation ID to Approve: ")

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
