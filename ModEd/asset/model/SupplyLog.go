package model

import (
	"time"

	"gorm.io/gorm"
)

type SupplyLog struct {
	gorm.Model
	LogID       string          `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;unique"`
	Timestamp   time.Time       `gorm:"autoCreateTime;not null"`
	RefUserID   string          `gorm:"type:uuid"`
	StaffUserID string          `gorm:"type:uuid;not null"`
	Action      SupplyLogAction `gorm:"not null"`
	SupplyID    string          `gorm:"type:uuid;not null"`
	Description string
	Quantity    int
	RefBorrowID string `gorm:"type:uuid"`
}
