package model

import (
	"github.com/google/uuid"
)

type ItemApprovalWorkflow struct {
	ItemApprovalWorkflowID uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ProcurementID          uuid.UUID      `gorm:"foreignKey:ProcurementID"`
	ItemApproversID        []uuid.UUID    `gorm:"foreignKey:InstructorId"` //TO-DO: Fix this data type
	Status                 ApprovalStatus `gorm:"foreignKey:ApprovalStatusID"`
}
