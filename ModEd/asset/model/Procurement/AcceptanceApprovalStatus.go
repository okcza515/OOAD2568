package model

import (
	"github.com/google/uuid"
)

type AcceptanceApprovalStatus struct {
	AcceptanceApprovalStatusID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	Status                     string    `gorm:"type:varchar(50);not null"`
	Description                string    `gorm:"type:text"`
}
