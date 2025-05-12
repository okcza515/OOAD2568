package controller

import (
	model "ModEd/asset/model"
	"fmt"
	"gorm.io/gorm"
)

type AcceptedInstrumentController struct {
	InstrumentController // 🔄 Extends InstrumentController
}

// ✅ Constructor for AcceptedInstrumentController
func NewAcceptedInstrumentController(db *gorm.DB) *AcceptedInstrumentController {
	return &AcceptedInstrumentController{
		InstrumentController: InstrumentController{
			db: db,
		},
	}
}

// ✅ List all instruments
func (c *AcceptedInstrumentController) ListAllInstruments() ([]model.Instrument, error) {
	var instruments []model.Instrument
	err := c.db.Find(&instruments).Error
	return instruments, err
}

// ✅ Create Instruments from an accepted acceptance request
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

		// Get the Budget Year from the approval time
		budgetYear := acceptance.ApprovalTime.Year()

		// Get Quotation Details linked to this procurement's TOR
		var quotationDetails []model.QuotationDetail
		if err := tx.Where("quotation_id IN (SELECT quotation_id FROM quotations WHERE tor_id = ?)", acceptance.Procurement.TORID).
			Find(&quotationDetails).Error; err != nil {
			return fmt.Errorf("failed to fetch quotation details: %w", err)
		}

		// Prepare a list of instruments to insert
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

		// ✅ Now you can call InsertMany, because it is inherited from InstrumentController
		if err := c.InsertMany(instruments); err != nil {
			return fmt.Errorf("failed to create instruments: %w", err)
		}

		// Mark as Instruments Created
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", acceptanceID).
			Update("InstrumentsCreated", true).Error; err != nil {
			return fmt.Errorf("failed to mark acceptance as created: %w", err)
		}

		fmt.Printf("%d Instruments successfully created for Acceptance ID %d\n", len(instruments), acceptanceID)
		return nil
	})
}
