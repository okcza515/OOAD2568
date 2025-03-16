package model

import "ModEd/common/model"

type StudentInfo struct {
	model.Student
	Gender      string `csv:"Gender"`
	CitizenID   string `csv:"CitizenID"`
	PhoneNumber string `csv:"PhoneNumber"`
}
