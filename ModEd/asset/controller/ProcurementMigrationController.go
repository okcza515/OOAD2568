// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"errors"

	"gorm.io/gorm"
)

type ProcurementMigrationController struct {
	db *gorm.DB
}

func (c *ProcurementMigrationController) migrateToDB() error {
	err := c.db.AutoMigrate(&model.InstrumentRequest{}, &model.InstrumentDetail{}, &model.Category{})

	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
