// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type BudgetAllocation struct {
	BudgetAllocationID uint           `gorm:"primaryKey"`
	ItemRequestID      uint           `gorm:"not null"`
	Amount             float64        `gorm:"not null"`
	ApproverName       string         `gorm:"type:varchar(255)"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
