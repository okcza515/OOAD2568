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
	return model.CommonModelGetAll[model.Student](c.DB)
}

func (c *StudentController) GetBy(field string, value interface{}) ([]*model.Student, error) {
	return model.GetRecordByField[model.Student](c.DB, field, value)
}

func (c *StudentController) Update(code string, updatedData map[string]any) error {
	return model.UpdateStudentByCode(c.DB, code, updatedData)
}

func (c *StudentController) DeleteByCode(code string) error {
	return model.DeleteStudentByCode(c.DB, code)
}

func (c *StudentController) Register(students []*model.Student) error {
	return model.CommonRegister(c.DB, students)
}

func (c *StudentController) Delete(field string, value interface{}) error {
	return model.DeleteRecordByField[model.Student](c.DB, field, value)
}

func (c *StudentController) Truncate() error {
	return model.TruncateModel(c.DB, "students")
}