package model

import (
	"gorm.io/gorm"
	"time"
)

type InternshipSchedule struct {
	gorm.Model
	InternshipScheduleId int           `gorm:"primaryKey autoIncrement"`
	Student              InternStudent `gorm:"foreignKey:InternStudentId"`
	StartDate            time.Time
	EndDate              time.Time
}
