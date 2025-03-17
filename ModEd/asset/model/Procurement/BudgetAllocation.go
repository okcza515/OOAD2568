package model

import (
	"github.com/google/uuid"
)

type BudgetAllocation struct {
	BudgetAllocationID uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProcurementID      uuid.UUID `gorm:"foreignKey:ProcurementID"`
	TotalBudget        float64   `gorm:"not null"`
	AllocatedAmount    float64   `gorm:"not null"`
	RemainingBalance   float64   `gorm:"not null"`
}
