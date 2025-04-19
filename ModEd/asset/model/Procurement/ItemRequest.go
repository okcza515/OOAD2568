// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type ItemRequest struct {
	ItemRequestID uint              `gorm:"primaryKey"` // Auto-increment ID
	Status        ItemRequestStatus `gorm:"type:varchar(50);default:'draft'"`
	Items         []ItemDetail      `gorm:"foreignKey:ItemRequestID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeleteAt      gorm.DeletedAt `gorm:"index"`                      // Soft delete
	DepartmentID  uint           `gorm:"type:varchar(255);not null"` // Foreign key (uint from gorm.Model)
	Departments   master.Department
}
