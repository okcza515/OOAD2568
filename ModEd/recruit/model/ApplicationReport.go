// MEP-1003 Student Recruitment
package model

import "ModEd/common/model"

type ApplicationReport struct {
	ApplicationReportID uint `gorm:"primaryKey"`

	ApplicantID uint
	Applicant   Applicant `gorm:"foreignKey:ApplicantID;references:ApplicantID"`

	ApplicationRoundsID uint
	ApplicationRound    ApplicationRound `gorm:"foreignKey:ApplicationRoundsID;references:RoundID"`

	FacultyID uint
	Faculty   *model.Faculty `gorm:"foreignKey:FacultyID;references:ID"`

	DepartmentID uint
	Department   *model.Department `gorm:"foreignKey:DepartmentID;references:ID"`

	Program *model.ProgramType `gorm:"type:varchar(20)"`

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
