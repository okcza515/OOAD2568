//MEP-1008
package model

import (
	commonModel "ModEd/common/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectCommittee struct {
	gorm.Model
	CommitteeId    int                          `gorm:"type:int;not null"`
	Committee      commonModel.Instructor       `gorm:"foreignKey:CommitteeId"`
	SeniorProjects []projectModel.SeniorProject `gorm:"many2many:project_committee_senior_projects"`
}
