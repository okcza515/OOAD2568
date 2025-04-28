package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type FacultyController struct {
	DB *gorm.DB
}

func CreateFacultyController(db *gorm.DB) *FacultyController {
	db.AutoMigrate(&model.Faculty{})
	return &FacultyController{DB: db}
}

func (c *FacultyController) GetAll() ([]*model.Faculty, error) {
	return model.GetAllCommonModels[model.Faculty](c.DB)
}

func (c *FacultyController) GetBy(field string, value interface{}) ([]*model.Faculty, error) {
	return model.GetRecordByField[model.Faculty](c.DB, field, value)
}

func (c *FacultyController) SetBudget(name string, budget int) error {
	return model.SetFacultyBudget(c.DB, name, budget)
}

func (c *FacultyController) UpdateBudget(name string, delta int) error {
	return model.UpdateFacultyBudget(c.DB, name, delta)
}

func (c *FacultyController) Register(faculties []*model.Faculty) error {
	return model.CommonRegister(c.DB, faculties)
}

func (c *FacultyController) Truncate() error {
	return model.TruncateModel(c.DB, "faculties")
}