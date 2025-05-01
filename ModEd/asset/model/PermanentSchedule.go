// MEP-1013
package model

import (
	master "ModEd/common/model"
	"ModEd/core"
	curriculum "ModEd/curriculum/model"
)

type PermanentSchedule struct {

	core.BaseModel
	TimeTableID   uint      			`gorm:"type:integer;not null;uniqueIndex" json:"time_table_id" csv:"time_table_id"`
	TimeTable     TimeTable 			`gorm:"foreignKey:TimeTableID;references:ID" json:"time_table"`
	FacultyID     uint      			`gorm:"type:integer;not null" json:"faculty_id" csv:"faculty_id"`
	Faculty       master.Faculty 		`gorm:"foreignKey:FacultyID;references:ID"`
	DepartmentID  uint      			`gorm:"type:integer;not null" json:"department_id" csv:"department_id"`
	Department    master.Department		`gorm:"foreignKey:DepartmentID;references:ID"`
	ProgramtypeID uint      			`gorm:"type:integer;not null" json:"programtype_id" csv:"programtype_id"`
	Programtype   master.ProgramType	`gorm:"foreignKey:ProgramtypeID;references:ID"`
	CourseId      uint      			`gorm:"type:integer;not null" json:"course_id" csv:"course_id"`
	Course        curriculum.Course 	`gorm:"foreignKey:CourseId;references:CourseId"`
	ClassId       uint      			`gorm:"type:integer;not null" json:"class_id" csv:"class_id"`
	Class         curriculum.Class 		`gorm:"foreignKey:ClassId;references:ClassId"`
}
