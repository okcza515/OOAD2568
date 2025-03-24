// MEP-1013
package spacemanagement

import (
	master "ModEd/common/model"
	curriculum "ModEd/curriculum/model"
	"time"

	"gorm.io/gorm"
)

type PermanentSchedule struct {
	gorm.Model
	ScheduleID  uint               `gorm:"type:integer ;primaryKey"`
	StartDate   time.Time          `gorm:"type:timestamp"`
	EndDate     time.Time          `gorm:"type:timestamp"`
	IsAvailable bool               `gorm:"type:boolean"`
	Faculty     master.Faculty     `gorm:"foreignKey:ID;references:ID"`
	Department  master.Department  `gorm:"foreignKey:ID;references:ID"`
	ProgramType master.ProgramType `gorm:"foreignKey:ProgramType;references:ProgramType"`
	Classroom   Room               `gorm:"foreignKey:RoomID;references:RoomID"`
	Course      curriculum.Course  `gorm:"foreignKey:CourseId;references:CourseId"`
	Class       curriculum.Class   `gorm:"foreignKey:ClassId;references:ClassId"`
}
