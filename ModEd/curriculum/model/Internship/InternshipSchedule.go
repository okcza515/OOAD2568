package model

import (
	"time"

	"gorm.io/gorm"
)

type InternshipSchedule struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
}
