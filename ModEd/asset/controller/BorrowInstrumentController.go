// MEP-1012 Asset
package controller

import (
	"ModEd/asset/model"
	"ModEd/core"

	"gorm.io/gorm"
)

type BorrowInstrumentController struct {
	db *gorm.DB
	*core.BaseController[model.BorrowInstrument]
}

type BorrowInstrumentControllerInterface interface {
	getAll() ([]model.BorrowInstrument, error)
	GetByID(ID uint) (*model.BorrowInstrument, error)
	//Create(body *model.BorrowInstrument) error
	//Update(ID uint, body *model.BorrowInstrument) error
	//Delete(ID uint) error
}

func (c *BorrowInstrumentController) getAll() ([]model.BorrowInstrument, error) {
	var borrowInstruments []model.BorrowInstrument
	result := c.db.Find(&borrowInstruments)
	return borrowInstruments, result.Error
}

func (c *BorrowInstrumentController) GetByID(ID uint) (*model.BorrowInstrument, error) {
	borrowInstrument := new(model.BorrowInstrument)
	result := c.db.First(&borrowInstrument, "ID = ?", ID)
	return borrowInstrument, result.Error
}

// func (c *BorrowInstrumentController) Create(body *model.BorrowInstrument) error {
// 	result := c.db.Create(body)
// 	return result.Error
// }

// func (c *BorrowInstrumentController) Update(ID uint, body *model.BorrowInstrument) error {
// 	body.ID = ID
// 	result := c.db.Updates(body)
// 	return result.Error
// }

// func (c *BorrowInstrumentController) Delete(ID uint) error {
// 	result := c.db.Model(&model.BorrowInstrument{}).Where("ID = ?", ID).Update("deleted_at", time.Now())
// 	return result.Error
// }
