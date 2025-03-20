// MEP-1014
package model

import (
	"time"

	"github.com/google/uuid"
)

type ItemApprovalWorkflow struct {
	ItemApprovalWorkflowID uuid.UUID   `gorm:"type:uuid;primaryKey"`
	ProcurementID          uuid.UUID   `gorm:"foreignKey:ProcurementID"`
	ItemApproversID        []uuid.UUID `gorm:"foreignKey:InstructorId"` //TODO: Fix this data type
	Status                 string      `gorm:"type:varchar(50);not null"`
	Description            string      `gorm:"type:text"`
	ApprovalTime           time.Time   `gorm:"type:time;not null"`
}
