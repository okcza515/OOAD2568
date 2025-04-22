package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReportType            ReportType             `gorm:"type:varchar(50);not null"`
	SubmissionDate        *time.Time             `gorm:"type:date"`
	DueDate               time.Time              `gorm:"type:date;not null"`
	ScoreReportAdvisors   []ScoreReportAdvisor   `gorm:"foreignKey:ReportId;references:ID"`
	ScoreReportCommittees []ScoreReportCommittee `gorm:"foreignKey:ReportId;references:ID"`
	SeniorProjectId       uint
	SeniorProject         SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}

func (r *Report) GetID() uint {
	return r.ID
}
func (r *Report) ToString() string {
	return ""
}
func (r *Report) Validate() error {
	if r.DueDate.Before(*r.SubmissionDate) {
		return fmt.Errorf("DueDate cannot be earlier than SubmissionDate")
	}
	return nil
}
func (r *Report) ToCSVRow() string {
	return fmt.Sprintf("%d,%s,%s,%s", r.ID, r.ReportType, r.SubmissionDate, r.DueDate)
}
func (r *Report) FromCSV(raw string) error {
	return nil
}
func (r *Report) ToJSON() string {
	return ""
}
func (r *Report) FromJSON(raw string) error {
	return nil
}
