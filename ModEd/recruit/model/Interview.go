// MEP-1003 Student Recruitment
package model

import (
	"time"

	"github.com/google/uuid"
)

// Interview struct defines an interview object
type Interview struct {
	ID                   uuid.UUID `gorm:"primaryKey"`
	InstructorID         uint      `gorm:"foreignKey:InstructorID"`
	ApplicantID          uuid.UUID `gorm:"foreignKey:ApplicantID"`
	ScheduledAppointment time.Time
	InterviewScore       *float64 `gorm:"default:null"` // Nullable score
	InterviewStatus      string   `gorm:"column:interview_status"`
}
