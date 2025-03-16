package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	InstructorId string      `gorm:"primaryKey;not null;unique" csv:"instructor_id" json:"instructor_id"`
	Name         string      `gorm:"not null" csv:"name" json:"name"`
	FirstName    string      `gorm:"not null" csv:"first_name" json:"first_name"`
	LastName     string      `gorm:"not null" csv:"last_name" json:"last_name"`
	Email        string      `gorm:"not null" csv:"email" json:"email"`
	StartDate    time.Time   `gorm:"not null" csv:"start_date" json:"start_date"`
	Faculty    	 Faculty     `gorm:"not null" json:"faculty"`
	Department 	 Department  `gorm:"not null" json:"department"`
	CoursesId    []uuid.UUID `gorm:"type:uuid[]" csv:"course_id" json:"course_id"` // UUID to avoid circular dependency
}