package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"fmt"
	"gorm.io/gorm"
)

type InstrumentController struct {
	db *gorm.DB
	*core.BaseController[model.Instrument]

	observers map[string]AssetObserver[model.BorrowInstrument]
}

type InstrumentControllerInterface interface {
	ListAll() ([]string, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.Instrument, error)
	RetrieveByID(id uint, preloads ...string) (model.Instrument, error)
	Insert(data model.Instrument) error
	UpdateByID(data model.Instrument) error
	DeleteByID(id uint) error
	InsertMany(data []model.Instrument) error

	addObserver(observer AssetObserver[model.BorrowInstrument])
	removeObserver(observer AssetObserver[model.BorrowInstrument])
}

func NewInstrumentController() *InstrumentController {
	observers := make(map[string]AssetObserver[model.BorrowInstrument])
	db := migration.GetInstance().DB
	return &InstrumentController{
		db:             db,
		BaseController: core.NewBaseController[model.Instrument](db),
		observers:      observers,
	}
}

func (c *InstrumentController) addObserver(observer AssetObserver[model.BorrowInstrument]) {
	c.observers[observer.GetObserverID()] = observer
}

func (c *InstrumentController) removeObserver(observer AssetObserver[model.BorrowInstrument]) {
	delete(c.observers, observer.GetObserverID())
}

func (c *InstrumentController) notifyAll(eventType model.InstrumentLogActionEnum, dataContext model.Instrument) {
	for _, observer := range c.observers {
		data := model.BorrowInstrument{Instrument: dataContext}
		observer.HandleEvent(string(eventType), data)
	}
}

func (c *InstrumentController) InsertMany(data []model.Instrument) error {
	err := c.BaseController.InsertMany(data)

	for _, inst := range data {
		c.notifyAll(model.INS_LOG_ADDNEW, inst)
	}

	return err
}

func (c *InstrumentController) ListAll() ([]string, error) {
	instruments := new([]model.Instrument)
	result := c.db.Find(&instruments)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, instrument := range *instruments {
		resultList = append(resultList, fmt.Sprintf("[%v] %v", instrument.InstrumentCode, instrument.InstrumentLabel))
	}

	return resultList, result.Error
}
