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
	ScheduleID uint      `gorm:"type:integer ;primaryKey" json:"schedule_id" csv:"schedule_id"`
	StartDate  time.Time `gorm:"type:timestamp" json:"start_date" csv:"start_date"`
	EndDate    time.Time `gorm:"type:timestamp" json:"end_date" csv:"end_date"`
	//RecurringDays []string           `gorm:"type:text[]" json:"recurring_days" csv:"recurring_days"`
	IsAvailable bool               `gorm:"type:boolean" json:"is_available" csv:"is_available"`
	Faculty     master.Faculty     `gorm:"foreignKey:Faculty;references:Faculty" json:"faculty" csv:"faculty"`
	Department  master.Department  `gorm:"foreignKey:Department;references:Department" json:"department" csv:"department"`
	ProgramType master.ProgramType `gorm:"foreignKey:ProgramType;references:ProgramType" json:"program_type" csv:"program_type"`
	Classroom   Room               `gorm:"foreignKey:RoomID;references:RoomID" json:"classroom" csv:"classroom"`
	Course      curriculum.Course  `gorm:"foreignKey:CourseId;references:CourseId" json:"course" csv:"course"`
	Class       curriculum.Class   `gorm:"foreignKey:ClassId;references:ClassId" json:"class" csv:"class"`
}
