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

func NewLeaveInstructorHRController(db *gorm.DB) *LeaveInstructorHRController {
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

func (c *LeaveInstructorHRController) SubmitInstructorLeaveRequest(instructorID, leaveType, reason, leaveDateStr string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {

		leaveController := NewLeaveInstructorHRController(tx)

		factory, err := model.GetFactory(1)
		if err != nil {
			return fmt.Errorf("failed to get instructor factory: %v", err)
		}

		reqInterface, err := factory.CreateLeave(instructorID, leaveType, reason, leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to create leave request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestLeaveInstructor)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor leave request")
		}

		if err := leaveController.insert(req); err != nil {
			return fmt.Errorf("failed to submit leave request within transaction: %v", err)
		}

		return nil
	})
}

func (c *LeaveInstructorHRController) ReviewInstructorLeaveRequest(
	tx *gorm.DB,
	requestID, action, reason string,
) error {
	return ReviewRequest(
		requestID,
		action,
		reason,
		func(id uint) (Reviewable, error) {
			return c.getByID(id)
		},
		func(r Reviewable) error {
			return tx.Save(r).Error
		},
	)
}
