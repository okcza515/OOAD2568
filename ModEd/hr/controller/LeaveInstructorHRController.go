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

func (c *LeaveInstructorHRController) insert(db *gorm.DB, request *model.RequestLeaveInstructor) error {
	return db.Create(request).Error
}
func (c *LeaveInstructorHRController) update(db *gorm.DB, request *model.RequestLeaveInstructor) error {
	return db.Save(request).Error
}
func (c *LeaveInstructorHRController) delete(db *gorm.DB, request *model.RequestLeaveInstructor) error {
	return db.Delete(request).Error
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

func (c *LeaveInstructorHRController) SubmitInstructorLeaveRequest(instructorID, leaveType, reason, leaveDateStr string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		factory, err := model.GetFactory("instructor")
		if err != nil {
			return fmt.Errorf("failed to get student factory: %v", err)
		}

		reqInterface, err := factory.CreateLeave(instructorID, leaveType, reason, leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to create resignation request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestLeaveInstructor)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor resignation request")
		}

		if err := c.insert(tx, req); err != nil {
			return fmt.Errorf("failed to insert resignation request within transaction: %v", err)
		}

		return nil
	})
}
