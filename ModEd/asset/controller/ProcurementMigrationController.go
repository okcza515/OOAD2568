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
	err := c.db.AutoMigrate(&model.InstrumentRequest{}, &model.InstrumentDetail{}, &model.Category{}, &model.BudgetApproval{}, &model.Procurement{}, &model.Quotation{}, &model.QuotationDetail{},&model.AcceptanceApproval{},&model.AcceptanceCriteria{},&model.Instrument{},)

	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
