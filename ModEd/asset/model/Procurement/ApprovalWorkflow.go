package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApprovalWorkflow struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ProcurementID uuid.UUID  `gorm:"type:uuid;not null"`
	Approvers     []Approver `gorm:"foreignKey:ApprovalWorkflowID"`
	Status        string     `gorm:"type:varchar(50);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// set the ID before creating a new record
func (a *ApprovalWorkflow) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return nil
}