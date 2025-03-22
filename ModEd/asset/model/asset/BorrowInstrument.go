// MEP-1012
package asset

import (
	"gorm.io/gorm"
	"time"
)

type BorrowInstrument struct {
	gorm.Model
	BorrowID           uint       `gorm:"primaryKey;not null;unique"`
	StaffUserID        uint       `gorm:"not null"`
	BorrowUserID       uint       `gorm:"not null"`
	ExpectedReturnDate *time.Time `gorm:"type:timestamp"`
	BorrowDate         time.Time  `gorm:"type:timestamp;not null"`
	ReturnDate         time.Time  `gorm:"type:timestamp;not null"`
	InstrumentID       uint       `gorm:"not null"`
	Description        *string    `gorm:"type:text"`
	BorrowObjective    string     `gorm:"type:text;not null"`
}
