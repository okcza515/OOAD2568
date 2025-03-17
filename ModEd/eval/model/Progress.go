// Sawitt Ngamvilaisiriwong 65070503469
// MEP-1006

package model

import (
	"ModEd/common/model"

	"gorm.io/gorm"

	"time"
)

type Progress struct {
	gorm.Model
	SID         model.Student `gorm:"foreignKey:SID"` // รหัสนักศึกษาและประกาศว่าตัวแปรเป็น foreign key
	Title       Assignment    // หัวข้อ assignment
	Status      string        // Status: Submitted, Late, Not Submitted
	LastUpdate  time.Time     `gorm:"autoUpdateTime"` // Update ล่าสุด
	TotalSubmit uint          // จำนวนคนส่ง
}
