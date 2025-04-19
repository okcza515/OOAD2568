// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type AcceptanceCriteria struct {
	AcceptanceCriteriaID uint           `gorm:"primaryKey"`
	CriteriaName         string         `gorm:"type:varchar(255);not null"`
	Description          string         `gorm:"type:text"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
