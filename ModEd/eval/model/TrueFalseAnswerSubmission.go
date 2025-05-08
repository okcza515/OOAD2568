package model

import "ModEd/core"

type TrueFalseAnswerSubmission struct {
	core.BaseModel
	QuestionID   uint
	SubmissionID uint
	Boolean      bool `json:"answer"`
}
