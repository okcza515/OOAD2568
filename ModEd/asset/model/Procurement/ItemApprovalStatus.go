package model

import (
	"time"

	"github.com/google/uuid"
)

type ItemApprovalStatus struct {
	ItemApprovalStatusID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	Status               string    `gorm:"type:varchar(50);not null"`
	Description          string    `gorm:"type:text"`
	ItemApprovalDate     time.Time `gorm:"type:time;not null"`
}
