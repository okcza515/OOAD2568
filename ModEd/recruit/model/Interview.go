package model

import (
	"time"
)

// Interview struct defines an interview object
type Interview struct {
	ID                   uint      `gorm:"primaryKey"`
	InstructorID         uint      `gorm:"not null"` // Foreign key referencing Instructor
	Instructor           Instructor `gorm:"foreignKey:InstructorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ApplicantID          uint      `gorm:"not null"` // Foreign key referencing Applicant
	Applicant           Applicant  `gorm:"foreignKey:ApplicantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ScheduledAppointment time.Time
	InterviewScore       *float64  `gorm:"default:null"` // Nullable score
	InterviewStatus      string    `gorm:"column:interview_status"`
}
