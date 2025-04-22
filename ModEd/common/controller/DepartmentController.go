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
	return model.GetAllDepartments(c.DB)
}

func (c *DepartmentController) GetByName(name string) (*model.Department, error) {
	return model.GetDepartmentByName(c.DB, name)
}
func (c *DepartmentController) SetBudget(name string, budget int) error {
	return model.SetDepartmentBudget(c.DB, name, budget)
}

func (c *DepartmentController) UpdateBudget(name string, delta int) error {
	return model.UpdateDepartmentBudget(c.DB, name, delta)
}

func (c *DepartmentController) Register(departments []*model.Department) error {
	return model.RegisterDepartments(c.DB, departments)
}

func (c *DepartmentController) Truncate() error {
	return model.TruncateDepartments(c.DB)
}
