// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type Procurement struct {
	ProcurementID uint              `gorm:"primaryKey"`
	TORID         uint              `gorm:"index"`
	ApproverID    *uint             `gorm:"index"`
	Approver      master.Instructor `gorm:"foreignKey:ApproverID"`
	Status        ProcurementStatus `gorm:"type:varchar(50);default:'pending'"`
	DeletedAt     gorm.DeletedAt    `gorm:"index"`
	ApprovalTime  *time.Time        `gorm:"type:time"`
	CreatedAt     time.Time
	TOR           TOR
}
