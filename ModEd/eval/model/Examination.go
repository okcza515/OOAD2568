// MEP-1007
package model

import (
	instructorModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
	"time"
	"ModEd/core"
)

type Examination struct {
	core.BaseModel
	ID          		uint 							`gorm:"primaryKey" csv:"id" json:"id"`
	ExamName   		    string 							`gorm:"not null" csv:"exam_name" json:"exam_name"`
	InstructorID 		uint 							`gorm:"not null" csv:"instructor_id" json:"instructor_id"`
	Instructor  		instructorModel.Instructor 		`gorm:"foreignKey:Instructor_id;references:ID" csv:"-" json:"-"`
	CourseID    		uint            				`gorm:"not null" csv:"course_id" json:"course_id"`
	Course      		curriculumModel.Course 			`gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Description 		string 							`gorm:"not null" csv:"description" json:"description"`
	ExamStatus          ExamStatus    					`gorm:"not null" csv:"exam_status" json:"exam_status"` // draft, published, closed
	Attempt  			uint     						`gorm:"not null" csv:"attempt" json:"attempt"`
	StartDate   		time.Time 						`gorm:"not null" csv:"start_date" json:"start_date"`
	EndDate   			time.Time 						`gorm:"not null" csv:"end_date" json:"end_date"`
}