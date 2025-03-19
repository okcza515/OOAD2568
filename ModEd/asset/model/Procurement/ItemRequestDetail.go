//MEP-1014
package model

import (
	"github.com/google/uuid"
)

type ItemRequestDetail struct {
	ItemRequestDetailID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	Quantity            int       `gorm:"not null"`
	ItemID              uuid.UUID `gorm:"type:uuid"` //TO-DO: dicuss with MEP-1012 asset about the itemid

	ItemName      string  `gorm:"type:varchar(255);not null"`
	UnitPrice     float64 `gorm:"not null"`
	TotalPrice    float64 `gorm:"not null"`
	Specification string  `gorm:"type:text"`
}
