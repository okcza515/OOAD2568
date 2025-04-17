// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type BudgetAllocation struct {
	BudgetAllocationID uint    `gorm:"primaryKey"`
	ItemRequestID      uint    `gorm:"not null"`
	Amount             float64 `gorm:"not null"` // Approved budget
	ApproverName       string  `gorm:"type:varchar(255)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"` // Soft delete
}
