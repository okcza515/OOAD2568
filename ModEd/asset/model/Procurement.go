// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Procurement struct {
	gorm.Model
	ProcurementID         uint              `gorm:"primaryKey"`
	QuotationID           uint              `gorm:"index"`
	InstrumentRequestID   uint              `gorm:"index"`
	ProcurementApprovalID uint              `gorm:"foreignKey:ProcurementApprovalID"`
	Status                ProcurementStatus `gorm:"type:varchar(50);default:'pending'"`
	DeletedAt             gorm.DeletedAt    `gorm:"index"`
	Quotation             Quotation         `gorm:"foreignKey:QuotationID"`
	InstrumentRequest     InstrumentRequest `gorm:"foreignKey:InstrumentRequestID"`
}
