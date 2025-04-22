// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"

	"gorm.io/gorm"
)

type WILProjectController struct {
	*core.BaseController
	Connector *gorm.DB
}

type WILProjectControllerInterface interface {
	RegisterWILProjects(projects []core.RecordInterface)
	Insert(data core.RecordInterface) error
	UpdateByID(data core.RecordInterface) error
	RetrieveByID(id uint, preloads ...string) (*core.RecordInterface, error)
	DeleteByID(id uint) error
	ListPagination(condition map[string]interface{}, page int, pageSize int)
}

func CreateWILProjectController(connector *gorm.DB) *WILProjectController {
	return &WILProjectController{
		Connector:      connector,
		BaseController: core.NewBaseController("WILProject", connector),
	}
}

func (repo WILProjectController) RegisterWILProjects(projects []core.RecordInterface) {
	for _, project := range projects {
		repo.Insert(project)
	}
}
