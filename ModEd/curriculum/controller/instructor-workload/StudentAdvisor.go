package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"
)

type StudentAdvisorController struct {
	DB *gorm.DB
}

func (sac *StudentAdvisorController) CreateStudentAdvisor(studentAdvisor model.StudentAdvisor) error {
	return sac.DB.Create(&studentAdvisor).Error
}

func (sac *StudentAdvisorController) GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error) {
	var studentAdvisors []model.StudentAdvisor
	err := sac.DB.Where("instructor_id = ?", instructorId).Preload("Students").Find(&studentAdvisors).Error
	if err != nil {
		return nil, err
	}
	return studentAdvisors, nil
}
