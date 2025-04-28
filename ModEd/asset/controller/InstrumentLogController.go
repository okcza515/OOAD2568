package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"gorm.io/gorm"
	"strconv"
)

type InstrumentLogController struct {
	db *gorm.DB
	*core.BaseController[model.InstrumentLog]
}

type InstrumentLogControllerInterface interface {
	getAll() ([]model.InstrumentLog, error)
	Insert(data model.InstrumentLog) error
	InsertMany(data []model.InstrumentLog) error
	RetrieveByID(id uint, preloads ...string) (model.InstrumentLog, error)
	List(condition map[string]interface{}, preloads ...string) ([]model.InstrumentLog, error)

	GetObserverID() string
	HandleEvent(eventType string, dataContext model.BorrowInstrument)
}

func NewInstrumentLogController() *InstrumentLogController {
	db := migration.GetInstance().DB
	return &InstrumentLogController{
		db:             db,
		BaseController: core.NewBaseController[model.InstrumentLog](db),
	}
}

func (c *InstrumentLogController) GetObserverID() string {
	return "InstrumentLogController"
}

func (c *InstrumentLogController) HandleEvent(eventType string, dataContext model.BorrowInstrument) {
	switch model.InstrumentLogActionEnum(eventType) {
	case model.INS_LOG_ADDNEW:
		log := model.InstrumentLog{
			InstrumentID: dataContext.Instrument.ID,
			Action:       model.INS_LOG_ADDNEW,
		}

		err := c.Insert(log)
		if err != nil {
			panic(err)
		}
	case model.INS_LOG_UPDATE:
		panic("not implemented")
	case model.INS_LOG_BORROW:
		panic("not implemented")
	case model.INS_LOG_RETURN:
		panic("not implemented")
	case model.INS_LOG_MOVE:
		panic("not implemented")
	case model.INS_LOG_BROKEN:
		panic("not implemented")
	case model.INS_LOG_REPAIR:
		panic("not implemented")
	case model.INS_LOG_LOST:
		panic("not implemented")
	case model.INS_LOG_FOUND:
		panic("not implemented")
	case model.INS_LOG_SALVAGING:
		panic("not implemented")
	case model.INS_LOG_SALVAGE:
		panic("not implemented")
	case model.INS_LOG_DONATE:
		panic("not implemented")
	}
}

func (c *InstrumentLogController) getAll() ([]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.db.Find(&logs)

	return *logs, result.Error
}

func (c *InstrumentLogController) ListAll() ([]string, error) {
	logs := new([]model.InstrumentLog)
	result := c.db.Find(&logs)

	if result.Error != nil {
		return nil, result.Error
	}

	var resultList []string

	for _, log := range *logs {
		resultList = append(resultList, "["+log.UpdatedAt.String()+"] "+string(log.Action)+" "+strconv.FormatUint(uint64(log.InstrumentID), 10))
	}

	return resultList, result.Error
}
