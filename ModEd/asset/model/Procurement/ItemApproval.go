// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemApproval struct {
	ItemApprovalID  uint           `gorm:"primaryKey"`
	ProcurementID   uint           `gorm:"foreignKey:ProcurementID"`
	ItemApproversID uint           `gorm:"foreignKey:ItemApproversID"` //TODO: Fix this data type
	Status          string         `gorm:"type:varchar(50);not null"`
	Description     string         `gorm:"type:text"`
	ApprovalTime    time.Time      `gorm:"type:time;not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
