package model

import (
	"ModEd/core"

	commonModel "ModEd/common/model"

	"fmt"
)

type Progress struct {
	StudentCode  string               `gorm:"column:student_code" csv:"student_code"`
	AssessmentId uint                 `gorm:"column:assessment_id" csv:"Assessment_id"`
	IsSubmitted  bool                 `gorm:"column:submitted" csv:"submitted"`
	TotalSubmit  uint                 `gorm:"column:total_submit" csv:"total_submit"`
	Student      commonModel.Student  `gorm:"foreignKey:StudentCode;references:StudentCode" csv:"-"`
	Assessment   Assessment           `gorm:"foreignKey:AssessmentId;references:AssessmentId" csv:"-"`
	SubmittedId  uint                 `gorm:"column:submitted_id" csv:"-"`
	Submitted    AssessmentSubmission `gorm:"foreignKey:SubmittedId;references:ID" csv:"-"`
	core.BaseModel
}

func (p Progress) GetID() uint {
	return p.ID
}

func (p Progress) ToString() string {
	return fmt.Sprintf("Progress{ID: %d, StudentCode: %s, AssessmentId: %d, Submitted: %v, TotalSubmit: %d}",
		p.ID, p.StudentCode, p.AssessmentId, p.IsSubmitted, p.TotalSubmit)
}

func (p Progress) Validate() error {
	if p.StudentCode == "" {
		return fmt.Errorf("student code is required")
	}
	if p.AssessmentId == 0 {
		return fmt.Errorf("assessment ID is required")
	}
	return nil
}

func (p Progress) ToCSVRow() string {
	return fmt.Sprintf("%d, %s, %d, %v, %d",
		p.ID, p.StudentCode, p.AssessmentId, p.IsSubmitted, p.TotalSubmit)
}

func (p Progress) FromCSV(raw string) error {
	// TODO: Implement CSV parsing
	return nil
}

func (p Progress) ToJSON() string {
	return fmt.Sprintf(`{"student_code":"%s","assessment_id":%d,"submitted":%v,"total_submit":%d}`,
		p.StudentCode, p.AssessmentId, p.IsSubmitted, p.TotalSubmit)
}

func (p Progress) FromJSON(raw string) error {
	return nil
}
