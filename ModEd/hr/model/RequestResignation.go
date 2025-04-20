package model

import (
	"time"
)

type RequestResignation struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID string    `gorm:"not null"` // อ้างถึง StudentInfo.StudentCode
	Reason    string    `gorm:"type:text"` // optional เหตุผลลาออก
	Status    string    `gorm:"default:Pending"` // Pending / Approved / Rejected
	CreatedAt time.Time
}
