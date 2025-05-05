package controller

import (
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

func (c *RaiseHRController) updateStatus(id uint, status string) error {
	return c.db.Model(&model.RequestRaiseInstructor{}).Where("id = ?", id).Update("status", status).Error
}

func (c *RaiseHRController) getAll() ([]model.RequestRaiseInstructor, error) {
	var requests []model.RequestRaiseInstructor
	err := c.db.Find(&requests).Error
	return requests, err
}
func (c *RaiseHRController) getByInstructorID(instructorID string) ([]model.RequestRaiseInstructor, error) {
	var requests []model.RequestRaiseInstructor
	err := c.db.Where("instructor_code = ?", instructorID).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (c *RaiseHRController) SubmitRaiseRequest(instructorID string, amount int, reason string) error {
	tm := &util.TransactionManager{DB: c.db}

	return tm.Execute(func(tx *gorm.DB) error {
		raiseController := NewRaiseHRController(tx)

		factory, err := model.GetFactory("instructor")
		if err != nil {
			return fmt.Errorf("failed to get factory: %v", err)
		}

		requestObj, err := factory.CreateRaise(instructorID, reason, amount)
		if err != nil {
			return fmt.Errorf("failed to create raise request using factory: %v", err)
		}

		request, ok := requestObj.(*model.RequestRaiseInstructor)
		if !ok {
			return fmt.Errorf("factory returned unexpected type for raise request")
		}

		if err := raiseController.insert(request); err != nil {
			return fmt.Errorf("failed to submit raise request: %v", err)
		}
		return nil
	})
}

func (c *RaiseHRController) ReviewInstructorRaiseRequest(
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

