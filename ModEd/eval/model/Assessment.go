package model

import (
	"time"

	commonModel "ModEd/common/model"
	"ModEd/core"
	curriculumModel "ModEd/curriculum/model"
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
	Status         AssessmentStatus       `gorm:"type:varchar(20);not null;default:'drafted'"`
	ClassId        curriculumModel.Class  `gorm:"foreignKey:ClassId;references:ClassId"`
	InstructorCode commonModel.Instructor `gorm:"foreignKey:InstructorCode;references:InstructorCode"`
	Submission     []AssessmentSubmission `gorm:"foreignKey:AssessmentId"`
	State          AssessmentState        `gorm:"-"`
}

type AssessmentState interface {
	HandleSubmission(assessment *Assessment, submission *AssessmentSubmission) error
	HandleStatusChange(assessment *Assessment, newStatus AssessmentStatus) error
}
