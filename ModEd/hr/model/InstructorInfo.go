package model

import (
	"ModEd/common/model"
)

type InstructorInfo struct {
	model.Instructor
	Gender             string             `csv:"Gender"`
	CitizenID          string             `csv:"CitizenID"`
	PhoneNumber        string             `csv:"PhoneNumber"`
	Salary             int                `csv:"Salary"`
	AcademicPosition   AcademicPosition   `csv:"AcademicPosition"`
	DepartmentPosition DepartmentPosition `csv:"DepartmentPosition"`
}
