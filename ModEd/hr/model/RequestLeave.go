package model

import (
	"time"
)

type RequestLeave struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID string    `gorm:"not null"` // รหัสนักเรียน
	Status    string    `gorm:"default:Pending"` // สถานะการลา (เช่น Pending, Approved, Rejected)
	LeaveType string    // ประเภทการลาออก (เช่น ลาออก, ลาเรียน, ลาอื่นๆ)
	Reason    string    // เหตุผลการลา
	Remarks   string    // หมายเหตุ
	ApprovedBy string    // ผู้อนุมัติ
	LeaveDate time.Time // วันที่ลา
	ApprovedAt time.Time // วันที่อนุมัติ
	CreatedAt time.Time
}
