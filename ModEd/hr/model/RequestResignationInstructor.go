package model

import (
	"gorm.io/gorm"
)

type RequestResignationInstructor struct {
	gorm.Model
	InstructorCode string `gorm:"not null"`        // อ้างถึง Instructor.InstructorID
	Reason         string `gorm:"type:text"`       // optional เหตุผลลาออก
	Status         string `gorm:"default:Pending"` // Pending / Approved / Rejected
}
