// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"fmt"

	"gorm.io/gorm"
)

type AcceptedInstrumentController struct {
	db *gorm.DB
}

func NewAcceptedInstrumentController(db *gorm.DB) *AcceptedInstrumentController {
	return &AcceptedInstrumentController{
		db: db,
	}
}

func (c *AcceptedInstrumentController) ListAllInstruments() ([]model.Instrument, error) {
	var instruments []model.Instrument
	if err := c.db.Find(&instruments).Error; err != nil {
		return nil, err
	}
	return instruments, nil
}

func (c *AcceptedInstrumentController) RetrieveByID(id uint) (*model.Instrument, error) {
	var instrument model.Instrument
	if err := c.db.Where("id = ?", id).First(&instrument).Error; err != nil {
		return nil, err
	}
	return &instrument, nil
}

func (c *AcceptedInstrumentController) CreateInstrumentsFromAcceptance(acceptanceID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		var acceptance model.AcceptanceApproval
		if err := tx.Preload("Procurement").First(&acceptance, acceptanceID).Error; err != nil {
			return fmt.Errorf("acceptance not found: %w", err)
		}

		if acceptance.Status != model.AcceptanceStatusApproved {
			return fmt.Errorf("acceptance is not approved")
		}

		if acceptance.InstrumentsCreated {
			return fmt.Errorf("instruments for this acceptance are already created")
		}

		if acceptance.ApprovalTime == nil {
			return fmt.Errorf("approval time is not set for this acceptance")
		}

		budgetYear := acceptance.ApprovalTime.Year()

		var quotationDetails []model.QuotationDetail
		if err := tx.Where("quotation_id IN (SELECT quotation_id FROM quotations WHERE tor_id = ?)", acceptance.Procurement.TORID).
			Find(&quotationDetails).Error; err != nil {
			return fmt.Errorf("failed to fetch quotation details: %w", err)
		}

		var instruments []model.Instrument
		for _, detail := range quotationDetails {
			for i := 0; i < detail.Quantity; i++ {
				instrument := model.Instrument{
					InstrumentLabel:  detail.InstrumentLabel,
					InstrumentCode:   fmt.Sprintf("%s-%d-%d", detail.InstrumentLabel, detail.QuotationDetailID, i+1),
					InstrumentStatus: model.INS_AVAILABLE,
					RoomID:           "Unassigned",
					Location:         "Unassigned",
					CategoryID:       detail.CategoryID,
					Cost:             detail.OfferedPrice,
					BudgetYear:       budgetYear,
					Category:         detail.Category,
				}
				instruments = append(instruments, instrument)
			}
		}

		if err := tx.Create(&instruments).Error; err != nil {
			return fmt.Errorf("failed to create instruments: %w", err)
		}

		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", acceptanceID).
			Update("InstrumentsCreated", true).Error; err != nil {
			return fmt.Errorf("failed to mark acceptance as created: %w", err)
		}

		fmt.Printf("%d Instruments successfully created for Acceptance ID %d\n", len(instruments), acceptanceID)
		return nil
	})
}

func (c *AcceptanceApprovalController) UpdateStatusToImported(acceptanceID uint) error {
	result := c.db.Model(&model.AcceptanceApproval{}).
		Where("acceptance_approval_id = ?", acceptanceID).
		Update("status", model.AcceptanceStatusImported)

	if result.Error != nil {
		return fmt.Errorf("failed to update status to Imported: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with AcceptanceApprovalID %d", acceptanceID)
	}

	fmt.Printf("Status updated to 'Imported' for Acceptance ID %d\n", acceptanceID)
	return nil
}
