package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type DepartmentController struct {
	DB *gorm.DB
}

func NewDepartmentController(db *gorm.DB) *DepartmentController {
	db.AutoMigrate(&model.Department{})
	return &DepartmentController{DB: db}
}

func (c *DepartmentController) GetAll() ([]*model.Department, error) {
	return model.CommonModelGetAll[model.Department](c.DB)
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

func (c *DepartmentController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return model.UpdateRecordByField[model.Department](c.DB, field, value, updatedData, model.Department{})
}

func (c *DepartmentController) Register(department []*model.Department) error {
	return model.CommonRegister(c.DB, department)
}

func (c *DepartmentController) Delete(field string, value interface{}) error {
	return model.DeleteRecordByField[model.Department](c.DB, field, value, model.Department{})
}

func (c *DepartmentController) Truncate() error {
	return model.TruncateModel(c.DB, "departments")
}

func (c *DepartmentController) GetByFaculty(faculty string) ([]*model.Department, error) {
	return model.GetDepartmentsByFaculty(c.DB, faculty)
}
