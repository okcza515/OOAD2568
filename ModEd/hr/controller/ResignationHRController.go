package controller

import (
	"gorm.io/gorm"
	"ModEd/hr/model"
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
