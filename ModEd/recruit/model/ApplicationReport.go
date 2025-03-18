package model

import (
	"github.com/google/uuid"
)

// ApplicationReport represents a report of an applicant's application status.
type ApplicationReport struct {
	ApplicationReportID uuid.UUID         `gorm:"primaryKey"`
	ApplicantID         uuid.UUID         `gorm:"foreignKey:ApplicantID"`
	ApplicationRoundsID uuid.UUID         `gorm:"foreignKey:ApplicationRoundsID"`
	FacultyID           uuid.UUID         `gorm:"foreignKey:FacultyID"`
	DepartmentID        uuid.UUID         `gorm:"foreignKey:DepartmentID"`
	ApplicationStatuses ApplicationStatus `gorm:"type:varchar(20)"`
}
