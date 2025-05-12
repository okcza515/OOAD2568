// MEP-1007
package model

import (
	"ModEd/core"
)

type ExamSection struct {
	core.BaseModel
	ExamID       uint       `gorm:"type:integer;not null" json:"exam_id" csv:"exam_id"`
	Exam         Exam       `gorm:"foreignKey:ExamID;references:ID" json:"exam" csv:"exam"`
	SectionNo    uint       `gorm:"type:integer;not null" json:"section_no" csv:"section_no"`
	Description  string     `gorm:"type:text;not null" json:"description" csv:"description"`
	NumQuestions int        `gorm:"type:integer;not null" json:"num_questions" csv:"num_questions"`
	Score        float64    `gorm:"type:decimal(5,2);not null" json:"score" csv:"score"`
	Questions    []Question `gorm:"foreignKey:SectionID;references:ID" json:"-" csv:"-"`
}
