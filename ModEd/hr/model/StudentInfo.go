package model

import (
	"ModEd/common/model"
	"ModEd/hr/util"
)

type StudentInfo struct {
	model.Student
	Gender      string `csv:"Gender" json:"Gender"`
	CitizenID   string `csv:"CitizenID" json:"CitizenID"`
	PhoneNumber string `csv:"PhoneNumber" json:"PhoneNumber"`
	// Advisor     model.Instructor `csv:"Advisor" json:"Advisor"`
	// Department  model.Department `csv:"Department" json:"Department"`
}

func NewStudentInfo(StudentCode string, Gender string, CitizenID string, PhoneNumber string) *StudentInfo {
	return &StudentInfo{
		Student: model.Student{
			StudentCode: StudentCode,
		},
		Gender:      Gender,
		CitizenID:   CitizenID,
		PhoneNumber: PhoneNumber,
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
	}
}
