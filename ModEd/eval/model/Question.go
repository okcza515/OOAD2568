// MEP-1007
package model

import (
	"ModEd/core"
)

type Question struct {
	core.BaseModel
	SectionID      uint         `gorm:"type:integer;not null" json:"section_id" csv:"section_id"`
	Section        ExamSection  `gorm:"foreignKey:SectionID;references:ID" json:"section" csv:"section"`
	Score          float64      `gorm:"type:decimal(5,2);not null" json:"score" csv:"score"`
	ActualQuestion string       `gorm:"type:text;not null" json:"actual_question" csv:"actual_question"`
	QuestionType   QuestionType `gorm:"type:text;not null" json:"question_type" csv:"question_type"`
}
