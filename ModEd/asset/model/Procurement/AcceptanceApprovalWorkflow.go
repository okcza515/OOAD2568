package model

import (
	"github.com/google/uuid"
)

type AcceptanceApprovalWorkflow struct {
	AcceptanceApprovalWorkflowID uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ProcurementID                uuid.UUID      `gorm:"type:uuid;not null"`
	ApproversID                  []uuid.UUID    `gorm:"foreignKey:InstructorId"` //TO-DO: Fix this data type
	Status                       ApprovalStatus `gorm:"foreignKey:ApprovalStatusID"`
}
