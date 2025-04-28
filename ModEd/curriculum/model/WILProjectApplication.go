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
	ProjectName       string                 `gorm:"not null"`
	ProjectDetail     string                 `gorm:"not null"`
	Semester          string                 `gorm:"not null"`
	CompanyId         uint                   `gorm:"not null"`
	Mentor            string                 `gorm:"not null"`
	Students          []WILProjectMember     `gorm:"foreignKey:WILProjectApplicationId"`
	AdvisorId         uint                   `json:"AdvisorId"`
	Advisor           commonModel.Instructor `json:"Advisor" gorm:"foreignKey:AdvisorId;references:InstructorCode"`
	ApplicationStatus string                 `gorm:"not null"`
	TurninDate        *time.Time             `gorm:"default:null" json:"TurninDate"`
}

func (application WILProjectApplication) ToString() string {
	return fmt.Sprintf("[%v - %v]\t%v\t%v | %v", application.ID, application.ApplicationStatus, application.ProjectName, application.Mentor, application.TurninDate)
}
