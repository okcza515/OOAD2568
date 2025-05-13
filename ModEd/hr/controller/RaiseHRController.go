package controller

import (
	"ModEd/core"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

type RaiseHRController struct {
	db *gorm.DB
}

func NewRaiseHRController(db *gorm.DB) *RaiseHRController {
	db.AutoMigrate(&model.RequestRaiseInstructor{})
	return &RaiseHRController{db: db}
}

func CreateRaiseInstructorHRController(db *gorm.DB) *RaiseHRController {
	db.AutoMigrate(&model.RequestRaiseInstructor{})
	return &RaiseHRController{db: db}
}

func (c *RaiseHRController) insert(req *model.RequestRaiseInstructor) error {
	return c.db.Create(req).Error
}

func (c *RaiseHRController) getByID(id uint) (*model.RequestRaiseInstructor, error) {
	var raise model.RequestRaiseInstructor
	err := c.db.First(&raise, id).Error
	return &raise, err
}

func (c *RaiseHRController) getAll() ([]*model.RequestRaiseInstructor, error) {
	var request []*model.RequestRaiseInstructor
	err := c.db.Find(&request).Error
	return request, err
}

func (c *RaiseHRController) updateStatus(id uint, status string) error {
	return c.db.Model(&model.RequestRaiseInstructor{}).Where("id = ?", id).Update("status", status).Error
}

func (c *RaiseHRController) getByInstructorID(instructorID string) ([]model.RequestRaiseInstructor, error) {
	var requests []model.RequestRaiseInstructor
	err := c.db.Where("instructor_code = ?", instructorID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *RaiseHRController) SubmitRaiseRequest(instructorID string, amount float64, reason string) error {
	tm := &util.TransactionManager{DB: c.db}

	return tm.Execute(func(tx *gorm.DB) error {
		raiseController := NewRaiseHRController(tx)

		requestFactory := model.RequestFactory{}

		params := model.CreateRequestParams{
			ID:           instructorID,
			Reason:       reason,
			TargetSalary: amount,
		}

		reqInterface, err := requestFactory.CreateRequest(model.RoleInstructor, model.RequestTypeLeave, params)
		if err != nil {
			return fmt.Errorf("failed to create raise request: %v", err)
		}

		req, ok := reqInterface.(*model.RequestRaiseInstructor)
		if !ok {
			return fmt.Errorf("failed to cast request to RequestRaiseInstructor")
		}

		if err := raiseController.insert(req); err != nil {
			return fmt.Errorf("failed to submit raise request: %v", err)
		}
		return nil
	})
}

func (c *RaiseHRController) ReviewInstructorRaiseRequest(requestID, action, reason string,
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

func (c *RaiseHRController) ExportInstructorRaiseRequests(filePath string) error {
	request, err := c.getAll()
	if err != nil {
		return fmt.Errorf("failed to retrieve instructor raise requests: %w", err)
	}

	mapper, err := core.CreateMapper[model.RequestRaiseInstructor](filePath)
	if err != nil {
		return fmt.Errorf("failed to create instructor raise request mapper: %w", err)
	}

	err = mapper.Serialize(request)
	if err != nil {
		return fmt.Errorf("failed to serialize instructor raise requests: %w", err)
	}

	fmt.Printf("Exported %d instructor raise requests to %s\n", len(request), filePath)

	return nil
}
