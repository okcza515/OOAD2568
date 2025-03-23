package model

import (
	curriculumModel "ModEd/curriculum/model"
	instructorModel "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type Examination struct {
	gorm.Model
	ID          		uint 							`gorm:"primaryKey" csv:"id" json:"id"`
	Exam_name   		string 							`gorm:"not null" csv:"exam_name" json:"exam_name"`
	Instructor_id 		uint 							`gorm:"not null" csv:"instructor_id" json:"instructor_id"`
	Instructor  		instructorModel.Instructor 		`gorm:"foreignKey:Instructor_id;references:InstructorID" csv:"-" json:"-"`
	CourseId    		uint            				`gorm:"not null" csv:"course_id" json:"course_id"`
	Course      		curriculumModel.Course 			`gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	CurriculumId    	uint            				`gorm:"not null" csv:"curriculum_id" json:"curriculum_id"`
	Curriculum  		curriculumModel.Curriculum 		`gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	Criteria    		string 							`gorm:"not null" csv:"criteria" json:"criteria"`
	Description 		string 							`gorm:"not null" csv:"description" json:"description"`
	Exam_date   		time.Time 						`gorm:"not null" csv:"exam_date" json:"exam_date"`
	Create_at   		time.Time 						`gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
}
