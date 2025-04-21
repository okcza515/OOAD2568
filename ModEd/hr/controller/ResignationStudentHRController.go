package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type ResignationStudentHRController struct {
	db *gorm.DB
}

func createResignationStudentHRController(db *gorm.DB) *ResignationStudentHRController {
	db.AutoMigrate(&model.RequestResignationStudent{})
	return &ResignationStudentHRController{db: db}
}

func (c *ResignationStudentHRController) insert(request *model.RequestResignationStudent) error {
	return c.db.Create(request).Error
}

func (c *ResignationStudentHRController) getByStudentID(id string) (*model.RequestResignationStudent, error) {
	var req model.RequestResignationStudent
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationStudentHRController) update(req *model.RequestResignationStudent) error {
	return c.db.Save(req).Error
}
