package model

import "ModEd/core"

type MultipleChoiceAnswerSubmission struct {
	core.BaseModel
	QuestionID   uint                 `gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	SubmissionID uint                 `gorm:"type:integer;not null" json:"submission_id" csv:"submission_id"`
	ChoiceID     uint                 `gorm:"type:integer;not null" json:"choice_id" csv:"choice_id"`
	Choice       MultipleChoiceAnswer `gorm:"foreignKey:ChoiceID;references:ID" json:"student_answer" csv:"student_answer"`
	Question     Question             `gorm:"foreignKey:QuestionID;references:ID" json:"-" csv:"-"`
	Submission   AnswerSubmission     `gorm:"foreignKey:SubmissionID;references:ID" json:"-" csv:"-"`
}
