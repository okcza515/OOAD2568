//MEP-1008
package controller

import (
	model "ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type StudentWorkloadService interface {
	CreateStudentAdvisor(studentAdvisor model.StudentAdvisor) error
	GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error)

	CreateStudentRequest(studentRequest model.StudentRequest) error
	GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error)
	ReviewStudentRequest(id uint, review string, comment string) error
}

type StudentWorkloadController struct {
	DB *gorm.DB
}

func NewStudentWorkloadController(db *gorm.DB) StudentWorkloadService {
	return &StudentWorkloadController{DB: db}
}

func (swc *StudentWorkloadController) CreateStudentAdvisor(studentAdvisor model.StudentAdvisor) error {
	return swc.DB.Create(&studentAdvisor).Error
}

func (swc *StudentWorkloadController) GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error) {
	var studentAdvisors []model.StudentAdvisor
	err := swc.DB.Where("instructor_id = ?", instructorId).Preload("Students").Find(&studentAdvisors).Error
	if err != nil {
		return nil, err
	}
	return studentAdvisors, nil
}

func (swc *StudentWorkloadController) CreateStudentRequest(studentRequest model.StudentRequest) error {
	return swc.DB.Create(&studentRequest).Error
}

func (swc *StudentWorkloadController) GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error) {
	var studentRequests []model.StudentRequest
	err := swc.DB.Where("instructor_id = ?", instructorId).Find(&studentRequests).Error
	if err != nil {
		return nil, err
	}
	return studentRequests, nil
}

func (swc *StudentWorkloadController) ReviewStudentRequest(id uint, review string, comment string) error {
	var studentRequest model.StudentRequest
	err := swc.DB.First(&studentRequest, id).Error
	if err != nil {
		return err
	}
	if review != "accept" && review != "reject" {
		return errors.New("invalid review value, must be 'accept' or 'reject'")
	}
	studentRequest.Review = review
	studentRequest.Comment = comment
	return swc.DB.Save(&studentRequest).Error
}
