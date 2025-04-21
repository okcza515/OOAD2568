// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type ProcurementApproval struct {
	ProcurementApprovalID uint           `gorm:"primaryKey"`
	ApproversID           uint           `gorm:"index"`
	Approver              Approver       `gorm:"foreignKey:ApproversID"`
	Status                string         `gorm:"type:varchar(50);not null"`
	Description           string         `gorm:"type:text"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	ApprovalTime          *time.Time
}
