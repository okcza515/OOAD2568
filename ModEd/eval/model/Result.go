package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StatusResult string
const (
	PENDING				StatusResult="Pending"
	SUCCESS				StatusResult="Success"
)

type GradeResult string
const (
	APLUS				GradeResult="A+"
	A					GradeResult="A"
	BPLUS				GradeResult="B+"
	B					GradeResult="B"
	CPLUS				GradeResult="C+"
	C					GradeResult="C"
	DPLUS				GradeResult="D+"
	D					GradeResult="D"
	F					GradeResult="F"
)

type Result struct {
	gorm.Model
	ID 				uint 					`gorm:"primaryKey"`
	Examination		Examination
	Student			commonModel.Student
	Status			StatusResult
	Grade			GradeResult
	Feedback		string
	Student_score	uint
}