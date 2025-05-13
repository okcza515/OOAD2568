package model

import (
	"ModEd/common/model"
	"ModEd/core"
	"ModEd/hr/util"

	"github.com/go-playground/validator/v10"
)

type InstructorInfo struct {
	model.Instructor
	core.BaseModel
	Gender             string             `csv:"Gender" validate:"required"`
	CitizenID          string             `csv:"CitizenID" validate:"required"`
	PhoneNumber        string             `csv:"PhoneNumber" validate:"required"`
	Salary             float64            `csv:"Salary" default:"0" validate:"required"`
	AcademicPosition   AcademicPosition   `csv:"AcademicPosition" default:"0" validate:"required"`
	DepartmentPosition DepartmentPosition `csv:"DepartmentPosition" default:"0" validate:"required"`
}

func (instructorInfo InstructorInfo) Validate() error {
	validate := validator.New()
	if err := validate.Struct(instructorInfo); err != nil {
		return err
	}
	return nil
}

func NewInstructorInfo(instr model.Instructor, Gender string, CitizenID string, PhoneNumber string, Salary float64, AcademicPosition AcademicPosition, DepartmentPosition DepartmentPosition) *InstructorInfo {
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
