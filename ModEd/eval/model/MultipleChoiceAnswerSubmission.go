package model

import "ModEd/core"

type MultipleChoiceAnswerSubmission struct {
	core.BaseModel
	QuestionID   uint
	SubmissionID uint
	ChoiceID     uint
}
