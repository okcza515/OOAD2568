package model

import (
	master "ModEd/common/model"
	curriculum "ModEd/curriculum/model"
	"time"

	"github.com/google/uuid"
)

type PermanentSchedule struct {
	ScheduleID  uuid.UUID          `gorm:"type:text;primaryKey"`
	StartDate   time.Time          `gorm:"type:timestamp"`
	EndDate     time.Time          `gorm:"type:timestamp"`
	IsAvailable bool               `gorm:"type:boolean"`
	Faculty     master.Faculty     `gorm:"foreignKey:FacultyID;references:FacultyID"`
	Department  master.Department  `gorm:"foreignKey:DepartmentID;references:DepartmentID"`
	ProgramType master.ProgramType `gorm:"foreignKey:ProgramType;references:ProgramType"`
	Classroom   string
	Course      curriculum.Course `gorm:"foreignKey:CourseID;references:CourseID"`
	Class       curriculum.Class  `gorm:"foreignKey:ClassID;references:ClassID"`
}
