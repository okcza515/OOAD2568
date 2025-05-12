package controller

import (
	"ModEd/core"
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
func (c *LeaveInstructorHRController) getAll() ([]*model.RequestLeaveInstructor, error) {
	var requests []*model.RequestLeaveInstructor
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
	err := c.db.Where("instructor_code = ?", instructorID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *LeaveInstructorHRController) SubmitInstructorLeaveRequest(instructorID, leaveType, reason, leaveDateStr string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {

		leaveController := NewLeaveInstructorHRController(tx)

		requestFactory := model.RequestFactory{}

		params := model.CreateRequestParams{
			ID:        instructorID,
			LeaveType: leaveType,
			Reason:    reason,
			DateStr:   leaveDateStr,
		}

		reqInterface, err := requestFactory.CreateRequest(model.RoleInstructor, model.RequestTypeLeave, params)
		if err != nil {
			return fmt.Errorf("failed to create leave request using factory: %w", err)
		}

		req, ok := reqInterface.(*model.RequestLeaveInstructor)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor leave request, got %T", reqInterface)
		}

		if err := leaveController.insert(req); err != nil {
			return fmt.Errorf("failed to submit leave request within transaction: %w", err)
		}

		return nil
	})
}

func (c *LeaveInstructorHRController) ReviewInstructorLeaveRequest(requestID, action, reason string,
) error {
	return ReviewRequest(
		requestID,
		action,
		reason,
		func(id uint) (Reviewable, error) {
			return c.getByID(id)
		},
		func(r Reviewable) error {
			return c.db.Save(r).Error
		},
	)
}

func (c *LeaveInstructorHRController) ExportInstructorLeaveRequests(filePath string) error {
	requests, err := c.getAll()
	if err != nil {
		return fmt.Errorf("failed to retrieve instructor leave requests: %w", err)
	}

	mapper, err := core.CreateMapper[model.RequestLeaveInstructor](filePath)
	if err != nil {
		return fmt.Errorf("failed to create instructor leave request mapper: %w", err)
	}

	err = mapper.Serialize(requests)
	if err != nil {
		return fmt.Errorf("failed to serialize instructor leave requests: %w", err)
	}

	fmt.Printf("Exported %d instructor leave requests to %s\n", len(requests), filePath)
	return nil
}
