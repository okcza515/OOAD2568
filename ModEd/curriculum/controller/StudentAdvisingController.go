// MEP-1008
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type StudentWorkloadService interface {
	ListStudentRequest(instructorId uint) ([]model.StudentRequest, error)
	Insert(data model.StudentAdvisor) error
	UpdateByID(data model.StudentAdvisor) error
	DeleteByID(id uint) error
	DeleteByStudentId(studentId uint) error
	RetrieveByID(id uint, preloads ...string) (*model.StudentAdvisor, error)
	GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error)
	CreateStudentRequest(data model.StudentRequest) error
	GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error)
	ReviewStudentRequest(id uint, review string, comment string) error
}

type StudentWorkloadController struct {
	*core.BaseController[*model.StudentAdvisor]
	Connector *gorm.DB
}

func CreateStudentWorkloadController(db *gorm.DB) *StudentWorkloadController {
	return &StudentWorkloadController{
		BaseController: core.NewBaseController[*model.StudentAdvisor](db),
		Connector:      db,
	}
}

func (swc *StudentWorkloadController) ListStudentRequest(instructorId uint) ([]model.StudentRequest, error) {
	var studentRequests []model.StudentRequest
	err := swc.Connector.Where("instructor_id = ?", instructorId).Find(&studentRequests).Error
	if err != nil {
		return nil, err
	}
	return studentRequests, nil
}

func (swc *StudentWorkloadController) CreateStudentAdvisor(studentAdvisor model.StudentAdvisor) error {
	return swc.Connector.Create(&studentAdvisor).Error
}

func (swc *StudentWorkloadController) UpdateStudentAdvisor(studentAdvisor model.StudentAdvisor) error {
	return swc.Connector.Save(&studentAdvisor).Error
}

func (swc *StudentWorkloadController) DeleteStudentAdvisor(id uint) error {
	var studentAdvisor model.StudentAdvisor
	err := swc.Connector.First(&studentAdvisor, id).Error
	if err != nil {
		return err
	}
	return swc.Connector.Delete(&studentAdvisor).Error
}

func (swc *StudentWorkloadController) GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error) {
	var studentAdvisors []model.StudentAdvisor
	err := swc.Connector.Where("instructor_id = ?", instructorId).Preload("Students").Find(&studentAdvisors).Error
	if err != nil {
		return nil, err
	}
	return studentAdvisors, nil
}

func (swc *StudentWorkloadController) CreateStudentRequest(studentRequest model.StudentRequest) error {
	return swc.Connector.Create(&studentRequest).Error
}

func (swc *StudentWorkloadController) GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error) {
	var studentRequests []model.StudentRequest
	err := swc.Connector.Where("instructor_id = ?", instructorId).Find(&studentRequests).Error
	if err != nil {
		return nil, err
	}
	return studentRequests, nil
}

func (swc *StudentWorkloadController) ReviewStudentRequest(id uint, review string, comment string) error {
	var studentRequest model.StudentRequest
	err := swc.Connector.First(&studentRequest, id).Error
	if err != nil {
		return err
	}
	if review != "accept" && review != "reject" {
		return errors.New("invalid review value, must be 'accept' or 'reject'")
	}
	studentRequest.Review = review
	studentRequest.Comment = comment
	return swc.Connector.Save(&studentRequest).Error
}
