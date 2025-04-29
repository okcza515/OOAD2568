package controller

import (
	"ModEd/hr/model"
	"gorm.io/gorm"
	"ModEd/hr/util"
	"fmt"
)

type RaiseHRController struct {
	db *gorm.DB
}

func createRaiseHRController(db *gorm.DB) *RaiseHRController {
	db.AutoMigrate(&model.RequestRaise{})
	return &RaiseHRController{db: db}
}

func (c *RaiseHRController) insert(req *model.RequestRaise) error {
	return c.db.Create(req).Error
}

func (c *RaiseHRController) getByID(id uint) (*model.RequestRaise, error) {
	var raise model.RequestRaise
	err := c.db.First(&raise, id).Error
	return &raise, err
}

func (c *RaiseHRController) updateStatus(id uint, status string) error {
	return c.db.Model(&model.RequestRaise{}).Where("id = ?", id).Update("status", status).Error
}

func (c *RaiseHRController) getAll() ([]model.RequestRaise, error) {
	var requests []model.RequestRaise
	err := c.db.Find(&requests).Error
	return requests, err
}

func (h *HRFacade) SubmitRaiseRequest(db *gorm.DB, instructorID string, amount int, reason string) error {
	tm := &util.TransactionManager{DB: db}

	return tm.Execute(func(tx *gorm.DB) error {
		raiseController := createRaiseHRController(tx)

		request := model.NewRequestRaise(instructorID, reason, amount)

		if err := raiseController.insert(request); err != nil {
			return fmt.Errorf("failed to submit raise request: %v", err)
		}
		return nil
	})
}
