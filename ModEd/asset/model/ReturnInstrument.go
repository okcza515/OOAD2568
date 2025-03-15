package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReturnInstrument struct {
	gorm.Model
	ReturnID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StaffUserID  uuid.UUID `gorm:"type:uuid;not null"`
	ReturnUserID uuid.UUID `gorm:"type:uuid;not null"`
	ReturnDate   *time.Time
	Description  uuid.UUID `gorm:"type:uuid;not null"`
}
