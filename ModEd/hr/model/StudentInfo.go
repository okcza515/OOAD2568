package model

import (
	"ModEd/common/model"
)

type StudentInfo struct {
	model.Student
	Gender      string           `csv:"Gender" json:"Gender"`
	CitizenID   string           `csv:"CitizenID" json:"CitizenID"`
	PhoneNumber string           `csv:"PhoneNumber" json:"PhoneNumber"`
	Year        int              `csv:"Year"`
	Advisor     model.Instructor `csv:"Advisor" json:"Advisor"`
	Department  model.Department `csv:"Department" json:"Department"`
}
