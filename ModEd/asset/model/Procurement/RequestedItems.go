package model

import (
	"github.com/google/uuid"
)

type RequestedItem struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	ProcurementID uuid.UUID `gorm:"type:uuid;not null"`
	ItemName      string    `gorm:"type:varchar(255);not null"`
	Quantity      int       `gorm:"not null"`
	UnitPrice     float64   `gorm:"not null"`
	TotalPrice    float64   `gorm:"not null"`
	Specification string    `gorm:"type:text"`
}
