package model

import (
	"ModEd/core"

	commonModel "ModEd/common/model"

	"fmt"
)

type Progress struct {
	core.BaseModel
	StudentCode  string
	Student      commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
	AssessmentId uint
	Assessment   Assessment       `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	Type         AssessmentType   `gorm:"column:type;not null"`
	Status       AssessmentStatus `gorm:"column:status;not null"`
	TotalSubmit  uint             `gorm:"column:total_submit;default:0"`
}

func (p Progress) GetID() uint {
	return p.ID
}

func (p Progress) ToString() string {
	return fmt.Sprintf("Progress{ID: %d, StudentCode: %s, AssessmentId: %d, Type: %s, Status: %s, TotalSubmit: %d}",
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.TotalSubmit)
}

func (p Progress) Validate() error {
	if p.StudentCode == "" {
		return fmt.Errorf("Student code is required")
	}
	if p.AssessmentId == 0 {
		return fmt.Errorf("Assessment ID is required")
	}
	if p.Type == "" {
		return fmt.Errorf("Assessment type is required")
	}
	if p.Status == "" {
		return fmt.Errorf("Status is required")
	}
	return nil
}

func (p Progress) ToCSVRow() string {
	return fmt.Sprintf("%d, %s, %d, %s, %s, %d",
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.TotalSubmit)
}

func (p Progress) FromCSV(raw string) error {
	// TODO: Implement CSV parsing
	return nil
}

func (p Progress) ToJSON() string {
	return fmt.Sprintf(`{"id":%d,"student_code":"%s","assessment_id":%d,"type":"%s","status":"%s", "total_submit":%d}`,
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.TotalSubmit)
}

func (p Progress) FromJSON(raw string) error {
	// TODO: Implement JSON parsing
	return nil
}
