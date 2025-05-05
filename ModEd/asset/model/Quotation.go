// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Quotation struct {
	QuotationID uint            `gorm:"primaryKey"`
	TORID       uint            `gorm:"index"`
	TOR         TOR             `gorm:"foreignKey:TORID"`
	SupplierID  uint            `gorm:"index"`
	Supplier    Supplier        `gorm:"foreignKey:SupplierID"`
	TotalOfferedPrice float64 `gorm:"type:decimal(12,2);default:0"`
	Status      QuotationStatus `gorm:"type:varchar(50);default:'pending'"`
	Details     []QuotationDetail `gorm:"foreignKey:QuotationID"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}
