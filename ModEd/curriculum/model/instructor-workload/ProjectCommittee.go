package model

import (
	commonModel "ModEd/common/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectCommittee struct {
	gorm.Model
	committeeId    int                          `gorm:"type:int;not null"`
	Committee      commonModel.Instructor       `gorm:"foreignKey:committeeId"`
	SeniorProjects []projectModel.SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
