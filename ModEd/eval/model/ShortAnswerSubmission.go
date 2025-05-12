// MEP-1007
package model

import (
	"ModEd/core"
)

type ShortAnswerSubmission struct {
	core.BaseModel
	QuestionID    uint             `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question      Question         `gorm:"foreignKey:QuestionID;references:ID" json:"-" csv:"-"`
	SubmissionID  uint             `gorm:"type:integer;not null" json:"submission_id" csv:"submission_id"`
	Submission    AnswerSubmission `gorm:"foreignKey:SubmissionID;references:ID" json:"-" csv:"-"`
	StudentAnswer string           `gorm:"type:text;not null" json:"student_answer" csv:"student_answer"`
}

func (s *ShortAnswerSubmission) SetQuestionID(id uint) {
	s.QuestionID = id
}

func (s *ShortAnswerSubmission) GetQuestionID() uint {
	return s.QuestionID
}

func (s *ShortAnswerSubmission) SetSubmmissionID(id uint) {
	s.QuestionID = id
}

func (s *ShortAnswerSubmission) GetSubmissionID() uint {
	return s.QuestionID
}
