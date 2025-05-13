package controller

import (
	"ModEd/core"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

type ResignationStudentHRController struct {
	db *gorm.DB
}

func NewResignationStudentHRController(db *gorm.DB) *ResignationStudentHRController {
	db.AutoMigrate(&model.RequestResignationStudent{})
	return &ResignationStudentHRController{db: db}
}

func (c *ResignationStudentHRController) insert(request *model.RequestResignationStudent) error {
	return c.db.Create(request).Error
}

func (c *ResignationStudentHRController) getAll() ([]*model.RequestResignationStudent, error) {
	var requests []*model.RequestResignationStudent
	err := c.db.Find(&requests).Error
	return requests, err
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
		studentController := NewResignationStudentHRController(tx)

		requestFactory := model.RequestFactory{}

		params := model.CreateRequestParams{
			ID:     studentID,
			Reason: reason,
		}

		reqInterface, err := requestFactory.CreateRequest(model.RoleStudent, model.RequestTypeResignation, params)
		if err != nil {
			return fmt.Errorf("failed to create resignation request using factory: %v", err)
		}

		req, ok := reqInterface.(*model.RequestResignationStudent)

		err = req.Validate()
		if err != nil {
			return fmt.Errorf("failed to validate resignation request: %v", err)
		}

		if !ok {
			return fmt.Errorf("factory returned unexpected type for instructor resignation request")
		}

		if err := studentController.insert(req); err != nil {
			return fmt.Errorf("failed to insert resignation request within transaction: %v", err)
		}

		return nil
	})
}

func (c *ResignationStudentHRController) ReviewStudentResignRequest(requestID, action, reason string,
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

func (c *ResignationStudentHRController) ExportStudentResignRequests(filePath string) error {
	resignationStudents, err := c.getAll()
	if err != nil {
		return fmt.Errorf("failed to retrieve resignation requests: %w", err)
	}

	mapper, err := core.CreateMapper[model.RequestResignationStudent](filePath)
	if err != nil {
		return fmt.Errorf("failed to create resignation student mapper: %w", err)
	}

	err = mapper.Serialize(resignationStudents)
	if err != nil {
		return fmt.Errorf("failed to serialize resignation students: %w", err)
	}

	fmt.Printf("Exported %d resignation student requests to %s\n", len(resignationStudents), filePath)

	return nil
}
