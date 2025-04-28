package model

import (
	"ModEd/common/model"
)

type StudentInfo struct {
	model.Student
	Gender      string           `csv:"Gender" json:"Gender"`
	CitizenID   string           `csv:"CitizenID" json:"CitizenID"`
	PhoneNumber string           `csv:"PhoneNumber" json:"PhoneNumber"`
	// Advisor     model.Instructor `csv:"Advisor" json:"Advisor"`
	// Department  model.Department `csv:"Department" json:"Department"`
}

func NewStudentInfo(StudentCode string, Gender string, CitizenID string, PhoneNumber string) *StudentInfo {
	return &StudentInfo{
		Student: model.Student{
			StudentCode: StudentCode,
		},
		Gender: Gender,
		CitizenID: CitizenID,
		PhoneNumber: PhoneNumber,
	}
}
