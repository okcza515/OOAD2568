// MEP-1007
package model

import "ModEd/core"

type ShortAnswer struct {
	core.BaseModel
	QuestionID     uint   `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	ExpectedAnswer string `gorm:"type:text;not null" json:"expected_answer" csv:"expected_answer"`
}
