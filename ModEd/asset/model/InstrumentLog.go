package model

import (
	"github.com/google/uuid"
	"time"
)

type InstrumentLog struct {
	LogID        uuid.UUID               `gorm:"type:uuid;default:gen_random_uuid();primaryKey;not null;unique" json:"log_id" csv:"log_id"`
	Timestamp    time.Time               `gorm:"type:timestamp;not null" json:"timestamp" csv:"timestamp"`
	RefUserID    *uuid.UUID              `gorm:"type:uuid" json:"ref_user_id,omitempty" csv:"ref_user_id"`
	StaffUserID  uuid.UUID               `gorm:"type:uuid;not null" json:"staff_user_id" csv:"staff_user_id"`
	Action       InstrumentLogActionEnum `gorm:"type:varchar;not null" json:"action" csv:"action"`
	InstrumentID uuid.UUID               `gorm:"type:uuid;not null" json:"instrument_id" csv:"instrument_id"`
	Description  string                  `gorm:"type:varchar;not null" json:"description" csv:"description"`
	RefBorrowID  *uuid.UUID              `gorm:"type:uuid" json:"ref_borrow_id,omitempty" csv:"ref_borrow_id"`
}
