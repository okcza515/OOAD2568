// MEP-1003 Student Recruitment
package model

// ApplicationReport represents a report of an applicant's application status.
type ApplicationReport struct {
	ApplicationReportID uint              `gorm:"primaryKey"`
	ApplicantID         uint              `gorm:"foreignKey:ApplicantID"`
	ApplicationRoundsID uint              `gorm:"foreignKey:ApplicationRoundsID"`
	FacultyID           uint              `gorm:"foreignKey:FacultyID"`
	DepartmentID        uint              `gorm:"foreignKey:DepartmentID"`
	ApplicationStatuses ApplicationStatus `gorm:"type:varchar(20)"`
}
