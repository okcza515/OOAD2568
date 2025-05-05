// MEP-1013
package model

import (
	"fmt"
	"strings"

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
	truncate := func(s string, maxLen int) string {
		if len(s) > maxLen {
			return s[:maxLen-3] + "..."
		}
		return s
	}

	idWidth := 5
	timetableWidth := 10
	facultyWidth := 20
	departmentWidth := 20
	programWidth := 15
	courseWidth := 20
	classWidth := 10

	headerBorder := "+" +
		strings.Repeat("-", idWidth+2) + "+" +
		strings.Repeat("-", timetableWidth+2) + "+" +
		strings.Repeat("-", facultyWidth+2) + "+" +
		strings.Repeat("-", departmentWidth+2) + "+" +
		strings.Repeat("-", programWidth+2) + "+" +
		strings.Repeat("-", courseWidth+2) + "+" +
		strings.Repeat("-", classWidth+2) + "+"

	headerRow := fmt.Sprintf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |",
		idWidth, "ID", timetableWidth, "TimeTable", facultyWidth, "Faculty", departmentWidth, "Department",
		programWidth, "Program", courseWidth, "Course", classWidth, "Section")

	dataRow := fmt.Sprintf("| %-*d | %-*d | %-*s | %-*s | %-*s | %-*s | %-*d |",
		idWidth, ps.ID,
		timetableWidth, ps.TimeTableID,
		facultyWidth, truncate(ps.Faculty.Name, facultyWidth),
		departmentWidth, truncate(ps.Department.Name, departmentWidth),
		programWidth, truncate(ps.Programtype.String(), programWidth),
		courseWidth, truncate(ps.Course.Name, courseWidth),
		classWidth, ps.Class.Section)

	result := headerBorder + "\n"
	result += headerRow + "\n"
	result += headerBorder + "\n"
	result += dataRow + "\n"
	result += headerBorder + "\n"

	return result
}
