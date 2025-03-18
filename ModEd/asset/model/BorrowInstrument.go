package model

import (
	"time"

	"github.com/google/uuid"
)

type BorrowInstrument struct {
	BorrowID           uuid.UUID        `gorm:"type:text;primaryKey;" json:"borrow_id" csv:"borrow_id"`
	StaffUserID        uuid.UUID        `gorm:"type:text;not null" json:"staff_user_id" csv:"staff_user_id"`
	BorrowUserID       uuid.UUID        `gorm:"type:text;not null" json:"borrow_user_id" csv:"borrow_user_id"`
	ExpectedReturnDate *time.Time       `gorm:"type:timestamp" json:"expected_return_date,omitempty" csv:"expected_return_date,omitempty"`
	BorrowDate         time.Time        `gorm:"type:timestamp;not null" json:"borrow_date" csv:"borrow_date"`
	ReturnID           *uuid.UUID       `gorm:"type:text" json:"return_id,omitempty" csv:"return_id,omitempty"`
	Description        *string          `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	BorrowObjective    string           `gorm:"type:text;not null" json:"borrow_objective" csv:"borrow_objective"`
	ReturnInstrument   ReturnInstrument `gorm:"foreignKey:ReturnID;references:ReturnID;constraint:OnUpdate:CASCADE;"`
}
