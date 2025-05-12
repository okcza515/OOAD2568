package model

import (
	"ModEd/core"

	commonModel "ModEd/common/model"
)

type Progress struct {
	StudentCode  string               `gorm:"column:student_code" csv:"student_code"`
	AssignmentId uint                 `gorm:"column:assignment_id" csv:"Assignment_id"`
	IsSubmitted  bool                 `gorm:"column:submitted" csv:"submitted"`
	TotalSubmit  uint                 `gorm:"column:total_submit" csv:"total_submit"`
	Student      commonModel.Student  `gorm:"foreignKey:StudentCode;references:StudentCode" csv:"-"`
	Assignment   Assignment           `gorm:"foreignKey:AssignmentId;references:AssignmentId" csv:"-"`
	SubmittedId  uint                 `gorm:"column:submitted_id" csv:"-"`
	Submitted    AssignmentSubmission `gorm:"foreignKey:SubmittedId;references:ID" csv:"-"`
	core.BaseModel
}
