// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type BudgetAllocation struct {
	BudgetAllocationID  uint              `gorm:"primaryKey"`
	InstrumentRequestID uint              `gorm:"index"`
	Amount              float64           `gorm:"not null"`
	TotalbudgetAmount   float64           `gorm:"not null"`
	ApproverID          uint              `gorm:"index"`
	Approver            master.Instructor `gorm:"foreignKey:ApproverID"`
	DeletedAt           gorm.DeletedAt    `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstrumentRequest   InstrumentRequest
}
