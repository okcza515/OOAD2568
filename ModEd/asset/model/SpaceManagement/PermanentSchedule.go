package space

import (
	master "ModEd/common/model"
	curriculum "ModEd/curriculum/model"
	"time"

	"gorm.io/gorm"
)

type PermanentSchedule struct {
	gorm.Model
	ScheduleID  uint               `gorm:"type:integer ;primaryKey" json:"schedule_id" csv:"schedule_id"`
	StartDate   time.Time          `gorm:"type:timestamp" json:"start_date" csv:"start_date"`
	EndDate     time.Time          `gorm:"type:timestamp" json:"end_date" csv:"end_date"`
	IsAvailable bool               `gorm:"type:boolean" json:"is_available" csv:"is_available"`
	Faculty     master.Faculty     `gorm:"foreignKey:FacultyID;references:FacultyID" json:"faculty" csv:"faculty"`
	Department  master.Department  `gorm:"foreignKey:DepartmentID;references:DepartmentID" json:"department" csv:"department"`
	ProgramType master.ProgramType `gorm:"foreignKey:ProgramType;references:ProgramType" json:"program_type" csv:"program_type"`
	Classroom   Room               `gorm:"foreignKey:RoomID;references:RoomID" json:"classroom" csv:"classroom"`
	Course      curriculum.Course  `gorm:"foreignKey:CourseID;references:CourseID" json:"course" csv:"course"`
	Class       curriculum.Class   `gorm:"foreignKey:ClassID;references:ClassID" json:"class" csv:"class"`
}
