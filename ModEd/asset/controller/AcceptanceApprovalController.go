// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AcceptanceApprovalController struct {
	db *gorm.DB
}

func (c *AcceptanceApprovalController) CreateAcceptanceRequest(req *model.AcceptanceApproval) error {
	return c.db.Create(req).Error
}

// func (c *AcceptanceApprovalController) ListAllApprovals() ([]model.AcceptanceApproval, error) {
// 	var approvals []model.AcceptanceApproval
// 	err := c.db.
// 		Preload("Procurement").
// 		Preload("Approver").
// 		Find(&approvals).Error
// 	return approvals, err
// }

func (c *AcceptanceApprovalController) ListAllApprovals() ([]model.AcceptanceApproval, error) {
    var approvals []model.AcceptanceApproval
    err := c.db.
        Preload("Procurement").
        Joins("LEFT JOIN procurements ON procurements.procurement_id = acceptance_approvals.procurement_id").
        Where("procurements.procurement_id IS NOT NULL").
        Find(&approvals).Error

    return approvals, err
}


func (c *AcceptanceApprovalController) ShowAcceptanceRequestList(procurementID uint) ([]model.AcceptanceApproval, error) {
    var approvals []model.AcceptanceApproval
    err := c.db.
        Preload("Procurement").
        Preload("Approver").
        Where("procurement_id = ?", procurementID).
        Find(&approvals).Error
    
    if err != nil {
        return nil, err
    }

    for _, approval := range approvals {
        if approval.Procurement == nil {
            fmt.Printf("Missing Procurement for AcceptanceApprovalID: %d\n", approval.AcceptanceApprovalID)
        }
    }

    return approvals, nil
}


func (c *AcceptanceApprovalController) ShowAcceptanceRequestStatus(id uint) (*model.AcceptanceApproval, error) {
	approval := new(model.AcceptanceApproval)
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		First(approval, id).Error
	return approval, err
}

func (c *AcceptanceApprovalController) ShowAcceptanceRequestByStatus(status model.AcceptanceStatus) ([]model.AcceptanceApproval, error) {
	var approvals []model.AcceptanceApproval
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		Where("status = ?", status).
		Find(&approvals).Error
	return approvals, err
}

func (c *AcceptanceApprovalController) DeleteAcceptanceRequest(id uint) error {
	return c.db.Model(&model.AcceptanceApproval{}).
		Where("acceptance_approval_id = ?", id).
		Update("deleted_at", time.Now()).Error
}

func (c *AcceptanceApprovalController) OnApproved(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":        model.AcceptanceStatusApproved,
				"approver_id":   approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.AcceptanceApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Procurement{}).
			Where("procurement_id = ?", approval.ProcurementID).
			Update("status", model.ProcurementStatusApproved).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *AcceptanceApprovalController) OnRejected(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":        model.AcceptanceStatusRejected,
				"approver_id":   approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.AcceptanceApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Procurement{}).
			Where("procurement_id = ?", approval.ProcurementID).
			Update("status", model.ProcurementStatusRejected).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *AcceptanceApprovalController) GetQuotationDetailsByProcurement(procurementID uint) ([]model.QuotationDetail, error) {
	var procurement model.Procurement
	if err := c.db.First(&procurement, procurementID).Error; err != nil {
		return nil, fmt.Errorf("failed to find procurement: %w", err)
	}

	var quotations []model.Quotation
	err := c.db.Preload("Details").
		Where("tor_id = ?", procurement.TORID).
		Find(&quotations).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get quotations: %w", err)
	}

	var details []model.QuotationDetail
	for _, quotation := range quotations {
		details = append(details, quotation.Details...)
	}

	return details, nil
}

func (c *AcceptanceApprovalController) PrintQuotationDetailsByProcurement(procurementID uint) {
    details, err := c.GetQuotationDetailsByProcurement(procurementID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(details) == 0 {
        fmt.Printf("No quotation details found for Procurement ID: %d\n", procurementID)
        return
    }

    fmt.Printf("Quotation Details for Procurement ID: %d\n", procurementID)
    for _, detail := range details {
        fmt.Printf("QuotationDetailID: %d\n", detail.QuotationDetailID)
        fmt.Printf("InstrumentLabel: %s\n", detail.InstrumentLabel)
        if detail.Description != nil {
            fmt.Printf("Description: %s\n", *detail.Description)
        } else {
            fmt.Println("Description: (none)")
        }
        fmt.Printf("CategoryID: %d\n", detail.CategoryID)
        fmt.Printf("Quantity: %d\n", detail.Quantity)
        fmt.Printf("Offered Price: %.2f\n", detail.OfferedPrice)
        fmt.Println("------")
    }
}

func (c *AcceptanceApprovalController) GetQuotationDetailsByAcceptance(acceptanceID uint) ([]model.QuotationDetail, error) {
    var acceptance model.AcceptanceApproval
    err := c.db.Preload("Procurement").First(&acceptance, acceptanceID).Error
    if err != nil {
        return nil, fmt.Errorf("failed to find Acceptance Request with ID %d: %w", acceptanceID, err)
    }

    if acceptance.Procurement == nil {
        return nil, fmt.Errorf("no associated Procurement found for Acceptance Request ID: %d", acceptanceID)
    }

    var quotations []model.Quotation
    err = c.db.Preload("Details").
        Where("tor_id = ?", acceptance.Procurement.TORID).
        Find(&quotations).Error
    if err != nil {
        return nil, fmt.Errorf("failed to get quotations for TOR ID %d: %w", acceptance.Procurement.TORID, err)
    }

    var details []model.QuotationDetail
    for _, quotation := range quotations {
        details = append(details, quotation.Details...)
    }

    return details, nil
}



