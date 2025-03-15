package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BorrowInstrument struct {
	gorm.Model
	BorrowID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StaffUserID        uuid.UUID `gorm:"type:uuid;not null"`
	BorrowUserID       uuid.UUID `gorm:"type:uuid;not null"`
	ExpectedReturnDate time.Time
	BorrowDate         time.Time `gorm:"not null"`
	ReturnID           uuid.UUID
	Description        string
	BorrowObjective    string `gorm:"not null"`
}
