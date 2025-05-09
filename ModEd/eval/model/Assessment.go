package model

import (
	"time"

	commonModel "ModEd/common/model"
	"ModEd/core"
	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type AssessmentStatus string

const (
	StatusDraft     AssessmentStatus = "drafted"
	StatusPublished AssessmentStatus = "published"
	StatusClosed    AssessmentStatus = "closed"
)

type Assessment struct {
	core.BaseModel
	AssessmentId   uint   `gorm:"not null;unique" csv:"Assessment_Id" json:"Assessment_Id"`
	Title          string `gorm:"type:varchar(20);not null"`
	Description    string `gorm:"type:varchar(255);"`
	PublishDate    time.Time
	DueDate        time.Time
	Status         AssessmentStatus
	ClassId        curriculumModel.Class  `gorm:"foreignKey:ClassId;references:ClassId"`
	InstructorCode commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
	Submission     []AssessmentSubmission
	State          AssessmentState
}

type AssessmentState interface {
	HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error
	HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error
}

type AssessmentSubmission struct {
	gorm.Model
	StudentCode string `gorm:"foreignKey:StudentCode;references:StudentCode"`
	FirstName   string
	LastName    string
	Email       string
	Answers     pathFile.PathFile
	Submitted   bool
	UpdatedAt   time.Time `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	Score       float64
	Feedback    string
}
