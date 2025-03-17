package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternshipSchedule struct {
	gorm.Model
	ScheduleId uuid.UUID     `gorm:"primaryKey"`
	Student    InternStudent `gorm:"foreignKey:StudentID"`
	StartDate  time.Time
	EndDate    time.Time
}
