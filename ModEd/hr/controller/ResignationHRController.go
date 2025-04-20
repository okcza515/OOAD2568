package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type ResignationHRController struct {
	db *gorm.DB
}

func CreateResignationHRController(db *gorm.DB) *ResignationHRController {
	db.AutoMigrate(&model.RequestResignation{})
	return &ResignationHRController{db: db}
}

func (c *ResignationHRController) Insert(request *model.RequestResignation) error {
	return c.db.Create(request).Error
}

func (c *ResignationHRController) GetByStudentID(id string) (*model.RequestResignation, error) {
	var req model.RequestResignation
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationHRController) Update(req *model.RequestResignation) error {
	return c.db.Save(req).Error
}
