// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type BudgetApproval struct {
	BudgetApprovalID uint                 `gorm:"primaryKey"`
	ItemRequestID    uint                 `gorm:"not null"`
	ApproverName     string               `gorm:"type:varchar(255);not null"`
	Status           BudgetApprovalStatus `gorm:"type:varchar(50);default:'pending'"`
	DeletedAt        gorm.DeletedAt       `gorm:"index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
