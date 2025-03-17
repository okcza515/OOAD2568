package model

import (
	"github.com/google/uuid"
)

type SupplierApproval struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	SupplierID uuid.UUID `gorm:"type:uuid;not null"`
	ApproverID uuid.UUID `gorm:"type:uuid;not null"`
	Status     string    `gorm:"type:varchar(50);not null"`
	Remarks    string    `gorm:"type:text"`
}
