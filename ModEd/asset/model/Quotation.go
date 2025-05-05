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
	Status      QuotationStatus `gorm:"type:varchar(50);default:'pending'"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	Supplier    Supplier
}
