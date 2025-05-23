// MEP-1014
package model

import (
	"ModEd/core/validation"
	"time"

	"gorm.io/gorm"
)

type TOR struct {
	TORID               uint           `gorm:"primaryKey"`
	InstrumentRequestID uint           `gorm:"index"`
	Scope               string         `gorm:"type:text;not null" validation:"not null"`
	Deliverables        string         `gorm:"type:text" validation:"not null"`
	Timeline            string         `gorm:"type:text" validation:"not null"`
	Committee           string         `gorm:"type:text" validation:"not null"`
	Quotations          []Quotation    `gorm:"foreignKey:TORID"`
	TotalPrice          float64        `gorm:"type:decimal(12,2);default:0"`
	Status              TORStatus      `gorm:"type:varchar(50);default:'announced'"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstrumentRequest   InstrumentRequest
}

func (tor *TOR) Validate() error {
	validator := validation.NewModelValidator()

	if err := validator.ModelValidate(tor); err != nil {
		return err
	}

	return nil
}
