// MEP-1007
package model

import "ModEd/core"

type TrueFalseAnswer struct {
	core.BaseModel
	QuestionID uint     `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID" json:"question" csv:"question"`
	IsExpected bool     `gorm:"type:boolean;not null" json:"is_expected" csv:"is_expected"`
}

func (a *TrueFalseAnswer) SetQuestionID(id uint) {
	a.QuestionID = id
}

func (a *TrueFalseAnswer) GetQuestionID() uint {
	return a.QuestionID
}
