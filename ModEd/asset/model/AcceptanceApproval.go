// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type AcceptanceApproval struct {
	AcceptanceApprovalID uint               `gorm:"primaryKey"`
	ProcurementID        uint               `gorm:"index;not null"`
	Procurement          *Procurement       `gorm:"foreignKey:ProcurementID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ApproverID           *uint              `gorm:"index"`
	Approver             *master.Instructor `gorm:"foreignKey:ApproverID;references:ID"`
	Status               AcceptanceStatus   `gorm:"type:varchar(50);default:'pending'"`
	ApprovalTime         *time.Time         `gorm:"type:datetime"`
	InstrumentsCreated   bool               `gorm:"default:false"`
	DeletedAt            gorm.DeletedAt     `gorm:"index"`
	CreatedAt            time.Time
}
