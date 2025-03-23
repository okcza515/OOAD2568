// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type ItemRequestDetail struct {
	ItemRequestDetailID uint `gorm:"primaryKey"` // PK
	Quantity            uint `gorm:"not null"`
	ItemID              uint `gorm:"foreignKey:ItemID"` //TO-DO: dicuss with MEP-1012 asset about the itemid

	ItemName      string         `gorm:"type:varchar(255);not null"`
	UnitPrice     float64        `gorm:"not null"`
	TotalPrice    float64        `gorm:"not null"`
	Specification string         `gorm:"type:text"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
