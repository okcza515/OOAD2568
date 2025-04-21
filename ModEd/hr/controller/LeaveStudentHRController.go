package controller

import (
	"gorm.io/gorm"
	"ModEd/hr/model"
)
type LeaveStudentHRController struct {
	db *gorm.DB
}

func CreateLeaveStudentHRController(db *gorm.DB) *LeaveStudentHRController {
	db.AutoMigrate(&model.RequestLeaveStudent{})
	return &LeaveStudentHRController{db: db}
}

func (c *LeaveStudentHRController) Insert(request *model.RequestLeaveStudent) error {
	return c.db.Create(request).Error
}
func (c *LeaveStudentHRController) Update(request *model.RequestLeaveStudent) error {
	return c.db.Save(request).Error
}
func (c *LeaveStudentHRController) Delete(request *model.RequestLeaveStudent) error {
	return c.db.Delete(request).Error
}
func (c *LeaveStudentHRController) GetAll() ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveStudentHRController) GetByID(id uint) (*model.RequestLeaveStudent, error) {
	var request model.RequestLeaveStudent
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveStudentHRController) GetByStudentID(studentID string) ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Where("student_id = ?", studentID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}