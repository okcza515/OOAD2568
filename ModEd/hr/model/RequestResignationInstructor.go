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

func (r RequestResignationInstructor) GetID() string {
	return r.InstructorCode
}

func (r RequestResignationInstructor) GetReason() string {
	return r.Reason
}

func (r RequestResignationInstructor) GetStatus() string {
	return r.Status
}