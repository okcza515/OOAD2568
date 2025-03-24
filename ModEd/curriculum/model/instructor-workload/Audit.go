package model

import (
	commonModel "ModEd/common/model"
)

type Audit struct {
	CreatedBy commonModel.Instructor `gorm:"foreignKey:CreatedBy; not null"`
	UpdatedBy commonModel.Instructor `gorm:"foreignKey:UpdatedBy; not null"`
	DeletedBy commonModel.Instructor `gorm:"foreignKey:DeletedBy"`
}
