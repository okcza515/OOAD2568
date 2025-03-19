//MEP-1014
package model

import (
	"github.com/google/uuid"
)

type ItemApprovalWorkflow struct {
	ItemApprovalWorkflowID uuid.UUID   `gorm:"type:uuid;primaryKey"`
	ProcurementID          uuid.UUID   `gorm:"foreignKey:ProcurementID"`
	ItemApproversID        []uuid.UUID `gorm:"foreignKey:InstructorId"` //TO-DO: Fix this data type
	Status                 uuid.UUID   `gorm:"foreignKey:ItemApprovalStatusID"`
}
