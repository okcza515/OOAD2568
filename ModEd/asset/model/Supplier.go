// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Supplier struct {
	SupplierID  uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	ContactInfo string         `gorm:"type:text"`
	Performance float64        `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Quotations  []Quotation
}
