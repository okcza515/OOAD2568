package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"gorm.io/gorm"
)

type SupplyLogController struct {
	db *gorm.DB
	*core.BaseController[model.SupplyLog]
}

type SupplyLogControllerInterface interface {
	getAll() (*[]model.SupplyLog, error)
	Insert(data model.SupplyLog) error
	InsertMany(data []model.SupplyLog) error
	RetrieveByID(id uint, preloads ...string) (model.SupplyLog, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.SupplyLog, error)

	GetObserverID() string
	HandleEvent(eventType string, dataContext model.Supply)
}

func NewSupplyLogController() *SupplyLogController {
	db := migration.GetInstance().DB
	return &SupplyLogController{
		db:             db,
		BaseController: core.NewBaseController[model.SupplyLog](db),
	}
}

func (c *SupplyLogController) GetObserverID() string {
	return "SupplyLogController"
}

func (c *SupplyLogController) HandleEvent(eventType string, dataContext model.Supply) {
	actionEnum, err := model.ToSupplyActionEnum(eventType)
	if err != nil {
		panic(err)
	}

	log := model.SupplyLog{
		SupplyID: dataContext.ID,
		Action:   actionEnum,
	}

	err = c.Insert(log)
	if err != nil {
		panic(err)
	}
}

func (c *SupplyLogController) getAll() (*[]model.SupplyLog, error) {
	suppliesLogs := new([]model.SupplyLog)
	result := c.db.Find(&suppliesLogs)
	return suppliesLogs, result.Error
}
