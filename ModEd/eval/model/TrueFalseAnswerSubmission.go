// MEP-1007
package model

import (
	"ModEd/core"
)

type TrueFalseAnswerSubmission struct {
	core.BaseModel
	QuestionID    uint     `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question      Question `gorm:"foreignKey:QuestionID;references:ID" json:"-" csv:"-"`
	SubmissionID  uint     `gorm:"type:integer;not null" json:"submission_id" csv:"submission_id"`
	StudentAnswer bool     `gorm:"type:boolean;not null" json:"student_answer" csv:"student_answer"`
}
