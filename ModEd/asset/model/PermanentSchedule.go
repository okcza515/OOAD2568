// MEP-1013
package model

import (
	"fmt"

	master "ModEd/common/model"
	"ModEd/core"
	curriculum "ModEd/curriculum/model"
)

type PermanentSchedule struct {
	core.BaseModel
	TimeTableID   uint               `gorm:"type:integer;not null;uniqueIndex" json:"time_table_id" csv:"time_table_id"`
	TimeTable     TimeTable          `gorm:"foreignKey:TimeTableID;references:ID" json:"time_table"`
	FacultyID     uint               `gorm:"type:integer;not null" json:"faculty_id" csv:"faculty_id"`
	Faculty       master.Faculty     `gorm:"foreignKey:FacultyID;references:ID"`
	DepartmentID  uint               `gorm:"type:integer;not null" json:"department_id" csv:"department_id"`
	Department    master.Department  `gorm:"foreignKey:DepartmentID;references:ID"`
	ProgramtypeID uint               `gorm:"type:integer;not null" json:"programtype_id" csv:"programtype_id"`
	Programtype   master.ProgramType `gorm:"foreignKey:ProgramtypeID;references:ID"`
	CourseId      uint               `gorm:"type:integer;not null" json:"course_id" csv:"course_id"`
	Course        curriculum.Course  `gorm:"foreignKey:CourseId;references:CourseId"`
	ClassId       uint               `gorm:"type:integer;not null" json:"class_id" csv:"class_id"`
	Class         curriculum.Class   `gorm:"foreignKey:ClassId;references:ClassId"`
}

func (ps PermanentSchedule) ToString() string {
	return fmt.Sprintf("==================================================================\n"+
		" PERMANENT SCHEDULE #%-43d \n"+
		"------------------------------------------------------------------\n"+
		" Created:                      %-25s \n"+
		" Updated:                      %-25s \n"+
		"------------------------------------------------------------------\n"+
		" Faculty:                      %-25s \n"+
		" Department:                   %-25s \n"+
		" Program Type:                 %-25s \n"+
		"------------------------------------------------------------------\n"+
		" Course:                       %-25s \n"+
		" Class/Section:                %-25d \n"+
		"------------------------------------------------------------------\n"+
		" Room:                         %-25s \n"+
		" Building:                     %-25s \n"+
		" Floor:                        %-25d \n"+
		"------------------------------------------------------------------\n"+
		" Start:                        %-25s \n"+
		" End:                          %-25s \n"+
		" Booking Type:                 %-25s \n"+
		"==================================================================",
		ps.ID,
		ps.CreatedAt.Format("2006-01-02 15:04:05"),
		ps.UpdatedAt.Format("2006-01-02 15:04:05"),
		getFacultyName(ps.Faculty),
		getDepartmentName(ps.Department),
		getProgramTypeName(ps.ProgramtypeID),
		getCourseName(ps.Course),
		ps.ClassId,
		ps.TimeTable.Room.RoomName,
		ps.TimeTable.Room.Building,
		ps.TimeTable.Room.Floor,
		ps.TimeTable.StartDate.Format("2006-01-02 15:04"),
		ps.TimeTable.EndDate.Format("2006-01-02 15:04"),
		string(ps.TimeTable.BookingType))
}

func getFacultyName(faculty master.Faculty) string {
	if faculty.Name == "" {
		return fmt.Sprintf("ID: %d", faculty.ID)
	}
	return faculty.Name
}

func getDepartmentName(department master.Department) string {
	if department.Name == "" {
		return fmt.Sprintf("ID: %d", department.ID)
	}
	return department.Name
}

func getProgramTypeName(programTypeID uint) string {
	switch programTypeID {
	case 0:
		return "Regular"
	case 1:
		return "International"
	default:
		return fmt.Sprintf("ID: %d", programTypeID)
	}
}

func getCourseName(course curriculum.Course) string {
	if course.Name == "" {
		return fmt.Sprintf("ID: %d", course.CourseId)
	}
	return course.Name
}
