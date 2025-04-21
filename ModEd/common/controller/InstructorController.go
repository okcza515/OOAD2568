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
	return model.GetAllInstructors(c.DB)
}

func (c *InstructorController) GetByCode(code string) (*model.Instructor, error) {
	return model.GetInstructorByCode(c.DB, code)
}

func (c *InstructorController) Create(instructor *model.Instructor) error {
	return model.CreateInstructor(c.DB, instructor)
}

func (c *InstructorController) Update(code string, updatedData map[string]any) error {
	return model.UpdateInstructorByCode(c.DB, code, updatedData)
}

func (c *InstructorController) DeleteByCode(code string) error {
	return model.DeleteInstructorByCode(c.DB, code)
}

func (c *InstructorController) Register(instructors *[]model.Instructor) error {
	return model.RegisterInstructors(c.DB, instructors)
}
