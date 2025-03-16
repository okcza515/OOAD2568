package model

import (
	"time"

	"github.com/google/uuid"
)

type ReturnInstrument struct {
	ReturnID     uuid.UUID  `gorm:"type:text;primaryKey;" json:"return_id" csv:"return_id"`
	StaffUserID  uuid.UUID  `gorm:"type:text;not null" json:"staff_user_id" csv:"staff_user_id"`
	ReturnUserID uuid.UUID  `gorm:"type:text;not null" json:"return_user_id" csv:"return_user_id"`
	ReturnDate   *time.Time `gorm:"type:timestamp" json:"return_date,omitempty" csv:"return_date,omitempty"`
	Description  *string    `gorm:"type:text;" json:"description,omitempty" csv:"description,omitempty"`
}
