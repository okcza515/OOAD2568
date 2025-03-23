// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type ItemBudgetAllocation struct {
	ItemBudgetAllocationID uint           `gorm:"primaryKey"`
	TotalBudget            float64        `gorm:"not null"` //TO-DO: link this to the budget in department
	AllocatedAmount        float64        `gorm:"not null"`
	RemainingBalance       float64        `gorm:"not null"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}
