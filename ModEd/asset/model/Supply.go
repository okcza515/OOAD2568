package model

// MEP-1012 Asset

import (
	"ModEd/core"
	"fmt"
)

type Supply struct {
	core.BaseModel
	SupplyLabel string `gorm:"not null"`
	Description *string
	Location    *string
	CategoryID  *uint
	Quantity    int         `gorm:"not null"`
	SupplyLog   []SupplyLog `gorm:"foreignKey:ID;references:ID"`
}

func (supp Supply) ToString() string {
	return fmt.Sprintf("[%d] \t%v\t%v", supp.ID, supp.Quantity, supp.SupplyLabel)
}
