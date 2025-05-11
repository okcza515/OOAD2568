// MEP-1014
package model

import (
	"time"
	"ModEd/core/validation"
	"gorm.io/gorm"
)

type InstrumentDetail struct {
	InstrumentDetailID  uint     `gorm:"primaryKey"`
	InstrumentRequestID uint     `gorm:"index"`
	InstrumentLabel     string   `gorm:"not null" validation:"not null"`
	Description         *string  `gorm:"type:text"`
	CategoryID          uint     `gorm:"type:text;not null" validation:"not null"`
	Category            Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE;"`
	Quantity            int      `gorm:"not null" validation:"not null"`
	EstimatedPrice      float64  `gorm:"type:decimal(10,2);default:0" validation:"not null"` 
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	InstrumentRequest InstrumentRequest
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (instrumentdetail *InstrumentDetail) Validate() error {
	validator := validation.NewModelValidator()

	if err := validator.ModelValidate(instrumentdetail); err != nil {
		return err
	}

	return nil
}