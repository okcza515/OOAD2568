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

	if adapter.WILProject.ProjectName == "" {
		return nil, errors.New("ProjectName is empty in WILProject")
	}

	return &seniorProjectModel.SeniorProject{
		Model: gorm.Model{
			ID:        adapter.WILProject.ID,
			CreatedAt: adapter.WILProject.CreatedAt,
			UpdatedAt: adapter.WILProject.UpdatedAt,
			DeletedAt: adapter.WILProject.DeletedAt,
		},
		GroupName: adapter.WILProject.ProjectName,
	}, nil
}
