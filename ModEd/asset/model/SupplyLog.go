package model

// MEP-1012 Asset

import (
	"ModEd/core"
	"fmt"
)

type SupplyLog struct {
	core.BaseModel
	RefUserID   *uint
	StaffUserID uint                `gorm:"not null"`
	Action      SupplyLogActionEnum `gorm:"not null"`
	SupplyID    uint                `gorm:"not null"`
	Description *string             `gorm:"not null"`
	Quantity    uint                `gorm:"not null"`
	Supply      Supply              `gorm:"foreignKey:SupplyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s SupplyLog) ToString() string {
	return fmt.Sprintf("[%v]\t%v\t%v\t%v", s.CreatedAt.Format("2006-01-02 15:04:05"), s.Action.String(), s.Supply.ID, s.Supply.SupplyLabel)
}
