// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Quotation struct {
	QuotationID uint            `gorm:"primaryKey"`
	Status      QuotationStatus `gorm:"type:varchar(50);default:'draft'"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
	TORID       uint            `gorm:"index"`
	SupplierID  uint            `gorm:"index"`
	TOR         TOR
	Supplier    Supplier
}
