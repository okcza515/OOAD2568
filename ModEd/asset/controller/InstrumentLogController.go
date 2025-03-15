package controller

import (
	"ModEd/asset/model"
	"errors"
	"gorm.io/gorm"
)

type InstrumentLogController struct {
	Db *gorm.DB
}

func (c *InstrumentLogController) MigrateToDB() error {
	err := c.Db.AutoMigrate(&model.InstrumentLog{}, &model.Instrument{})
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}

func (c *InstrumentLogController) GetAll() (*[]model.InstrumentLog, error) {
	logs := new([]model.InstrumentLog)
	result := c.Db.Find(&logs)

	return logs, result.Error
}
