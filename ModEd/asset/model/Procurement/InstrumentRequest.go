// MEP-1014
package model

import (
	master "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type InstrumentRequest struct {
	InstrumentRequestID uint                    `gorm:"primaryKey"`
	DepartmentID        uint                    `gorm:"type:varchar(255);not null"`
	Status              InstrumentRequestStatus `gorm:"type:varchar(50);default:'draft'"`
	Instruments         []InstrumentDetail      `gorm:"index"`
	DeleteAt            gorm.DeletedAt          `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Departments         master.Department
}
