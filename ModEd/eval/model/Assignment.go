package model

import (
	"time"

	commonModel "ModEd/common/model"
	"ModEd/core"
)

type AssignmentStatus string

const (
	StatusDraft     AssignmentStatus = "drafted"
	StatusPublished AssignmentStatus = "published"
	StatusClosed    AssignmentStatus = "closed"
)

type Assignment struct {
	core.BaseModel
	AssignmentId   uint   `gorm:"not null;unique" csv:"Assignment_Id" json:"Assignment_Id"`
	Title          string `gorm:"type:varchar(20);not null"`
	Description    string `gorm:"type:varchar(255);"`
	PublishDate    time.Time
	DueDate        time.Time
	Status         AssignmentStatus `gorm:"type:varchar(20);not null;default:'drafted'"`
	ClassId        uint             `gorm:"not null"`
	InstructorCode string
	Instructor     commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
	Submission     []AssignmentSubmission `gorm:"foreignKey:AssignmentId"`
	State          AssignmentState        `gorm:"-"`
}

type AssignmentState interface {
	HandleSubmission(assignment *Assignment, submission *AssignmentSubmission) error
	HandleStatusChange(assignment *Assignment, newStatus AssignmentStatus) error
}
