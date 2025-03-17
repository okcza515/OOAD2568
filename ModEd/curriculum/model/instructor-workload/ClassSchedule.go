package model

import (
	"ModEd/curriculum/model"

	"github.com/google/uuid"
)

type ClassSchedule struct {
	ScheduleId    uuid.UUID   `gorm:"type:string;default:uuid_generate_v4();primaryKey"`
	Class         model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	StartDateTime string      `gorm:"type:timestamp;not null"`
	EndDateTime   string      `gorm:"type:timestamp;not null"`
}
