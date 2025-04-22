package assignment

import (
	"time"
)

type EvalSubmittable interface {
	GetID() uint
	GetTitle() string
	GetDescription() string
	GetStartDate() time.Time
	GetDueDate() time.Time
	GetMaxScore() float64
	GetType() string
	GetCourseID() string
}

type EvalSubmission interface {
	GetID() uint
	GetStudentID() string
	GetFirstName() string
	GetLastName() string
	GetSubmittedAt() time.Time
	GetContent() string
	IsLate() bool
	GetStatus() string
	GetScore() *float64
	GetFeedback() string
}

type EvalEvaluator interface {
	GetID() string
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetStartDate() time.Time
	GetFaculty() string
	GetDepartment() string
	GetCourseID() string
}

type EvalEvaluation interface {
	GetID() uint
	GetSubmissionID() uint
	GetEvaluatorID() string
	GetStudentID() string
	GetScore() float64
	GetFeedback() string
	GetEvaluatedAt() time.Time
	GetStatus() string
	GetGrade() string
	GetCourseID() string
}
