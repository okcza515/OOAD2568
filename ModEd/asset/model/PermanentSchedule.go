// MEP-1013
package model

import (
	"fmt"

	master "ModEd/common/model"
	"ModEd/core"
	curriculum "ModEd/curriculum/model"

	"github.com/go-playground/validator/v10"
)

type PermanentSchedule struct {
	core.BaseModel
	TimeTableID   uint               `gorm:"type:integer;not null;uniqueIndex" json:"time_table_id" csv:"time_table_id" validate:"-"`
	TimeTable     TimeTable          `gorm:"foreignKey:TimeTableID;references:ID" json:"time_table" validate:"-"`
	FacultyID     uint               `gorm:"type:integer;not null" json:"faculty_id" csv:"faculty_id" validate:"required"`
	Faculty       master.Faculty     `gorm:"foreignKey:FacultyID;references:ID" validate:"-"`
	DepartmentID  uint               `gorm:"type:integer;not null" json:"department_id" csv:"department_id" validate:"required"`
	Department    master.Department  `gorm:"foreignKey:DepartmentID;references:ID" validate:"-"`
	ProgramtypeID uint               `gorm:"type:integer;not null" json:"programtype_id" csv:"programtype_id" validate:"required"`
	Programtype   master.ProgramType `gorm:"foreignKey:ProgramtypeID;references:ID" validate:"-"`
	CourseId      uint               `gorm:"type:integer;not null" json:"course_id" csv:"course_id" validate:"required"`
	Course        curriculum.Course  `gorm:"foreignKey:CourseId;references:CourseId" validate:"-"`
	ClassId       uint               `gorm:"type:integer;not null" json:"class_id" csv:"class_id" validate:"required"`
	Class         curriculum.Class   `gorm:"foreignKey:ClassId;references:ClassId" validate:"-"`
}

func (ps PermanentSchedule) Validate() error {
	validate := validator.New()

	// Validate struct fields using v10 validator
	if err := validate.Struct(ps); err != nil {
		return err
	}

	return nil
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
