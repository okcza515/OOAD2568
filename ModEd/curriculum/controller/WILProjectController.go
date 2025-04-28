// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectController struct {
	*core.BaseController[model.WILProject]
	Connector *gorm.DB
}

type WILProjectControllerInterface interface {
	RetrieveAllWILProjects() ([]model.WILProject, error)
	Insert(data model.WILProject) error
	UpdateByID(data model.WILProject) error
	RetrieveByID(id uint, preloads ...string) (*model.WILProject, error)
	DeleteByID(id uint) error
}

func CreateWILProjectController(connector *gorm.DB) *WILProjectController {
	return &WILProjectController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.WILProject](connector),
	}
}

func (controller *WILProjectController) RetrieveAllWILProjects() ([]model.WILProject, error) {
	var wilProjects []model.WILProject

	if err := controller.Connector.Find(&wilProjects).Error; err != nil {
		return nil, err
	}

	return wilProjects, nil
}
