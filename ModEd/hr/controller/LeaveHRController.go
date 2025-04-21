package controller

import (
	"gorm.io/gorm"
	"ModEd/hr/model"
)
type LeaveHRController struct {
	db *gorm.DB
}

func CreateLeaveHRController(db *gorm.DB) *LeaveHRController {
	db.AutoMigrate(&model.RequestLeave{})
	return &LeaveHRController{db: db}
}

func (c *LeaveHRController) Insert(request *model.RequestLeave) error {
	return c.db.Create(request).Error
}
func (c *LeaveHRController) Update(request *model.RequestLeave) error {
	return c.db.Save(request).Error
}
func (c *LeaveHRController) Delete(request *model.RequestLeave) error {
	return c.db.Delete(request).Error
}
func (c *LeaveHRController) GetAll() ([]model.RequestLeave, error) {
	var requests []model.RequestLeave
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveHRController) GetByID(id uint) (*model.RequestLeave, error) {
	var request model.RequestLeave
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveHRController) GetByStudentID(studentID string) ([]model.RequestLeave, error) {
	var requests []model.RequestLeave
	err := c.db.Where("student_id = ?", studentID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}