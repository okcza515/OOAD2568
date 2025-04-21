package model

import (
	"time"
	"gorm.io/gorm"
)

type RequestLeaveInstructor struct {
	gorm.Model
	InstructorCode string    `gorm:"not null"` // รหัสนักเรียน
	Status    string    `gorm:"default:Pending"` // สถานะการลา (เช่น Pending, Approved, Rejected)
	LeaveType string    // ประเภทการลาออก (เช่น ลาออก, ลาเรียน, ลาอื่นๆ)
	Reason    string    // เหตุผลการลา
	LeaveDate time.Time // วันที่ลา
}
