package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Approver struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	ApprovalWorkflowID uuid.UUID `gorm:"type:uuid;not null"`
	Name               string    `gorm:"type:varchar(255);not null"`
	Role               string    `gorm:"type:varchar(100);not null"`
	Email              string    `gorm:"type:varchar(255);unique"`
}
