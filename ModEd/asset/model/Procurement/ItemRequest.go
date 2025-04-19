// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type ItemRequest struct {
	ItemRequestID uint              `gorm:"primaryKey"`
	DepartmentID  uint              `gorm:"type:varchar(255);not null"`
	Status        ItemRequestStatus `gorm:"type:varchar(50);default:'draft'"`
	Items         []ItemDetail      `gorm:"foreignKey:ItemRequestID"`
	DeleteAt      gorm.DeletedAt    `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Departments   master.Department
}
