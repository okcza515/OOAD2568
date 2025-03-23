// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type TOR struct {
	TORID        uint               `gorm:"primaryKey"`
	SupplierID   uint               `gorm:"foreignKey:SupplierID"`
	Scope        string             `gorm:"type:text"`
	Deliverables AcceptanceCriteria `gorm:"type:text"` //to-do: make it work properly.
	Timeline     string             `gorm:"type:text"`
	DeletedAt    gorm.DeletedAt     `gorm:"index"`
}
