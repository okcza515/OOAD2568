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
		idWidth, "ID", timetableWidth, "TimeTable", facultyWidth, "Faculty",
		departmentWidth, "Department", programWidth, "Program",
		courseWidth, "Course", classWidth, "Section")

	// Faculty
	faculty := fmt.Sprintf("%d", ps.FacultyID)
	if ps.Faculty.Name != "" {
		faculty = ps.Faculty.Name
	}

	// Department
	department := fmt.Sprintf("%d", ps.DepartmentID)
	if ps.Department.Name != "" {
		department = ps.Department.Name
	}

	// Program type
	programType := "Unknown"
	if ps.ProgramtypeID == 0 {
		programType = "Regular"
	} else if ps.ProgramtypeID == 1 {
		programType = "International"
	}

	// Course
	course := fmt.Sprintf("%d", ps.CourseId)
	if ps.Course.Name != "" {
		course = ps.Course.Name
	}

	// Section
	section := ps.Class.Section

	dataRow := fmt.Sprintf("| %-*d | %-*d | %-*s | %-*s | %-*s | %-*s | %-*d |",
		idWidth, ps.ID,
		timetableWidth, ps.TimeTableID,
		facultyWidth, truncate(faculty, facultyWidth),
		departmentWidth, truncate(department, departmentWidth),
		programWidth, truncate(programType, programWidth),
		courseWidth, truncate(course, courseWidth),
		classWidth, section)

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		headerBorder, headerRow, headerBorder, dataRow, headerBorder)
}
