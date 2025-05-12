// MEP-1007
package model

import "ModEd/core"

type TrueFalseAnswer struct {
	core.BaseModel
	QuestionID uint `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	IsExpected bool `gorm:"type:boolean;not null" json:"is_expected" csv:"is_expected"`
}
