package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type LeaveStudentHRController struct {
	db *gorm.DB
}

func CreateLeaveStudentHRController(db *gorm.DB) *LeaveStudentHRController {
	db.AutoMigrate(&model.RequestLeaveStudent{})
	return &LeaveStudentHRController{db: db}
}

func (c *LeaveStudentHRController) insert(request *model.RequestLeaveStudent) error {
	return c.db.Create(request).Error
}
func (c *LeaveStudentHRController) update(request *model.RequestLeaveStudent) error {
	return c.db.Save(request).Error
}
func (c *LeaveStudentHRController) delete(request *model.RequestLeaveStudent) error {
	return c.db.Delete(request).Error
}
func (c *LeaveStudentHRController) getAll() ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (c *LeaveStudentHRController) getByID(id uint) (*model.RequestLeaveStudent, error) {
	var request model.RequestLeaveStudent
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
func (c *LeaveStudentHRController) getByStudentID(studentID string) ([]model.RequestLeaveStudent, error) {
	var requests []model.RequestLeaveStudent
	err := c.db.Where("student_id = ?", studentID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *LeaveStudentHRController) SubmitStudentLeaveRequest(studentID, leaveType, reason, leaveDateStr string) error {

	tm := &util.TransactionManager{DB: c.db}

	return tm.Execute(func(tx *gorm.DB) error {
		leaveController := CreateLeaveStudentHRController(tx)

		factory, err := model.GetFactory("student")
		if err != nil {
			return fmt.Errorf("failed to get student factory: %v", err)
		}

		reqInterface, err := factory.CreateLeave(studentID, leaveType, reason, leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to create leave request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestLeaveStudent)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for student leave request")
		}

		if err := leaveController.insert(req); err != nil {
			return fmt.Errorf("failed to submit leave request within transaction: %v", err)
		}
		return nil
	})
}

func (c *LeaveStudentHRController) ReviewStudentLeaveRequest(tx *gorm.DB, requestID string, action string, reason string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		id, err := strconv.ParseUint(requestID, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid request ID: %v", err)
		}

		request, err := c.getByID(uint(id))
		if err != nil {
			return fmt.Errorf("failed to get leave request: %v", err)
		}

		if action == "approve" {
			request.Status = action
		} else if action == "reject" {
			request.Status = action
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
