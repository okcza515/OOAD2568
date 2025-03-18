package model

import (
	"github.com/google/uuid"
	"time"
)

type InstrumentLog struct {
	LogID            uuid.UUID               `gorm:"type:text;primaryKey" json:"log_id" csv:"log_id"`
	Timestamp        time.Time               `gorm:"type:timestamp;not null" json:"timestamp" csv:"timestamp"`
	RefUserID        *uuid.UUID              `gorm:"type:text" json:"ref_user_id,omitempty" csv:"ref_user_id"`
	StaffUserID      uuid.UUID               `gorm:"type:text;not null" json:"staff_user_id" csv:"staff_user_id"`
	Action           InstrumentLogActionEnum `gorm:"type:text;not null" json:"action" csv:"action"`
	InstrumentID     uuid.UUID               `gorm:"type:text;not null" json:"instrument_id" csv:"instrument_id"`
	Description      string                  `gorm:"type:text;not null" json:"description" csv:"description"`
	RefBorrowID      *uuid.UUID              `gorm:"type:text" json:"ref_borrow_id,omitempty" csv:"ref_borrow_id"`
	BorrowInstrument BorrowInstrument        `gorm:"foreignKey:RefBorrowID;references:BorrowID;constraint:OnUpdate:CASCADE;"`
}
