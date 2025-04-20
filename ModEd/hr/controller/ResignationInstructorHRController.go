package controller

import (
	"ModEd/hr/model"

	"gorm.io/gorm"
)

type ResignationInstructorHRController struct {
	db *gorm.DB
}

func CreateResignationInstructorHRController(db *gorm.DB) *ResignationInstructorHRController {
    db.AutoMigrate(&model.RequestResignationInstructor{})
    return &ResignationInstructorHRController{db: db}
}

func (c *ResignationInstructorHRController) Insert(request *model.RequestResignationInstructor) error {
	return c.db.Create(request).Error
}

func (c *ResignationInstructorHRController) GetByInstructorID(id string) (*model.RequestResignationInstructor, error) {
	var req model.RequestResignationInstructor
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationInstructorHRController) Update(req *model.RequestResignationInstructor) error {
	return c.db.Save(req).Error
}