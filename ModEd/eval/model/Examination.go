package model

import (
	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
	"time"

	"gorm.io/gorm"
)

type QuestionType string
const (
	MULTIPLECHOICE		QuestionType="Multiple_choice"
	SHORTANSWER			QuestionType="Short_answer"
	TRUEFALSE			QuestionType="True_false"
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

type Examination struct {
	gorm.Model
	ID 				uint					`gorm:"primaryKey"`
	// Instructor_id model.Instructor
	Exam_name 		string
	Course 			curriculumModel.Course
	Curriculum 		curriculumModel.Curriculum
	Criteria 		string
	Description		string
	Exam_date		time.Time
	Create_at		time.Time
}

type Question struct {
	gorm.Model
	ID 				uint					`gorm:"primaryKey"`
	Examination 	Examination
	Question_detail	string
	Question_type	QuestionType
	Correct_answer	string
	Score			float64	
}

type Answer struct {
	gorm.Model
	ID 				uint					`gorm:"primaryKey"`
	Question		Question
	Student			commonModel.Student
	TheAnswer		string
}

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