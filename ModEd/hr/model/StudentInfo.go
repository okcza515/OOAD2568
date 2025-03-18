package model

import (
	"ModEd/common/model"
)

type StudentInfo struct {
	model.Student
	Gender      string `csv:"Gender" json:"Gender"`
	CitizenID   string `csv:"CitizenID" json:"CitizenID"`
	PhoneNumber string `csv:"PhoneNumber" json:"PhoneNumber"`
	Year        int    `csv:"Year"`
	// DepartmentID uuid.UUID        `csv:"DepartmentID" json:"DepartmentID"`
	// Department   model.Department `csv:"Department" gorm:"foreignKey:DepartmentID;references:DepartmentId"`
	// AdvisorID    string           `csv:"AdvisorID" json:"AdvisorID"`
	// Advisor      model.Instructor `csv:"Advisor" gorm:"foreignKey:AdvisorID;references:InstructorId"`
}
