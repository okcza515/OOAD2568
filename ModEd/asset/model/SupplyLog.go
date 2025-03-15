package model

import (
	"time"

	"github.com/google/uuid"
)

type SupplyLog struct {
	LogID       uuid.UUID       `gorm:"primaryKey;type:text;" json:"log_id" csv:"log_id"`
	Timestamp   time.Time       `gorm:"type:timestamp;not null" json:"timestamp" csv:"timestamp"`
	RefUserID   *uuid.UUID      `gorm:"type:text" json:"ref_user_id,omitempty" csv:"ref_user_id,omitempty"`
	StaffUserID uuid.UUID       `gorm:"type:text;not null" json:"staff_user_id" csv:"staff_user_id"`
	Action      SupplyLogAction `gorm:"not null" json:"action" csv:"action"`
	SupplyID    uuid.UUID       `gorm:"type:text;not null" json:"supply_id" csv:"supply_id"`
	Description *string         `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	Quantity    int             `gorm:"not null;" json:"quantity" csv:"quantity"`
}
