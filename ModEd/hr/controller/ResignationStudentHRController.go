package controller

import (
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type ResignationStudentHRController struct {
	db *gorm.DB
}

func CreateResignationStudentHRController(db *gorm.DB) *ResignationStudentHRController {
	db.AutoMigrate(&model.RequestResignationStudent{})
	return &ResignationStudentHRController{db: db}
}

func (c *ResignationStudentHRController) insert(request *model.RequestResignationStudent) error {
	return c.db.Create(request).Error
}

func (c *ResignationStudentHRController) getByID(id uint) (*model.RequestResignationStudent, error) {
	var request model.RequestResignationStudent
	err := c.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
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

func (c *ResignationStudentHRController) SubmitResignationStudent(studentID string, reason string) error {
	tm := &util.TransactionManager{DB: c.db}
	return tm.Execute(func(tx *gorm.DB) error {
		studentController := CreateResignationStudentHRController(tx)

		factory, err := model.GetFactory("student")
		if err != nil {
			return fmt.Errorf("failed to get student factory: %v", err)
		}

		reqInterface, err := factory.CreateResignation(studentID, reason)
		if err != nil {
			return fmt.Errorf("failed to create resignation request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestResignationStudent)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor resignation request")
		}

		if err := studentController.insert(req); err != nil {
			return fmt.Errorf("failed to insert resignation request within transaction: %v", err)
		}

		return nil
	})
}

func (c *ResignationStudentHRController) ReviewStudentResignRequest(tx *gorm.DB, requestID string, action string, reason string) error {
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

		if err := tx.Save(&request).Error; err != nil {
			return fmt.Errorf("failed to update resignation request: %v", err)
		}

		return nil
	})
}
