package model

import (
    "errors"
    seniorProjectModel "ModEd/project/model"

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
        SeniorProjectId: adapter.WILProject.ID,
        GroupName:       adapter.WILProject.ProjectName,
    }, nil
}