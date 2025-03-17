package model

import (
	CommonModel "ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternStudent struct {
	gorm.Model
	StudentID    uint                `gorm:"index"`
	Student      CommonModel.Student `gorm:"foreignKey:StudentID"`
	InternID     uuid.UUID           `gorm:"type:uuid"`
	InternStatus InternStatus        `gorm:"type:varchar(20)"`
}
