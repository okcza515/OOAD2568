// MEP-1008

package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassMaterialControllerInterface interface {
	Insert(data model.ClassMaterial) error
	RetrieveByID(id uint, preloads ...string) (*model.ClassMaterial, error)
	DeleteByID(id uint) error
	UpdateByID(data model.ClassMaterial) error
	List(condition map[string]interface{}, preloads ...string) ([]model.ClassMaterial, error)
}

type ClassMaterialController struct {
	*core.BaseController[*model.ClassMaterial]
	connector *gorm.DB
}

func NewClassMaterialController(db *gorm.DB) *ClassMaterialController {
	return &ClassMaterialController{
		BaseController: core.NewBaseController[*model.ClassMaterial](db),
		connector:      db,
	}
}
