package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type DepartmentController struct {
	DB *gorm.DB
}

func CreateDepartmentController(db *gorm.DB) *DepartmentController {
	db.AutoMigrate(&model.Department{})
	return &DepartmentController{DB: db}
}

func (c *DepartmentController) GetAll() ([]*model.Department, error) {
	return model.GetAllCommonModels[model.Department](c.DB)
}

func (c *DepartmentController) GetBy(field string, value interface{}) ([]*model.Department, error) {
	return model.GetRecordByField[model.Department](c.DB, field, value)
}

func (c *DepartmentController) SetBudget(name string, budget int) error {
	return model.SetDepartmentBudget(c.DB, name, budget)
}

func (c *DepartmentController) UpdateBudget(name string, delta int) error {
	return model.UpdateDepartmentBudget(c.DB, name, delta)
}

func (c *DepartmentController) GetByFaculty(faculty string) ([]*model.Department, error) {
	return model.GetDepartmentsByFaculty(c.DB, faculty)
}
