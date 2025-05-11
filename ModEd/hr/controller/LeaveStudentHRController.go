package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

type LeaveStudentHRController struct {
	db *gorm.DB
}

func NewLeaveStudentHRController(db *gorm.DB) *LeaveStudentHRController {
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
		leaveController := NewLeaveStudentHRController(tx)

		requestFactory := model.RequestFactory{}

		params := model.CreateRequestParams{
			ID:        studentID,
			LeaveType: leaveType,
			Reason:    reason,
			DateStr:   leaveDateStr,
		}

		reqInterface, err := requestFactory.CreateRequest(model.RoleInstructor, model.RequestTypeLeave, params)
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

func (c *LeaveStudentHRController) ReviewStudentLeaveRequest(requestID, action, reason string,
) error {
	return ReviewRequest(
		requestID,
		action,
		reason,
		// fetch
		func(id uint) (Reviewable, error) {
			return c.getByID(id)
		},
		// save
		func(r Reviewable) error {
			return c.db.Save(r).Error
		},
	)
}
