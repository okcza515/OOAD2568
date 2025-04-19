// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemDetail struct {
	ItemDetailID  uint           `gorm:"primaryKey"`
	ItemRequestID uint           `gorm:"not null"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Quantity      int            `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
