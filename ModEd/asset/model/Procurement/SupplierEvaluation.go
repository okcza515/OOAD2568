package model

import (
	"github.com/google/uuid"
)

type SupplierEvaluation struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	SupplierID uuid.UUID `gorm:"type:uuid;not null"`
	CriteriaID uuid.UUID `gorm:"type:uuid;not null"`
	Score      float64   `gorm:"not null"`
}