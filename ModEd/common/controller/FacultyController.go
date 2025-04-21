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
	return model.GetAllFaculties(c.DB)
}

func (c *FacultyController) GetByName(name string) (*model.Faculty, error) {
	return model.GetFacultyByName(c.DB, name)
}

func (c *FacultyController) Create(faculty *model.Faculty) error {
	return model.CreateFaculty(c.DB, faculty)
}

func (c *FacultyController) SetBudget(name string, budget int) error {
	return model.SetFacultyBudget(c.DB, name, budget)
}

func (c *FacultyController) UpdateBudget(name string, delta int) error {
	return model.UpdateFacultyBudget(c.DB, name, delta)
}

func (c *FacultyController) Register(faculties []*model.Faculty) error {
	return model.RegisterFaculties(c.DB, faculties)
}
