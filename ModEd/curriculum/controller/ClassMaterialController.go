// MEP-1008

package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassMaterialService interface {
	//AddClassMaterial(material *model.ClassMaterial) error
	Insert(data model.ClassMaterial) error
	//GetClassMaterialsByClassId(classId uint) ([]model.ClassMaterial, error)
	RetrieveByID(id uint, preloads ...string) (*model.ClassMaterial, error)
	//DeleteClassMaterial(materialId uint) error
	DeleteByID(id uint) error
	//UpdateClassMaterial(material *model.ClassMaterial) error
	UpdateByID(data model.ClassMaterial) error
}

type ClassMaterialController struct {
	*core.BaseController[*model.ClassMaterial]
	Connector *gorm.DB
}

func CreateClassMaterialController(db *gorm.DB) *ClassMaterialController {
	return &ClassMaterialController{
		BaseController: core.NewBaseController[*model.ClassMaterial](db),
		Connector:      db,
	}
}

// func (c *ClassMaterialController) AddClassMaterial(material *model.ClassMaterial) error {
// 	if err := c.Connector.Create(material).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *ClassMaterialController) GetClassMaterialsByClassId(classId uint) ([]model.ClassMaterial, error) {
// 	var materials []model.ClassMaterial
// 	if err := c.Connector.Where("class_id = ?", classId).Find(&materials).Error; err != nil {
// 		return nil, err
// 	}
// 	return materials, nil
// }

// func (c *ClassMaterialController) DeleteClassMaterial(materialId uint) error {
// 	if err := c.Connector.Delete(&model.ClassMaterial{}, materialId).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *ClassMaterialController) UpdateClassMaterial(material *model.ClassMaterial) error {
// 	if err := c.Connector.Save(material).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
