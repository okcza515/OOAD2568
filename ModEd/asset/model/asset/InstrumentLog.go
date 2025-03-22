package asset

import (
	"gorm.io/gorm"
)

type InstrumentLog struct {
	gorm.Model
	RefUserID        *uint
	StaffUserID      uint                    `gorm:"not null"`
	Action           InstrumentLogActionEnum `gorm:"not null"`
	InstrumentID     uint                    `gorm:"not null"`
	Description      string                  `gorm:"not null"`
	RefBorrowID      *uint                   `gorm:"type:integer"`
	BorrowInstrument BorrowInstrument        `gorm:"foreignKey:RefBorrowID;references:BorrowID;constraint:OnUpdate:CASCADE;"`
}
