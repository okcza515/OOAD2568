package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func CreateStudentController(db *gorm.DB) *StudentController {
	db.AutoMigrate(&model.Student{})
	return &StudentController{DB: db}
}

func (c *StudentController) GetAll() ([]*model.Student, error) {
	return model.GetAllStudents(c.DB)
}

func (c *StudentController) GetByCode(code string) (*model.Student, error) {
	return model.GetStudentByCode(c.DB, code)
}

func (c *StudentController) Create(student *model.Student) error {
	return model.CreateStudent(c.DB, student)
}

func (c *StudentController) Update(code string, updatedData map[string]any) error {
	return model.UpdateStudentByCode(c.DB, code, updatedData)
}

func (c *StudentController) DeleteByCode(code string) error {
	return model.DeleteStudentByCode(c.DB, code)
}

func (c *StudentController) Register(students []*model.Student) error {
	return model.RegisterStudents(c.DB, students)
}

func (c *StudentController) Truncate() error {
	return model.TruncateStudents(c.DB)
}
