// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type Quotation struct {
	QuotationID uint           `gorm:"primaryKey"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	TORID       uint           `gorm:"index"`
	SupplierID  uint           `gorm:"index"`
	TOR         TOR
	Supplier    Supplier
}
