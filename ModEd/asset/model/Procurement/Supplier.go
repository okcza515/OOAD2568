package model

import (
	"github.com/google/uuid"
)

type Supplier struct {
	SupplierID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	ContactInfo string    `gorm:"type:text"`
	Performance float64   `gorm:"not null"`
}
