// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"
)

type WILProjectApplication struct {
	core.BaseModel
	ProjectName       string                 `gorm:"not null" validation:"required"`
	ProjectDetail     string                 `gorm:"not null" validation:"required"`
	Semester          string                 `gorm:"not null" validation:"required"`
	CompanyId         uint                   `gorm:"not null" validation:"required,uint"`
	Mentor            string                 `gorm:"not null" validation:"required"`
	Students          []WILProjectMember     `gorm:"foreignKey:WILProjectApplicationId"`
	AdvisorId         uint                   `json:"AdvisorId" validation:"required,uint"`
	Advisor           commonModel.Instructor `json:"Advisor" gorm:"foreignKey:AdvisorId;references:InstructorCode"`
	ApplicationStatus string                 `gorm:"not null" validation:"required"`
	TurninDate        *time.Time             `gorm:"default:null" json:"TurninDate"`
}

func (application WILProjectApplication) ToString() string {
	return fmt.Sprintf("[%v - %v]\t%v\t%v | %v", application.ID, application.ApplicationStatus, application.ProjectName, application.Mentor, application.TurninDate)
}
