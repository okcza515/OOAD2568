package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassSchedule struct {
	gorm.Model
	ActionTracker
	ClassId   uint        `gorm:"index"`
	Class     model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	StartTime string      `gorm:"type:timestamp;not null"`
	EndTime   string      `gorm:"type:timestamp;not null"`
}
