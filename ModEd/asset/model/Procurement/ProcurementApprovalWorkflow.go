package model

import (
	"github.com/google/uuid"
)

type ProcurementApprovalWorkflow struct {
	ProcurementApprovalWorkflowID uuid.UUID                `gorm:"type:uuid;primaryKey"`
	ApproversID                  []uuid.UUID              `gorm:"foreignKey:InstructorId"` //TO-DO: Fix this data type
	Status                       ProcurementApprovalStatus `gorm:"foreignKey:AcceptanceApprovalStatusID"`
}
