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

func NewInstructorInfo(instr model.Instructor, Gender string, CitizenID string, PhoneNumber string, Salary int, AcademicPosition AcademicPosition, DepartmentPosition DepartmentPosition) *InstructorInfo {
	return &InstructorInfo{
		Instructor:         instr,
		Gender:             Gender,
		CitizenID:          CitizenID,
		PhoneNumber:        PhoneNumber,
		Salary:             Salary,
		AcademicPosition:   AcademicPosition,
		DepartmentPosition: DepartmentPosition,
	}
}

func (InstructorInfo) TableName() string {
	return "instructor_infos"
}
