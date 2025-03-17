package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetAllocation struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProcurementID    uuid.UUID `gorm:"type:uuid;not null"`
	TotalBudget      float64   `gorm:"not null"`
	AllocatedAmount  float64   `gorm:"not null"`
	RemainingBalance float64   `gorm:"not null"`
}
