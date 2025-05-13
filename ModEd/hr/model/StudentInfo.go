package model

import (
	"ModEd/common/model"
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/util"
)

type StudentInfo struct {
	model.Student
	core.BaseModel
	Gender      string           `csv:"Gender" json:"Gender"`
	CitizenID   string           `csv:"CitizenID" json:"CitizenID"`
	PhoneNumber string           `csv:"PhoneNumber" json:"PhoneNumber" validation:"phone"`
	AdvisorCode string           `csv:"AdvisorCode" json:"AdvisorCode"`
	Advisor     model.Instructor `csv:"Advisor" json:"Advisor" gorm:"foreignKey:AdvisorCode;references:InstructorCode"`
}

func NewStudentInfo(Stu model.Student, Gender string, CitizenID string, PhoneNumber string, advisorCode string) *StudentInfo {
	return &StudentInfo{
		Student:     Stu,
		Gender:      Gender,
		CitizenID:   CitizenID,
		PhoneNumber: PhoneNumber,
		AdvisorCode: advisorCode,
	}
}

func NewUpdatedStudentInfo(
	studentInfo *StudentInfo,
	firstName, lastName, gender, citizenID, phoneNumber, email string,
) *StudentInfo {
	return &StudentInfo{
		Student: model.Student{
			StudentCode: studentInfo.StudentCode,
			FirstName:   util.IfNotEmpty(firstName, studentInfo.FirstName),
			LastName:    util.IfNotEmpty(lastName, studentInfo.LastName),
			Email:       util.IfNotEmpty(email, studentInfo.Email),
			StartDate:   studentInfo.StartDate,
			BirthDate:   studentInfo.BirthDate,
			Program:     studentInfo.Program,
			Department:  studentInfo.Department,
			Status:      studentInfo.Status,
		},
		Gender:      util.IfNotEmpty(gender, studentInfo.Gender),
		CitizenID:   util.IfNotEmpty(citizenID, studentInfo.CitizenID),
		PhoneNumber: util.IfNotEmpty(phoneNumber, studentInfo.PhoneNumber),
		AdvisorCode: studentInfo.AdvisorCode,
	}
}

func (StudentInfo) TableName() string {
	return "student_infos"
}

func (s *StudentInfo) Validate() error {
	modelValidator := validation.NewModelValidator()

	if err := modelValidator.ModelValidate(s); err != nil {
		return err
	}

	return nil
}
