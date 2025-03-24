package asset

import (
	"gorm.io/gorm"
)

type InstrumentLog struct {
	gorm.Model
	RefUserID        *uint
	StaffUserID      uint                    `gorm:"type:integer;not null"`
	Action           InstrumentLogActionEnum `gorm:"not null"`
	InstrumentID     uint                    `gorm:"type:integer;not null"`
	Description      string                  `gorm:"not null"`
	RefBorrowID      *uint                   `gorm:"type:integer"`
	BorrowInstrument BorrowInstrument        `gorm:"foreignKey:RefBorrowID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Instrument       Instrument              `gorm:"foreignKey:InstrumentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
