package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type ResignationStudentHRController struct {
	db *gorm.DB
}

func CreateResignationStudentHRController(db *gorm.DB) *ResignationStudentHRController {
	db.AutoMigrate(&model.RequestResignation{})
	return &ResignationStudentHRController{db: db}
}

func (c *ResignationStudentHRController) Insert(request *model.RequestResignationStudent) error {
	return c.db.Create(request).Error
}

func (c *ResignationStudentHRController) GetByStudentID(id string) (*model.RequestResignationStudent, error) {
	var req model.RequestResignationStudent
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationStudentHRController) Update(req *model.RequestResignationStudent) error {
	return c.db.Save(req).Error
}
