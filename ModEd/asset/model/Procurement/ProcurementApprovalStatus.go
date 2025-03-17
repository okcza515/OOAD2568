package model

import (
	"github.com/google/uuid"
)

type ProcurementApprovalStatus struct {
	ProcurementApprovalStatusID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	Status                     string    `gorm:"type:varchar(50);not null"`
	Description                string    `gorm:"type:text"`
}