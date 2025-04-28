package model

import (
	"ModEd/common/model"
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

func (s *StudentInfo) SetStudent(commomStudent model.Student) *StudentInfo {
	s.FirstName = commomStudent.FirstName
	s.LastName = commomStudent.LastName
	s.Email = commomStudent.Email
	s.StartDate = commomStudent.StartDate
	s.BirthDate = commomStudent.BirthDate
	s.Program = commomStudent.Program
	s.Department = commomStudent.Department
	s.Status = commomStudent.Status
	return s
}

func (s *StudentInfo) SetFirstName(firstName string) *StudentInfo {
	s.FirstName = firstName
	return s
}

func (s *StudentInfo) SetLastName(lastName string) *StudentInfo {
	s.LastName = lastName
	return s
}

func (s *StudentInfo) SetGender(gender string) *StudentInfo {
	s.Gender = gender
	return s
}

func (s *StudentInfo) SetCitizenID(citizenID string) *StudentInfo {
	s.CitizenID = citizenID
	return s
}

func (s *StudentInfo) SetPhoneNumber(phoneNumber string) *StudentInfo {
	s.PhoneNumber = phoneNumber
	return s
}

func (s *StudentInfo) SetEmail(email string) *StudentInfo {
	s.Email = email
	return s
}
