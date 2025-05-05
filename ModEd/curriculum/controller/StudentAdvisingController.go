// MEP-1008
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"
	"errors"

	"gorm.io/gorm"
)

type StudentWorkloadService interface {
	Insert(data model.StudentAdvisor) error
	UpdateByID(data model.StudentAdvisor) error
	DeleteByID(id uint) error
	RetrieveByID(id uint, preloads ...string) (*model.StudentAdvisor, error)
	DeleteByStudentId(studentId uint) error
	GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error)
	CreateStudentRequest(data model.StudentRequest) error
	GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error)
	ReviewStudentRequest(id uint, review string, comment string) error
}

type StudentRequestController struct {
	*core.BaseController[*model.StudentAdvisor]
	db *gorm.DB
}

func CreateStudentWorkloadController(db *gorm.DB) StudentWorkloadService {
	return &StudentRequestController{
		BaseController: core.NewBaseController[*model.StudentAdvisor](db),
		db:             db,
	}
}

func (c *StudentRequestController) Insert(data model.StudentAdvisor) error {
	return c.BaseController.Insert(&data)
}

func (c *StudentRequestController) UpdateByID(data model.StudentAdvisor) error {
	return c.BaseController.UpdateByID(&data)
}

func (c *StudentRequestController) DeleteByID(id uint) error {
	return c.BaseController.DeleteByID(id)
}

func (c *StudentRequestController) RetrieveByID(id uint, preloads ...string) (*model.StudentAdvisor, error) {
	return c.BaseController.RetrieveByID(id, preloads...)
}

func (c *StudentRequestController) DeleteByStudentId(studentId uint) error {
	return c.db.Where("student_id = ?", studentId).Delete(&model.StudentAdvisor{}).Error
}

func (c *StudentRequestController) GetStudentUnderSupervisionByInstructorId(instructorId uint) ([]model.StudentAdvisor, error) {
	var advisors []model.StudentAdvisor
	err := c.db.Where("instructor_id = ?", instructorId).Preload("Students").Find(&advisors).Error
	if err != nil {
		return nil, err
	}
	return advisors, nil
}

func (c *StudentRequestController) CreateStudentRequest(data model.StudentRequest) error {
	return c.db.Create(&data).Error
}

func (c *StudentRequestController) GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error) {
	var requests []model.StudentRequest
	err := c.db.Where("instructor_id = ?", instructorId).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *StudentRequestController) ReviewStudentRequest(id uint, review string, comment string) error {
	var req model.StudentRequest
	if err := c.db.First(&req, id).Error; err != nil {
		return err
	}
	if review != "accept" && review != "reject" {
		return errors.New("invalid review value, must be 'accept' or 'reject'")
	}
	req.Review = review
	req.Comment = comment
	return c.db.Save(&req).Error
}
