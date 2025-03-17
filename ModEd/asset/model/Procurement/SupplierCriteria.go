package model

import (
	"github.com/google/uuid"
)

type SupplierCriteria struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	CriteriaName string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text"`
}
