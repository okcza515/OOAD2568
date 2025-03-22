package model

import (
	CommonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type InternStudent struct {
	gorm.Model
	InternStatus InternStatus        `gorm:"type:varchar(20)"`
	StudentCode  string              `gorm:"not null;unique" csv:"student_code"`
	Student      CommonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
