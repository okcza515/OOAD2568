// Sawitt Ngamvilaisiriwong 65070503469

package model

import (
	"ModEd/common/model"

	"gorm.io/gorm"

	"time"
)

type Progress struct {
	gorm.Model
	SID        string        `gorm:"not null"` // รหัสนักศึกษา
	Student    model.Student `gorm:"foreignKey:SID"`
	Title      string        // ชื่อ Assignment
	Assignment Assignment    `gorm:"foreignKey:Title"`
	Status     string        // Status: Submitted, Late, Not Submitted
	LastUpdate time.Time     `gorm:"autoUpdateTime"` // Update ล่าสุด
	SubmitNum  uint          // จำนวนคนส่ง
}
