// MEP-1008

package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassWorkloadService interface {
	AddClassMaterial(material *model.ClassMaterial) error
	GetClassMaterialsByClassId(classId uint) ([]model.ClassMaterial, error)
	AddClassLecture(lecture *model.ClassLecture) error
	GetClassLecturesByClassId(classId uint) ([]model.ClassLecture, error)
	DeleteClassMaterial(materialId uint) error
	DeleteClassLecture(lectureId uint) error
	UpdateClassMaterial(material *model.ClassMaterial) error
	UpdateClassLecture(lecture *model.ClassLecture) error
}

type ClassWorkloadController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateClassWorkloadController(db *gorm.DB) ClassWorkloadService {
	return &ClassWorkloadController{
		BaseController: core.NewBaseController("ClassWorkload", db),
		Connector:      db,
	}
}

func (c *ClassWorkloadController) AddClassMaterial(material *model.ClassMaterial) error {
	if err := c.Connector.Create(material).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClassWorkloadController) GetClassMaterialsByClassId(classId uint) ([]model.ClassMaterial, error) {
	var materials []model.ClassMaterial
	if err := c.Connector.Where("class_id = ?", classId).Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}

func (c *ClassWorkloadController) AddClassLecture(lecture *model.ClassLecture) error {
	if err := c.Connector.Create(lecture).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClassWorkloadController) GetClassLecturesByClassId(classId uint) ([]model.ClassLecture, error) {
	var lectures []model.ClassLecture
	if err := c.Connector.Where("class_id = ?", classId).Find(&lectures).Error; err != nil {
		return nil, err
	}
	return lectures, nil
}

func (c *ClassWorkloadController) DeleteClassMaterial(materialId uint) error {
	if err := c.Connector.Delete(&model.ClassMaterial{}, materialId).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClassWorkloadController) DeleteClassLecture(lectureId uint) error {
	if err := c.Connector.Delete(&model.ClassLecture{}, lectureId).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClassWorkloadController) UpdateClassMaterial(material *model.ClassMaterial) error {
	if err := c.Connector.Save(material).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClassWorkloadController) UpdateClassLecture(lecture *model.ClassLecture) error {
	if err := c.Connector.Save(lecture).Error; err != nil {
		return err
	}
	return nil
}
