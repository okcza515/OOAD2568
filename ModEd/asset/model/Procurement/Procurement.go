// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Procurement struct {
	gorm.Model
	ProcurementID         uint           `gorm:"primaryKey"`
	TORcandidate          uint           `gorm:"foreignKey:TORID"`
	ItemRequestID         uint           `gorm:"foreignKey:ItemRequestID"`
	ProcurementApprovalID uint           `gorm:"foreignKey:ProcurementApprovalID"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
