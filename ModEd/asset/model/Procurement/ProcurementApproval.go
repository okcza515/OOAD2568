// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type ProcurementApproval struct {
	ProcurementApprovalID uint           `gorm:"primaryKey"`
	ApproversID           uint           `gorm:"foreignKey:ApproversID"` //TODO: Fix this data type
	Status                string         `gorm:"type:varchar(50);not null"`
	Description           string         `gorm:"type:text"`
	ApprovalTime          time.Time      `gorm:"type:time;not null"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
