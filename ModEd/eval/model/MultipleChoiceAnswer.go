// MEP-1007
package model

import "ModEd/core"

type MultipleChoiceAnswer struct {
	core.BaseModel
	QuestionID  uint   `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	AnswerLabel string `gorm:"type:text;not null" json:"answer_label" csv:"answer_label"`
	IsExpected  bool   `gorm:"type:boolean;not null" json:"is_expected" csv:"is_expected"`
}
