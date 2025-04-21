package model

import (
	"gorm.io/gorm"
)

type RequestResignationStudent struct {
	gorm.Model
	StudentCode string `gorm:"not null"`        // อ้างถึง StudentInfo.StudentCode
	Reason      string `gorm:"type:text"`       // optional เหตุผลลาออก
	Status      string `gorm:"default:Pending"` // Pending / Approved / Rejected
}
