// MEP-1014
package controller

import (
	model "ModEd/asset/model"

	"gorm.io/gorm"
)

type BudgetAllocationController struct {
	db *gorm.DB
}

func (c *BudgetAllocationController) AllocateBudget(allocation *model.BudgetAllocation) error {
	return c.db.Create(allocation).Error
}

func (c *BudgetAllocationController) UpdateBudget(id uint, newAmount float64) error {
	return c.db.Model(&model.BudgetAllocation{}).Where("budget_allocation_id = ?", id).Update("Amount", newAmount).Error
}

func (c *BudgetAllocationController) UpdateTotalBudget(id uint) error {
	var allocation model.BudgetAllocation

	if err := c.db.First(&allocation, "budget_allocation_id = ?", id).Error; err != nil {
		return err
	}

	newTotal := allocation.Amount * 1
	return c.db.Model(&model.BudgetAllocation{}).
		Where("budget_allocation_id = ?", id).
		Update("TotalbudgetAmount", newTotal).Error
}

func (c *BudgetAllocationController) GetByID(id uint) (*model.BudgetAllocation, error) {
	var alloc model.BudgetAllocation
	err := c.db.
		Preload("InstrumentRequest.Departments").
		Preload("Approver").
		First(&alloc, "budget_allocation_id = ?", id).Error

	return &alloc, err
}
