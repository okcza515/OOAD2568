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

func CreateResignationInstructorHRController(db *gorm.DB) *ResignationInstructorHRController {
	db.AutoMigrate(&model.RequestResignationInstructor{})
	return &ResignationInstructorHRController{db: db}
}

func (c *ResignationInstructorHRController) insert(request *model.RequestResignationInstructor) error {
	return c.db.Create(request).Error
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

		factory := &model.RequestResignationFactory{}
		req, err := factory.Create("instructor", instructorID, reason)
		if err != nil {
			return fmt.Errorf("failed to build resignation request: %v", err)
		}

		if err := c.insert(req.(*model.RequestResignationInstructor)); err != nil {
			return fmt.Errorf("failed to insert resignation request: %v", err)
		}

		return nil
	})
}
