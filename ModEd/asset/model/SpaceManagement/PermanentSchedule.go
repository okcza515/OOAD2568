// MEP-1013
package spacemanagement

import (
	master "ModEd/common/model"
	curriculum "ModEd/curriculum/model"

	"gorm.io/gorm"
)

type PermanentSchedule struct {
	gorm.Model
	TimeTableID uint       		   `gorm:"type:integer" json:"time_table_id" csv:"time_table_id"` 
	TimeTable   TimeTable  		   `gorm:"foreignKey:ID;references:ID" json:"time_table"`
	Faculty     master.Faculty     `gorm:"foreignKey:ID;references:ID"`
	Department  master.Department  `gorm:"foreignKey:ID;references:ID"`
	ProgramType master.ProgramType `gorm:"foreignKey:ProgramType;references:ProgramType"`
	Classroom   Room               `gorm:"foreignKey:RoomID;references:RoomID"`
	Course      curriculum.Course  `gorm:"foreignKey:CourseId;references:CourseId"`
	Class       curriculum.Class   `gorm:"foreignKey:ClassId;references:ClassId"`
}
