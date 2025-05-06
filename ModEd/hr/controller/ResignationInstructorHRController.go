package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

type ResignationInstructorHRController struct {
	db *gorm.DB
}

func NewResignationInstructorHRController(db *gorm.DB) *ResignationInstructorHRController {
	db.AutoMigrate(&model.RequestResignationInstructor{})
	return &ResignationInstructorHRController{db: db}
}

// Use the passed db object (which will be 'tx' in the transaction context)
func (c *ResignationInstructorHRController) insert(request *model.RequestResignationInstructor) error {
	return c.db.Create(request).Error
}

func (c *ResignationInstructorHRController) getByID(id uint) (*model.RequestResignationInstructor, error) {
	var request model.RequestResignationInstructor
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (c *ResignationInstructorHRController) getByInstructorID(id string) (*model.RequestResignationInstructor, error) {
	var req model.RequestResignationInstructor
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationInstructorHRController) update(req *model.RequestResignationInstructor) error {
	return c.db.Save(req).Error
}

func (c *ResignationInstructorHRController) SubmitResignationInstructor(instructorID string, reason string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {

		instructorController := NewResignationInstructorHRController(tx)

		factory, err := model.GetFactory(1)
		if err != nil {
			return fmt.Errorf("failed to get instructor factory: %v", err)
		}

		reqInterface, err := factory.CreateResignation(instructorID, reason)
		if err != nil {
			return fmt.Errorf("failed to create resignation request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestResignationInstructor)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor resignation request")
		}

		if err := instructorController.insert(req); err != nil {
			return fmt.Errorf("failed to insert resignation request within transaction: %v", err)
		}

		return nil
	})
}

func (c *ResignationInstructorHRController) ReviewInstructorResignRequest(
	tx *gorm.DB,
	requestID, action, reason string,
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
			return tx.Save(r).Error
		},
	)
}
