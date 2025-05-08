package model

import (
	"ModEd/core"
)

type ShortAnswerSubmission struct {
	core.BaseModel
	QuestionID    uint
	SubmissionID  uint
	StudentAnswer string
}
