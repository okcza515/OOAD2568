// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type BudgetApproval struct {
	BudgetApprovalID    uint                 `gorm:"primaryKey"`
	InstrumentRequestID uint                 `gorm:"index"`
	ApproverID          *uint                `gorm:"index"`
	Approver            master.Instructor    `gorm:"foreignKey:ApproverID"`
	Status              BudgetApprovalStatus `gorm:"type:varchar(50);default:'pending'"`
	DeletedAt           gorm.DeletedAt       `gorm:"index"`
	ApprovalTime        *time.Time           `gorm:"type:datetime"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstrumentRequest   InstrumentRequest
}
