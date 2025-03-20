package model

import (
	"ModEd/curriculum/model"
)

type ClassSchedule struct {
	ScheduleId    int         `gorm:"primaryKey;autoIncrement"`
	Class         model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	StartDateTime string      `gorm:"type:timestamp;not null"`
	EndDateTime   string      `gorm:"type:timestamp;not null"`
}
