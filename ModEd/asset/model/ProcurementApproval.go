// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type ProcurementApproval struct {
	ProcurementApprovalID uint              `gorm:"primaryKey"`
	ProcurementID         uint              `gorm:"index"`
	Procurement           Procurement       `gorm:"foreignKey:ProcurementID"`
	ApproverID            uint              `gorm:"index"`
	Approver              master.Instructor `gorm:"foreignKey:ApproverID"`
	Status                string            `gorm:"type:varchar(50);not null"`
	Description           string            `gorm:"type:text"`
	DeletedAt             gorm.DeletedAt    `gorm:"index"`
	ApprovalTime          *time.Time
}
