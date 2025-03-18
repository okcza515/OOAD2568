package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type StudentController struct {
	Connector *gorm.DB
}

func CreateStudentController(connector *gorm.DB) *StudentController {
	student := StudentController{Connector: connector}
	connector.AutoMigrate(&model.Student{})
	return &student
}

func (student StudentController) GetAll() ([]*model.Student, error) {
	students := []*model.Student{}
	result := student.Connector.Find(&students)
	return students, result.Error
}

func (student StudentController) GetByStudentId(sid string) (*model.Student, error) {
	s := &model.Student{}
	result := student.Connector.Where("s_id = ?", sid).First(student)
	return s, result.Error
}

func (student StudentController) Create(s *model.Student) error {
	return student.Connector.Create(s).Error
}

func (student StudentController) Update(sid string, updatedData map[string]any) error {
	return student.Connector.Model(&model.Student{}).
		Where("s_id = ?", sid).
		Updates(updatedData).Error
}

func (student StudentController) DeleteByStudentId(sid string) error {
	return student.Connector.Where("s_id = ?", sid).Delete(&model.Student{}).Error
}
