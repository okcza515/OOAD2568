// MEP-1009 Student Internship
package model

import "ModEd/core"

type SupervisorReview struct {
	core.BaseModel
	InstructorScore int `gorm:"type:int"`
	MentorScore     int `gorm:"type:int"`
}
