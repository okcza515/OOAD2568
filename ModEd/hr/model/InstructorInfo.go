package model

import (
	"ModEd/common/model"
	"ModEd/hr/util"
)

type InstructorInfo struct {
	model.Instructor
	Gender             string             `csv:"Gender"`
	CitizenID          string             `csv:"CitizenID"`
	PhoneNumber        string             `csv:"PhoneNumber"`
	Salary             int                `csv:"Salary" default:"0"`
	AcademicPosition   AcademicPosition   `csv:"AcademicPosition" default:"0"`
	DepartmentPosition DepartmentPosition `csv:"DepartmentPosition" default:"0"`
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

func NewUpdatedInstructorInfo(
	instructorInfo *InstructorInfo,
	firstName string, lastName string, email string,
	gender string, citizenID string, phoneNumber string, academicPos AcademicPosition, departmentPos DepartmentPosition,
) *InstructorInfo {
	return &InstructorInfo{
		Instructor: model.Instructor{
			InstructorCode: instructorInfo.InstructorCode,
			FirstName:      util.IfNotEmpty(firstName, instructorInfo.FirstName),
			LastName:       util.IfNotEmpty(lastName, instructorInfo.LastName),
			Email:          util.IfNotEmpty(email, instructorInfo.Email),
			StartDate:      instructorInfo.StartDate,
			Department:     instructorInfo.Department,
		},
		Gender:             util.IfNotEmpty(gender, instructorInfo.Gender),
		CitizenID:          util.IfNotEmpty(citizenID, instructorInfo.CitizenID),
		PhoneNumber:        util.IfNotEmpty(phoneNumber, instructorInfo.PhoneNumber),
		Salary:             instructorInfo.Salary,
		AcademicPosition:   AcademicPosition(util.IfNotZero(int(academicPos), int(instructorInfo.AcademicPosition))),
		DepartmentPosition: DepartmentPosition(util.IfNotZero(int(departmentPos), int(instructorInfo.DepartmentPosition))),
	}
}

func (InstructorInfo) TableName() string {
	return "instructor_infos"
}
