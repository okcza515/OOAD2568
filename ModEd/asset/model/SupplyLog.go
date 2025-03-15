package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplyLog struct {
	gorm.Model
	LogID       uuid.UUID       `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null;unique" json:"log_id,omitempty" csv:"log_id,omitempty"`
	Timestamp   time.Time       `gorm:"autoCreateTime;not null" json:"timestamp,omitempty" csv:"timestamp,omitempty"`
	RefUserID   uuid.UUID       `gorm:"type:uuid" json:"ref_user_id,omitempty" csv:"ref_user_id,omitempty"`
	StaffUserID uuid.UUID       `gorm:"type:uuid;not null" json:"staff_user_id,omitempty" csv:"staff_user_id,omitempty"`
	Action      SupplyLogAction `gorm:"not null" json:"action,omitempty" csv:"action,omitempty"`
	SupplyID    uuid.UUID       `gorm:"type:uuid;not null" json:"supply_id,omitempty" csv:"supply_id,omitempty"`
	Description string          `json:"description,omitempty" csv:"description,omitempty"`
	Quantity    int             `json:"quantity,omitempty" csv:"quantity,omitempty"`
	RefBorrowID uuid.UUID       `gorm:"type:uuid" json:"ref_borrow_id,omitempty" csv:"ref_borrow_id,omitempty"`
}
