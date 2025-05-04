package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)
type LeaveInstructorHRController struct {
	db *gorm.DB
}

func CreateLeaveInstructorHRController(db *gorm.DB) *LeaveInstructorHRController {
	db.AutoMigrate(&model.RequestLeaveInstructor{})
	return &LeaveInstructorHRController{db: db}
}

func (c *LeaveInstructorHRController) insert(request *model.RequestLeaveInstructor) error {
	return c.db.Create(request).Error
}
func (c *LeaveInstructorHRController) update(request *model.RequestLeaveInstructor) error {
	return c.db.Save(request).Error
}
func (c *LeaveInstructorHRController) delete(request *model.RequestLeaveInstructor) error {
	return c.db.Delete(request).Error
}
func (c *LeaveInstructorHRController) getAll() ([]model.RequestLeaveInstructor, error) {
	var requests []model.RequestLeaveInstructor
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveInstructorHRController) getByID(id uint) (*model.RequestLeaveInstructor, error) {
	var request model.RequestLeaveInstructor
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveInstructorHRController) getByInstructorID(instructorID string) ([]model.RequestLeaveInstructor, error) {
	var requests []model.RequestLeaveInstructor
	err := c.db.Where("instructor_id = ?", instructorID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func SubmitInstructorLeaveRequest(db *gorm.DB,instructorID, leaveType, reason, leaveDateStr string) error {

	tm := &util.TransactionManager{DB:db}

	return tm.Execute(func(tx *gorm.DB) error {
		instructorController := CreateLeaveInstructorHRController(tx)
		factory := &model.RequestLeaveFactory{}

		req, err := factory.Create("instructor", instructorID, leaveType, reason, leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to build leave request: %v", err)
		}

		if err := instructorController.insert(req.(*model.RequestLeaveInstructor)); err != nil {
			return fmt.Errorf("failed to submit leave request: %v", err)
		}
		return nil
	})
}

	