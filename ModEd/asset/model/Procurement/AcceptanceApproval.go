// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type AcceptanceApproval struct {
	AcceptanceApprovalID uint           `gorm:"primaryKey"`
	Approver             Approver       `gorm:"foreignKey:ApproversID"`
	Status               string         `gorm:"type:varchar(50);not null"`
	Description          string         `gorm:"type:text"`
	ApprovalTime         time.Time      `gorm:"type:time;not null"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
