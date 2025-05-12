// MEP-1007
package model

import "ModEd/core"

type MultipleChoiceAnswerSubmission struct {
	core.BaseModel
	QuestionID   uint                 `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question     Question             `gorm:"foreignKey:QuestionID;references:ID" json:"-" csv:"-"`
	SubmissionID uint                 `gorm:"type:integer;not null" json:"submission_id" csv:"submission_id"`
	Submission   AnswerSubmission     `gorm:"foreignKey:SubmissionID;references:ID" json:"-" csv:"-"`
	ChoiceID     uint                 `gorm:"type:integer;not null" json:"choice_id" csv:"choice_id"`
	Choice       MultipleChoiceAnswer `gorm:"foreignKey:ChoiceID;references:ID" json:"student_answer" csv:"student_answer"`
}

func (s *MultipleChoiceAnswerSubmission) SetQuestionID(id uint) {
	s.QuestionID = id
}

func (s *MultipleChoiceAnswerSubmission) GetQuestionID() uint {
	return s.QuestionID
}

func (s *MultipleChoiceAnswerSubmission) SetSubmmissionID(id uint) {
	s.QuestionID = id
}

func (s *MultipleChoiceAnswerSubmission) GetSubmissionID() uint {
	return s.QuestionID
}
