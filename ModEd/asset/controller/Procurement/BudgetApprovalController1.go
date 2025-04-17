// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"
	"time"

	"gorm.io/gorm"
)

type BudgetApprovalController struct {
	db *gorm.DB
}

// Create a new budget approval entry
func (c *BudgetApprovalController) CreateBudgetApproval(body *model.BudgetApproval) error {
	result := c.db.Create(body)
	return result.Error
}

// Update the status of a specific budget approval
func (c *BudgetApprovalController) UpdateApprovalStatus(id uint, newStatus string) error {
	result := c.db.Model(&model.BudgetApproval{}).Where("id = ?", id).Update("status", newStatus)
	return result.Error
}

// Get a specific budget approval by ID
func (c *BudgetApprovalController) GetByID(id uint) (*model.BudgetApproval, error) {
	approval := new(model.BudgetApproval)
	result := c.db.Preload("Approver").First(approval, "id = ?", id)
	return approval, result.Error
}

// Get a list of all budget approvals
func (c *BudgetApprovalController) ListAll() ([]model.BudgetApproval, error) {
	var approvals []model.BudgetApproval
	result := c.db.Preload("Approver").Find(&approvals)
	return approvals, result.Error
}

// Delete budget approval by ID (soft delete)
func (c *BudgetApprovalController) Delete(id uint) error {
	now := time.Now()
	result := c.db.Model(&model.BudgetApproval{}).Where("id = ?", id).Update("deleted_at", now)
	return result.Error
}

// CreateBudgetApprovalController - initializes controller with DB connection
func CreateBudgetApprovalController(db *gorm.DB) *BudgetApprovalController {
	return &BudgetApprovalController{
		db: db,
	}
}
