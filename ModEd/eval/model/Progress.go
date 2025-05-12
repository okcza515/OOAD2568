package model

import (
	"ModEd/core"

	commonModel "ModEd/common/model"
)

type Progress struct {
	core.BaseModel
	StudentCode  string               `gorm:"column:student_code" csv:"student_code"`
	AssessmentId uint                 `gorm:"column:assessment_id" csv:"Assessment_id"`
	IsSubmitted  bool                 `gorm:"column:submitted" csv:"submitted"`
	TotalSubmit  uint                 `gorm:"column:total_submit" csv:"total_submit"`
	Student      commonModel.Student  `gorm:"foreignKey:StudentCode;references:StudentCode" csv:"-"`
	Assessment   Assessment           `gorm:"foreignKey:AssessmentId;references:AssessmentId" csv:"-"`
	SubmittedId  uint                 `gorm:"column:submitted_id" csv:"-"`
	Submitted    AssessmentSubmission `gorm:"foreignKey:SubmittedId;references:ID" csv:"-"`
}
