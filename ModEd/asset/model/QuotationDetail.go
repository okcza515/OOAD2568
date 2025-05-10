// MEP-1014
package model

type QuotationDetail struct {
	QuotationDetailID uint     `gorm:"primaryKey"`
	QuotationID       uint     `gorm:"index"`
	InstrumentLabel   string   `gorm:"not null"`
	Description       *string  `gorm:"type:text"`
	CategoryID        uint     `gorm:"not null"`
	Category          Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE;"`
	Quantity          int      `gorm:"not null"`
	OfferedPrice      float64  `gorm:"type:decimal(10,2);not null"`
}
