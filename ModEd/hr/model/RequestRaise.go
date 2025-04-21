package model

import (
	"gorm.io/gorm"
)

type RequestRaise struct {
	gorm.Model
	InstructorCode string `gorm:"not null"`        // อ้างถึง StudentInfo.StudentCode
	Reason         string `gorm:"type:text"`       // optional เหตุผลลาออก
	Status         string `gorm:"default:Pending"` // Pending / Approved / Rejected
	TargetSalary   int    `gorm:"not null"`        // เป้าหมายเงินเดือนที่ต้องการ
}
