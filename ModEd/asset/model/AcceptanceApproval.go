// MEP-1014
package model

import (
	"time"
	master "ModEd/common/model"

	"gorm.io/gorm"
)

type AcceptanceApproval struct {
	AcceptanceApprovalID uint           `gorm:"primaryKey"`
	ProcurementID        uint           `gorm:"index"`
	Procurement          Procurement  
	ApproverID            *uint             `gorm:"index"`
	Approver              master.Instructor `gorm:"foreignKey:ApproverID"`	
	Status                AcceptanceStatus `gorm:"type:varchar(50);default:'pending'"`
	ApprovalTime         time.Time      `gorm:"type:time;not null"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	CreatedAt            time.Time
}
