// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type InstrumentRequest struct {
	InstrumentRequestID uint                    `gorm:"primaryKey"`
	Status              InstrumentRequestStatus `gorm:"type:varchar(50);default:'pending'"`
	Instruments         []InstrumentDetail      `gorm:"foreignKey:InstrumentRequestID"`
	TotalEstimatedPrice float64                 `gorm:"type:decimal(12,2);default:0"`
	DeleteAt            gorm.DeletedAt          `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DepartmentID        uint
	IsLinkedToTOR       bool `gorm:"default:false"`
}
