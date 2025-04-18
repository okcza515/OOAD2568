// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type ItemDetail struct {
	ItemDetailID  uint   `gorm:"primaryKey"` // Auto-increment ID
	ItemRequestID uint   `gorm:"not null"`   // Foreign key to ItemRequest
	Name          string `gorm:"type:varchar(255);not null"`
	Quantity      int    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"` // Soft delete
}
