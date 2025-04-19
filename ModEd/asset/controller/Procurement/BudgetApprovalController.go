// MEP-1014
package procurement

import (
	"time"

	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type BudgetApprovalController struct {
	db *gorm.DB
}

func (c *BudgetApprovalController) CreateBudgetRequest(req *model.BudgetApproval) error {
	return c.db.Create(req).Error
}

func (c *BudgetApprovalController) BudgetApprove(id uint, status model.BudgetApprovalStatus) error {
	return c.db.Model(&model.BudgetApproval{}).Where("budget_approval_id = ?", id).Update("status", status).Error
}

func (c *BudgetApprovalController) ShowBudgetRequestList(itemRequestID uint) (*[]model.BudgetApproval, error) {
	var approvals []model.BudgetApproval
	err := c.db.Where("item_request_id = ?", itemRequestID).Find(&approvals).Error
	return &approvals, err
}

func (c *BudgetApprovalController) ShowBudgetRequestStatus(id uint) (*model.BudgetApproval, error) {
	approval := new(model.BudgetApproval)
	err := c.db.First(&approval, "budget_approval_id = ?", id).Error
	return approval, err
}

func (c *BudgetApprovalController) DeleteBudgetRequest(id uint) error {
	return c.db.Model(&model.BudgetApproval{}).Where("budget_approval_id = ?", id).Update("deleted_at", time.Now()).Error
}
