package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type InstructorController struct {
	DB *gorm.DB
}

func CreateInstructorController(db *gorm.DB) *InstructorController {
	db.AutoMigrate(&model.Instructor{})
	return &InstructorController{DB: db}
}

func (c *InstructorController) GetAll() ([]*model.Instructor, error) {
	return model.GetAllCommonModels[model.Instructor](c.DB)
}

func (c *InstructorController) GetBy(field string, value interface{}) ([]*model.Instructor, error) {
	return model.GetRecordByField[model.Instructor](c.DB, field, value)
}

func (c *InstructorController) Update(code string, updatedData map[string]any) error {
	return model.UpdateInstructorByCode(c.DB, code, updatedData)
}

func (c *InstructorController) DeleteByCode(code string) error {
	return model.DeleteInstructorByCode(c.DB, code)
}
