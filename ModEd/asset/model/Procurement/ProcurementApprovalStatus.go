//MEP-1014
package model

import (
	"time"

	"github.com/google/uuid"
)

type ProcurementApprovalStatus struct {
	ProcurementApprovalStatusID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	Status                      string    `gorm:"type:varchar(50);not null"`
	Description                 string    `gorm:"type:text"`
	ApprovalTime                time.Time `gorm:"type:time;not null"`
}
