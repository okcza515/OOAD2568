package controller

import (
	"ModEd/common/model"
	"ModEd/core"
	"errors"

	"gorm.io/gorm"
)

type FacultyController struct {
	*core.BaseController[model.Faculty]
}

func NewFacultyController(db *gorm.DB) *FacultyController {
	db.AutoMigrate(&model.Faculty{})
	return &FacultyController{
		BaseController: core.NewBaseController[model.Faculty](db),
	}
}

func (c *FacultyController) GetAll() ([]model.Faculty, error) {
	return c.List(nil)
}

func (c *FacultyController) GetBy(field string, value interface{}) ([]model.Faculty, error) {
	return c.List(map[string]interface{}{field: value})
}

func (c *FacultyController) SetBudget(name string, budget int) error {
	return c.UpdateByCondition(map[string]interface{}{
		"name": name,
	}, model.Faculty{
		Budget: budget,
	})
}

func (c *FacultyController) UpdateBudget(name string, delta int) error {
	faculty, err := c.RetrieveByCondition(map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}

	newBudget := faculty.Budget + delta
	if newBudget < 0 {
		return errors.New("budget cannot be negative")
	}

	return c.SetBudget(name, newBudget)
}

func (c *FacultyController) UpdateByField(field string, value interface{}, updatedData map[string]any) error {
	return c.UpdateByCondition(map[string]interface{}{field: value}, model.Faculty{})
}

func (c *FacultyController) Register(faculties []model.Faculty) error {
	return c.InsertMany(faculties)
}

func (c *FacultyController) Truncate() error {
	return c.DeleteByCondition(map[string]interface{}{})
}

func (c *FacultyController) GetByName(name string) (model.Faculty, error) {
	return c.RetrieveByCondition(map[string]interface{}{
		"name": name,
	})
}

func (c *FacultyController) UpdateByName(name string, faculty model.Faculty) error {
	return c.UpdateByCondition(map[string]interface{}{
		"name": name,
	}, faculty)
}

func (c *FacultyController) DeleteByName(name string) error {
	return c.DeleteByCondition(map[string]interface{}{
		"name": name,
	})
}
