// MEP-1007
package model

import "ModEd/core"

type AnswerSubmission struct {
	core.BaseModel
	StudentID uint `json:"student_id"`
	ExamID    uint `json:"exam_id"`
	Score     float64
}
