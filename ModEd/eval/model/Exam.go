// MEP-1007
package model

import (
	instructorModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
	"time"
	"ModEd/core"
)

type Exam struct {
	core.BaseModel
	ExamName   		    string 							`gorm:"type:text;not null" json:"exam_name" csv:"exam_name"`
	InstructorID 		uint 							`gorm:"type:integer;not null" json:"instructor_id" csv:"instructor_id"`
	Instructor  		instructorModel.Instructor 		`gorm:"foreignKey:InstructorID;references:ID" csv:"instructor" json:"instructor"`
	ClassID    			uint            				`gorm:"type:integer;not null" json:"class_id" csv:"class_id"`
	Class      			curriculumModel.Class 			`gorm:"foreignKey:ClassID;references:ID" csv:"class" json:"class"`
	Description 		string 							`gorm:"type:text;not null" json:"description" csv:"description"`
	ExamStatus          ExamStatus    					`gorm:"type:text;not null" json:"exam_status" csv:"exam_status"` // draft, published, closed
	Attempt  			uint     						`gorm:"type:integer;not null" json:"attempt" csv:"attempt"`
	StartDate   		time.Time 						`gorm:"type:timestamp;not null" json:"start_date" csv:"start_date"`
	EndDate   			time.Time 						`gorm:"type:timestamp;not null" json:"end_date" csv:"end_date"`
}