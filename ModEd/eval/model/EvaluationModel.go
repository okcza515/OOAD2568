// 65070503445
// MEP-1006
package model

import (
	"time"
)

type Evaluation struct {
	StudentCode    string    `csv:"student_code"`
	InstructorCode string    `csv:"instructor_code"`
	AssignmentId   *uint     `csv:"assignment_id"`
	QuizId         *uint     `csv:"quiz_id"`
	Score          uint      `csv:"score"`
	Comment        string    `csv:"comment"`
	EvaluatedAt    time.Time `csv:"evaluated_at"`
}
