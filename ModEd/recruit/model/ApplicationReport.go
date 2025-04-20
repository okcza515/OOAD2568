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

func (i *ApplicationReport) GetID() uint {
	return i.ApplicantID
}
func (i *ApplicationReport) FromCSV(csvData string) error {
	return nil
}
func (i *ApplicationReport) ToCSVRow() string {
	return ""
}
func (i *ApplicationReport) FromJSON(jsonData string) error {
	return nil
}
func (i *ApplicationReport) ToJSON() string {
	return ""
}
func (i *ApplicationReport) Validate() error {
	return nil
}
func (i *ApplicationReport) ToString() string {
	return ""
}
