package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

type ResignationStudentHRController struct {
	db *gorm.DB
}

func createResignationStudentHRController(db *gorm.DB) *ResignationStudentHRController {
	db.AutoMigrate(&model.RequestResignationStudent{})
	return &ResignationStudentHRController{db: db}
}

func (c *ResignationStudentHRController) insert(request *model.RequestResignationStudent) error {
	return c.db.Create(request).Error
}

func (c *ResignationStudentHRController) getByStudentID(id string) (*model.RequestResignationStudent, error) {
	var req model.RequestResignationStudent
	if err := c.db.Where("id = ?", id).First(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (c *ResignationStudentHRController) update(req *model.RequestResignationStudent) error {
	return c.db.Save(req).Error
}

func (h *HRFacade) SubmitResignationStudent(db *gorm.DB,studentID string, reason string) error {
	tm := &util.TransactionManager{DB: db}
	return tm.Execute(func(tx *gorm.DB) error {
		
		controller := createResignationStudentHRController(tx)
		factory := &model.RequestResignationFactory{}
		req, err := factory.Create("student", studentID, reason)
		if err != nil {
			return fmt.Errorf("failed to build resignation request: %v", err)
		}
		if err := controller.insert(req.(*model.RequestResignationStudent)); err != nil {
			return fmt.Errorf("failed to insert resignation request: %v", err)
		}
		return nil
	})
}
