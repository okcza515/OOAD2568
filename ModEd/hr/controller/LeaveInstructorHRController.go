package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
	"strconv"

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

func (c *LeaveInstructorHRController) SubmitInstructorLeaveRequest(instructorID, leaveType, reason, leaveDateStr string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {

		leaveController := CreateLeaveInstructorHRController(tx)

		factory, err := model.GetFactory("instructor")
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

func (c *LeaveInstructorHRController) ReviewInstructorLeaveRequest(tx *gorm.DB, requestID string, action string, reason string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		id, err := strconv.ParseUint(requestID, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid request ID: %v", err)
		}

		request, err := c.getByID(uint(id))
		if err != nil {
			return fmt.Errorf("failed to find leave request: %v", err)
		}

		if action == "approve" {
			request.Status = "approved"
		} else if action == "reject" {
			request.Status = "rejected"
			request.Reason = reason
		} else {
			return fmt.Errorf("invalid action: %s", action)
		}

		if err := c.update(request); err != nil {
			return fmt.Errorf("failed to update leave request: %v", err)
		}
		return nil
	})
}
