package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplyLog struct {
	gorm.Model
	LogID       uuid.UUID       `gorm:"primaryKey; type:uuid; default:gen_random_uuid(); not null;unique" json:"log_id" csv:"log_id"`
	Timestamp   time.Time       `gorm:"autoCreateTime;not null" json:"timestamp" csv:"timestamp"`
	RefUserID   *uuid.UUID      `gorm:"type:uuid" json:"ref_user_id,omitempty" csv:"ref_user_id,omitempty"`
	StaffUserID uuid.UUID       `gorm:"type:uuid;not null" json:"staff_user_id" csv:"staff_user_id"`
	Action      SupplyLogAction `gorm:"not null" json:"action" csv:"action"`
	SupplyID    uuid.UUID       `gorm:"type:uuid;not null" json:"supply_id" csv:"supply_id"`
	Description *string         `json:"description,omitempty" csv:"description,omitempty"`
	Quantity    int             `gorm:"not null;" json:"quantity" csv:"quantity"`
}
