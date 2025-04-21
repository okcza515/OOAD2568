package controller

import (
	"gorm.io/gorm"
	"ModEd/hr/model"
)
type LeaveInstructorHRController struct {
	db *gorm.DB
}

func CreateLeaveInstructorHRController(db *gorm.DB) *LeaveInstructorHRController {
	db.AutoMigrate(&model.RequestLeaveInstructor{})
	return &LeaveInstructorHRController{db: db}
}

func (c *LeaveInstructorHRController) Insert(request *model.RequestLeaveInstructor) error {
	return c.db.Create(request).Error
}
func (c *LeaveInstructorHRController) Update(request *model.RequestLeaveInstructor) error {
	return c.db.Save(request).Error
}
func (c *LeaveInstructorHRController) Delete(request *model.RequestLeaveInstructor) error {
	return c.db.Delete(request).Error
}
func (c *LeaveInstructorHRController) GetAll() ([]model.RequestLeaveInstructor, error) {
	var requests []model.RequestLeaveInstructor
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveInstructorHRController) GetByID(id uint) (*model.RequestLeaveInstructor, error) {
	var request model.RequestLeaveInstructor
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveInstructorHRController) GetByStudentID(studentID string) ([]model.RequestLeaveInstructor, error) {
	var requests []model.RequestLeaveInstructor
	err := c.db.Where("student_id = ?", studentID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}