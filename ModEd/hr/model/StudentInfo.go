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
	Department  string           `csv:"Department"`
	AdvisorID   uint             `csv:"AdvisorID" json:"AdvisorID"`          // Foreign key field
	Advisor     model.Instructor `csv:"Advisor" gorm:"foreignKey:AdvisorID"` // Association using AdvisorIDs
}
