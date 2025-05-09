// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type AcceptanceApproval struct {
	AcceptanceApprovalID uint `gorm:"primaryKey"`
	ProcurementID        uint `gorm:"index"`
	Procurement          *Procurement
	ApproverID           *uint             `gorm:"index"`
	Approver             master.Instructor `gorm:"foreignKey:ApproverID"`
	Status               AcceptanceStatus  `gorm:"type:varchar(50);default:'pending'"`
	ApprovalTime         *time.Time        `gorm:"type:datetime"`
	DeletedAt            gorm.DeletedAt    `gorm:"index"`
	CreatedAt            time.Time
}
