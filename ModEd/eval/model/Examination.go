package model

import (
	curriculumModel "ModEd/curriculum/model"
	"time"

	"gorm.io/gorm"
)

type Examination struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Exam_name   string
	Course      curriculumModel.Course
	Curriculum  curriculumModel.Curriculum
	Criteria    string
	Description string
	Exam_date   time.Time
	Create_at   time.Time
}
