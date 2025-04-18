// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"

	"gorm.io/gorm"
)

type BudgetAllocationController struct {
	db *gorm.DB
}

func (c *BudgetAllocationController) AllocateBudget(allocation *model.BudgetAllocation) error {
	return c.db.Create(allocation).Error
}

func (c *BudgetAllocationController) UpdateBudget(id uint, newAmount float64) error {
	return c.db.Model(&model.BudgetAllocation{}).Where("id = ?", id).Update("amount", newAmount).Error
}
