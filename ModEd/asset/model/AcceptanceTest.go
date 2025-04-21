// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type AcceptanceTest struct {
	AcceptanceTestID     uint           `gorm:"primaryKey"`
	ProcurementID        uint           `gorm:"foreignKey:ProcurementID"`
	TORID                uint           `gorm:"foreignKey:TORID"`
	AcceptanceCriteriaID uint           `gorm:"foreignKey:AcceptanceCriteriaID"`
	AcceptanceApprovalID uint           `gorm:"foreignKey:AcceptanceApprovalID"`
	Results              string         `gorm:"type:text"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
