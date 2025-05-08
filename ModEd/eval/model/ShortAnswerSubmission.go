package model

import (
	"ModEd/core"
)

type ShortAnswer struct {
	core.BaseModel
	QuestionID    uint
	SubmissionID  uint
	StudentAnswer string
}
