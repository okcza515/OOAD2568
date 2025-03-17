package model

import (
	"github.com/google/uuid"
)

type AcceptanceApprovalWorkflow struct {
	AcceptanceApprovalWorkflowID uuid.UUID                `gorm:"type:uuid;primaryKey"`
	ProcurementID                Procurement              `gorm:"foreignKey:ProcurementID"`
	ApproversID                  []uuid.UUID              `gorm:"foreignKey:InstructorId"` //TO-DO: Fix this data type
	Status                       AcceptanceApprovalStatus `gorm:"foreignKey:AcceptanceApprovalStatusID"`
}
