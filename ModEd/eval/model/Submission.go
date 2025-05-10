package model

import (
	"ModEd/common/model"
	"ModEd/core"
)

type AnswerSubmission struct {
	core.BaseModel
	StudentID   uint          `gorm:"type:integer;not null" json:"student_id" csv:"student_id"`
	ExamID      uint          `gorm:"type:integer;not null" json:"exam_id" csv:"exam_id"`
	Score       float64       `gorm:"type:decimal(5,2);not null" json:"score" csv:"score"`
	Student     model.Student `gorm:"foreignKey:StudentID;references:ID" json:"-" csv:"-"`
	Examination Exam          `gorm:"foreignKey:ExamID;references:ID" json:"-" csv:"-"`
}
