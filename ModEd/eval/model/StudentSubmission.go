package model

import (
	"time"

	"ModEd/common/model"

	"gorm.io/gorm"
)

type SubmissionStatus string

const (
	SubmissionStatusDraft     SubmissionStatus = "draft"
	SubmissionStatusSubmitted SubmissionStatus = "submitted"
	SubmissionStatusGraded    SubmissionStatus = "graded"
)

type Submission struct {
	gorm.Model
	//SubmissionID   uuid.UUID               `gorm:"type:text;primaryKey;not null;unique" json:"submission_id"`
	SID            model.Student    `gorm:"foreignKey:SID"`
	SubmissionDate time.Time        `gorm:"not null" json:"submission_date"`
	Status         SubmissionStatus `gorm:"type:text;default:'draft'" json:"status"`
	Content        string           `json:"content"`
	Feedback       string           `json:"feedback,omitempty"`
	Score          *float64         `json:"score,omitempty"`
	EvaluatedAt    *time.Time       `json:"evaluated_at,omitempty"`
	EvaluatedBy    model.Instructor `gorm:"foreignKey:EvaluatedByID" json:"evaluated_by,omitempty"`
	EvaluatedByID  string           `gorm:"not null" json:"evaluated_by_id,omitempty"`
	CourseID       string           `gorm:"not null" json:"course_id"`
}
