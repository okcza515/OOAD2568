package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"strconv"

	"gorm.io/gorm"
)

type SupplyController struct {
	db *gorm.DB
	*core.BaseController[model.Supply]

	observers map[string]AssetObserver[model.Supply]
}

type SupplyControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.Supply, error)
	RetrieveByID(id uint, preloads ...string) (model.Supply, error)
	Insert(data model.Supply) error
	UpdateByID(data model.Supply) error
	DeleteByID(id uint) error
	InsertMany(data []model.Supply) error

	addObserver(observer AssetObserver[model.Supply])
	removeObserver(observer AssetObserver[model.Supply])
}

func NewSupplyController() *SupplyController {
	observers := make(map[string]AssetObserver[model.Supply])
	db := migration.GetInstance().DB
	return &SupplyController{
		db:             db,
		BaseController: core.NewBaseController[model.Supply](db),
		observers:      observers,
	}
}

func (c *SupplyController) addObserver(observer AssetObserver[model.Supply]) {
	c.observers[observer.GetObserverID()] = observer
}

func (c *SupplyController) removeObserver(observer AssetObserver[model.Supply]) {
	delete(c.observers, observer.GetObserverID())
}

func (c *SupplyController) notifyAll(eventType model.SupplyLogActionEnum, dataContext model.Supply) {
	for _, observer := range c.observers {
		observer.HandleEvent(eventType.String(), dataContext)
	}
}

func (c *SupplyController) ListAll() ([]string, error) {
	supplies := new([]model.Supply)
	result := c.db.Find(&supplies)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, supply := range *supplies {
		resultList = append(resultList, "["+supply.UpdatedAt.String()+"] "+string(supply.SupplyLabel)+" "+strconv.FormatUint(uint64(supply.ID), 10))
	}

	return resultList, result.Error
}
