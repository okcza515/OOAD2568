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

func (s *TrueFalseAnswerSubmission) SetQuestionID(id uint) {
	s.QuestionID = id
}

func (s *TrueFalseAnswerSubmission) GetQuestionID() uint {
	return s.QuestionID
}

func (s *TrueFalseAnswerSubmission) SetSubmmissionID(id uint) {
	s.SubmissionID = id
}

func (s *TrueFalseAnswerSubmission) GetSubmissionID() uint {
	return s.SubmissionID
}
