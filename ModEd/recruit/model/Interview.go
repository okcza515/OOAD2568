// MEP-1003 Student Recruitment
package model

import (
	"ModEd/common/model"
	"encoding/json"
	"time"
)

type Interview struct {
	InterviewID          uint              `gorm:"primaryKey"`
	InstructorID         uint              `gorm:"not null"` // Foreign key referencing Instructor
	Instructor           *model.Instructor `gorm:"foreignKey:InstructorID;references:InstructorCode"`
	ApplicationReportID  uint              `gorm:"not null"` // Foreign key referencing ApplicationReport
	ApplicationReport    ApplicationReport `gorm:"foreignKey:ApplicationReportID;references:ApplicationReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ScheduledAppointment time.Time
	CriteriaScores       string
	TotalScore           float64
	EvaluatedAt          time.Time
	InterviewStatus      ApplicationStatus `gorm:"type:varchar(20)"`
}

func (i *Interview) GetID() uint {
	return i.InterviewID
}
func (i *Interview) FromCSV(csvData string) error {
	return nil
}
func (i *Interview) ToCSVRow() string {
	return ""
}
func (i *Interview) FromJSON(jsonData string) error {
	return nil
}
func (i *Interview) ToJSON() string {
	return ""
}
func (i *Interview) Validate() error {
	return nil
}
func (i *Interview) ToString() string {
	return ""
}

func (i *Interview) SetCriteriaScores(scores map[string]float64) error {
	jsonData, err := json.Marshal(scores)
	if err != nil {
		return err
	}
	i.CriteriaScores = string(jsonData)
	return nil
}

func (i *Interview) GetCriteriaScores() (map[string]float64, error) {
	var scores map[string]float64
	err := json.Unmarshal([]byte(i.CriteriaScores), &scores)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
