// MEP-1007
package model

import (
	"ModEd/core"
)

type ExamSection struct {
	core.BaseModel	
	ExamID   		    string 			`gorm:"not null" csv:"exam_id" json:"exam_id"`
	Exam				Examination 	`gorm:"foreignKey:ExamID;references:ID" csv:"-" json:"-"`
	SectionNo           uint			`gorm:"not null" csv:"section_no" json:"section_no"`
	Description			string			`gorm:"not null" csv:"description" json:"description"`
	NumQuestions		int				`gorm:"not null" csv:"num_question" json:"num_question"`
	Score				float64			`gorm:"not null" csv:"score" json:"score"`
}