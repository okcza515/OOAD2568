package controller

import (
	"ModEd/common/model"
	"ModEd/core"
	"errors"

	"gorm.io/gorm"
)

type DepartmentController struct {
	*core.BaseController[model.Department]
}

func newDepartmentController(db *gorm.DB) *DepartmentController {
	db.AutoMigrate(&model.Department{})
	return &DepartmentController{
		BaseController: core.NewBaseController[model.Department](db),
	}
}

func (c *DepartmentController) GetAll() ([]model.Department, error) {
	return c.List(nil)
}

func (c *DepartmentController) GetBy(field string, value interface{}) ([]model.Department, error) {
	return c.List(map[string]interface{}{field: value})
}

func (c *DepartmentController) SetBudget(name string, budget int) error {
	if budget < 0 {
		return errors.New("budget cannot be negative")
	}
	return c.UpdateByCondition(map[string]interface{}{
		"name": name,
	}, model.Department{
		Budget: budget,
	})
}

func (c *DepartmentController) UpdateBudget(name string, delta int) error {
	department, err := c.RetrieveByCondition(map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}

	newBudget := department.Budget + delta
	if newBudget < 0 {
		return errors.New("budget cannot be negative")
	}

	return c.SetBudget(name, newBudget)
}

func (c *DepartmentController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{field: value}, model.Department{})
}

func (c *DepartmentController) Register(departments []model.Department) error {
	return c.InsertMany(departments)
}

func (c *DepartmentController) Delete(field string, value interface{}) error {
	return c.DeleteByCondition(map[string]interface{}{field: value})
}

func (c *DepartmentController) Truncate(db *gorm.DB) error {
	return model.TruncateModel(db, "departments")
}

func (c *DepartmentController) GetByFaculty(faculty string) ([]model.Department, error) {
	return c.List(map[string]interface{}{
		"faculty": faculty,
	})
}

func (c *DepartmentController) GetByName(name string) (model.Department, error) {
	return c.RetrieveByCondition(map[string]interface{}{
		"name": name,
	})
}

func (c *DepartmentController) UpdateByName(name string, department model.Department) error {
	return c.UpdateByCondition(map[string]interface{}{
		"name": name,
	}, department)
}

func (c *DepartmentController) DeleteByName(name string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"name": name,
	})
}
