// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	seniorProjectModel "ModEd/project/model"
	"errors"

	"gorm.io/gorm"
)

type WILProjectApplicationToSeniorProjectAdapter struct {
	WILProject *WILProjectApplication
}

func (adapter *WILProjectApplicationToSeniorProjectAdapter) ToSeniorProject() (*seniorProjectModel.SeniorProject, error) {
	if adapter.WILProject == nil {
		return nil, errors.New("WILProject is nil")
	}

	return &seniorProjectModel.SeniorProject{
		Model: gorm.Model{
			CreatedAt: adapter.WILProject.CreatedAt,
			UpdatedAt: adapter.WILProject.UpdatedAt,
			DeletedAt: adapter.WILProject.DeletedAt,
		},
		GroupName: adapter.WILProject.ProjectName,
		Members: []seniorProjectModel.GroupMember{},
		Advisors: []seniorProjectModel.Advisor{
			{
				
			},
		},
		Committees: []seniorProjectModel.Committee{},
		Assignments: []seniorProjectModel.Assignment{},
		Presentations: []seniorProjectModel.Presentation{},
		Reports: []seniorProjectModel.Report{},
		Assessments: []seniorProjectModel.Assessment{},
	}, nil
}
