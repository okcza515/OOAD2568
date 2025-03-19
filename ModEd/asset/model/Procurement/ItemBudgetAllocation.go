//MEP-1014
package model

import (
	"github.com/google/uuid"
)

type ItemBudgetAllocation struct {
	ItemBudgetAllocationID uuid.UUID `gorm:"type:uuid;primaryKey"`
	TotalBudget            float64   `gorm:"not null"` //TO-DO: link this to the budget in department
	AllocatedAmount        float64   `gorm:"not null"`
	RemainingBalance       float64   `gorm:"not null"`
}
