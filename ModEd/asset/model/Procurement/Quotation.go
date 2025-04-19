// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Quotation struct {
	QuotationID uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	ContactInfo string         `gorm:"type:text"`
	Performance float64        `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	SupplierID  uint           `gorm:"index"`
	Supplier    Supplier
}
