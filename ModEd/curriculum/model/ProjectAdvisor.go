//MEP-1008
package model

import (
	commonModel "ModEd/common/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectAdvisor struct {
	gorm.Model
	AdvisorId      int                          `gorm:"type:int;not null"`
	Advisor        commonModel.Instructor       `gorm:"foreignKey:AdvisorId"`
	SeniorProjects []projectModel.SeniorProject `gorm:"foreignKey:SeniorProjectId"`
}
